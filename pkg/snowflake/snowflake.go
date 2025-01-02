package snowflake

import (
	"github.com/sony/sonyflake"
	"log"
)

var sf *sonyflake.Sonyflake

// InitSnowflake 初始化雪花算法生成器
func InitSnowflake() {
	stSettings := sonyflake.Settings{}

	var err error
	sf, err = sonyflake.New(stSettings)
	if err != nil {
		log.Fatalf("Failed to create sonyflake: %v", err)
	}

	if sf == nil {
		log.Fatalf("Failed to initialize sonyflake")
	}

	log.Println("Snowflake initialized")
}

// GenerateSnowflakeID 生成雪花 ID
func GenerateSnowflakeID() (uint, error) {
	id, err := sf.NextID()
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
