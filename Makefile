build:
	GOOS=linux GOARCH=amd64 go build -o ./.build/www ./*.go
	zip -j ./.build/www.zip ./.build/www *.html.tmpl

run:
	qtc
	go run **.go
