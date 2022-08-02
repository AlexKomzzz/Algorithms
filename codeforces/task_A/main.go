// A. Все любят спать
// https://codeforces.com/contest/1714/problem/0
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Alarm struct {
	hour, minute int8
}

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var CountData uint8
	fmt.Fscan(in, &CountData)

	for t := uint8(0); t < CountData; t++ {

		var CountAlarm, Hsleep, Msleep int8
		fmt.Fscan(in, &CountAlarm, &Hsleep, &Msleep)

		AlarmSlice := make([]Alarm, CountAlarm)

		for alarm := int8(0); alarm < CountAlarm; alarm++ {
			fmt.Fscan(in, &AlarmSlice[alarm].hour)
			fmt.Fscan(in, &AlarmSlice[alarm].minute)
		}

		// сортируем будильники
		sort.Slice(AlarmSlice, func(i int, j int) bool {
			if AlarmSlice[i].hour == AlarmSlice[j].hour {
				return AlarmSlice[i].minute < AlarmSlice[j].minute
			}
			return AlarmSlice[i].hour < AlarmSlice[j].hour
		})

		// Найдем будильник с часов большим или равным часу сна
		indHour := sort.Search(int(CountAlarm)-1, func(i int) bool {
			if AlarmSlice[i].hour == Hsleep {
				return AlarmSlice[i].minute >= Msleep
			}
			return AlarmSlice[i].hour >= Hsleep
		})

		var hourRes, minRes int8

		// Проверка на поиск
		if (int8(indHour) < CountAlarm) && (AlarmSlice[indHour].hour == Hsleep) && (AlarmSlice[indHour].minute == Msleep) {
			fmt.Fprintln(out, "0 0")
		} else if (AlarmSlice[indHour].hour == Hsleep && AlarmSlice[indHour].minute > Msleep) || AlarmSlice[indHour].hour > Hsleep {
			// Если мы нашли будильник позже в этих же сутках, то из него вычитаем время сна
			minRes = AlarmSlice[indHour].minute - Msleep
			// если разность минут отрицательная значит считаем так
			if minRes < 0 {
				hourRes = AlarmSlice[indHour].hour - Hsleep - 1
				minRes = 60 + minRes // учитываем, что минуты со знаком минус!!!!
			} else {
				hourRes = AlarmSlice[indHour].hour - Hsleep
			}
			fmt.Fprintf(out, "%d %d\n", hourRes, minRes)

		} else {
			// Если ничего не нашлось, то к первому самому раннему бульнику к 00:00 прибавляем разницу от 24:00 и времени сна
			hourRes = 23 - Hsleep + AlarmSlice[0].hour
			minRes = 60 - Msleep + AlarmSlice[0].minute
			if minRes > 59 {
				minRes -= 60
				hourRes++
			}

			fmt.Fprintf(out, "%d %d\n", hourRes, minRes)
		}
	}

}
