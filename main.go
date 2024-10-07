package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type KubeHaven struct {
	name	string
	version	string
}

func newKubeHaven() KubeHaven {
	kh := KubeHaven{
		name: "KubeHaven",
		version: "v0.01",
	}
	return kh
}

func (kh *KubeHaven) Name() string {
	return kh.name
}

func (kh *KubeHaven) Version() string {
	return kh.version
}

func (kh *KubeHaven) KubernetesHeadRequest() string {
	res, err := http.Get("https://www.google.de/")
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	return strconv.Itoa(res.StatusCode)
}

func main() {
	fmt.Println("asd")
	kh := newKubeHaven()
	fmt.Println(kh.Name())
	fmt.Println(kh.Version())
	fmt.Println(kh.KubernetesHeadRequest())
}