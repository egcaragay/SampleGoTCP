package main

import "net"

type Player struct {
	x float32
	y float32
	z float32
	id int
	health int

	conn net.Conn
}
