package main

import "fmt"

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

func main() {
	fmt.Println("asd")
	kh := newKubeHaven()
	fmt.Println(kh.Name())
	fmt.Println(kh.Version())
}