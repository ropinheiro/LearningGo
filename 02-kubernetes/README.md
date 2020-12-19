-----------------------------------------------------------------

To edit this file (and the *.go files) inside Powershell, you
must have WSL2 installed and properly running, then, e.g.:

> bash -c "pico README.md"

-----------------------------------------------------------------

To run Server code:

> go run server.go

Then, open a browser and navigate to:

> http://localhost:8080 (will say "not found")
> http://localhost:8080/hello (will say "hello")

To build it into an executable (in Windows):

> go build .
> mv 02-kubernetes.exe server.exe

To run it:

> ./server

-----------------------------------------------------------------
Reference:
https://medium.com/swlh/zero-to-kubernetes-in-5-mins-dcff81b4508
-----------------------------------------------------------------
