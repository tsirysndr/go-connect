package main

import (
	"fmt"

	"github.com/tsirysndr/go-connect"
)

func main() {
	c := connect.NewConnectionManager()
	err := c.Tethering.EnableTethering("wifi", "raspberry_wifi", "hotspot123")
	fmt.Println(err)
}
