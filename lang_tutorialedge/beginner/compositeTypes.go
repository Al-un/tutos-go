package beginner

import "fmt"

// RunCompositeTypes https://tutorialedge.net/golang/go-complex-types-tutorial/
func RunCompositeTypes() {
	youtubeSubscribers := map[string]int{
		"TutorialEdge":     2240,
		"MKBHD":            6580350,
		"Fin fun Function": 171220,
	}

	key := "MKBHD"
	fmt.Printf("Number of subscriber of %s, %v\n", key, youtubeSubscribers[key])
	key = "Pouet"
	fmt.Printf("Number of subscriber of %s, %v\n", key, youtubeSubscribers[key])
}
