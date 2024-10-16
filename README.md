## Web crawler

To build the program run:

```go build -o crawler```

To execute the crawler run next command:

```./crawler <baseURL> <maxConcurrency> <maxPages>```

Where:
- **baseUrl** is the URL from which script should start crawling links.
- **maxConcurrency** is the quantity of goroutines to run at the same time.
- **maxPages** the max quantity of links to crawl.

As an output the script will return:
- A formatted report about how much links it crawled started in **baseUrl**

Example:
```
=============================
  REPORT for https://crawler-test.com/
=============================
Found 1 internal links to crawler-test.com//mobile/separate_mobile_with_mobile_not_subdomain
Found 2 internal links to crawler-test.com//mobile/no_mobile_with_amp
Found 10 internal links to crawler-test.com//mobile/responsive
...
```

P.S. with some modification this script can be a solid base to recursively scrap any HTML you want, as all needed modules are already developed.