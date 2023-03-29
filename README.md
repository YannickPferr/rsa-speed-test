# Benchmark for RSA decryption and encryption

How to run the benchmark:
```
# clone the repo
git clone https://github.com/YannickPferr/rsa-speed-test.git
cd rsa-speed-test
# build the container
make build
# run the speed test for a 4096 bit RSA key
docker run rsa-speed-test 4096 
```
