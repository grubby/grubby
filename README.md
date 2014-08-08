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
* [x] booleans
* [x] arguments on a method decl have optional parens
* [x] call expressions have optional parens
* [x] unary operators
    * [x] !
    * [x] +
    * [x] -
    * [x] ~
* [ ] binary operators
    * [x] +
    * [x] -
    * [ ] *
    * [ ] /
    * [ ] %
    * [ ] <<
    * [ ] >>
    * [ ] & (bitwise and)
    * [ ] | (bitwise or)
    * [ ] ^ (bitwise exclusive)
    * [ ] <= >= and <>
    * [ ] &&
    * [ ] ||
    * [ ] <=> == === != =~ !~
    * [ ] fix associativity of unary + and - (try "abc - 123")
* [x] hashes
    * [x] => syntax
      [x] ruby 1.9 {key: value} syntax
* [x] arrays
* [ ] blocks
    * [ ] with block params
* [ ] heredoc
* [ ] globals ($: and $?, et cetera)
* [ ] backtics
* [ ] string interpolation
* [ ] semicolons
* [ ] method calls
* [x] comments
* [ ] go docs (see: http://godoc.org/github.com/robertkrimen/otto)

Grubby interpreter
##################
* [ ] write a simple REPL / interpreter
* [ ] write some simple tests based on Otto: http://github.com/robertkrimen/otto
* [ ] what does ObjectSpace need to look like?
* [ ] what are the builtins we need to have by default?
    (Object has 97 class methods, 56 instance methods)
    [ ] by default method calls go to Kernel
    [ ] builtins
        [ ] sprintf
        [ ] format
        [ ] Integer
        [ ] Float
        [ ] String
        [ ] Array
        [ ] Hash
        [ ] warn
        [ ] raise
        [ ] fail
        [ ] global_variables
        [ ] __method__
        [ ] __callee__
        [ ] __dir__
        [ ] eval
        [ ] local_variables
        [ ] iterator?
        [ ] block_given?
        [ ] catch
        [ ] throw
        [ ] loop
        [ ] trace_var
        [ ] untrace_var
        [ ] at_exit
        [ ] syscall
        [ ] open
        [ ] printf
        [ ] print
        [ ] putc
        [ ] puts
        [ ] gets
        [ ] readline
        [ ] select
        [ ] readlines
        [ ] `
        [ ] p
        [ ] test
        [ ] srand
        [ ] rand
        [ ] trap
        [ ] exec
        [ ] fork
        [ ] exit!
        [ ] system
        [ ] spawn
        [ ] sleep
        [ ] exit
        [ ] abort
        [ ] load
        [ ] require
        [ ] require_relative
        [ ] autoload
        [ ] autoload?
        [ ] proc
        [ ] lambda
        [ ] binding
        [ ] caller
        [ ] caller_locations
        [ ] Rational
        [ ] Complex
        [ ] set_trace_func
        [ ] freeze
        [ ] ===
        [ ] ==
        [ ] <=>
        [ ] <
        [ ] <=
        [ ] >
        [ ] >=
        [ ] to_s
        [ ] inspect
        [ ] included_modules
        [ ] include?
        [ ] name
        [ ] ancestors
        [ ] instance_methods
        [ ] public_instance_methods
        [ ] protected_instance_methods
        [ ] private_instance_methods
        [ ] constants
        [ ] const_get
        [ ] const_set
        [ ] const_defined?
        [ ] const_missing
        [ ] class_variables
        [ ] remove_class_variable
        [ ] class_variable_get
        [ ] class_variable_set
        [ ] class_variable_defined?
        [ ] public_constant
        [ ] private_constant
        [ ] singleton_class?
        [ ] include
        [ ] prepend
        [ ] module_exec
        [ ] class_exec
        [ ] module_eval
        [ ] class_eval
        [ ] method_defined?
        [ ] public_method_defined?
        [ ] private_method_defined?
        [ ] protected_method_defined?
        [ ] public_class_method
        [ ] private_class_method
        [ ] instance_method
        [ ] public_instance_method
        [ ] nil?
        [ ] =~
        [ ] !~
        [ ] eql?
        [ ] hash
        [ ] class
        [ ] singleton_class
        [ ] clone
        [ ] dup
        [ ] taint
        [ ] tainted?
        [ ] untaint
        [ ] untrust
        [ ] untrusted?
        [ ] trust
        [ ] frozen?
        [ ] methods
        [ ] singleton_methods
        [ ] protected_methods
        [ ] private_methods
        [ ] public_methods
        [ ] instance_variables
        [ ] instance_variable_get
        [ ] instance_variable_set
        [ ] instance_variable_defined?
        [ ] remove_instance_variable
        [ ] instance_of?
        [ ] kind_of?
        [ ] is_a?
        [ ] tap
        [ ] send
        [ ] public_send
        [ ] respond_to?
        [ ] extend
        [ ] display
        [ ] method
        [ ] public_method
        [ ] singleton_method
        [ ] define_singleton_method
        [ ] object_id
        [ ] to_enum
        [ ] enum_for
        [ ] equal?
        [ ] !
        [ ] !=
        [ ] instance_eval
        [ ] instance_exec
        [ ] __send__
        [ ] __id__
* [ ] benchmark against jruby / mri for something simple
