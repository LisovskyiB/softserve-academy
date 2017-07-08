package tasks

import (
	"os"
	"fmt"
	"time"
	"math"
	"bufio"
	"strconv"
	"math/rand"
)

func RunTask() {

	scanner := bufio.NewScanner(os.Stdin)

	var number string

	for number != "q"{

		fmt.Print("Enter number of task or q to exit: ")

		scanner.Scan()

		number = scanner.Text()

		if number == "178" {

			fmt.Print("Enter letter: ")

			scanner.Scan()

			letter := scanner.Text()

			if letter == "d" || letter == "D" {

				fmt.Println("You choose :", number, letter, "task")

				fmt.Print("Enter slice size: ")

				scanner.Scan()

				sizeStr := scanner.Text()

				size,_ := strconv.Atoi(sizeStr)

				fmt.Printf("%+v", "result : ")

				task178D(size)

			}else if letter == "e" || letter == "E" {

				fmt.Println("You choose :", number, letter, "task")

				fmt.Print("Enter slice size: ")

				scanner.Scan()

				sizeStr := scanner.Text()

				size,_ := strconv.Atoi(sizeStr)

				fmt.Printf("%+v", "result : ")

				task178E(size)
			}

		}else if number == "555" {

			fmt.Print("Enter n: ")

			scanner.Scan()

			size := scanner.Text()

			triangleSize,_  := strconv.Atoi(size)

			fmt.Printf("%+v\n", "result : ")

			task555(triangleSize)
		}
	}
}

func returnRandomSlice(size int) []int{

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r.Perm(size)
}

func task178D(size int) {

	numbers := returnRandomSlice(size)

	count := 0

	for i, value := range numbers {

		if i == len(numbers)-1 {
			break
		}

		if i == 0 {
			continue
		}

		if value < (numbers[i - 1] + numbers[i + 1]) / 2 {

			count ++
		}

		time.Sleep(time.Millisecond * 5)
	}

	fmt.Printf("%+v\n", count)
}

func task178E(size int) {

	numbers := returnRandomSlice(size)

	count := 0

	for i, value := range numbers {

		if i == len(numbers) - 1 {
			break
		}

		x := factorial(uint64(value))

		if math.Pow(2, float64(value)) < float64(value) && x > uint64(value) {

			count ++
		}
	}

	fmt.Printf("%+v\n", count)
}

func task555(n int) {// Pifagor treangle

	var triangle [ 100 ][ 100 ]int

	var i, j, k, sp int

	triangle[ 0 ][ 0 ] = 1

	for i = 1; i <= n; i++ {

		for j = 0; j <= i; j++ {

			if 0 == j || j == i {

				triangle[ i ][ j ] = 1

			} else {

				triangle[ i ][ j ] = triangle[ i - 1][ j ] + triangle[ i - 1 ][ j - 1 ]
			}
		}
	}

	sp = n

	for i = 0; i <= n; i++ {

		for k = sp; k >= 0; k -- {

			fmt.Printf(" ")
		}

		sp --

		for j = 0; j <= i; j++ {

			fmt.Printf("%d ", triangle[i][j])
		}

		fmt.Printf("\n")
	}

}

func factorial(n uint64)(result uint64) {

	if n > 0 {
		result = n * factorial(n-1)

		return result
	}

	return 1
}
