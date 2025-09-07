Absolutely! You can use **Enum** types in your schema to restrict a column to a set of predefined values. Enums are common for fields like `status`, `currency`, etc.

Below is an example **using your schema dialect** (like [dbdiagram.io](https://dbdiagram.io/docs#enums)), showing how to define and use enums.

---

## **1. Define an Enum**

Let's add an enum for `currency` and for `orders.status`:

```sql
Enum currency_type {
  USD
  EUR
  PKR
  INR
}

Enum order_status {
  pending
  completed
  cancelled
}
```

---

## **2. Use Enum in Tables**

Modify your existing tables to use these enums:

```sql
Table account as A {
  id bigserial [pk]
  owner varchar [not null]
  balancer bigint [not null] 
  currency currency_type [not null]  // Now uses enum
  created_at timestampz [not null, default:`now()`] 
}

Table orders {
  id int [pk]
  user_id int [not null, unique]
  status order_status // Now uses enum
  created_at varchar [note: 'when order created']
}
```

---

## **3. Example: Insert Statements**

```sql
INSERT INTO account (id, owner, balancer, currency, created_at)
VALUES (1, 'Alice', 10000, 'USD', now());

INSERT INTO orders (id, user_id, status, created_at)
VALUES (1, 99, 'pending', '2025-09-07 05:00:00');
```

- If you try to insert a `currency` not in the enum (`currency_type`), the DB will reject it.
- Same for `status` in orders.

---

## **Summary Table**

| Enum Name      | Possible Values            | Used In      | Example Value |
|----------------|---------------------------|--------------|--------------|
| currency_type  | USD, EUR, PKR, INR        | account      | 'USD'        |
| order_status   | pending, completed, cancelled | orders   | 'pending'    |

---

**In summary:**  
Enums help you enforce data integrity by limiting a column to a defined set of values in your schema.  
If you want SQL/PostgreSQL syntax examples or further explanation, let me know!
