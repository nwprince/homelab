// Application which greets you.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nwprince/homelab/cmd/server/panic"
	"github.com/nwprince/homelab/cmd/server/routes"
	"github.com/yarf-framework/yarf"
)

const TLS bool = false

func main() {
	y := yarf.New()
	y.NotFound = panic.Exception

	certFile := getCertFile()

	keyFile := getKeyFile()

	y.Add("/", new(routes.Portfolio))

	fmt.Println("starting server")
	s := &http.Server{
		Addr:           ":8080",
		Handler:        y,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	var err error
	if TLS {
		err = s.ListenAndServeTLS(certFile, keyFile)
	} else {
		err = s.ListenAndServe()
	}
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Ending program")
}

func getCertFile() string {
	certFile := os.Getenv("TLS_CERT_FILE")
	if len(certFile) == 0 {
		log.Fatal("Expected certfile path")
	}
	return certFile
}

func getKeyFile() string {
	keyFile := os.Getenv("TLS_KEY_FILE")
	if len(keyFile) == 0 {
		log.Fatal("Expected keyFile path")
	}

	return keyFile
}
