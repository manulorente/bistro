# -*- coding: utf-8 -*-
"""
Created on Fri May 29 21:20:10 2020

@author: Manuel
"""
import os

class Config(object):
    SECRET_KEY = os.environ.get('SECRET_KEY') or 'you-will-never-guess'