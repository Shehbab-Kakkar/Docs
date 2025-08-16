using a matrix strategy per language is a clean and scalable way to build and test each component independently (or in parallel) while maintaining a shared Cypress E2E step if needed.

🧠 Strategy

We'll break this down into:

Matrix job to build/test per language:

node, python, go, rust, java

An optional Cypress job that depends on successful builds of all services.

Each language job sets up and runs its service (optionally exposes via port).

Cypress runs after all are healthy.

✅ Matrix GitHub Actions Workflow: .github/workflows/cypress-matrix.yml
name: Polyglot CI with Cypress

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
        language: [node, python, go, rust, java]

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

      - name: Set up Node (if needed)
        if: matrix.language == 'node'
        uses: actions/setup-node@v4
        with:
          node-version: 20

      - name: Set up Python (if needed)
        if: matrix.language == 'python'
        uses: actions/setup-python@v5
        with:
          python-version: 3.11

      - name: Set up Go (if needed)
        if: matrix.language == 'go'
        uses: actions/setup-go@v5
        with:
          go-version: 1.22

      - name: Set up Rust (if needed)
        if: matrix.language == 'rust'
        uses: actions-rs/toolchain@v1
        with:
          toolchain: stable
          override: true

      - name: Set up Java (if needed)
        if: matrix.language == 'java'
        uses: actions/setup-java@v4
        with:
          distribution: 'temurin'
          java-version: '21'

      # Node
      - name: Install & Run Node App
        if: matrix.language == 'node'
        working-directory: ./frontend
        run: |
          npm install
          nohup npm start &

      # Python (Django)
      - name: Install & Run Django App
        if: matrix.language == 'python'
        working-directory: ./backend-python
        env:
          DATABASE_URL: postgres://user:password@localhost:5432/testdb
        run: |
          pip install -r requirements.txt
          python manage.py migrate
          nohup python manage.py runserver 0.0.0.0:8000 &

      # Go
      - name: Build & Run Go Service
        if: matrix.language == 'go'
        working-directory: ./backend-go
        run: |
          go build -o service
          nohup ./service &

      # Rust
      - name: Build & Run Rust Service
        if: matrix.language == 'rust'
        working-directory: ./backend-rust
        run: |
          cargo build --release
          nohup ./target/release/your-rust-service &

      # Java
      - name: Build & Run Java Service
        if: matrix.language == 'java'
        working-directory: ./backend-java
        run: |
          ./gradlew build
          nohup java -jar build/libs/your-java-service.jar &

      - name: Wait for service to boot
        run: sleep 10

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

🔧 Notes

build-and-run job runs once per language via matrix.

Each language service runs in the background via nohup.

cypress job runs after all language jobs succeed (needs: build-and-run).

You can extend matrix values (include:) to pass custom config like working-directory, port, etc.

🚀 Optional Enhancement: Matrix include for Config

If your services have different ports or directories, extend the matrix like this:

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


And access them via matrix.dir, matrix.port, etc.
