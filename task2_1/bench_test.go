package task2_1

import (
	"testing"
)

var (
	intChSl = func() (result []interface{}) {
		for i := 0; i < 10; i++ {
			result = append(result, make(chan int))
		}
		return
	}()
	stringChSl = func() (result []interface{}) {
		for i := 0; i < 10; i++ {
			result = append(result, make(chan string))
		}
		return
	}()
	interfaceChSl = func() (result []interface{}) {
		for i := 0; i < 10; i++ {
			result = append(result, make(chan interface{}))
		}
		return
	}()
)

func BenchmarkCreateUnitedChannel_intChSl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		close(CreateUnitedChannel(intChSl...))
	}
}
func BenchmarkCreateUnitedChannel_stringChSl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		close(CreateUnitedChannel(stringChSl...))
	}
}

func BenchmarkCreateUnitedChannel_interfaceChSl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		close(CreateUnitedChannel(interfaceChSl...))
	}
}
