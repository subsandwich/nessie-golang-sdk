package atm

import(
    "fmt"
    "net/http"
    "io/ioutil"
    "strconv"
    "../shared"
)

const baseUrl = "http://api.reimaginebanking.com/atms"
var apiKey = shared.ApiKey

//GET: Returns all of the Capital One ATMs in the speified search area (Pages not implemented yet)
func GetAllATMs(lat float64, lng float64, rad int) (string, error) {
	
    var latString = strconv.FormatFloat(lat,'f',4,64)
    var lngString = strconv.FormatFloat(lng,'f',4,64)
    var radString = strconv.Itoa(rad)

	var url = baseUrl + "?lat=" + latString + "&lng=" + lngString + "&rad=" + radString + "&key=" + apiKey
	req, err := http.NewRequest("GET", url, nil)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(body), nil
}

//GET: Returns the ATM with the specific id
func GetATMInfo(atmId string) (string, error) { 

    var url = baseUrl + "/" + atmId + "?key=" + apiKey

    req, err := http.NewRequest("GET", url, nil)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(body), nil
}