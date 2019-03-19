# password [![GoDoc](https://godoc.org/github.com/1800alex/go-utilities-password?status.svg)](https://godoc.org/github.com/1800alex/go-utilities-password) [![Build Status](https://travis-ci.com/1800alex/go-utilities-password.png?branch=master)](https://travis-ci.com/1800alex/go-utilities-password)
Package go-utilities-password provides a library for generating high-entropy random password strings via the crypto/rand package.

Download:
```shell
go get github.com/1800alex/go-utilities-password
```

* * *
Package go-utilities-password provides a library for generating high-entropy random
password strings via the crypto/rand package.

forked from github.com/sethvargo/go-password/password

Most functions are safe for concurrent use.





# Examples

Generate
Code:

```
{
	res, err := Generate(64, true, true, false, false)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res)
}
```


Generator Generate
Code:

```
{
	gen, err := NewGenerator(nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := gen.Generate(64, true, true, false, false)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res)
}
```


NewGenerator custom
Code:

```
{
	gen, err := NewGenerator(&GeneratorInput{Symbols: "!@#$%^()"})
	if err != nil {
		log.Fatal(err)
	}
	_ = gen
}
```


NewGenerator nil
Code:

```
{
	gen, err := NewGenerator(nil)
	if err != nil {
		log.Fatal(err)
	}
	_ = gen
}
```



