# Wexlgit

Wexlgit is a lightweight, Git-like version control system (VCS) implemented in Go. It provides basic version control functionality, such as initializing a repository, staging files, and creating commits. Designed for learning and experimentation, Wexlgit demonstrates core concepts of version control systems in a simple and modular codebase.

---

## Features

- **Repository Initialization**: Initialize a new repository with `wgit init`.
- **File Staging**: Stage files for the next commit with `wgit add`.
- **Commit Creation**: Create snapshots of the repository with `wgit commit`.
- **Tag Creation**: Create tags for specific commits with `wgit tag`.
- **Lightweight and Fast**: Built in Go for performance and simplicity.
- **Modular Design**: Clean separation of concerns for handlers, utilities, and object management.

---

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/Wexlersolk/wexlgit.git
   ```
2. Navigate to the project directory:
   ```bash
   cd wexlgit
   ```
3. Build the application:
   ```bash
   go build -o wgit
   ```
4. Add the binary to your PATH (optional):
   ```bash
   sudo mv wgit /usr/local/bin/
   ```

---

## Usage

### Initialize a Repository
To initialize a new Wexlgit repository, run:
```bash
wgit init
```
This creates a `.wgit` directory with the necessary structure for version control.

### Stage Files
To stage files for the next commit, use:
```bash
wgit add <file1> <file2> ...
```
You can also stage entire directories:
```bash
wgit add .
```

### Create a Commit
To create a new commit with a message, run:
```bash
wgit commit -m "Your commit message"
```
This creates a snapshot of the staged files and updates the repository history.

### Create a Tag
To create a new tag for a specific commit, use:
```bash
wgit tag -m "Tag message" <tag-name>
```
This assigns a tag to the latest commit or a specified commit hash.

---
