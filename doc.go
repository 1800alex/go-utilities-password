// Package go-utilities-password provides a library for generating high-entropy random
// password strings via the crypto/rand package.
//
// forked from github.com/sethvargo/go-password/password
//
//    res, err := Generate(64, true, true, false, false)
//    if err != nil  {
//      log.Fatal(err)
//    }
//    log.Printf(res)
//
// Most functions are safe for concurrent use.
