package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
	"unicode"
)

func main() {

	buf, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	rBuffer := []rune(buf)
	//var trune []rune
	/* s1 := string(rBuffer[0:19])
	s2 := string(rBuffer[20:39])  */

	//timeString := strings.ReplaceAll(buf, " ", "")

	//
	count := 0
	var Min, Sec []rune
	for i := 0; i < len(rBuffer); i++ {
		if unicode.IsDigit(rBuffer[i]) && count == 0 {
			Min = append(Min, rBuffer[i])
		}
		if !unicode.IsDigit(rBuffer[i]) {
			count++
		}
		if unicode.IsDigit(rBuffer[i]) && count > 0 {
			Sec = append(Sec, rBuffer[i])
		}

	}
	//fmt.Println(string(Min), " min, ", string(Sec), " sec")
	iMin, _ := strconv.Atoi(string(Min))
	iSec, _ := strconv.Atoi(string(Sec))

	//sectime:=time.Unix(int64(iMin*60+iSec),0)

	const UnixDate = "Mon Jan _2 15:04:05 MST 2006"
	firstTime := time.Unix(1589570165+int64(iMin*60+iSec), 0)
	fmt.Println(firstTime.Format("Mon Jan _2 15:04:05 MST 2006"))
	//fmt.Println(rBuffer[0:19])
	//fmt.Println(rBuffer[20:39])
	//     13.03.2018 14:00:15,12.03.2018 14:00:15

	//13.03.2018 14:00:15,12.03.2018 14:00:15
	//02.01.2006 15:04:05

	//fmt.Println(s)
	//s := buf

	//firstTime, _ := time.Parse("02.01.2006 15:04:05", s1)
	//secondTime, _ := time.Parse("02.01.2006 15:04:05", s2)

	//fmt.Println(firstTime.Format("02.01.2006 15:04:05"))
	//fmt.Println(secondTime.Format("02.01.2006 15:04:05"))

	//13.03.2018 14:00:15,12.03.2018 14:00:15
	//12.03.2018 14:00:15,13.03.2018 14:00:15

	/*
		res := -time.Since(firstTime) + time.Since(secondTime)
		if res > 0 {
			fmt.Println(res.Round(time.Second))
		} else {
			fmt.Println(-1 * res.Round(time.Second))
		}
	*/

	//fmt.Println(res.Round(time.Second))
	//fmt.Println(s1)
	//fmt.Println(s2)

	/*
		if firstTime.Hour() > 13 {
			nextime := firstTime.AddDate(0, 0, 1)

		} else {
			fmt.Println(firstTime.Format("2006-01-02 15:04:05"))
		}
	*/
}
