package main

import (
	"io/ioutil"
	"os"

	"github.com/yosssi/gohtml"
)

func main() {
	cont, _ := ioutil.ReadAll(os.Stdin)
	result := gohtml.Format(string(cont))
	os.Stdout.WriteString(result)
}
