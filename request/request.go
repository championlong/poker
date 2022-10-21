package request

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"poker/model"
	"strconv"
	"time"
)

const (
	urlDomain = "https://1024.hellogroup.com"
	UserToken = "7f4afd39-727e-4d4a-9e60-e76d79ac4d78"
)

var client = &http.Client{}

// ApplyRouter 报名接口
func ApplyRouter(ctx context.Context) (string, error) {
	url := urlDomain + "/card/ready?userToken=" + UserToken
	request, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		respBody, _ := ioutil.ReadAll(resp.Body) //true/flase
		return string(respBody), nil
	}
	return "", errors.New("err apply router")
}

// CardInfo 获取整体⽐赛信息接口
func CardInfo(ctx context.Context) (model.CardInfoResp, error) {
	cardInfoResp := model.CardInfoResp{}
	url := urlDomain + "/card/info?userToken=" + UserToken + "&startTime=" +
		strconv.FormatInt(time.Now().Add(-time.Hour).Unix() * 1000, 10) + "&endTime=" + strconv.FormatInt(time.Now().Unix() * 1000, 10)
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := client.Do(request)
	if err != nil {
		return cardInfoResp, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(respBody, &cardInfoResp)
		if err != nil {
			return cardInfoResp, err
		}
		return cardInfoResp, nil
	}
	return cardInfoResp, errors.New("err cardInfo")
}

// RoundInfo 获取某场比赛信息接口
func RoundInfo(ctx context.Context, roundNum int) (model.RoundInfoResp, error) {
	roundResp := model.RoundInfoResp{}
	url := urlDomain + "/card/roundInfo?userToken=" + UserToken + "&roundNum=" + strconv.Itoa(roundNum)
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := client.Do(request)
	if err != nil {
		return roundResp, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(respBody, &roundResp)
		if err != nil {
			return roundResp, err
		}
		return roundResp, nil
	}
	return roundResp, errors.New("err roundInfo")
}

// CardOperate 修改某场比赛出牌顺序
func CardOperate(ctx context.Context, roundNum int, cardsInfo string) (string, error) {
	url := urlDomain + "/card/operate?userToken=" + UserToken + "&roundNum=" +
		strconv.Itoa(roundNum) + "&cardsInfo=" + cardsInfo
	request, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		respBody, _ := ioutil.ReadAll(resp.Body) //ture/false
		return string(respBody), nil
	}
	return "", errors.New("err cardOperate")
}
