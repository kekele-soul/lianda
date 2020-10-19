package util

import "time"

/**
* 根据你做需要的格式，生成相应的日期
 */
const TIME_FORMAT_ONE  = "2006年01月02日 15:01:05"
const TIME_FORMAT_TWO  = "2006/01/02 15:01:05"
const TIME_FORMAT_THREE  = "2006-01-02 15:01:05"
func TimeNow(format string) string{
	return time.Now().Format(format)
}
/**
 *将int64整形的时间格式转换成相应的格式
 */
func TimeFormat(sec int64,nsec int64, format string)string{
	return time.Unix(sec,nsec).Format(format)
}