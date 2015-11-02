package transfer

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

//GET: Returns the transfers that you are involved in
func GetTransfersByAccount(accountId string){
	
	url := baseUrl + "accounts/" + accountId + "/transfers?key=" + apiKey

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

//GET: Returns the transfer with the specific id
func GetTransferById(transferId string){
	
	url := baseUrl + "transfers/" + transferId + "?key=" + apiKey

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

//POST: Creates a transfer where the account with the ID specified is the payer
//Optional POST Param transaction_date, status, description, use empty sting "" if omitted
func CreateTransfer(accountId string, medium string, payeeId string, amount float64, transaction_date string, status string, description string){

    url := baseUrl + "accounts/" + accountId + "/transfers?key=" + apiKey

    fmt.Println("URL:>", url)

    var amountStr = strconv.FormatFloat(amount,'f',4,64)

    var payloadStr = `{"medium":"` + medium + `", "payee_id": "` + payeeId + `", "amount": ` + amountStr

    if len(transaction_date) > 0{
        payloadStr = payloadStr + `, "transaction_date":"` + transaction_date + `"`
    }

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

//PUT: Updates the specific purchase
//For optional Params, use empty string "" and blankNumber for optional float
//NOTE: You don't have to update all fields. Any fields you don't include in the body will stay the same
func UpdateTransfer(transferId string, medium string, payeeId string, amount float64, description string){

    url := baseUrl + "transfers/" + transferId + "?key=" + apiKey

    fmt.Println("URL:>", url)

    amountStr := strconv.FormatFloat(amount,'f',4,64)

    payloadStr := `{ ` 

    if len(medium) > 0 {

    	payloadStr = payloadStr + `"medium": "` + medium + `"`  
    } 

    if len(payeeId) > 0 {
    	if len(medium) > 0{
        	payloadStr = payloadStr + `,"payee_id":"` + payeeId + `"`
    	} else {
    		payloadStr = payloadStr + `"payee_id":"` + payeeId + `"`
    	}
	}

    if amount != math.SmallestNonzeroFloat64{
    	
    	if len(medium) > 0 || len(payeeId) > 0{
    	 	payloadStr = payloadStr + `, "amount": ` + amountStr
    	} else {
    		payloadStr = payloadStr + ` "amount": ` + amountStr
    	}
    }
    
    if len(description) > 0{
    	if(amount!= -999 || len(medium) > 0 || len(payeeId) > 0){
        	payloadStr = payloadStr + `, "description": "` + description + `"`
    	} else{
    		payloadStr = payloadStr + `, "description": "` + description + `"`
    	}
    } 
    
    payloadStr = payloadStr + `}`
    
    fmt.Println(string(payloadStr))
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

//DELETE: Deletes the specific transfer
func DeleteTransfer(transferId string){

    url := baseUrl + "transfers/" + transferId + "?key=" + apiKey

    req, err := http.NewRequest("DELETE", url, nil)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response Body: Account was succesfully deleted")
}