# UII CLI client
This is a command line client for Virtomize Unattended Install Images API. 
It is intended both for day to day use, as well as to demonstrate the usage of the API for other projects. 

## Requirements
In order to use this client, you need an API token form the cloud api.
A token can be created using the [Virtomize website](virtomize.com).
This token can be passed to each  command using `--token=123456`.
Alternatively, the client will use the value provided in the `UIIAPITOKEN` environment variable.


## Usage
The client supports the following commands:

### Help
Use `--help` or `-h` to print a list of all available commands.

### Building an ISO

The `build` command can be used to build an ISO.
The following example command will build an ISO file named `foo.iso`, which will install Debian linux in version 9 using `hostname` as the hostname. 

Example: 
```shell
uii-go-cli build foo.iso debian 9 hostname --token=1234567890ABCDEFGHIJKL
```

### List available operation systems

The `ls os` command can be used to list all available operating systems.

Example:
```shell
uii-go-cli ls os
```

### List available packages for an operating systems

The `ls package` command can be used to list all available operating systems.

Example:
```shell 
uii-go-cli ls package debian 9
```

This command can produce quite a bit of output.
Piping the result into less is a good way to address this. 

```shell
uii-go-cli ls package debian 9 | less
```
