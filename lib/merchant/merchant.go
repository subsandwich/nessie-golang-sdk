package merchant

import(
    "fmt"
    "net/http"
    "bytes"
    "io/ioutil"
    "strconv"
    "../shared"
)

var baseUrl = "http://api.reimaginebanking.com/merchants"
var apiKey = shared.ApiKey

