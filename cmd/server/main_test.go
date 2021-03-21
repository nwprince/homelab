package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("PANIC: %s", r)
			}
		}()
		main()
	}()
}

func Test_getCertFile(t *testing.T) {
	certFile := getCertFile()
	if len(certFile) == 0 {
		t.Error("getCertFile should return a string. Nil value returned")
	}
}

func Test_getKeyFile(t *testing.T) {
	keyFile := getKeyFile()
	if len(keyFile) == 0 {
		t.Error("getKeyFile should return a string. Nil value returned")
	}
}
