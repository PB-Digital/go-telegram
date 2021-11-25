# go-telegram #

[![Test Status](https://github.com/OpenSourcePB/go-telegram/workflows/tests/badge.svg)](https://github.com/OpenSourcePB/go-telegram/actions?query=workflow%3Atests)
[![Test Coverage](https://codecov.io/gh/OpenSourcePB/go-telegram/branch/master/graph/badge.svg?token=MMlSlIbqOT)](https://codecov.io/gh/OpenSourcePB/go-telegram)

go-telegram is a Go client library for accessing the [Telegram API][].

## Installation ##

go-telegram is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/OpenSourcePB/go-telegram
```

will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/OpenSourcePB/go-telegram"
```

## Usage ##

```go
import "github.com/OpenSourcePB/go-telegram"
```

Construct a new Telegram client, then use the various services on the client to
access different parts of the Telegram API. 

For example:

```go
client := telegram.NewClient("TELEGRAM_BOT_TOKEN")

// send plain message to Telegram channel or group
client.SendMessage(telegram.Message{
    ChatId:    "TELEGRAM_GROUP_OR_CHANNEL_ID",
    Body:      "Hello there!",
})
```

For more sample code snippets, head over to the
[example](https://github.com/OpenSourcePB/go-telegram/tree/master/example) directory.

## License ##

This library is licensed under the MIT License. See
[LICENSE](https://github.com/OpenSourcePB/go-telegram/blob/master/LICENSE) for the full
license text.

[Telegram API]: https://core.telegram.org/api
