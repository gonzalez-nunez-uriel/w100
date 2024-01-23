# Second Package Complete

2024/01/23 14:35

PASS
ok  	github.com/gonzalez-nunez-uriel/w100/internal/linematch	0.001s

A little anticlimactic. Did it really work? I was expecting another failing test.

What now?

Refactoring. I don't like the way the code looks. But I am still concerned about the correctness of
the code. I had a lapsus brutus when I was working on the else clause. That is the reason why I
placed the two comments at L42 and L43. I was baby-talking myself into accepting that just
returning the current line number was the right answer. The only reason why that is the right answer
is because of the two if-guards at the begginging of the function. That was also the reason that I
placed the comment at L14. Without those if guards that would not be valid ... I think. This is the
reason why I need to refactor the code to make it easier to understand.

But I still feel uncomfortable about the correctness of the code. I think I am going to come up with
a couple more tests before I do the refactoring just to be sure. No. I don't have to wait for more
tests before I refactor. I can refactor now and do more tests latter. I have good discipline when
writting the tests, so it should be fine.

In regards to the tests, they all have the same format. Should I extract them into a file with a
format? I mean, I am already behind schedule. Even if the tests are the same now, they could change
in the future. Specially if the requirements of the package change. I think that is unlikely. I do
like how you can grow the tests one by one in code. All you need to know is the format that the
tests follow. Should I do some nested testdata and use files? That makes sense, but is overkill. I
think the application is fine as it is. I am just going to write some more tests that use more empty
lines, since that was the edge case that I was not considering.
