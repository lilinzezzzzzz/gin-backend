package snowflake

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

var sf *sonyflake.Sonyflake

// 初始化雪花算法生成器
func init() {
	var st sonyflake.Settings
	st.StartTime = time.Now()
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("Failed to initialize sonyflake")
	}
}

// GenerateSnowflakeID 生成雪花 ID
func GenerateSnowflakeID() uint {
	id, err := sf.NextID()
	if err != nil {
		panic(fmt.Sprintf("Failed to generate snowflake ID: %v", err))
	}
	return uint(id)
}
