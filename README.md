# learning-docker

## Terms

### Image

An image is a read only definition of a container.

### Container

A virtualized read-write environment.

Containers are different from VMs in that they are virtualized on top of the hosts operating system,
instead of running on emulated hardware from a hypervisor.

### Volume

Volumes allow us to have "persistent state" for a container. Otherwise containers are considered "stateless."

### Bridge Network

An internal network that allows for communication betwen containers without accessing the open internet.

### Dockerfile

A file that lists instructions for how to build an image.

## Commands

### See local images

```bash
docker images
```

### Start a container

```bash
docker run -d -p <host_port>:<container_port> <image>
```

- `-d` runs the container in "detatched" mode, so it doesn't take over your current terminal
- `-p` applies port forwarding from the `host_port` into the `container_port`, thereby exposing the container for external access
- `-e` sets environment variables
- `-v` binds container directories to volumes. `<volume_name>:<container_path>`
    - This will create the volume if it doesn't exist
    - You can use a path on your machine as the volume name as well. So `$PWD/some_file.txt:/bin/config/config.txt` would copy `some_file.txt` in the current directory to `config.txt` in the container
- `--network` will attach the container to a network
    - When specifying a network you would only use `-p` for anything that you would want to be an entrypoint into the network. For example, an application load balancer.
- `--name` will allow you to assign a name to the container
    - Once you've specified a name for the container you can run it again by using just the name: `docker run <container_name>`
    - This will also map to the host name of the container for the bridge network that it is in.

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

### Create a bridge network

```bash
docker network create <network_name>
```

### List networks

```bash
docker network ls
```

### Build an image from a Dockerfile
 
```bash
docker build . -t <image_name>:(version)
```

This will use a `Dockerfile` located in the current directory to create an image with `image_name`. Usually you would specify the `version` as `latest` if you do specify it at all as this is optional.

For example:

```bash
docker build . -t helloworld:latest
```

This creates an image called `helloworld` with the `latest` tag.

You can then run the new image as you would any other:

```bash
docker run helloworld
```

### Push an image to docker hub

Once you've built an image locally, you can push after loging into the corresponding registry.

Login with the following command and complete the prompts that follow:

```bash
docker login
```

You can also pass the username and password into the command with `--username` and `--password`

After logging in, you can push a local image with the following command:

```bash
docker push <image_name>
```

You will want to make sure that you prefix the image with your username followed by a `/`, then the image name when naming the image.

