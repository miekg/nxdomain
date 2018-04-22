# nxdomain

## Name

*nxdomain* - shortcut domain name by returning NXDOMAIN directly

## Description

*nxdomain* takes a list of domain and immediately return NXDOMAIN for any name in them
instead of taking the long road of trying to resolve them.

## Syntax

~~~ txt
nxdomain DOMAIN [DOMAIN]...
~~~

## Examples

NXDOMAIN everything in the `.com` zone:

~~~ corefile
. {
    nxdomain com
    whoami
}
~~~

NXDOMAIN *everything* (might be a bad idea):

~~~ corefile
. {
    nxdomain .
    whoami
}
~~~
