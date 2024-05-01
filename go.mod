module github.com/exampleowner/data-sync-tool

go 1.19

replace (
	github.com/exampleowner/config-tool => ../config-tool
)

require (
	github.com/exampleowner/config-tool v0.0.0-20231012010748-2e6ae8d99ead
	github.com/leaanthony/clir v1.6.0
	github.com/manifoldco/promptui v0.9.0
)

require (
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	golang.org/x/sys v0.0.0-20181122145206-62eef0e2fa9b // indirect
)
