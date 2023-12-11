# manglename(1) - mangle file names

Kozmix Go, 15-JUL-2021

```
manglename [-u -l -c ] [-L lang -v -e -h] name ...
```


<a name="description"></a>

# Description

**manglename**
changes the capitalisation of file names passed on the command line
according to the option given:


* **-u**  
  make file names uppercase
* **-l**  
  make file names lowercase
* **-c**  
  capitalise file names
  

Exactly one of these three options must be given. In addition:


* **-e**  
  mangle the file extension as well
* **-L lang**  
  use capitalisation rules for language "lang" (defaults to English)
* **-v**  
  be verbose about what is happening
* **-h**  
  print some help
  

<a name="notes"></a>

# Notes

If
**manglename**
is invoked as
**capitalise,**
it will act as if
**manglename -c**
had been called. This is for backwards compatibility with an earlier
command by that name.


<a name="see-also"></a>

# See Also

**capit**(1)


<a name="bugs"></a>

# Bugs

None.


<a name="author"></a>

# Author

svm
