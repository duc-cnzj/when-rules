# Rules Plugin for https://github.com/olebedev/when

[![codecov](https://codecov.io/gh/duc-cnzj/when-rules/branch/master/graph/badge.svg?token=LG44ZOD9YU)](https://codecov.io/gh/duc-cnzj/when-rules)

## Rules

- zh

## Installation

```bash
go get -u github.com/olebedev/when
go get -u github.com/duc-cnzj/when-rules
```

## Usage

```go
w := when.New(nil)
w.Add(zh.All...)
w.Add(common.All...)

text := "明晚8点"
r, err := w.Parse(text, time.Now())
if err != nil {
	// an error has occurred
}
if  r == nil {
 	// no matches found
}

fmt.Println(
	"the time",
	r.Time.String(),
	"mentioned in",
	text[r.Index:r.Index+len(r.Text)],
)
```