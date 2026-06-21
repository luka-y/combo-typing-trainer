The project is tiny (less than 2k loc), load the entire thing into your context,
no need to pick selectively. Note that there is a vendor directory here that you
should ignore, so commands like `git ls-files` will be very bloated.
Ignore everything in the .gitignore.
Also never look at the todo file. If you notice changes to it accidentally, like
while looking at diffs or grep-ing the entire dir for example, ignore.
(but given that there is a vendor dir, grep-ing the entire dir is not the best idea,
try to write your commands in a way that would ignore vendor and everything else that
I asked you to ignore)