# dsp

Tool for generating and processing discrete signals

## 🚀 Instalation

### Option 1 System-wide Installation (Recommended)

This method installs the binaries to `/usr/local/bin`, making them available to all users.

1. Build and install

```bash
sudo make install
```

2. Verify

```bash
dsp --help
```

### Option 2: Local Installation (No Root/Sudo)

Use this if you don't have administrator privileges or want to keep the tool restricted to your user.

1. Install to your home directory:
```bash
make install PREFIX=$HOME/.local
```

2. Update your PATH (if needed):
Ensure $HOME/.local/bin is in your shell configuration (.bashrc or .zshrc):

```bash
export PATH=$PATH:$HOME/.local/bin
```

### Option 3: Manual Build

If you only want to compile the binaries without moving them to system folders:

```bash
make build
```

The executable files will be generated in the `./bin/` directory.

### 🗑 Uninstallation
To remove the installed binaries from your system:
If installed via sudo:
```bash
sudo make uninstall
```

If installed locally:

```bash
make uninstall PREFIX=$HOME/.local
```

🛠 Cross-Compilation
The Makefile supports cross-compilation via environment variables. For example, to build for ARM64:
```bash
GOARCH=arm64 make build
```