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
