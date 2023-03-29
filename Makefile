build:
	docker build --no-cache --tag rsa-speed-test .

run-tiny-container:
	docker build --no-cache --tag rsa-speed-test .
	docker run -it --memory=128m --cpus=0.25 rsa-speed-test 4096

run-normal-container:
	docker build --no-cache --tag rsa-speed-test .
	docker run rsa-speed-test 4096