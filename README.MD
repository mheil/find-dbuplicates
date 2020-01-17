find-duplicates
===============

desciption
----------

`find-duplicates` is a command line utilities that searches multiline strings and prints the
line-numbers of the occurances for non-unique lines. `find-duplicates` either process the 
first command line argument, or stdin, of no command line arguments have been given.

usage
-----

```
find-duplicates "content1
content2
content1
content3"
``` 

will output:
`line 'content1' has been found in multiple lines: [1 3]`

