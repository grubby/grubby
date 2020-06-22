grubby [![Build Status](https://secure.travis-ci.org/grubby/grubby.png?branch=trunk)](http://travis-ci.org/grubby/grubby)
======


Grubby is an experimental ruby written in Golang

What? Why?
----------

Why not? Go is basically a modern C (given that it has Unicode support and concurrency primitives), and C is the language that a lot of other languages are implemented in. It seems like a worthy experiment to see if Go is a suitable language to implement other languages in.

Yes but Why Golang?
-------------------

Golang doesn't always feel like the best language to implement another language in. I do enjoy that it feels like a very modern C (albeit one without guarantees on how long **any** function can take because of heap allocation and GC). Ultimately, my goal is to learn more about Golang and Ruby through this project. If I find out that Golang is not ideal for implementing other languages, then that is fine.

How you can help!
-----------------
1. Find some ruby behavior you would like to write
2. Write a failing spec for it
3. Write the code that that makes it pass
4. Issue a pull request. It will bring a smile to my face.

Running Tests
-------------

Just run `bin/test` from the root directory. This builds the necessary files (currently just the parser) with `goyacc` first.

Contributions
-------------

I'm currently trying to get rubyspec to run against grubby, which requires a small, but sizeable amount of behavior initially. Run `bin/rubyspec` and you should be able to see the next behavior that needs to be implemented there.

Unfortunately, that's not very easy to parallelize between multiple people. Feel free to write a test on the VM for the functionality you'd like to implement. If you're comfortable implementing it as well, that's great, otherwise a pull request to discuss how it might work is more than welcome. :)

Some great things to work on include:

* more builtin classes and methods (e.g.: String, Array, Hash, Enumerable methods)
* raising and rescuing from exceptions
* finding edge cases in the parser (using grubby with existing gems would be quite helpful)

If you're not sure where to start, opening up an issue or pull request with failing tests is also fine. :)

Contributors
------------

* [John Foley](https://github.com/jfoley)
* [Brian Butz](https://github.com/butzopower)
* [Trent Ogren](https://github.com/misfo)
* [Tim Jarratt](https://github.com/tjarratt)
