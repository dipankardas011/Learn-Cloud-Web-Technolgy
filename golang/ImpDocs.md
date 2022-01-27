# GoLang Improtant Differences

[Golang tour](https://go.dev/tour/)
[Introduction Docs](https://go.dev/doc/effective_go#introduction)
## IF Statements
Since if and switch accept an initialization statement, it's common to see one used to set up a local variable.
```go
if err := file.Chmod(0664); err != nil {
  log.Print(err)
  return err
}
```
In the Go libraries, you'll find that when an if statement doesn't flow into the next statement—that is, the body ends in break, continue, goto, or return—the unnecessary else is omitted.
```go
f, err := os.Open(name)
if err != nil {
  return err
}
codeUsing(f)
```
This is an example of a common situation where code must guard against a sequence of error conditions. The code reads well if the successful flow of control runs down the page, eliminating error cases as they arise. Since error cases tend to end in return statements, the resulting code needs no else statements.
```go
f, err := os.Open(name)
if err != nil {
  return err
}
d, err := f.Stat()
if err != nil {
  f.Close()
  return err
}
codeUsing(f, d)
```

## Redeclaration and reassignment

An aside: The last example in the previous section demonstrates a detail of how the := short declaration form works. The declaration that calls os.Open reads,
```go
f, err := os.Open(name)
```
This statement declares two variables, f and err. A few lines later, the call to f.Stat reads,
```go
d, err := f.Stat()
```
which looks as if it declares d and err. Notice, though, that err appears in both statements. This duplication is legal: err is declared by the first statement, but only re-assigned in the second. This means that the call to f.Stat uses the existing err variable declared above, and just gives it a new value.

In a := declaration a variable v may appear even if it has already been declared, provided:

this declaration is in the same scope as the existing declaration of v (if v is already declared in an outer scope, the declaration will create a new variable §),
the corresponding value in the initialization is assignable to v, and
there is at least one other variable that is created by the declaration.

## For
The Go for loop is similar to—but not the same as—C's. It unifies for and while and there is no do-while. There are three forms, only one of which has semicolons.
```go
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }
```
Short declarations make it easy to declare the index variable right in the loop.
```go
sum := 0
for i := 0; i < 10; i++ {
sum += i
}
```
If you're looping over an array, slice, string, or map, or reading from a channel, a range clause can manage the loop.
```go
for key, value := range oldMap {
newMap[key] = value
}
```
If you only need the first item in the range (the key or index), drop the second:
```go
for key := range m {
if key.expired() {
delete(m, key)
}
}
```
If you only need the second item in the range (the value), use the blank identifier, an underscore, to discard the first:
```go
sum := 0
for _, value := range array {
sum += value
}
```
The blank identifier has many uses, as described in a later section.

For strings, the range does more work for you, breaking out individual Unicode code points by parsing the UTF-8. Erroneous encodings consume one byte and produce the replacement rune U+FFFD. (The name (with associated builtin type) rune is Go terminology for a single Unicode code point. See the language specification for details.) The loop
```go
for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}
```
prints

character U+65E5 '日' starts at byte position 0
character U+672C '本' starts at byte position 3
character U+FFFD '�' starts at byte position 6
character U+8A9E '語' starts at byte position 7
Finally, Go has no comma operator and ++ and -- are statements not expressions. Thus if you want to run multiple variables in a for you should use parallel assignment (although that precludes ++ and --).
```go
// Reverse a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
a[i], a[j] = a[j], a[i]
}
```