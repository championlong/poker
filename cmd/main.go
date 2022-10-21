package main

import (
	"context"
	"fmt"
	"poker/request"
)

const (
	team = "艾斯奥特曼"
)

func main() {
	a,err := request.CardInfo(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}

