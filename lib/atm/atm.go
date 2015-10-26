package atm

import(
    "fmt"
    "net/http"
    "io/ioutil"
    "strconv"
    "../shared"
)

var baseUrl = "http://api.reimaginebanking.com/atms"
var apiKey = shared.ApiKey

//GET: Returns all of the Capital One ATMs in the speified search area (Pages not implemented yet)
func GetAllBranches(lat float64, lng float64, rad int){
	
    var latString = strconv.FormatFloat(lat,'f',4,64)
    var lngString = strconv.FormatFloat(lng,'f',4,64)
    var radString = strconv.Itoa(rad)

	var url = baseUrl + "?lat=" + latString + "&lng=" + lngString + "&rad=" + radString + "&key=" + apiKey
	req, err := http.NewRequest("GET", url, nil)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}

//GET: Returns the ATM with the specific id
func GetATMInfo(atmId string){

    var url = baseUrl + "/" + atmId + "?key=" + apiKey

    req, err := http.NewRequest("GET", url, nil)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}