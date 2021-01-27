package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	caCertFile     string // ca证书
	serverKeyFile  string // 服务端秘钥文件地址
	serverCertFile string // 服务端证书文件地址
)

func main() {
	flag.StringVar(&caCertFile, "ca", "./out/ca.crt", "ca cert file")
	flag.StringVar(&serverKeyFile, "key", "./out/server.key", "server private key file")
	flag.StringVar(&serverCertFile, "cert", "./out/server.crt", "server cert file")
	flag.Parse()
	if caCertFile == "" || serverKeyFile == "" || serverCertFile == "" {
		flag.PrintDefaults()
		return
	}
	s := http.Server{
		Addr: ":8181",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("start")
			_, err := w.Write([]byte("hello you !"))
			if err != nil {
				panic(err)
			}
			fmt.Println("end")
		}),
		TLSConfig: getTLSConfig(),
	}
	fmt.Println(s.ListenAndServeTLS(serverCertFile, serverKeyFile, ))
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
	return &tls.Config{
		ClientCAs:  ca,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
}
