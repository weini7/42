package utils

import (
	"fmt"
	"strconv"
)

//伪解决 73.10 * float64(100) = 7309.999999999999   的问题
//伪解决 73.10 * float64(100.0) = 7309              的问题
func IntegerFloat64ToInt64(value float64) int64 {
	jie, _ := strconv.ParseFloat(fmt.Sprintf("%.f", value), 64)
	return int64(jie)
}
