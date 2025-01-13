# Docker Instructions

## Building

Build image:
`docker build --tag vette-tracker-services .`

To see a list of images:
`docker images`

You should see an image named `vette-tracker-services`.

To remove an image:
`docker rmi vette-tracker-services`

To build multistage image:
`docker build -t vette-tracker-services:multistage -f Dockerfile.multistage .`

## Running

To run image:
`docker run -p 8080:8080 vette-tracker-services`

To run in detached mode:
`docker run -d -p 8080:8080 --name rest-server vette-tracker-services`

To run the multistage image in detached mode:
`docker run -d -p 8080:8080 --name rest-server vette-tracker-services:multistage`s

## Containers

To see the containers that are running on your machine:
`docker ps`

To see all containers:
`docker ps -a`

To stop a container:
`docker stop <container_id>`

To remove a container:
`docker rm <container_id>`
