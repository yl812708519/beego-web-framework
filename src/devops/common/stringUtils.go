package common

import "github.com/satori/go.uuid"

/*
Version 1, based on timestamp and MAC address (RFC 4122)
Version 2, based on timestamp, MAC address and POSIX UID/GID (DCE 1.1)
Version 3, based on MD5 hashing (RFC 4122)
Version 4, based on random numbers (RFC 4122)
Version 5, based on SHA-1 hashing (RFC 4122)
version 1  只有最开始几个字符不一样，timestamp精度不够。大概一秒内的请求会产生重复的uuid
 */
func GenUUID() string {
	return uuid.NewV3(uuid.NewV1(), "uuid").String()
}