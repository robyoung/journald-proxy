
build:
	docker build -t journald-proxy .
	docker run --rm --entrypoint cat journald-proxy:latest /usr/local/src/journald-proxy > journald-proxy
	chmod a+x journald-proxy

