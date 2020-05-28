#Import libraries
from flask import Flask, render_template
import datetime

#Create Flask object
app = Flask(__name__)

#Run the home() function when someone accesses the root URL ('/') of the server
@app.route('/')
def home():
	now = datetime.datetime.now()
	timeString = now.strftime("%Y-%m-%d %H:%M")
	templateData = {
      'message' : 'This is the landing page',
      'time': timeString
      }
	return render_template('index.html', **templateData)

# Creating a new static page
@app.route('/example')
def static_page():
	now = datetime.datetime.now()
	timeString = now.strftime("%Y-%m-%d %H:%M")
	templateData = {
      'message' : 'Example static page',
      'time': timeString
      }
	return render_template('index.html', **templateData)
	
# Passing objects to html
@app.route('/example/<message>')
def dynamic_page(message):
	now = datetime.datetime.now()
	timeString = now.strftime("%Y-%m-%d %H:%M")
	templateData = {
      'message' : message,
      'time': timeString
      }
	return render_template('index.html', **templateData)
	
if __name__ == '__main__':
	# Starts listen on port 80
    app.run(debug=True, port=80, host='0.0.0.0')