package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	caCertFile     string // ca证书
	clientKeyFile  string // 客户端秘钥文件地址
	clientCertFile string // 客户端证书文件地址
)

func main() {
	flag.StringVar(&caCertFile, "ca", "./out/ca.crt", "ca cert file")
	flag.StringVar(&clientKeyFile, "key", "./out/client.key", "client private key file")
	flag.StringVar(&clientCertFile, "cert", "./out/client.crt", "client cert file")
	flag.Parse()
	if caCertFile == "" || clientKeyFile == "" || clientCertFile == "" {
		flag.PrintDefaults()
		return
	}
	// https
	go func() {
		c := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: getTLSConfig(),
			}, Timeout: 1 * time.Second}
		req, err := http.NewRequest("GET", "https://127.0.0.1:8181", nil)
		if err != nil {
			panic(err)
		}
		for ; ; {
			<-time.After(1 * time.Second)
			fmt.Println("https client start")
			resp, err := c.Do(req)
			if err != nil {
				fmt.Println("https client err", err)
				continue
			}
			bs, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			resp.Body.Close()
			fmt.Println("https client end", string(bs))
		}
	}()
	// https no auth
	go func() {
		c := &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}, Timeout: 1 * time.Second}
		req, err := http.NewRequest("GET", "https://127.0.0.1:8181", nil)
		if err != nil {
			panic(err)
		}
		for ; ; {
			<-time.After(1 * time.Second)
			fmt.Println("https no auth client start")
			resp, err := c.Do(req)
			if err != nil {
				fmt.Println("https no auth client err", err)
				continue
			}
			bs, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			resp.Body.Close()
			fmt.Println("https no auth client end", string(bs))
		}
	}()
	select {}
}

func getTLSConfig() *tls.Config {

	var ca *x509.CertPool
	data, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		panic(err)
	}
	ca = x509.NewCertPool()
	if ok := ca.AppendCertsFromPEM(data); !ok {
		panic(err)
	}
	crt, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		RootCAs:      ca,
		Certificates: []tls.Certificate{crt},
	}
}
