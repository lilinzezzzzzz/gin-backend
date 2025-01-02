package main

import (
	"fmt"
	"golang-backend/pkg/snowflake"
	"log"
)

func main() {
	// 初始化 雪花算法
	snowflake.InitSnowflake()
	id, err := snowflake.GenerateSnowflakeID()
	if err != nil {
		log.Fatalf("GenerateSnowflakeID fail, err: %v", err)
	}
	fmt.Println(id)
}
