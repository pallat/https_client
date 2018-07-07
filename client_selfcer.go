package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	certFile = flag.String("im", "intermediate CA", "A PEM eoncoded certificate file.")
)

func main() {
	flag.Parse()

	cert, err := ioutil.ReadFile(*certFile)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(cert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	url := "https://my.internal.com/api/login"

	payload := strings.NewReader("{\n  \"username\":\"myname\",\n  \"password\":\"mypassword\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
