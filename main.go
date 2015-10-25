package main

import(
	//comment out if not used
    "./lib/account"
	//"./lib/atm"
    //"./lib/branch"
)

func main() {

    //Demo Code for Requests (uncomment to run)

    //======================Account=======================
    //account.GetAllAccounts()
    //account.GetAccountWithId("56241a13de4bf40b1711287b")
	//account.GetAccountsOfCustomer("56241a12de4bf40b17111f9d")
    //account.CreateAccount("56241a12de4bf40b17111f9d", "Checking", "Account to Delete", 44444, 44444, "")
    //account.CreateAccount("56241a12de4bf40b17111f9d", "Checking", "Account to Delete", 7777, 77777, "8888444455551111")
    //account.UpdateAccount("562cf1190afebb140066cd81", "Iron Man's Account", "")
    //account.UpdateAccount("562cf1190afebb140066cd81", "Mario's Account", "1234567812345678")
    //account.DeleteAccount("562b073f0afebb140066cd58")

    //=======================ATM=========================
    //atm.GetAllBranches(38.9283, -77.1753, 1)
    //atm.GetATMInfo("56241a12de4bf40b17111c65")

    //======================Branch======================= (DONE)
    //branch.GetAllBranches()
    //branch.GetBranchWithId("56241a12de4bf40b17111eb2")
}
