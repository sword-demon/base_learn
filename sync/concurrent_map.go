package main

import (
	"fmt"
	"github.com/orcaman/concurrent-map"
	"log"
	"time"
)

func main() {
	m := cmap.New()

	//m.Set("foo", "bar")
	//
	//if tmp, ok := m.Get("foo"); ok {
	//	bar := tmp.(string)
	//	fmt.Println(bar)
	//}
	//
	//m.Remove("foo")

	go func() {
		for i := 0; i < 10000; i++ {
			key := fmt.Sprintf("key_%d", i)
			m.Set(key, i)
		}
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			key := fmt.Sprintf("key_%d", i)
			v, exists := m.Get(key)
			if exists {
				log.Printf("[%s=%v]", key, v)
			}
		}
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			key := fmt.Sprintf("key_%d", i)
			m.Set(key, i)
		}
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			key := fmt.Sprintf("key_%d", i)
			m.Set(key, i)
		}
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			key := fmt.Sprintf("key_%d", i)
			m.Set(key, i)
		}
	}()

	go func() {
		for i := 0; i < 10000; i++ {
			key := fmt.Sprintf("key_%d", i)
			m.Set(key, i)
		}
	}()

	time.Sleep(1 * time.Hour)
}
