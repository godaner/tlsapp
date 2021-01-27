build:buildclientwin buildclientamd buildserverwin buildserveramd clearbuild
buildclientwin:
	-rm ./bin/cwin64.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/cwin64.exe ./cliapp/client.go
	-upx -9 ./bin/cwin64.exe
buildclientamd:
	-rm ./bin/clinux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/clinux ./cliapp/client.go
	-upx -9 ./bin/clinux
buildserverwin:
	-rm ./bin/swin64.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/swin64.exe ./serapp/server.go
	-upx -9 ./bin/swin64.exe
buildserveramd:
	-rm ./bin/slinux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/slinux ./serapp/server.go
	-upx -9 ./bin/slinux
clearbuild:
	-rm *.upx

cert:
	-go get -u github.com/square/certstrap
	certstrap init --common-name "ca" --expires "20 years"
	certstrap request-cert -cn server -ip 127.0.0.1 -domain "*.example.com"
	certstrap sign server --CA ca
	certstrap request-cert -cn client
	certstrap sign client --CA ca
certclear:
	rm -r out/