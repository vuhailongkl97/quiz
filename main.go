package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func quiz1() {
	file, err := os.Open("problems.csv")

	if err != nil {
		fmt.Printf("error %v\n", err)
		log.Fatal(err)
	}
	defer file.Close()
	rder := csv.NewReader(file)
	record, err := rder.ReadAll()
	for err != nil {
		log.Fatal(err)
	}
	numberCorrectedAnswer := 0
	for _, i := range record {
		if len(i) != 2 {
			fmt.Printf("error (%v) len %v != 2\n", i, len(i))
			break
		}

		iisplit := strings.Split(i[0], "+")
		if len(iisplit) != 2 {
			fmt.Printf("error (%v) len %v != 2\n", iisplit, len(iisplit))
			break
		}

		num1, err := strconv.ParseInt(iisplit[0], 10, 32)

		if err != nil {
			log.Fatal(err)
			continue
		}
		num2, err := strconv.ParseInt(iisplit[1], 10, 32)

		if err != nil {
			log.Fatal(err)
			continue
		}

		//fmt.Printf("Num1 %v , num2 %v,expected sum = %v, reality %v \n", num1, num2, num1+num2, i[1])
		sum, err := strconv.ParseInt(i[1], 10, 32)

		if err != nil {
			log.Fatal(err)
			continue
		}
		if num1+num2 == sum {
			numberCorrectedAnswer++
		}
		fmt.Println()
	}
	fmt.Printf("correct %v/%v\n", numberCorrectedAnswer, len(record))

}

func quiz2() {
	file, err := os.Open("problems.csv")

	if err != nil {
		fmt.Printf("error %v\n", err)
		log.Fatal(err)
	}
	defer file.Close()
	rder := csv.NewReader(file)
	record, err := rder.ReadAll()
	for err != nil {
		log.Fatal(err)
	}
	user := bufio.NewReader(os.Stdin)

	res := make(chan int32)
	go func() {
		for {
			text, _ := user.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			your_ret, err := strconv.ParseInt(text, 10, 32)

			if err != nil {
				fmt.Print("next one : ")
				res <- 0
				continue
			}
			res <- int32(your_ret)
		}
	}()
	for _, i := range record {
		if len(i) != 2 {
			fmt.Printf("error (%v) len %v != 2\n", i, len(i))
			break
		}

		iisplit := strings.Split(i[0], "+")
		if len(iisplit) != 2 {
			fmt.Printf("error (%v) len %v != 2\n", iisplit, len(iisplit))
			break
		}

		num1, err := strconv.ParseInt(iisplit[0], 10, 32)

		if err != nil {
			log.Fatal(err)
			continue
		}
		num2, err := strconv.ParseInt(iisplit[0], 10, 32)

		if err != nil {
			log.Fatal(err)
			continue
		}

		fmt.Printf(" %v + %v = ?", num1, num2)

		t := time.NewTimer(3 * time.Second)

		select {
		case <-t.C:
			fmt.Println("time's up, next question")
			break
		case ans := <-res:
			if ans == int32(num1+num2) {
				fmt.Println("correct ")
			} else {
				fmt.Println("incorrect")
			}
			t.Stop()
			break
		}
	}
}
func main() {
	quiz2()

}
