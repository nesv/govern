# govern

A silly, proof-of-concept project to see how difficult it would be to implement
an [Ansible](http://www.ansible.com/home)-like system in Go, while not relying
on a remote Python interpreter. Trying to do this all without some form of
remotely-executable code would be silly, so instead of relying on Python, or
Ruby, or TCL, or whatever, just use sh-compatible scripting and execute the
scripts using bash or zsh, or the like.

## The goals

* Just SSH for the transport layer
* YAML for configs/"plays"
* Compile config file templates locally, then ship them out to the hosts (this
  would need to be done after getting information about the host)

## Things to consider

* Getting information about a host (refer to Ansible's
  [setup](http://docs.ansible.com/setup_module.html) module)
* Stick to YAML for an inventory file; don't switch between YAML and INI-style

## Pipe dreams

* Full compatibility with Ansible; should be able to use `govern` in place of
  the `ansible` and `ansible-playbook` commands
* Be successful

## Not-quite pipe dreams

* Provide a silly, little script that can be run through cron to send machine
  information to an [etcd](https://github.com/coreos/etcd) cluster, via cURL,
  essentially creating a remote inventory, which would mean...
* Configurable inventory providers; files are nice and dependable, but in the
  spirit of service-oriented architectures, it would be nice to have a "live"
  inventory somewhere, so that we don't try to connect to boxes that are
  unresponsive
