```
mkdir tutorialpb
protoc -I=. --go_out=tutorialpb --go_opt=paths=source_relative addressbook.proto
```