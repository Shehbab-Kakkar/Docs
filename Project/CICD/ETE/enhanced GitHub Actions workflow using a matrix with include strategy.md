Perfect! Here's the enhanced GitHub Actions workflow using a matrix with include strategy, allowing you to configure:

Custom working directories

Custom ports

Easier scaling of your services

✅ Full GitHub Actions Workflow: .github/workflows/cypress-matrix-enhanced.yml
name: Polyglot CI with Cypress (Matrix Enhanced)

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

          - language: python
            dir: backend-python
            port: 8000

          - language: go
            dir: backend-go
            port: 8080

          - language: rust
            dir: backend-rust
            port: 5000

          - language: java
            dir: backend-java
            port: 8081

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

      ### Setup environments based on language ###
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

      ### Build & Start Services ###
      - name: Run Node.js Service
        if: matrix.language == 'node'
        working-directory: ${{ matrix.dir }}
        run: |
          npm install
          nohup npm start > node.log 2>&1 &

      - name: Run Django Service
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

      - name: Wait for Service to be Ready
        run: |
          echo "Waiting for port ${{ matrix.port }}..."
          for i in {1..20}; do
            nc -zv 127.0.0.1 ${{ matrix.port }} && echo "Service is up!" && exit 0
            sleep 3
          done
          echo "Service failed to start"
          exit 1

  cypress:
    name: Cypress E2E Tests
    runs-on: ubuntu-latest
    needs: build-and-run

    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Set up Node.js
        uses: actions/setup-node@v4
        with:
          node-version: 20

      - name: Install dependencies
        working-directory: ./frontend
        run: npm install

      - name: Run Cypress Tests
        uses: cypress-io/github-action@v6
        with:
          working-directory: ./frontend
          start: npm start
          wait-on: 'http://localhost:3000'
          wait-on-timeout: 120

📦 Assumptions

All services bind to 0.0.0.0

Cypress targets localhost:3000

You can adjust the wait-on: URL to test against other service endpoints too if needed.

🔍 Tip: Logs for Debugging

Each service writes to a log file (node.log, django.log, etc.) — you can add a cat command if debugging is needed:

- name: Print Logs
  run: cat django.log


Would you like to add service health checks, or have Cypress run against multiple service endpoints?
