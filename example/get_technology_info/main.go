package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-connect"
)

func main() {
	c := connect.NewConnectionManager()
	res, _ := c.GetTechnologyInfo("wifi")
	info, _ := json.Marshal(res)
	fmt.Println(string(info))
}
