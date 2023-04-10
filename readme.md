# OBD Connection API

This is an API for the OBD connection application. It uses Buf build and Connect build, and the endpoint definitions can be found at [buf.build/ahmeddarwish/obd/docs/main:obd.v1](https://buf.build/ahmeddarwish/obd/docs/main:obd.v1).

The base URL for connecting to the REST endpoints is `obd.exploremelon.com/`.

All requests are POST requests with the `baseurl/{servicename}` format, and accept JSON bodies like the request message on the documentation link. Responses are also returned in JSON format.

## Dependencies

This project uses SQLC and PostgreSQL for its database.

## Usage

To use this project, clone the repository and run the following commands:

```bash
# Build the Docker image
docker build -t obd-connection-api .

# Run the Docker container
docker run -p 9091:9091 obd-connection-api