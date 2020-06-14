package AddressRecover

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetIP method will return the user IP address.
func GetAddress() string {
	resp, err := http.Get("https://ipinfo.io/json")
	if err != nil {
		panic("Something went wrong!")
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(string(body)), &dat); err != nil {
		panic(err)
	}

	return dat["ip"].(string)
}
