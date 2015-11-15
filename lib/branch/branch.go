package branch

import(
    "fmt"
    "net/http"
    "io/ioutil"
    "../shared"
)

const baseUrl = "http://api.reimaginebanking.com/branches/"
var apiKey = shared.ApiKey

//GET: Returns all of the Capital One branches.
func GetAllBranches() string {
	
	var url = baseUrl + "?key=" + apiKey

	req, err := http.NewRequest("GET", url, nil)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response Status:", resp.Status)
    body, _ := ioutil.ReadAll(resp.Body)
    var response = string(body)
    //fmt.Println("Response Body:", response)
    return response
}

//GET: Returns the branch with the specific id
func GetBranchWithId(branchId string) string {

	var url = baseUrl + branchId + "?key=" + apiKey

	req, err := http.NewRequest("GET", url, nil)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("Response Status:", resp.Status)
    body, _ := ioutil.ReadAll(resp.Body)
    var response = string(body)
    //fmt.Println("Response Body:", response)
    return response
}