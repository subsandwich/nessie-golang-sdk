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
func GetAllBranches(){
	
	var url = baseUrl + "?key=" + apiKey

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

//GET: Returns the branch with the specific id
func GetBranchWithId(branchId string){

	var url = baseUrl + branchId + "?key=" + apiKey

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