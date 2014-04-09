# Host inventory

The first step in using govern is to declare your hosts, and organize them into
groups.

When executing govern, it will look in your current working directory for a file
called `hosts`, but you can specify a different file by passing the `-i` flag to
govern, like so:

	$ govern -i=/path/to/inventory-file
	
Here is a very basic example of a hosts file:

	---
	- group: webservers
	  hosts:
	    - web01
		- web02
	
	- group: loadbalancers
	  hosts:
	    - lb01
		  
