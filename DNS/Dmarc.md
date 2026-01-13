---

## What DMARC is (1-line recap)

DMARC is a DNS **TXT record** that tells receiving mail servers how to handle emails that fail SPF/DKIM checks.

---

## Example Domain

We’ll use this example domain:

```
example.com
```

DMARC records are **always published on a subdomain**:

```
_dmarc.example.com
```

---

## Sample DMARC TXT Record (Basic)

This is a **safe starter DMARC policy** (monitoring only):

```
v=DMARC1; p=none; rua=mailto:dmarc@example.com; ruf=mailto:dmarc@example.com; fo=1
```

---

## How It Looks in NS1 (DNS UI)

When adding the record in **NS1**, you typically fill out fields like this:

### 1. Record Type

```
TXT
```

### 2. Domain / Host / Name

```
_dmarc.example.com
```

*(Some NS1 UIs only require `_dmarc` if the zone is already `example.com`)*

### 3. TTL

```
3600
```

(or “Auto” is fine)

### 4. TXT Answer / Value

```
v=DMARC1; p=none; rua=mailto:dmarc@example.com; ruf=mailto:dmarc@example.com; fo=1
```

✅ **Do not add quotes manually** unless NS1 explicitly requires them (NS1 usually handles this for you).

---

## Stricter DMARC Policy Examples

### Quarantine failing emails

```
v=DMARC1; p=quarantine; pct=100; rua=mailto:dmarc@example.com
```

### Reject failing emails (strongest)

```
v=DMARC1; p=reject; pct=100; rua=mailto:dmarc@example.com
```

---

## What Each Tag Means (Quick Reference)

| Tag        | Meaning                                 |
| ---------- | --------------------------------------- |
| `v=DMARC1` | DMARC version (required)                |
| `p=`       | Policy (`none`, `quarantine`, `reject`) |
| `rua=`     | Aggregate report email                  |
| `ruf=`     | Forensic report email                   |
| `pct=`     | % of messages policy applies to         |
| `fo=1`     | Failure reporting option                |

---

## How to Verify After Adding

After publishing, you can check with:

```
dig TXT _dmarc.example.com
```

or an online DMARC checker.

---

If you want, tell me:

* your **real domain**
* whether you use **Google Workspace, Microsoft 365, or another mail provider**

…and I’ll generate a **production-ready DMARC record tailored to your setup**.
