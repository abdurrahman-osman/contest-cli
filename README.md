# contest-cli 

Contest-cli is a tool to test network connection from source to destinations.

## Main Goals:
1. Human Developed: contest-cli will be developed by a human. There will be no vibe coding. no hallucination. As an opensource project should be. AI will only be an assistant.

2. Various protocol support: TCP, UDP, WebSocket, HTTP, HTTPs, Postgres, Mongo, Kafka, Redis etc.

3. Various source options: local, k8s, ssh, etc.

## Current Status:
- Protocols: TCP & UDP are ready.

- Source: Local, SSH(TCP Only).
## Build


## Examples:
```
contest ssh --target=google.com:80 --proto=tcp --port=443 --hosts 10.218.16.80 --key=~/.ssh/key --user=username
```

## TODO
- Ability to parse multiple hosts for ssh.
- Multi host targets.
- Multi port targets.
- Multi protocol targets.

