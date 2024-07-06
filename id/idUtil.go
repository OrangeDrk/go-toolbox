package idUtil

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

// RandomUUID 随机获取UUID
func RandomUUID() string {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		panic(err)
	}

	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

// SimpleUUID  简化的UUID，去掉了横线
func SimpleUUID() string {
	uuid := RandomUUID()
	return hex.EncodeToString([]byte(uuid))
}

// ObjectID 创建MongoDB ID生成策略实现
func ObjectID() string {
	timestamp := time.Now().Unix()
	return fmt.Sprintf("%x", timestamp)
}

// getSnowflakeInstance 返回单例的 Snowflake 实例
func getSnowflakeInstance() *Snowflake {
	once.Do(func() {
		workId, err := getMacAddress()
		if err != nil {
			workId = 1
		}
		instance = newSnowflake(workId, workId)
	})
	return instance
}

// 初始化一个新的Snowflake实例
func newSnowflake(workerID, datacenterID int64) *Snowflake {
	if workerID > maxWorkerID || workerID < 0 {
		return nil
	}
	if datacenterID > maxDatacenterID || datacenterID < 0 {
		return nil
	}
	return &Snowflake{
		timestamp:    0,
		workerID:     workerID,
		datacenterID: datacenterID,
		sequence:     0,
	}
}

// GetSnowFlakeNextId 获取雪花ID int值
func GetSnowFlakeNextId() int64 {
	s := getSnowflakeInstance()
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixNano() / 1e6
	if now == s.timestamp {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sequence = 0
	}
	s.timestamp = now
	id := ((now - epoch) << timestampLeftShift) |
		(s.datacenterID << datacenterIDShift) |
		(s.workerID << workerIDShift) |
		s.sequence

	return id
}

// GetSnowFlakeNextIdStr 获取雪花ID字符串
func GetSnowFlakeNextIdStr() string {
	id := GetSnowFlakeNextId()
	return fmt.Sprintf("%d", id)
}
