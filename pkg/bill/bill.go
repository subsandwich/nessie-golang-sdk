package bill

import(
    "fmt"
    "math"
    "net/http"
    "bytes"
    "io/ioutil"
    "strconv"
    "../shared"
)

const baseUrl = "http://api.reimaginebanking.com/"
var apiKey = shared.ApiKey
const blankNumber = math.SmallestNonzeroFloat64

//GET: Returns the bills that are tied to the specific account
func GetBillsOfAccount(accountId string) (string, error) {

    var url = baseUrl + "accounts/" + accountId + `/bills` + "?key=" + apiKey
   
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

//GET: Returns the bill with the specific id
func GetBillWithId(billId string) (string, error) {

    var url = baseUrl + "bills/" + billId + "?key=" + apiKey
   
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

//GET: Returns the bill with the specific id
func GetBillsOfCustomer(customerId string) (string, error) {

    var url = baseUrl + "customers/" + customerId + "/bills?key=" + apiKey
   
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

//POST: Creates a bill
//For Optional params, use empty string "" or blankNumber for recurring_date
func CreateBill(customerId string, status string, payee string, nickname string, payment_date string, recurring_date int, payment_amount float64) (string, error) {

    var url = baseUrl + "accounts/" + customerId + "/bills?key=" + apiKey


    var recurring_dateStr = strconv.Itoa(recurring_date)
    var payment_amountStr = strconv.FormatFloat(payment_amount,'f',4,64)

    var payloadStr = `{"status":"` + status + `","payee":"` + payee + `"`

    if len(nickname) > 0{
    	payloadStr = payloadStr + `, "nickname":"` + nickname + `"`
    }

    if len(payment_date) > 0{
    	payloadStr = payloadStr + `,"payment_date":"` + payment_date + `"`
    }
    
    if recurring_date > 0 && recurring_date < 31 {
    	payloadStr = payloadStr + `, "recurring_date": ` + recurring_dateStr
    } 
    
    payloadStr = payloadStr + `, "payment_amount":` + payment_amountStr + `}`

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

//PUT: Updates the specific bill
func UpdateBill(billId string, status string, payee string, nickname string, payment_date string, recurring_date int, payment_amount float64) (string, error) {

    var url = baseUrl + "bills/" + billId + "?key=" + apiKey

    var recurring_dateStr = strconv.Itoa(recurring_date)
    var payment_amountStr = strconv.FormatFloat(payment_amount,'f',4,64)

    var payloadStr = `{"status":"` + status + `","payee":"` + payee + `"`

    if len(nickname) > 0{
    	payloadStr = payloadStr + `, "nickname":"` + nickname + `"`
    }

    if len(payment_date) > 0{
    	payloadStr = payloadStr + `,"payment_date":"` + payment_date + `"`
    }
    
    if recurring_date > 0 && recurring_date < 31 {
    	payloadStr = payloadStr + `, "recurring_date": ` + recurring_dateStr
    } 
    
    payloadStr = payloadStr + `, "payment_amount":` + payment_amountStr + `}`


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

//DELETE: Deletes the specific bill
func DeleteBill(billId string) (string, error) {

	var url = baseUrl + "bills/" + billId + "?key=" + apiKey
   
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