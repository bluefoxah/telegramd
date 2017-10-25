./codegen_encode_decode.py -i scheme.tl -o ./
gofmt -w codec_schema.tl.pb.go
mv codec_schema.tl.pb.go ..

