package main

import (
	"context"
	"log"
	"runtime"
	"time"

	cpuv3 "github.com/shirou/gopsutil/v3/cpu"
)

func main() {
	cpup := func() {
		v, err := cpuv3.PercentWithContext(context.Background(), 0, false)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(v)
	}

	cpup()

	done := make(chan struct{}, 1)
	quit, cancel := context.WithCancel(context.Background())
	for i := 0; i < runtime.NumCPU()*10000; i++ {
		go func() {
			for {
				select {
				case <-quit.Done():
					done <- struct{}{}
					return
				default:
				}
			}
		}()
	}

	time.Sleep(time.Second * 30)
	cancel()
	<-done

	cpup()
}
