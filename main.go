package main

import (
	"fmt"

	"./AddressRecover"
	"./Logger"
)

func main() {
	defer fmt.Println("Operation completed. Good bye!")
	var ip string = AddressRecover.GetAddress()
	Logger.Info(ip)
}
