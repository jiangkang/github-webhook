[English](README.md)|[中文版本](README-zh.md) 

## github-webhook
A Github webhook implemention written by Go

## what github-webhook can do
You can use "github-webhook" binary executable file to receive the notification 
when your github project trigger the push event.
Once you can receive the push event,you can deploy the website project auto.

## How to install

- build from source
```shell
git clone https://github.com/jiangkang/github-webhook.git
cd github-webhook
go build
```

## How to use
1. create the hook_config.json in the same path of the github-webhook binary
```json
{
  "path_repo" : "/var/www",
  "port": "5678"
}

```
2. in path_repo directory mentioned above,you have to git pull the website project from github

3. execute the script

```
./github-webhook
```
4. add webhook info in your github project

url : "yourdomain/payload"

such as :

```
https://jiangkang.tech/payload
```
