

## running them

### Gucumber

* [Gucumber](https://github.com/gucumber/gucumber)
* these examples are in the internal features dir:
 ` gucumber `

### GoConvey

* [GoConvey](https://github.com/smartystreets/goconvey/wiki)
* these examples are in the tests dir:
 ` $GOPATH/bin/goconvey -workDir tests/ -port 14900 `

### Ginko

* [Ginko](http://onsi.github.io/ginkgo/)
* [Agoutis](http://agouti.org/)
*  cd into the specs dir
* run: `$GOPATH/bin/ginkgo -v `


## Deps

* [Gucumber](https://github.com/gucumber/gucumber)
* [Ginko]()
* [Agoutis](http://agouti.org/)
* [GoConvey](https://github.com/smartystreets/goconvey/wiki)

## Notes

* Bumped up the Eventually polling and timeout by adding this:

```go
	// waits 60 seconds before throwing element not found
	SetDefaultEventuallyTimeout(60000 * time.Millisecond) 
	// polls ever 10 milliseconds for whatever the SetDefaultEventuallyTimeout is
	SetDefaultEventuallyPollingInterval(10 * time.Millisecond) 
```
