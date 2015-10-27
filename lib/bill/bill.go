package bill

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

//POST: Creates a bill
//For Optional params, use empty string "" or -999 for recurring_date
func CreateBill(customerId string, status string, payee string, nickname string, payment_date string, recurring_date int, payment_amount float64){

    var url = baseUrl + "accounts/" + customerId + "/bills?key=" + apiKey

    fmt.Println("url", string(url))

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

//PUT: Updates the specific bill
func UpdateBill(billId string, status string, payee string, nickname string, payment_date string, recurring_date int, payment_amount float64){

    var url = baseUrl + "bills/" + billId + "?key=" + apiKey

    fmt.Println("url", string(url))

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

//DELETE: Deletes the specific bill
func DeleteBill(billId string){

	var url = baseUrl + "bills/" + billId + "?key=" + apiKey
   
    req, err := http.NewRequest("DELETE", url, nil)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Bill Deleted Successfully")
    fmt.Println("response Body:", string(body))
}