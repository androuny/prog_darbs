# My project

This project contains 2 repos - client and server

## Client

Implemented in [golang](https://go.dev/).
This client calculates Pi using [Monte Carlo method](https://www.geeksforgeeks.org/estimating-value-pi-using-monte-carlo/) and sends results to the server

**How to run client app**

``cd client``

``go mod tidy``

``go mod run main.go``

## Server

Implemented in Python using [flask](https://flask.palletsprojects.com/en/3.0.x/).
This server is used to collect data from clients and to store it in a mongodb database, i used [mongodb atlas free tier](https://www.mongodb.com/atlas/database)

**How to run server app**

``cd server``

``pip install -r requirements.txt``

``python app.py``
