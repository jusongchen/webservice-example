package main

import (
	"log"

	restful "github.com/emicklei/go-restful/v3"
)

//Payload declares the payload struct
type Payload struct{ Name string }

func helloYaml(req *restful.Request, resp *restful.Response) {

	payload := Payload{}

	err := req.ReadEntity(&payload)
	if err != nil {
		resp.WriteErrorString(500, err.Error())
		return
	}
	log.Printf("helloYaml:got payload:%v", payload)

	err = resp.WriteEntity(payload)
	if err != nil {
		resp.WriteErrorString(500, err.Error())
		return
	}
}
