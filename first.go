package main

//Каюков А.И. ИВТ-23, "Написать программу которая выводит текущую дату и время"
import (
	"fmt"
	"time"
)

func first() {
	t := time.Now()
	fmt.Println("\nТекущее время: " + t.Format(time.RFC1123+"\n"))
}
