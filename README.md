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

Roadmap
-------

Here's a short list of areas I'm planning on working on in the near future:

(NB: This will likely get out of date quickly)

Grubby yacc
###########
* [x] parse an integer
* [x] parse a float
* [x] remove warnings for DIGIT and FLOAT (?)
* [x] remove any dead code that isn't helping (because it assumes int or w/e)
* [x] throw some ast nodes into the parser
* [x] refactor lexer to be less error prone && fewer lines of code
* [x] parse simple strings
* [x] parse symbols
* [x] parse a bare reference to **something**
* [x] parse call expressions
    * [x] one arg
    * [x] with parens
    * [x] many args
    * [x] no args
* [x] parse method definitions
    * [x] no args
    * [x] one arg
    * [x] many args
* [x] parse a class
    * [x] with a super class
    * [x] with a namespace
    * [x] with a body
    * [x] class + instance methods
    * [x] class + class methods
* [x] define a module
* [x] assignment
* [ ] postfix // prefix operations
* [ ] binary operators
* [ ] hashes
* [ ] arrays
* [ ] blocks
* [ ] arguments on a method decl have optional parens
* [ ] call expressions have optional parens
* [ ] heredoc
* [ ] globals ($: and $?)
* [ ] backtics
* [ ] string interpolation
* [ ] semicolons
* [ ] method calls (message passing?)
* [ ] comments

Grubby interpreter
##################
* [ ] write a simple REPL / interpreter
* [ ] what does ObjectSpace need to look like?
* [ ] what are the builtins we need to have by default?
    (Object has 97 class methods, 56 instance methods)
* [ ] benchmark against jruby / mri for something simple
