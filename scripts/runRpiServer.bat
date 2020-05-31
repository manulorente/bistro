pscp -pw "raspberry" -r %cd%/..bin/app/ pi@192.168.1.144:rpiWebServer/app
pscp -pw "raspberry" -r %cd%/..bin/app/webServer.py pi@192.168.1.144:rpiWebServer/
plink -v pi@192.168.1.144 -pw "raspberry"  (cd rpiWebServer; export FLASK_APP=run.py; flask run -h 192.168.1.144)
Echo Press any key to continue...
PAUSE