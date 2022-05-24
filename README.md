# UUI CLI client
Command line client for cloud api.

## Requirements
In order to use this client, you need an API token form the cloud api.
This token can be created using the website. 

## Usage
The client supports the following commands:

### Help
Use `--help` or `-h` to get help for the available commands.

Example: 
```buildiso -h```

### Building an ISO

The `build` command can be used to build an ISO.
The following example command will build an ISO file named `foo.iso`, which will install Debian linux in version 9 using `hostname` as the hostname. 

Example: 
```buildiso build foo.iso debian 9 hostname --token=1234567890ABCDEFGHIJKL```

### List available operation systems

The `ls os` command can be used to list all available operating systems.

Example:
```buildiso ls os```

### List available packages for an operating systems

The `ls package` command can be used to list all available operating systems.

Example:
```buildiso ls package debian 9```

This command can produce quite a bit of output.
Piping the result into less is a good way to address this. 

```shell
buildiso ls package debian 9 | less
```
