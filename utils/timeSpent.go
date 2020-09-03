package utils

import (
	"fmt"
	"time"
)

type MeasureTime interface {
	Start()
	End()
}

type runtime struct {
	funcstartTime time.Time
}

func (r *runtime) Start() {
	r.funcstartTime = time.Now()
}

func (r *runtime) End() {
	fmt.Printf("Time Taken to execute: %v\n", time.Now().Sub(r.funcstartTime))
}

func GetRunTimeCalculator() MeasureTime {
	return &runtime{}
}
