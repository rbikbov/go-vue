package models

import (
	"fmt"
	//"os"
	"math/rand"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		//os.Exit(0)
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
