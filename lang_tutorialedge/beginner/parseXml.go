package beginner

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// RunParseXML https://tutorialedge.net/golang/parsing-xml-with-golang/
// structs are fined in parseJson.jso
func RunParseXML() {
	// ----- Open JSON file
	// jsonFileName := "/home/al-un/git/go/src/github.com/al-un/tutos-go/lang_tutorialedge/beginner/parseXml.users.xml"
	jsonFileName := "./lang_tutorialedge/beginner/parseXml.users.xml"
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

	xml.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}

	// ----- Parse it with nothing
	var result map[string]interface{}
	xml.Unmarshal(byteValue, &result)
	fmt.Println(result)
}

// Users : many dudes
