# govern

A silly, proof-of-concept project for writing a configuration management
system, in Go.

## TODO

* [ ] Pick a decent configuration syntax (I'm thinking [HCL](https://github.com/hashicorp/hcl)

## Goals

*	Stick to the Go standard library, as much as possible

## Pipe dreams

*	Be successful

## Not-quite pipe dreams

* Provide a silly, little script that can be run through cron to send machine
  information to an [etcd](https://github.com/coreos/etcd) cluster, via cURL,
  essentially creating a remote inventory, which would mean...
* Configurable inventory providers; files are nice and dependable, but in the
  spirit of service-oriented architectures, it would be nice to have a "live"
  inventory somewhere, so that we don't try to connect to boxes that are
  unresponsive

## Documentation

To view the documentation for govern (which is still a *huge* work in progress),
they are generated by [viewdocs](http://viewdocs.io).

[Documentation](http://nesv.viewdocs.io/govern)
