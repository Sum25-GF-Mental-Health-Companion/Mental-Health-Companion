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

#### Initial Backend Configuration

**Configure .env file**
* PROXY_API_KEY - key obtained from _ProxyAPI_ website
* PROXY_API_URL - URL to the needed model _(example: [Anthropic](https://api.proxyapi.ru/anthropic/v1))_
* DB_HOST - DB hostname, set to ```postgres``` for simplicity
* DB_PORT - port on which you would like to start PostgreSQL DB
* DB_USER - PostgreSQL user, ```postgres``` by default
* DB_PASSWORD - PostgreSQL DB password
* DB_NAME - name of the PostgreSQL DB
* DATABASE_URL - _optional variable_; URL to the PostgreSQL DB - set if you are going to run migrations

#### Server application startup

**Windows**  
```
cd backend/mental-health-api
docker-compose up --build -d
cd ../../application
flutter run
```  

**Unix**  
```
cd backend/mental-health-api
sudo docker-compose up --build -d
cd ../../application
flutter run
```  

**Apply migrations using Goose**  
```
goose -dir db/migrations postgres "$DATABASE_URL" up
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
* CI/CD: GitHub Actions (Lint, Test)

## 🙌 Team
* Damir - LLM integration
* Vladimir - Flutter application
* Semyon - Flutter application, CI/CD
* Magomedgadzhi - server API
* Pavel - Database, documentation
