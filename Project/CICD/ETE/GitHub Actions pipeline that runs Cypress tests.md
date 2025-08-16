GitHub Actions pipeline that runs Cypress tests for a polyglot monorepo/project that includes:

Node.js

Python (Django)

Go (Golang)

Rust

Java

This assumes you're testing a frontend with Cypress and also building backend services written in various languages.

📁 Directory Structure Assumption
.
├── frontend/            # Node.js app (Cypress tests live here)
├── backend-python/      # Django app
├── backend-go/          # Go service
├── backend-rust/        # Rust service
├── backend-java/        # Java service

✅ GitHub Actions Workflow: .github/workflows/cypress-multi-lang.yml
name: CI Pipeline with Cypress

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  setup:
    name: Setup and Build All Services
    runs-on: ubuntu-latest
    services:
      postgres:
        image: postgres:13
        ports:
          - 5432:5432
        env:
          POSTGRES_USER: user
          POSTGRES_PASSWORD: password
          POSTGRES_DB: testdb
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5

    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      ### Node.js Setup ###
      - name: Set up Node.js
        uses: actions/setup-node@v4
        with:
          node-version: 20

      - name: Install frontend dependencies
        working-directory: ./frontend
        run: npm install

      ### Python / Django Setup ###
      - name: Set up Python
        uses: actions/setup-python@v5
        with:
          python-version: '3.11'

      - name: Install Django dependencies
        working-directory: ./backend-python
        run: |
          python -m pip install --upgrade pip
          pip install -r requirements.txt

      - name: Run Django Migrations
        working-directory: ./backend-python
        env:
          DATABASE_URL: postgres://user:password@localhost:5432/testdb
        run: |
          python manage.py migrate
          python manage.py runserver &

      ### Go Setup ###
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.22'

      - name: Build Go service
        working-directory: ./backend-go
        run: go build -v ./...

      - name: Start Go service
        working-directory: ./backend-go
        run: ./your-go-binary &

      ### Rust Setup ###
      - name: Set up Rust
        uses: actions-rs/toolchain@v1
        with:
          toolchain: stable
          override: true

      - name: Build Rust service
        working-directory: ./backend-rust
        run: cargo build --release

      - name: Run Rust service
        working-directory: ./backend-rust
        run: ./target/release/your-rust-service &

      ### Java Setup ###
      - name: Set up Java
        uses: actions/setup-java@v4
        with:
          distribution: 'temurin'
          java-version: '21'

      - name: Build Java service
        working-directory: ./backend-java
        run: ./gradlew build

      - name: Run Java service
        working-directory: ./backend-java
        run: java -jar build/libs/your-java-service.jar &

      ### Wait for all services to start
      - name: Wait for services
        run: sleep 20

      ### Cypress Tests ###
      - name: Run Cypress tests
        uses: cypress-io/github-action@v6
        with:
          working-directory: ./frontend
          start: npm start
          wait-on: 'http://localhost:3000'
          wait-on-timeout: 120


📝 Notes

Cypress: Runs against the frontend (npm start assumed).

Each service is built and started in background (&) — useful for monolith + microservices setups.

Database support via PostgreSQL service container.

Make sure all services bind to 0.0.0.0 if accessed externally (important in GitHub Actions).

✅ Customize

Replace ./your-go-binary, your-rust-service, and your-java-service.jar with actual names.

Modify DATABASE_URL or ports depending on your service setup.

Add service health checks or use wait-on to ensure they're ready.
