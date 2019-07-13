package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

const versionNumber = 0.1

var version = flag.Bool("v", false, "output version information")
var all = flag.Bool("a", false, "list all items beginning with '.'")
var list = flag.Bool("l", true, ` use a long listing format`)

func usage() {
	fmt.Println("Usage:")
	fmt.Println("List the contents of a directory.")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage

	flag.Parse()
	if *version {
		fmt.Printf("Version Number: %v\n", versionNumber)
		return
	}

	var dir string
	lenArgs := len(os.Args)
	if lenArgs > 2 {
		dir = os.Args[lenArgs-1]
	} else {
		dir = "."
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, f := range files {
		if !*all && strings.HasPrefix(f.Name(), ".") {
			continue
		}

		if *list {
			u, _ := user.LookupId(strconv.FormatInt(int64(f.Sys().(*syscall.Stat_t).Uid), 10))
			group, _ := user.LookupId(strconv.FormatInt(int64(f.Sys().(*syscall.Stat_t).Gid), 10))
			fmt.Println(f.Mode().String(),
				f.Sys().(*syscall.Stat_t).Nlink,
				u.Username,
				group.Username,
				f.Sys().(*syscall.Stat_t).Blksize,
				f.ModTime().Month(),
				f.ModTime().Day(),
				f.ModTime().Hour(),
				f.Size(),
				f.Name())
		} else {
			fmt.Println(f.Name())
		}

	}
}
