

## **1. Table: `account` as `A`**

```sql
Table account as A {
  id bigserial [pk]
  owner varchar [not null]
  balancer bigint [not null] 
  currency varchar [not null]
  created_at timestampz [not null, default:`now()`] 
}
```
- **Primary Key:** `id`
- This table stores account information.
- Each account has a unique `id`.

---

## **2. Table: `entries`**

```sql
Table entries {
  big bigserial [pk]
  account_id bigint [ref: > A.id]
  amount bigint [not null]
  created_at timestampz [not null, default:`now()`]
}
```
- **Primary Key:** `big`
- **Foreign Key:**  
  - `account_id bigint [ref: > A.id]`
    - This means `account_id` in the `entries` table **references** the `id` column in the `account` table (aliased as `A`).
    - **Purpose:** Every entry must be linked to a valid account.
    - **Effect:** If you try to insert an entry with an `account_id` that does not exist in `account.id`, the database will reject it.

**Relationship:**  
- **One account** can have **many entries** (one-to-many).
- **Each entry** belongs to **one account**.

---

## **3. Table: `transfers`**

```sql
Table transfers {
 id bigserial [pk]
 from_account_id bigint [ref: > A.id]
 to_account_id bigint [ref: > A.id]
 created_at timestampz [not null, default:`now()`]
}
```
- **Primary Key:** `id`
- **Foreign Keys:**  
  - `from_account_id bigint [ref: > A.id]`
  - `to_account_id bigint [ref: > A.id]`
    - Both columns reference `id` in the `account` table.
    - **Purpose:** Records a transfer between two accounts.
    - **Effect:** Both the sending and receiving accounts must exist in the `account` table.

**Relationship:**  
- **Each transfer** is linked to **two accounts** (source and destination).
- **Accounts** can appear in many transfers (again, one-to-many from accounts to transfers).

---

## **4. Table: `orders`**

```sql
Table orders {
 id int [pk]
 user_id in [not null, unique]
 status varchar
 created_at varchar [note: 'when order created']
}
```
- **Primary Key:** `id`
- **Foreign Key:**  
  - `user_id in [not null, unique]`
    - This is **not a valid foreign key** syntax.  
    - `in` is not a type; typically, you’d see:  
      ```sql
      user_id int [not null, unique, ref: > users.id]
      ```
    - You need to specify which table and column `user_id` references.

---

## **How Do Foreign Keys Work Here?**

- **Foreign keys define relationships between tables** by referencing the primary key of another table.
- **They enforce referential integrity:** You can’t insert a value in a child table if it doesn’t exist in the parent table.
- **In your schema:**
    - `entries.account_id`, `transfers.from_account_id`, and `transfers.to_account_id` **must** match an `account.id`.
    - If you try to delete an `account` that is referenced by an `entry` or `transfer`, the database will block the operation unless you specify `ON DELETE CASCADE` or similar behavior.

---

## **Diagram**

```
+---------+      +---------+     +-----------+
| account |<-----| entries |     | transfers |
+---------+      +---------+     +-----------+
| id      |      | account_id |-->| from_account_id (FK to account.id)
| ...     |      | ...      |    | to_account_id   (FK to account.id)
+---------+      +---------+     +-----------+
```

---

## **Summary Table**

| Table      | Column Name         | Foreign Key Reference      | Relationship                |
|------------|--------------------|---------------------------|-----------------------------|
| entries    | account_id         | account.id                | Many entries per account    |
| transfers  | from_account_id    | account.id                | Many transfers per account  |
| transfers  | to_account_id      | account.id                | Many transfers per account  |
| orders     | user_id            | (missing FK; should be users.id) | (should reference users)    |

---

**In summary:**  
Foreign keys in your schema ensure that entries and transfers always reference valid accounts, enforcing data integrity and representing real-world relationships between the tables.  
The `orders` table is missing a correct foreign key definition for `user_id`. 

If you want an example query or a correction for the `orders` table, let me know!
