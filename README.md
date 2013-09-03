goreutils
=========

An implementation of GNU Coreutils(go + coreutils = goreutils. Get it?
[Hahaha...][hh]) in the Go programming language. This project is not meant to
replace coreutils just help me check out what Go is like.

The resulting binaries are a lot bigger than the coreutils versions, largely due
to static linking. They are also a bit slower but still plenty fast.

For now there's only:

+ ``wc``
+ ``fold`` - only the -w flag is implemented.

[hh]: https://dl.dropboxusercontent.com/u/1280403/haha.jpg
