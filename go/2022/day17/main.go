package main

import "fmt"

func ShiftRight(message []int, bit int) []int {
	if len(message) > 1 {
		for i, item := range message {
			if item == bit {
				if i == len(message)-1 {
					i = 0
				}
				message[i] = message[i+1]
				message[i+1] = item
				break
			}
		}
	}
	return message
}

func ShiftLeft(message []int, bit int) []int {
	if len(message) > 1 {
		for i, item := range message {
			if item == bit {
				if i == 0 {
					i = len(message) - 1
				}
				message[i] = message[i-1]
				message[i-1] = item
				break
			}
		}
	}
	return message
}

func main() {
	msg := []int{1, 2, 5, 7, 4, 3}
	nMsg := []int{}

	for i := 0; i < len(msg); i++ {
		for j := 0; j < msg[i]; j++ {
			nMsg = ShiftRight(msg, msg[i])
		}
	}

	fmt.Println(nMsg)
}
