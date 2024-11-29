# minimal-server

A distroless web server container for testing that serves the <code>TEXT</code> environment variable on the specified <code>PORT</code>.

### Example Usage

```bash
docker run -p 3000:3000 -e TEXT="howdy partner" -e PORT=3000 doron/minimal-server