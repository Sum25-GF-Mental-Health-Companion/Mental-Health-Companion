# Mental Health Companion

## Overview
A cross-platform virtual AI psychologist app tailored for students, providing timed, private therapy-like sessions with AI, summaries, and session history. Built with Flutter (Web & Mobile), Go (REST API), PostgreSQL, and LLM integration.

## 🌟 Features

  🎯 Login & registration

  💬 Chat with AI psychologist

  ⏲️ Fixed-time sessions (non-interruptible)

  🧾 Summarized session history

  🌙 Dark/light themes + RU/EN language switch

  📶 Offline support with SQLite/Hive

## 🛠️ Setup Instructions
### Prerequisites

* Flutter SDK (>=3.2.0)
* Go 1.21+
* PostgreSQL 15+
* Docker + Docker Compose (for backend + db)

### 🔧 Backend Setup

```
cd backend/mental-health-api
cp .env.example .env  # configure DB and API keys
go run main.go
```

OR run via Docker:
```
docker-compose up --build
```

Apply migrations (using Goose):
```
goose -dir db/migrations postgres "postgres://user:pass@localhost:5432/mentalhealth?sslmode=disable" up
```

### 📱 Frontend Setup

```
cd frontend
flutter pub get
flutter run -d chrome  # or your preferred device
```

## 🖼️ Features in screenshots
#### Login Screen

#### Chat with AI

#### Timer Widget

#### Session History

    All assets are located in assets/screenshots/

## 📡 API Documentation

    Base URL: http://localhost:8000

* Auth
  * POST	/auth/register - Register user
  * POST	/auth/login	- Login, returns JWT
* Session
  * GET	/session/start - Begin session, returns session ID
  * POST	/session/end - Ends session, returns summary
* Messaging
  * POST	/message - Send user message, get AI response
* History
  * GET	/sessions - Get past session summaries

📘 Swagger docs available at http://localhost:8000/docs

## 🧱 Architecture

[Flutter App (Web + Mobile)]
       ↓↑ REST API (JWT)
[Go Backend Server — Fiber]
       ↓
[PostgreSQL Database]
       ↓
[External LLM API (OpenAI/Claude)]

## ⚙️ Tech Stack

* Flutter + Riverpod
* Go + Fiber
* PostgreSQL + Goose
* OpenAI API / Claude API
* Docker + GitHub Actions (CI/CD)

## 🗂️ Folder Structure
Paste after finishing project

## 🧠 LLM Context Strategy

* Fetches summaries of last 50sessions 
* Sends to LLM with user message for reply + summary generation

## 🚀 Deployment
* Flutter Web: Vercel / GitHub Pages
* Backend: Railway / Render / Fly.io
* CI/CD: GitHub Actions (Lint, Test, Docker Build/Push)

## 🙌 Team
* Damir - LLM integration
* Vladimir - Flutter app
* Syoma - Flutter app + project deployment
* Magomed - server API 
* Pavel - Database + documentation
