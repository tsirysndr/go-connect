package main

import (
	"fmt"

	"github.com/tsirysndr/go-connect"
)

func main() {
	c := connect.NewConnectionManager()
	err := c.Tethering.DisableTethering("wifi")
	fmt.Println(err)
}
