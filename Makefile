build:
	go build
docker:
	GOOS=linux go build
	docker build -t cloud-native-go:1.0.0 .
	rm -rf Cloud-Native-Go
clean:
	rm -rf Cloud-Native-Go
