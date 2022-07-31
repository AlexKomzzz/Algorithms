// A. Андрей и кислота
// https://contest.yandex.ru/contest/28412/problems/?nc=ytaCvoSB

package main

import (
	"fmt"
	"os"
)

func main() {
	var rezerv int
	fmt.Fscan(os.Stdin, &rezerv)

	VolumeSlice := make([]int, rezerv)

	OldValue := -1
	Ok := true

	for n := 0; n < rezerv; n++ {
		var volume int
		fmt.Fscan(os.Stdin, &volume)

		// если новое значение меньше правого, то невозможно налить
		if volume < OldValue {
			Ok = false
		}
		OldValue = volume
		VolumeSlice[n] = volume
	}

	if !Ok {
		fmt.Println(-1)
		return
	} else {
		fmt.Println(VolumeSlice[rezerv-1] - VolumeSlice[0])
	}
}
