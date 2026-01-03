package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateEmployeeCode() string {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(90000) + 10000 // 10000 - 99999
	return fmt.Sprintf("EMP-%d", number)
}
