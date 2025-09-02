module github.com/vivgrid/yc-mcp

go 1.24.2

toolchain go1.24.5

require (
	github.com/vivgrid/yc v1.15.2
	github.com/yomorun/yomo v1.20.6
)

replace github.com/vivgrid/yc => ../yc

require (
	github.com/sashabaranov/go-openai v1.40.1 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
)
