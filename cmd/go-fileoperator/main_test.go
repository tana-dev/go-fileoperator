package main

import "testing"

func TestMain(m *testing.M) {
	println("[test start]")
	m.Run()
	println("[test finish]")
}
