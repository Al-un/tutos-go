## Simple containerizing

### Docker config

See `lang_tutorialedge/Dockerfile` for Docker configuration

### Docker image

To build the Docker image:

```sh
# locate terminal to tutos-go/lang_tutorialedge
#   al-un   : namespace
#   go-...  : image name
#   0.2     : image tag
sudo docker build -t al-un/go-simple-web-server-tls:0.2
```

Check the image with

```sh
sudo docker images
```

which should output something like

| REPOSITORY                     | TAG              | IMAGE_ID     | CREATED       | SIZE  |
| ------------------------------ | ---------------- | ------------ | ------------- | ----- |
| al-un/go-simple-web-server-tls | 0.2              | 13bebd261404 | 9 minutes ago | 359MB |
| golang                         | 1.12.0-alpine3.9 | 2205a315f9c7 | 7 months ago  | 347MB |

### Docker container

To create a container from the previously created image:

```sh
sudo docker run -p 8000:8080 -it al-un/go-simple-web-server-tls:0.2
```

- Image tag is required as `latest` is not yet defined
- `https://localhost:8000` would be mapped to Docker's 8080 port
- `-i` stands for _interactive_
- `-t` for _timeout_?

Stop the container with `Ctrl+C`

Check the containers:

```
sudo docker ps -a
```

| CONTAINER_ID | IMAGE                              | COMMAND                | CREATED       | STATUS                   | PORTS | NAMES             |
| ------------ | ---------------------------------- | ---------------------- | ------------- | ------------------------ | ----- | ----------------- |
| 0797b1880ad7 | al-un/go-simple-web-server-tls:0.2 | "/go/src/github.com/…" | 9 minutes ago | Exited (2) 8 minutes ago |       | elegant_antonelli |

To start a container:

```sh
# with container name
sudo docker container start elegant_antonelli
# with container ID
sudo docker container start 0797
```
