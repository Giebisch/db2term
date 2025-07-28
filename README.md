# db2term

Deutsche Bahn to Terminal – A TUI for quickly viewing your favorite
German public transport connections on the command line.

Written in Go!

## 📸 Demo

Here's a preview of db2term in action:

![demo.gif](docs/demo.gif)

## ✨ Overview

db2term is a lightweight Go program that brings Deutsche Bahn (DB) public transport schedules directly to your terminal. Whether you commute daily or frequently travel via train, S-Bahn, U-Bahn, or bus, db2term helps you stay updated on your preferred routes with just a few keystrokes.

**🎯 Features**

 🚌 Support for all DB transit modes: Train, S-Bahn, U-Bahn, buses, etc

 ⭐ Favorite routes: Save your most-used trips for quick access.

 💨 Blazingly fast CLI: Optimized for performance and minimal output latency.

 💻 Terminal-first experience: Pretty TUI for quickly showing the next connections

## 🛠 Installation

Prerequisites

    Go 1.24

You will also have to enable [Nerdfonts](https://www.nerdfonts.com/) in your terminal.

## 🚀 Usage

You can use `make build` to build a binary.

Recommended usage is to alias it in your `~/.zshrc` or `~/.bash_profile`.

Right now you have to manually input the station codes in the created `<temp>/db2term/config.json`

`db2term` automatically creates a new folder in `os.UserConfigDir()`.
Check, where this in your OS.

This only needs to be done once, but I still want to add the functionality to add station codes via a search bar within the TUI, so you don't have to manually retrieve them from the DB website.

Example `config.json`:

```json

{
  "trips": [
    {
      "from": "München Hbf",
      "to": "Dachau Stadt",
      "fromCode": "A=1@O=München Hbf@X=11558339@Y=48140229@U=80@L=8000261@B=1@p=1752089579@i=U×008020347@",
      "toCode": "A=1@O=Dachau Stadt@X=11439762@Y=48266662@U=80@L=8001355@B=1@p=1750894836@i=U×008020455@"
    },
    {
      "from": "Dachau Stadt",
      "to": "München Hbf",
      "fromCode": "A=1@O=Dachau Stadt@X=11439762@Y=48266662@U=80@L=8001355@B=1@p=1750894836@i=U×008020455@",
      "toCode": "A=1@O=München Hbf@X=11558339@Y=48140229@U=80@L=8000261@B=1@p=1752089579@i=U×008020347@"
    }
  ]
}

```

## 🙌 Contributing

Pull requests are welcome! Feel free to fork the repo and create a PR.
