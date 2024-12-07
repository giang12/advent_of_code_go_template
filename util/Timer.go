package util

import (
	"log"
	"time"
)

func Timer() func() {
	start := time.Now()
	return func() {
		elapsed := time.Since(start)
		log.Printf("%s took %v", GetCurrentFuncName(), elapsed)
	}
}
