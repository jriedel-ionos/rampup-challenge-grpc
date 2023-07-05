# gRPC

## Usage
Run `./build.sh` in root directory to generate files

To run the server execute:   
`go run server/server.go -port=8080` (standard port is `8080`) 

and for frontend server:   
`go run frontend/frontend.go -port=8081` (standard port is `8081`)

### Command to make request (install `grpcurl` beforehand):
`grpcurl -plaintext -d '{"variable_name": "SHELL"}' localhost:8080 EnvVariable/GetEnvironmentVariable`

---

To show an environment variable in the browser:   
`localhost:8081/yourvariable` for example `SHELL` --> `localhost:8081/SHELL`