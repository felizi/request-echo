# request-echo
Simple request echo for troubleshooting/debuging

### Prerequisites
This project requires **Golang**:

- [Golang](https://golang.org/)

### Compiling
For compile, execute the command bellow then execute generated binary
```
go build main.go
```

### Running
You can pass parameters for executation:
* Port
```
-p "8888"
```
* Sleep before response (milliseconds)
```
-s 10
```

For running without compiling, execute the command bellow
```
go run main.go
```
