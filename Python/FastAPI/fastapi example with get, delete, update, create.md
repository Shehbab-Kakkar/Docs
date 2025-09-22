fastapi example with get, delete, update, create.

Here is a simple FastAPI example implementing all CRUD operations—GET (read), POST (create), PUT (update), and DELETE—using an in-memory list to store data. This provides a quick way to learn how FastAPI endpoints work for typical resource management scenarios[1][2][3].

### Full Example: FastAPI CRUD (GET, POST, PUT, DELETE)

```python
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from typing import List, Optional

app = FastAPI()

# Pydantic model for data validation
class Item(BaseModel):
    id: int
    name: str
    description: Optional[str] = None

# In-memory storage for items
items = []

# CREATE - Add new item
@app.post("/items", response_model=Item)
async def create_item(item: Item):
    items.append(item)
    return item

# READ - Get all items
@app.get("/items", response_model=List[Item])
async def read_items():
    return items

# READ - Get item by ID
@app.get("/items/{item_id}", response_model=Item)
async def read_item(item_id: int):
    for item in items:
        if item.id == item_id:
            return item
    raise HTTPException(status_code=404, detail="Item not found")

# UPDATE - Update item by ID (replace)
@app.put("/items/{item_id}", response_model=Item)
async def update_item(item_id: int, updated_item: Item):
    for idx, item in enumerate(items):
        if item.id == item_id:
            items[idx] = updated_item
            return updated_item
    raise HTTPException(status_code=404, detail="Item not found")

# DELETE - Remove item by ID
@app.delete("/items/{item_id}")
async def delete_item(item_id: int):
    for idx, item in enumerate(items):
        if item.id == item_id:
            del items[idx]
            return {"message": "Item deleted"}
    raise HTTPException(status_code=404, detail="Item not found")
```
- Use `uvicorn main:app --reload` to run this application[3].
- Test endpoints using `/docs` (Swagger UI, built into FastAPI)[2].

### Key Points

- `@app.post("/items")`: Adds a new item (Create)[1].
- `@app.get("/items")`: Retrieves all items (Read)[1][2].
- `@app.get("/items/{item_id}")`: Retrieves single item by ID (Read)[2].
- `@app.put("/items/{item_id}")`: Updates an item by ID (Update)[1][2].
- `@app.delete("/items/{item_id}")`: Deletes an item by ID (Delete)[1][2].

This approach keeps things in-memory for simplicity; for real applications, connect with a database like SQLite or Postgres using ORM tools such as SQLAlchemy[4][5].

Citations:
[1] How to Create CRUD Operations with FastAPI Quickly https://apidog.com/blog/how-to-quickly-implement-crud-operations-with-fastapi/
[2] FastAPI - CRUD Operations https://www.tutorialspoint.com/fastapi/fastapi_crud_operations.htm
[3] Build a Simple CRUD API Using FastAPI (Python) - Shift Asia https://shiftasia.com/community/build-a-simple-crud-api-using-fastapi-python-tutorial/
[4] FastAPI - Crud Operations https://www.geeksforgeeks.org/python/fastapi-crud-operations/
[5] Build a CRUD App with FastAPI and SQLAlchemy https://github.com/wpcodevo/fastapi_sqlalchemy
[6] My First CRUD App With Fast API https://betterprogramming.pub/my-first-crud-app-with-fast-api-74ac190d2dcc
[7] FastAPI Essentials - Creating a CRUD API with GET, POST ... https://www.youtube.com/watch?v=gV-EpY2TeQ0
[8] Build an App With FastAPI for Python https://kinsta.com/blog/fastapi/
[9] Crud Basics Best Practices https://notes.kodekloud.com/docs/Python-API-Development-with-FastAPI/FastAPI-Basics/Crud-Basics-Best-Practices
[10] Developing a Single Page App with FastAPI and React https://testdriven.io/blog/fastapi-react/
[11] How to implement all CRUD Create, Read, Update and ... https://www.youtube.com/watch?v=ZlfSG_ptueI
[12] Building a CRUD API with FastAPI and MongoDB https://dev.to/aquibpy/building-a-crud-api-with-fastapi-and-mongodb-32n
