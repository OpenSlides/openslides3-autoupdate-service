dev-cert:
	mkdir -p ./cert/
	mkcert -cert-file ./cert/autoupdate.pem -key-file ./cert/autoupdate-key.pem localhost
