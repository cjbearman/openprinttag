# tagtool
For reading, writing and otherwise working with the following tags:
* ICode SLIX
* ICode SLIX2
* ST25DV04
* ST25DV16
* ST25DV64

Using an ACS ACR-1552-U USB Reader, on Windows, Mac or Linux

# Installation
Install with go install ./...

For linux, you will need to install the following packages:
* pcscd 
* libpcsclite1
* pcsc-tools
* libacsccid1 1.1.13-1~bpo12+1 (or higher)

N.B. The required version of libacsccid1 will probably require a manual download from https://acsccid.sourceforge.io/

Mac and Windows should work without dependencies

# General operation

## Reading a tag
Read a tag's user memory space and output to specified filename. Use "-" for the filename to output to stdout. Add -hex option to dump output in hex.
```
ntagtool -r filename
```

To limit the number of bytes read use the -n option.

To output a hex dump of the data, add -hex


## Writing a tag
Write a tag's user memory space from the contents of the specified filename. Use "-" to accept data from stdin.
```
ntagtool -w filename
```

## ID a tag
```
ntagtool -i
```
