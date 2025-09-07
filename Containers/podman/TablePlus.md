To **list all databases and their tables** in PostgreSQL (using TablePlus, psql, or any SQL client), you’ll need to:

1. Query to get all databases.
2. For each database, list the tables.

### **1. List All Databases**

Run this query (works in TablePlus and any SQL editor):

```sql
SELECT datname FROM pg_database WHERE datistemplate = false;
```

### **2. List All Tables in Each Database**

You need to connect to each database and then run:

```sql
SELECT schemaname, tablename
FROM pg_catalog.pg_tables
WHERE schemaname NOT IN ('pg_catalog', 'information_schema');
```

This will list all user tables in the current database.

---

## **Automated Script (psql only)**

If you’re using the `psql` terminal and want to see all tables in all databases in one go, run this shell script on your server (replace `root` and `password` as needed):

```sh
PGPASSWORD=password psql -U root -d postgres -t -c "SELECT datname FROM pg_database WHERE datistemplate = false;" | while read dbname; do
  echo "Database: $dbname"
  PGPASSWORD=password psql -U root -d "$dbname" -c "\dt"
done
```

---

## **In TablePlus GUI**
- Connect to your Postgres server.
- The left sidebar will show all databases.
- Click on each database, then expand “Tables” to see all tables in that database.

---

## **Summary Table**

| Task         | Query/Action                                                                                 |
|--------------|---------------------------------------------------------------------------------------------|
| List DBs     | `SELECT datname FROM pg_database WHERE datistemplate = false;`                              |
| List tables  | `SELECT schemaname, tablename FROM pg_catalog.pg_tables WHERE schemaname NOT IN ('pg_catalog', 'information_schema');` (run inside each db) |

---

Let me know if you want a custom SQL function or script to view all tables in all databases in a single query output!

--
```go
SELECT column_name, data_type, is_nullable, column_default
FROM information_schema.columns
WHERE table_name = 'accounts';

SELECT column_name, data_type, is_nullable, column_default
FROM information_schema.columns
WHERE table_name = 'entries';

SELECT column_name, data_type, is_nullable, column_default
FROM information_schema.columns
WHERE table_name = 'transfers';
```
