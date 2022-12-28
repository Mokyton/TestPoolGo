package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Args struct {
	dFlag  bool
	fFlag  bool
	slFlag bool
	ext    string
	path   string
}

func main() {
	opt, err := New()
	if err != nil {
		log.Fatal(err)
	}
	err = MyFind(&opt)
	if err != nil {
		log.Fatal(err)
	}
}

func New() (Args, error) {
	dPtr := flag.Bool("d", false, "only dir mode")
	slPtr := flag.Bool("sl", false, "only symlinks mode")
	fPtr := flag.Bool("f", false, "only files mode")
	extPtr := flag.String("ext", "", "special extension mode")
	flag.Parse()
	if !*fPtr && *extPtr != "" {
		return Args{}, errors.New("Error: flag ext used in a wrong condition !")
	}
	return Args{*dPtr, *fPtr, *slPtr, *extPtr, os.Args[len(os.Args)-1]}, nil
}

func MyFind(opt *Args) error {
	err := filepath.Walk(opt.path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Mode().Type() == os.ModeDir && opt.dFlag && os.ModePerm != 0 {
				fmt.Println(path)
			} else if info.Mode().Type() == os.ModeSymlink && opt.slFlag && os.ModePerm != 0 {
				location, err := os.Readlink(info.Name())
				if err != nil {
					location = "[broken]"
				}
				fmt.Println(path, "->", location)
			} else if os.ModePerm != 0 && opt.fFlag && opt.ext != "" && !info.IsDir() && info.Mode().Type() != os.ModeSymlink {
				if strings.HasSuffix(path, "."+opt.ext) {
					fmt.Println(path)
				}
			} else if os.ModePerm != 0 && !info.IsDir() && info.Mode().Type() != os.ModeSymlink && opt.fFlag {
				fmt.Println(path)
			}
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}
