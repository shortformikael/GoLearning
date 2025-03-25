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
var menu *models.Menu = models.NewMenu()

func main() {
	fmt.Println("=== START ===")
	Run()
	fmt.Println("=== END ===")
}

func Run() {
	for {
		fmt.Printf("Current Status: %s \n", menu.Status)
		if string(menu.Status) == "EXIT" {
			break
		}
		menu.Print()
		SelectChoice(Choose())
	}
}

func Choose() interface{} {
	input := ReadStdInput()
	num, err := strconv.Atoi(input)
	if err != nil {
		return strings.ToUpper(input)
	} else if 1 <= num && num <= len(menu.Items) {
		return num
	}
	return "Out of bounds ERROR" // Return invalid input
}

func SelectChoice(pChoice interface{}) {
	pType := reflect.TypeOf(pChoice).Kind()
	if pType == reflect.Int || pType == reflect.String {
		switch pChoice := pChoice.(type) {
		case int:
			menu.Status = menu.Items[pChoice-1]
		case string:
			for item := range menu.Items {
				if string(menu.Items[item]) == pChoice {
					menu.Status = menu.Items[item]
				}
			}

		}
	}
}

func ReadStdInput() string {
	fmt.Print("Choose option: ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
