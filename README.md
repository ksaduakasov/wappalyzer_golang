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
2021/07/24 16:44:47 [{"url":"https://www.geeksforgeeks.org/how-to-print-string-with-double-quotes-in-golang","technologies":[{"slug":"google-tag-manager","name":"Google Tag Manager","versions":[],"trafficRank":24699,"confirmedAt":1625052312,"categories":[{"id":42,"slug":"tag-managers","name":"Tag managers"}]},{"slug":"amp","name":"AMP","versions":[],"trafficRank":31002,"confirmedAt":1625052312,"categories":[{"id":12,"slug":"javascript-frameworks","name":"JavaScript frameworks"}]},{"slug":"ubuntu","name":"Ubuntu","versions":[],"trafficRank":31133,"confirmedAt":1625031462,"categories":[{"id":28,"slug":"operating-systems","name":"Operating systems"}]},{"slug":"nginx","name":"Nginx","versions":["1.18.0"],"trafficRank":31133,"confirmedAt":1625031462,"categories":[{"id":22,"slug":"web-servers","name":"Web servers"},{"id":64,"slug":"reverse-proxies","name":"Reverse proxies"}]},{"slug":"akamai","name":"Akamai","versions":[],"trafficRank":31133,"confirmedAt":1625031462,"categories":[{"id":31,"slug":"cdn","name":"CDN"}]},{"slug":"twitter-emoji-twemoji","name":"Twitter Emoji (Twemoji)","versions":[],"trafficRank":31273,"confirmedAt":1625052312,"categories":[{"id":17,"slug":"font-scripts","name":"Font scripts"}]},{"slug":"disqus","name":"Disqus","versions":[],"trafficRank":31303,"confirmedAt":1625052312,"categories":[{"id":15,"slug":"comment-systems","name":"Comment systems"}]},{"slug":"buysellads","name":"BuySellAds","versions":[],"trafficRank":31308,"confirmedAt":1625052312,"categories":[{"id":36,"slug":"advertising","name":"Advertising"}]},{"slug":"typekit","name":"Typekit","versions":[],"trafficRank":31385,"confirmedAt":1625052312,"categories":[{"id":17,"slug":"font-scripts","name":"Font scripts"}]},{"slug":"google-analytics","name":"Google Analytics","versions":[],"trafficRank":31389,"confirmedAt":1625052312,"categories":[{"id":10,"slug":"analytics","name":"Analytics"}]},{"slug":"jquery","name":"jQuery","versions":["1.11.2","1.11.3","3.1.1"],"trafficRank":31396,"confirmedAt":1625052312,"categories":[{"id":59,"slug":"javascript-libraries","name":"JavaScript libraries"}]},{"slug":"mysql","name":"MySQL","versions":[],"trafficRank":31396,"confirmedAt":1625052312,"categories":[{"id":34,"slug":"databases","name":"Databases"}]},{"slug":"wordpress","name":"WordPress","versions":["4.9.8"],"trafficRank":31396,"confirmedAt":1625052312,"categories":[{"id":1,"slug":"cms","name":"CMS"},{"id":11,"slug":"blogs","name":"Blogs"}]},{"slug":"jquery-migrate","name":"jQuery Migrate","versions":["3.0.0"],"trafficRank":31396,"confirmedAt":1625052312,"categories":[{"id":59,"slug":"javascript-libraries","name":"JavaScript libraries"}]},{"slug":"google-plus","name":"Google Plus","versions":[],"trafficRank":31396,"confirmedAt":1625052312,"categories":[{"id":5,"slug":"widgets","name":"Widgets"}]},{"slug":"requirejs","name":"RequireJS","versions":["2.1.14"],"trafficRank":31396,"confirmedAt":1625052312,"categories":[{"id":12,"slug":"javascript-frameworks","name":"JavaScript frameworks"}]},{"slug":"google-sign-in","name":"Google Sign-in","versions":[],"trafficRank":31396,"confirmedAt":1625052312,"categories":[{"id":69,"slug":"social-logins","name":"Social logins"}]},{"slug":"php","name":"PHP","versions":[],"trafficRank":31396,"confirmedAt":1625052312,"categories":[{"id":27,"slug":"programming-languages","name":"Programming languages"}]}]}]
BenchmarkProcess-8   	       1	3027784279 ns/op
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
