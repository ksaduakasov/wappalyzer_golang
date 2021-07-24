## Benchmark results
* Pure wappalyzer api with api key was used, along with other 2 similar native golang approaches
* Even though tests were different and depended on website url that we are looking for, Wappalyzer api showed the slowest results

### Pure WAPPALYZER api
```go
goos: darwin
goarch: amd64
pkg: awesomeProject20/wappalyzer
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkProcess
[{[{google-tag-manager Google Tag Manager [] 24699 1625052312 [{42 tag-managers Tag managers}]} {amp AMP [] 31002 1625052312 [{12 javascript-frameworks JavaScript frameworks}]} {ubuntu Ubuntu [] 31133 1625031462 [{28 operating-systems Operating systems}]} {nginx Nginx [1.18.0] 31133 1625031462 [{22 web-servers Web servers} {64 reverse-proxies Reverse proxies}]} {akamai Akamai [] 31133 1625031462 [{31 cdn CDN}]} {twitter-emoji-twemoji Twitter Emoji (Twemoji) [] 31273 1625052312 [{17 font-scripts Font scripts}]} {disqus Disqus [] 31303 1625052312 [{15 comment-systems Comment systems}]} {buysellads BuySellAds [] 31308 1625052312 [{36 advertising Advertising}]} {typekit Typekit [] 31385 1625052312 [{17 font-scripts Font scripts}]} {google-analytics Google Analytics [] 31389 1625052312 [{10 analytics Analytics}]} {jquery jQuery [1.11.2 1.11.3 3.1.1] 31396 1625052312 [{59 javascript-libraries JavaScript libraries}]} {mysql MySQL [] 31396 1625052312 [{34 databases Databases}]} {wordpress WordPress [4.9.8] 31396 1625052312 [{1 cms CMS} {11 blogs Blogs}]} {jquery-migrate jQuery Migrate [3.0.0] 31396 1625052312 [{59 javascript-libraries JavaScript libraries}]} {google-plus Google Plus [] 31396 1625052312 [{5 widgets Widgets}]} {requirejs RequireJS [2.1.14] 31396 1625052312 [{12 javascript-frameworks JavaScript frameworks}]} {google-sign-in Google Sign-in [] 31396 1625052312 [{69 social-logins Social logins}]} {php PHP [] 31396 1625052312 [{27 programming-languages Programming languages}]}]}]
BenchmarkProcess-8   	       1	2426424923 ns/op
PASS

Process finished with exit code 0
```

### Projectdiscovery/wappalyzergo
```go
goos: darwin
goarch: amd64
pkg: awesomeProject20/projectdiscovery
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkProcess
map[Akamai:{} MySQL:{} Nginx:{} PHP:{} WordPress:{}]
BenchmarkProcess-8   	       1	1543221677 ns/op
PASS

Process finished with exit code 0
```

### unstppbl/gowap
```go
goos: darwin
goarch: amd64
pkg: awesomeProject20/unstppbl
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkProcess
2021/07/24 16:44:21 [*] Result for https://www.geeksforgeeks.org/how-to-print-string-with-double-quotes-in-golang/:
"[{\"categories\":[\"Web servers\",\"Reverse proxies\"],\"name\":\"Nginx\",\"version\":\"\"}]"
BenchmarkProcess-8   	       1	1570434924 ns/op
PASS

Process finished with exit code 0```
