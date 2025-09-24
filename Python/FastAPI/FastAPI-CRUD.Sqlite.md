Here's a **complete example** of a FastAPI application that performs **CRUD operations** using **SQLite** with **SQLAlchemy** as the ORM.

---

## ✅ Features

* Uses **FastAPI** for building the API
* Uses **SQLite** as the database (in-memory or file)
* Uses **SQLAlchemy** for ORM
* Basic **CRUD operations**: Create, Read, Update, Delete
* Pydantic for data validation

---

## 📁 Project Structure

```
fastapi_sqlite_crud/
├── main.py
└── models.py
```

---

## 🐍 Step 1: Install Requirements

```bash
pip install fastapi uvicorn sqlalchemy pydantic
```

---

## 📄 `models.py`

```python
from sqlalchemy import Column, Integer, String, create_engine
from sqlalchemy.orm import declarative_base, sessionmaker

DATABASE_URL = "sqlite:///./test.db"

Base = declarative_base()
engine = create_engine(DATABASE_URL, connect_args={"check_same_thread": False})
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)


class Item(Base):
    __tablename__ = "items"

    id = Column(Integer, primary_key=True, index=True)
    name = Column(String, index=True)
    description = Column(String, index=True)
```

---

## 📄 `main.py`

```python
from fastapi import FastAPI, HTTPException, Depends
from sqlalchemy.orm import Session
from pydantic import BaseModel
from models import Base, engine, SessionLocal, Item

# Create the database tables
Base.metadata.create_all(bind=engine)

app = FastAPI()

# Dependency to get DB session
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

# Pydantic schemas
class ItemCreate(BaseModel):
    name: str
    description: str

class ItemRead(ItemCreate):
    id: int

    class Config:
        orm_mode = True


# Create Item
@app.post("/items/", response_model=ItemRead)
def create_item(item: ItemCreate, db: Session = Depends(get_db)):
    db_item = Item(name=item.name, description=item.description)
    db.add(db_item)
    db.commit()
    db.refresh(db_item)
    return db_item

# Read all Items
@app.get("/items/", response_model=list[ItemRead])
def read_items(skip: int = 0, limit: int = 10, db: Session = Depends(get_db)):
    return db.query(Item).offset(skip).limit(limit).all()

# Read one Item
@app.get("/items/{item_id}", response_model=ItemRead)
def read_item(item_id: int, db: Session = Depends(get_db)):
    item = db.query(Item).filter(Item.id == item_id).first()
    if not item:
        raise HTTPException(status_code=404, detail="Item not found")
    return item

# Update Item
@app.put("/items/{item_id}", response_model=ItemRead)
def update_item(item_id: int, updated_item: ItemCreate, db: Session = Depends(get_db)):
    item = db.query(Item).filter(Item.id == item_id).first()
    if not item:
        raise HTTPException(status_code=404, detail="Item not found")
    item.name = updated_item.name
    item.description = updated_item.description
    db.commit()
    db.refresh(item)
    return item

# Delete Item
@app.delete("/items/{item_id}")
def delete_item(item_id: int, db: Session = Depends(get_db)):
    item = db.query(Item).filter(Item.id == item_id).first()
    if not item:
        raise HTTPException(status_code=404, detail="Item not found")
    db.delete(item)
    db.commit()
    return {"message": "Item deleted"}
```

---

## ▶️ Run the app

```bash
uvicorn main:app --reload
```

Visit `http://127.0.0.1:8000/docs` to access the **interactive Swagger UI**.

---

## ✅ Example JSON (for POST or PUT)

```json
{
  "name": "Test Item",
  "description": "A test item description"
}
```

---

Would you like to turn this into a package or add user authentication next?

Great! Let's break this down into **two goals**:

---

## ✅ Goal 1: Package the App (Modular Structure)

We'll restructure your app into a clean **package-style project**.

### 📦 Project Structure (after refactor)

```
fastapi_sqlite_crud/
├── app/
│   ├── __init__.py
│   ├── main.py
│   ├── models.py
│   ├── database.py
│   ├── schemas.py
│   ├── crud.py
│   └── auth/
│       ├── __init__.py
│       ├── models.py
│       ├── schemas.py
│       ├── utils.py
│       └── routes.py
├── requirements.txt
└── run.sh
```

---

## ✅ Goal 2: Add Authentication (JWT + Password Hashing)

We’ll add:

* **User Registration**
* **Login with JWT**
* **Password Hashing (using `passlib`)**
* **JWT Auth Middleware (using `fastapi.security`)**

---

## 🛠 Step-by-step Breakdown

---

### 1. `app/database.py` – Database Setup

```python
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker, declarative_base

DATABASE_URL = "sqlite:///./test.db"

engine = create_engine(DATABASE_URL, connect_args={"check_same_thread": False})
SessionLocal = sessionmaker(bind=engine, autocommit=False, autoflush=False)
Base = declarative_base()
```

---

### 2. `app/models.py` – Item Model

```python
from sqlalchemy import Column, Integer, String
from .database import Base

class Item(Base):
    __tablename__ = "items"
    id = Column(Integer, primary_key=True, index=True)
    name = Column(String, index=True)
    description = Column(String, index=True)
```

---

### 3. `app/schemas.py` – Item Schemas

```python
from pydantic import BaseModel

class ItemCreate(BaseModel):
    name: str
    description: str

class ItemRead(ItemCreate):
    id: int

    class Config:
        orm_mode = True
```

---

### 4. `app/crud.py` – CRUD Logic

```python
from sqlalchemy.orm import Session
from . import models, schemas

def create_item(db: Session, item: schemas.ItemCreate):
    db_item = models.Item(name=item.name, description=item.description)
    db.add(db_item)
    db.commit()
    db.refresh(db_item)
    return db_item

def get_item(db: Session, item_id: int):
    return db.query(models.Item).filter(models.Item.id == item_id).first()

def get_items(db: Session, skip: int = 0, limit: int = 10):
    return db.query(models.Item).offset(skip).limit(limit).all()

def update_item(db: Session, item_id: int, item_data: schemas.ItemCreate):
    item = get_item(db, item_id)
    if item:
        item.name = item_data.name
        item.description = item_data.description
        db.commit()
        db.refresh(item)
    return item

def delete_item(db: Session, item_id: int):
    item = get_item(db, item_id)
    if item:
        db.delete(item)
        db.commit()
    return item
```

---

### 5. `app/auth/models.py` – User Model

```python
from sqlalchemy import Column, Integer, String
from ..database import Base

class User(Base):
    __tablename__ = "users"
    id = Column(Integer, primary_key=True, index=True)
    username = Column(String, unique=True, index=True)
    hashed_password = Column(String)
```

---

### 6. `app/auth/schemas.py` – User Schemas

```python
from pydantic import BaseModel

class UserCreate(BaseModel):
    username: str
    password: str

class UserRead(BaseModel):
    id: int
    username: str

    class Config:
        orm_mode = True

class Token(BaseModel):
    access_token: str
    token_type: str = "bearer"
```

---

### 7. `app/auth/utils.py` – Security Utilities

```python
from passlib.context import CryptContext
from jose import JWTError, jwt
from datetime import datetime, timedelta

# Password hashing
pwd_context = CryptContext(schemes=["bcrypt"], deprecated="auto")

SECRET_KEY = "super-secret-key"  # You should use env vars
ALGORITHM = "HS256"
ACCESS_TOKEN_EXPIRE_MINUTES = 30

def hash_password(password: str):
    return pwd_context.hash(password)

def verify_password(plain_password, hashed_password):
    return pwd_context.verify(plain_password, hashed_password)

def create_access_token(data: dict, expires_delta: timedelta | None = None):
    to_encode = data.copy()
    expire = datetime.utcnow() + (expires_delta or timedelta(minutes=15))
    to_encode.update({"exp": expire})
    return jwt.encode(to_encode, SECRET_KEY, algorithm=ALGORITHM)
```

---

### 8. `app/auth/routes.py` – Auth Routes

```python
from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from ..database import SessionLocal
from . import schemas, models, utils
from fastapi.security import OAuth2PasswordBearer, OAuth2PasswordRequestForm

router = APIRouter()
oauth2_scheme = OAuth2PasswordBearer(tokenUrl="auth/login")

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

@router.post("/register", response_model=schemas.UserRead)
def register(user: schemas.UserCreate, db: Session = Depends(get_db)):
    db_user = db.query(models.User).filter(models.User.username == user.username).first()
    if db_user:
        raise HTTPException(status_code=400, detail="Username already registered")
    hashed_pw = utils.hash_password(user.password)
    new_user = models.User(username=user.username, hashed_password=hashed_pw)
    db.add(new_user)
    db.commit()
    db.refresh(new_user)
    return new_user

@router.post("/login", response_model=schemas.Token)
def login(form_data: OAuth2PasswordRequestForm = Depends(), db: Session = Depends(get_db)):
    user = db.query(models.User).filter(models.User.username == form_data.username).first()
    if not user or not utils.verify_password(form_data.password, user.hashed_password):
        raise HTTPException(status_code=401, detail="Invalid credentials")
    token = utils.create_access_token(data={"sub": user.username})
    return {"access_token": token, "token_type": "bearer"}
```

---

### 9. `app/main.py` – Entry Point

```python
from fastapi import FastAPI, Depends, HTTPException
from sqlalchemy.orm import Session
from .database import Base, engine, SessionLocal
from . import models, schemas, crud
from .auth import routes as auth_routes, models as user_models

# Create tables
Base.metadata.create_all(bind=engine)

app = FastAPI()

# Register auth routes
app.include_router(auth_routes.router, prefix="/auth", tags=["auth"])

# Dependency
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()

# Protected CRUD endpoints
@app.post("/items/", response_model=schemas.ItemRead)
def create(item: schemas.ItemCreate, db: Session = Depends(get_db)):
    return crud.create_item(db, item)

@app.get("/items/", response_model=list[schemas.ItemRead])
def read_items(skip: int = 0, limit: int = 10, db: Session = Depends(get_db)):
    return crud.get_items(db, skip, limit)

@app.get("/items/{item_id}", response_model=schemas.ItemRead)
def read_item(item_id: int, db: Session = Depends(get_db)):
    item = crud.get_item(db, item_id)
    if not item:
        raise HTTPException(status_code=404, detail="Item not found")
    return item

@app.put("/items/{item_id}", response_model=schemas.ItemRead)
def update(item_id: int, item_data: schemas.ItemCreate, db: Session = Depends(get_db)):
    item = crud.update_item(db, item_id, item_data)
    if not item:
        raise HTTPException(status_code=404, detail="Item not found")
    return item

@app.delete("/items/{item_id}")
def delete(item_id: int, db: Session = Depends(get_db)):
    item = crud.delete_item(db, item_id)
    if not item:
        raise HTTPException(status_code=404, detail="Item not found")
    return {"message": "Item deleted"}
```

---

### 🔐 Secure Routes with JWT (Optional)

If you want to secure CRUD endpoints:

```python
from fastapi.security import OAuth2PasswordBearer
from fastapi import Security
from jose import jwt, JWTError

oauth2_scheme = OAuth2PasswordBearer(tokenUrl="/auth/login")

def get_current_user(token: str = Depends(oauth2_scheme)):
    try:
        payload = jwt.decode(token, utils.SECRET_KEY, algorithms=[utils.ALGORITHM])
        username: str = payload.get("sub")
        if username is None:
            raise HTTPException(status_code=401, detail="Invalid token")
        return username
    except JWTError:
        raise HTTPException(status_code=401, detail="Invalid token")
```

---

### 📦 `requirements.txt`

```txt
fastapi
uvicorn
sqlalchemy
pydantic
passlib[bcrypt]
python-jose
```

