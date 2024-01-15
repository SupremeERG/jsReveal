# jsReveal

Reveals interesting JS code to ethical hackers

# install

```
go install github.com/SupremeERG/jsReveal@latest
```
Or build through source code
```
git clone https://github.com/SupremeERG/jsReveal.git && \
cd jsReveal && \
go install
```

# Usage
`jsReveal -f <input js file>`
```
-f              -- Path to target JS file
-l              -- Path to a file with JS URLs
-u              -- URL to a singular JS file
-v              -- Enable Verbosity
--endpoint  -- Use predefined regex file for API endpoints and directories
--api-key       -- Use predefined regex file for API keys
-o                      -- Send output to file (JSON)
```
