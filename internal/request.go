package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mtslzr/pokeapi-go/structs"
)

func GetResourceList(url string) structs.Resource {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	resourceList := structs.Resource{}
	json.Unmarshal(body, &resourceList)
	return resourceList
}
