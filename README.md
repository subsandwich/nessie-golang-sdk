# Nessie Golang SDK
Provides Wrapper Functions for Nessie API

## Usage
You can follow the code in main.go as examples
1. Change the API Key to the one given to you upon sign up on [api.reimaginebanking.com](http://api.reimaginebanking.com/ "Title")  
	You can change that in /lib/shared/shared.go 
	````
	var ApiKey = {YOUR_API_KEY_HERE}
	````
2. Import the package you need using the path of its location relative to your class
	````
	import(
    	"./lib/withdrawal"
    	"./lib/customer"
	)
	````


# Progress:
* Account: 6/6 - Done
* ATM: 2/2 - In Progress (Need to Implement Paginated Response)
* Bill: 6/6 - Done
* Branch: 2/2 - Done
* Customer: 5/5 - Done
* Deposit: 5/5 - Done
* Merchant: 4/4 - Done
* Purchase: 5/5 - Done
* Transfer: 5/5 - Done
* Withdrawl: 5/5 - Done

Test ApiKey: 00515c501bdde5a46e9e56394c140932


Set your own API Key in lib/shared/shared.go
