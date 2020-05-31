![](https://github.com/manulorente/bistro/blob/master/resources/logo.jpeg?raw=true)

**Bistro**
================================================
This web application is build in Angular + Node.js.  
This front-end will make requests against Golang backend which is powered by NoSQL.

## Requirements
* Golang 1.14+  
* Node.js 12.17+
* Angular 9.1+

## Taxonomy
| File      		| Description |
| ----------- 		| ----------- |
| *readme*  		| This file |
| *changelog*  		| Changes tracking file |
| *resources*  		| Documentation|
| *scripts*  		| All scripts to build the project|
| *app*  			| Source code|


## Getting started
0.	Clone/download proyect **branch** on your desired location

1. 	Install all requirements to deploy the application

2. Set Golan workspace for this project running `virtualEnvironment.bat` script. 

3. Change directory to *server* folder:  
``cd [YOUR_PROJECT_PATH]\server``

4. Download all of our Golang dependencies. Using the Command Prompt or Terminal, execute the following:  
`go get github.com/couchbase/gocb`  
`go get github.com/gorilla/handlers`  
`go get github.com/gorilla/mux`  
`go get github.com/satori/go.uuid`  
The above will get the Couchbase Go SDK, a library for making RESTful API endpoints easier to create, a library for handling cross origin resource sharing (CORS) and a library for generating unique id values that will represent NoSQL document keys.
	
5.	Launch unitary tests to check everythig is ok:  
`go test -v`
		
6.	Open web browser to view server changes in *http://127.0.0.1:8080/* 

7.	From now on you are able to change the code and user any browser as web viewer.

## Serving content over HTTPS
It is needed to generate self-signed certs locally using openssl:  
`openssl genrsa -out server.key 2048`  
`openssl ecparam -genkey -name secp384r1 -out server.key`  
`openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650`  
	