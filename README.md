# gRPC

## Usage
Run `./build.sh` in root directory to generate files

To run the server execute:   
`go run server/server.go --port=8080` (standard port is `8080`) 

and for frontend server:   
`go run frontend/frontend.go --port=8081` (standard port is `8081`)

### Command to make request (install `grpcurl` beforehand):
`grpcurl -plaintext -d '{"variable_name": "SHELL"}' localhost:8080 EnvVariable/GetEnvironmentVariable`

---

To show an environment variable in the browser:   
`localhost:8081/yourvariable` for example `SHELL` --> `localhost:8081/SHELL`

## Build docker images

### Server
`docker build -t ghcr.io/jriedel-ionos/rampup-challenge-grpc/server:latest -f Dockerfile.server .`
### Frontend
`docker build -t ghcr.io/jriedel-ionos/rampup-challenge-grpc/frontend:latest -f Dockerfile.frontend .`

The variable `TEST` is guaranteed existing (it's defined in the `docker-compose.yml`, bc. some linux environments don't have some env variables.