-----------------------------------------------------------------

To edit this file (and the *.go files) inside Powershell, you
must have WSL2 installed and properly running, then, e.g.:

> bash -c "pico README.md"

-----------------------------------------------------------------

To run Server code:

> go run server.go

Then, open a browser and navigate to:

> http://localhost:8080

> http://localhost:8080/hello

-----------------------------------------------------------------

To prepare (if not done previously) the Quote code: 

> go mod init quote

Then, to run it:

> go run quote.go

-----------------------------------------------------------------
Reference: https://golang.org/doc/tutorial/getting-started
-----------------------------------------------------------------
