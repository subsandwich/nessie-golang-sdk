package main

import(
    "fmt"
    "net/http"
    "bytes"
    "io/ioutil"
    "strconv"
)

func main() {
    url := "http://api.reimaginebanking.com/customers/56241a12de4bf40b17111f96/accounts?key=034a5205a4dce82c13a3de1a7f46711b"
    fmt.Println("URL:>", url)

    var nickname = "Min Pay Account"
    var rewards = 10
    var balance = 10
    rewardsString  := strconv.Itoa(rewards)
    balanceString := strconv.Itoa(balance)
    
    var ttype = "Credit Card"
    //var payloadStr = `{"type":"Credit Card","nickname":"Min Pay Account","rewards":10, "balance":10}`
    var payloadStr = `{"type":"` + ttype + `","nickname":"` + nickname + `","rewards":` + rewardsString + `, "balance":` + balanceString + `}`
    
    fmt.Println(string(payloadStr))
    var jsonStr = []byte(payloadStr)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}