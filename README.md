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

## 🖼️ Features in Screenshots
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

### Static View Diagram
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
 
### Dynamic View Diagram
You can access the diagram at [docs/architecture/MHC_dynamic_view](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/blob/4a9664571118bb356634a3e6c5191a8430acb6fc/docs/architecture/MHC_dynamic_view.png)

#### Diagram Explanation
This diagram illustrates the real-time communication workflow from the moment a user sends a message until the AI responds.

1) User Interaction:
   * The user types a message and presses "Send" in the ChatInput widget of the Flutter frontend.
   
2) Frontend → Backend Request:
   * A POST /message request is sent to the MessageController (Go backend) with the session_id and the user's message.
    
3) Backend Processing – Save & Context Building:
   * The message is saved to the messages table in PostgreSQL with sender='user'.
   * All previous messages from the current session are fetched to maintain contextual continuity.
   * Additionally, the backend retrieves the last 20 summary records (from previous sessions) to give the AI broader conversation history.

4) LLM Interaction:
   * The LLMClient composes a prompt that includes:
     * Full session context
     * Compressed summaries
     * The latest user message
   * It sends this prompt to the Calude.

5) AI Response Handling:
   * Claude responds with a generated reply.
   * The backend logs this response in the messages table with sender='ai'.

6) Response Delivery:
   * The AI's reply is sent back to the frontend.
   * The Flutter UI displays the AI response in the chat.
 
## ⚙️ Tech Stack

* Flutter + Riverpod
* Go + Fiber
* PostgreSQL + Goose
* OpenAI API / Claude API
* Docker + GitHub Actions (CI/CD)

## 🧠 LLM Context Strategy

* Fetches summaries of last 50 sessions 
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
  * Check the Go backend component at [backend/mental-health-api](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/tree/d296742e5d7aa6e4e97bb8372a88ae15188cedd4/backend/mental-health-api)
- [X] RESTful API with Swagger documentation (1 point)
  * You can see API ducomentataion and link to Swagger documentation above in the [📡 API Documentation](#api-documentation) section.
- [X] PostgreSQL database with proper schema design (1 point)
  * Check the database files at [backend/mental-health-api/database](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/tree/d296742e5d7aa6e4e97bb8372a88ae15188cedd4/backend/mental-health-api/database)
- [X] JWT-based authentication and authorization (1 point)
  * Check the JWT-middleware at [backend/mental-health-api/middleware](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/tree/d296742e5d7aa6e4e97bb8372a88ae15188cedd4/backend/mental-health-api/middleware)
- [X] Comprehensive unit and integration tests (1 point)
  * You can find unit tests for the backend functionality in the corresponding folders, e.g.:
    * [backend/mental-health-api/handlers](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/tree/d296742e5d7aa6e4e97bb8372a88ae15188cedd4/backend/mental-health-api/handlers)
    * [backend/mental-health-api/internal/llm](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/tree/d296742e5d7aa6e4e97bb8372a88ae15188cedd4/backend/mental-health-api/internal/llm)

#### Frontend development (8 points)
- [X] Flutter-based cross-platform application (mobile + web) (3 points)
  * Check the Flutter cross-platform application at [\application](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/tree/d296742e5d7aa6e4e97bb8372a88ae15188cedd4/application)
- [X] Responsive UI design with custom widgets (1 point)
  * You can find the most important widgets of our app at [application/lib/widgets](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/tree/d296742e5d7aa6e4e97bb8372a88ae15188cedd4/application/lib/widgets)
- [x] State management implementation (1 point)
  * All screens in our app are stateful widgets and can be accessed at [application/lib/screens/](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/tree/3cff30ce7828a5c26b770325e3b60a5d907e723e/application/lib/screens)
- [ ] Offline data persistence (1 point)
- [x] Unit and widget tests (1 point)
  * You can find all the flutter tests at [application/tests](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/tree/3cff30ce7828a5c26b770325e3b60a5d907e723e/application/test)
- [X] Support light and dark mode (1 point)
  * You can switch themes on the login screen of our app (sun button).

#### DevOps & deployment (4 points)
- [x] Docker compose for all services (1 point)
  * We have Docker compose for all backend services. See our docker files at:
    * [backend/mental-health-api/Dockerfile](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/blob/3d10ae5913dcc7de7d238fa61467f5997138a4ec/backend/mental-health-api/Dockerfile)
    * [backend/mental-health-api/docker-compose.yml](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/blob/3d10ae5913dcc7de7d238fa61467f5997138a4ec/backend/mental-health-api/docker-compose.yml)
- [x] CI/CD pipeline implementation (1 point)
  * We added CI/CD pipeline, you can check its presence in actions and you can find workflows configs at [.github/workflows](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/tree/51f3c92a45d297d9a9b29698b8670ec8ddf5153d/.github/workflows)
- [x] Environment configuration management using config files (1 point)
  * You cannot find .env file in this repository, because we added it to secrets.
- [ ] GitHub pages for the project (1 point)

### Non-Technical Requirements (10 points)
#### Project management (4 points)
- [x] GitHub organization with well-maintained repository (1 point)
  * You are here and you can check this.
- [x] Regular commits and meaningful pull requests from all team members (1 point)
  * All of our team members contributed to the project and made commits. You can check list of contributors (one of our guys has two accounts).
- [x] Project board (GitHub Projects) with task tracking (1 point)
  * You can check our project board at [Mental-Health-Companion Project board](https://github.com/orgs/Sum25-GF-Mental-Health-Companion/projects/1/views/1?system_template=team_planning)
- [x] Team member roles and responsibilities documentation (1 point)
  * Check section "Team" above

#### Documentation (4 points)
- [x] Project overview and setup instructions (1 point)
  * You can find overview and setup instructions in the sections "Overview" and "Setup instructions" above.
- [x] Screenshots and GIFs of key features (1 point)
  * You can find link to screenshots of most relevant screens and features in the section "Features in Screenshots" above
  * Instructions can be found in the [🛠️ Setup Instructions](#setup-instructions) section
- [ ] Screenshots and GIFs of key features (1 point)
- [x] API documentation (1 point)
  * Detailed API documentation can be found in the [📡 API Documentation](#api-documentation) section
- [x] Architecture diagrams and explanations (1 point)
  * You can find links to architecture diagrams (static and dynamic) and their explanations in the section "Architecture" above.

#### Code quality (2 points)
- [x] Consistent code style and formatting during CI/CD pipeline (1 point)
  * Code style and practices are following the best practices learned in the course
- [ ] Code review participation and resolution (1 point)

### Bonus Features (up to 10 points)
- [x] Localization for Russian (RU) and English (ENG) languages (2 points)
  * Language packages can be found in [application/lib/l10n](https://github.com/Sum25-GF-Mental-Health-Companion/Mental-Health-Companion/tree/main/application/lib/l10n) directory
- [x] Good UI/UX design (up to 3 points)
  * Stunning, minimalistic, and, most important, simple user interface
- [x] Integration with external APIs (fitness trackers, health devices) (up to 5 points)
  * [ProxyAPI](https://proxyapi.ru/) was integrated in the project _(Ollama API is anticipated to be integrated in future builds)_
- [ ] Comprehensive error handling and user feedback (up to 2 points)
- [ ] Advanced animations and transitions (up to 3 points)
- [ ] Widget implementation for native mobile elements (up to 2 points)

Total points implemented: XX/30 (excluding bonus points)

Note: For each implemented feature, provide a brief description or link to the relevant implementation below the checklist.
