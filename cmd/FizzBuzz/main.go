package main

import "fmt"

/*FizzBuzz starter code to become familiar with go
- Stanford Solar Car Project - Data Team
Sabrina Flemming
November 2017

Great source: https://graysonkoonce.com/fizzbuzz-in-golang/
 */

func main() {
		fmt.Println("vim-go")
		i := 1								//initialize the counter i
		for ; i <= 100; {					//loop will end when i hits 101
			result := ""					//initialize result as an empty string
			if i%3 == 0 {					//if i is a multiple of 3
				result += "Fizz"
			}
			if i%4 == 0 { 					//if i is a multiple of 4
				result += "Buzz"
			}

			if result != "" {				//if result is not an empty string, print result, otherwise print i
				fmt.Println(result)
			}else {
				fmt.Println(i)
			}

			i++								//increment i
		}
}
