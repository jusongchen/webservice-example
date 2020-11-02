# webservice-example

## test manually

go build -o /tmp/yaml-hello-world
/tmp/yaml-hello-world &
curl -d "name: Merlin"  -H "Accept: application/json"   -H 'Content-Type: application/x-yaml' localhost:8080/helloyaml
pkill -9 -f "yaml-hello-world"