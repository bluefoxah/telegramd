./codegen_encode_decode.py -i scheme.tl -o ./codec_schema.tl.pb.go
gofmt -w codec_schema.tl.pb.go
mv codec_schema.tl.pb.go ..

