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
