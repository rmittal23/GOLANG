package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type Person struct {
		fname string
		lname string
	}
	sli := make([]Person, 0)
	fmt.Println("Name of the file")
	var file string
	fmt.Scan(&file)
	f, _ := os.Open(file)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		scanner.Split(bufio.ScanWords)
		p1 := Person{
			fname: strings.Split(scanner.Text(), " ")[0],
			lname: strings.Split(scanner.Text(), " ")[1],
		}
		sli = append(sli, p1)
	}
	for _, v := range sli {
		fmt.Printf("First Name is:%v and Last Name is: %v", v.fname, v.lname)
	}
	f.Close()
}
