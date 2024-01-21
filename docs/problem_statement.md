# Problem Statement

## What do you want?

I want to keep my markdown files at a 100 char line limit.

## What do you not want?

I don't want to manually adjust words in each line every time that I modify a file.

Right now I have to manually edit each line of a file to make sure it is under 100 chars per line.
When I add or remove content, I have to move words around from the next line to make sure the 100
chars aligment is maintained. The longer the paragraph, the more work I have to do to maintain the
margin of the file.

## What needs to be implemented?

A formater that moves words around to make sure that the margin is respected. Given a file, the
formater will open the file and adjust the placement of newline characters to ensure that the text
follows the char limit.
