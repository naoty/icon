build:
	go build -o bin/icon

release: build
	tar czvf icon.tar.gz bin/
