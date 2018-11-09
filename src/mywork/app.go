package main

import (
	 . "mywork/router"
	 "github.com/braintree/manners"
)

func main() {
	router := InitRouter()
	//router.Run("0.0.0.0:8888")
	manners.ListenAndServe("0.0.0.0:8888",router)
}
