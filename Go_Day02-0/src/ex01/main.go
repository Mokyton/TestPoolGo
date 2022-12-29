package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"unicode/utf8"
)

type Args struct {
	w     bool
	l     bool
	m     bool
	paths []string
}

func main() {
	cr, err := New()
	if err != nil {
		log.Fatal(err)
	}
	start(cr)
}

func New() (*Args, error) {
	wFlag := flag.Bool("w", false, "for counting words")
	lFlag := flag.Bool("l", false, "for counting lines")
	mFlag := flag.Bool("m", false, "for counting characters")
	flag.Parse()
	return Validation(&Args{*wFlag, *lFlag, *mFlag, os.Args[2:]})
}

func Validation(args *Args) (*Args, error) {
	if (args.w && args.l) || (args.w && args.m) || (args.l && args.m) {
		return nil, errors.New("To many flags you can use only one ")
	}
	if !args.m && !args.w && !args.l {
		args.w = true
		args.paths = os.Args[1:]
	}
	return args, nil
}

func start(args *Args) error {
	var wg sync.WaitGroup
	var lock sync.Mutex
	var output []string
	for i := 0; i < len(args.paths); i++ {
		wg.Add(1)
		if args.m {
			go CountChars(args.paths[i], &wg, &output, &lock)
		} else if args.w {
			go CountWords(args.paths[i], &wg, &output, &lock)
		} else if args.l {
			go CountLines(args.paths[i], &wg, &output, &lock)
		}
	}
	wg.Wait()
	for i := 0; i < len(output); i++ {
		fmt.Println(output[i])
	}
	return nil
}

func CountWords(fileName string, group *sync.WaitGroup, res *[]string, lock *sync.Mutex) {
	defer group.Done()
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	counter := 0
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		counter++
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	lock.Lock()
	*res = append(*res, fmt.Sprintf("%d\t%s", counter, fileName))
	lock.Unlock()
}

func CountLines(fileName string, group *sync.WaitGroup, res *[]string, lock *sync.Mutex) {
	defer group.Done()
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	counter := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		counter++
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	lock.Lock()
	*res = append(*res, fmt.Sprintf("%d\t%s", counter, fileName))
	lock.Unlock()
}

func CountChars(fileName string, group *sync.WaitGroup, res *[]string, lock *sync.Mutex) {
	defer group.Done()
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	counter := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		counter += utf8.RuneCountInString(scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	lock.Lock()
	*res = append(*res, fmt.Sprintf("%d\t%s", counter, fileName))
	lock.Unlock()
}
