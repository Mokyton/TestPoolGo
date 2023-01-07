package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	args, err := getArgs()
	if err != nil {
		log.Fatal(err)
	}
	res, err := Handler(os.Args[1], args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func getArgs() ([]string, error) {
	args := os.Args[2:]
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		args = append(args, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return args, nil
}

func Handler(command string, args []string) (string, error) {
	cmd := exec.Command(command, args...)

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", errors.New(stderr.String())
	}
	return out.String(), nil
}
