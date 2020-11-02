package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/yaml.v2"

	"github.com/emicklei/go-restful/v3"
	"github.com/jusongchen/webservice-example/yaml-hello-world/restyaml"
	"github.com/stretchr/testify/require"
)

func TestWebService_Yaml(t *testing.T) {

	ws := new(restful.WebService)

	restful.RegisterEntityAccessor(restyaml.MediaTypeApplicationYaml, restyaml.NewYamlReaderWriter(restyaml.MediaTypeApplicationYaml))

	ws.Path("/").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.POST(pathYAML).Consumes(restyaml.MediaTypeApplicationYaml).To(helloYaml))

	// create test server
	handler := restful.NewContainer()
	handler.Add(ws)

	ts := httptest.NewServer(handler)
	defer ts.Close()

	payload := []byte(`name: Merlin`)

	want := Payload{}
	err := yaml.Unmarshal(payload, &want)
	require.NoError(t, err)

	resp, err := http.Post(ts.URL+pathYAML, "application/x-yaml", bytes.NewReader(payload))
	t.Logf("Post response:%v", resp)

	require.NoError(t, err)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Post reqeust failed, status code: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	log.Printf("got reponse body:%s", string(data))

	got := Payload{}
	err = json.Unmarshal(data, &got)
	require.NoError(t, err)
	require.Equal(t, want, got)

}
