# check_tftp

An Icinga check plugin to check TFTP servers.
This is basically a rewrite of the existing perl check plugin check_tftp: http://william.leibzon.org/nagios/

## Usage 

```bash
Usage:
    check_tftp [flags]

Flags:
    -h, --help              help for check_tftp
    -H, --hostname string   Hostname or IP-Address of the TFTP server
    -f, --file string       File to download from the TFTP server
    -C, --checksum string   Checksum of the File
    -v, --version           version of check_tftp

## Build it yourself

Prerequisites: Glibc and golang
Inside the plugin directory simply type
```go
go build .
```

## License

Copyright (c) 2023 [NETWAYS GmbH](mailto:info@netways.de)

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
License as published by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program. If not,
see [gnu.org/licenses](https://www.gnu.org/licenses/).

