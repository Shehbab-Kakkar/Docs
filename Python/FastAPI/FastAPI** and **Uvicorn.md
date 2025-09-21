To get started with **FastAPI** and **Uvicorn**, here's the **standard installation and setup process**:

---

## ✅ Step-by-Step Installation

### 1. **Create a virtual environment (optional but recommended)**

```bash
python -m venv venv
source venv/bin/activate  # On Windows use: venv\Scripts\activate
```

---

### 2. **Install FastAPI and Uvicorn**

```bash
pip install fastapi uvicorn
```

> Optionally, for full development support (e.g. auto-reload):

```bash
pip install "uvicorn[standard]"
```

---

### 3. **Create a simple FastAPI app**

Create a file called `main.py`:

```python
# main.py
from fastapi import FastAPI

app = FastAPI()

@app.get("/")
def read_root():
    return {"message": "Hello, FastAPI!"}
```

---

### 4. **Run the app with Uvicorn**

```bash
uvicorn main:app --reload
```

> `--reload` enables auto-reload during development (only use in dev, not production).

---

## 📌 Notes

* `main:app` refers to:

  * `main.py` file
  * `app` is the FastAPI instance inside it.
* Open your browser at: [http://127.0.0.1:8000](http://127.0.0.1:8000)
* Interactive API docs available at:

  * Swagger: [http://127.0.0.1:8000/docs](http://127.0.0.1:8000/docs)
  * ReDoc: [http://127.0.0.1:8000/redoc](http://127.0.0.1:8000/redoc)

---

Let me know if you want Docker setup, production tips, or anything more advanced!
