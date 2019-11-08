package filesystem

import (
	"fmt"
	"io/ioutil"
	"os"
)

// RunTemporaryFiles https://tutorialedge.net/golang/temporary-files-directories-go-111/
func RunTemporaryFiles() {
	// --- Create a temporary file
	file, err := ioutil.TempFile("filesystem/car-images", "car-*.png")
	if err != nil {
		fmt.Println("TempFile Error: ", err)
	}
	defer os.Remove(file.Name())

	fmt.Println("Created temporary file: ", file.Name())

	// --- Write some stuff
	if _, err := file.Write([]byte("hello world\n")); err != nil {
		fmt.Println(err)
	}

	// --- Check it out!
	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		fmt.Println("TempFile Error: ", err)
	}

	fmt.Println(string(data))
}
