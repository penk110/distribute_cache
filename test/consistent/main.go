package main

import (
	"fmt"
	"stathat.com/c/consistent"
)

type Hash struct {
	consistent *consistent.Consistent
}

func (h *Hash) Add(key string) {
	h.consistent.Add(key)
}

func (h *Hash) Get(key string) (string, error) {
	return h.consistent.Get(key)
}

func (h *Hash) Remove(key string) {
	h.consistent.Remove(key)
}

func NewHash() *Hash {
	consistent := consistent.New()
	h := &Hash{consistent: consistent}
	return h
}

func main() {
	hash := NewHash()
	hash.Add("aaa")
	hash.Add("abc")
	hash.Add("acd")

	fmt.Println(hash.Get("1"))
	fmt.Println(hash.Get("2"))
	fmt.Println(hash.Get("3"))
	fmt.Println(hash.Get("4"))
	fmt.Println(hash.Get("5"))
	fmt.Println(hash.Get("6"))
}
