clearbit-cli
===

A small wrapper around the `clearbit-go` library. Useful for quick lookups of the clearbit api right from your command line.

Requires `go v1.11`

```
go install github.com/kiernanmcgowan/clearbit-cli
```

Usage
==

Set your clearbit key to the environment variable `CLEARBIT_KEY`

```
CLEARBIT_KEY=sk_magic
```

And call clearbit's like so:

```
clearbit-cli company <domain>
clearbit-cli person <email>
clearbit-cli reveal <ip>
clearbit-cli discovery <query>
clearbit-cli autocomplete <query>
```

