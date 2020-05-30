# -*- coding: utf-8 -*-
# To run --> flask run -h 192.168.1.144
import os
import locale
from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from flask_login import LoginManager
import logging
from app.config import Config

locale.setlocale(locale.LC_TIME, locale.getlocale())

app = Flask(__name__)

app.config.from_object(Config)

from app import models
from app import views

