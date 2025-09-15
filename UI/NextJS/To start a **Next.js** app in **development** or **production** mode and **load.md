To start a **Next.js** app in **development** or **production** mode and **load environment variables based on the environment**, follow this complete guide:

---

## ✅ 1. **Environment Variable Setup**

Create the following environment files in your root directory:

```
.env.local           → for local development (ignored by git)
.env.development     → for dev mode
.env.production      → for production mode
```

### Example `.env.development`

```env
NEXT_PUBLIC_API_URL=http://localhost:3000/api
```

### Example `.env.production`

```env
NEXT_PUBLIC_API_URL=https://yourdomain.com/api
```

### ⚠️ Prefix with `NEXT_PUBLIC_` for client-side usage:

Only variables starting with `NEXT_PUBLIC_` will be available in the **browser**. Others are **server-only**.

---

## ✅ 2. **Using Environment Variables in Code**

```js
const apiUrl = process.env.NEXT_PUBLIC_API_URL;
console.log('API URL:', apiUrl);
```

---

## ✅ 3. **Starting the App**

### 🔧 Development Mode

```bash
npm run dev
# or
yarn dev
# or
pnpm dev
```

This uses `.env.local` and `.env.development` (if both exist, `.env.local` takes precedence).

---

### 🚀 Production Mode

Build first, then start:

```bash
npm run build
npm run start
```

This uses `.env.production`.

If you want to override envs during runtime, set them like this:

```bash
NODE_ENV=production NEXT_PUBLIC_API_URL=https://api.example.com npm run start
```

---

## ✅ 4. **Custom Scripts (optional)**

In your `package.json`, you can define custom scripts:

```json
"scripts": {
  "dev": "next dev",
  "build": "next build",
  "start": "next start",
  "start:prod": "NODE_ENV=production next start"
}
```

---

## ✅ 5. **Environment Priority Order**

Next.js loads env vars in this order:

1. `.env.local` (local overrides)
2. `.env.development` or `.env.production` (based on `NODE_ENV`)
3. `.env`

So:

* `.env.local` > `.env.[env]` > `.env`

---

## ✅ 6. **Extra Tip: Dynamic Loading with `dotenv` (optional)**

If you want to fully control loading, install `dotenv`:

```bash
npm install dotenv
```

And in `next.config.js`:

```js
require('dotenv').config({
  path: `.env.${process.env.NODE_ENV}`
});

module.exports = {
  env: {
    CUSTOM_VAR: process.env.CUSTOM_VAR,
  },
};
```

---

Would you like a ready-made boilerplate project structure with this setup?
