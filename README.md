# serve-init

Initialize a scratchpad server.

## Prerequisites

- Install ansible.

  On macOS:

  ```sh
  brew install ansible
  ```

  On Ubuntu:

  ```sh
  sudo apt-add-repository ppa:ansible/ansible \
      && sudo apt update \
      && sudo apt install ansible
  ```

## Initialize the server

- Run setup:

  ```sh
  make init
  ```

- Upgrade OS to the latest Ubuntu LTS:

  ```sh
  make upgrade
  ```

- Add the current user's ssh to the server:

  ```sh
  make add_ssh
  ```

- Install docker:

  ```sh
  make add_docker
  ```

- Setup a simple 'hello world' service with Caddy:

  ```sh
  make add_caddy
  ```

## Deployment

The template `.github/workflows/deploy.yml` uses GitHub Actions to deploy a simple REST API
defined in the `app` directory. 
