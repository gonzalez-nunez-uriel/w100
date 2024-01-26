# Reflection

The implementation was simple. Straighforward.

I liked how I splited things around. The refactoring was acceptable. The core of the algo takes only
27 lines of code, but I am still concerned about its readability. Can people figure out what it is
supposed to be doing? I still have to properly document the package but just from the naming
conventions and the function names, is this good enough? I am satisfied, but I don't want to be a
conformist.

## Thoughts on software specification

I don't think that it was particularly difficult, but there was some mental adjustments that I had
to make, and was not expecting, that indicate that there is more to learn about software
specification. I know that I went on a tanget doing all of these utilities but the peace of mind
that I am getting is worth it. I know that by the end of this small project I am going to be
confident that I can modify the code without loosing any functionality or compromising the integrity
of my program. I also feel confident that any future contributor, if they have the wrong mental
model, will trip the tests and correct their mental model. That is, it is going to help any future
contributor to corroborate that they understand what the code is supposed to do (assuming that they
learn what the code does and what the code is about before modifiying the test - note to self: never
modify tests without understanding the code base).

Something that I was not expecting was how much working on the tests helpe discover the algorithm
and how it helped me make architectural decisions about the code. In particular, it helped me create
the if guards at the entry point of the function. I am aware of the potential performance gains from
doing something that is more involved with traversing the string char by char using a state
automaton, but this is a great example of 'get working code right away'. Now that I have working
code, I can take my time to develop the automaton.

The most critical lesson was when I went on a tangent with the edge cases. Fortunately, by thinking
critically about *why* I was writting the package in the first place, I was able to determine that I
had made a mistake in the way that I was thinking about the problem, which leaked into the test
specifications. As an architect, you have to constantly question *why* you are doing things the way
you are doing them. It needs to become a having - to be constantly on the lookout for derailments of
thougth and design.

There is something about looking ahead that is really important when it comes to specification. The
linematch package was created as a utility. Even though it is possible to make sense of it on its
own, the context is really important.

I am thinking of making an improvement. The utility that consumes FindLineMismatch makes a
comparison of the string, checking for equality. But FindLineMismatch is already doing that. Taking
inspiration from strings.Cut/2, instead of returning just the integer we can use a bool to indicate
that a mismatch was found. Even though technically speaking the result of -1 is good enough to
handle this case, I think it using a boolean value is a better API design. The cost is minimal and
the benefits are immesurable.
