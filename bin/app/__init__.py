# -*- coding: utf-8 -*-
"""
Created on Fri May 29 21:22:02 2020

@author: Manuel
"""

from flask import Flask
from app.config import Config

app = Flask(__name__)
app.config.from_object(Config)

from app import routes
