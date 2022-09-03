# golang-webserver-performance

Project for fun to compare performance of GRPC and HTTP servers.
Can also function as a demo project.

## Takeaways

### Structure of code
- domain driven design oriented

### Code generation 
- GRPC server
- HTTP server
- clients are not generated yet but can/will be

### Load testing tools for HTTP and GRPC
- autocannon
- ghz

### Automating flows with go-task
- see Taskfile.yml

### Usage of trace
- to measure and compare performance gotrace is used (see main.go)