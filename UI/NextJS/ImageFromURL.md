To configure external image loading in **Next.js** using `next.config.js`, you can use the `images.domains` setting.

This tells Next.js to allow images to be loaded from external sources (i.e., another domain) when using the `next/image` component.

---

### ✅ Example: Load Images from an External Website

Let's say you're loading images from `https://example.com`. Your `next.config.js` should look like this:

```js
// next.config.js

/** @type {import('next').NextConfig} */
const nextConfig = {
  images: {
    domains: ['example.com'],
  },
};

module.exports = nextConfig;
```

---

### ⚠️ Important Notes

* **Do not include `https://` or `www.`** — just the domain name.
* This setting is **only needed for `next/image`**. If you're using a regular `<img>` tag, it doesn't apply.
* For subdomains, include the full subdomain: `cdn.example.com` if that’s where the image comes from.

---

### 🧪 Example Usage in a Component

```jsx
import Image from 'next/image';

export default function MyComponent() {
  return (
    <Image
      src="https://example.com/images/picture.jpg"
      alt="Example image"
      width={800}
      height={600}
    />
  );
}
```

---

If you’re using multiple domains or a CDN, let me know — I can help you tailor the config.
