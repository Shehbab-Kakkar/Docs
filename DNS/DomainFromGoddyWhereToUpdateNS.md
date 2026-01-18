You update the **nameservers at GoDaddy**, not in Google Cloud.

Here’s how the setup works and what to do, step by step:

---

## Where to update nameservers (important)

* **Nameservers are always managed at the domain registrar**
* Since your domain was bought from **GoDaddy**, you must update them **in your GoDaddy account**

Google Cloud (GCP) only provides the DNS nameservers — it does not control your domain registration.

---

## Typical setup with Google Cloud DNS

### Step 1: Create a DNS zone in Google Cloud

1. Go to **Google Cloud Console**
2. Navigate to **Network Services → Cloud DNS**
3. Create a **Public Zone**
4. Enter your domain name (example: `example.com`)

After creation, Google will show **4 nameservers**, for example:

```
ns-cloud-a1.googledomains.com
ns-cloud-a2.googledomains.com
ns-cloud-a3.googledomains.com
ns-cloud-a4.googledomains.com
```

---

### Step 2: Update nameservers in GoDaddy

1. Log in to **GoDaddy**
2. Go to **My Products → Domains**
3. Click your domain
4. Find **Nameservers**
5. Choose **Change → Custom**
6. Replace GoDaddy’s nameservers with the **Google Cloud DNS nameservers**
7. Save

✅ This is the **only place** you change nameservers.

---

## Alternative: If you are NOT using Cloud DNS

If instead you are:

* Hosting on **Cloud Run / App Engine / GCE**
* And only pointing the domain to an IP or endpoint

Then:

* **Do NOT change nameservers**
* Just add **DNS records (A / CNAME)** in GoDaddy’s DNS manager

---

## How to know which approach you need

Use **Google Cloud DNS** → Change nameservers at GoDaddy
Use **GoDaddy DNS only** → Add records in GoDaddy

---

## Propagation time

* Nameserver changes: **up to 24–48 hours**
* DNS record changes: usually **minutes to a few hours**

---

If you tell me:

* Which GCP service you’re using (Cloud Run, GCE, App Engine, Load Balancer)
* Whether you already created a Cloud DNS zone

I can give you **exact DNS records** you need.
