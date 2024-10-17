# Project Setup

Ensure you have the following installed on your system

- `Docker`
- `Go (version 1.18+)`

# Step 1: Configure Environment Variables

Before running the application, you need to set up the environment variables in the `.env.example` file to `.env`

Add the following values to the file:

```bash
CLIENT_ID=
CLIENT_SECRET=
```

# Step 2: Start Docker Services

Once the environment is configured, you can start the required services (such as `PostgreSQL`, `MinIO`) using Docker.

to start a docker container with the following command it contians the following

```bash
## This command will run all the services defined in the docker-compose.yml file in detached mode.
docker compose up -d
```

# Step 3: Import base data in database

First You need to import the following data from folder CSV_data folder in the repository directory and then import the data to database:
`user_groups_202410180220.csv` this file will contain the following data: [user groups] import the following data from the repository directory to table user_groups

Second You need to import the following data from folder CSV_data folder in the repository directory and then import the data to database:
`users_202410180227.csv` this file will contain the following data: [example users] import the following data from the repository directory to table users

# Step 4: Run the Go Application

With the services running, you can now start the Go application. Run the following command from your project directory:

```bash
go run cmd/main.go
```

# Step 5: If you want to read the documentation for the API

You can read the documentation at `http://localhost:3000/api/docs`
