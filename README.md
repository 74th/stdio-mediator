# Stdio Tee Logger

prog1 --execute command-> prog2

prog1 --execute-> Logger --execute-> prog2

prog1 shuld execute `stdio-tee-logger prog2`.

```
13:39:49.187327 >: Content-Length: 3031
13:39:49.187340 >:
13:39:49.281470 <: Content-Length: 73
13:39:49.281494 <:
13:39:49.305603 >: {"jsonrpc":"2.0","id":0,"method":"initialize","params": ...
13:39:49.305639 >:
13:39:49.305992 >: {"jsonrpc":"2.0","method":"initialized","params":{}}Content-Length: 99
13:39:49.306005 >:
13:39:49.306498 >: {"jsonrpc":"2.0","method":"workspace/didChangeConfiguration","params":...
13:39:49.306516 >:
```
