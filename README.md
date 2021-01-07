# Todo To Docs
Tool to inspect a package for todos and compiles all found todos into a single document.

This tool currently only works on `go` code but can be easily expanded to analyse other code bases.

## Install

`go get github.com/status-im/todo-to-docs`

## Run

Currently the tool is hardcoded to process the `status-go` repo and assumes that the repo is present on the user's machine.

`cd` into the `github.com/status-im/todo-to-docs` repo.

Run the following

```bash
go build
todo-to-doc
```

## Sample Output

- 📁 params<br/>
  - 📃 config.go<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (c *NodeConfig) updatePeerLimits()`<br/>
      *On line*              : 723 <br/>
      *Description*          : TODO(dshulyak) register mailserver limits when we will change how they are handled.<br/>
- 📁 peers<br/>
  - 📃 peerpool.go<br/>
    - ⬜ Todo:<br/>
      *On line*              : 249 <br/>
      *Description*          : TODO(adam): split it into peers and discovery management loops. This should<br/>
       simplify the whole logic and allow to remove `timeout` field from `PeerPool`.<br/>
  - 📃 topicpool.go<br/>
    - ⬜ Todo:<br/>
      *Related func or type* : `func (t *TopicPool) AddPeerFromTable(server *p2p.Server) *discv5.Node`<br/>
      *On line*              : 401 <br/>
      *Description*          : TODO(adam): investigate if it's worth to keep the peer in the queue<br/>
       until the server confirms it is added and in the meanwhile only adjust its priority.<br/>

For a full sample output [see here](sample/sample.md)
