set GOPATH=%cd%/../
set GOBIN=%GOPATH%/bin
PATH=%PATH%:%GOPATH%:%GOBIN%
set PATH
go env -w GOBIN=%cd%/../bin
go env -w GOPATH=%cd%/../src
Echo Press any key to continue...
PAUSE