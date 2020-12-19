-----------------------------------------------------------------

Pre-requisites / other assumptions / my environment:
- Windows 10
- WSL2 installed with Ubuntu 20.04 running
- Windows Terminal
- Docker Desktop
- Golang
- Commands executed in Powershell, in the root folder.

Root folder is .\02-kubernetes\

-----------------------------------------------------------------

To edit files (e.g. *.md, *.go, Dockerfile) inside the console:

> bash -c "pico README.md"

-----------------------------------------------------------------

To run Server code:

> go run ./server/server.go

Then, open a browser and navigate to:

> http://localhost:8080 (will say "not found")
> http://localhost:8080/hello (will say "hello")

To build it into an executable (in Windows):

> go build ./server
> mv -Force ./server.exe ./server/server.exe

To run it:

> ./server/server

-----------------------------------------------------------------

To containerize it:

> docker build . -t ropinheiro/kubernetes-tutorial

Check it:

> docker images

-----------------------------------------------------------------
Reference:
https://medium.com/swlh/zero-to-kubernetes-in-5-mins-dcff81b4508
-----------------------------------------------------------------
