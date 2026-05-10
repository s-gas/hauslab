# continuous integration

CI is done through GitHub Actions.

## Pipeline

On every push, the pipeline:

1. Runs tests for Go services
2. Builds and pushes the Docker images to Docker Hub

## Secrets

The following action secrets need to be added the GitHub repository:

- `DOCKERHUB_USERNAME`
- `DOCKERHUB_PASSWORD`

To add an action secret, in the projects repository go to **Settings**->**Secrets and variables**->**Actions**.

## New service

When developing a new service, remember to add a step in `ci.yaml` for building and pushing its image.
