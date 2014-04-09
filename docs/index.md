# Govern

Welcome to the documentation for govern!

Govern aims to be a simple, easy-to-use system
orchestration/remote-code-execution tool to help you manage your servers with
ease.

The documentation aims to explain all of the features and workings of govern,
in a procedural fashion; it will try its best to avoid talking about things you
will learn in topics that have yet to be explained.

## Topics

* [Host inventory](/govern/inventory)
* [Plays](/govern/plays)
* [Roles](/govern/roles)
* [Tasks](/govern/tasks)
* [Handlers](/govern/handlers)
* [Variables](/govern/variables)
* [Modules](/govern/modules)

## Introduction

One of the over-arching goals of govern is to prevent you &mdash; the beloved
user &mdash; from having to repeat yourself. To accomplish such a goal, there is
almost always one, single place for you to define or declare something.

To quickly illustrate the how this goal is applied, govern uses the following
"inheritance" model:

* *Plays* assign *roles* to hosts
* *Roles* contain *tasks*
* *Tasks* can notify *handlers*

Other important things worth noting, is that *variables* can only be defined in a
few places; this is to minimize the occurrence of the "hunt-and-peck" process of
trying to figure out why (for example) Apache is not listening on port 80, even
though that is the port you are pretty sure you told it to listen on.

Additionally, *tasks* and *handlers* are nothing more than conditions you can
place around calling a *module*.

## Design details

Govern tries its best to keep the number of extra tools you require on the
managed hosts, to a minimum. As wonderful, and great, of a tool
[Ansible](http://www.ansible.com) is, it still requires you to have a Python
interpreter installed on the managed hosts (you may also have to install a
few Python modules to use some of Ansible's stock modules).

Python is everywhere (and it is a great language!), but we live in an imperfect
world, where larger companies and enterprises may not have the flexibility 
to upgrade to the "latest and greatest" version of their
*operating system-du-jour*, and may be stuck with an ancient version of
Python.

None of that, here! Govern wants to be as plug-n-play as possible. What is more
common than an up-to-date version of Python? A shell.

### Files

The host inventory, plays, roles, tasks, handlers, and variables are all defined
in YAML files. YAML is fairly common these days, and it is easy to read and
write.

### Modules

Govern's modules are nothing more than sh-compatible scripts; while you may not
be able to use the coolest features from bash, or zsh, the idea here is that
modules should be able to run across as many hosts, with different operating
systems as possible.

Granted, even targeting as something as ubiquitous as a POSIX-standard shell is
limiting. Programming languages like Python and Ruby give you extra horsepower,
and that immediately limits usability for managing operating systems that do not
have a POSIX-compliant shell, like Windows.

Truth be told, govern *will not* immediately target Windows machines, so please
do not event ask.

## Execution model

Using nothing more than trusty, old SSH, govern copies the runnable module to
the remote machine, and passes any provided arguments to the script.

For example, let's assume the module `package` has the arguments `name=apache
state=latest`. Govern will use the `scp(1)` command to copy the `package.sh`
script over to the remote machine's `/tmp` directory, then call it, like so:

	/tmp/package.sh --name=apache --state=latest
	
At this point, it is up to the `package.sh` script to figure out what to do.

Govern captures the STDOUT and STDERR streams from the command, and in the event
of a non-zero exit status code, the output is presented to the user.
