# Lina backend

## About
This is the backend API for Lina, a project for NASA Space Apps Challenge 2024, which purpose is to provide an endpoint for the frontend to chat with the AI model. The model must advertise the farmers about possible weather changes and how to prevent them.

## Disclaimer
This project is a prototype and is not intended to be used in production. The model used is a simple one and is not accurate, it uses hardcoded data to simulate a real model.

## How to run
To run the code, you must have Golang 1.20 or higher installed in your machine. 

The environment variables are:
    - `PORT`: The port where the server will run. Default is 8080.
    - `GEMINI_API_KEY`: The API key for the Gemini API. It is mandatory.

To run the code, you must run the following commands:
```bash
go run cmd/main.go
```