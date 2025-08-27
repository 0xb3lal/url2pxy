# url2pxy

**url2pxy** is a lightweight tool written in Go that forwards a list of URLs through a specified proxy (e.g., Burp Suite or ZAP).  
It is useful for quickly replaying a set of URLs into your proxy for testing and inspection.

---

## Features
- Simple and lightweight.
- Sends all URLs from a list through a specified proxy.
- Supports multi-threading for faster processing (`-t`).
- Skips empty lines in the URL list.
- Ignores HTTPS certificate verification (`InsecureSkipVerify: true`) to avoid issues with self-signed certs (useful when testing with Burp/ZAP).
- Forwards requests as-is without modifying them.
- Colored output for success, error, and failures.

---

## Installation

You need [Go](https://go.dev/dl/) installed.

```bash
go install -v github.com/0xb3lal/url2pxy/cmd/url2pxy@latest
```

This will install `url2pxy` into your `$GOPATH/bin`.  
Make sure it’s in your `$PATH`:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

**(Optional) Run from anywhere with a symlink:**
```bash
sudo ln -s $(go env GOPATH)/bin/url2pxy /usr/local/bin/url2pxy
```

**Build from source (alternative):**
```bash
git clone https://github.com/0xb3lal/url2pxy.git
cd url2pxy/cmd/url2pxy
go build -o url2pxy
```

---

## Usage

```bash
url2pxy -p [Proxy] -l [urls_list] -t [Threads]
```

### Options
- `-p` : Proxy address (e.g., `127.0.0.1:8080` for Burp Suite).  
  *(use `http://127.0.0.1:8080` if scheme is required by your proxy setup)*
- `-l` : Path to the list of URLs.
- `-t` : Number of threads to use (default: 10).

---

## Examples

Send requests through Burp Suite (listening on `127.0.0.1:8080`):

```bash
url2pxy -p 127.0.0.1:8080 -l urls.txt -t 20
```

You should now see the requests in **Burp Proxy → HTTP history**.

---

## License
MIT License © 2025 [0xb3lal](https://github.com/0xb3lal)
