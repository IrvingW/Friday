package cron_task

import "fmt"

func Hello() (err error){
	fmt.Printf("哈喽大家好，我是一个定时任务\n")
	return nil
}
