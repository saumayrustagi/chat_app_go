all: server/server client/client

server/server: server/server.go helper/helper.go
	cd server && go build server.go

client/client: client/client.go helper/helper.go
	cd client && go build client.go

clean:
	rm -f server/server client/client
