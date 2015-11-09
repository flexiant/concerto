# Concerto CLI / Go Library
[![Build Status](https://drone.io/github.com/flexiant/concerto/status.png)][clilatest] [![GoDoc](https://godoc.org/github.com/flexiant/concerto?status.png)](https://godoc.org/github.com/flexiant/concerto)

Flexiant Concerto Command Line Interface allows you to interact with Concerto features, and build your own scripts calling Concerto's API.

If you already know Concerto CLI, and only want to obtain the latest version, download Concerto from [this link][clilatest]

If you want to build the CLI using the source code, please, take into account that the master branch is used for development, it is unstable and might be broken. Download stable tagged versions to use Flexiant Concerto CLI.

# Setup

## Pre-requisites
Before setting up the CLI, we will need a Flexiant Concerto account, and an API key associated with your account.

You can create a free account in less than a minute following the steps in Flexiant Concerto [login page](https://start.concerto.io/).

<img src="./docs/images/signup.png" alt="sign up" width="300px" >

Once your account have been provisioned, navigate the menus to `Settings` > `User Details`
and scroll down until you find the `New API key` button.

<img src="./docs/images/newAPIkey.png" alt="new API key" width="500px" >

Pressing `New API Key` will download a compressed file that contains the necessary files to manage authenticate with Concerto API and manage your infrastructure. Keep it safe.

Extract the contents with your zip compressor of choice and continue using the setup guide for your O.S.

## Linux and OSX
### Configuration
Concerto configuration will usually be located in your personal folder under `.concerto`. If you are using root, concerto will look for contiguration files under `/etc/concerto`.
We will assume that you are not root, so create the folder and drop the certificates to this location:
```
$ mkdir -p ~/.concerto/ssl/
$ unzip -x api-key.zip -d ~/.concerto/ssl
```

Concerto CLI expects a configuration file to be present containing:
 - API Endpoint
 - Log file
 - Log level
 - Certificate location

This command will generate the file `~/.concerto/client.xml` with suitable contents for most users:
```
$ cat <<EOF > ~/.concerto/client.xml
<concerto version="1.0" server="https://clients.concerto.io:886/" log_file="/var/log/concerto-client.log" log_level="info">
 <ssl cert="$HOME/.concerto/ssl/cert.crt" key="$HOME/.concerto/ssl/private/cert.key" server_ca="$HOME/.concerto/ssl/ca_cert.pem" />
</concerto>
EOF
```

We should have in our `.concerto` folder this structure:
```
/Users/myuser/.concerto
├── client.xml
└── ssl
    ├── ca_cert.pem
    ├── cert.crt
    └── private
        └── cert.key
```
### Binaries
Download linux binaries from this [URL][clilatest] and place it in your path.
```
sudo curl -o /usr/bin/concerto https://drone.io/github.com/flexiant/concerto/files/concerto.x64.linux
sudo chmod +x /usr/bin/concerto
```

To test the binary execute concerto without parameters
```
$ concerto
NAME:
   concerto - Manages comunication between Host and Concerto Platform
USAGE:
   concerto [global options] command [command options] [arguments...]
VERSION:
   0.1.0
AUTHOR(S):
   Concerto Contributors <https://github.com/flexiant/concerto>
COMMANDS:
   firewall	Manages Firewall Policies within a Host
   scripts	Manages Execution Scripts within a Host
...
```
To test that certificates are valid, and that we can communicate with Concerto server, obtain the list of workspaces at your Concerto account using this command
```
$ concerto cloud  workspaces list
ID                         NAME                  DEFAULT        DOMAIN ID                  SSH PROFILE ID             FIREWALL PROFILE ID
56017273fef51ac13400002c   default               true           5601726ffef51ac134000028   56017273fef51ac13400002b   56017273fef51ac13400002a
56388361cfda105f6e000502   Wordpress_workspace   false          5601726ffef51ac134000028   56017273fef51ac13400002b   56388360cfda105f6e000501
```

###Troubleshooting
If you got an error executing concerto CLI:
 - execute `which concerto` or `whereis concerto` to make sure that the binary is installed
 - execute ls -l /path/to/concerto with the output from the previous command, and check that you have execute permissions
 - execute `$PATH` and search for the path where concerto is installed. If concerto isn't in the path, move it to a $PATH location.
 - check that your internet connection can reach clients.concerto.io
 - make sure that your firewall lets you access to https://clients.concerto.io:886
 - check that  client.xml is pointing to the correct certificates location
 - if concerto executes but only shows server commands, you are probably trying to use concerto from a commissioned server, and the configuration is being read from `/etc/concerto`. If that's the case, you should leave concerto configuration untouched so that server commands are available for our remote management.


# Usage
We will include some use cases here. If you can't wait for them, please, contact us at <contact@flexiant.com>.

# Contribute

To contribute
 - Find and open issue, or report a new one. Include proper information about the environment, at least: operating system, CLI version, steps to reproduce the issue and related issues. Avoid writing multi-issue reports, and make sure that the issue is unique.
 - Fork the repository to your account
 - Commit scoped chunks, adding concise and clear comments
 - Remember to add tests to your contributed code
 - Push changes to the forked repository
 - Submit the PR to Concerto CLI
 - Let the maintainers give you the LGTM.

Please, use gofmt, golint, go vet, and folow [go style](https://github.com/golang/go/wiki/CodeReviewComments) advices

[clilatest]: https://drone.io/github.com/flexiant/concerto/latest
