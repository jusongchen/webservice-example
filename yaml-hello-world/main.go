package main

import (
	"io"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
	"github.com/jusongchen/webservice-example/yaml-hello-world/restyaml"
)

const (
	pathYAML = "/helloyaml"
)

func main() {
	ws := new(restful.WebService)

	restful.RegisterEntityAccessor(restyaml.MediaTypeApplicationYaml, restyaml.NewYamlReaderWriter(restyaml.MediaTypeApplicationYaml))
	ws.Path("/").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/hello").To(hello))
	ws.Route(ws.POST(pathYAML).Consumes(restyaml.MediaTypeApplicationYaml).To(helloYaml))

	restful.Add(ws)
	port := ":8080"
	log.Println("starting server on port" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "world")
}
