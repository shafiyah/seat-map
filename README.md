# Seat Map Assignment

The system processes flight seat map data from a provided JSON file, stores it in a relational database, and exposes multiple REST APIs following a clean architecture and MVC pattern. It also includes a basic frontend interface built with React to visualize seat maps and allow users to select seats interactively.

## Technologies Used

#### Backend (Go):
- Gin – HTTP web framework
- GORM – ORM for database operations
- PostgreSQL – Relational Database

#### Frontend (React):
- React (Vite + JSX)
- Axios – API communication
- Tailwind CSS – Styling

##  Backend API Overview

**1. `POST /flights/import-data`**
- Trigger data import from the provided JSON into the PostgreSQL database.
- JSON file must be placed in project directory as `SeatMapResponse.json`

**2. `GET /flights/:flightId/seats`**
- Return structured seat map per flight, including seat code, price, and availability.

**3. `GET /passenger/:userId`**
- Retrieve passenger details by ID.

**4. `POST /seat/selection`**
- Store selected seat(s) for a passenger.
- Payload format:
```
{
  "flightId": 1,
  "passengerId": 1,
  "seatId": 5
}
```
## Run Backend Project 

#### Prerequisites:
- Go 1.20+
- PostgreSQL (create database seatmap)

#### Steps:
```
# Clone project
git clone https://github.com/shafiyah/seat-map.git
cd seat-map/backend

# Run import to load JSON data
go run cmd/import.go

# Run server
go run cmd/server.go

# API accessible at http://localhost:8080
```

## Run Frontend Project 
#### Tech Stack:
- React + Vite
- TailwindCSS
- Axios

#### Run Frontend:
```
cd seat-map/frontend
npm install
npm run dev

# App runs at http://localhost:5173

```
## Result Dokumentation 
[Photo](https://drive.google.com/drive/folders/11XBoADCk3mS7B2CasPWtluXjoFmg8pR5?usp=sharing)


