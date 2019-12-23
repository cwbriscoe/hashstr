# hashstr

A faster conversion of []byte to hex string for cryptographic sum functions in golang.

It is over 3x faster than using `Sprintf` on my machine.

### Benchmarks:

```
BenchmarkSprintfSha1Upper-12            12686316                85.6 ns/op
BenchmarkSprintfSha1Lower-12            14707431                81.0 ns/op
BenchmarkSha1ToUpperString-12           46048754                25.4 ns/op
BenchmarkSha1ToLowerString-12           45884714                25.6 ns/op
BenchmarkSprintfMd5Upper-12             16712253                72.2 ns/op
BenchmarkSprintfMd5Lower-12             17371088                70.6 ns/op
BenchmarkMd5ToUpperString-12            66416130                20.0 ns/op
BenchmarkMd5ToLowerString-12            59609260                19.9 ns/op
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