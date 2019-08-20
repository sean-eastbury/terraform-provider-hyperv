HyperV Terraform Provider
=========================
Forked from https://github.com/taliesins/terraform-provider-hyperv
#This is beta code.

- [Website](https://github.com/sean-eastbury/terraform-provider-hyperv)
- [Releases](https://github.com/sean-eastbury/terraform-provider-hyperv/releases)
- [Documentation](https://github.com/sean-eastbury/terraform-provider-hyperv/tree/master/website/docs)
- [Issues](https://github.com/sean-eastbury/terraform-provider-hyperv/issues)

![Hashi Logo](https://cdn.rawgit.com/taliesins/terraform-provider-hyperv/master/website/logo-hashicorp.svg "Hashi Logo")

Features
------------
- Remote scheduled task powershell runner does not run into issues with escaping variables or escaping between the different scripting layers.
- Changed Winrmcp to use Powershell commands directly rather then use base64 encoded strings as we want to prevent Powershell progress leaking.
-Changed Winrmcp to return path of files on remote box as the location of $env:temp can change in Powershell depending on the session instance.
- Runs all HyperV commands remotely i.e. so the provider can run on a linux machine and connect remotely to a windows machine running HyperV.
- Almost all functionality of Powershell HyperV commandlets for the resources is exposed via Terraform resources.
- Resource - Network Switch
- Resource - VHD
- Resource - Virtual Machine Instance
  - Network adaptors
  - Hard drives
  - Dvd drives

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)
-   Connectivity and credentials to a server running HyperV on Windows 10, Windows Server 2016 or newer

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/sean-eastbury/terraform-provider-hyperv`

```sh
$ mkdir -p $GOPATH/src/github.com/tidalf; cd $GOPATH/src/github.com/tidalf
$ git clone https://github.com/sean-eastbury/terraform-provider-hyperv.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/taliesins/terraform-provider-hyperv
$ make build
```

Using the provider
----------------------
## Fill in for each provider

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-hyperv
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
