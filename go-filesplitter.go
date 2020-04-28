package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	line        int = 0
	splitNumber int = 0
	fileLines   int = 0
)

func main() {
	e := echo.New()
	e.POST("/api/fileupload", fileupload)
	e.Logger.Fatal(e.Start(":1323"))
}

func fileupload(c echo.Context) error {

	s := c.QueryParam("splitNumber")
	splitNumber, _ = strconv.Atoi(s)

	//
	// file upload
	//
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	filename := filepath.Join("tmp", file.Filename)
	dst, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	//
	line, err := countLine(filename)
	if err != nil {
		return err
	}

	fmt.Println(line)

	//
	//	fileLines = line / splitNumber
	//	remainderLines = line % splitNumber

	//
	// 既存ファイルオープン
	//
	fp, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fp.Close()

	// 新規ファイルオープン
	newfile := filepath.Join("tmp", "hello.txt")
	newfp, err := os.OpenFile(newfile, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer newfp.Close()

	// 書き込み
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		newfp.WriteString(scanner.Text())
	}

	// 書き込み
	if err = scanner.Err(); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "file has uploaded")
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
