package customer

import(
    "fmt"
    "net/http"
    "io/ioutil"
    "bytes"
    "../shared"
)

var baseUrl = "http://api.reimaginebanking.com/customers"
var apiKey = shared.ApiKey

//GET: Returns the customer that the account belongs to
func GetCustomerOfAccount(accountId string){

    var url = "http://api.reimaginebanking.com/accounts/" + accountId + "/customer?key=" + apiKey

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

//GET: Returns the customers that have been assigned to you
func GetAllCustomers(){

    var url = baseUrl + "?key=" + apiKey

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

//GET: Returns the customer with the specific id
func GetCustomerWithId(customerId string){

    var url = baseUrl + `/` + customerId + "?key=" + apiKey

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

//POST: Creates a customer
func CreateCustomer(firstName string, lastName string, street_number string, street_name string, city string, state string, zip string){

    var url = baseUrl + "?key=" + apiKey

    fmt.Println("URL:>", url)

    var address = `{"street_number": "` + street_number + `", "street_name":"` + street_name + `", "city": "` + city + `", "state":"` + state + `", "zip":"` + zip + `"}`

    var payloadStr = `{"first_name":"` + firstName + `", "last_name": "` + lastName + `", "address":` + address + `}`

    fmt.Println("payload:", string(payloadStr))

    var jsonStr = []byte(payloadStr)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response Status:", resp.Status)
    fmt.Println("Response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response Body:", string(body))
}

//PUT: Updates the specific customer
func UpdateCustomer(customerId string, street_number string, street_name string, city string, state string, zip string){

    var url = baseUrl + `/` + customerId + "?key=" + apiKey

    fmt.Println("URL:>", url)

    var address = `{"street_number": "` + street_number + `", "street_name":"` + street_name + `", "city": "` + city + `", "state":"` + state + `", "zip":"` + zip + `"}`

    var payloadStr = `{"address":` + address + `}`

    fmt.Println("payload:", string(payloadStr))

    var jsonStr = []byte(payloadStr)
    req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response Status:", resp.Status)
    fmt.Println("Response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response Body:", string(body))
}