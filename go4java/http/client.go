package main

import (
	"bytes"
	"crypto/tls"
	"log"
	"net/http"
)

func callServer(certFile, keyFile []byte, url string) {
	clientSessionCache, buffer := tls.NewLRUClientSessionCache(-1), &bytes.Buffer{}
	cert, _ := tls.X509KeyPair(certFile, keyFile)
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true, ClientSessionCache: clientSessionCache}
	tlsConfig.BuildNameToCertificate()
	client := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}}
	request, _ := http.NewRequest("POST", url, bytes.NewBufferString("payload"))
	resp, _ := client.Do(request)
	buffer.ReadFrom(resp.Body)
	log.Printf("Got status %d and response: %s", resp.StatusCode, buffer.String())
}

func main() {
	
}
