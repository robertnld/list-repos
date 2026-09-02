# Read me

## Introduction

The app has simple and limited functionality to read Git repositories. Git is not required.

## Logic

### Get latest commit message

Read HEAD - Use reference to get latest reference -> hash of object -> get object -> get commit message


## Information

### File format

```regex
^(blob|commit|tree) + whitespace + length_object + \0 + data$
```


## To do

- Support detached HEAD. getHead() returns hash of the commit object.
