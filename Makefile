app-cover:
	go test ./cmd/autoupdate/ --coverpkg=./... --coverprofile=app.cover
	go tool cover -html=app.cover

build-dev:
	docker build . --target development --tag os3-autoupdate-dev
