package bill

import(
    "fmt"
    "net/http"
    //"bytes"
    "io/ioutil"
    //"strconv"
    "../shared"
)

var baseUrl = "http://api.reimaginebanking.com/"
var apiKey = shared.ApiKey

//GET: Returns the bills that are tied to the specific account
func GetBillsOfAccount(accountId string){

    var url = baseUrl + "accounts/" + accountId + `/bills` + "?key=" + apiKey
   
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

//GET: Returns the bill with the specific id
func GetBillWithId(billId string){

    var url = baseUrl + "bills/" + billId + "?key=" + apiKey
   
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

//GET: Returns the bill with the specific id
func GetBillsOfCustomer(customerId string){

    var url = baseUrl + "customers/" + customerId + "/bills?key=" + apiKey
   
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