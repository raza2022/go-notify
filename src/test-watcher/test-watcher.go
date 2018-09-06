package test_watcher

import "fmt"

func main() {
	hello()
	fmt.Println("fahed")
}

func hello() {
	fmt.Println("am changed why not triggered")
}
