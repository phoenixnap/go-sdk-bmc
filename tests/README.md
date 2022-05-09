# Tests

## Setup

To run tests you'll need to run an instance of [mockserver](https://github.com/mock-server/mockserver) running at port 1080. This can be done using [Docker](https://www.docker.com/):
- [Install docker](https://www.docker.com/get-started/)
- Start `mockserver`:
   ```
   docker run --name mockserver -d --rm -p 1080:1080 mockserver/mockserver:5.13.0
   ```

To view mockserver logs:
```
docker logs -f mockserver
```

Alternatively, you can also run an instance of [portainer](https://www.portainer.io/) by running:
```
docker run -d -p 9000:9000 --restart always -v /var/run/docker.sock:/var/run/docker.sock -v /opt/portainer:/data portainer/portainer
```
This will provide a GUI at [localhost:9000](localhost:9000) where one can start, stop, and view logs for `mockserver`.

## Running

Each go script in this directory contains tests for a specific api.
Tests can be executed by running each script. Example:
```
go run auditapi_test.go
```