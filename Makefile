run: build
	docker run --name atlas-dns -d -p 9000:9000 atlas-dns

build: rm
	docker build . -t atlas-dns

rm: stop
	-docker rm atlas-dns

stop:
	-docker stop atlas-dns

# run_test:
# 	cd tests
# 	go test ./... -v -coverpkg=./...