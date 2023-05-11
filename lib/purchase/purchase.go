package purchase

import(
    "fmt"
    "math"
    "net/http"
    "io/ioutil"
    "bytes"
    "strconv"
    "../shared"
)

const baseUrl = "http://api.reimaginebanking.com/"
var apiKey = shared.ApiKey
const blankNumber = math.SmallestNonzeroFloat64

//GET: Returns the purchases that you are involved in
func GetPurchasesByAccount(accountId string) (string, error) {

	var url = baseUrl + "accounts/" + accountId + "/purchases?key=" + apiKey

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

//GET: Returns the purchase with the specific id
func GetPurchaseById(purchaseId string) (string, error) {

	var url = baseUrl + "purchases/" + purchaseId + "?key=" + apiKey

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

//POST: Creates a purchase where the account with the ID specified is the payer
//For optional Params, use empty string ""
func CreatePurchase(accountId string, merchant_id string, medium string, purchase_date string, amount float64,
     status string, description string) (string, error) {

    url := baseUrl + "accounts/" + accountId + "/purchases?key=" + apiKey

    amountStr := strconv.FormatFloat(amount,'f',4,64)

    payloadStr := `{"merchant_id":"` + merchant_id +  `", "medium":"` + medium + `"`

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

//PUT: Updates the specific purchase
//For optional Params, use empty string "" and blankNumber for optional float
//NOTE: You don't have to update all fields. Any fields you don't include in the body will stay the same
func UpdatePurchase(purchaseId string, payerId string, medium string, amount float64, description string) (string, error) {

    url := baseUrl + "purchases/" + purchaseId + "?key=" + apiKey

    amountStr := strconv.FormatFloat(amount,'f',4,64)

    payloadStr := `{ ` 

    if len(payerId) > 0 {

    	payloadStr = payloadStr + `"payer_id": "` + payerId + `"`  
    } 

    if len(medium) > 0 {
    	if len(payerId) > 0{
        	payloadStr = payloadStr + `,"medium":"` + medium + `"`
    	} else {
    		payloadStr = payloadStr + `"medium":"` + medium + `"`
    	}
	}

    if amount != blankNumber{
    	
    	if len(medium) > 0 || len(payerId) > 0{
    	 	payloadStr = payloadStr + `, "amount": ` + amountStr
    	} else {
    		payloadStr = payloadStr + ` "amount": ` + amountStr
    	}
    }
    
    if len(description) > 0{
    	if(amount!= blankNumber || len(medium) > 0 || len(payerId) > 0){
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

//DELETE: Deletes the specific purchase
func DeletePurchase(purchaseId string) (string, error) {

    url := baseUrl + "purchases/" + purchaseId+ "?key=" + apiKey

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