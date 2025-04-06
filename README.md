
# kmail

Simple command-line tool for sending emails with file attachments.

I called it `kmail` because I'm Končal :)

## Features

- Send emails from the command line
- Attach any files
- Supports custom email subject and body
- Optional contact aliases via environment variable
- Works with Gmail SMTP (but can be adapted)

## Usage

```bash
kmail --to=m@m.sk --subject="Hello" --body="Check this out" ./file1.txt ./file2.log
```

This will send an email to `m@m.sk` with `file1.txt` and `file2.log` attached.

### Using contact aliases

You can define your contacts in an environment variable `CONTACTS` like this:

```bash
export CONTACTS="matej m@m.sk;boss boss@company.com"
```

Then you can send an email like:

```bash
kmail --to=matej ./log.txt
```

It will resolve `matej` to `m@m.sk`.

## Required environment variables

| Variable         | Description                 |
|-----------------|-----------------------------|
| `GMAIL_USERNAME` | Your Gmail address         |
| `GMAIL_PASSWORD` | Your Gmail app password    |
| `CONTACTS`       | (Optional) Contact aliases |

> Note: For Gmail, you need to create an "App Password" — regular passwords won't work with SMTP.

## Install

```bash
go install github.com/matejkoncal/kmail@latest
```

This will install `kmail` binary into your `$GOPATH/bin`  
(or `$HOME/go/bin` if you're using default Go setup).

Make sure your `$PATH` contains it:

```bash
export PATH=$PATH:$HOME/go/bin
```

## Example

```bash
kmail --to=boss --subject="Report" --body="See attachment" ./report.pdf
```

---

Made with love (and laziness) by Končal.
