package main

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
)

func callServer(certFile, keyFile []byte, url string) {
	clientSessionCache := tls.NewLRUClientSessionCache(-1)
	cert, _ := tls.X509KeyPair(certFile, keyFile)
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true, ClientSessionCache: clientSessionCache}
	tlsConfig.BuildNameToCertificate()
	client := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}}
	request, _ := http.NewRequest("POST", url, bytes.NewBufferString("payload"))
	resp, _ := client.Do(request)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	log.Printf("Got status %d and response: %s", resp.StatusCode, string(body))
}

func main() {

}
