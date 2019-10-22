package beginner

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type users struct {
	XMLName xml.Name `xml:"users"`
	Users   []user   `json:"users" xml:"user"`
}

type user struct {
	XMLName xml.Name `xml:"user"`
	Name    string   `json:"name"`
	Type    string   `json:"type" xml:"type,attr"`
	Age     int      `json:"Age"`
	Social  social   `json:"social" xml:"social"`
}

type social struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `json:"facebook" xml:"facebook"`
	Twitter  string   `json:"twitter" xml:"twitter"`
}

// RunParseJSON https://tutorialedge.net/golang/parsing-json-with-golang/
func RunParseJSON() {
	// ----- Open JSON file
	// jsonFileName := "/home/al-un/git/go/src/github.com/al-un/tutos-go/lang_tutorialedge/beginner/parseJson.users.json"
	jsonFileName := "./lang_tutorialedge/beginner/parseJson.users.json"
	jsonFile, err := os.Open(jsonFileName)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Successfully opened %v\n", jsonFileName)

	defer jsonFile.Close()

	// ----- Parse it with struct
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users users

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}

	// ----- Parse it with nothing
	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)
	fmt.Println(result)
}

// Users : many dudes
