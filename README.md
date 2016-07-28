License: WTFPL

Ordinary string:

```go
type face struct {
	Id		 int
	Firstname	 string
	Lastname	 string
	Patronymic	*string

}

[..{..]
	patronymic := ""
	if (f.Patronymic != nil) {
		patronymic = *f.Patronymic
	}
[..}..]
```

Handy string pointer

```go
import (
	h "github.com/xaionaro-go/handyPointers"
)

type face struct {
	Id		 int
	Firstname	 string
	Lastname	 string
	Patronymic	*h.String

}

[..{..]
	patronymic := f.Patronymic.String()
[..}..]
```
