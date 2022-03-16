package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Car struct {
	Id      string `json:"car_id"` //json tag
	Mileage int    `json:"car_mileage"`
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func main() {
	//Fill up the request
	id := "B3QR80"
	mileAge := 50002
	car := Car{id, mileAge}
	//
	bytesData, err := json.Marshal(car)
	fmt.Println(bytesData)
	checkErr(err)
	reader := bytes.NewReader(bytesData)
	url := "https://localhost:8080/put"
	request, err := http.NewRequest("PUT", url, reader)
	checkErr(err)
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")

	pool := x509.NewCertPool()
	caCertPath := "./ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)
	cliCrt, err := tls.LoadX509KeyPair("./client.crt", "./client.key")
	checkErr(err)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(request)
	checkErr(err)
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
