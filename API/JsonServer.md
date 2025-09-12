Here’s a properly formatted, step-by-step guide to installing json-server, setting up your first API, and customizing routes—along with concise explanations for each part.

## Step 1: Install json-server

**Code:**

```sh
npm install -g json-server  # global install
# OR
npm install --save-dev json-server  # local/project install
```

**Explanation:**  
This installs json-server as a CLI tool. Global installation allows you to run it anywhere in your terminal, while a project-level install is better if you want to share or version-control your setup with others.[1][2]

## Step 2: Prepare Your Data

**Code:**  
Create a file named **db.json** in your project root with the following content:

```json
{
  "posts": [
    { "id": 1, "title": "Hello, json-server!", "author": "Amit" },
    { "id": 2, "title": "Developing APIs is easy", "author": "Mohit" }
  ],
  "comments": [
    { "id": 1, "body": "Nice post!", "postId": 1 },
    { "id": 2, "body": "Very helpful!", "postId": 2 }
  ],
  "profile": { "name": "Amit Kumar" }
}
```

**Explanation:**  
Each top-level key (`posts`, `comments`, `profile`) becomes a RESTful endpoint. Json-server automatically exposes standard CRUD routes for each.[2][1]

## Step 3: Start the Server

**Code:**

```sh
json-server --watch db.json --port 3000
```

**Explanation:**  
This launches your mock API server on port 3000. The `--watch` flag tells json-server to reload if db.json changes. The server prints available endpoints on startup, and you can access the UI at **http://localhost:3000**.[1][2]

## Step 4: Test Your API

**Examples:**

- **GET** all posts:  
  `http://localhost:3000/posts`
- **GET** a single post:  
  `http://localhost:3000/posts/1`
- **POST** a new post:  
  `curl -X POST -H "Content-Type: application/json" -d "{\"title\":\"Test post\",\"author\":\"Test Author\"}" http://localhost:3000/posts`
- **PATCH** (update) a post:  
  `curl -X PATCH -H "Content-Type: application/json" -d "{\"title\":\"Updated title\"}" http://localhost:3000/posts/1`
- **DELETE** a post:  
  `curl -X DELETE http://localhost:3000/posts/1`

**Explanation:**  
Json-server provides full RESTful routes for each resource automatically. These routes match standard web API conventions.[3][1]

## Step 5: Customize Routes (Optional)

**Code:**  
Create a **routes.json** file in your project root:

```json
{
  "/api/*": "/$1",
  "/articles/:id": "/posts/:id"
}
```

Start the server with custom routes:

```sh
json-server --watch db.json --routes routes.json --port 3000
```

**Explanation:**  
The `routes.json` file lets you map any URL structure to your data. For example, `/api/posts` will serve `POST /posts`, and `/articles/2` will serve `POST /posts/2`.[4][5]

## Step 6: Run as npm Script (Optional)

**Code:**  
Add to **package.json**:

```json
"scripts": {
  "api": "json-server --watch db.json --port 3000"
}
```

**Explanation:**  
This lets you start your API with just `npm run api`—useful for automation and CI/CD pipelines.[6][7]

## Step 7: Programmatic Usage (Optional)

**Code:**  
Create a **server.js** file:

```js
const jsonServer = require('json-server')
const server = jsonServer.create()
const router = jsonServer.router('db.json')
const middlewares = jsonServer.defaults()

server.use(middlewares)
server.use('/api', router)
server.listen(3000, () => {
  console.log('JSON Server is running')
})
```

**Explanation:**  
You can run json-server as a Node module for advanced usage—custom middleware, rewriting, and serving static files with more control.[3][4]

***

## Summary Table

| Step              | Command / File                                | Purpose                                 | Endpoint Example           |
|-------------------|-----------------------------------------------|-----------------------------------------|----------------------------|
| Install           | `npm install json-server`                     | Add the tool                            | —                          |
| Prepare Data      | `db.json`                                     | Define mock data                        | —                          |
| Start Server      | `json-server --watch db.json --port 3000`     | Launch API                              | `/posts`, `/comments`      |
| Custom Routes     | `routes.json`                                 | Map custom paths                        | `/api/posts`, `/articles/1` |
| npm Script        | `package.json` scripts                        | Automate startup                        | —                          |
| Programmatic      | `server.js`                                   | Advanced customization                  | `/api/posts`               |

***

### What You’ve Learned

- **json-server** instantly gives you a mock REST API from a JSON file—no backend needed.
- Each top-level key in **db.json** becomes a resource with full CRUD endpoints.
- **routes.json** enables custom URL structures for integration with existing frontend code.
- You can launch it from the CLI, via npm script, or programmatically in Node.js.
- Great for frontend development, testing, prototyping, and learning REST APIs.

This setup is production-grade for prototyping, testing, and even local development when you need a quick backend without the overhead of coding one from scratch.[8][2][1]

[1](https://www.npmjs.com/package/json-server)
[2](https://www.geeksforgeeks.org/node-js/json-server-setup-and-introduction/)
[3](https://github.com/typicode/json-server)
[4](https://saltsthlm.github.io/protips/jsonServer.html)
[5](https://github.com/typicode/json-server/issues/1512)
[6](https://stackoverflow.com/questions/62701824/how-can-i-run-start-json-server-and-live-server-using-npm-scripts-start-in-packa)
[7](https://dev.to/avinashvagh/json-server-getting-started-4475)
[8](https://www.freecodecamp.org/news/json-server-for-frontend-development/)
