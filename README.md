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


## Benchmark
With this command:
```
docker run -it --memory=128m --cpus=0.25 rsa-speed-test 4096
```
The output is:
```
RSA 4096 BENCHMARK
managed to encrypt 3690 keys in 10s; encryption rate 369 keys/s 
managed to decrypt 81 keys in 10s; decryption rate 8 keys/s 
```
