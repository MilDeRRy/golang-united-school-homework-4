/*package string_sum

import (
	"errors"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	return "", nil
}*/

/*package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

func main() {
	example := "3+5"
	krot, err := strconv.Atoi(example)
	if err != nil {
		log.Fatal(errorEmptyInput)
	}
	for i := 0 ; i <= len(example) - 1; i++ {
		res := (int(example[i]))
		output += res
	}
	//ero, grim := StringSum(example)
	//fmt.Println(grim, ero)
	//fmt.Println(output)
	fmt.Println(krot)
}*/

/*func StringSum(input string) (output string, err error) {
	krot, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(krot)
	return "", nil
}*/

/*package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func main() {
	s := "1+2"
	grim := ReverseString(s)
	fmt.Println(grim)
}
func ReverseString(input string) (output int) {
	em := []rune(input)
	pattern := (regexp.QuoteMeta(`+`))
	for i := 0; i <= len(em)-1; i++ {
		res := (string(em[i]))
		krot, err := strconv.Atoi(res)
		if err != nil {
			log.Fatal(err)

		}
		//fmt.Println(pattern)
		//if x == pattern {
		//	output += krot
		//}
		//output += krot

		switch {
		case pattern == +:
			output += krot
		}

	}
	return output
}*/

/*package main

import (
	"fmt"
	"os"
	"regexp"
)

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	emails := []string{
		"3-1",
		"3",
		"1",
		"3+1",
	}

	pattern := `^\d`
	for _, email := range emails {
		matched, err := regexp.Match(pattern, []byte(email))
		check(err)
		if matched {
			fmt.Printf("√ '%s' is a valid email\n", email)
		} else {
			fmt.Printf("X '%s' is not a valid email\n", email)
		}
	}
}*/

/*package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func main() {
	s := "2+2"
	grim := ReverseString(s)
	fmt.Println(grim)
}
func ReverseString(input string) (output int) {
	em := []rune(input)
	//pattern := (regexp.QuoteMeta(`+`))
	for i := 0; i <= len(em)-1; i++ {

		res := (string(em[i]))
		krot, err := strconv.Atoi(res)
		if err != nil {
			log.Fatal(err)
		}

		matched, err := regexp.Match(`\+`, []byte(input))
		//fmt.Println(matched, err)
		switch matched {
		case true:
			output = 2 + 3
			fmt.Println(output)
		}
		matched, err := regexp.Match(`\+`, []byte(input))
		//fmt.Println(matched, err)
		switch matched {
		case true:
			output = 2 + 3
			fmt.Println(output)
		matched, err := regexp.Match(`\+`, []byte(input))
		//fmt.Println(matched, err)
		switch matched {
		case true:
			output = 2 + 3
			fmt.Println(output)
		//fmt.Println(pattern)
		//if x == pattern {
		//	output += krot
		//}
		output += krot
	}
	return output
}*/

package string_sum

import (
	"fmt"
	"regexp"
	"strconv"
)

/*func main() {
	s := "6+4"
	grim := ReverseString(s)
	fmt.Println(grim)
}*/
func ReverseString(input string) (output int) {
	em := []rune(input)
	for i := range input {
		if i%2 == 0 {
			res := (string(em[i]))
			krot, err := strconv.Atoi(res)
			if err == nil {
				fmt.Println(err)
			}
			matched, err := regexp.Match(`^\d\+\d$`, []byte(input))
			switch matched {
			case true:
				output += krot
			}
			//matched1, err := regexp.Match(`^\d\-\d$`, []byte(input))
			//switch matched1 {
			//case true:
			//	output += krot
			//fmt.Println(krot)
			//}
			//matched2, err := regexp.Match(`^\-\d\-\d$`, []byte(input))
			//switch matched2 {
			//case true:
			//	output -= krot
			//}
		}
	}
	return output
}
