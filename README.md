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
* ```PROXY_API_KEY``` - key obtained from _ProxyAPI_ website
* ```PROXY_API_URL``` - URL to the needed model _(example: [Anthropic](https://api.proxyapi.ru/anthropic/v1))_
* ```DB_HOST``` - DB hostname, set to ```postgres``` for simplicity
* ```DB_PORT``` - port on which you would like to start PostgreSQL DB
* ```DB_USER``` - PostgreSQL user, ```postgres``` by default
* ```DB_PASSWORD``` - PostgreSQL DB password
* ```DB_NAME``` - name of the PostgreSQL DB
* ```DATABASE_URL``` - _optional variable_; URL to the PostgreSQL DB - set if you are going to run migrations

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

## 📡 API Documentation

Base URL: http://localhost:8080 (/app)

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

📘 Swagger docs available at http://localhost:8080/docs/docs

## 🧱 Architecture
### Static view diagram
You can access the diagram at [docs/architecture/MHC_static_view](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/blob/329de58270d14e2d5dcf010f64890b04010dd629/docs/architecture/MHC_static_view.png)

#### Diagram Explanation:
* Frontend (Flutter):
  * Includes key UI screens like LoginScreen, SessionScreen, SessionHistory, and SummaryScreen.
  * Subcomponents like ChatInput and TimerWidget are embedded in relevant views.
  * These screens interact with the backend to perform authentication, initiate chat sessions, and retrieve session summaries.
* Backend (Go + Fiber):
  * Composed of controllers: AuthController, SessionController, and MessageController, each responsible for handling corresponding routes.
  * The JWTMiddleware ensures that only authenticated users access protected endpoints.
  * MessageController delegates message handling to the LLMClient, which communicates with an external AI service (e.g., Claude API) to simulate therapeutic responses.
  * SummaryService handles the generation and saving of session summaries.
* Database (PostgreSQL):
  * Stores all persistent data: users, sessions, messages, and summaries.
  * Each controller interacts with the appropriate table to read/write data.
* External APIs:
  * The LLMClient can connect to external large language model APIs such as Claude or OpenAI, to generate responses and session summaries based on recent user messages.
 
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

## 🔮 Future Plans
- [ ] Add a local LLM support via Ollama API
- [ ] Introduce notifications about an everyday session
- [ ] Enhance the communication wrapping a user message

## Implementation checklist

### Technical requirements (20 points)
#### Backend development (8 points)
- [X] Go-based backend (3 points)
  Check the Go backend component at [backend/mental-health-api](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/tree/d296742e5d7aa6e4e97bb8372a88ae15188cedd4/backend/mental-health-api)
- [X] RESTful API with Swagger documentation (1 point)
- [X] PostgreSQL database with proper schema design (1 point)
  Check the database files at [backend/mental-health-api/database](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/tree/d296742e5d7aa6e4e97bb8372a88ae15188cedd4/backend/mental-health-api/database)
- [X] JWT-based authentication and authorization (1 point)
  Check the JWT-middleware at [backend/mental-health-api/middleware](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/tree/d296742e5d7aa6e4e97bb8372a88ae15188cedd4/backend/mental-health-api/middleware)
- [X] Comprehensive unit and integration tests (1 point)
  You can find unit tests for the backend functionality in the corresponding folders, e.g.:
  [backend/mental-health-api/handlers](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/tree/d296742e5d7aa6e4e97bb8372a88ae15188cedd4/backend/mental-health-api/handlers)
  [backend/mental-health-api/internal/llm](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/tree/d296742e5d7aa6e4e97bb8372a88ae15188cedd4/backend/mental-health-api/internal/llm)

#### Frontend development (8 points)
- [X] Flutter-based cross-platform application (mobile + web) (3 points)
  Check the Flutter cross-platform application at [\application](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/tree/d296742e5d7aa6e4e97bb8372a88ae15188cedd4/application)
- [X] Responsive UI design with custom widgets (1 point)
  You can find the most important widgets of our app at [application/lib/widgets](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/tree/d296742e5d7aa6e4e97bb8372a88ae15188cedd4/application/lib/widgets)
- [x] State management implementation (1 point)
- [ ] Offline data persistence (1 point)
- [x] Unit and widget tests (1 point)
- [X] Support light and dark mode (1 point)
  You can switch themes on the login screen of our app (sun button). 

#### DevOps & deployment (4 points)
- [x] Docker compose for all services (1 point)
- [x] CI/CD pipeline implementation (1 point)
- [x] Environment configuration management using config files (1 point)
- [ ] GitHub pages for the project (1 point)

### Non-Technical Requirements (10 points)
#### Project management (4 points)
- [x] GitHub organization with well-maintained repository (1 point)
- [x] Regular commits and meaningful pull requests from all team members (1 point)
- [ ] Project board (GitHub Projects) with task tracking (1 point)
- [x] Team member roles and responsibilities documentation (1 point)

#### Documentation (4 points)
- [x] Project overview and setup instructions (1 point)
- [ ] Screenshots and GIFs of key features (1 point)
- [x] API documentation (1 point)
- [x] Architecture diagrams and explanations (1 point)

#### Code quality (2 points)
- [x] Consistent code style and formatting during CI/CD pipeline (1 point)
- [ ] Code review participation and resolution (1 point)

### Bonus Features (up to 10 points)
- [x] Localization for Russian (RU) and English (ENG) languages (2 points)
- [x] Good UI/UX design (up to 3 points)
- [x] Integration with external APIs (fitness trackers, health devices) (up to 5 points)
- [ ] Comprehensive error handling and user feedback (up to 2 points)
- [ ] Advanced animations and transitions (up to 3 points)
- [ ] Widget implementation for native mobile elements (up to 2 points)

Total points implemented: XX/30 (excluding bonus points)

Note: For each implemented feature, provide a brief description or link to the relevant implementation below the checklist.
