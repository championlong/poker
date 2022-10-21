package main

import (
	"context"
	"fmt"
	"net/http"
	"poker/request"
)

const (
	team = "艾斯奥特曼"
)

var client = &http.Client{}


func main() {
	a,err := request.CardInfo(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}

