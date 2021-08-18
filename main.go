package main

import (
	"fmt"

	"github.com/iIIusi0n/backkit/network"
)

func main() {
	publicIP, _ := network.GetPublicIP()
	privateIP, _ := network.GetPrivateIP()
	fmt.Println(publicIP, privateIP)
}
