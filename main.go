package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"reflect"
	"time"
)

func main() {
	//bungieId := "20814900"
	//etbungieId := fmt.Sprintf(`/User/GetBungieNetUserById/?id=%s/`, bungieId)
	//var getbungieIdUrl = fmt.Sprintf(`https://www.bungie.net/platform%s`, getbungieId)
	client_id := "38993"
	//fmt.Println("lets scan destiny 2's api")
	//fmt.Println(newRequest("/User/GetMembershipFromHardLinkedCredential/SteamId/sweatleaf666/"))
	//rootPath := "https://www.bungie.net/Platform"
	//localPath := "https://www.bungie.net/platform/User/GetMembershipFromHardLinkedCredential/SteamId/sweatleaf666/"
	oAuthUrl := fmt.Sprintf("https://www.bungie.net/en/oauth/authorize?client_id=%s&response_type=code&state=6i0mkD4rT1Kppj9WVeHrzHG4", client_id)
	//call(getbungieIdUrl, "GET")
	callOauth(oAuthUrl, "GET")
}
func callOauth(url, method string) error {
	client_id := "38993"
	oAuthUrl := fmt.Sprintf("https://www.bungie.net/en/oauth/authorize?client_id=%s&response_type=code&state=6i0mkD4rT1Kppj9WVeHrzHG4", client_id)

	fmt.Println(oAuthUrl, " ", client_id)
	client := &http.Client{
		Timeout: time.Second * 15,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("Error in callOauth after defining request: %s", err.Error())
	}
	req.Header.Set("X-API-Key", "7973756ba5f149d0860c8f815440e8f4")
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("Error after sending the request", err.Error())
	}
	bodyBytes, err := io.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	ech := "echo"
	direc := ">"
	filepath := "oAuth.html"
	//fmt.Printf("bodyString: %s", bodyString)
	cmd := exec.Command(ech, bodyString, direc, filepath)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("Error after sending the request", err.Error())
	}
	fmt.Print(string(stdout))
	defer response.Body.Close()
	return nil
}
func call(url, method string) error {
	client := &http.Client{
		Timeout: time.Second * 15,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("Got error %s", err.Error())
	}
	req.Header.Set("X-API-Key", "7973756ba5f149d0860c8f815440e8f4")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("Got error %s", err.Error())
	}
	bodyBytes, err := io.ReadAll(response.Body)

	/* if err != nil {
		log.Fatal(err)
	}
	*/
	bodyString := string(bodyBytes)
	fmt.Printf("bodyString2: %s\n", bodyString)
	fmt.Printf("response.Body type: %s\n", reflect.TypeOf(response.Body))
	defer response.Body.Close()
	return nil
}
