# Superbank Simple Dashboard

## Description

This project consists of a multi-service application using Docker Compose. It includes a PostgreSQL database, Redis cache, authentication service, and account dashboard service, along with a frontend application.

## Project Decisions

- PostgreSQL is used as the database for storing user and account data.
- Redis is used for caching purposes, improving the performance of the authentication service and dashboard service.
- The authentication service handles login and token generation for secure access to the system.
- The account dashboard service provides the dashboard data for the user.
- The frontend application is built using Next Js and communicates with the backend services via HTTP API endpoints.

## Prerequisites

- Docker
- Docker Compose

Ensure you have Docker and Docker Compose installed on your machine before proceeding. You can follow the installation instructions from the official Docker documentation:

- [Install Docker](https://docs.docker.com/get-docker/)
- [Install Docker Compose](https://docs.docker.com/compose/install/)

## Setup

1. Clone this repository:

   ```bash
   git clone <repository_url>
   cd <project_directory>
   ```

2. Ensure that the `docker-compose.yml` file is present in your project’s root folder.

3. Add the required configuration files in the project:

   - **config.yaml**: The `config.yaml` file should be placed in the `./backend/authenticationService/configs/` and `./backend/accountDashboardService/configs/` directory. This file contains environment-specific settings for your application, such as database and Redis connection details. Also, the `./frontend/search-app/.env.local` file should be placed in the `./frontend/search-app/` directory. This file contains environment-specific settings for your application, such as authentication and dashboard API connection details.

   Example `config.yaml` for authentication service:

   ```yaml
   jwt_secret: secret

   app_name: auth_service
   environment: dev
   log_level: debug

   host: postgres_db_compose
   port: 5432
   username: root
   password: root
   dbname: postgres

   redis:
     address: redis://redis_db_compose:6379/1
     timeout: 3
     pool_size: 100

   token_cache:
     access_token_timeout: 15
     refresh_token_timeout: 1440
   ```

   Example `config.yaml` for account dashboard service:

   ```yaml
   host: postgres_db_compose
   port: 5432
   username: root
   password: root
   dbname: postgresapp_name: account_dashboard_service
   environment: dev
   log_level: debug

   host: postgres_db_compose
   port: 5432
   username: root
   password: root
   dbname: postgres

   auth:
     address: http://authentication_service:8091
     timeout: 10
   ```

   Example `.env.local` for frontend:

   ```env
   NEXT_PUBLIC_AUTH_API_URL=http://localhost:8091/v1/auth
   NEXT_PUBLIC_DASHBOARD_API_URL=http://localhost:8090/v1/dashboard
   ```

   **Note:** If the file is not present, the services may fail to start.

## Running the Application

Once you have added the `config.yaml` file, you can start the application using Docker Compose:

1. **Build and start the services**:

   Run the following command to build the containers and start the services:

   ```bash
   docker-compose up --build
   ```

   This will start the following services:

   - **PostgreSQL** (`postgres_db_compose`)
   - **Redis** (`redis_db_compose`)
   - **Account Dashboard Service** (`account_dashboard_service`)
   - **Authentication Service** (`authentication_service`)
   - **Frontend** (`frontend_app`)

2. **Run SQL Queries for Data Population**:

   Before testing the app, ensure you run any necessary SQL queries inside the `data_population` folder to set up your database. These queries will populate your database with the required schema and data.

   Example (adjust the path if needed):

   ```bash
   psql -h localhost -p 54321 -U root -d postgres -f ./data_population/your_sql_file.sql
   ```

3. **Access the Services**:

   - **Account Dashboard Service** will be available at: `http://localhost:8090`
   - **Authentication Service** will be available at: `http://localhost:8091`
   - **Frontend** will be available at: `http://localhost:3000`

4. **Check Logs**:

   To view logs for any service, use the following command:

   ```bash
   docker-compose logs <service_name>
   ```

   Example: To view logs for the authentication service:

   ```bash
   docker-compose logs authentication_service
   ```

5. **Stopping the Services**:

   To stop all running services, use:

   ```bash
   docker-compose down
   ```

   This will stop and remove all containers, but the data in the `./postgres_db_volume` directory will persist.

## Troubleshooting

- **Permission Denied Error**: If you encounter permission issues with the volume, run the following commands to set the correct permissions:

  ```bash
  sudo chmod -R 777 ./postgres_db_volume
  sudo chown -R $USER:$USER ./postgres_db_volume
  ```

## Conclusion

This setup uses Docker Compose to manage multiple services in a containerized environment. You can easily scale the application or modify the configuration as needed. If you need to make changes to the database or Redis configurations, update the `config.yaml` file and restart the containers.
