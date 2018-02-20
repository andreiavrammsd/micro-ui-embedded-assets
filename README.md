## Micro UI

Small web app with all assets embedded into the binary

### Development and build 

[sudo apt install go-bindata -y](https://github.com/jteeuwen/go-bindata)

go get -t

[./run.sh](./run.sh)

### Production build and deploy

[./deploy.sh](./deploy.sh)

### Run

./micro-ui (default addr = :1234)

./micro-ui -addr=host:8080

./micro-ui -addr=:8081

### Access

http://host:port/
