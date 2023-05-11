package account

import(
    "fmt"
    "net/http"
    "bytes"
    "io/ioutil"
    "strconv"
    "../shared"
)

const baseUrl = "http://api.reimaginebanking.com/"
var apiKey = shared.ApiKey

//GET: Returns the accounts that have been assigned to you
func GetAllAccounts() (string, error) {

    url := "http://api.reimaginebanking.com/accounts?key=" + apiKey

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

//GET: Returns the account with the specific id
func GetAccountWithId(accountId string) (string, error) {

    url := baseUrl + "accounts/" + accountId+ "?key=" + apiKey

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

//GET: Returns the accounts associated with the specific customer
func GetAccountsOfCustomer(customerId string) (string, error) {

    url := baseUrl + "/customers/" + customerId+ "/accounts?key=" + apiKey

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

//POST: Creates an account for the customer with the id provided
//Optional POST Param account_number, use empty sting "" if omitted
func CreateAccount(customerId string, accountType string, nickname string, rewards int, balance int, account_number string) (string, error) {

    url := baseUrl + "/customers/" + customerId + "/accounts?key=" + apiKey


    rewardsString  := strconv.Itoa(rewards)
    balanceString := strconv.Itoa(balance)
    
    var payloadStr = ""

    if len(account_number) > 0 {
        payloadStr = `{"type":"` + accountType + `","nickname":"` + nickname + `","rewards":` + rewardsString + `, "balance":` + balanceString + `, "account_number":"` + account_number + `"}`
    } else{
        payloadStr = `{"type":"` + accountType + `","nickname":"` + nickname + `","rewards":` + rewardsString + `, "balance":` + balanceString + `}`
    }
    
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

//PUT: Updates the specific account
//Optional PUT Param account_number, use empty sting "" if omitted
func UpdateAccount(accountId string, nickname string, account_number string) (string, error) {

    url := baseUrl + "accounts/" + accountId+ "?key=" + apiKey

    var payloadStr = ""

    if len(account_number) > 0 {
        payloadStr = `{"nickname":"` + nickname + `", "account_number":"` + account_number + `"}`
    } else {
        payloadStr = `{"nickname":"` + nickname + `"}`
    }

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

//DELETE: Deletes the specific account
func DeleteAccount(accountId string) (string, error) {

    url := baseUrl + "accounts/" + accountId+ "?key=" + apiKey

    req, err := http.NewRequest("DELETE", url, nil)

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