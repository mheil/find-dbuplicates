find-duplicates
===============

desciption
----------

`find-duplicates` is a command line utility that searches for non-unique lines
in texts read from stdin or file and prints the line numbers in which they occurred.

usage
-----

```
echo "content1
content2
content1
content3" | find-duplicates
``` 

will output:
`line 'content1' has been found in multiple lines: [1 3]`

