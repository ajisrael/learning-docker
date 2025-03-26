# learning-docker

## Terms

### Image

An image is a read only definition of a container

### Container

A virtualized read-write environment

Containers are different from VMs in that they are virtualized on top of the hosts operating system,
instead of running on emulated hardware from a hypervisor.

### Volume

Volumes allow us to have "persistent state" for a container. Otherwise containers are considered "stateless."

## Commands

### List running containers

```bash
docker ps
```

### Execute a command on a running container

```bash
docker exec <CONTAINER_ID> <command>
```

Example:

This will execute the `ls` command on the `1234ABC` container

```bash
docker exec 1234ABC ls
```

You can also use `exec` to start an interactive shell

```bash
docker exec -it <CONTAINER_ID> /bin/sh
```

- `-i` makes the exec command interactive
- `-t` gives us a tty (keyboard) interface
- `/bin/sh` is the path to the command we're running. After all, a command line shell is just a program. sh is a more basic version of bash.

__ Note: to leave simply use the `exit` command __  

### Get stats of containers

```bash
docker stats
```
