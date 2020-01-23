# lion

[![godoc.org][godoc-badge]][godoc]

`lion` finds functions which are not tested.

```
$ go test -coverprofile=cover.out pkgname
$ go vet -vettool=`lion` lion.coverprofile=cover.out pkgname
```

<!-- links -->
[godoc]: https://godoc.org/github.com/gostaticanalysis/lion
[godoc-badge]: https://img.shields.io/badge/godoc-reference-4F73B3.svg?style=flat-square&label=%20godoc.org

