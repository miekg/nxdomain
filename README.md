# nxdomain

## Name

*nxdomain* - shortcut domain resolution by returning NXDOMAIN directly.

## Description

*nxdomain* takes a list of domains and immediately returns NXDOMAIN for any name under them instead
 of taking the long road of trying to resolve them.

## Syntax

~~~ txt
nxdomain [ZONE]...
~~~

## Examples

NXDOMAIN everything in the `.com` zone:

~~~ corefile
com {
    nxdomain
    whoami
}
~~~

NXDOMAIN *everything* (might be a bad idea):

~~~ corefile
. {
    nxdomain
    whoami
}
~~~

# Bugs

The list of zones is just a slice that is traversed, meaning this plugin will get slow when a lof of
names are to be shortcut.
