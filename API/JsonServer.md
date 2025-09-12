The quickest way to spin up a mock REST API is to use json-server with a simple db.json file and run it locally on a chosen port. Below are copy-paste steps for install, creating endpoints, running, and customizing routes.[1][2]

## Quick start
- Install json-server globally or as a dev dependency: 
  - Global: npm install -g json-server[3][4]
  - Local: npm install --save-dev json-server[5][6]
- Create a db.json file with collections; each top-level key becomes an endpoint.[7][3]
- Run the server: json-server --watch db.json --port 3000[1][3]

## Example db.json
- Create db.json in the project root with sample collections:[7][3]
  
  {
    "posts": [
      { "id": 1, "title": "JSON-Server", "author": "Amit" },
      { "id": 2, "title": "Node.js", "author": "Mohit" }
    ],
    "comments": [
      { "id": 1, "body": "Great post!", "postId": 1 },
      { "id": 2, "body": "Informative!", "postId": 2 }
    ],
    "profile": { "name": "Amit Kumar" }
  }
  
- This gives endpoints: /posts, /comments, /profile and standard REST routes.[3][1]

## Run and test
- Start: json-server --watch db.json --port 3000[1][3]
- Default base URL: http://localhost:3000[3][1]
- Built-in REST routes per resource:[2][1]
  - GET /posts (list), GET /posts/1 (single)[2][1]
  - POST /posts (create), PUT/PATCH /posts/1 (update), DELETE /posts/1 (delete)[2][1]
- The CLI prints available resources and a Home page showing routes.[8][1]

## Querying and filters
- Pagination: ?_page=1&_limit=10[9][1]
- Sorting: ?_sort=title&_order=asc[9][1]
- Full-text search: ?q=node[9][1]
- Relational: /comments?postId=1 filters by foreign key fields.[1][3]

## Custom routes (rewrite)
- Create routes.json to map friendly or nested URLs:[10][11]
  
  {
    "/api/*": "/$1",
    "/carts/:cartId/items/:itemId": "/items/:itemId"
  }
  
- Run: npx json-server --watch db.json --routes routes.json --port 8080[11][10]
- This helps when a frontend expects specific paths like /api/heroes.[12][13]

## Run via npm scripts
- Install locally and add scripts to package.json:[14][15]
  
  {
    "scripts": {
      "api": "json-server --watch db.json --port 3000"
    }
  }
  
- Start with: npm run api[15][14]

## Programmatic server (Node API)
- For advanced control, use json-server in a small server.js:[10][2]
  
  const jsonServer = require('json-server');
  const server = jsonServer.create();
  const router = jsonServer.router('db.json');
  const middlewares = jsonServer.defaults();
  server.use(middlewares);
  server.use(jsonServer.rewrite({"/api/*": "/$1"}));
  server.use(router);
  server.listen(3000, () => console.log('JSON Server is running'));
  
- Run: node server.js[10][2]

## Tips and extras
- Static hosting: place assets in ./public or use --static ./public to serve HTML/CSS/JS alongside the API.[9][1]
- Snapshot: press s in the terminal to create a data snapshot while running.[8][1]
- Great for frontend dev and Postman testing with zero backend code.[16][7]

If a tailored db.json schema is needed, share the entities and fields, and a ready-to-run file plus routes.json will be provided.[7][3]

[1](https://www.npmjs.com/package/json-server)
[2](https://github.com/typicode/json-server)
[3](https://www.geeksforgeeks.org/node-js/json-server-setup-and-introduction/)
[4](https://json-server.dev/json-server-installation/)
[5](https://www.digitalocean.com/community/tutorials/json-server)
[6](https://www.dhiwise.com/post/how-to-use-json-server-in-frontend-development)
[7](https://www.freecodecamp.org/news/json-server-for-frontend-development/)
[8](https://dev.to/aguowisdom/creating-a-restful-api-with-json-server-in-nodejs-a-comprehensive-guide-for-beginners-2goc)
[9](https://www.npmjs.com/package/json-server/v/0.17.4)
[10](https://saltsthlm.github.io/protips/jsonServer.html)
[11](https://www.sitepoint.com/json-server-example/)
[12](https://github.com/typicode/json-server/issues/1512)
[13](https://stackoverflow.com/questions/57005091/path-with-in-json-server-db-json)
[14](https://dev.to/avinashvagh/json-server-getting-started-4475)
[15](https://stackoverflow.com/questions/62701824/how-can-i-run-start-json-server-and-live-server-using-npm-scripts-start-in-packa)
[16](https://www.linkedin.com/pulse/setting-up-local-json-server-testing-api-postman-laljan-basha-shaik-fkbmc)
[17](https://www.geeksforgeeks.org/node-js/how-to-create-a-rest-api-using-json-server-npm-package/)
[18](https://vsys.host/how-to/json-server-installation-setting-up-mock-apis)
[19](https://contabo.com/blog/json-server-installation-setting-up-mock-apis/)
[20](http://codesandbox.io/p/github/cloud-walker/json-server)
