package main

import (
	"fmt"
	"net/http"
)

func NewRequest(endpoint string) *http.Request {
	res, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	res.Header.Add("X-API-Key", "7973756ba5f149d0860c8f815440e8f4")
	fmt.Print(res.Body)
	return res
}
func main() {
	fmt.Println("lets scan destiny 2's api")
	//rootPath := "https://www.bungie.net/Platform"
	localPath := "https://www.bungie.net/Platform/GetAvailableLocales/"
	fmt.Println(NewRequest(localPath).Body)
}
