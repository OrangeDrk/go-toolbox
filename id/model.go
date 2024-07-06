package idUtil

import (
	"crypto/md5"
	"fmt"
	"net"
	"sync"
)

// 雪花算法单例
var (
	once     sync.Once
	instance *Snowflake
	//workerId int64 = 1 // 设定你的 worker ID
)

// Snowflake Snowflake结构体
type Snowflake struct {
	mu           sync.Mutex
	timestamp    int64
	workerID     int64
	datacenterID int64
	sequence     int64
}

const (
	workerIDBits     = 5  // 机器ID所占的位数
	datacenterIDBits = 5  // 数据中心ID所占的位数
	sequenceBits     = 12 // 序列号所占的位数

	maxWorkerID     = -1 ^ (-1 << workerIDBits)
	maxDatacenterID = -1 ^ (-1 << datacenterIDBits)
	maxSequence     = -1 ^ (-1 << sequenceBits)

	workerIDShift      = sequenceBits
	datacenterIDShift  = sequenceBits + workerIDBits
	timestampLeftShift = sequenceBits + workerIDBits + datacenterIDBits
	epoch              = int64(1288834974657) // 开始时间戳 (2010-11-04T01:42:54.657Z)
)

// 获取本地 MAC 地址
func getMacAddress() (int64, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return 0, err
	}

	for _, inter := range interfaces {
		if len(inter.HardwareAddr) > 0 {
			hash := md5.Sum([]byte(inter.HardwareAddr))
			macInt := int64(hash[0]) | int64(hash[1])<<8 | int64(hash[2])<<16 | int64(hash[3])<<24
			macInt = macInt & maxWorkerID // 确保 workerId 在允许的范围内
			return macInt, nil
		}
	}

	return 0, fmt.Errorf("no MAC address found")
}
