`wget spider` usually refers to **`wget --spider`**, a mode where `wget` **checks URLs without downloading files**.

### What `--spider` does

* Sends HTTP requests to test if a URL exists and is reachable
* Does **not** save any files
* Useful for link checking, monitoring, or validating mirrors

### Basic example

```bash
wget --spider https://example.com
```

If the URL is reachable, you’ll see something like:

```
HTTP request sent, awaiting response... 200 OK
```

### Check multiple URLs (from a file)

```bash
wget --spider -i urls.txt
```

### Recursive link checking (like a crawler)

```bash
wget --spider --recursive --no-verbose https://example.com
```

### Common useful options

* `-q` → quiet mode (no output)
* `-S` → show server response headers
* `--timeout=10` → set timeout
* `--tries=1` → don’t retry
* `--recursive` → follow links
* `--level=1` → limit recursion depth

Example with exit code checking:

```bash
wget --spider -q https://example.com || echo "Site is down"
```

### Typical use cases

* Check if a website or file is available
* Validate download links in scripts
* Monitor uptime in cron jobs
* Test mirrors before downloading large files

If you want, I can show:

* how to use it in a **bash script**
* how it compares to `curl -I`
* or how to spider a whole site safely
