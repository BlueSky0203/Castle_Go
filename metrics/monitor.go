package metrics

import (
	"context"
	"encoding/json"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"

	"Castle_Go/utils"
)

type SystemStats struct {
	CPU       float64 `json:"cpu"`
	Memory    float64 `json:"memory"`
	DiskUsed  float64 `json:"disk_used"`
	NetRecvMB float64 `json:"net_recv_mbps"`
	NetSentMB float64 `json:"net_sent_mbps"`
}

func StartMonitoringPublisher() {
	var prevRecv, prevSent uint64
	var prevTime time.Time

	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		cpuPercent, _ := cpu.Percent(0, false)
		memStats, _ := mem.VirtualMemory()
		diskStats, _ := disk.Usage("/") // 根目錄磁碟使用率

		netIOs, _ := net.IOCounters(false) // 只取總計
		var recvSpeed, sentSpeed float64
		now := time.Now()

		if len(netIOs) > 0 {
			recv := netIOs[0].BytesRecv
			sent := netIOs[0].BytesSent

			if !prevTime.IsZero() {
				interval := now.Sub(prevTime).Seconds()
				recvSpeed = float64(recv-prevRecv) / interval / (1024 * 1024) // MB/s
				sentSpeed = float64(sent-prevSent) / interval / (1024 * 1024) // MB/s
			}

			prevRecv = recv
			prevSent = sent
			prevTime = now
		}

		stats := SystemStats{
			CPU:       cpuPercent[0],
			Memory:    memStats.UsedPercent,
			DiskUsed:  diskStats.UsedPercent,
			NetRecvMB: recvSpeed,
			NetSentMB: sentSpeed,
		}

		data, _ := json.Marshal(stats)
		utils.RedisClient.Publish(context.Background(), "metrics_channel", data)
	}
}
