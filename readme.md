# South-facing roof detector

## Technologies used
- Go + OpenCV bindings
- NextJS + RSC
- Mantine UI + TailwindCSS
- gRPC API

## Preparing to run

1) Make sure [OpenCV](https://opencv.org/get-started/) and the [Proto compiler](https://grpc.io/docs/protoc-installation/) are installed on the system.
2) Fetch dependencies:
```bash
cd backend && go mod download
```
```bash
cd frontend && npm install
```
3) Re-generate the gRPC & Protobuf files:
```bash
cd backend && make
```
```bash
cd frontend && npm run codegen
```
## Option 1: Run one-shot analysis
```bash
cd backend && go run cmd/analyze/main.go data/1.png data/1.blob mask.png
```
## Option 2: Run the server and the frontend

1) Run the backend:
```bash
cd backend && BLOBSTORE_ROOT_DIR=data BLOBSTORE_TYPE=FILESYSTEM LOG_LEVEL=DEBUG LOG_MODE=development LOG_OUTPUT=stdout PORT=8080 go run cmd/serve/main.go
```
2) Run the frontend:
```bash
cd frontend && PROTO_PATH=../shared/roof.proto NEXT_PUBLIC_HOST="localhost:8080" npm run dev
```