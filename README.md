# robo

`robo` is using OpenAI API to convert natural language to shell commands.

## Usage
Get `robo`:
```shell
go install github.com/damejeras/robo@latest
```
To use `robo` you have to set `OPENAI_API_TOKEN` environment variable.
```shell
export OPENAI_API_TOKEN=<your_api_token>
```
Use `robo`:
```shell
$ robo show process that is using port 8080
lsof -i :8080
```
