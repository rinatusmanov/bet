package task2_1

import (
	"reflect"
	"testing"
	"time"
)

func TestCreateUnitedChannel(t *testing.T) {
	var (
		intChData    = []int{1, 2, 3, 4, 5, 6, 7, 4, 5234, 6235}
		stringChData = []string{"1", "2", "3", "4"}
		intMap       = make(map[int]int)
		stringMap    = make(map[string]int)
	)
	intCh := make(chan int, 1000)
	stringCh := make(chan string, 1000)
	go func() {
		for _, data := range intChData {
			if _, ok := intMap[data]; !ok {
				intMap[data] = 0
			}
			intMap[data]++
			intCh <- data
		}
		for _, data := range stringChData {
			if _, ok := stringMap[data]; !ok {
				stringMap[data] = 0
			}
			stringMap[data]++
			stringCh <- data
		}
	}()

	ch := CreateUnitedChannel(intCh, stringCh)
ForSwitch:
	for {
		select {
		case outData := <-ch:
			switch reflect.ValueOf(outData).Kind() {
			case reflect.String:
				out := outData.(string)
				if _, ok := stringMap[out]; ok {
					stringMap[out]--
					if stringMap[out] < 0 {
						t.Errorf("не корректный ответ, обнаружены мусорные данные %v", outData)
					}
				} else {
					t.Errorf("не корректный ответ, обнаружены мусорные данные %v", outData)
				}
			case reflect.Int:
				out := outData.(int)
				if _, ok := intMap[out]; ok {
					intMap[out]--
					if intMap[out] < 0 {
						t.Errorf("не корректный ответ, обнаружены мусорные данные %v", outData)
					}
				} else {
					t.Errorf("не корректный ответ, обнаружены мусорные данные %v", outData)
				}
			default:
				t.Errorf("не корректный ответ, не совпадение типов %v", outData)
			}
		case <-time.After(2 * time.Second):
			break ForSwitch
		}
	}
}
