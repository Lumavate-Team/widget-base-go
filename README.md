## Build the container containers

Clone the repo, and from the root directory of the repo, run the following command.

```
docker build --no-cache --rm -t gobasewidget:1.0 .
```

This command will use the Dockerfile to build the image.  We are specifying that we
do not want to use cache when building containers and that we want to remove intermediate containers after a successful builld.

[Docker Build Options](https://docs.docker.com/engine/reference/commandline/build/)


## Run the container

After the container has been built, you can now run the container.  Run the following command.

```
docker run -d -p 5000:8080 --volume "$(pwd)"/widget:/go/src/widget gobasewidget:1.0
```

This command will run the container in detached mode.  It will map port 5000 on your machine to port 8080 on the container.
It will map the widget directory to the /go/src/widget directory inside the container.  This will will allow you to modify files in your local widget directory, and it will reload the process when the files change.  You should be able to go to http://localhost:5000 

[Docker Run Options](https://docs.docker.com/engine/reference/commandline/run/)

## Check the logs

To stream the logs to your terminal, you first need to run

```
docker ps
```

which will result in something like

```
CONTAINER ID        IMAGE                  COMMAND                  CREATED             STATUS              PORTS                    NAMES
676f62d88565        gobasemac4:dev021418   "/bin/sh -c 'bee run'"   15 minutes ago      Up 16 minutes       0.0.0.0:5000->8080/tcp   dreamy_albattani
```

you can then use the Container ID to stream the logs

```
docker logs -f 676
```



