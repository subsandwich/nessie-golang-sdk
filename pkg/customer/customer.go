package customer

import(
    "fmt"
    "net/http"
    "io/ioutil"
    "bytes"
    "../shared"
)

const baseUrl = "http://api.reimaginebanking.com/customers"
var apiKey = shared.ApiKey

//GET: Returns the customer that the account belongs to
func GetCustomerOfAccount(accountId string) (string, error) {

    var url = "http://api.reimaginebanking.com/accounts/" + accountId + "/customer?key=" + apiKey

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

//GET: Returns the customers that have been assigned to you
func GetAllCustomers() (string, error) {

    var url = baseUrl + "?key=" + apiKey

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

//GET: Returns the customer with the specific id
func GetCustomerWithId(customerId string) (string, error) {

    var url = baseUrl + `/` + customerId + "?key=" + apiKey

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

//POST: Creates a customer
func CreateCustomer(firstName string, lastName string, street_number string, street_name string, city string, state string, zip string) (string, error) {

    var url = baseUrl + "?key=" + apiKey


    var address = `{"street_number": "` + street_number + `", "street_name":"` + street_name + `", "city": "` + city + `", "state":"` + state + `", "zip":"` + zip + `"}`

    var payloadStr = `{"first_name":"` + firstName + `", "last_name": "` + lastName + `", "address":` + address + `}`

    req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(payloadStr)))
    req.Header.Set("Content-Type", "application/json")

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

//PUT: Updates the specific customer
func UpdateCustomer(customerId string, street_number string, street_name string, city string, state string, zip string) (string, error) {

    var url = baseUrl + `/` + customerId + "?key=" + apiKey


    var address = `{"street_number": "` + street_number + `", "street_name":"` + street_name + `", "city": "` + city + `", "state":"` + state + `", "zip":"` + zip + `"}`

    var payloadStr = `{"address":` + address + `}`

    req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(payloadStr)))
    req.Header.Set("Content-Type", "application/json")

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