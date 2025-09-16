MongoDB is a **NoSQL** database that uses a **document-oriented** model, rather than a relational one. It stores data in **JSON-like BSON (Binary JSON)** documents, which makes it flexible, scalable, and efficient for many modern applications.

Here's a breakdown of the **MongoDB architecture**:

---

## 🔧 **1. Core Components**

### **a. Documents**

* The basic unit of data.
* Stored in BSON format (Binary JSON).
* Example:

  ```json
  {
    "_id": ObjectId("507f191e810c19729de860ea"),
    "name": "John Doe",
    "age": 30,
    "email": "john@example.com"
  }
  ```

### **b. Collections**

* Groups of documents, equivalent to tables in relational databases.
* Do not enforce schema, but documents in a collection are often related.

### **c. Databases**

* A database is a group of collections.
* Each MongoDB server can host multiple databases.

---

## 🗃️ **2. Storage Engine**

MongoDB uses different storage engines. The default is **WiredTiger**, which provides:

* Compression
* Document-level locking
* Efficient concurrency

---

## 🌐 **3. MongoDB Server (mongod)**

The core process that handles:

* Data requests
* Management
* Background tasks
  This is the main **daemon** process for MongoDB.

---

## 🔄 **4. MongoDB Client (mongo / Drivers)**

Clients (like shell or application drivers) connect to the MongoDB server using:

* The `mongo` shell
* MongoDB drivers (for Node.js, Python, Java, etc.)

---

## 🧩 **5. Replication (High Availability)**

### **Replica Set**

* A group of mongod instances that maintain the same data.
* Provides **high availability** and **data redundancy**.

#### Components:

* **Primary**: Handles all write operations.
* **Secondary**: Replicates data from the primary.
* **Arbiter**: Participates in elections but doesn’t hold data.

**Automatic failover** occurs if the primary goes down.

---

## 🧵 **6. Sharding (Horizontal Scalability)**

Sharding allows you to **scale horizontally** by distributing data across multiple machines.

### Key Components:

* **Shard**: Each shard holds part of the dataset.
* **mongos**: A routing service that directs queries to the correct shard.
* **Config servers**: Store metadata and configuration settings.

---

## 🔐 **7. Security Components**

* **Authentication**: User identity verification
* **Authorization**: Role-based access control
* **Encryption**: Data at rest and in transit

---

## 📦 MongoDB Architecture Overview Diagram (Text-based)

```
             ┌────────────┐
             │   Client   │
             └────┬───────┘
                  │
                  ▼
             ┌────────────┐
             │   mongos   │ ← Query Router (in sharded cluster)
             └────┬───────┘
        ┌─────────┴─────────┐
        ▼                   ▼
   ┌──────────┐       ┌──────────┐
   │ Shard 1  │       │ Shard 2  │   ← Each shard is a replica set
   └────┬─────┘       └────┬─────┘
        ▼                 ▼
┌──────────────┐    ┌──────────────┐
│Replica Set 1 │    │Replica Set 2 │
└──────────────┘    └──────────────┘
```

---

## ✅ Summary

| Feature          | Description                               |
| ---------------- | ----------------------------------------- |
| **Storage**      | BSON documents in collections             |
| **Scalability**  | Horizontal via sharding                   |
| **Availability** | Replica sets with automatic failover      |
| **Flexibility**  | Schema-less document model                |
| **Security**     | Authentication, Authorization, Encryption |

---

