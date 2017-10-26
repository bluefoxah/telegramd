./codegen_rpc_server.py account -i ./scheme.tl -o ./server/account_service_impl.go
gofmt -w ./server/account_service_impl.go
./codegen_rpc_server.py auth -i ./scheme.tl -o ./server/auth_service_impl.go
gofmt -w ./server/auth_service_impl.go
./codegen_rpc_server.py bots -i ./scheme.tl -o ./server/bots_service_impl.go
gofmt -w ./server/bots_service_impl.go
./codegen_rpc_server.py channels -i ./scheme.tl -o ./server/channels_service_impl.go
gofmt -w ./server/channels_service_impl.go
./codegen_rpc_server.py contacts -i ./scheme.tl -o ./server/contacts_service_impl.go
gofmt -w ./server/contacts_service_impl.go
./codegen_rpc_server.py langpack -i ./scheme.tl -o ./server/langpack_service_impl.go
gofmt -w ./server/langpack_service_impl.go
./codegen_rpc_server.py messages -i ./scheme.tl -o ./server/messages_service_impl.go
gofmt -w ./server/messages_service_impl.go
./codegen_rpc_server.py payments -i ./scheme.tl -o ./server/payments_service_impl.go
gofmt -w ./server/payments_service_impl.go
./codegen_rpc_server.py phone -i ./scheme.tl -o ./server/phone_service_impl.go
gofmt -w ./server/phone_service_impl.go
./codegen_rpc_server.py photos -i ./scheme.tl -o ./server/photos_service_impl.go
gofmt -w ./server/photos_service_impl.go
./codegen_rpc_server.py stickers -i ./scheme.tl -o ./server/stickers_service_impl.go
gofmt -w ./server/stickers_service_impl.go
./codegen_rpc_server.py updates -i ./scheme.tl -o ./server/updates_service_impl.go
gofmt -w ./server/updates_service_impl.go
./codegen_rpc_server.py upload -i ./scheme.tl -o ./server/upload_service_impl.go
gofmt -w ./server/upload_service_impl.go
./codegen_rpc_server.py users -i ./scheme.tl -o ./server/users_service_impl.go
gofmt -w ./server/users_service_impl.go
