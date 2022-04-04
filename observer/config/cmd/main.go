package main

import (
	"fmt"
	"patterns/observer/config"
	"time"
)

type Observable interface {
	Add(o config.ConfigObserver)
	Remove(o config.ConfigObserver)
	Notify()
	UpdateCAFile(newFile string) 
}

func main() {
	observable := config.NewConfigManager()
	for i := 1; i <= 4; i++ {
		observable.Add(config.NewConfigObserver(i))
	}
	
	for j := 0; j < 10; j++ {
		observable.UpdateCAFile(fmt.Sprintf("new ca file %v", j))
		observable.Notify()
		time.Sleep(1 * time.Second)
	}
}