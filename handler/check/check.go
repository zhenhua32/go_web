package check

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

const (
	B  = 1
	KB = B * 1024
	MB = KB * 1024
	GB = MB * 1024
)

func HealthCheck(ctx *gin.Context) {
	message := "OK"
	ctx.String(http.StatusOK, message)
}

func DiskCheck(ctx *gin.Context) {
	usage, _ := disk.Usage("/")

	ctx.JSON(http.StatusOK, usage)
}

func CPUCheck(ctx *gin.Context) {
	counts, _ := cpu.Counts(true)
	precent, _ := cpu.Percent(time.Second*1, false)

	ctx.JSON(http.StatusOK, gin.H{
		"count":   counts,
		"precent": precent,
	})
}

func MemoryCheck(ctx *gin.Context) {
	state, _ := mem.VirtualMemory()

	ctx.JSON(http.StatusOK, state)
}
