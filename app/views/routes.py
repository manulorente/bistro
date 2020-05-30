# -*- coding: utf-8 -*-
"""
Created on Fri May 29 21:24:55 2020

@author: Manuel
"""
from app import app
from flask import render_template, flash, redirect
from app.forms import LoginForm
import datetime

#Run the home() function when someone accesses the root URL ('/') of the server
@app.route("/")
@app.route("/home")
@app.route("/index")
def index():
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

@app.route('/login')
def login():
    form = LoginForm()
    if form.validate_on_submit():
        flash('Login requested for user {}, remember_me={}'.format(
            form.username.data, form.remember_me.data))
        return redirect('/index')
    return render_template('login.html', title='Sign In', form=form)