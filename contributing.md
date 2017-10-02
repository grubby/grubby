History
-------

Grubby was a two year experiment, starting just before my son was born and running until just after his second birthday. The primary hypothesis was "it is possible to build a programming language in 30 minutes a day, using TDD". This was a resounding success, by any measure.

The secondary hypothesis was "it is possible to build a beautiful codebase as a sleep-deprived parent". As far as that is concerned, well that's why you're reading this, right ?

Who am I ?
----------

As of this writing, I'm Tim Jarratt - an engineering manager working at Pivotal Labs in Paris, France. Around the time I became both a father and a manager, I stopped having as much time for grubby, and while that makes me sad, I am committed to helping others contribute to the project. Consider these an initial pass at making the project easier to contribute to - but I'm often available for 10 or 30 minutes of consultation a day if you're really interested in contributing to grubby.

How to contribute
-----------------

Let's use, as an example, the concrete case of adding a `#reverse` method to `Array`.

* Try as much as possible to follow pure TDD
* Find the appropriate test file (in thise case, it would be `vm/array_test.go`
* Write the smallest possible spec that captures the behavior you believe is important
* Really, it's fine to have lots of small tests. Try to make it as small as possible.
* Ensure you have a correctly failing test before writing any code
* Go to the implementation (`vm/builtings/array.go`) and see where builtin methods are added
* Define a new method that you think implements the behavior correctly
* Run the tests (and continue writing just the code you need until it passes)
* Refactor until you are satisfied with the feature
* Continue adding test cases until you are satisfied


For more complicated cases
--------------------------

Let's assume you want to implement a more complicated method, such as `Kernel#class`

* fire up MRI's irb and run this code `method(:class)`
* keep track of which class or module defines the method
* poke around in grubby's interpreter/vm/builtins directory and see if that module or class is defined
* start by defining the module if it's missing
* write a small test for the behavior you want to see
* write some simple, (often naive) code that makes that test pass
* continue previous steps, refactoring until you are happy
* BAM, you just wrote a feature. Commit!


Hope you find these loose, adhoc instructions helpful!
