package maps

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "https://www.google.com",
		"Amazon Web Services": "https://aws.amazon.com",
	}

	fmt.Println(websites)

	websites["Facebook"] = "https://www.facebook.com"

	fmt.Println(websites["Facebook"])

	delete(websites, "Google")
	fmt.Println(websites)
}