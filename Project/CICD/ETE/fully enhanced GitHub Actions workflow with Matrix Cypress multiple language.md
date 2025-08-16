fully enhanced GitHub Actions workflow with:

✅ Matrix per language
✅ Custom directories and ports
✅ Health checks on service endpoints
✅ Cypress matrix to test against each backend (Django, Go, Rust, Java)
✅ Support for environment variable injection (BACKEND_API_URL) into Cypress

✅ Final Workflow File: .github/workflows/cypress-matrix-full.yml
name: Polyglot CI with Cypress (Full Matrix)

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  build-and-run:
    name: Build & Run ${{ matrix.language }}
    runs-on: ubuntu-latest

    strategy:
      matrix:
        include:
          - language: node
            dir: frontend
            port: 3000
            health_url: http://localhost:3000

          - language: python
            dir: backend-python
            port: 8000
            health_url: http://localhost:8000/health/

          - language: go
            dir: backend-go
            port: 8080
            health_url: http://localhost:8080/health

          - language: rust
            dir: backend-rust
            port: 5000
            health_url: http://localhost:5000/health

          - language: java
            dir: backend-java
            port: 8081
            health_url: http://localhost:8081/actuator/health

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

      ### Setups ###
      - name: Set up Node.js
        if: matrix.language == 'node'
        uses: actions/setup-node@v4
        with:
          node-version: 20

      - name: Set up Python
        if: matrix.language == 'python'
        uses: actions/setup-python@v5
        with:
          python-version: 3.11

      - name: Set up Go
        if: matrix.language == 'go'
        uses: actions/setup-go@v5
        with:
          go-version: 1.22

      - name: Set up Rust
        if: matrix.language == 'rust'
        uses: actions-rs/toolchain@v1
        with:
          toolchain: stable
          override: true

      - name: Set up Java
        if: matrix.language == 'java'
        uses: actions/setup-java@v4
        with:
          distribution: 'temurin'
          java-version: 21

      ### Build & Run Services ###
      - name: Run Node.js App
        if: matrix.language == 'node'
        working-directory: ${{ matrix.dir }}
        run: |
          npm install
          nohup npm start > node.log 2>&1 &

      - name: Run Django App
        if: matrix.language == 'python'
        working-directory: ${{ matrix.dir }}
        env:
          DATABASE_URL: postgres://user:password@localhost:5432/testdb
        run: |
          pip install -r requirements.txt
          python manage.py migrate
          nohup python manage.py runserver 0.0.0.0:${{ matrix.port }} > django.log 2>&1 &

      - name: Run Go Service
        if: matrix.language == 'go'
        working-directory: ${{ matrix.dir }}
        run: |
          go build -o service
          nohup ./service > go.log 2>&1 &

      - name: Run Rust Service
        if: matrix.language == 'rust'
        working-directory: ${{ matrix.dir }}
        run: |
          cargo build --release
          nohup ./target/release/your-rust-service > rust.log 2>&1 &

      - name: Run Java Service
        if: matrix.language == 'java'
        working-directory: ${{ matrix.dir }}
        run: |
          ./gradlew build
          nohup java -jar build/libs/your-java-service.jar > java.log 2>&1 &

      ### Health Check ###
      - name: Wait for Service Health Endpoint
        run: |
          echo "Waiting for health at ${{ matrix.health_url }}"
          for i in {1..20}; do
            curl -fsS ${{ matrix.health_url }} && echo "Service is healthy!" && exit 0
            sleep 3
          done
          echo "Service health check failed."
          exit 1

  cypress:
    name: Cypress E2E (Test ${{ matrix.api_name }})
    runs-on: ubuntu-latest
    needs: build-and-run

    strategy:
      matrix:
        include:
          - api_name: Django
            api_url: http://localhost:8000

          - api_name: Go
            api_url: http://localhost:8080

          - api_name: Rust
            api_url: http://localhost:5000

          - api_name: Java
            api_url: http://localhost:8081

    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Set up Node.js
        uses: actions/setup-node@v4
        with:
          node-version: 20

      - name: Install frontend dependencies
        working-directory: ./frontend
        run: npm install

      - name: Run Cypress Tests
        uses: cypress-io/github-action@v6
        with:
          working-directory: ./frontend
          start: npm start
          wait-on: 'http://localhost:3000'
          wait-on-timeout: 120
        env:
          BACKEND_API_URL: ${{ matrix.api_url }}

🧪 Cypress Configuration (cypress.config.js)

Make sure your Cypress project reads the backend URL like this:

const { defineConfig } = require("cypress");

module.exports = defineConfig({
  e2e: {
    baseUrl: "http://localhost:3000",
    env: {
      backendUrl: process.env.BACKEND_API_URL || "http://localhost:8000",
    },
  },
});


You can then use in tests:

cy.request(`${Cypress.env('backendUrl')}/api/something`)

🔍 Debugging Tips

Use cat *.log in each matrix job to inspect service output if things go wrong.

Add echo ${{ matrix.api_url }} or curl to verify backend response before Cypress runs.
