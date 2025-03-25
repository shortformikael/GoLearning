package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"./models"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var menu *models.Menu = models.NewMenu()

func main() {
	fmt.Println("=== START ===")
	run()
	fmt.Println("=== END ===")
}

func run() {
	for {
		fmt.Printf("Current Status: %s \n", menu.Status)
		menu.Print()
		SelectChoice(Choose())
		break
	}
}

func Choose() interface{} {
	input := ReadStdInput()
	num, err := strconv.Atoi(input)
	if err != nil {
		return input
	}
	return num
}

func SelectChoice(pChoice interface{}) {
	pType := reflect.TypeOf(pChoice).Kind()
	if pType == reflect.Int || pType == reflect.String {
		//Exec Slection
	}
}

func ReadStdInput() string {
	fmt.Print("Choose option: ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
