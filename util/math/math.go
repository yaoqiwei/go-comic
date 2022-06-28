package math

import (
	"fmt"
	m "math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// 小数点后 n 位 - 四舍五入
func RoundedFixed(val float64, n int) float64 {
	shift := m.Pow(10, float64(n))
	fv := 0.0000000001 + val //对浮点数产生.xxx999999999 计算不准进行处理
	return m.Floor(fv*shift+.5) / shift
}

// 小数点后 n 位 - 舍去
func TruncRound(val float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n+1)+"f", val)
	temp := strings.Split(floatStr, ".")
	var newFloat string
	if len(temp) < 2 || n >= len(temp[1]) {
		newFloat = floatStr
	} else {
		newFloat = temp[0] + "." + temp[1][:n]
	}
	inst, _ := strconv.ParseFloat(newFloat, 64)
	return inst
}

func GetRandomInt(l int) string {

	rand.Seed(time.Now().UnixNano())
	result := make([]byte, l)

	for i := 0; i < l; i++ {
		result[i] = byte(48 + rand.Intn(10))
	}
	return string(result)
}

func GetRandomStr(l int) string {

	rand.Seed(time.Now().UnixNano())
	result := make([]byte, l)

	for i := 0; i < l; i++ {
		rand := rand.Intn(36)
		if rand < 10 {
			result[i] = byte(48 + rand)
		} else {
			result[i] = byte(55 + rand)
		}

	}
	return string(result)
}

func GetRandomStri(l int) string {

	rand.Seed(time.Now().UnixNano())
	result := make([]byte, l)

	for i := 0; i < l; i++ {
		rand := rand.Intn(62)
		if rand < 10 {
			result[i] = byte(48 + rand)
		} else if rand < 36 {
			result[i] = byte(55 + rand)
		} else {
			result[i] = byte(61 + rand)
		}
	}
	return string(result)
}

func Rand(min, max int) int {
	if min > max {
		panic("min: min cannot be greater than max")
	}
	if int31 := 1<<31 - 1; max > int31 {
		panic("max: max can not be greater than " + strconv.Itoa(int31))
	}
	if min == max {
		return min
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max+1-min) + min
}
