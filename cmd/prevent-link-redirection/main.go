package main

import (
	"github.com/Floor-Gang/prevent-link-redirection/internal"
	util "github.com/Floor-Gang/utilpkg"
)

func main() {
	internal.Start()
	util.KeepAlive()
}
