package purchase

import(
    "fmt"
    "net/http"
    "io/ioutil"
    "bytes"
    "strconv"
    "../shared"
)

var baseUrl = "http://api.reimaginebanking.com/"
var apiKey = shared.ApiKey

//GET: Returns the purchases that you are involved in
func GetPurchasesByAccount(accountId string){

	var url = baseUrl + "accounts/" + accountId + "/purchases?key=" + apiKey

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

//GET: Returns the purchase with the specific id
func GetPurchaseById(purchaseId string){

	var url = baseUrl + "purchases/" + purchaseId + "?key=" + apiKey

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

//POST: Creates a purchase where the account with the ID specified is the payer
//For optional Params, use empty string ""
func CreatePurchase(accountId string, merchant_id string, medium string, purchase_date string, amount float64, status string, description string){

    url := baseUrl + "accounts/" + accountId + "/purchases?key=" + apiKey

    fmt.Println("URL:>", url)

    var amountStr = strconv.FormatFloat(amount,'f',4,64)

    var payloadStr = `{"merchant_id":"` + merchant_id +  `", "medium":"` + medium + `"`

    if len(purchase_date) > 0{
        payloadStr = payloadStr + `, "purchase_date":"` + purchase_date + `"`
    }

    payloadStr = payloadStr + `, "amount": ` + amountStr

    if len(status) > 0{
        payloadStr = payloadStr + `,"status":"` + status + `"`
    }

    
    if len(description) > 0{
        payloadStr = payloadStr + `, "description": "` + description + `"`
    } 
    
    payloadStr = payloadStr + `}`
    
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