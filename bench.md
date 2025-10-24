# Benchmarks
```
goos: linux
goarch: amd64
pkg: github.com/pippellia-btc/smallset
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz
```

## Add
```
BenchmarkAdd/size=10/slice_set-4                        66852817                16.93 ns/op            0 B/op          0 allocs/op
BenchmarkAdd/size=10/slice_set_custom-4                 49168137                23.78 ns/op            0 B/op          0 allocs/op
BenchmarkAdd/size=10/map_set-4                          50352103                24.28 ns/op            0 B/op          0 allocs/op

BenchmarkAdd/size=100/slice_set-4                       55642534                20.30 ns/op            0 B/op          0 allocs/op
BenchmarkAdd/size=100/slice_set_custom-4                30528352                37.21 ns/op            0 B/op          0 allocs/op
BenchmarkAdd/size=100/map_set-4                         43945383                25.93 ns/op            0 B/op          0 allocs/op

BenchmarkAdd/size=1000/slice_set-4                      20402848                57.40 ns/op            0 B/op          0 allocs/op
BenchmarkAdd/size=1000/slice_set_custom-4               14062644                83.91 ns/op            0 B/op          0 allocs/op
BenchmarkAdd/size=1000/map_set-4                        37119403                31.35 ns/op            0 B/op          0 allocs/op
```
## Remove
```
BenchmarkRemove/size=10/slice_set-4                     145053924               8.293 ns/op           0 B/op          0 allocs/op
BenchmarkRemove/size=10/slice_set_custom-4              57739856                19.68 ns/op            0 B/op          0 allocs/op
BenchmarkRemove/size=10/map_set-4                       86700064                14.18 ns/op            0 B/op          0 allocs/op

BenchmarkRemove/size=100/slice_set-4                    99090374                11.54 ns/op            0 B/op          0 allocs/op
BenchmarkRemove/size=100/slice_set_custom-4             40757954                28.56 ns/op            0 B/op          0 allocs/op
BenchmarkRemove/size=100/map_set-4                      67044292                16.79 ns/op            0 B/op          0 allocs/op

BenchmarkRemove/size=1000/slice_set-4                   78700677                14.71 ns/op            0 B/op          0 allocs/op
BenchmarkRemove/size=1000/slice_set_custom-4            31307943                36.97 ns/op            0 B/op          0 allocs/op
BenchmarkRemove/size=1000/map_set-4                     80246607                14.09 ns/op            0 B/op          0 allocs/op
```
## Contains
```
BenchmarkContains/size=10/slice_set-4                   171060940                6.946 ns/op           0 B/op          0 allocs/op
BenchmarkContains/size=10/slice_set_custom-4            64184013                18.06 ns/op            0 B/op          0 allocs/op
BenchmarkContains/size=10/map_set-4                     38921341                28.50 ns/op            8 B/op          1 allocs/op

BenchmarkContains/size=100/slice_set-4                  100000000               10.24 ns/op            0 B/op          0 allocs/op
BenchmarkContains/size=100/slice_set_custom-4           43042246                26.72 ns/op            0 B/op          0 allocs/op
BenchmarkContains/size=100/map_set-4                    35867744                35.07 ns/op            8 B/op          1 allocs/op

BenchmarkContains/size=1000/slice_set-4                 84404246                13.92 ns/op            0 B/op          0 allocs/op
BenchmarkContains/size=1000/slice_set_custom-4          32386256                38.09 ns/op            0 B/op          0 allocs/op
BenchmarkContains/size=1000/map_set-4                   33945051                29.59 ns/op            8 B/op          1 allocs/op
```
## Intersect
```
BenchmarkIntersect/size=10/slice_set-4                  12542593               94.35 ns/op           104 B/op          2 allocs/op
BenchmarkIntersect/size=10/slice_set_custom-4            7760259               148.8 ns/op           112 B/op          2 allocs/op
BenchmarkIntersect/size=10/map_set-4                     2642293               457.4 ns/op           136 B/op          3 allocs/op

BenchmarkIntersect/size=100/slice_set-4                  2824398             422.3 ns/op             920 B/op          2 allocs/op
BenchmarkIntersect/size=100/slice_set_custom-4            898191              1120 ns/op             928 B/op          2 allocs/op
BenchmarkIntersect/size=100/map_set-4                     190958              6062 ns/op            1531 B/op         13 allocs/op

BenchmarkIntersect/size=1000/slice_set-4                  281451              3919 ns/op            8216 B/op          2 allocs/op
BenchmarkIntersect/size=1000/slice_set_custom-4            94188             12684 ns/op            8224 B/op          2 allocs/op
BenchmarkIntersect/size=1000/map_set-4                     18294             66013 ns/op           24024 B/op         40 allocs/op
```
## Union
```
BenchmarkUnion/size=10/slice_set-4                      11148476               104.8 ns/op           104 B/op          2 allocs/op
BenchmarkUnion/size=10/slice_set_custom-4                9169765               129.3 ns/op           112 B/op          2 allocs/op
BenchmarkUnion/size=10/map_set-4                         1814629               667.8 ns/op           298 B/op          4 allocs/op

BenchmarkUnion/size=100/slice_set-4                      2978028               403.2 ns/op           920 B/op          2 allocs/op
BenchmarkUnion/size=100/slice_set_custom-4               1596964               788.8 ns/op           928 B/op          2 allocs/op
BenchmarkUnion/size=100/map_set-4                         190177                6006 ns/op          2506 B/op         14 allocs/op

BenchmarkUnion/size=1000/slice_set-4                      254397                4702 ns/op            8216 B/op          2 allocs/op
BenchmarkUnion/size=1000/slice_set_custom-4               129920                8746 ns/op            8224 B/op          2 allocs/op
BenchmarkUnion/size=1000/map_set-4                         19371               64830 ns/op           34867 B/op         35 allocs/op
```
## Difference
```
BenchmarkDifference/size=10/slice_set-4                 10448563                131.8 ns/op           104 B/op          2 allocs/op
BenchmarkDifference/size=10/slice_set_custom-4           6336138                176.7 ns/op           112 B/op          2 allocs/op
BenchmarkDifference/size=10/map_set-4                    2285030                475.0 ns/op           136 B/op          3 allocs/op

BenchmarkDifference/size=100/slice_set-4                 1890544                598.8 ns/op             920 B/op          2 allocs/op
BenchmarkDifference/size=100/slice_set_custom-4           847808                 1333 ns/op             928 B/op          2 allocs/op
BenchmarkDifference/size=100/map_set-4                    196634                 5979 ns/op            1531 B/op         13 allocs/op

BenchmarkDifference/size=1000/slice_set-4                 182808                 5963 ns/op            8216 B/op          2 allocs/op
BenchmarkDifference/size=1000/slice_set_custom-4           85860                13769 ns/op            8224 B/op          2 allocs/op
BenchmarkDifference/size=1000/map_set-4                    17430                66287 ns/op           24026 B/op         40 allocs/op
```
## SymmetricDifference
```
BenchmarkSymmetricDifference/size=10/slice_set-4                        9436933             120.1 ns/op             184 B/op          2 allocs/op
BenchmarkSymmetricDifference/size=10/slice_set_custom-4                 6550690             192.8 ns/op             192 B/op          2 allocs/op
BenchmarkSymmetricDifference/size=10/map_set-4                          1000000              1006 ns/op             298 B/op          4 allocs/op

BenchmarkSymmetricDifference/size=100/slice_set-4                       2032152             609.2 ns/op            1816 B/op          2 allocs/op
BenchmarkSymmetricDifference/size=100/slice_set_custom-4                 680652              1743 ns/op            1824 B/op          2 allocs/op
BenchmarkSymmetricDifference/size=100/map_set-4                           95606             12262 ns/op            3175 B/op         20 allocs/op

BenchmarkSymmetricDifference/size=1000/slice_set-4                       163326              6380 ns/op           16408 B/op          2 allocs/op
BenchmarkSymmetricDifference/size=1000/slice_set_custom-4                 78106             14553 ns/op           16416 B/op          2 allocs/op
BenchmarkSymmetricDifference/size=1000/map_set-4                           9148            134705 ns/op           47869 B/op         68 allocs/op
```