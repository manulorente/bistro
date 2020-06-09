![](https://github.com/manulorente/bistro/blob/master/resources/logo.jpeg?raw=true)

**Bistro**
================================================
This web application is build following MVC layered architecture. 
This front-end will make requests against Golang backend which is powered by NoSQL.

### Tech stack
* Go 1.14+  		- Backend language
* React v16+ 		- Frontend technology
* Auth0 			- User authentication
* Postgres SQL 		- Database

### Info
| Item     		| Description |
| ----------- 	| ----------- |
| *Author*  	| Manuel Lorente <manloralm@outlook.com> |
| *Copyright*  	| © 2020, Manuel Lorente.       |
| *Date*  		| 2020-June |
| *Version*  	| 0.0.0 |


Project overview
==============
### Structure
| File      		| Description |
| ----------- 		| ----------- |
| *README*  		| This file |
| *resources*  		| Documentation |
| *scripts*  		| Useful batch scripts|
| *app*  			| Source code |


### Getting started
0. Clone/download proyect **branch** on your desired location

1. Set Go environment for this project running `goSetup.bat` script. 

2. Change directory to *backend* folder:  
``cd [YOUR_PROJECT_PATH]\backend``

3. Compile server program using the Makefile
`make build`  

4. Run server located in app folder
`bistro.exe`  

5. Open web browser to view server changes in *http://127.0.0.1:3000/* 


### Serving content over HTTPS
It is needed to generate self-signed certs locally using openssl:  
`openssl genrsa -out server.key 2048`  
`openssl ecparam -genkey -name secp384r1 -out server.key`  
`openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650`  

Changelog
==============
0.0.0  
		*Initial version*