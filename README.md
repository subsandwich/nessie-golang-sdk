# Nessie Golang SDK
Provides Wrapper Functions for Nessie API

## Usage

1. Change the API Key to the one given to you upon sign up on [api.reimaginebanking.com](http://api.reimaginebanking.com/ "Title")  
	You can change that in /lib/shared/shared.go 
	````go
	var ApiKey = {YOUR_API_KEY_HERE}
	````

2. Import the package you need using the path of its location relative to your class
	````go
	import(
    	"./lib/withdrawal"
    	"./lib/customer"
	)
	````

3. Use the code as shown in main.go
	````go
	customer.GetAllCustomers()
	````
	Will return something like this:
	````json
	{"_id":"56241a12de4bf40b17111f9d","address":{"street_number":"1112","street_name":"Infinity Loop","city":"Richmond","state":"VA","zip":"22211"},"first_name":"Anna","last_name":"Kendrick"},{"_id":"562ee88d0afebb140066cda2","first_name":"Jennifer","last_name":"Lawrence","address":{"street_number":"32","street_name":"Hunger Rd.","city":"Purham","state":"VA","zip":"11111"}},{"_id":"562ee9970afebb140066cda3","first_name":"Robert","last_name":"Frost","address":{"street_number":"1111","street_name":"Infinity Loop","city":"Richmond","state":"VA","zip":"22211"}},{"_id":"56241a12de4bf40b17111f9c","address":{"city":"Homeworth","street_number":"24500","street_name":"Bowman Street Northeast","state":"Ohio","zip":"44634"},"first_name":"Matthew","last_name":"McConaughey"},{"_id":"56241a12de4bf40b17111f9e","address":{"street_number":"1111","street_name":"Hollywood Ave","city":"San Jose","state":"CA","zip":"90211"},"first_name":"Karen","last_name":"Gillan"}
	````