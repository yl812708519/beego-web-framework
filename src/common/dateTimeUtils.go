package common

import "time"

func GetTimeStamp() int64 {
	return int64(time.Now().UnixNano()/1000)
}