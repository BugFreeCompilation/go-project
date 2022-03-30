# go-project

A small project to practise go. It uses sqlite3 and adds/removes DVD titles.

To run a small test, go into `/service` and run `go test`.

```
Create docker:
docker build -t go_project .

Run docker: (os port 8080 will be redirected to docker port 8000)
docker run -p8080:8000 go_project
```

```
Install postman:
yay -S postman-bin

Post request can be done with postman (X->Body). Send this to `localhost:8000/data`:
{
    "TITLE": "Google the Movie"
}
```

```
Creating go.mod and go.sum:
go mod init
go mod tidy
```