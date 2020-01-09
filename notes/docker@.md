## Containers
Use docker container my_command
~~~sh
docker create # - Create a container from an image.
docker start # - Start an existing container.
docker run # - Create a new container and start it.
docker ls # - List running containers.
docker inspect # - See lots of info about a container.
docker logs # - Print logs.
docker stop # - Gracefully stop running container.
docker kill # - Stop main process in container abruptly.
docker rm # - Delete a stopped container.
~~~
## Images
Use docker image my_command
~~~sh
docker build # - Build an image.
docker push # - Push an image to a remote registry.
docker ls # - List images.
docker history # - See intermediate image info.
docker inspect # - See lots of info about an image, including the layers.
docker rm # - Delete an image.
~~~
## Misc
~~~sh
docker version # - List info about your Docker Client and Server versions.
docker login # - Log in to a Docker registry.
docker system prune # - Delete all unused containers, unused networks, and dangling images.
~~~
## Run
Run NGinX container with config - startup after reboot machine / expose port 8080 from container port 80 
~~~sh
docker run -d -p 8080:80 --restart=unless-stopped --name nginx_alpine nginx:alpine
~~~
### Several commands
~~~sh
# - Create a container from an image.
docker container create my_repo/my_image:my_tag

# I’ll shorten my_repo/my_image:my_tag to my_image for the rest of the article.
# There are a lot of possible flags you could pass to create.
docker container create -a STDIN my_image
# -a is short for --attach. Attach the container to STDIN, STDOUT or STDERR.
~~~
Now that we’ve created a container let’s start it.
~~~sh
# - Start an existing container.
docker container start my_container

# Note that the container can be referred to by either the container’s ID or the container’s name.
docker container start my_container
~~~