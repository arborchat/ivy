default: test

test:
	go run main.go

server:
	arbor

nc:
	nc localhost 7777
