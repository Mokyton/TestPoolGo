package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	killSig := make(chan os.Signal)
	signal.Notify(killSig, syscall.SIGTERM, os.Interrupt)
	go func() {
		<-killSig
		fmt.Println("kill")
		cancel()
	}()

	in := make(chan string)

	go func(ctx context.Context, in chan string) {
		for {
			id := rand.Intn(100)
			select {
			case <-ctx.Done():
				return
			case in <- "https://api.narutodb.xyz/character/" + strconv.Itoa(id):
			}
		}
	}(ctx, in)

	data := crawlWeb(ctx, in)

	for {
		select {
		case <-ctx.Done():
			return
		case str := <-data:
			fmt.Println(str)
		}
	}

}

func crawlWeb(ctx context.Context, in chan string) chan string {
	out := make(chan string)
	maxGoriutines := make(chan struct{}, 8)
	for i := 0; i < 8; i++ {
		maxGoriutines <- struct{}{}
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case v := <-in:
				<-maxGoriutines
				go func(ctx context.Context, url string, pool chan struct{}) {
					r, _ := http.Get(url)
					defer r.Body.Close()
					body, _ := ioutil.ReadAll(r.Body)
					select {
					case out <- string(body):
					case <-ctx.Done():
						return

					}
					pool <- struct{}{}
				}(ctx, v, maxGoriutines)
			}
		}

	}()
	return out
}
