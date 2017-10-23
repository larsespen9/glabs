package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

const timeOnly = "15:04:05"
const event string = "2013/07/20, 21:56:13, 252.126.91.56, HDMI_Status: 0"
const event2 string = "2013/07/20, 20:45:13, 252.126.91.56, HDMI_Status: 0"
const timeLayout = "2006/01/02, 15:04:05"
const dateFormat = "2006/01/02"

func main() {
	heleStringen := fjernMellomrom(event)
	heleStringen2 := fjernMellomrom(event2)
	stringListe := strings.Split(heleStringen, ",")
	stringListe2 := strings.Split(heleStringen2, ",")

	tidogdato := stringListe[0] + ", " + stringListe[1]
	tidogdato2 := stringListe2[0] + ", " + stringListe2[1]
	tid2, err := time.Parse(timeLayout, tidogdato2)
	tid, err := time.Parse(timeLayout, tidogdato)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(tid)
	}
	fmt.Println(tid.Date())
	duration := tid.Sub(tid2)
	fmt.Println(duration)
	fmt.Println(tid.Date())
	fmt.Println(tid.Format(dateFormat))
	/* time1, err := time.Parse(timeOnly, stringListe[1])
	if err != nil {
		return
	}
	time2, err2 := time.Parse(timeOnly, stringListe2[1])
	if err2 != nil {
		return
	}
	fmt.Println(time1)
	fmt.Println(time2) */
	//fmt.Println(timeOnly2)
	//fmt.Println(timeOnly)

	/*
		fmt.Println(stringListe[0])
		timeString := stringListe[0] + ", " + stringListe[1]
		fmt.Println(timeString)
		dateString := stringListe[0]
		fmt.Println(dateString)

		ipadr := stringListe[2]
		fmt.Println(ipadr)
		fmt.Println(len(stringListe))

		if len(stringListe) > 4 {
			toChan := stringListe[3]
			fromChan := stringListe[4]
			fmt.Println(toChan)
			fmt.Println(fromChan)
		} else {
			statusChange := stringListe[3]
			status := strings.Split(stringListe[3], ":")
			statustype := status[0]
			value := status[1]
			fmt.Println(statustype)
			fmt.Println(value)
			fmt.Println(statusChange)

		}
		time, err := time.Parse(timeOnly, stringListe[1])
		if err != nil {
			fmt.Errorf(err.Error())
		}
		fmt.Println(time)*/
}

func fjernMellomrom(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
