To get a list of users from **Okta** using Python, you'll need to use the **Okta API**. Here's a basic Python script that does this using the `requests` library.

---

### ✅ Prerequisites:

1. **Okta Domain** – something like `https://your-domain.okta.com`
2. **API Token** – from Okta Admin dashboard
3. Python environment with `requests` installed:

   ```bash
   pip install requests
   ```

---

### 🐍 Python Script: Get Okta Users

```python
import requests

# === CONFIGURATION ===
OKTA_DOMAIN = "https://your-domain.okta.com"  # Replace with your actual Okta domain
API_TOKEN = "your_api_token_here"             # Replace with your API token

# Base URL for Okta API
BASE_URL = f"{OKTA_DOMAIN}/api/v1/users"

# Headers for the API call
HEADERS = {
    "Authorization": f"SSWS {API_TOKEN}",
    "Accept": "application/json",
    "Content-Type": "application/json"
}

def get_okta_users():
    users = []
    url = BASE_URL

    while url:
        response = requests.get(url, headers=HEADERS)
        if response.status_code != 200:
            raise Exception(f"Error fetching users: {response.status_code} - {response.text}")
        
        data = response.json()
        users.extend(data)

        # Pagination: look for 'next' link in headers
        links = response.links
        url = links['next']['url'] if 'next' in links else None

    return users

# === Run the script ===
if __name__ == "__main__":
    users = get_okta_users()
    for user in users:
        print(f"{user['profile']['login']} - {user['profile'].get('firstName', '')} {user['profile'].get('lastName', '')}")
```

---

### 📌 Notes:

* This handles **pagination** — Okta returns 200 users per page.
* Replace placeholders (`your-domain.okta.com` and `your_api_token_here`) with your actual values.
* You can expand the script to export to CSV or filter specific users.

Let me know if you want it to save to a file or include deactivated users.
