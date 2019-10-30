package check

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

var hostname string

func init() {
	name, err := os.Hostname()
	if err != nil {
		name = "unknow"
	}
	hostname = name
}

// HealthCheck 返回心跳响应
func HealthCheck(ctx *gin.Context) {
	message := fmt.Sprintf("OK from %s", hostname)
	ctx.String(http.StatusOK, message)
}

// DiskCheck 返回磁盘信息
func DiskCheck(ctx *gin.Context) {
	usage, _ := disk.Usage("/")

	ctx.JSON(http.StatusOK, usage)
}

// CPUCheck 返回 CPU 信息
func CPUCheck(ctx *gin.Context) {
	counts, _ := cpu.Counts(true)
	precent, _ := cpu.Percent(time.Second*1, false)

	ctx.JSON(http.StatusOK, gin.H{
		"count":   counts,
		"precent": precent,
	})
}

// MemoryCheck 返回内存信息
func MemoryCheck(ctx *gin.Context) {
	state, _ := mem.VirtualMemory()

	ctx.JSON(http.StatusOK, state)
}
