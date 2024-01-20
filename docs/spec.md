# Specification

This application opens markdown files and makes sure that all lines in the file do now exceed 100
characters.It also makes sure that there are no trailing whitespaces in between the last word in a
line and the newline character.

Whenever a characted in a word exceeds the 100th character limit then that word must go into the
next line. This way we avoid chopping off words.

