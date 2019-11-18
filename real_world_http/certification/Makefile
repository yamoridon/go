all: ca.crt server.crt client.crt

clean:
	rm -f *.key *.csr *.crt *.srl

ca.key:
	openssl genrsa -out $@ 2048

ca.csr: ca.key openssl.cnf
	openssl req -batch -new -sha256 -key $< -out $@ -config openssl.cnf -subj "/C=JP/ST=Tokyo/L=Itabashi/O=example.com/emailAddress=webmaster@example.com"

ca.crt: ca.csr ca.key openssl.cnf
	openssl x509 -req -days 365 -in $< -sha256 -out $@ -signkey ca.key -extfile openssl.cnf -extensions CA

server.key:
	openssl genrsa -out $@ 2048

server.csr: server.key openssl.cnf
	openssl req -batch -new -nodes -sha256 -key $< -out $@ -config openssl.cnf -subj "/C=JP/ST=Tokyo/L=Itabashi/O=example.com/emailAddress=webmaster@example.com/CN=localhost"

server.crt: server.csr ca.crt ca.key openssl.cnf
	openssl x509 -req -days 365 -in $< -sha256 -out $@ -CA ca.crt -CAkey ca.key -CAcreateserial -extfile openssl.cnf -extensions Server

client.key:
	openssl genrsa -out $@ 2048

client.csr: client.key openssl.cnf
	openssl req -batch -new -nodes -sha256 -key $< -out $@ -config openssl.cnf -subj "/C=JP/ST=Tokyo/L=Itabashi/O=example.com/emailAddress=webmaster@example.com"

client.crt: client.csr ca.crt ca.key openssl.cnf
	openssl x509 -req -days 365 -in $< -sha256 -out $@ -CA ca.crt -CAkey ca.key -CAcreateserial -extfile openssl.cnf -extensions Client

.PHONY: all clean
