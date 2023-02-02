package main

import "planigo/api"

func main() {
	//token := utils.GenerateJWT(&utils.TokenPayload{ID: "123"})
	//
	//println(token)
	//
	//p, err := utils.VerifyJWT(token)
	//if err != nil {
	//	println(err.Error())
	//}
	//println(p.ID)
	api.Start()
}
