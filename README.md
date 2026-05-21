# fl-launches

See the next 5 rocket launches from Cape Canaveral, right in your terminal.

```
  #1  FALCON 9 BLOCK 5 | STARLINK GROUP 10-31
      Falcon 9 Block 5  ·  Space Launch Complex 40
      Thu May 21, 2026 at 6:04 AM EDT  |  T- 7h 24m

      A batch of 29 satellites for the Starlink mega-constellation -
      SpaceX's project for space-based Internet communication system.
```

## Installation

### macOS — Homebrew

```bash
brew tap patlehmann1/fl-launches
brew install fl-launches
```

### macOS — Direct download (no Homebrew)

```bash
# Apple Silicon (M1/M2/M3)
curl -L https://github.com/patlehmann1/space_coast_fl_launches/releases/latest/download/fl-launches_darwin_arm64.tar.gz | tar -xz
sudo mv fl-launches /usr/local/bin/

# Intel Mac
curl -L https://github.com/patlehmann1/space_coast_fl_launches/releases/latest/download/fl-launches_darwin_amd64.tar.gz | tar -xz
sudo mv fl-launches /usr/local/bin/
```

### Linux — Homebrew (Linuxbrew)

```bash
brew tap patlehmann1/fl-launches
brew install fl-launches
```

### Linux — Direct download (no Homebrew)

```bash
# x86_64
curl -L https://github.com/patlehmann1/space_coast_fl_launches/releases/latest/download/fl-launches_linux_amd64.tar.gz | tar -xz
sudo mv fl-launches /usr/local/bin/

# ARM64 (Raspberry Pi, etc.)
curl -L https://github.com/patlehmann1/space_coast_fl_launches/releases/latest/download/fl-launches_linux_arm64.tar.gz | tar -xz
sudo mv fl-launches /usr/local/bin/
```

## Usage

```bash
fl-launches              # show next 5 launches
fl-launches --count 3    # show next 3 launches (1–10)
fl-launches --version
```

## Data source

Launch data from [Launch Library 2](https://thespacedevs.com/llapi) by The Space Devs.
