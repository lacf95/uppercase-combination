FROM golang:1.17.1-alpine3.14

WORKDIR /var/app

COPY . ./

CMD ["go", "test", "-bench=."]
