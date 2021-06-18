# reread

[![Go Reference](https://pkg.go.dev/badge/github.com/sanggonlee/reread.svg)](https://pkg.go.dev/github.com/sanggonlee/reread)

Read again from `io.Reader`, as many times as you want!

## Usage

```
r, err = reread.With(r, func(r io.Reader) error {
  // Read from r
  _, err = ioutil.ReadAll(r)
  if err != nil {
    return err
  }

  return nil
})

// You can read from r again
bytes, _ := ioutil.ReadAll(r)
```