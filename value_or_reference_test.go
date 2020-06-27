package go_study

import "testing"

// array不变
// slice map变
func TestArray(t *testing.T) {
	var arr [3]int = [3]int{1, 2, 3}
	t.Log(arr)
	changeArray(arr)
	t.Log(arr)
}

func changeArray(arr [3]int) {
	arr[0] = 10
}

func TestSlice(t *testing.T) {
	var slice []int = []int{1, 2, 3}
	t.Log(slice)
	changeSlice(slice)
	t.Log(slice)
}

func changeSlice(slice []int) {
	slice[0] = 10
}

func TestMap(t *testing.T) {
	var m map[string]string = map[string]string{"1": "11", "2": "22"}
	t.Log(m)
	changeMap(m)
	t.Log(m)
}

func changeMap(m map[string]string) {
	m["1"] = "33"
}

func TestChan(t *testing.T) {
	var channel chan int = make(chan int,10)
	channel <- 2
	t.Log(channel)
	changeChan(channel)
	t.Log(channel)
}

func changeChan(channel chan int) {

}
