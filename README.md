grubby [![Build Status](https://secure.travis-ci.org/grubby/grubby.png?branch=master)](http://travis-ci.org/grubby/grubby)
======


Grubby is an experimental ruby written in Golang

Wat
---

Sure, why not? I've been interested in implementing a language for several years now and have finally reached a point where I feel comfortable exploring the idea further. While rewriting a ruby project in golang, I joked to some coworkers about just writing a Ruby implementation in Go to speed up our efforts; eventually the joke became real.

Why Golang?
-----------

Golang doesn't always feel like the best language to implement another language in. I do enjoy that it feels like a very modern C (albeit one without guarantees on how long **any** function can take because of heap allocation and GC). Ultimately, my goal is to learn more about Golang and Ruby through this project. If I find out that Golang is not ideal for implementing other languages, then that is fine.

Running Tests
-------------

Just run `bin/test` from the root directory. This builds the necessary files (currently just the parser) with `goyacc` first.

Contributions
-------------

I'm currently trying to get rubyspec to run against grubby, which requires a small, but sizeable amount of behavior initially. Run `bin/rubyspec` and you should be able to see the next behavior that needs to be implemented there.

Unfortunately, that's not very easy to parallelize between multiple people. Feel free to write a test on the VM for the functionality you'd like to implement. If you're comfortable implementing it as well, that's great, otherwise a pull request to discuss how it might work is more than welcome. :)

Some great things to work on include:

* line numbers in stack traces
* more builtin classes and methods (anything on Class, Kernel, Module, Object, etc... honestly)
* raising and rescuing from exceptions
* while loops
* blocks

If you're not sure where to start, opening up a pull request is also fine. :)
