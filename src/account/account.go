package account

import(
    "fmt"
    "net/http"
    "bytes"
    "io/ioutil"
    "strconv"
    "../shared"
)

var baseUrl = "http://api.reimaginebanking.com/"
var apiKey = shared.ApiKey

//GET: Returns the accounts that have been assigned to you
func GetAllAccounts(){

    url := "http://api.reimaginebanking.com/accounts?key=" + apiKey

    req, err := http.NewRequest("GET", url, nil)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response Body:", string(body))
}

//GET: Returns the customer with the specific id
func GetCustomerWithId(accountId string){

    url := baseUrl + "accounts/" + accountId+ "?key=" + apiKey

    req, err := http.NewRequest("GET", url, nil)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response Body:", string(body))
}

//GET: Returns the accounts associated with the specific customer
func GetAccountsWithId(customerId string){

    url := baseUrl + "/customers/" + customerId+ "/accounts?key=" + apiKey

    req, err := http.NewRequest("GET", url, nil)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response Body:", string(body))
}

//POST: Creates an account for the customer with the id provided (optional param not implemented yet)
func CreateAccount(customerId string, accountType string, nickname string, rewards int, balance int){

    url := baseUrl + "/customers/" + customerId + "/accounts?key=" + apiKey

    fmt.Println("URL:>", url)

    rewardsString  := strconv.Itoa(rewards)
    balanceString := strconv.Itoa(balance)
    
    var payloadStr = `{"type":"` + accountType + `","nickname":"` + nickname + `","rewards":` + rewardsString + `, "balance":` + balanceString + `}`
    
    fmt.Println(string(payloadStr))
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

//PUT: Updates the specific account (Optional param not implemented yet)
func UpdateAccount(accountId string, nickname string){

    url := baseUrl + "accounts/" + accountId+ "?key=" + apiKey
    var payloadStr = `{"nickname":"` + nickname + `"}`

    var jsonStr = []byte(payloadStr)
    req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response Body:", string(body))
}

//DELETE: Deletes the specific account
func DeleteAccount(accountId string){

    url := baseUrl + "accounts/" + accountId+ "?key=" + apiKey

    req, err := http.NewRequest("DELETE", url, nil)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response Body: Account was succesfully deleted")
}