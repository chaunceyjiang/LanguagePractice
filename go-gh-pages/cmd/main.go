package main

import (
	"github.com/labstack/gommon/log"
	"go-gh-pages"
	"os"
	"os/exec"
	"path/filepath"
)
func main() {
	_ = exec.Command("rm", "-rf", "/tmp/.cache").Run()
	base, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	go_gh_pages.Publish(base,nil, func(err error) {
		log.Error(err)
	})

}
