package main

import (
	"fmt"

	"github.com/tsirysndr/go-connect"
)

func main() {
	c := connect.NewConnectionManager()
	// PATH example:
	// /net/connman/service/wifi_b827eb51e6c7_7a786865687541596858695a31797a677379494c_managed_psk
	err := c.Wifi.Remove("<PATH>")
	fmt.Println(err)
}
