![](https://github.com/manulorente/bistro/blob/master/resources/logo.jpeg?raw=true)

**Project description**
================================================
This web application is build in Angular + Node.js.  
This front-end will make requests against Golang backend which is powered by NoSQL.

## Requirements
* Couchbase Server 4.5+  
* Golang 1.7+  
* Node.js 6.0+  
	
## Taxonomy
| File      		| Description |
| ----------- 		| ----------- |
| *readme*  		| This file |
| *changelog*  		| Changes tracking file |
| *resources*  		| Documentation and scripts to build the project |
| *fonts*  			| All the fonts in the EOT, SVG, TTF, WOFF, etc, formats |
| *scripts*  		| All the AngularJS code |
| *server*  		| All the Go code |
| *styles*  		| All styles code |
| *views*  			| All the HTML code |


## Getting started
0.	Clone/download proyect **branch** on your desired location

1. 	Install all requirements to deploy the application

3. Check server running in *http://localhost:8091*  
		
2. 	Check Python installation typing in a new command line:  
	``python -m pip install --upgrade pip``

3.	Change directory to **branch** folder:  
	``cd [YOUR_PROJECT_PATH]``

4.	Install all project dependencies from command line:   
	``pip install -r requirements.txt``  
	
5.	Launch script to run server every time you modified a file:  
	``runLocalServer.bat``
		
6.	Open web browser to view server changes in *http://127.0.0.1:5000/* 

7.	From now on you are able to change the code and user any browser as web viewer.

#### Do not forget to commit and push changes and do pull request when your development is finished
	