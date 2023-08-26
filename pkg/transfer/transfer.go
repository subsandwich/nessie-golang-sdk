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
func GetTransfersByAccount(accountId string) (string, error) {
	
	url := baseUrl + "accounts/" + accountId + "/transfers?key=" + apiKey

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

//GET: Returns the transfer with the specific id
func GetTransferById(transferId string) (string, error) {
	
	url := baseUrl + "transfers/" + transferId + "?key=" + apiKey

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

//POST: Creates a transfer where the account with the ID specified is the payer
//Optional POST Param transaction_date, status, description, use empty sting "" if omitted
func CreateTransfer(accountId string, medium string, payeeId string, amount float64, transaction_date string, 
        status string, description string) (string, error) {

    url := baseUrl + "accounts/" + accountId + "/transfers?key=" + apiKey

    amountStr := strconv.FormatFloat(amount,'f',4,64)

    payloadStr := `{"medium":"` + medium + `", "payee_id": "` + payeeId + `", "amount": ` + amountStr

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

//PUT: Updates the specific transfer
//For optional Params, use empty string "" and blankNumber for optional float
//NOTE: You don't have to update all fields. Any fields you don't include in the body will stay the same
func UpdateTransfer(transferId string, medium string, payeeId string, amount float64, description string) (string, error) {

    url := baseUrl + "transfers/" + transferId + "?key=" + apiKey

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

//DELETE: Deletes the specific transfer
func DeleteTransfer(transferId string) (string, error) {

    url := baseUrl + "transfers/" + transferId + "?key=" + apiKey

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