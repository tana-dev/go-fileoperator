package api

import (
	"archive/zip"
	"bufio"
	"bytes"
	crand "crypto/rand"
	"encoding/base32"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

var (
	splitLine int = 0
	splitFile int = 0
	fileLines int = 0
	sessionID string
)

func PostFilesplit(c echo.Context) error {

	s := c.QueryParam("splitNumber")
	splitFile, _ = strconv.Atoi(s)
	sessionID, err := getSessionID()
	if err != nil {
		return err
	}

	targetDir := filepath.Join("tmp", sessionID)
	if err := os.Mkdir(targetDir, 0777); err != nil {
		return err
	}
	defer os.RemoveAll(targetDir)

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	filename := filepath.Join("tmp", sessionID, file.Filename)
	dst, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	fileLines, err := countLine(filename)
	if err != nil {
		return err
	}

	splitLine = fileLines / splitFile
	splitFiles, err := splitRowFile(filename)
	if err != nil {
		return err
	}

	w := c.Response()
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=split.zip")

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	for _, s := range splitFiles {
		if err := addToZip(s, zipWriter); err != nil {
			return err
		}
	}

	return nil
}

func countLine(file string) (int, error) {

	var (
		line int
		err  error
	)

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return 0, err
	}

	line = bytes.Count(b, []byte{'\n'})

	return line, err
}

func splitRowFile(filename string) ([]string, error) {

	var l int = 1
	var fileNumber int = 0
	var newFiles []string

	newFilesOpener := []*os.File{}
	for i := 1; i <= splitFile; i++ {
		newfile := filename + "." + strconv.Itoa(i)
		newfp, err := os.OpenFile(newfile, os.O_RDWR|os.O_CREATE, 0664)
		if err != nil {
			return newFiles, err
		}
		defer newfp.Close()
		newFilesOpener = append(newFilesOpener, newfp)
		newFiles = append(newFiles, newfile)
	}

	fp, err := os.Open(filename)
	if err != nil {
		return newFiles, err
	}
	defer fp.Close()

	r := bufio.NewReader(fp)
	for {

		if l == splitLine && fileNumber != splitFile-1 {
			fileNumber++
			l = 1
		}

		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return newFiles, err
		}

		newFilesOpener[fileNumber].WriteString(line)

		l++
	}

	return newFiles, err
}

func getSessionID() (string, error) {

	b := make([]byte, 32)
	_, err := io.ReadFull(crand.Reader, b)
	if err != nil {
		return "", err
	}
	id := strings.TrimRight(base32.StdEncoding.EncodeToString(b), "=")

	return id, nil
}

func readfile(srcpath string) []byte {

	src, err := os.Open(srcpath)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	contents, _ := ioutil.ReadAll(src)

	return contents
}

func addToZip(filename string, zipWriter *zip.Writer) error {
	src, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer src.Close()

	writer, err := zipWriter.Create(filepath.Base(filename))
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, src)
	if err != nil {
		return err
	}

	return nil
}
