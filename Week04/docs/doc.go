package docs

// 生成 grpc pb文件，generate 需要结尾空行
//go:generate protoc -I ../third_party -I ../api/order/v1 --go_out ../api/order/v1/ --go_opt paths=source_relative --go-grpc_out ../api/order/v1/ --go-grpc_opt paths=source_relative ../api/order/v1/*.proto

// 生成 grpc-gateway 文件
//go:generate protoc -I ../third_party -I ../api/order/v1 --grpc-gateway_out ../api/order/v1/ --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative ../api/order/v1/*.proto
