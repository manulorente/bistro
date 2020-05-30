**Instructions to run/debug app branch**
================================================
## Taxonomy
This repository consists in a set of folders with specific purpose described below:
| File      	| Description |
| ----------- 	| ----------- |
| *readme.md:*  | This file       |
| *changelog.md:*  | Changes tracking file       |
| *doc:*  | Stuff documentation       |
| *tags:*  | Releases tagged versions |
| *trunk:*  | Main development source files       |
| *branches:*  | Copy of trunk source files to implement in parallel extra functionalities      |

## Getting started
0.	Clone/download GitHub repository on your project desired location

1. 	Download latest **Python3.X** version from official page \ 
	*Select Add Python to Windows Path checkbox during installation* \
		https://www.python.org/downloads/ \ 
	*Add Python to Windows Path just in case not added by default* \
		https://datatofish.com/add-python-to-windows-path/ \
		
2. 	Check Python installation typing in a new command line:\
	``python -m pip install --upgrade pip``
	
3.	Install **Flask** and all project dependencies\
	``python install Flask``\
	``python install flask_wtf``
	
4.	Change directory to your **devBranch** in **branches** folder\
	``cd [YOUR_PROJECT_PATH]\branches\devBranch``
	
5.	Override FLASK variable to easily deploy server:\
	``set FLASK_APP=webServer.py``
	
6.	Run server just typing in command line:\
	``flask run``
	
7.	Open web browser to contact the server which address is *http://127.0.0.1:5000/* 

8.	From now on you are able to change the code and user any browser as web viewer.

#### Do not forget to commit and push changes once your development is finished
	