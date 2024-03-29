package main

import (
	"os"
	"os/signal"
	"reflect"
	"strconv"
	"strings"
	"syscall"
)

func AddIntoInt64Arr(arr []int64, val int64) ([]int64, bool) {
	res := []int64{}
	seen := false
	for _, s := range arr {
		res = append(res, s)
		if s == val {
			seen = true
		}
	}
	if !seen {
		res = append(res, val)
	}
	return res, !seen
}

func DelFromInt64Arr(arr []int64, val int64) ([]int64, bool) {
	res := []int64{}
	seen := false
	for _, s := range arr {
		if s != val {
			res = append(res, s)
		} else {
			seen = true
		}
	}
	return res, seen
}

func ParseInt64ArrToStr(arr []int64) string {
	res := []string{}
	for _, s := range arr {
		res = append(res, strconv.FormatInt(s, 10))
	}
	return strings.Join(res, ",")
}

func ParseStrToInt64Arr(str string) []int64 {
	res := []int64{}
	resStr := strings.Split(str, ",")
	for _, s := range resStr {
		i, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			res = append(res, i)
		}
	}

	return res
}

func I64In(arr *[]int64, target int64) bool {
	for _, i := range *arr {
		if i == target {
			return true
		}
	}
	return false
}

func MakeSysChan() chan os.Signal {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	return sigCh
}

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func MaxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Abs(a int64) int64 {
	if a >= 0 {
		return a
	}
	return -a
}

func Type(to interface{}) string {
	if to == nil {
		return "nil"
	}
	return reflect.TypeOf(to).String()
}
