# Figuring out the algoruithm by working with case two

## Manual replication

Moving to line one.

It has 48 chars.

It does not exceed the char limit.

There is no need to trim the line.

Do I need to merge the next line?

The next line is empty.

Therefore, I do not need to merge the next line.

Moving to line two.

It is empty.

Therefore, there is nothing to be done to the second line.

Moving to line three.

The line exceeds the char limit.

The first char of the first instance of the word "coughing" is at the 101th char place.

The closest whitespace to the left of the char limit is at the 100th char position, after the first
instance of the word "from".

That whitespace needs to be replaced by a newline char.

Now the third line satisfies the two conditions: 1) the line does not exceed the char limit and 2)
the first word of the next line does not fit in the leftover space that this line has.

Moving to line four.

The line exceeds the char limit.

The char at the 101th position is a whitespace.

That whitespace needs to be replaced with a newline char.

Now the fourth line satisfies the two conditions: 1) the line does not exceed the char limit and 2)
the first word of the next line does not fit in the leftover space that this line has.

Moving to line 5.

The line exceeds the char limit.

The char at the 101th position is not a whitespace.

The closest whitespace to the left of the char limit is at the 97th char position, after the first
instance of the word "two".

That whitespace needs to be replaced by a newline char.

Now we can move to the next line.

*I think that is enough to extract the algorithm*

## Extracting the algorithm

The main point of interest is to verify that every line in a file satisfies the following
conditions:

1. The line does not have more than 100 characters
1. There are no whitespaces between the last character and the newline charcter
1. The first word in the next line cannot be added to the current line without breaking the
   character limit; with the exception of an empty line. In the event of an empty line, the current
   line satisfies this by default.
   
There are two operations that we might perform on a line: 1) trim it or 2) merge it with the next
line. Trimming is when you replace a whitesapce char with a newline char so that the current line
satisfies the char limit. When a line is trimmed, the algorithm can move on to the next line and
ignore all the previous lines in the file, since all of them satisfy all of the given conditions.
The same cannot be said after a line is merged. If we merge the next line with the current line,
there is a chance that the current line exceeds the char limit. After merging the next line with the
current line, the status of the line is 'reset', although we do not need to change the current state
of the algorithm. The process of trimming and merging involves swappings between newline and
whitespace characters. The location of the closest whitespace character, the current character, and
its location are the only pieces of information that the algorithm needs in order to decide whether
to trim or merge.

Given a starting point, we traverse the document character by character, keeping track of the char
number. If we encounter a newline char in a location before the 101th character, then the line has
space to merge.

If the next char is also a newline char, that means that the next line is empty. We can move the
starting point to the char next to the second newline char, if it exists. Otherwise we are at the
end of the file and we are done.

If the next char is not a newline char, then we can replace the current char(a newline) with a
whitespace char and proceed with the algorithm.

If we encounter a char at the 101th location, then the line might exceed the character limit. If the
char is a newline char, then the line is full and satisfies all the conditions. If there is a next
character, then we can move the start point to that char and continue the algorithm as if all
previous lines do not exist.

If the char at the 101th location is a whitespace, then we replace it with a newline. If there is a
next character, then we can move the start point to that char and continue the algorithm as if all
previous lines do not exist.

If the char at the 101th location is not a whitespace or a newline, then the line exceeds the char
limit and it needs to be trimmed at the closest whitespace before the 101th location. It would be
beneficial to keep track of where the last whitespace has occured in order to perform this
operation.

Therfore, the state machine that we require is as follows:

We use variable to keep track of the current char number and another variable to keep track of the
latest whitespace that we have encountered. We are provided a starting point in a string and a char
number of 1. We check the char number. If the char is a whitespace, we save the char number in the
most-recent-whitespace var. If the char number is less than 101 and is not the newline char, we
increase the char number and we repeat. If the char number is ...

