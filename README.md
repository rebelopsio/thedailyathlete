```markdown
# The Daily Athlete

This repository contains a full-stack application with a Nuxt.js frontend and a Go backend using PostgreSQL.

## Directory Structure
```

/my-app
/backend
/handlers
/models
/routes
main.go
go.mod
go.sum
/frontend
/components
/layouts
/pages
/static
/store
nuxt.config.js
package.json
yarn.lock
.gitignore
README.md

````

## Setup

### Prerequisites

- Go 1.16+
- Node.js 14+
- PostgreSQL 12+
- Yarn (or npm)
- Docker (optional, for running PostgreSQL)

### Backend (Go)

#### 1. Install Go Dependencies

Navigate to the `backend` directory and install the Go dependencies:

```bash
cd backend
go mod tidy
````

#### 2. Configure PostgreSQL

Make sure you have PostgreSQL installed and running. You can use Docker for this:

```bash
docker run --name my-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres
```

Create a `.env` file in the `backend` directory with your PostgreSQL connection details:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=mysecretpassword
DB_NAME=myapp
```

#### 3. Run the Backend

```bash
go run main.go
```

The backend server will be running on `http://localhost:8081`.

### Frontend (Nuxt.js)

#### 1. Install Node.js Dependencies

Navigate to the `frontend` directory and install the dependencies:

```bash
cd frontend
yarn install
# or
npm install
```

#### 2. Run the Frontend

```bash
yarn dev
# or
npm run dev
```

The frontend server will be running on `http://localhost:3000`.

## API Endpoints

### Login

**URL:** `/api/login`  
**Method:** `POST`  
**Body:**

```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

### Logout

**URL:** `/api/logout`  
**Method:** `POST`

## File Descriptions

### Backend

- **main.go**: Entry point for the Go application. Sets up the router and starts the server.
- **handlers/**: Contains the HTTP handlers for different routes.
- **models/**: Contains the data models and database interactions.
- **routes/**: Defines the routes and their corresponding handlers.

### Frontend

- **components/**: Contains reusable Vue components.
- **layouts/**: Contains layout components.
- **pages/**: Contains the pages of the Nuxt.js application.
- **static/**: Contains static files.
- **store/**: Contains Vuex store modules.
- **nuxt.config.js**: Nuxt.js configuration file.
- **package.json**: Node.js dependencies and scripts.

## Running the Application

1. Start the PostgreSQL database (if using Docker):

   ```bash
   docker start my-postgres
   ```

2. Start the backend server:

   ```bash
   cd backend
   go run main.go
   ```

3. Start the frontend server:

   ```bash
   cd frontend
   yarn dev
   # or
   npm run dev
   ```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
