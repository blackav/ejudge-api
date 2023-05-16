all : main

main : main.go go.mod go.sum
	~/go/bin/swag init
	go build main.go

clean : 
	-rm -f main
