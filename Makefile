dev-cert:
	mkdir -p ./cert/
	mkcert -cert-file ./cert/autoupdate.pem -key-file ./cert/autoupdate-key.pem localhost

app-cover:
	go test ./cmd/autoupdate/ --coverpkg=./... --coverprofile=app.cover
	go tool cover -html=app.cover
