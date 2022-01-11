# proxym8

_a proxy mate_

This is a tool to be used for generating static go module proxy files, which can be used to resolve go modules stored in
git on various platforms.

## usage

Install it using `go install proxym8.ruud.ninja/proxym8`

Then run it with `proxym8 -in m8.yaml -out ./some-out-folder`

*Note: The output folder needs to be cleaned manually*

See the `m8.yaml` file for an example on how to configure.

## generate docs for project

`go run main.go -out docs`