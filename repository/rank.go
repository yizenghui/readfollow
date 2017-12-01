package repository

import (
	"math"
	"time"
)

//Rank 获取排名
func Rank(vote int, devote int, timestamp int64) float64 {

	// 赞成与否定差
	voteDiff := vote - devote

	//争议度(赞成/否定)
	var voteDispute float64
	if voteDiff != 0 {
		voteDispute = math.Abs(float64(voteDiff))
	} else {
		voteDispute = 1
	}

	// 项目开始时间 2017-06-01
	projectStartTime, _ := time.Parse("2006-01-02", "2017-09-01")
	fund := projectStartTime.Unix() - 8*3600
	survivalTime := timestamp - fund
	// 如果文章发布时间小于项目创建时间，时差统算为0
	if survivalTime < 0 {
		survivalTime = 0
	}

	// 投票方向与时间造成的系数差
	var timeMagin float64

	// 倍率
	var rate = float64(survivalTime) / float64(86400)

	if voteDiff > 0 {
		timeMagin = rate
	} else if voteDiff < 0 {
		timeMagin = -1 * rate
	} else {
		timeMagin = 0
	}

	vateMagin := math.Log10(voteDispute)

	//详细算法
	socre := vateMagin + float64(timeMagin)
	return socre
}
