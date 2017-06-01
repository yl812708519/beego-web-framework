package common

import "time"


var nanoMil int64 = 1000000



func GetTimeStamp() int64 {
	return int64(time.Now().UnixNano()/ nanoMil)
}

func Time2Date_yyyyMMDDHHMMSS(t int64) string {
	t *= nanoMil
	timer := time.Unix(0,t)

	return timer.Format("2006-01-02 15:04:05")
}

func Time2Date_yyyyMMDD(t int64) string {
	t *= nanoMil
	timer := time.Unix(0,t)

	return timer.Format("2006-01-02")
}