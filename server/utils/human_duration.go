package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseDuration(d string) (time.Duration, error) {
	// 去掉首尾的空格
	d = strings.TrimSpace(d)
	// 这个用来解析时间字符串
	dr, err := time.ParseDuration(d)
	fmt.Println(dr)
	if err == nil {
		return dr, nil
	}
	// 检查字符串d 中是否包含子串 "d"
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}
