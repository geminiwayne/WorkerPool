package ProgressBar

import (
	"fmt"
)

const SQUARE = "█"

type Progress struct {
	subTaskNum int
	TaskName   string
	TotalNum   int
}

var progressStr [100]string

// init progress
func (p *Progress) Init(progress Progress) {
	for i := 0; i < 100; i++ {
		progressStr[i] = " "
	}
}

// the UI for the progress bar
func (p *Progress) ProgressBar(progress Progress) {
	percent := 100 * progress.subTaskNum / progress.TotalNum
	var taskStr = ""
	for i := 0; i < percent; i++ {
		progressStr[i] = SQUARE
	}
	for _, item := range progressStr {
		taskStr = taskStr + item
	}
	if percent != 100 {
		fmt.Printf("\r Task: %s Progress: [%s]%d%%", progress.TaskName, taskStr, percent)
	} else if percent == 100 {
		fmt.Printf("\r Task: %s Progress: [%s]%d%% Done!", progress.TaskName, taskStr, percent)
	} else {
		fmt.Printf("\r Task: %s Progress: [%s]%d%% Failed!", progress.TaskName, taskStr, percent)
	}
}

func (p *Progress) Add(num int, progress Progress) {
	progress.subTaskNum = num
	progress.ProgressBar(progress)
}
