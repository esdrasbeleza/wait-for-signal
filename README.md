# wait-for-signal

These are two binaries to be used when an application needs to wait for some other task to finish before it starts, probably in different hosts.

When I created these tools, I intended to use them in different Docker containers executed by `docker-compose`. An application needed to wait for another container to be fully provisioned before it was able to start.

This situation is similar to the ones handled by [wait-for-it](https://github.com/vishnubob/wait-for-it), but instead of waiting to a service to be running and listening in some port, I want a task A to be finished (like a database provisioning) before a task B could start (like a backend application).

## Building

You can clone the source code locally and run:

```
$ cd build
$ GOOS=linux GOARCH=amd64 ./build.sh
```

You should replace `GOOS` and `GOARCH` for your desired values of platform/architecture. `build-all.sh` will build binaries for Linux and Darwin (macOS).

You can also download the binaries from [releases page](https://github.com/esdrasbeleza/wait-for-signal/releases). The [example](https://github.com/esdrasbeleza/wait-for-signal/tree/master/example) folder has instructions for how you can download the binaries using [wget](https://www.gnu.org/software/wget/wget.html) in a Dockerfile (like [this](https://github.com/esdrasbeleza/wait-for-signal/blob/9d9cb730de638759b49f30882d18f7c0989f6035/example/main-application/Dockerfile#L12)).

## Usage

In a host called `host1`, run `wait-for-signal`. In this scenario, we want to wait for a signal called `database-is-ready` which will be sent to port 12345:

```
host1:/myproject/bin$ ./wait-for-signal database-is-ready 12345
```

In another host, called `host2`, we'll provision our database and send a signal to `host1`:

```
host2:/myproject-db/bin$ ./database-provision.sh
Database is ready!
host2:/myproject-db/bin$ ./send-signal database-is-ready host1:12345
Sending "database-is-ready"... Done
```

An example application using `docker-compose` can be found in [example](https://github.com/esdrasbeleza/wait-for-signal/tree/master/example) folder.
