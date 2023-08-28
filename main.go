package main

import (
	"context"
	"fmt"

	swagger "github.com/fairDataSociety/fairOS-dfs-clients/go-client"
)

func main() {
	// create FairOS-dfs client
	cnf := swagger.NewConfiguration()
	cnf.BasePath = "http://localhost:9090"
	client := swagger.NewAPIClient(cnf)

	// login
	loginRequest := swagger.CommonUserLoginRequest{
		Password: "passwordpassword",
		UserName: "usernameusername",
	}
	apiResp, httpResponse, err := client.UserApi.UserLoginV2(context.TODO(), loginRequest)
	if err != nil {
		panic(err)
	}
	fmt.Println(apiResp)
	// get cookie
	cookie := httpResponse.Header.Get("Set-Cookie")

	// get pod list
	podList, podListResp, err := client.PodApi.PodListHandler(context.Background(), cookie)
	if err != nil {
		panic(err)
	}
	fmt.Println(podList)
	fmt.Println(podListResp)
}
