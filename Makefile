proto:
	protoc -I. --go_out=plugins=grpc:. \
	  service/proto/chat.proto
