package merchant

import(
    "fmt"
    "math"
    "net/http"
    "io/ioutil"
    "strconv"
    "bytes"
    "strings"
    "../shared"
)

const baseUrl = "http://api.reimaginebanking.com/merchants"
var apiKey = shared.ApiKey
const blankNumber = math.SmallestNonzeroFloat64

func IsNumeric (text string) bool{
    if _, err := strconv.Atoi(text); err == nil {
        return true
    }
    return false
}

//GET: Returns the merchants that have been assigned to you
func GetAllMerchants(lat float64, lng float64, rad int) (string, error) {
	
    latString := strconv.FormatFloat(lat,'f',4,64)
    lngString := strconv.FormatFloat(lng,'f',4,64)
    radString := strconv.Itoa(rad)

	url := baseUrl + "?lat=" + latString + "&lng=" + lngString + "&rad=" + radString + "&key=" + apiKey

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

//GET: Returns the merchant with the specific id
func GetMerchantInfo(merchantId string) (string, error) {

    var url = baseUrl + "/" + merchantId + "?key=" + apiKey

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

//POST: Creates a merchant
//For optional Params, use empty string "" and blankNumber for empty lat/lng
func CreateMerchant(merchantName string, categories []string, street_number string, street_name string, city string, state string, zip string,
         lat float64, lng float64) (string, error) {

    if len(categories) == 0 {
        return "", fmt.Errorf("Categories field cannot be empty")
    }

    if len(state) > 2 {
        return "", fmt.Errorf("State field needs to be the two letter abbreviation of the state")
    }

    if len(zip) != 5 || !IsNumeric(zip) {
        return "", fmt.Errorf("Zip code field needs to be numeric and have a length of 5")
    }

    formattedCategories := make([]string, len(categories))

    for _, category := range categories {
        formattedCategories = append(formattedCategories, `"`+ category +`"`)
    }

    formattedCategories = append(formattedCategories[:0], formattedCategories[2:]...)

    categoriesString := "["
    categoriesString = categoriesString + strings.Join(formattedCategories, ",") + "]"

    url := baseUrl + "?key=" + apiKey


    var latString = strconv.FormatFloat(lat,'f',4,64)
    var lngString = strconv.FormatFloat(lng,'f',4,64)

    var geocode = `{"lat": ` + latString + `, "lng": ` + lngString + `}`
    var address = `{"street_number": "` + street_number + `", "street_name":"` + street_name + `", "city": "` + city + `", "state":"` + state + `", "zip":"` + zip + `"}`

    if lat == blankNumber || lng == blankNumber {
    	geocode = ""
    }

    var payloadStr = `{"name":"` + merchantName + `"`

    payloadStr = payloadStr + `, "category":` + categoriesString

    if len(street_number) > 0{
    	payloadStr = payloadStr + `,"address":` + address
    }
    
    if len(geocode) > 0{
    	payloadStr = payloadStr + `, "geocode": ` + geocode
    } 
    
    payloadStr = payloadStr + `}`

    var jsonStr = []byte(payloadStr)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
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

//PUT: Updates a specific merchant
//For optional Params, use empty string "" and blankNumber for empty lat/lng
func UpdateMerchant(merchantId string, merchantName string, categories []string, street_number string, street_name string, city string, state string, zip string,
         lat float64, lng float64) (string, error) {

    url := baseUrl + "/" + merchantId + "?key=" + apiKey

    if len(categories) == 0 {
        return "", fmt.Errorf("Categories field cannot be empty")
    }

    if len(state) > 2 {
        return "", fmt.Errorf("State field needs to be the two letter abbreviation of the state")
    }

    if len(zip) != 5 || !IsNumeric(zip) {
        return "", fmt.Errorf("Zip code field needs to be numeric and have a length of 5")
    }

    formattedCategories := make([]string, len(categories))

    for _, category := range categories {
        formattedCategories = append(formattedCategories, `"`+ category +`"`)
    }

    formattedCategories = append(formattedCategories[:0], formattedCategories[2:]...)

    categoriesString := "["
    categoriesString = categoriesString + strings.Join(formattedCategories, ",") + "]"

    var latString = strconv.FormatFloat(lat,'f',4,64)
    var lngString = strconv.FormatFloat(lng,'f',4,64)

    var geocode = `{"lat": ` + latString + `, "lng": ` + lngString + `}`
    var address = `{"street_number": "` + street_number + `", "street_name":"` + street_name + `", "city": "` + city + `", "state":"` + state + `", "zip":"` + zip + `"}`

    if lat == blankNumber || lng == blankNumber {
    	geocode = ""
    }

    var payloadStr = `{"name":"` + merchantName + `"`

    payloadStr = payloadStr + `, "category":` + categoriesString

    if len(street_number) > 0{
    	payloadStr = payloadStr + `,"address":` + address
    }
    
    if len(geocode) > 0{
    	payloadStr = payloadStr + `, "geocode": ` + geocode
    } 
    
    payloadStr = payloadStr + `}`

    var jsonStr = []byte(payloadStr)
    req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
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
