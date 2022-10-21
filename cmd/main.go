package main

import (
	"context"
	"fmt"
	"math/rand"
	"poker/request"
	"poker/utils"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	team = "艾斯奥特曼"
)

var roundNum int

func main() {
	for {
		applyReport()
		fmt.Println("========报名成功===========")

		roundNum = getRoundNum()
		fmt.Println(roundNum)

		getRoundInfo()

		time.Sleep(28*time.Second)
	}
}

func applyReport(){
	apply, err := request.ApplyRouter(context.Background())
	if err != nil {
		fmt.Println("ApplyRouter", err)
	}
	if apply != "true" && roundNum != 0{
		applyReport()
		return
	}
	return
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
		return
	}
	proks := roundInfo.PerCardInfo[team]
	prokStr := getSortPorkV2(proks)
	operate, err := request.CardOperate(context.Background(), roundNum, prokStr)
	if err != nil {
		fmt.Println("CardOperate", err)
	}
	fmt.Println("============CardOperate result", operate)
}

func getRandPork(porks []int) string {
	tmpArray := rand.Perm(13)
	result := make([]string, 13)
	for i, v := range tmpArray {
		result[i] = strconv.Itoa(porks[v])
	}
	return strings.Join(result, ",")
}

func getSortPork(porks []int) string {
	result := make([]string, 13)
	sort.Ints(porks)
	for i, v := range porks {
		result[i] = strconv.Itoa(v)
	}
	return strings.Join(result, ",")
}


func getSortPorkV2(porks []int) string {
	result := make([]string, 13)
	sort.Ints(porks)
	for i, v := range porks {
		result[i] = strconv.Itoa(v)
	}

	return strings.Join(result[6:11], ",")+","+strings.Join(descSort(result[:6]),",")+","+result[11]+","+result[12]
}


func getSortDescPork(porks []int) string {
	result := make([]string, 13)
	sort.Sort(sort.Reverse(sort.IntSlice(porks)))
	for i, v := range porks {
		result[i] = strconv.Itoa(v)
	}
	return strings.Join(result, ",")
}

func descSort(tmp []string)  []string{

	sort.Strings(tmp)

	ret := make([]string,0)

	for i:=len(tmp)-1;i>=0;i-- {
		ret = append(ret, tmp[i])
	}

	return ret


}
