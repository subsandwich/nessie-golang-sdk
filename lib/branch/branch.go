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
func GetAllBranches() (string, error) {
	
	var url = baseUrl + "?key=" + apiKey

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

//GET: Returns the branch with the specific id
func GetBranchWithId(branchId string) (string, error) {

	var url = baseUrl + branchId + "?key=" + apiKey

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