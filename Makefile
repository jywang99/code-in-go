free:
	go build -o bin/free src/*.go
pro:
	go build -o bin/pro -tags pro src/*.go
enterprise:
	go build -o bin/enterprise -tags pro,enterprise src/*.go
