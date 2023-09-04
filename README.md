# NotesTask

NotesTask is a simple CRUD (Create, Read, Update, Delete) application built using the GIN web framework. It provides endpoints to manage notes.

## Getting Started

Follow these steps to run the NotesTask application using Docker Compose:

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/abdullabasim/NotesTask.git
   cd NotesTask-master
   ```

2. Build and start the application using Docker Compose:

   ```bash
   docker-compose up --build
   ```

3. Once the application is running, you can access the Swagger API documentation by visiting [http://localhost:8080/docs/index.html](http://localhost:8080/docs/index.html) in your web browser. This provides an interactive interface to explore and test the available API endpoints.

## API Endpoints

The following API endpoints are available:

- **GET** `/api/v1/`: Returns a simple message for the main endpoint.
- **POST** `/api/v1/note`: Creates a new note with a title and text.
- **GET** `/api/v1/notes`: Retrieves a list of notes with pagination.
- **PUT** `/api/v1/note/{id}`: Updates an existing note by ID.
- **DELETE** `/api/v1/notes/`: Deletes one or more notes by IDs.
