sudo apt install protobuf-compiler
export GO111MODULE=on  # Enable module mode

go get -u google.golang.org/protobuf
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --version
export PATH="$PATH:$HOME/.local/bin"

# mkdir core
protoc --go_out=./core --go-grpc_out=./core node.proto
