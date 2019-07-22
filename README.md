# senv-template

To help you to populate your environment variables using AWS SSM Parameter Store. 

## Installation

### Using go

Execute the following command

```shell
$ go get -u github.com/luisc09/senv-template
```

### Using bash 

Execute the following command to install in Linux

```shell
$ curl -L https://github.com/luisc09/senv-template/releases/download/latest/senv-template-linux-amd64 > /usr/local/bin/senv-template && chmod +x /usr/local/bin/senv-template
```

To install it on MacOS execute the following command

```shell
$ curl -L https://github.com/luisc09/senv-template/releases/download/latest/senv-template-darwin-amd64 > /usr/local/bin/senv-template && chmod +x /usr/local/bin/senv-template
```

### Usage

1. Set your credentials through the AWS CLI. 
1. The element to replace must be the full SSM Parameter Store path surrounded by two brackets.
   ```
   db_password={{/dev/db/passowrd}}
   ``` 
1. Execute the following command to create the environment file. 
   ```
   senv-template --file env.dev.tpl --output .env
   ```
1. A new file will be created with all the elements replaced.