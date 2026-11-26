<p align="right"><strong>English</strong> | <a href="README_CN.md">简体中文</a></p>

# SpearX - Modern Security & Developer Tool Manager

<div align="center">

<img src="https://img.shields.io/badge/SpearX-Tool%20Manager-6366f1?style=for-the-badge&logo=data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjQiIGhlaWdodD0iMjQiIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHBhdGggZD0iTTEyIDJMMTMuMDkgOC4yNkwyMCA5TDEzLjA5IDE1Ljc0TDEyIDIyTDEwLjkxIDE1Ljc0TDQgOUwxMC45MSA4LjI2TDEyIDJaIiBmaWxsPSJ3aGl0ZSIvPgo8L3N2Zz4K" alt="SpearX Logo" />

[![Wails](https://img.shields.io/badge/Wails-v2.10.1-FF6B6B?style=flat&logo=data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjQiIGhlaWdodD0iMjQiIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0ibm9uZSI+CjxwYXRoIGQ9Ik0xMiAyTDEzLjA5IDguMjZMMjAgOUwxMy4wOSAxNS43NEwxMiAyMkwxMC45MSAxNS43NEw0IDlMMTAuOTEgOC4yNkwxMiAyWiIgZmlsbD0iI0ZGNkI2QiIvPgo8L3N2Zz4K)](https://wails.io/)
[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Vue 3](https://img.shields.io/badge/Vue.js-3.0+-4FC08D?style=flat&logo=vue.js)](https://vuejs.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg?style=flat)](LICENSE)

**A modern desktop tool manager built for developers and security researchers**
**Supports multiple tool types, smart categorization, directory scanning, and one-click execution**

**Created by Spe4r** ⚡

[📖 Features](#-features) • [🚀 Quick Start](#-quick-start) • [📚 User Guide](#-user-guide) • [🔨 Development](#-development--build) • [🤝 Contributing](#-contributing)

</div>

---

## 📋 Table of Contents

- [✨ Features](#-features)
- [🎯 Key Advantages](#-key-advantages)
- [📸 Screenshots](#-screenshots)
- [🚀 Quick Start](#-quick-start)
- [📚 User Guide](#-user-guide)
- [⚙️ Configuration](#️-configuration)
- [🔨 Development & Build](#-development--build)
- [🏗️ Architecture](#️-architecture)
- [❓ FAQ](#-faq)
- [🤝 Contributing](#-contributing)

## ✨ Features

### 🎯 Core Capabilities
- **🔧 Multi-Type Tool Execution**
  - ☕ **Java Applications** (Java 8 / 11 / 17) — Execute `.jar` packages with per-tool Java versions and custom JVM flags.
  - 🖥️ **Terminal Tools** (`openterm`) — Open your terminal directly in the tool's working directory.
  - 🌐 **Web Apps & Bookmarks** (`Browser`) — Open remote URLs or local HTML reports in the default browser.
  - 📱 **Native Apps & Files** (`Open`) — Launch `.app` bundles, documents, and folders via OS associations.
  - 🔗 **Custom Commands** (`Custom`) — Run arbitrary shell/Docker/Python commands with variable substitution.

- **📁 Smart Organization**
  - 📂 Category-based management
  - 🏷️ Tagging system
  - 🔍 Instant search across names, tags, and descriptions
  - 📝 Per-tool Markdown notes
  - 🔄 Smart directory scanning to auto-import tools

- **🎨 Modern UI**
  - 🌟 Frosted glass / glassmorphism design
  - 🎭 Smooth animations & responsive layout
  - 🌙 Clean visual hierarchy powered by Vue 3 + Element Plus

### 🎯 Key Advantages

| Feature | Description | Benefit |
| --- | --- | --- |
| **Flexible Paths** | Relative paths, absolute paths, and URLs | Organize tools however you prefer |
| **Smart Scanning** | Auto-discovers tools in folders | Rapidly import existing toolkits |
| **Multi-Java Support** | Supports Java 8 / 11 / 17 | Handles legacy and modern security JARs |
| **Real-Time Search** | Search by name, tag, or description | Locate tools in milliseconds |
| **YAML-Backed** | Human-readable `tool.yml` config | Easy backup and visual editing |

## 📸 Screenshots

### Main Interface
> Modern tool management workspace with category filtering and instant search

![SpearX Main UI](https://s2.loli.net/2025/08/15/NSps9d6oXIRqzxM.png)

## 🚀 Quick Start

### 📋 Requirements

| Platform | Minimum Version |
| --- | --- |
| **macOS** | 10.15 Catalina |

### 📦 Installation

#### Option 1: Download Prebuilt Binary (Recommended)
1. Go to the [Releases](https://github.com/sspsec/Spear/releases) page.
2. Download the latest archive for your platform.
3. Extract and launch `SpearX`.

#### Option 2: Build from Source
```bash
# 1. Clone the repository
git clone https://github.com/sspsec/Spear.git
cd Spear/Spear\ X

# 2. Install Wails CLI (if not installed)
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# 3. Build the app
wails build -platform darwin/amd64 -clean

# 4. Run
./build/bin/SpearX
```

### ⚡ Quick Walkthrough

1. **First Launch**: Creates a default configuration automatically.
2. **Add Tools**: Click ➕ to add a tool manually.
3. **Scan Folders**: Click 🔄 to scan an existing tools directory.
4. **Launch**: Click any tool card to run it.

## 📚 User Guide

### 📖 Basic Operations

#### 1️⃣ Adding Tools
- ➕ **Manual Add** — Configure tool metadata and launch parameters one by one.
- 🔄 **Scan Default Directory** — Scan the built-in `resources` directory.
- 📂 **Scan Custom Directory** — Pick any folder on disk to import tools.

#### 2️⃣ Execution Modes

##### ☕ Java Applications (Java8 / Java11 / Java17)

```yaml
- ToolName: Burp Suite Professional
  PATH: /Applications/Security/BurpSuite
  FileName: burpsuite_pro.jar
  VALUE: Java11
  Optional: "-Xmx4g -XX:+UseG1GC"
```

##### 🖥️ Terminal Tools (`openterm`)

```yaml
- ToolName: Nmap Scanner
  PATH: /usr/local/bin/nmap
  FileName: ""
  VALUE: openterm
  Optional: ""
```

##### 🌐 Web Apps (`Browser`)

```yaml
- ToolName: GitHub
  PATH: https://github.com
  VALUE: Browser

- ToolName: Local Report Viewer
  PATH: /Users/tools/reports/index.html
  VALUE: Browser
```

##### 📱 System Default (`Open`)

```yaml
- ToolName: Wireshark
  PATH: /Applications/Wireshark.app
  VALUE: Open

- ToolName: Tool Manual
  PATH: /Users/docs/manual.pdf
  VALUE: Open
```

##### 🔗 Custom Commands (`Custom`)

```yaml
- ToolName: Docker Tool Container
  PATH: /path/to/tool
  COMMAND: "docker run -it --rm -v {path}:/workspace tool:latest"
  VALUE: Custom
  Optional: "--network host"

- ToolName: Python Scanner
  PATH: /tools/scanner
  FileName: scanner.py
  COMMAND: "python3 {filename} --config config.json"
  VALUE: Custom
```

Supported template variables: `{path}`, `{filename}`, `{name}`, `{optional}`.

---

### 🔧 Execution Priority & Fallback

| Priority | Detection Method | Fallback |
| --- | --- | --- |
| **1** | Explicit `VALUE` | Execute as configured |
| **2** | File extension | `.jar` → Java, `.app` → Open |
| **3** | File inspection | Check header & executable permissions |
| **4** | OS default | System file association |

### 📂 Recommended Directory Structure & Smart Scanning

When scanning a directory, SpearX automatically detects executables (`.jar`, `.exe`, `.app`, `.sh`, `.py`, etc.), documentation (`README.*`, `*.md`, `*.pdf`), web files (`*.html`), configs (`*.yaml`, `*.json`), and shortcuts (`.url`, `.webloc`), and categorizes folders by keywords such as `info`/`recon`, `scan`/`vulnerability`, `exploit`/`penetration`, `web`, and `system`.

## ⚙️ Configuration

### 📄 `tool.yml`

```yaml
JavaPaths:
  Java8: "/usr/bin/java"
  Java11: "/usr/bin/java"
  Java17: "/usr/bin/java"

Categories:
- CategoryName: Recon
  Tools:
  - ToolName: WebFinder
    PATH: resources/info/webfinder
    FileName: webfinder-next.jar
    VALUE: Java8
    COMMAND: -jar
    Optional: "-Xmx2g"
```

Leave `JavaPaths` empty to auto-detect Java from `PATH`, or specify explicit JDK binary paths.

## 🔨 Development & Build

### 🛠️ Prerequisites
- **Go 1.24+**
- **Node.js 18+**
- **Wails CLI v2.10.1** (`go install github.com/wailsapp/wails/v2/cmd/wails@latest`, verify with `wails doctor`)

### 📦 Production Build

```bash
wails build                          # Current platform
wails build -platform darwin/arm64   # macOS Apple Silicon
wails build -platform darwin/amd64   # macOS Intel
wails build -platform windows/amd64  # Windows x64
wails build -platform linux/amd64    # Linux x64
```

### 🔧 Tech Stack
- **Frontend**: Vue 3 (Composition API), Element Plus, Vite, CSS3
- **Backend**: Go 1.24, Wails v2, YAML

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐
│   Frontend      │    │   Backend       │
│   (Vue 3)       │◄──►│   (Go)          │
├─────────────────┤    ├─────────────────┤
│ • UI Components │    │ • Tool Manager  │
│ • State Mgmt    │    │ • File Ops      │
│ • Events        │    │ • Runner Engine │
│ • Glassmorphism │    │ • Config Parser │
└─────────────────┘    └─────────────────┘
         ▲                       ▲
         │      Wails Bridge     │
         ▼                       ▼
┌─────────────────┐    ┌─────────────────┐
│ Desktop Runtime │    │ OS Integration  │
└─────────────────┘    └─────────────────┘
```

## ❓ FAQ

**Q: macOS says the app is damaged or cannot verify the developer on first launch.**
```bash
xattr -rd com.apple.quarantine Spear.app
# or on macOS Sequoia:
codesign --sign - Spear.app
```

**Q: Java tools fail to launch.**
Check that the configured Java version is installed and paths in Settings (or system `PATH`) are valid.

**Q: Where are tool notes saved?**
As Markdown files (`<ToolName>.md`) inside each tool's directory.

**Q: Looking for legacy Fyne V3 / V4 / V5 notes?**
<details>
<summary>Legacy Version Archive (V3 – V5)</summary>

Legacy builds (`Spear-V3`, `Spear-V4`) are still preserved in this repository. For full Chinese release notes, iTerm/dark-mode setup, and migration tips from V5 `resources`/`tool.yml`, please see [README_CN.md](README_CN.md).

</details>

## 🤝 Contributing

Contributions are welcome! Feel free to open an [Issue](https://github.com/sspsec/Spear/issues) for bug reports/feature ideas or submit a Pull Request.

## 🙏 Acknowledgments

Built with [Wails](https://wails.io/), [Vue.js](https://vuejs.org/), [Element Plus](https://element-plus.org/), [Go](https://golang.org/), and [Vite](https://vitejs.dev/).

---

<div align="center">

**If SpearX helps your workflow, please consider giving it a ⭐ Star!**

<img width="790" height="463" alt="Support QR" src="https://github.com/user-attachments/assets/6e2b5a9e-8eec-4eb7-8627-0f64b8c8b8b5" />

[![Star History Chart](https://api.star-history.com/svg?repos=sspsec/Spear&type=Date)](https://star-history.com/#sspsec/Spear&Date)

[🐛 Report Bug](https://github.com/sspsec/Spear/issues) • [💡 Feature Request](https://github.com/sspsec/Spear/issues)

**Created with ❤️ by Spe4r**

</div>
