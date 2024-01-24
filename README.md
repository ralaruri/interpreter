## Interpreter in Go

Working through the https://interpreterbook.com/

## Additional Resources on Intepreters. 
- Parsing: https://www.youtube.com/watch?v=bxpc9Pp5pZM


## Chapter 1 Lexing

### Ch 1.1 Lexing Analysis
- Lexer identifies tokens from text and categorizes them.

### Ch 1.2 Defining Our Tokens
- We don't care if a number is 5, 10, 20 etc. we care that is it a nmber and identifies them as such.
- Token Type is a string so we can take advantage and classify anything as a stirng Identifiers, Ints, Commas (,) 
- String is not the most peformant but it's the easiest to work with in the case of this book.

### Ch 1.3 The Lexer
- We formally write out lexer here. We are takeing source input and --> outputting tokens that reperesen tthe source code.
- let x = 5 --> "LET", "IDENT", "EQUAL", "INT"
- We have a testing framework to test the final format of our langauge we are building to test it. 
  

### Ch 1.4 Extending our Token Set and Lexer

### Ch 1.5 Start of a REPL


## Chapter 2 Parsing

### Ch 2.1 Parsers
- Parsers take in input data (text) and create a data strucutre like a abstract syntax tree.

Javascript Example:

```
“> var input = '{"name": "Thorsten", "age": 28}';
> var output = JSON.parse(input);
> output
{ name: 'Thorsten', age: 28 }
> output.name
'Thorsten'
> output.age
28
>
Our”

Excerpt From
Writing An Interpreter In Go
Thorsten Ball
https://itunes.apple.com/WebObjects/MZStore.woa/wa/viewBook?id=0
This material may be protected by copyright.
```
