# Bank Statement Viewer

A full-stack application that allows users to upload bank statement files, view transaction history, and analyze transaction data with sorting and filtering capabilities.

## Features

- 📤 Upload bank statement files (CSV format)
- 📊 View transaction history with sorting and pagination
- 💰 Track account balance
- 🚦 Transaction status tracking (Success, Pending, Failed)
- 🎨 Clean, responsive UI with pure CSS

## Tech Stack

### Frontend
- React 19 with TypeScript
- Vite for build tooling
- Pure CSS (no UI frameworks)
- React Hooks for state management

### Backend
- Go (Golang)
- Gin web framework
- GORM for database operations
- SQLite for data storage

## Getting Started

### Prerequisites

- Node.js 18+ (for frontend)
- Go 1.21+ (for backend)
- Git

### Backend Setup

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

### Frontend Setup

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
│   │   └── service/      # Business logic
│   └── go.mod           # Go module definition
│
├── frontend/             # Frontend source code
│   ├── public/           # Static assets
│   ├── src/              # React application
│   │   ├── components/   # Reusable UI components
│   │   ├── App.tsx       # Main application component
│   │   └── main.tsx      # Application entry point
│   └── package.json      # NPM dependencies
│
└── README.md             # This file
```

## Architecture Decisions

### Frontend
- **Component-Based Architecture**: The UI is built using React functional components with hooks for state management.
- **Pure CSS**: Styling is implemented with pure CSS for better control and smaller bundle size.
- **Type Safety**: TypeScript is used throughout the codebase for better developer experience and code quality.
- **Responsive Design**: The application is designed to work on both desktop and mobile devices.

### Backend
- **Layered Architecture**: Follows a clean architecture pattern with clear separation of concerns.
- **RESTful API**: Implements REST principles for API design.
- **SQLite**: Chosen for its simplicity and zero-configuration requirements for development.
- **Gin Framework**: Provides good performance and a clean API for routing and middleware.

## Environment Variables

### Backend
- `PORT`: Port to run the server on (default: 8080)
- `DB_PATH`: Path to SQLite database file (default: `./data/transactions.db`)

### Frontend
- `VITE_API_BASE`: Base URL for API requests (default: `http://localhost:8080`)

## Development

### Running Tests

#### Backend Tests
```bash
cd backend
go test ./...
```

### Linting

#### Frontend
```bash
cd frontend
npm run lint
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
