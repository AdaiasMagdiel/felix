<p align="center">
  <img src="assets/banner.jpg" alt="Felix the Recursive Cat Banner" width="100%">
</p>

# 🐱 Felix — A Recursive Cat

[![Go Version](https://img.shields.io/badge/Go-1.24%2B-blue.svg)](https://go.dev/)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-green.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![Build](https://img.shields.io/badge/build-passing-brightgreen.svg)](#)

**Felix** is a small but powerful command-line tool written in Go 🧰  
It recursively reads files that match wildcard patterns (like `*.md`, `**/*.txt`, etc.)  
and prints their contents — just like the Unix `cat`, but recursive and smart 🐾

> “Felix” comes from Latin, meaning _happy_ or _lucky_.

> It’s also a nod to our feline friends — perfect for a recursive `cat`. 🐾

---

## ✨ Features

- 📂 Recursive glob pattern support (`**/*.md`)
- 🧾 Multiple file patterns (`*.md`, `*.txt`, etc.)
- ⚙️ Output redirection (`-o output.txt`)
- 🚫 Automatic binary file detection and skip
- 💙 Pure Go — no external dependencies besides `doublestar`

---

## 🚀 Installation

### Prerequisites

- [Go 1.24+](https://go.dev/dl/)

### Install directly from GitHub

```bash
go install github.com/AdaiasMagdiel/felix@latest
```

---

## 🧰 Usage

```bash
felix [options] <pattern1> [<pattern2> ...]
```

### Examples:

```bash
felix **/*.md
felix *.txt *.log
felix -o combined.txt **/*.md
```

### Output:

```
===== docs/readme.md =====
# Introduction
Welcome to Felix!

===== notes/todo.txt =====
- [x] Implement recursive globbing
- [x] Detect and skip binary files
- [x] Add output redirection
```

---

## ⚙️ Build from source

```bash
git clone https://github.com/AdaiasMagdiel/felix.git
cd felix
go build ./cmd/felix
```

---

## 🧱 Project structure

```
felix/
├── assets/
│   └── banner.jpg
├── main.go
├── go.mod
├── go.sum
├── COPYRIGHT
├── LICENSE
└── README.md
```

---

## 📜 License

Licensed under the **GNU General Public License v3.0 (GPLv3)**.
See the [LICENSE](./LICENSE) file for details.
