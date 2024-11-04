# Cohnect (A declarative scriptable backend game server.)

## About

- Is a scriptable backend game server written in go
- Scriptable via modules
- Provided client sdks in different languages
- Communicates over UDP with custom encrypted packets and algorithms to overcome some limitations
  - Packet Structure
    - 1byte: VERSION NUMBER
    - 36bytes: CLIENT/SERVER UUID
    - 2byte: OPCODE
    - byte[]: BODY
  - BuiltIn Operations:
    - GET (0x01): Retrieving some data that has been stored
    - SET (0x02): Setting some data
    - SET_CLIENT_TAGS (0x03)
    - EXECUTE (0x04): Execute some action on the server via a registered module
    - STREAM (0x05): Register client to listen for packets based on a specified topic until the explicitly closed or the server finishes streaming
    - BROADCAST (0x06): Broadcast a message through a pipeline of registered modules
- Modules Written in LUA+GO
  - Module functionality
    - Interacting with clients by sending packets to them via direct, pattern matching via client tags, all clients
    - Interface with Registered Systems (Other Modules)
    - Interface with Databases
    - Interface with Message Queues
    - Interface with Caches
    - Interface with Server State (Ex list of clients)
    - Interface with Server Feature Flags
    - Interface with Server Instances (Ex Sharding, Clusters, Spawn New Instances etc.)
    - Interface with Server Hooks (Startup, Shutdown, Etc)
    - Logging via Sinks (Console, Files, APIs, Databases)
  - Module Types:
    - Pipelines
    - Background Long Running
    - Cron Jobs
    - Action Scripts
- Initial Client SDKs
  - Rust?
