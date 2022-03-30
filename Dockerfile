FROM golang:latest

ENV PORT 8000

# unset go path to avoid failed build
ENV GOPATH=

# copy contents of current folder into docker
COPY . .

RUN go mod download
RUN go build

CMD ["./go-project"]