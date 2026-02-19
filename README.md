# Go Task CLI 🚀

A lightweight, high-performance Command Line Interface (CLI) tool built in Go for managing daily tasks. This project demonstrates clean backend architecture, JSON-based data persistence, and a fully automated CI/CD pipeline.



## ✨ Features
- **Task Management:** Create, list, complete, and delete tasks instantly.
- **Smart Search:** Built-in case-insensitive search to find specific tasks.
- **Data Persistence:** Tasks are stored in a local `tasks.json` file—no database setup required.
- **Zero Dependencies:** Built entirely with the Go Standard Library for maximum portability and speed.
- **Automated CI/CD:** Integrated with GitHub Actions to provide pre-compiled binaries for every update.

## 🛠️ Technical Highlights
- **Language:** Go 1.22+
- **Architecture:** Decoupled logic (internal package) and interface (cmd package).
- **Serialization:** JSON encoding/decoding for local state management.
- **Automation:** GitHub Actions workflow for automated builds and artifact generation.

## 🚀 Getting Started

### Installation
You can download the latest executable from the **Actions** tab on GitHub, or build it locally:

# Clone the repository
git clone [https://github.com/viccy2/go-task-cli.git](https://github.com/viccy2/go-task-cli.git)

# Navigate to the project
cd go-task-cli

# Build the executable
go build -o task ./cmd/task
