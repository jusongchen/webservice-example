package restyaml

import (
	"io"
	"io/ioutil"

	"log"

	"github.com/emicklei/go-restful/v3"
	"gopkg.in/yaml.v3"
)

// MediaTypeApplicationYaml is a Mime Type for YAML.
const MediaTypeApplicationYaml = "application/x-yaml"

// YamlReaderWriter implements EntityReaderWriter for YAML objects to be used by restful.
type YamlReaderWriter struct {
	contentType string
}

// NewYamlReaderWriter creates new instance.
func NewYamlReaderWriter(contentType string) restful.EntityReaderWriter {
	return YamlReaderWriter{contentType: contentType}
}

func closeWithErrHandle(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Println("Unable to close resource: ", err)
	}
}

// Read a serialized version of the value from the request.
// The Request may have a decompressing reader. Depends on Content-Encoding.
func (e YamlReaderWriter) Read(req *restful.Request, v interface{}) error {
	defer closeWithErrHandle(req.Request.Body)
	bytes, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		return err
	}
	got := string(bytes)
	_ = got
	err = yaml.Unmarshal(bytes, v)
	return err
}

// Write a serialized version of the value on the response.
// The Response may have a compressing writer. Depends on Accept-Encoding.
// status should be a valid Http Status code
func (e YamlReaderWriter) Write(resp *restful.Response, status int, v interface{}) error {
	bytes, err := yaml.Marshal(v)
	if err != nil {
		return err
	}

	resp.WriteHeader(status)
	_, err = resp.Write(bytes)
	return err
}
