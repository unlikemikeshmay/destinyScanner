package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func main() {
	//bungieId := "20814900"
	//etbungieId := fmt.Sprintf(`/User/GetBungieNetUserById/?id=%s/`, bungieId)
	//var getbungieIdUrl = fmt.Sprintf(`https://www.bungie.net/platform%s`, getbungieId)
	//client_id := "38993"
	//fmt.Println("lets scan destiny 2's api")
	//fmt.Println(newRequest("/User/GetMembershipFromHardLinkedCredential/SteamId/sweatleaf666/"))
	//rootPath := "https://www.bungie.net/Platform"
	//localPath := "https://www.bungie.net/platform/User/GetMembershipFromHardLinkedCredential/SteamId/sweatleaf666/"
	//oAuthUrl := fmt.Sprintf("https://www.bungie.net/en/oauth/authorize?client_id=%s&response_type=code&state=6i0mkD4rT1Kppj9WVeHrzHG4", client_id)
	//call(getbungieIdUrl, "GET")

	openbrowser("GET https://www.bungie.net/en/oauth/authorize?client_id=12345&response_type=code&state=6i0mkLx79Hp91nzWVeHrzHG4")
}

/* func callOauth(url, method string) error {
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

	filepath := "oAuth.html"

	f, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("Error creating filepath: %s", err.Error())
	}

	_, err2 := f.WriteString(bodyString)
	if err2 != nil {
		fmt.Println(err2.Error())
		return fmt.Errorf("Error writing body string to file: %s", err2)
	}
	defer f.Close()

	defer response.Body.Close()
	return nil
} */
/* func call(url, method string) error {
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

	bodyString := string(bodyBytes)
	fmt.Printf("bodyString2: %s\n", bodyString)
	fmt.Printf("response.Body type: %s\n", reflect.TypeOf(response.Body))
	defer response.Body.Close()
	return nil
} */

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", "https://www.bungie.net/en/oauth/authorize?client_id=12345&response_type=code&state=6i0mkLx79Hp91nzWVeHrzHG4").Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
