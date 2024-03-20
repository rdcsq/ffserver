# ffserver

A (quick and dirty) web server to run ffmpeg commands remotely using remote sources.

## Features

- Get streams
- Generate a thumbnail

## Usage

See [docker-compose.yml](docker-compose.yml) for an example.

### Instructions:

- Build the Docker image (`docker build -t ffserver .`)
- Set the `DOMAIN` environment variable as the domain where you're hosting this server (example: `DOMAIN=http://localhost:3000`)
- Create a secure string to be used as an authentication secret for authentication and set it as the `AUTH_SECRET` environment variable.
- Every time the server is started, a new JWT will be printed to the console/terminal. You can use the new one or any other previously generated (as long as they were generated with the same `AUTH_SECRET`)
- On the client, add the JWT in the Authorization header (`Authorization=Bearer {token}`) to all requests. See [API.md](API.md)

## Motivation

Using ffmpeg in more restricted environments (like Vercel), without fiddling with installing it (causing other issues like exceeding the size of a serverless function, being able to use it in the Edge runtime or finding/compiling static builds of FFmpeg with working DNS resolution)
