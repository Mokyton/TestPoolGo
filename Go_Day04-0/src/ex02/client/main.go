package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"github.com/lizrice/secure-connections/utils"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Data struct {
	Money      int    `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int    `json:"candyCount"`
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	k := flag.String("k", "", "candy name")
	c := flag.Int("c", 0, "count of candies")
	m := flag.Int("m", 0, "amount of money")
	flag.Parse()

	candy := NewData(*k, *m, *c)

	if !isFlagPassed("k") || !isFlagPassed("c") || !isFlagPassed("m") {
		log.Fatalln("Wrong arguments")
	}
	var data bytes.Buffer
	err := json.NewEncoder(&data).Encode(candy)
	if err != nil {
		log.Fatal(err)
	}

	client := getClietn()
	//client := &http.Client{}
	resp, err := client.Post("https://localhost:3333/buy_candy", "application/json", bytes.NewBuffer(data.Bytes()))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

}

func getClietn() *http.Client {
	data, err := ioutil.ReadFile("../ca/minica.pem")
	if err != nil {
		log.Println(err)
	}
	cp, err := x509.SystemCertPool()
	if err != nil {
		log.Println(err)
	}
	cp.AppendCertsFromPEM(data)

	config := &tls.Config{
		InsecureSkipVerify:    true,
		ClientAuth:            tls.RequireAndVerifyClientCert,
		RootCAs:               cp,
		GetCertificate:        utils.CertReqFunc("../ca/server/cert.pem", "../ca/server/key.pem"),
		VerifyPeerCertificate: utils.CertificateChains,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}
	return client
}

func NewData(name string, money, count int) *Data {
	return &Data{CandyType: name, Money: money, CandyCount: count}
}
