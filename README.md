goreutils
=========

An implementation of GNU Coreutils in the Go programming language. This project
is not meant to replace coreutils just help me check out what Go is like.

The resulting binaries are a lot bigger than the coreutils versions, largely due
to static linking. They are also a bit slower but still plenty fast.

For now there's only:

+ ``wc`` - with only multibyte counting missing (as far as I know).
+ ``fold``
