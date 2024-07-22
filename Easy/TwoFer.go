package Easy

import "fmt"

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	if name != "" {
		return fmt.Sprintf("One for %s, one for me.", name)
	} else {
		return "One for you, one for me."
	}
}

func Easy() {
	name := "amanda"
	fmt.Println(ShareWith(name))
}
