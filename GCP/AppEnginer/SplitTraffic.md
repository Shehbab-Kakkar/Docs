# 🚀 How to Split Traffic Between App Engine Versions Using `gcloud`

Traffic splitting in Google App Engine (GAE) lets you distribute incoming requests between multiple deployed versions of your service. This is ideal for **canary deployments**, **A/B testing**, and **safe rollbacks**.

---

## 📋 Prerequisites

Before you start, make sure:

- You have deployed **at least two versions** of your app:
    ```bash
    gcloud app deploy --version=v2
    gcloud app deploy --version=v3
    ```
- The `gcloud` CLI is **installed and authenticated**.
- You have selected the **right project**:
    ```bash
    gcloud config set project YOUR_PROJECT_ID
    ```

---

## 🧪 Split Traffic 50/50 Between v2 and v3

### Command

```bash
gcloud app services set-traffic default \
  --splits v3=0.5,v2=0.5 \
  --split-by=random
```

- **default** is the service name (change it if your service is named differently).
- **v3** and **v2** are version names.
- **0.5** means 50% of traffic to each version.
- `--split-by=random` randomly routes each request.

---

## 📝 Explanation

- **default**: Your App Engine service name. Replace with your own (e.g., `api`, `web`) if needed.
- **--splits**: Assigns a proportion (between `0` and `1`) of total traffic to each version. The values must sum to `1`.
- **--split-by**:
    - `random`: Each request is randomly routed.
    - `ip`: Each unique IP always gets the same version.
    - `cookie`: Sticky sessions based on cookies.

---

## 🔄 Modifying or Reverting the Split

To send **100% of traffic back to v2**:

```bash
gcloud app services set-traffic default \
  --splits v2=1
```

---

## ✅ Why Use Traffic Splitting?

- **Canary Releases**: Gradually roll out new features to a subset of users.
- **Performance Testing**: Compare live performance between versions.
- **Safe Rollback**: Instantly shift traffic if a deployment goes wrong.

---

## 🔍 Tips

- **Monitor logs and errors** in Google Cloud Console during and after a split.
- Use **Cloud Monitoring** to track metrics for each version.
- Combine with **feature flags** for even finer control.

---

## 📚 Final Thoughts

Traffic splitting in App Engine is a powerful strategy for safer, more controlled deployments. Use it for testing, staged rollouts, and rapid rollbacks without redeploying your application.

**Master this command to increase your deployment confidence and flexibility!**

---

> **References**  
> - [Google App Engine Documentation](https://cloud.google.com/appengine/docs)
> - [gcloud app services set-traffic](https://cloud.google.com/sdk/gcloud/reference/app/services/set-traffic)
