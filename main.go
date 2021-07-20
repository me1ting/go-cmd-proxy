package main

import (
	"os"
	"os/exec"
)

var proxy = "http://localhost:1081"
var target = "original-go"

var debug = true

func main() {
	os.Setenv("http_proxy", proxy)
	os.Setenv("https_proxy", proxy)

	cmd := exec.Command(target)
	cmd.Args = os.Args
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}
