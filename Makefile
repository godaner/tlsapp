build:buildclientwin64 buildclientlinux buildserverwin64 buildserverlinux clearupx
buildclientwin64:
	-rm ./bin/cwin64.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/cwin64.exe ./cliapp/client.go
	-upx -9 ./bin/cwin64.exe
buildclientlinux:
	-rm ./bin/clinux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/clinux ./cliapp/client.go
	-upx -9 ./bin/clinux
buildserverwin64:
	-rm ./bin/swin64.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/swin64.exe ./serapp/server.go
	-upx -9 ./bin/swin64.exe
buildserverlinux:
	-rm ./bin/slinux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/slinux ./serapp/server.go
	-upx -9 ./bin/slinux
clearupx:
	-rm *.upx

cert:
	-go get -u github.com/square/certstrap
	certstrap init --common-name "ca" --expires "20 years"
	certstrap request-cert -cn server -ip 127.0.0.1 -domain "*.example.com"
	certstrap sign server --CA ca
	certstrap request-cert -cn client
	certstrap sign client --CA ca
clearcert:
	rm -r out/