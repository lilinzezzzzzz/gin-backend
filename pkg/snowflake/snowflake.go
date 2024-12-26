package snowflake

import (
	"github.com/sony/sonyflake"
	"log"
	"time"
)

var sf *sonyflake.Sonyflake

// InitSnowflake 初始化雪花算法生成器
func InitSnowflake() {
	var st sonyflake.Settings
	st.StartTime = time.Now()
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		log.Fatalf("Failed to initialize sonyflake")
	}
	log.Println("Snowflake initialized")
}

// GenerateSnowflakeID 生成雪花 ID
func GenerateSnowflakeID() uint {
	id, err := sf.NextID()
	if err != nil {
		log.Fatalf("Failed to generate snowflake ID: %v", err)
	}
	return uint(id)
}
