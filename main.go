package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/shortformikael/GoLearning/models"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var menu *models.menu = models.NewMenu()

func main() {
	fmt.Println("=== START ===")
	Run()
	fmt.Println("=== END ===")
}

func Run() {
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
