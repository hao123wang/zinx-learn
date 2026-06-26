package main

import "zinx-learn/znet"

type A interface{}

func main() {
	s := znet.NewServer("zinx-learn", "0.0.0.0", 8080)
	s.Start()
}
