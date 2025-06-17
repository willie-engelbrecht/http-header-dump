Application listens on port 8080, or you can override the environment variable `PORT` to another value. 

To run the application:

```
go run main.go
```

And it will print the following headers:
```
Request from GET /metrics
{
  "Accept": [
    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"
  ],
  "Accept-Encoding": [
    "gzip, deflate, br, zstd"
  ],
  "Accept-Language": [
    "en-GB,en-ZA;q=0.9,en-US;q=0.8,en;q=0.7"
  ],
  "Cache-Control": [
    "no-cache"
  ],
  "Connection": [
    "keep-alive"
  ],
  "Pragma": [
    "no-cache"
  ],
  "Sec-Ch-Ua": [
    "\"Google Chrome\";v=\"137\", \"Chromium\";v=\"137\", \"Not/A)Brand\";v=\"24\""
  ],
  "Sec-Ch-Ua-Mobile": [
    "?0"
  ],
  "Sec-Ch-Ua-Platform": [
    "\"Windows\""
  ],
  "Sec-Fetch-Dest": [
    "document"
  ],
  "Sec-Fetch-Mode": [
    "navigate"
  ],
  "Sec-Fetch-Site": [
    "none"
  ],
  "Sec-Fetch-User": [
    "?1"
  ],
  "Upgrade-Insecure-Requests": [
    "1"
  ],
  "User-Agent": [
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36"       
  ]
}
Request from GET /favicon.ico
{
  "Accept": [
    "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8"
  ],
  "Accept-Encoding": [
    "gzip, deflate, br, zstd"
  ],
  "Accept-Language": [
    "en-GB,en-ZA;q=0.9,en-US;q=0.8,en;q=0.7"
  ],
  "Cache-Control": [
    "no-cache"
  ],
  "Connection": [
    "keep-alive"
  ],
  "Pragma": [
    "no-cache"
  ],
  "Referer": [
    "http://localhost:8080/metrics"
  ],
  "Sec-Ch-Ua": [
    "\"Google Chrome\";v=\"137\", \"Chromium\";v=\"137\", \"Not/A)Brand\";v=\"24\""
  ],
  "Sec-Ch-Ua-Mobile": [
    "?0"
  ],
  "Sec-Ch-Ua-Platform": [
    "\"Windows\""
  ],
  "Sec-Fetch-Dest": [
    "image"
  ],
  "Sec-Fetch-Mode": [
    "no-cors"
  ],
  "Sec-Fetch-Site": [
    "same-origin"
  ],
  "User-Agent": [
    "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36"       
  ]
}
```
