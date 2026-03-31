[Tue Mar 31 00:43:11.398631 2026] [auth_mellon:warn] [pid 747554:tid 747603] [client 10.10.12.12:49624] User has disabled cookies, or has lost the cookie before returning from the SAML2 login server., referer: https://test.okta.com/


That log message is coming from **mod_auth_mellon**, and it indicates a problem with the SAML authentication flow—specifically, the session cookie required to complete login is missing.

### What it means

* During SAML login (e.g., via Okta), the SP (your Apache server with Mellon) sets a **session cookie** before redirecting the user to the IdP.
* When the user returns after authentication, that cookie must still be present.
* This warning means:

  * The cookie was **never stored**, or
  * It was **lost / not sent back** with the request.

---

### Common causes

1. **Cookies disabled in browser**

   * User has blocked cookies entirely.

2. **SameSite cookie issues**

   * Modern browsers block cookies in cross-site redirects unless properly configured.
   * Missing or incorrect `SameSite=None; Secure`.

3. **Domain / path mismatch**

   * Cookie set for one domain, returned on another (e.g., `test.okta.com` vs your app domain).

4. **HTTP vs HTTPS**

   * Secure cookies won’t be sent over HTTP.

5. **Load balancer / reverse proxy issues**

   * Cookie not preserved across nodes (no sticky sessions).
   * Headers like `X-Forwarded-Proto` misconfigured.

6. **Clock skew or session timeout**

   * Cookie expires before return.

---

### How to fix

#### ✅ 1. Configure SameSite correctly

In Apache config:

```apache
MellonCookieSameSite None
MellonSecureCookie On
```

#### ✅ 2. Ensure HTTPS everywhere

* Your SP endpoint must be HTTPS.
* Cookies marked `Secure` require HTTPS.

#### ✅ 3. Check cookie domain

Make sure it matches your application domain:

```apache
MellonCookieDomain your-app-domain.com
```

#### ✅ 4. Verify browser behavior

* Test in incognito with cookies enabled.
* Check DevTools → Application → Cookies.

#### ✅ 5. Load balancer setup

* Enable sticky sessions OR shared session storage.
* Forward headers properly:

  * `X-Forwarded-Proto: https`

#### ✅ 6. Increase logging for debugging

```apache
LogLevel auth_mellon:debug
```

---

### Quick diagnostic checklist

* Does the cookie appear before redirect to Okta?
* Is it present when returning?
* Is it blocked due to SameSite?
* Is the domain identical?

---

If you want, share your Apache Mellon config (sanitized), and I can pinpoint exactly what's causing it.
