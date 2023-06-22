.PHONY: build run
build:
	docker run --rm -v `pwd`:/go/src/go-filter -w /go/src/go-filter \
		golang:1.20 \
		go build -v -o libsimple.so -buildmode=c-shared -buildvcs=false .
run:
	docker run --rm -v `pwd`/envoy.yaml:/etc/envoy/envoy.yaml \
		-v `pwd`/libsimple.so:/etc/envoy/libsimple.so \
		-p 10000:10000 \
		-e "GODEBUG=cgocheck=0" \
		envoyproxy/envoy:contrib-v1.26.2 \
		envoy -c /etc/envoy/envoy.yaml -l trace
trace:
	docker run --rm -v `pwd`/envoy.yaml:/etc/envoy/envoy.yaml \
		-v `pwd`/libsimple.so:/etc/envoy/libsimple.so \
		-p 10000:10000 \
		envoyproxy/envoy:contrib-v1.26-latest \
		envoy -c /etc/envoy/envoy.yaml -l trace
dclean:
	docker run -rf $(docker ps -qa)
