package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {

	var idc string
	var iType int
	var checkSum int

	iType = rand.Intn(7) + 1
	checkSum += iType * 13
	idc = strconv.Itoa(iType)
	for i := 12; i > 1; i-- {
		iDigit := rand.Intn(9)
		checkSum += iDigit * i
		idc += strconv.Itoa(iDigit)
	}

	checkSum = (checkSum % 11)
	if checkSum != 0 {
		checkSum = 11 - checkSum
	}
	idc += strconv.Itoa(checkSum)
	fmt.Println(idc)
}
