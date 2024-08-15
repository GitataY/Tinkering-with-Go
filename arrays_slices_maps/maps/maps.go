package maps

import "fmt"

func maps() {
	websites := map[string]string{
		"SitePoint": "https://www.sitepoint.com",
		"Google":    "https://www.google.com",
		"Facebook":  "https://www.facebook.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	websites["LinkedIn"] = "https://www.linkedin.com"
	fmt.Println(websites)

	delete(websites, "Facebook")
	fmt.Println(websites)
}