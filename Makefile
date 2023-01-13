export CGO_ENABLED=1

all: usage 

usage: clean
	env GOOS=linux GOARCH=arm GOARM=6 go build  -o usage ./main.go

macos: clean
	env GOOS=darwin GOARCH=amd64 go build -o usage ./main.go

clean:
	rm -f usage
