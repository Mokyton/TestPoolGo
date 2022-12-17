package main

import (
	"bufio"
	"cringe.com/mean"
	"cringe.com/median"
	"cringe.com/mode"
	"cringe.com/sd"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const (
	MAX = 100000
	MIN = -100000
)

type Result struct {
	mode   int
	mean   float64
	median float64
	sd     float32
}

func main() {
	flags := flag.String("n", "", "choose what to print")
	flag.Parse()
	var res Result
	seq, err := ReadSequence(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	res.mean = mean.Mean(seq)
	res.median = median.Median(seq)
	res.mode = mode.Mode(seq)
	res.sd = sd.SD(seq)

	switch *flags {
	case "1":
		fmt.Println("Mean:", res.mean)
	case "2":
		fmt.Println("Media:", res.median)
	case "3":
		fmt.Println("Mode:", res.mode)
	case "4":
		fmt.Println("SD:", res.sd)
	default:
		fmt.Printf("Mean: %.2f\n", res.mean)
		fmt.Printf("Median: %.2f\n", res.median)
		fmt.Printf("Mode: %d\n", res.mode)
		fmt.Printf("SD: %.2f\n", res.sd)
	}
}

func ReadSequence(reader io.Reader) ([]int, error) {
	sequence := make([]int, 0)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		text, err := scanner.Text(), scanner.Err()
		if err == io.EOF {
			if len(sequence) == 0 {
				return nil, errors.New("empty string")
			}
			break
		} else if err != nil {
			return nil, err
		}
		num, err := strconv.Atoi(text)
		if err != nil {
			return nil, err
		}
		if num > MAX || num < MIN {
			return nil, errors.New("num out of range")
		}
		sequence = append(sequence, num)
	}
	return sequence, nil
}
