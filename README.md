# ctx

A project context switcher for terminal nerds. One command to drop into your project — right directory, right environment, right tmux layout.

```bash
ctx switch myproject
```

---

## How it works

`ctx` reads a TOML config file for each project, outputs shell commands for your shell to eval, and creates (or attaches to) a tmux session with your defined pane layout.

Because `ctx` runs as a child process, your shell needs a small wrapper in `.zshrc` to apply the directory and env changes:

```bash
function ctx() {
    eval "$(command ctx "$@")"
}
```

---

## Installation

```bash
git clone https://github.com/vladstefanc/ctx.git
cd ctx
go build -o ctx .
sudo mv ctx /usr/local/bin/ctx
```

Then add the shell wrapper to your `~/.zshrc`:

```bash
function ctx() {
    eval "$(command ctx "$@")"
}
```

Reload your shell:

```bash
source ~/.zshrc
```

---

## Commands

| Command             | Description                 |
| ------------------- | --------------------------- |
| `ctx switch <name>` | Switch to a context         |
| `ctx list`          | List all available contexts |

---

## Config reference

Contexts live in `~/.config/ctx/contexts/<name>.toml`.

```toml
[context]
name = "myproject"
root = "~/Projects/myproject"

[env]
GO_ENV = "development"
DB_URL = "postgres://localhost/myproject"

[[panes]]
path = "~/Projects/myproject"
command = "nvim ."

[[panes]]
path = "~/Projects/myproject/cmd"
command = ""
split = "h"

[[panes]]
path = "~/Projects/myproject"
command = ""
split = "v"
```

### Fields

**`[context]`**

- `name` — context name, must match the filename
- `root` — directory to cd into on switch

**`[env]`**

- Any key/value pairs — exported as environment variables on switch

**`[[panes]]`**

- `path` — starting directory for this pane
- `command` — command to run on launch, leave empty for a plain shell
- `split` — `h` for horizontal (side by side), `v` for vertical (stacked). Defaults to `h`

---

## Demo

![ctx demo](demo.gif)

## Contributing

PRs and issues welcome. Keep it focused — `ctx` is intentionally minimal.
