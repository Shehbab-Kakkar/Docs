Perfect! Here's a simple example of **Next.js App Router navigation** with:

* A **Home page** at `/`
* A **Login page** at `/login`
* Navigation links between them
* No CSS — just pure structure for understanding the concept

---

## ✅ Project Structure (App Router)

Assuming you're using the `App Router` with `src/` folder:

```
/src
  /app
    /login
      page.js      ← Login page
    /home
      page.js      ← Home page
    layout.js      ← Shared layout (with links)
    page.js        ← Default route (optional, can redirect to /home)
```

---

## 1. 📄 `src/app/layout.js`

This file defines a **shared layout**, and appears on **every page**.

```jsx
// src/app/layout.js

export const metadata = {
  title: 'Simple Navigation Example',
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>
        <nav>
          <a href="/home">Home</a> | <a href="/login">Login</a>
        </nav>
        <hr />
        {children}
      </body>
    </html>
  );
}
```

### 🔍 Explanation:

* `<nav>` contains simple links using native `<a>` tags.
* This layout wraps all your pages.

---

## 2. 🏠 `src/app/home/page.js` – Home Page

```jsx
// src/app/home/page.js

import React from 'react';

const HomePage = () => {
  return (
    <div>
      <h1>Home Page</h1>
      <p>Welcome to the home page!</p>
    </div>
  );
};

export default HomePage;
```

---

## 3. 🔐 `src/app/login/page.js` – Login Page

```jsx
// src/app/login/page.js

import React from 'react';

const LoginPage = () => {
  return (
    <div>
      <h1>Login Page</h1>
      <p>Please enter your login credentials.</p>
    </div>
  );
};

export default LoginPage;
```

---

## 4. 🔁 (Optional) `src/app/page.js` – Redirect to Home (optional)

This is what loads when visiting `/`. You can either:

### Option A: Just show a welcome

```jsx
// src/app/page.js

export default function DefaultPage() {
  return (
    <div>
      <h1>Welcome!</h1>
      <p>Click on Home or Login above.</p>
    </div>
  );
}
```

### Option B: Redirect to `/home`

```jsx
// src/app/page.js

import { redirect } from 'next/navigation';

export default function Page() {
  redirect('/home');
}
```

---

## 🧪 Test It Out

Start your dev server:

```bash
npm run dev
```

Then visit:

* `/` → shows default or redirects to `/home`
* `/home` → shows Home page
* `/login` → shows Login page
* Nav links work with full page reloads (using `<a>`)

---

## 🚀 Bonus: Use Next.js `<Link>` for Client Navigation

You can also use the `<Link>` component from Next.js for **faster, client-side navigation**:

### Update `layout.js`:

```jsx
// src/app/layout.js
import Link from 'next/link';

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body>
        <nav>
          <Link href="/home">Home</Link> | <Link href="/login">Login</Link>
        </nav>
        <hr />
        {children}
      </body>
    </html>
  );
}
```

> ✅ Now, links won't cause a full page reload.

---

## ✅ Summary

| Path      | File                | Purpose             |
| --------- | ------------------- | ------------------- |
| `/`       | `app/page.js`       | Welcome or redirect |
| `/home`   | `app/home/page.js`  | Home page           |
| `/login`  | `app/login/page.js` | Login page          |
| All pages | `app/layout.js`     | Shared layout/nav   |

---

