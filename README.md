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

#### /auth-app

```bash
    # PORT
    PORT=8080

    # TOKEN GENERATION
    SECRET=EFISHERYBACKENDENGINEERTASK 2022
```
