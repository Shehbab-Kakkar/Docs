You can perform insert, update, and delete operations on a database using a Python database connector, such as sqlite3 for a SQLite database or psycopg2 for PostgreSQL. Here's a general guide on how to perform these operations using the sqlite3 module.
Prerequisites
First, you'll need to import the sqlite3 module and establish a connection to your database. You'll also need a cursor object, which allows you to execute SQL commands.
import sqlite3

# Connect to the database (creates a file if it doesn't exist)
conn = sqlite3.connect('example.db')

# Create a cursor object
cursor = conn.cursor()

# Create a sample table
cursor.execute('''
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL,
        email TEXT NOT NULL
    )
''')

# Commit the changes to the database
conn.commit()

Insert Operation
To insert a new record into a table, you use the INSERT INTO SQL command. It's best practice to use placeholders (like ? for sqlite3) to prevent SQL injection attacks.
# SQL query to insert a new user
sql_insert = "INSERT INTO users (name, email) VALUES (?, ?)"

# Data to be inserted
new_user = ('Alice', 'alice@example.com')

# Execute the query
cursor.execute(sql_insert, new_user)

# Commit the changes
conn.commit()
print("Record inserted successfully!")

Update Operation
To modify an existing record, you use the UPDATE SQL command. The WHERE clause is crucial here, as it specifies which record(s) to update. Without it, you would update every record in the table.
# SQL query to update a user's email
sql_update = "UPDATE users SET email = ? WHERE name = ?"

# New email and the name of the user to update
updated_data = ('alice.smith@example.com', 'Alice')

# Execute the query
cursor.execute(sql_update, updated_data)

# Commit the changes
conn.commit()
print("Record updated successfully!")

Delete Operation
To remove a record from a table, you use the DELETE FROM SQL command. Like with the update operation, the WHERE clause is essential to specify which record to delete.
# SQL query to delete a user
sql_delete = "DELETE FROM users WHERE name = ?"

# Name of the user to delete
user_to_delete = ('Alice',)

# Execute the query
cursor.execute(sql_delete, user_to_delete)

# Commit the changes
conn.commit()
print("Record deleted successfully!")

Closing the Connection
After you've finished all your operations, you should always close the cursor and the database connection to free up resources.
# Close the cursor and the connection
cursor.close()
conn.close()

