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

### See local images

```bash
docker images
```

### Start a container

```bash
docker run -d -p <host_port>:<container_port> <image>
```

- `-e` sets environment variables
- `-v` binds container directories to volumes. `<volume_name>:<container_path>`
    - This will create the volume if it doesn't exist

### List running containers

```bash
docker ps
```

You can add a `-a` to the end to see all containers (running or not)

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

_Note: to leave simply use the `exit` command_  

### Get stats of containers

```bash
docker stats
```

### Restart container

```bash
docker restart <CONTAINER_ID>
```

### Remove container

```bash
docker rm <CONTAINER_ID>
```

### Create a volume

```bash
docker volume create <volume_name>
```

### List volumes

```bash
docker volume ls
```

### Inspect a volume

```bash
docker volume inspect <volume_name>
```

### Remove a volume

```bash
docker volume rm <volume_name>
```
