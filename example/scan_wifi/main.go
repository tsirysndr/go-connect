package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-connect"
)

func main() {
	c := connect.NewConnectionManager()
	c.Wifi.Scan()
	res, _ := c.Wifi.GetNetworks()
	r, _ := json.Marshal(res)
	fmt.Println(string(r))
}
