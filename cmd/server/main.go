// Application which greets you.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/nwprince/homelab/cmd/server/panic"
	"github.com/nwprince/homelab/cmd/server/routes"
	"github.com/yarf-framework/yarf"
)

const TLS bool = false

var mblnDevelopment bool = false

func main() {
	y := yarf.New()
	y.NotFound = panic.Exception

	mblnDevelopment = strings.ToLower(os.Getenv("DEVELOPMENT_MACHINE")) == "true"

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
		certFile := getCertFile()
		keyFile := getKeyFile()
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
	if !mblnDevelopment && len(certFile) == 0 {
		log.Fatal("Expected certfile path")
	}
	return certFile
}

func getKeyFile() string {
	keyFile := os.Getenv("TLS_KEY_FILE")
	if !mblnDevelopment && len(keyFile) == 0 {
		log.Fatal("Expected keyFile path")
	}

	return keyFile
}
