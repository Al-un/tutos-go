package beginner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readStdInReader() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("----------------------------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// Convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("Hello yourself")
		} else if strings.Compare("exit", text) == 0 {
			break
		}
	}
}

func readStdInSingleChar() {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(char)

	switch char {
	case 'A':
		fmt.Println("'A' key pressed")
		break
	case 'a':
		fmt.Println("'a' key pressed")
		break
	}
}

func readStdInScanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("input: ", text)

		if text == "exit" {
			fmt.Println("bye")
			break
		}
	}
}

// RunReadingConsoleInput https://tutorialedge.net/golang/reading-console-input-golang/
func RunReadingConsoleInput() {
	// readStdInReader()
	// readStdInSingleChar()
	readStdInScanner()
}
