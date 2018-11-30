# Router hardware hash perf testing #

As I am too lazy to set up cross-compilers, I am doing some performance
benchmarking of my home routers using Go. I am assuming performance of Go
implementation is within order of magnitude of C one; typically this is
true.

What is being tested: 4 different packet sizes.

- 'small' packet: 40 byte payload
- 'medium': 200
- 'large': 500
- ~'IPv6 MTU (given tunnels and whatnot)': 1200


## Raw data: My laptop

```
Hardware: Retina Macbook 2018, 2.3GHz Core i5
Tested: 30.11.2018

goos: darwin
goarch: amd64
pkg: github.com/fingon/go-hashperf
BenchmarkSha256/sha256-40-8         	10000000	       207 ns/op	 193.22 MB/s
BenchmarkSha256/sha256-200-8        	 2000000	       688 ns/op	 290.48 MB/s
BenchmarkSha256/sha256-500-8        	 1000000	      1320 ns/op	 378.79 MB/s
BenchmarkSha256/sha256-1200-8       	  500000	      3028 ns/op	 396.27 MB/s
BenchmarkBlake2B/b2b-40-8           	10000000	       159 ns/op	 250.74 MB/s
BenchmarkBlake2B/b2b-200-8          	 5000000	       289 ns/op	 691.66 MB/s
BenchmarkBlake2B/b2b-500-8          	 3000000	       551 ns/op	 906.85 MB/s
BenchmarkBlake2B/b2b-1200-8         	 1000000	      1314 ns/op	 912.70 MB/s
BenchmarkBlake2S/b2s-40-8           	10000000	       126 ns/op	 315.53 MB/s
BenchmarkBlake2S/b2s-200-8          	 5000000	       396 ns/op	 504.22 MB/s
BenchmarkBlake2S/b2s-500-8          	 2000000	       752 ns/op	 664.30 MB/s
BenchmarkBlake2S/b2s-1200-8         	 1000000	      1695ns/op		 707.56 MB/s
```

## Raw data: My CER

```
Hardware: Turris Omnia (2016 high-end router)
/proc/cpuinfo:
model name	: ARMv7 Processor rev 1 (v7l)
1.6GHz dual-core

Tested: 30.11.2018

root@cer:~# /tmp/arm.test -test.bench .
goos: linux
goarch: arm
pkg: github.com/fingon/go-hashperf
BenchmarkSha256/sha256-40-2         	  300000	      4099 ns/op	   9.76 MB/s
BenchmarkSha256/sha256-200-2        	  100000	     13198 ns/op	  15.15 MB/s
BenchmarkSha256/sha256-500-2        	   50000	     25133 ns/op	  19.89 MB/s
BenchmarkSha256/sha256-1200-2       	   30000	     57877 ns/op	  20.73 MB/s
BenchmarkBlake2B/b2b-40-2           	  300000	      5783 ns/op	   6.92 MB/s
BenchmarkBlake2B/b2b-200-2          	  200000	     11076 ns/op	  18.06 MB/s
BenchmarkBlake2B/b2b-500-2          	  100000	     21353 ns/op	  23.42 MB/s
BenchmarkBlake2B/b2b-1200-2         	   30000	     52068 ns/op	  23.05 MB/s
BenchmarkBlake2S/b2s-40-2           	 1000000	      2180 ns/op	  18.35 MB/s
BenchmarkBlake2S/b2s-200-2          	  200000	      7406 ns/op	  27.00 MB/s
BenchmarkBlake2S/b2s-500-2          	  100000	     14303 ns/op	  34.96 MB/s
BenchmarkBlake2S/b2s-1200-2         	   50000	     33578 ns/op	  35.74 MB/s
```

## Raw data: My old CER

```
Hardware: Buffalo WZR-600DHP (2012 mid-end router)
/proc/cpuinfo:
cpu model		: MIPS 24Kc V7.4
isa			: mips1 mips2 mips32r1 mips32r2
450MHz

soft-float :-p

Tested 30.11.2018


root@OpenWrt:/bin# /tmp/linux.mips.softfloat.test -test.bench .
goos: linux
goarch: mips
pkg: github.com/fingon/go-hashperf
BenchmarkSha256/sha256-40         	   50000	     26842 ns/op	   1.49 MB/s
BenchmarkSha256/sha256-200        	   20000	     74550 ns/op	   2.68 MB/s
BenchmarkSha256/sha256-500        	   10000	    138320 ns/op	   3.61 MB/s
BenchmarkSha256/sha256-1200       	    5000	    309335 ns/op	   3.88 MB/s
BenchmarkBlake2B/b2b-40           	   50000	     24243 ns/op	   1.65 MB/s
BenchmarkBlake2B/b2b-200          	   30000	     44980 ns/op	   4.45 MB/s
BenchmarkBlake2B/b2b-500          	   20000	     85033 ns/op	   5.88 MB/s
BenchmarkBlake2B/b2b-1200         	   10000	    203795 ns/op	   5.89 MB/s
BenchmarkBlake2S/b2s-40           	  200000	     10901 ns/op	   3.67 MB/s
BenchmarkBlake2S/b2s-200          	   50000	     35445 ns/op	   5.64 MB/s
BenchmarkBlake2S/b2s-500          	   20000	     68000 ns/op	   7.35 MB/s
BenchmarkBlake2S/b2s-1200         	   10000	    156789 ns/op	   7.65 MB/s
```


