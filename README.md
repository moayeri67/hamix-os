# 🐧 Hamix OS
**A Training Operating System Implemented in Go**

Hamix OS is a **learning-first, UNIX-inspired operating system simulator** written in Go.  
It runs entirely in user space and is designed to teach **core operating system concepts through real, working code** rather than theory alone.

This project is ideal for:
- Engineers refreshing OS fundamentals
- Developers learning Go through systems programming
- Students exploring kernel architecture, scheduling, memory, and filesystems

---

## ✨ Key Features

- 🧠 **Kernel Architecture**
    - Modular kernel core
    - Explicit subsystem boundaries
    - Deterministic control flow

- ⚙️ **Process Model**
    - Goroutine-based process simulation
    - PID management
    - Process lifecycle tracking

- 🕒 **Scheduler (Planned)**
    - Round-robin scheduling
    - Priority-based scheduling
    - Cooperative and simulated preemption

- 💾 **Memory Management (Planned)**
    - Virtual address spaces
    - Page tables
    - Simulated page faults
    - Copy-on-write (advanced)

- 📁 **Virtual File System (Planned)**
    - Inodes and directories
    - File descriptors
    - Device files (`/dev/console`, `/dev/null`)
    - Mount points

- 🔗 **Inter-Process Communication (Planned)**
    - Pipes
    - Message queues
    - Signals
    - Shared memory

- 🖥️ **Interactive Shell**
    - First userland process
    - OS-style command interface
    - Live system interaction

---

## 🏗️ Architecture Overview

Hamix OS follows a **clean, layered kernel design** inspired by real UNIX systems:

```
+------------------------+
|        Shell           |  ← User Space
+------------------------+
|     Syscall API        |
+------------------------+
|      Kernel Core       |
|  -----------------     |
|  Scheduler             |
|  Process Manager       |
|  Memory Manager        |
|  VFS                   |
|  IPC Manager           |
+------------------------+
|   Device Drivers       |
|  (Console, Timer)      |
+------------------------+
|     Go Runtime         |
|     Host OS            |
+------------------------+
```

All access to kernel resources happens through a **system call interface**.  
Subsystems are isolated and coordinated only by the kernel core.

---

## 📂 Project Structure

```
hamix-os/
├── cmd/
│   └── hamix/          # Bootloader (main entry point)
│       └── main.go
├── internal/
│   ├── kernel/         # Kernel core and syscall dispatcher
│   ├── process/        # Process model and lifecycle
│   ├── scheduler/      # Scheduling algorithms
│   ├── memory/         # Virtual memory and paging
│   ├── vfs/            # Virtual filesystem
│   ├── ipc/            # Inter-process communication
│   ├── drivers/        # Device drivers
│   ├── shell/          # Userland shell
│   └── platform/       # Logging, config, utilities
├── docs/               # Architecture and design docs
├── test/               # Integration and E2E tests
├── Makefile
└── go.mod
```

---

## 🚀 Getting Started

### Requirements
- Go 1.22+
- macOS / Linux / Windows

### Clone & Run

```bash
git clone https://github.com/moayeri67/hamix-os.git
cd hamix-os
go run cmd/hamix/main.go
```

You should see:

```
[BOOT] Hamix OS v0.1 is booting...
[KERNEL] Initializing kernel core
[KERNEL] Registered process: shell
[KERNEL] Starting scheduler
[PROCESS 1] Starting process
hamix>
```

---

## 🧪 Example Commands

Currently supported:
```
hamix> hello
You typed: hello
```

Planned:
- `ps` — list processes
- `kill` — terminate process
- `meminfo` — memory usage
- `ls` — filesystem browsing

---

## 🛣️ Learning Roadmap

### Phase 1 — Boot & Process Model ✅
- Kernel core
- Process creation
- Shell

### Phase 2 — System Calls
- Syscall dispatcher
- Kernel request handling
- Command execution model

### Phase 3 — Scheduling
- Round-robin scheduler
- Priority queues
- Simulated preemption

### Phase 4 — Memory
- Virtual memory model
- Page tables
- Page faults

### Phase 5 — Filesystem
- Inodes
- File descriptors
- Device files

### Phase 6 — IPC
- Pipes
- Signals
- Shared memory

---

## 🎯 Project Goals

This is **not** a production OS.  
It is a **teaching kernel** designed to:
- Make OS concepts visible
- Keep architecture explicit
- Favor clarity over performance
- Encourage experimentation

---

## 🤝 Contributing

Contributions are welcome:
- Bug fixes
- New subsystems
- Documentation improvements
- Teaching examples

Feel free to open an issue or pull request.

---

## 📜 License
MIT License — use it, break it, learn from it, improve it.

---

## 👤 Author
**Hamed Moayeri**  
Built with ❤️, ☕, and too many kernel metaphors.
