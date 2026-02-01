Ah, I see—you want to **log out from the GitHub Container Registry (GHCR)** using Helm. The correct command is:

```bash
helm registry logout ghcr.io
```

A couple of things to note:

* `ghcr.io` is the host for GitHub Container Registry.
* `helm logout registry ghcr` as you wrote it won’t work; Helm expects the full registry host.

After this, you can verify by trying to `helm pull` or `helm login`—it should prompt you again for credentials.


