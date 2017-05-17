# passt - The **pas**sword **st**ore
passt - The **pas**sword **st**ore - is a password manager following the Unix philosophy.
In addition it aims to be compatible to
[pass](www.passwordstore.org), i.e. its password file storage tree **but** replace
- bash and
- GnuPG.

passt consist of the following components,
- passt - the password manager itself
- passgen - a tool to generate passwords. *Working title*
- totp - a wrapper for passt to be used for Time-based One-Time Password. *Working title*

Example usage:

```sh
passgen | passt add gmail
totp github
```

Example file tree:

```
.
|-- .pgpkey
|-- gmail
`-- totp
    `-- github
```
