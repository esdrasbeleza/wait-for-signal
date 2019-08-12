# wait-for-signal

These are two binaries to be used when an application needs to wait for some other task to finish before it starts, probably in different hosts.

When I created these tools, I intended to use them in different Docker containers executed by `docker-compose`. An application needed to wait for another container to be fully provisioned before it was able to start.

This situation is similar to the ones handled by [wait-for-it](https://github.com/vishnubob/wait-for-it), but instead of waiting to a service to be running and listening in some port, I want a task A to be finished (like a database provisioning) before a task B could start (like a backend application).

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
