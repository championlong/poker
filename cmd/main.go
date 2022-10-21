package main

import (
	"context"
	"fmt"
	"math/rand"
	"poker/request"
	"poker/utils"
	"strconv"
	"strings"
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
		getRoundInfo()
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
	} else {
		for _, v := range cardInfo.RoundInfos {
			if v.WinnerGroup == "" && utils.InSlice(team, v.GroupNames) {
				if roundNum < v.RoundNum {
					return v.RoundNum
				}
			}
		}
	}
	return getRoundNum()
}

func getRoundInfo() {
	roundInfo, err := request.RoundInfo(context.Background(), roundNum)
	if err != nil {
		fmt.Println("RoundInfo", err)
	}

	if _, ok := roundInfo.PerCardInfo[team]; !ok {
		getRoundInfo()
	}
	proks := roundInfo.PerCardInfo[team]
	prokStr := getRandPork(proks)
	operate, err := request.CardOperate(context.Background(), roundNum, prokStr)
	if err != nil {
		fmt.Println("CardOperate", err)
	}
	fmt.Println("CardOperate result", operate)
}

func getRandPork(porks []int) string {
	tmpArray := rand.Perm(13)
	result := make([]string, 13)
	for i, v := range tmpArray {
		result[i] = strconv.Itoa(porks[v])
	}
	return strings.Join(result, ",")
}
