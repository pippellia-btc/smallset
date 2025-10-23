# Benchmarks
goos: linux
goarch: amd64
pkg: github.com/pippellia-btc/smallset
cpu: Intel(R) Core(TM) i5-4690K CPU @ 3.50GHz

## Add

BenchmarkAdd/size=10/slice_set-4                70601794                17.08 ns/op            0 B/op          0 allocs/op
BenchmarkAdd/size=10/map_set-4                  46366634                25.51 ns/op            0 B/op          0 allocs/op

BenchmarkAdd/size=100/slice_set-4               58374398                19.92 ns/op            0 B/op          0 allocs/op
BenchmarkAdd/size=100/map_set-4                 42435817                25.31 ns/op            0 B/op          0 allocs/op

BenchmarkAdd/size=1000/slice_set-4              20369235                55.69 ns/op            0 B/op          0 allocs/op
BenchmarkAdd/size=1000/map_set-4                37806518                29.92 ns/op            0 B/op          0 allocs/op

## Remove

BenchmarkRemove/size=10/slice_set-4             149210940               8.002 ns/op            0 B/op          0 allocs/op
BenchmarkRemove/size=10/map_set-4               85598176                13.42 ns/op            0 B/op          0 allocs/op

BenchmarkRemove/size=100/slice_set-4            99354484                11.12 ns/op            0 B/op          0 allocs/op
BenchmarkRemove/size=100/map_set-4              74732918                15.09 ns/op            0 B/op          0 allocs/op

BenchmarkRemove/size=1000/slice_set-4           78863067                14.28 ns/op            0 B/op          0 allocs/op
BenchmarkRemove/size=1000/map_set-4             85198021                13.62 ns/op            0 B/op          0 allocs/op

## Contains

BenchmarkContains/size=10/slice_set-4           174974682               6.758 ns/op            0 B/op          0 allocs/op
BenchmarkContains/size=10/map_set-4             39767090                31.23 ns/op            8 B/op          1 allocs/op

BenchmarkContains/size=100/slice_set-4          121152996               9.852 ns/op            0 B/op          0 allocs/op
BenchmarkContains/size=100/map_set-4            36737862                30.82 ns/op            8 B/op          1 allocs/op

BenchmarkContains/size=1000/slice_set-4         89497611                12.94 ns/op            0 B/op          0 allocs/op
BenchmarkContains/size=1000/map_set-4           39638977                30.81 ns/op            8 B/op          1 allocs/op

## Intersect

BenchmarkIntersect/size=10/slice_set-4          11503231               118.3 ns/op           104 B/op          2 allocs/op
BenchmarkIntersect/size=10/map_set-4             2399421               510.1 ns/op           136 B/op          3 allocs/op

BenchmarkIntersect/size=100/slice_set-4          2831347               462.7 ns/op           920 B/op          2 allocs/op
BenchmarkIntersect/size=100/map_set-4             191391               6895 ns/op            1530 B/op         13 allocs/op

BenchmarkIntersect/size=1000/slice_set-4          274294               4349 ns/op            8216 B/op         2 allocs/op
BenchmarkIntersect/size=1000/map_set-4             16614               77464 ns/op           24023 B/op        40 allocs/op

## Union

BenchmarkUnion/size=10/slice_set-4               9726332               106.9 ns/op           104 B/op          2 allocs/op
BenchmarkUnion/size=10/map_set-4                 1739325               791.2 ns/op           298 B/op          4 allocs/op

BenchmarkUnion/size=100/slice_set-4              4022416               296.7 ns/op           920 B/op          2 allocs/op
BenchmarkUnion/size=100/map_set-4                 194602              6651 ns/op            2506 B/op         14 allocs/op

BenchmarkUnion/size=1000/slice_set-4              539636              2136 ns/op            8216 B/op          2 allocs/op
BenchmarkUnion/size=1000/map_set-4                 16410             75800 ns/op           34871 B/op         35 allocs/op

## Difference

BenchmarkDifference/size=10/slice_set-4         11171020               112.4 ns/op           104 B/op          2 allocs/op
BenchmarkDifference/size=10/map_set-4            2503006               485.5 ns/op           136 B/op          3 allocs/op

BenchmarkDifference/size=100/slice_set-4         2399856               501.2 ns/op           920 B/op          2 allocs/op
BenchmarkDifference/size=100/map_set-4            189381              6158 ns/op            1531 B/op         13 allocs/op

BenchmarkDifference/size=1000/slice_set-4         260264              4109 ns/op            8216 B/op          2 allocs/op
BenchmarkDifference/size=1000/map_set-4            15559             77662 ns/op           24024 B/op         40 allocs/op

## SymmetricDifference

BenchmarkSymmetricDifference/size=10/slice_set-4                 9396140               140.2 ns/op           184 B/op          2 allocs/op
BenchmarkSymmetricDifference/size=10/map_set-4                   1087268              1066 ns/op             298 B/op          4 allocs/op

BenchmarkSymmetricDifference/size=100/slice_set-4                1994282               600.4 ns/op          1816 B/op          2 allocs/op
BenchmarkSymmetricDifference/size=100/map_set-4                    99508             12511 ns/op            3175 B/op         20 allocs/op

BenchmarkSymmetricDifference/size=1000/slice_set-4                203006              5063 ns/op           16408 B/op          2 allocs/op
BenchmarkSymmetricDifference/size=1000/map_set-4                    7915            150412 ns/op           47870 B/op         68 allocs/op