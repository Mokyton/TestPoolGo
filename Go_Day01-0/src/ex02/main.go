package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	oldFile := flag.String("old", "txtFiles/snapshot1.txt", "path to old database")
	newFile := flag.String("new", "txtFiles/snapshot2.txt", "path to new database")
	flag.Parse()
	o1, err := exec.Command("sort", *oldFile).CombinedOutput()
	o2, err := exec.Command("sort", *newFile).CombinedOutput()
	f, _ := os.Create("tmp1")
	defer f.Close()
	f.Write(o1)
	d, _ := os.Create("tmp2")
	defer d.Close()
	d.Write(o2)
	defer os.Remove(f.Name())
	defer os.Remove(d.Name())
	output, err := exec.Command("diff", f.Name(), d.Name()).CombinedOutput()
	//cmd := exec.Command("diff", "snapshot1.txt", "snapshot2.txt")
	//stdout, err := cmd.Output()
	storage := strings.Split(string(output), "\n")
	for i := 0; i < len(storage); i++ {
		fmt.Println(storage[i])
	}
	if err != nil {
		log.Fatal(err)
	}

	//
	//// Print the output
	//fmt.Println(string(stdout))
}

//<(sort File1.txt)
