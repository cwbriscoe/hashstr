# hashstr

# Before I wrote this library, I did not know about the `encoding/hex` package.  As you can see by the last benchmark, it is the fastest option and does not require a separate library.  I would recommend using that. 

A faster conversion of []byte to hex string for cryptographic sum functions in golang.

It is over 3x faster than using `Sprintf` on my machine.

### Benchmarks:

```
BenchmarkSprintfSha1-12         12970434                77.6 ns/op           112 B/op          3 allocs/op
BenchmarkSprintfMd5-12          19735413                63.2 ns/op            64 B/op          3 allocs/op
BenchmarkSha1-12                50396240                24.7 ns/op            48 B/op          1 allocs/op
BenchmarkMd5-12                 71154963                17.7 ns/op            32 B/op          1 allocs/op
BenchmarkHexEncoder-12          89067684                16.4 ns/op            32 B/op          1 allocs/op
```

### Example:

```
content := []byte("This is some file content")
sum := sha1.Sum(content)
str := hashstr.Sha1ToLowerString(sum)
fmt.Println(str)
```

### Result:

```
76174c39570ee4e682afb90dc71eb5ff7c70b660
```

### Legal Information

The license is the MIT license. Refer to the `LICENSE` file for more details.

**© 2019 Christopher Briscoe**