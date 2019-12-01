# Single Command

Single command is a way to ensure two instances of the same command won't run
at the same time.

## Installation

### Unix users (Linux, BSDs and MacOSX)

Unix users may download and install latest *single* release with command:

```bash
sh -c "$(curl http://sweetohm.net/dist/single/install)"
```

If *curl* is not installed on you system, you might run:

```bash
sh -c "$(wget -O - http://sweetohm.net/dist/single/install)"
```

**Note:** Some directories are protected, even as *root*, on **MacOSX** (since *El Capitan* release), thus you can't install *single* in */usr/bin* for instance.

### Binary package

Otherwise, you can download latest binary archive at <https://github.com/c4s4/single/releases>. Unzip the archive, put the binary of your platform somewhere in your *PATH* and rename it *single*.

## Usage

To ensure that command *build args* only runs once at a time, you would type:

```bash
single 12345 build args
```

Where:

- *12345* is a port number that should be the same for given command. Must be
  greater than 1024 if not running as root.
- *build args* is the command to run with arguments.

This command will:

- Open a server socket on given port *12345*. So that if another single command
  is already listening this port, this will fail.
- Run given command.
- Release the port when done.

*Enjoy!*
