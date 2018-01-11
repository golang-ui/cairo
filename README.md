## cairo bindings

Work In Progress.

Full Cairo bindings with examples:
* Quartz-backed rendering for OS X;
* GLFW crossplatform OpenGL 2.1 surface support.

### Installation

To use libcairo and dependencies shipped with this package (`darwin_amd64` at the moment):
```
$ go get github.com/golang-ui/cairo/cmd/cairo-example-glfw
```

To dynamically link against external libcairo (via pkg-config, `darwin` systems at the moment):
```
$ go get -tags external github.com/golang-ui/cairo/cmd/cairo-example-glfw
```

### Examples

<img alt="cairo-example-glfw" width="400px" src="/cmd/cairo-example-glfw/screenshot.png" />

### LICENSE

MIT
