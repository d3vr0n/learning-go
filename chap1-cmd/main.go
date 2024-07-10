package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	println("Hello, World!")
	var a = 2
	b := "3"
	c, err := strconv.ParseInt(b, 10, 32)
	d, err1 := strconv.Atoi(b)
	if err != nil || err1 != nil {
		fmt.Println("Error converting string to int", err, err1)
		if err != nil { // does not support ternary operator
			log.Fatal(err)
		} else {
			log.Fatal(err1)
		}
	}
	fmt.Println(a, b, c, d)

	e := int64(a) == c
	fmt.Println(e)

	for i := 0; i < 1; i++ {
		fmt.Println("Loop iteration:", i)
	}

	currentDate := time.Now()
	utcDate := currentDate.UTC()
	formattedDate := currentDate.Format("2006-01-02") // must use 2006-01-02 15:04:05 as reference or Mon Jan 2 15:04:05 MST 2006

	fmt.Println("Current Date:", currentDate)
	fmt.Println("UTC Date:", utcDate)
	fmt.Println("Formatted Date:", formattedDate)

	const (
		g = iota     // 0 iota defines positions in a const block
		h            // 1 = iota
		i = 3 * iota // 6
	)
	fmt.Println(g, h, i)

	fmt.Println("What would you like to do?")
	in := bufio.NewReader(os.Stdin)
	s, err := in.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	fmt.Println(s + "!")

	return
	path := flag.String("path", "myapp.log", "The path for log file")
	level := flag.String("level", "ERROR", "The log level to search for")

	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}
