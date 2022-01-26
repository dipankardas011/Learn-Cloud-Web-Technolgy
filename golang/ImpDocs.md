# GoLang Improtant Differences

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