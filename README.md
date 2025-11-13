# Bank Statement Viewer

A full-stack web application that allows users to upload and analyze bank statements, view transaction summaries, and inspect transaction issues - built with Go and React (TypeScript).

## Features

- 📤 Upload CSV-based bank statements
- 📊 View transaction history with pagination and sorting
- 💰 Compute account balance (credits − debits for successful transactions)
- 🚦 Detect transaction issues (Pending & Failed)
- 🎨 Responsive, clean UI with pure CSS
- 🧪 Unit testing for backend services
- 🐳 Containerized with Docker Compose
- ⚙️ CI/CD pipeline with GitHub Actions

## Tech Stack

### Frontend
- **Framework**: React 19 with TypeScript
- **Build Tool**: Vite
- **Styling**: Pure CSS (no UI frameworks)

### Backend
- **Language**: Go 1.21+
- **Web Framework**: Standard Library HTTP
- **Testing**: Standard Library testing package

## Getting Started

### Prerequisites

- **Development**
  - Node.js 18+ (for frontend development)
  - Go 1.21+ (for backend development)
  - Git
  - Docker & Docker Compose (for containerized development and deployment)

### Quick Start with Docker Compose

The easiest way to get started is using Docker Compose:

```bash
# Clone the repository
git clone https://github.com/bagindaisfa/flip-bank-statement-viewer.git
cd flip-bank-statement-viewer

# Start the application
docker-compose up --build
```

The application will be available at:
- Frontend: http://localhost:5173
- Backend API: http://localhost:8080

### Manual Setup

#### Backend Setup

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Start the backend server:
   ```bash
   go run cmd/server/main.go
   ```
   The backend will start on `http://localhost:8080` by default.

#### Frontend Setup

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the development server:
   ```bash
   npm run dev
   ```
   The frontend will be available at `http://localhost:5173`

## Project Structure

```
.
├── backend/               # Backend source code
│   ├── cmd/              # Application entry points
│   ├── internal/         # Private application code
│   │   ├── handler/      # HTTP request handlers
│   │   ├── models/       # Database models
│   │   ├── repository/   # Data access layer
│   │   ├── service/      # Business logic
│   │   └── utils/        # Utility functions
│   ├── Dockerfile        # Production Dockerfile
│   └── go.mod           # Go module definition
│
├── frontend/             # Frontend source code
│   ├── public/           # Static assets
│   ├── src/              # React application
│   │   ├── components/   # Reusable UI components
│   │   ├── App.tsx       # Main application component
│   │   └── main.tsx      # Application entry point
│   ├── Dockerfile        # Production Dockerfile
│   └── package.json      # NPM dependencies
│
├── .github/workflows/    # GitHub Actions workflows
│   └── ci.yml           # CI/CD pipeline
│
├── docker-compose.yml    # Development and production compose file
└── README.md            # This file
```

## Architecture Decisions

### Frontend Architecture
- **Component-Based**: Built with functional React components and hooks for state management
- **Type Safety**: TypeScript for better developer experience and code quality
- **Performance**: Code splitting and lazy loading with React.lazy and Suspense
- **Responsive Design**: Mobile-first approach with pure CSS for better performance
- **State Management**: Local component state for simplicity, with prop drilling for shared state

### Backend Architecture
- **Layered Architecture**: Clear separation of concerns with handler, service, and repository layers
- **RESTful API**: Follows REST principles with consistent response formats
- **Error Handling**: Centralized error handling with appropriate HTTP status codes
- **Input Validation**: Request validation at the handler level
- **In-Memory Storage**: Simple in-memory storage for development, easily swappable with a database

### CI/CD Pipeline
- **Automated Testing**: Runs unit tests for both frontend and backend
- **Linting**: Enforces code style and catches potential issues
- **Docker Builds**: Verifies that Docker images can be built successfully
- **Deployment**: Ready to be extended for automated deployments

### Containerization
- **Docker Compose**: Simplified local development and production deployment

## Environment Variables

### Frontend
- `VITE_API_BASE`: Base URL for API requests (default: `http://localhost:8080`)


### Running Tests

#### Backend Tests
```bash
cd backend/internal/service
go test ./...
```

## API Endpoints

| Method | Endpoint       | Description |
|--------|----------------|-------------|
| POST   | `/upload`      | Upload a CSV file containing transaction data |
| GET    | `/balance`     | Get the total account balance |
| GET    | `/issues`      | Get paginated list of transactions with issues |

## Deployment

### Production with Docker Compose

1. Update environment variables in `docker-compose.yml` if needed
2. Run:
   ```bash
   docker-compose -f docker-compose.yml up --build -d
   ```
