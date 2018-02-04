package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	"io"
	"strconv"
	"./algorithms/bubblesort"
	"./algorithms/qsort"
	"time"
	"math/rand"
)

var infile = flag.String("i", "unsorted.dat", "File contains values for sorting")
var outfile = flag.String("o", "sorted.dat", "File to receive sorted values")
var algorithm = flag.String("a", "qsort", "Sort algorithm")
var num = flag.Int("n", 100, "The number of nums to sort")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("File to open input file ", infile)
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}
		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
	}
	defer file.Close()
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func generateRandomFile(num int) {
	nums := make([]int, num)
	for i := 0; i < num; i++ {
		nums[i] = random(1, num)
	}
	writeValues(nums, *infile)
}

func main() {
	flag.Parse()
	generateRandomFile(*num)

	nums, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(nums)
		case "bubblesort":
			bubblesort.BubbleSort(nums)
		default:
			fmt.Println("Unknown sort algorithm")
		}
		t2 := time.Now()
		fmt.Println("Done! ", t2.Sub(t1))
		writeValues(nums, "sorted.dat")
	} else {
		fmt.Println(err)
	}
}
