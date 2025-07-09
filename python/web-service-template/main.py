#!/usr/bin/env python3
"""
Python Daemon Template
A proper daemon implementation following Unix daemon conventions
"""

from flask import Flask
app = Flask(__name__)

@app.route("/")
def hello_world():
    return "<p>Hello, World!</p>"
