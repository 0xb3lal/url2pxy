# url2pxy

url2pxy is a lightweight tool written in Go that forwards a list of URLs through a specified proxy (e.g., Burp Suite or ZAP).  
It is useful for quickly replaying a set of URLs into your proxy for testing and inspection.

## Features
- Simple and lightweight.
- Sends all URLs from a list through a specified proxy.
- Supports multi-threading for faster processing (`-t`).
- Skips empty lines in the URL list.
- Ignores HTTPS certificate verification (`InsecureSkipVerify: true`) to avoid issues with self-signed certs (useful when testing with Burp/ZAP).
- Forwards requests as-is without modifying them.
- Colored output for success, error, and failures.

## Installation
You need [Go](https://go.dev/dl/) installed.

```bash
go install -v github.com/0xb3lal/url2pxy/cmd/url2pxy@latest
```

This will install `url2pxy` into your `$GOPATH/bin` (make sure it’s in your `$PATH`).

## Usage
```bash
url2pxy -p [Proxy] -l [urls_list] -t [Threads]
```

### Options
- `-p` : Proxy address (e.g., `127.0.0.1:8080` for Burp Suite).
- `-l` : Path to the list of URLs.
- `-t` : Number of threads to use (default: 10).

### Example
```bash
url2pxy -p 127.0.0.1:8080 -l urls.txt -t 20
```

## License
MIT License
