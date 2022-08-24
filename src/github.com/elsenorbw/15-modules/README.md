
 Modules.. are weird :)


 So there's some horrible logic around including modules, really not a fan so far..

 To get a package in one folder to be seen in another had to do : 

 1982  go mod init github.com/elsenorbw/greetings
 1983  ll
 1984  echo $GOPATH
 1985  ll
 1986  cd hello/
 1987  go mod init example.com/hello
 1988  ll
 1989  go run hello.go 
 1990  vi ~/.bashrc
 1991  cd development/
 1992  cd go-academy/
 1993  ll
 1994  code .
 1995  exit
 1996  exit
 1997  ll
 1998  go run hello.go 
 1999  go mod edit -replace github.com/elsenorbw/15-modules/greetings=../greetings
 2000  go run hello.go 
 2001  go mod tidy
 2002  go run .
 2003  go run hello.go 
 

I'm sure there's a good reason for all this, just that it's not immediately obvious what that might be.

Was following the guide here : https://go.dev/doc/tutorial/create-module
