package model

// RoundInfoResp 获取某场比赛信息
type RoundInfoResp struct {
	Id            int              `json:"id"` //场次序号
	StartTime     string           `json:"startTime"` //开始时间
	EndTime       string           `json:"endTime"` //结束时间
	GroupNames    []string         `json:"groupNames"` //参赛小组
	PerCardInfo   map[string][]int `json:"perCardInfo"` //各队伍发牌的信息，key为组名，value为牌
	FinalCardInfo map[string][]int `json:"finalCardInfo"`//各队伍最终牌的信息，key为组名，value为牌
	WinnerGroup   string           `json:"winnerGroup"` //获胜者
}

// RoundInfoResp 获取全部比赛信息
type CardInfoResp struct {
	RoundInfos []RoundInfo `json:"roundInfos"` //比赛场次
	QueueInfo  []string    `json:"queueInfo"`	//排队信息
	GroupRank  []GroupRank `json:"groupRank"` //队伍排名
}

type RoundInfo struct {
	RoundNum    int      `json:"roundNum"` //场次序号
	StartTime   string   `json:"startTime"` //开始时间
	GroupNames  []string `json:"groupNames"` //参赛小组
	WinnerGroup string   `json:"winnerGroup"` //获胜者
}

type GroupRank struct {
	GroupName string `json:"groupName"` //小组名
	Score     int    `json:"score"` //分数
}
