# Deployment

## Deployment Setup

This app runs on a Digital Ocean droplet using Docker and Docker Compose. The API is built using a multi-stage Dockerfile to create a production-ready image. The image is then pushed to the GitHub Container Registry, from where it can be pulled onto the server.

## Automated Deployment with GitHub Actions

The deployment process has been automated using GitHub Actions to build, push, and deploy the application whenever changes are pushed to the main branch. See the [.github/workflows/deploy.yml](../.github/workflows/deploy.yml) file for the details of the workflow.

### Setup Process

A dedicated user called `gh-deploy` was created on the Droplet to manage deployments. A dedicated SSH key was created for this user, and the public key was added to the `~/.ssh/authorized_keys` file on the server. The private key is stored as a GitHub secret called `SSL_PRIVATE_KEY`.

## Manual Deployment Process

1. Build API Image Locally
2. Push Image to GitHub Container Registry
3. Pull Image from GitHub Container Registry on DO Droplet
4. Run Container from Pulled Image using server docker compose file
5. Verify Container is Running

### Commands

To build the image:

```sh
# Build your API service image
# docker build -t ghcr.io/amchappell9/vette-tracker-services:latest -f Dockerfile.multistage .

# Build and push for multiple architectures in one command
docker buildx build --platform linux/amd64,linux/arm64 \
  -t ghcr.io/amchappell9/vette-tracker-services:latest \
  -f Dockerfile.multistage \
  --push .
```

Login to GitHub Container Registry:

```sh
echo $CR_PAT | docker login ghcr.io -u amchappell9 --password-stdin
```

To push the image:

```sh
# Push the built image to GitHub Container Registry
docker push ghcr.io/amchappell9/vette-tracker-services:latest
```

On server:

```sh
# Pull the image from GitHub Container Registry
cd /opt/vette-tracker
docker compose pull
docker compose up -d
```

## Database Backups

A GitHub action creates a backup of the database on the server, and pushes it to Digital Ocean SPaces. The backup is created using the `pg_dump` command, which creates a SQL file that can be used to restore the database.

## Initial Digital Ocean Setup steps

1. Create a droplet
2. Generate SSH keys (1 time thing)
3. SSH into the droplet
4. Install Docker - https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-20-04
5. Install Docker Compose - https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-compose-on-ubuntu-22-04
6. Setup docker-compose.yml file and .env file
7. Run docker compose up -d
