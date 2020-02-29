package task2_1

import (
	"fmt"
	"reflect"
	"time"
)

func CreateUnitedChannel(intCh ...interface{}) (ch chan interface{}) {
	var cap int
	for _, channel := range intCh {
		if reflect.TypeOf(channel).Kind() == reflect.Chan {
			capabality := reflect.ValueOf(ch).Cap()
			if capabality == 0 {
				capabality = 1
			}
			cap += capabality
		}
	}
	ch = make(chan interface{}, cap)
	for _, channel := range intCh {
		if reflect.TypeOf(channel).Kind() == reflect.Chan {
			go func(ch interface{}, output chan interface{}) {
				chv := reflect.ValueOf(ch)
				for {
					v, ok := chv.Recv()
					if ok {
						output <- v.Interface()
					} else {
						return
					}
				}
			}(channel, ch)
		}
	}
	return
}

func main() {
	intCh := make(chan int)
	stringCh := make(chan string)
	go func() {
		intCh <- 7
		stringCh <- "3333"
		close(intCh)
	}()
	go func() {
		time.Sleep(2 * time.Hour)
	}()
	ch := CreateUnitedChannel(intCh, stringCh)
	for {
		select {
		case msg := <-ch:
			fmt.Println("received message", msg)
		case <-time.After(time.Hour):
			return
		}

	}
}
