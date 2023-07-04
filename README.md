# gRPC

## Usage
Run `./build.sh` in root directory

`go run server/main.go` to run the server


### Command to make request (install `grpcurl` beforehand):
`grpcurl -plaintext -d '{"variable_name": "SHELL"}' localhost:8080 EnvVariable/GetEnvironmentVariable`