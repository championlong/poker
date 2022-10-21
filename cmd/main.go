package main

import (
	"context"
	"fmt"
	"poker/request"
	"poker/utils"
)

const (
	team = "艾斯奥特曼"
)

var roundNum int

func main() {
	ctx := context.Background()
	apply, err := request.ApplyRouter(ctx)
	if err != nil {
		fmt.Println("ApplyRouter", err)
	}
	if apply == "true" {
		roundNum = getRoundNum()
	}

	fmt.Println(roundNum)
}

func getRoundNum() int {
	cardInfo, err := request.CardInfo(context.Background())
	if err != nil {
		fmt.Println("CardInfo", err)
	}
	if utils.InSlice(team, cardInfo.QueueInfo) {
		return getRoundNum()
	}else {
		for _, v := range cardInfo.RoundInfos {
			if v.WinnerGroup == "" && utils.InSlice(team, v.GroupNames) {
				if roundNum <  v.RoundNum {
					return v.RoundNum
				}
			}
		}
	}
	return getRoundNum()
}
