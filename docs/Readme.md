# Documentation
Here the concepts is explained briefly for future reference. This folder might require some refactoring.

# Concepts
The ZML is named the `Zettelkasten Manipulation Language`. The language specification is not out yet. However, this project contains the underlying implementation for the operators.

The ZMLE is known as `Zettelkasten Manipulation Language Executor` is the executor engine that performs the operations on the set of documents.

A `Command` is essentially one or a group of `ZMLE Statements`. Which can be nested inside each other. A command can also receive arguments to use in the query.

A `Group` is a text file that contains `Sets` which are composed of `Symbols`. A set of symbols are called `String`. If a `Set` is a `String` that ends with a newline (`\n`) character. In this implementation groups corresponds to files, sets to lines, and symbols to UTF-8 characters.

# ZMLE Statements
They are used to transfer a set of groups into another set of groups.

Here is an example of an simple and a complex query in `ZML` (no context highlighting yet).

```ruby
FROM ("mynotes/mynote.md") AS f
THEN 
    MATCH TAG("unread") AS m 
    REPLACE TAG("read") 
    AFTER 
        TOGROUP
THEN
    SAVE
```

This one replaces the tag unread to read in the file `mynotes/mynote.md`.

Now lets go to a more complex one.

```ruby
# Get the .md files under "mynotes1" and put the matched name into f1. So it will be accessible as f1.notename
FROM ("mynotes1/", {notename}, ".md") AS f1

# Same for the mynotes2.
FROM ("mynotes2/", {notename}, ".md") AS f2

# Then merge 2 groups by the set union operator and call it f.
THEN 
    MERGE f1+f2 AS f 

# Then map the name f to the new filename. Since both groups have 'notename' variable, it is guaranteed to be there. If there is a variable only included in one group, then if it cannot be accessed, it is skipped.
THEN 
    MAP ("mynotes3/", f.notename, ".md")

# Now the group operations are done, we will do set operations

# Then among the sets of those groups try to match this grammar. TAG essentially means a string starting with '#' symbol. One crucial info is that MATCH statement retuns sets not groups.
THEN 
    # Here MATCH takes a tree of operators to match and names it as m.
    MATCH TAG("type/", {ttype}) AS m 
    # WHERE is used to add custom logic.
    WHERE m.ttype in ("note", "meeting", "journal")
    # Then the SELECT keyword selects everyting since it is given without a match expression. 
    SELECT 
    # Then the sets are converted to groups. More detailed info is required about this. 
    TOGROUP

# Save the changes to the files.
THEN 
    SAVE

```

Here the goal of this script is to merge all of the files from two different directories and paste them into the same place.

> So after further consideration my opinion is that this language is too verbose. I know that managing documents needs so much functionality that it is an extremely complex process. However, I think that there can be a solution to this.