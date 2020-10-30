package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func getUserString(msg string, scanner *bufio.Scanner) string {
	var user_string string
	fmt.Printf("%s", msg)
	scanner.Scan()
	user_string = scanner.Text()
	return user_string
}

func reverse(l []string) []string {
	var revesed_slice []string
	for h := (len(l) - 1); h >= 0; h-- {
		revesed_slice = append(revesed_slice, l[h])
	}
	return revesed_slice
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func join(s []string, join_character string) (r_string string) {
	for _, current_string := range s {
		r_string += fmt.Sprintf("%s%s", current_string, join_character)
	}
	return
}

func saveSlice(l []string) {
	fa, err := os.Create("Acendente.txt")
	check(err)
	defer fa.Close()

	fd, err := os.Create("Decendente.txt")
	check(err)
	defer fd.Close()

	_, err = fa.WriteString(join(l, "\n"))
	check(err)

	_, err = fd.WriteString(join(reverse(l), "\n"))
	check(err)

}

func main() {
	var user_strings []string
	var user_strings_length int
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	user_strings_length, _ = strconv.Atoi(getUserString("Hey give me a number: ", scanner))
	for h := 0; h < user_strings_length; h++ {
		user_strings = append(user_strings, getUserString("string :", scanner))
	}
	sort.Strings(user_strings)
	saveSlice(user_strings)
}
