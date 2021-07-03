### Simple Fiber Endpoint setup

Project demonstrating the [Fiber](https://gofiber.io/) Web framework.

For now just have an endpoint which can be extended.

The endpoint serves,

    - / (root)


Running the main.go program gives terminal output,

```bash
Hello there buddy!

 ┌───────────────────────────────────────────────────┐ 
 │                    Fiber v2.7.1                   │ 
 │               http://127.0.0.1:8000               │ 
 │       (bound on host 0.0.0.0 and port 8000)       │ 
 │                                                   │ 
 │ Handlers ............. 2  Processes ........... 1 │ 
 │ Prefork ....... Disabled  PID .............. 5708 │ 
 └───────────────────────────────────────────────────┘ 

```

Hitting the service,

```bash
$ curl http://localhost:8000
```
gives
```json
{
  "message": "You are at the endpoint 😉",
  "success": true
}
```