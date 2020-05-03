package main

import "github.com/tana-dev/go-filesplitter/route"

func main() {
	e := route.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
