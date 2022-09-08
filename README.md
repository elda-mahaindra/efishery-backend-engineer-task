# eFishery Back End Engineer Task

A group of two services: auth-app and fetch-app.

## Environment Variables

To run this project, the following environment variables needed to be added.

### 1. Auth App

Create `app.env` files inside auth-app directory (`/auth-app`) and add the following environment variables:

```bash
    # PORT
    PORT=<PORT>

    # TOKEN GENERATION
    SECRET=<SECRET>
```

Notes:

- `PORT`is the port number where the net will listen to.
- `SECRET` is a jwt key to sign the generated token.

The environment variables will be something like the below:

```bash
    # PORT
    PORT=8080

    # TOKEN GENERATION
    SECRET=EFISHERYBACKENDENGINEERTASK 2022
```

### 2. Fetch App

Create `.env` files inside fetch-app directory (`/fetch-app`) and add the following environment variables:

```bash
    PORT=<PORT>
    SECRET=<SECRET>
```

Notes:

- `PORT`is the port number where the net will listen to.
- `SECRET` is a jwt key to sign the generated token.

The environment variables will be something like the below:

```bash
    PORT=4000
    SECRET=EFISHERYBACKENDENGINEERTASK 2022
```

## Docker Compose File

Change the configuration inside the `docker-compose.yaml` file according to your environment.

Make sure you set the bind mount configuration of `auth-app` service to fit your environment.

```bash
    volumes:
      - <TARGET>:/app/data
```

Notes:

- `TARGET`is the target directory to save the user data.

The environment variables will be something like the below:

```bash
    volumes:
      - D:\Codes\Web\Others\efishery-backend-engineer-task\efishery-backend-engineer-task\data:/app/data
```

## Start Servers

Make sure the environment variables have been set up before trying to start the servers. To start the servers, run the following command from the root directory where the `docker-compose.yaml` file is placed.

For the plugin version (space compose):

```bash
  docker compose up
```

For the standalone version (dash compose):

```bash
  docker-compose up
```

## Stop Servers

### 1. Docker Compose Down

To **remove all** existing containers and networks related to this project, Run the following command from the root directory where the `docker-compose.yaml` file is placed.

For the plugin version (space compose):

```bash
  docker compose down
```

For the standalone version (dash compose):

```bash
  docker-compose down
```

### 2. Remove Docker Images

The following command structure is used to remove a docker image:

```bash
  docker rmi <IMAGE_NAME>
```

The commands to remove all images related to this project will be something like the below:

```bash
  docker rmi efishery-backend-engineer-task_auth-app
```

```bash
  docker rmi efishery-backend-engineer-task_fetch-app
```
