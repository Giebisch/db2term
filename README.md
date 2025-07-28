# db2term

Deutsche Bahn to Terminal – A CLI tool for quickly viewing your favorite
German public transport connections.

Written in Go 🧳🚉💻

## ✨ Overview

db2term is a lightweight Go program that brings Deutsche Bahn (DB) public transport schedules directly to your terminal. Whether you commute daily or frequently travel via train, S-Bahn, U-Bahn, or bus, db2term helps you stay updated on your preferred routes with just a few keystrokes.
🎯 Features

- 🚌 Support for all DB transit modes: Train, S-Bahn, U-Bahn, and buses.
- ⭐ Favorite routes: Save your most-used trips for quick access.
- 💨 Blazingly fast CLI: Optimized for performance and minimal output latency.
- 💻 Terminal-first experience: No GUI, just clean and readable output for your shell or tmux setup.

## 📸 Demo

Here's a preview of db2term in action:

    Want to contribute a better GIF or terminal theme? PRs welcome!

## 🛠 Installation

Prerequisites

    Go 1.20 or higher

Install via go install

go install github.com/<your-username>/db2term@latest

Or clone and build manually

git clone <https://github.com/><your-username>/db2term.git
cd db2term
go build -o db2term

and Nerdfonts

## 🚀 Usage

Basic command

db2term <from> <to>

Example:

db2term berlin hbf potsdam hbf

With saved favorites

db2term fav home-office

You can configure your favorites using a simple config file (~/.db2term/config.yml):

favorites:
  home-office:
    from: "Berlin Hbf"
    to: "Potsdam Griebnitzsee"
  office-home:
    from: "Potsdam Griebnitzsee"
    to: "Berlin Hbf"

## 🧠 How It Works

db2term interacts with Deutsche Bahn’s public transport APIs to fetch live connection data. It parses and formats the results to present them clearly in your terminal, without needing to visit a website or open an app.
🧪 Example Output

🚆 ICE 1723 | Berlin Hbf → München Hbf | Dep: 14:30 | Arr: 18:45 | Platform: 8 | On Time
🚇 U7 | Mehringdamm → Rathaus Spandau | Dep: 14:36 | Arr: 15:05 | On Time

🔮 Planned Features

- Live delay tracking

- Station search

## 🙌 Contributing

Pull requests are welcome! Feel free to fork the repo and create a PR.
