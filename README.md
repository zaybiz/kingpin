# Kingpin - A Go (golang) command line and flag parser [![Build Status](https://travis-ci.org/alecthomas/kingpin.png)](https://travis-ci.org/alecthomas/kingpin)

## Features

- POSIX-style short flag combining.
- Parsed, type-safe flags.
- Parsed, type-safe positional arguments.
- Support for required flags and required positional arguments
- Callbacks per command, flag and argument.
- Help output that isn't as ugly as sin.

## Simple Example

Kingpin can be used for simple flag+arg applications like so:

```
$ ping --help
usage: ping [<flags>] <ip> [<count>]

Flags:
  --debug            Enable debug mode.
  --help             Show help.
  -t, --timeout=5s   Timeout waiting for ping.

Args:
  <ip>        IP address to ping.
  [<count>]   Number of packets to send
$ ping 1.2.3.4 5
Would ping: 1.2.3.4 with timeout 5s%
```

From the following source:

```go
package main

import (
  "fmt"

  "github.com/alecthomas/kingpin"
)

var (
  debug   = kingpin.Flag("debug", "Enable debug mode.").Bool()
  timeout = kingpin.Flag("timeout", "Timeout waiting for ping.").Default("5s").MetaVarFromDefault().Short('t').Duration()
  ip      = kingpin.Arg("ip", "IP address to ping.").Required().IP()
  count   = kingpin.Arg("count", "Number of packets to send").Int()
)

func main() {
  kingpin.Parse()
  fmt.Printf("Would ping: %s with timeout %s", *ip, *timeout)
}
```

## Complex Example

Kingpin can also produce complex command-line applications with global flags,
subcommands, and per-subcommand flags, like this:

```
$ chat
usage: chat [<flags>] <command> [<flags>] [<args> ...]

A command-line chat application.

Flags:
  --debug              enable debug mode
  --help               Show help.
  --server=127.0.0.1   server address

Commands:
  help <command>
    Show help for a command.

  post [<flags>] <channel>
    Post a message to a channel.

  register <nick> <name>
    Register a new user.

$ chat help post
usage: chat [<flags>] post [<flags>] <channel> [<text>]

Post a message to a channel.

Flags:
  --image=IMAGE   image to post

Args:
  <channel>   channel to post to
  [<text>]    text to post
$ chat post --image=~/Downloads/owls.jpg pics
...
```

From this code:

```go
package main

import (
  "os"
  "github.com/alecthomas/kingpin"
)

var (
  app      = kingpin.New("chat", "A command-line chat application.")
  debug    = app.Flag("debug", "enable debug mode").Default("false").Bool()
  serverIP = app.Flag("server", "server address").Default("127.0.0.1").MetaVarFromDefault().IP()

  register     = app.Command("register", "Register a new user.")
  registerNick = register.Arg("nick", "nickname for user").Required().String()
  registerName = register.Arg("name", "name of user").Required().String()

  post        = app.Command("post", "Post a message to a channel.")
  postImage   = post.Flag("image", "image to post").File()
  postChannel = post.Arg("channel", "channel to post to").Required().String()
  postText    = post.Arg("text", "text to post").String()
)

func main() {
  switch kingpin.MustParse(app.Parse(os.Argv[1:])) {
  // Register user
  case "register":
    println(*registerNick)

  // Post message
  case "post":
    if *postImage != nil {
    }
    if *postText != "" {
    }
  }
}
```

## Parsers

Kingpin supports both flag and positional argument parsers for converting to
Go types. For example, some included parsers are `Int()`, `Float()`,
`Duration()` and `ExistingFile()`.

Parsers conform to Go's [`flag.Value`](http://godoc.org/flag#Value)
interface, so any existing implementations will work.

For example, a parser for accumulating HTTP header values might look like this:

```go
type HTTPHeaderValue http.Header

func (h *HTTPHeaderValue) Set(value string) error {
  parts := strings.SplitN(value, ":", 2)
  if len(parts) != 2 {
    return fmt.Errorf("expected HEADER:VALUE got '%s'", value)
  }
  (*http.Header)(h).Add(parts[0], parts[1])
  return nil
}

func (h *HTTPHeaderValue) String() string {
  return ""
}
```

As a convenience, I would recommend something like this:

```go
func HTTPHeader(s Settings) (target *http.Header) {
  target = new(http.Header)
  s.SetValue((*HTTPHeaderValue)(target))
  return
}
```

You would use it like so:

```go
headers = HTTPHeader(kingpin.Flag("--header", "Add a HTTP header to the request.").Short('-H'))
```

## Default Values

The default value is the zero value for a type. This can be overridden with
the `Default(value)` function on flags and arguments. This function accepts a
string, which is parsed by the value itself, so it *must* be compliant with
the format expected.
