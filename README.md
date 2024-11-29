# minimal-server

This web server takes 2 environment variables TEXT and PORT and serves the value of TEXT on the port specified in value of PORT.

It is a "distroless" container you can use for testing purposes.

### Example Usage

<code>docker run -p 3000:3000 -e TEXT="howdy partner" -e PORT=3000 doron/minimal-server</code>