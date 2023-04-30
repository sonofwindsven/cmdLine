package cmd

import (
	"cmdLine/internal"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := internal.GetNowTime()
		log.Printf("输出结果：%s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTime time.Time
		var layout = "2006-01-02 15:04:05"
		
		if calculateTime == "" {
			currentTime = time.Now()
		} else {
			var err error
			if !strings.Contains(calculateTime, " ") {
				layout = "2006-01-02"
			}
			currentTime, err = time.ParseInLocation(layout, calculateTime, internal.Tz)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTime = time.Unix(int64(t), 0)
			}
		}
		calculateTime, err := internal.GetCalculateTime(currentTime, duration)
		if err != nil {
			log.Fatalf("time.GetCalculateTime err: %v", err)
		}
		log.Printf("输出结果：%s, %d", calculateTime.In(internal.Tz).Format(layout), calculateTime.Unix())
	},
}

var calculateTime, duration string

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算的时间，有效单位为时间戳或已经格式化后的时间")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效单位为"ns"、"us"(or "us")、"ms", "s"、"m"、"h"`)
}
