package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

var (
	line int = 0
)

func main() {
	e := echo.New()
	e.POST("/api/fileupload", fileupload)
	e.Logger.Fatal(e.Start(":1323"))
}

func fileupload(c echo.Context) error {

	// ファイルをアップロード
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

	//  読み込むファイルを開く
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	//  スキャナライブラリを作成
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line++
	}

	//  エラーが有ったかチェック
	if err = scanner.Err(); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "file has uploaded")
}
