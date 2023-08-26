package deposit

import(
    "fmt"
    "math"
    "net/http"
    "io/ioutil"
    "strconv"
    "bytes"
    "../shared"
)

const baseUrl = "http://api.reimaginebanking.com/"
var apiKey = shared.ApiKey
const blankNumber = math.SmallestNonzeroFloat64

//GET: Returns the deposits that you are involved in
func GetDepositOfAccount(accountId string) (string, error) {

    var url = baseUrl + "accounts/" + accountId + "/deposits?key=" + apiKey

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

//GET: Returns the deposit with the specific id
func GetDepositById(depositId string) (string, error) {

    var url = baseUrl + "deposits/" + depositId + "?key=" + apiKey

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
//Optional POST Param transaction_date, status, description, use empty sting "" if omitted
func CreateDeposit(accountId string, medium string, transaction_date string, status string, amount float64, description string) (string, error) {

    url := baseUrl + "accounts/" + accountId + "/deposits?key=" + apiKey

    var amountStr = strconv.FormatFloat(amount,'f',4,64)

    var payloadStr = `{"medium":"` + medium + `"`

    if len(transaction_date) > 0{
        payloadStr = payloadStr + `, "transaction_date":"` + transaction_date + `"`
    }

    if len(status) > 0{
        payloadStr = payloadStr + `,"status":"` + status + `"`
    }

    payloadStr = payloadStr + `, "amount": ` + amountStr
    
    if len(description) > 0{
        payloadStr = payloadStr + `, "description": "` + description + `"`
    } 
    
    payloadStr = payloadStr + `}`

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

//PUT: Updates the specific deposit
//For optional Params, use empty string "" and "blankNumber" for optional float
//NOTE: You don't have to update all fields. Any fields you don't include will stay the same
func UpdateDeposit(depositId string, medium string, amount float64, description string) (string, error) {

    url := baseUrl + "deposits/" + depositId + "?key=" + apiKey

    amountStr := strconv.FormatFloat(amount,'f',4,64)

    payloadStr := `{ ` 

    if len(medium) > 0 {
        payloadStr = payloadStr + `"medium":"` + medium + `"`
    }

    if amount != blankNumber{
        
        if len(medium) > 0 {
            payloadStr = payloadStr + `, "amount": ` + amountStr
        } else {
            payloadStr = payloadStr + ` "amount": ` + amountStr
        }
    }
    
    if len(description) > 0{
        if amount!= blankNumber || len(medium) > 0 {
            payloadStr = payloadStr + `, "description": "` + description + `"`
        } else{
            payloadStr = payloadStr + `, "description": "` + description + `"`
        }
    } 
    
    payloadStr = payloadStr + `}`

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

//DELETE: Deletes the specific deposit
func DeleteDeposit(depositId string) (string, error) {

    url := baseUrl + "purchases/" + depositId+ "?key=" + apiKey

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