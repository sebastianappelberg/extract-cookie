# Extract Cookie

`extract-cookie` helps you export the full set of cookies from your browser, even for browsers that encrypt their cookies.

**NOTE:** The cookies that are exported using this tool should be considered highly sensitive information, as they'll allow whoever has them to 
start a browser session as if they were you. This tool is a bona fide foot gun, use with caution. 

## Motivation

Writing a scraper that requires you to be logged in can be annoying, especially when 2FA is required. With Playwright, you can start 
a session based on previously stored cookies. See https://playwright.dev/docs/auth for more info. For whatever reason, it can be annoying to extract those cookies 
from your browser. This tool was created to make it less annoying. 

## Usage

```
extract-cookie -b <browser> -f <format>
```

```
Flags:
  -h, --help    help for cookie-extractor
  -b, -browser  Browser to extract cookies from: chrome, chromium, edge, firefox. (default "firefox").  
  -f, -format   Output format: json or text. (default "json"). 
```

Using `-f json` will output the cookies in the following format:

```json
{
  "cookies": [
    {
      "Name": "NameA",
      "Value": "ValueA",
      "Domain": ".domainA.com",
      "Path": "/",
      "Expires": 1777554369,
      "Secure": true,
      "HttpOnly": true
    }
  ]
}
```
And using `-f text`, this format:
```
Name: NameA
Value: ValueA
Domain: .domainA.com
Path: /
Expires: 2026-04-30 15:00:00
Secure: true
HttpOnly: true
```

## Installation

Using Go:
```
go install github.com/sebastianappelberg/extract-cookie@latest
```

## TODO

- [ ] macOS support.
- [ ] Linux support.
- [ ] Additional browsers.
- [ ] Ability to extract cookies from Chromium-based browsers on Windows while the browser is running.

## Contributing

The easiest way to contribute is to take something from the [TODO](#todo) list. Other contributions are of course welcome. Please make sure your pull request
contains a description containing the motivation for the changes.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.

