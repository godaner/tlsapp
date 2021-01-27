build:buildclientwin buildclientamd buildserverwin buildserveramd clearbuild
buildclientwin:
	-rm c.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o c.exe ./cliapp/client.go
	-upx -9 c.exe
buildclientamd:
	-rm c
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o c ./cliapp/client.go
	-upx -9 c
buildserverwin:
	-rm s.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o s.exe ./serapp/server.go
	-upx -9 s.exe
buildserveramd:
	-rm s
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o s ./serapp/server.go
	-upx -9 s
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