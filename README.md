# Casual Chess
[![Wails](https://img.shields.io/badge/wails-%23242526.svg?style=for-the-badge&logo=wails&logoColor=red&labelColor=%23242526)](https://wails.io/)
[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Svelte](https://img.shields.io/badge/svelte-%23f1413d.svg?style=for-the-badge&logo=svelte&logoColor=white)](https://svelte.dev)
[![TypeScript](https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white)](https://www.typescriptlang.org/)
[![SASS](https://img.shields.io/badge/SASS-hotpink.svg?style=for-the-badge&logo=SASS&logoColor=white)](https://sass-lang.com/)
[![SQLite](https://img.shields.io/badge/sqlite-%2307405e.svg?style=for-the-badge&logo=sqlite&logoColor=white)](https://www.sqlite.org/)

A desktop chess application that lets you play casual, untimed games with friends. Built with Golang, Wails, and Svelte.

## Features

- Play untimed chess games with friends
- Pick up and continue games at your convenience
- Google OAuth authentication
- Cross-platform desktop application
- Real-time game state synchronization

## Tech Stack

### Backend
- Golang with Gin framework
- SQLite database
- RESTful API architecture

### Frontend
- Wails (Golang desktop framework)
- Svelte
- TypeScript
- SASS
- bbolt (key-value store)

## Prerequisites

- Latest version of Go (1.x or higher)
- Wails - Follow the [official installation guide](https://wails.io/docs/gettingstarted/installation)

## Project Structure

```
casual-chess-golang/
├── backend/         # Golang backend server
└── frontend/        # Wails + Svelte frontend
```

## Building and Running

### Backend

```bash
cd backend
go build backend
./backend
```

### Frontend

For development:
```bash
cd frontend
wails dev
```

For production build:
```bash
cd frontend
wails build
```

## License

This project is licensed under the Apache License 2.0 - see the LICENSE file for details.

## Acknowledgments

- [Wails](https://wails.io/) - The framework that made this desktop application possible
- [Svelte](https://svelte.dev/) - For the reactive frontend
- [Gin](https://gin-gonic.com/) - The web framework used for the backend
