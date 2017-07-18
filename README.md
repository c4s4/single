Single Command
==============

Single command is a way to ensure two instances of the same command won't run
at the same time.

Installation
------------

Drop the binary for your platform in the *bin* directory of the archive
somewhere in your `PATH` and rename it *single*.

Usage
-----

To ensure that command *build args* only runs once at a time, you would type:

    $ single 12345 build args

Where:

- *12345* is a port number that should be the same for given command. Must be
  greater than 1014 if not running as root.
- *build args* is the command to run with arguments.

This command will:

- Open a server socket on given port *12345*. So that if another single command
  is already listening this port, this will fail.
- Run given command.
- Release the port when done.

*Enjoy!*
