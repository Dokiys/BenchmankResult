# Benchmank
Help you view benchmark result as a real man!

```bash
go get github.com/Dokiys/benchmank@latest
```

```go
func main() {
    results := []testing.BenchmarkResult{
    {N: 303759, T: 1098217418, Bytes: 6890293, MemAllocs: 11846681, MemBytes: 1103259792},
    {N: 103902, T: 1195154370, Bytes: 1093904, MemAllocs: 13091854, MemBytes: 1427216584},
    {N: 28524, T: 1201525659, Bytes: 258498, MemAllocs: 13035570, MemBytes: 1653488272},
    {N: 6651, T: 1187788036, Bytes: 138442, MemAllocs: 12650256, MemBytes: 1752303368},
    {N: 1200, T: 1187804459, Bytes: 68493, MemAllocs: 11994096, MemBytes: 1902036864},
    {N: 183, T: 1198201708, Bytes: 13432, MemAllocs: 11134804, MemBytes: 1883085512},
    {N: 25, T: 1200116249, Bytes: 3849, MemAllocs: 10936051, MemBytes: 2031376064},
    {N: 3, T: 1195828333, Bytes: 438, MemAllocs: 10779369, MemBytes: 2129925992},
    {N: 1, T: 3840851542, Bytes: 138, MemAllocs: 33674740, MemBytes: 7124489904}}
    
    for _, bResult := range results {
        fmt.Printf("%s\n", benchmank.String(bResult))
    }
}
```
Output:
```
n: 303,759 	     3.615 µs/op	     3.632 KB/op	           39 allocs/op
n: 103,902 	    11.502 µs/op	    13.736 KB/op	          126 allocs/op
n:  28,524 	    42.123 µs/op	    57.968 KB/op	          457 allocs/op
n:   6,651 	   178.587 µs/op	   263.464 KB/op	        1,902 allocs/op
n:   1,200 	   989.837 µs/op	     1.585 MB/op	        9,995 allocs/op
n:     183 	     6.548 ms/op	    10.290 MB/op	       60,845 allocs/op
n:      25 	    48.005 ms/op	    81.255 MB/op	      437,442 allocs/op
n:       3 	   398.609 ms/op	   709.975 MB/op	    3,593,123 allocs/op
n:       1 	     3.841  s/op	     7.124 GB/op	   33,674,740 allocs/op
```