#get a base image
FROM golang:1.16-buster

 

WORKDIR /go/src/app
COPY . .

RUN go get -d -v
RUN go build -v
RUN echo $PATH
RUN ls
RUN pwd

CMD ["./todo-project"]