## Benchmark results
* Pure wappalyzer api with api key was used, along with other 2 similar native golang approaches
* Even though tests were different and depended on website url that we are looking for, Wappalyzer api showed the slowest results
* UPDATED: aside Wappalyzer, I conducted a big research, and finally I found out some pretty decent alternatives for Wappalyer:
    * Builtwith.com
    * Whatcms.com
    * Similartech.com

### Pure WAPPALYZER API
```go
goos: darwin
goarch: amd64
pkg: awesomeProject20/wappalyzer
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkProcess
[
    {
        "url": "https://www.geeksforgeeks.org/how-to-print-string-with-double-quotes-in-golang",
        "technologies": [
            {
                "slug": "google-ads-conversion-tracking",
                "name": "Google Ads Conversion Tracking",
                "versions": [],
                "trafficRank": 0,
                "confirmedAt": 1628689448,
                "categories": [
                    {
                        "id": 10,
                        "slug": "analytics",
                        "name": "Analytics"
                    }
                ]
            },
            {
                "slug": "google-remarketing-tag",
                "name": "Google Remarketing Tag",
                "versions": [],
                "trafficRank": 0,
                "confirmedAt": 1628689448,
                "categories": [
                    {
                        "id": 77,
                        "slug": "retargeting",
                        "name": "Retargeting"
                    }
                ]
            },
            {
                "slug": "http-2",
                "name": "HTTP/2",
                "versions": [],
                "trafficRank": 433,
                "confirmedAt": 1627556558,
                "categories": [
                    {
                        "id": 19,
                        "slug": "miscellaneous",
                        "name": "Miscellaneous"
                    }
                ]
            },
            {
                "slug": "ubuntu",
                "name": "Ubuntu",
                "versions": [],
                "trafficRank": 435,
                "confirmedAt": 1626246256,
                "categories": [
                    {
                        "id": 28,
                        "slug": "operating-systems",
                        "name": "Operating systems"
                    }
                ]
            },
            {
                "slug": "amazon-advertising",
                "name": "Amazon Advertising",
                "versions": [],
                "trafficRank": 485,
                "confirmedAt": 1627703691,
                "categories": [
                    {
                        "id": 36,
                        "slug": "advertising",
                        "name": "Advertising"
                    }
                ]
            },
            {
                "slug": "google-publisher-tag",
                "name": "Google Publisher Tag",
                "versions": [],
                "trafficRank": 486,
                "confirmedAt": 1627703691,
                "categories": [
                    {
                        "id": 36,
                        "slug": "advertising",
                        "name": "Advertising"
                    }
                ]
            },
            {
                "slug": "google-adsense",
                "name": "Google AdSense",
                "versions": [],
                "trafficRank": 512,
                "confirmedAt": 1627703691,
                "categories": [
                    {
                        "id": 36,
                        "slug": "advertising",
                        "name": "Advertising"
                    }
                ]
            },
            {
                "slug": "nginx",
                "name": "Nginx",
                "versions": [
                    "1.18.0"
                ],
                "trafficRank": 704,
                "confirmedAt": 1627703691,
                "categories": [
                    {
                        "id": 22,
                        "slug": "web-servers",
                        "name": "Web servers"
                    },
                    {
                        "id": 64,
                        "slug": "reverse-proxies",
                        "name": "Reverse proxies"
                    }
                ]
            },
            {
                "slug": "akamai",
                "name": "Akamai",
                "versions": [],
                "trafficRank": 704,
                "confirmedAt": 1627703691,
                "categories": [
                    {
                        "id": 31,
                        "slug": "cdn",
                        "name": "CDN"
                    }
                ]
            },
            {
                "slug": "google-tag-manager",
                "name": "Google Tag Manager",
                "versions": [],
                "trafficRank": 722,
                "confirmedAt": 1627703691,
                "categories": [
                    {
                        "id": 42,
                        "slug": "tag-managers",
                        "name": "Tag managers"
                    }
                ]
            },
            {
                "slug": "twitter-emoji-twemoji",
                "name": "Twitter Emoji (Twemoji)",
                "versions": [],
                "trafficRank": 830,
                "confirmedAt": 1627703691,
                "categories": [
                    {
                        "id": 17,
                        "slug": "font-scripts",
                        "name": "Font scripts"
                    }
                ]
            },
            {
                "slug": "amp",
                "name": "AMP",
                "versions": [],
                "trafficRank": 856,
                "confirmedAt": 1627703691,
                "categories": [
                    {
                        "id": 12,
                        "slug": "javascript-frameworks",
                        "name": "JavaScript frameworks"
                    }
                ]
            },
            {
                "slug": "buysellads",
                "name": "BuySellAds",
                "versions": [],
                "trafficRank": 860,
                "confirmedAt": 1627703691,
                "categories": [
                    {
                        "id": 36,
                        "slug": "advertising",
                        "name": "Advertising"
                    }
                ]
            },
            {
                "slug": "typekit",
                "name": "Typekit",
                "versions": [],
                "trafficRank": 860,
                "confirmedAt": 1627703691,
                "categories": [
                    {
                        "id": 17,
                        "slug": "font-scripts",
                        "name": "Font scripts"
                    }
                ]
            },
            {
                "slug": "google-sign-in",
                "name": "Google Sign-in",
                "versions": [],
                "trafficRank": 865,
                "confirmedAt": 1627703691,
                "categories": [
                    {
                        "id": 69,
                        "slug": "social-logins",
                        "name": "Social logins"
                    }
                ]
            },
            {
                "slug": "disqus",
                "name": "Disqus",
                "versions": [],
                "trafficRank": 865,
                "confirmedAt": 1627703691,
                "categories": [
                    {
                        "id": 15,
                        "slug": "comment-systems",
                        "name": "Comment systems"
                    }
                ]
            },
            {
                "slug": "wordpress",
                "name": "WordPress",
                "versions": [
                    "4.9.8"
                ],
                "trafficRank": 865,
                "confirmedAt": 1627703691,
                "categories": [
                    {
                        "id": 1,
                        "slug": "cms",
                        "name": "CMS"
                    },
                    {
                        "id": 11,
                        "slug": "blogs",
                        "name": "Blogs"
                    }
                ]
            },
            {
                "slug": "mysql",
                "name": "MySQL",
                "versions": [],
                "trafficRank": 865,
                "confirmedAt": 1627703691,
                "categories": [
                    {
                        "id": 34,
                        "slug": "databases",
                        "name": "Databases"
                    }
                ]
            },
            {
                "slug": "jquery",
                "name": "jQuery",
                "versions": [
                    "1.11.2",
                    "3.1.1"
                ],
                "trafficRank": 865,
                "confirmedAt": 1627703691,
                "categories": [
                    {
                        "id": 59,
                        "slug": "javascript-libraries",
                        "name": "JavaScript libraries"
                    }
                ]
            },
            {
                "slug": "google-analytics",
                "name": "Google Analytics",
                "versions": [],
                "trafficRank": 865,
                "confirmedAt": 1627703691,
                "categories": [
                    {
                        "id": 10,
                        "slug": "analytics",
                        "name": "Analytics"
                    }
                ]
            },
            {
                "slug": "requirejs",
                "name": "RequireJS",
                "versions": [
                    "2.1.14"
                ],
                "trafficRank": 865,
                "confirmedAt": 1627703691,
                "categories": [
                    {
                        "id": 12,
                        "slug": "javascript-frameworks",
                        "name": "JavaScript frameworks"
                    }
                ]
            },
            {
                "slug": "google-plus",
                "name": "Google Plus",
                "versions": [],
                "trafficRank": 865,
                "confirmedAt": 1627703691,
                "categories": [
                    {
                        "id": 5,
                        "slug": "widgets",
                        "name": "Widgets"
                    }
                ]
            },
            {
                "slug": "php",
                "name": "PHP",
                "versions": [],
                "trafficRank": 865,
                "confirmedAt": 1627703691,
                "categories": [
                    {
                        "id": 27,
                        "slug": "programming-languages",
                        "name": "Programming languages"
                    }
                ]
            },
            {
                "slug": "jquery-migrate",
                "name": "jQuery Migrate",
                "versions": [
                    "3.0.0"
                ],
                "trafficRank": 865,
                "confirmedAt": 1627703691,
                "categories": [
                    {
                        "id": 59,
                        "slug": "javascript-libraries",
                        "name": "JavaScript libraries"
                    }
                ]
            }
        ]
    }
]BenchmarkProcess-8   	       1	2924812979 ns/op
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

Process finished with exit code 0
```

### Pure "builtwith.com" API 
```go
goos: darwin
goarch: amd64
pkg: awesomeProject20/builtWith.com
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkProcess
{
    "Results": [
        {
            "Result": {
                "IsDB": "False",
                "Spend": 0,
                "Paths": [
                    {
                        "Technologies": [
                            {
                                "IsPremium": "no",
                                "Name": "HTML5 DocType",
                                "Description": "The DOCTYPE is a required preamble for HTML5 websites.",
                                "Link": "https://www.w3.org/TR/html51/",
                                "Tag": "docinfo",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "UTF-8",
                                "Description": "UTF-8 (8-bit UCS/Unicode Transformation Format) is a variable-length character encoding for Unicode. It is the preferred encoding for web pages.",
                                "Link": "http://en.wikipedia.org/wiki/UTF-8",
                                "Tag": "encoding",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "Viewport Meta",
                                "Description": "This page uses the viewport meta tag which means the content may be optimized for mobile content.",
                                "Link": "https://developers.google.com/speed/docs/insights/ConfigureViewport",
                                "Tag": "mobile",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "IPhone / Mobile Compatible",
                                "Description": "The website contains code that allows the page to support IPhone / Mobile Content.",
                                "Link": "http://apple.com",
                                "Tag": "mobile",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "Open Graph Protocol",
                                "Description": "The Open Graph protocol enables any web page to become a rich object in a social graph, a open protocol supported by Facebook",
                                "Link": "http://opengraphprotocol.org/",
                                "Tag": "docinfo",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "Javascript Defer",
                                "Description": "The defer attribute gives a hint to the browser that the script does not create any content so the browser can optionally defer interpreting the script. This can improve performance by delaying execution of scripts until after the body content is parsed and rendered.",
                                "Link": "http://www.websiteoptimization.com/speed/tweak/defer/",
                                "Tag": "docinfo",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "Google API",
                                "Description": "The website uses some form of Google APIs to provide interaction with the many API\u0027s Google Providers.",
                                "Link": "http://code.google.com",
                                "Tag": "javascript",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "Categories": [
                                    "Login"
                                ],
                                "IsPremium": "no",
                                "Name": "Google Identity Platform",
                                "Description": "Google Sign-In is a secure authentication system that enables users to sign in with their Google account.",
                                "Link": "https://developers.google.com/identity/",
                                "Tag": "widgets",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "CDN JS",
                                "Description": "CloudFlare\u0027s CDN with popular javascript frameworks available.",
                                "Link": "http://cdnjs.com",
                                "Tag": "cdn",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "Cloudflare JS",
                                "Description": "Loads content from Cloudflare CDN.",
                                "Link": "https://cloudflare.com",
                                "Tag": "cdn",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "RequireJS",
                                "Description": "RequireJS is a JavaScript file and module loader.",
                                "Link": "http://requirejs.org/",
                                "Tag": "javascript",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "Categories": [
                                    "Translation"
                                ],
                                "IsPremium": "no",
                                "Name": "Google Translate Widget",
                                "Description": "Provides a widget which lets a visitor translate your webpage into a different language.",
                                "Link": "http://translate.google.com/translate_tools",
                                "Tag": "widgets",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "Typekit",
                                "Description": "Typekit is the easiest way to use real fonts on the web. It\u0027s a subscription-based service for linking to high-quality Open Type fonts from some of the world\u0027s best type foundries.",
                                "Link": "http://typekit.com",
                                "Tag": "widgets",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "SEO_TITLE",
                                "Description": "SEO Title",
                                "Link": "",
                                "Tag": "seo_title",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "Friends Network",
                                "Description": "XHTML Friends Network is a simple way to represent human relationships using hyperlinks. ",
                                "Link": "http://www.gmpg.org/xfn/",
                                "Tag": "docinfo",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "JSON-LD",
                                "Description": "JSON-LD, or JavaScript Object Notation for Linked Data, is a method of transporting Linked Data using JSON.",
                                "Link": "http://en.wikipedia.org/wiki/JSON-LD",
                                "Tag": "docinfo",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "Categories": [
                                    "Schema"
                                ],
                                "IsPremium": "no",
                                "Name": "Organization Schema",
                                "Description": "Organization i.e. school, NGO, Corporation.",
                                "Link": "https://schema.org/Organization",
                                "Tag": "framework",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "Categories": [
                                    "Schema"
                                ],
                                "IsPremium": "no",
                                "Name": "Person Schema",
                                "Description": "A human being.",
                                "Link": "https://schema.org/Person",
                                "Tag": "framework",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "Javascript",
                                "Description": "JavaScript is a scripting language most often used for client-side web development.",
                                "Link": "http://en.wikipedia.org/wiki/JavaScript",
                                "Tag": "docinfo",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "Meta Description",
                                "Description": "The description attribute provides a concise explanation of the page content.",
                                "Link": "http://en.wikipedia.org/wiki/Meta_tag#The_description_attribute",
                                "Tag": "docinfo",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "Canonical Content Tag",
                                "Description": "Public specification of a preferred URL for a page allows search engines to understand the original location for content.",
                                "Link": "http://googlewebmastercentral.blogspot.com/2009/02/specify-your-canonical.html",
                                "Tag": "docinfo",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "Categories": [
                                    "Open Source",
                                    "Blog"
                                ],
                                "IsPremium": "no",
                                "Name": "WordPress",
                                "Description": "WordPress is a state-of-the-art semantic personal publishing platform with a focus on aesthetics, web standards, and usability.",
                                "Link": "http://wordpress.org",
                                "Tag": "cms",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "Parent": "jQuery",
                                "Categories": [
                                    "JavaScript Library"
                                ],
                                "IsPremium": "no",
                                "Name": "jQuery",
                                "Description": "JQuery is a fast, concise, JavaScript Library that simplifies how you traverse HTML documents, handle events, perform animations, and add Ajax interactions to your web pages. jQuery is designed to change the way that you write JavaScript.",
                                "Link": "http://jquery.com",
                                "Tag": "javascript",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "jQuery CDN",
                                "Description": "The jQuery Amazon S3 Content Delivery Network",
                                "Link": "http://blog.jquery.com/2008/11/19/cloudfront-cdn-for-jquery/",
                                "Tag": "cdn",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "Parent": "jQuery",
                                "IsPremium": "no",
                                "Name": "jQuery 3.1.1",
                                "Description": "jQuery version 3.1.1",
                                "Link": "https://blog.jquery.com/2016/09/22/jquery-3-1-1-released/",
                                "Tag": "javascript",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "Live Writer Support",
                                "Description": "Windows Live Writer Tagging Support Schema",
                                "Link": "http://windowslivewriter.spaces.live.com",
                                "Tag": "feeds",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "oEmbed",
                                "Description": "oEmbed is a format for allowing an embedded representation of a URL on third party sites.",
                                "Link": "http://oembed.com/",
                                "Tag": "docinfo",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "Apple Mobile Web Clips Icon",
                                "Description": "This page contains an icon for iPhone, iPad and iTouch devices.",
                                "Link": "http://developer.apple.com/library/iOS/#documentation/AppleApplications/Reference/SafariWebContent/ConfiguringWebApplications/ConfiguringWebApplications.html",
                                "Tag": "mobile",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "Windows 8 Pinning",
                                "Description": "Windows 8 pinning functionality for websites.",
                                "Link": "http://www.buildmypinnedsite.com/",
                                "Tag": "docinfo",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "Categories": [
                                    "Dynamic Creative Optimization"
                                ],
                                "IsPremium": "no",
                                "Name": "Adpushup",
                                "Description": "Advertising optimization service.",
                                "Link": "http://adpushup.com",
                                "Tag": "ads",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "Categories": [
                                    "Tag Management"
                                ],
                                "IsPremium": "no",
                                "Name": "Google Tag Manager",
                                "Description": "Tag management that lets you add and update website tags without changes to underlying website code.",
                                "Link": "http://google.com/tagmanager",
                                "Tag": "widgets",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "Global Site Tag",
                                "Description": "Google\u0027s primary tag for Google Measurement/Conversion Tracking, Adwords and DoubleClick.",
                                "Link": "https://support.google.com/adwords/answer/7548399",
                                "Tag": "analytics",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "Parent": "Google Conversion Tracking",
                                "Categories": [
                                    "Advertiser Tracking",
                                    "Conversion Tracking"
                                ],
                                "IsPremium": "no",
                                "Name": "Google AdWords Conversion",
                                "Description": "Adwords conversion tracking code.",
                                "Link": "https://support.google.com/adwords/answer/6095821#get_tag",
                                "Tag": "analytics",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "Categories": [
                                    "Conversion Optimization",
                                    "Conversion Tracking"
                                ],
                                "IsPremium": "no",
                                "Name": "Google Conversion Tracking",
                                "Description": "This free tool in AdWords can show you what happens after customers click your ad (for example, whether they purchased your product, called from a mobile phone or downloaded your app).",
                                "Link": "https://support.google.com/adwords/answer/1722054",
                                "Tag": "analytics",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "WAI-ARIA",
                                "Description": "A way to make Web content and Web applications more accessible to people with disabilities. It especially helps with dynamic content and advanced user interface controls developed with Ajax, HTML, JavaScript, and related technologies.",
                                "Link": "http://www.w3.org/WAI/intro/aria",
                                "Tag": "docinfo",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "SEO_H2",
                                "Description": "SEO H2",
                                "Link": "",
                                "Tag": "seo_headers",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "HTML 5 Specific Tags",
                                "Description": "This page contains tags that are specific to an HTML 5 implementation.",
                                "Link": "http://dev.w3.org/html5/html-author/",
                                "Tag": "docinfo",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "SEO_H1",
                                "Description": "SEO H1",
                                "Link": "",
                                "Tag": "seo_headers",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "Parent": "Google Analytics",
                                "IsPremium": "no",
                                "Name": "Google Analytics Classic",
                                "Description": "Classic Google Analytics - sites that are using non-universal analytics code.",
                                "Link": "https://developers.google.com/analytics/devguides/collection/gajs/",
                                "Tag": "analytics",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "Categories": [
                                    "Application Performance",
                                    "Audience Measurement",
                                    "Visitor Count Tracking"
                                ],
                                "IsPremium": "no",
                                "Name": "Google Analytics",
                                "Description": "Google Analytics offers a host of compelling features and benefits for everyone from senior executives and advertising and marketing professionals to site owners and content developers.",
                                "Link": "http://google.com/analytics",
                                "Tag": "analytics",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "Categories": [
                                    "Contextual Advertising"
                                ],
                                "IsPremium": "no",
                                "Name": "Google Adsense",
                                "Description": "A contextual advertising solution for delivering Google AdWords ads that are relevant to site content pages.",
                                "Link": "http://google.com/adsense",
                                "Tag": "ads",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "Parent": "Google Adsense",
                                "IsPremium": "no",
                                "Name": "Google Adsense Asynchronous",
                                "Description": "Fully asynchronous version of the AdSense ad code.",
                                "Link": "https://support.google.com/adsense/answer/3221666?hl=en",
                                "Tag": "ads",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "Categories": [
                                    "Ad Blocking"
                                ],
                                "IsPremium": "no",
                                "Name": "AdRecover",
                                "Description": "Monetize adblocked impressions.",
                                "Link": "https://www.adrecover.com/",
                                "Tag": "ads",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "X-Frame-Options",
                                "Description": "The X-Frame-Options HTTP response header can be used to indicate whether or not a browser should be allowed to render a page in a frame or iframe. Sites can use this to avoid clickjacking attacks, by ensuring that their content is not embedded into other sites.",
                                "Link": "https://developer.mozilla.org/en/the_x-frame-options_response_header",
                                "Tag": "docinfo",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "PHP",
                                "Description": "PHP is a widely used general-purpose scripting language that is especially suited for Web development and can be embedded into HTML.",
                                "Link": "http://www.php.net",
                                "Tag": "framework",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "Content Delivery Network",
                                "Description": "This page contains links that give the impression that some of the site contents are stored on a content delivery network.",
                                "Link": "http://en.wikipedia.org/wiki/Content_Delivery_Network",
                                "Tag": "CDN",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            },
                            {
                                "IsPremium": "no",
                                "Name": "nginx",
                                "Description": "nginx [engine x] is a HTTP server and mail proxy server written by Igor Sysoev.",
                                "Link": "http://nginx.net/",
                                "Tag": "Web Server",
                                "FirstDetected": 1629191942561,
                                "LastDetected": 1629191942561
                            }
                        ],
                        "FirstIndexed": 1629191942561,
                        "LastIndexed": 1629191942561,
                        "Domain": "geeksforgeeks.org/how-to-print-string-with-double-quotes-in-golang/",
                        "Url": "how-to-print-string-with-double-quotes-in-golang/",
                        "SubDomain": "geeksforgeeks.org"
                    }
                ]
            },
            "Meta": {
                "Majestic": 0,
                "Umbrella": 0,
                "Vertical": "",
                "Social": null,
                "CompanyName": null,
                "Telephones": null,
                "Emails": [],
                "City": null,
                "State": null,
                "Postcode": null,
                "Country": null,
                "Names": null,
                "ARank": 0,
                "QRank": 0
            },
            "Attributes": {
                "MJRank": 0,
                "MJTLDRank": 0,
                "RefSN": 0,
                "RefIP": 0,
                "Followers": 0,
                "Sitemap": 0,
                "GTMTags": 0,
                "QubitTags": 0,
                "TealiumTags": 0,
                "AdobeTags": 0,
                "CDimensions": 0,
                "CGoals": 0,
                "CMetrics": 0,
                "ProductCount": 0
            },
            "FirstIndexed": 1629191942561,
            "LastIndexed": 1629191942561,
            "Lookup": "geeksforgeeks.org/how-to-print-string-with-double-quotes-in-golang/",
            "SalesRevenue": 0
        }
    ],
    "Errors": []
}BenchmarkProcess-8   	       1	1761254724 ns/op
PASS

Process finished with exit code 0
```

### Pure "whatcms.com" API 
```go
goos: darwin
goarch: amd64
pkg: awesomeProject20/whatcms.com
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkProcess
{
    "request": "https:\/\/whatcms.org\/API\/Tech?key=8291b3042eaf6d82b37f6af754422e479f2e454a7f094e0b8d8e978636806b79f54943&url=https:\/\/www.geeksforgeeks.org\/how-to-print-string-with-double-quotes-in-golang\/",
    "request_web": "https:\/\/whatcms.org\/?s=https%3A%2F%2Fwww.geeksforgeeks.org%2Fhow-to-print-string-with-double-quotes-in-golang%2F",
    "result": {
        "code": 200,
        "msg": "Success"
    },
    "results": [
        {
            "name": "WordPress",
            "version": "4.9.8",
            "categories": [
                "Blog",
                "CMS"
            ],
            "url": "https:\/\/whatcms.org\/c\/WordPress"
        },
        {
            "name": "PHP",
            "categories": [
                "Programming Language"
            ],
            "url": "https:\/\/whatcms.org\/c\/PHP"
        },
        {
            "name": "MySQL",
            "categories": [
                "Database"
            ],
            "url": "https:\/\/whatcms.org\/c\/MySQL"
        },
        {
            "name": "Nginx",
            "categories": [
                "Web Server"
            ],
            "url": "https:\/\/whatcms.org\/c\/Nginx"
        }
    ],
    "meta": {
        "social": [
            {
                "network": "twitter",
                "url": "https:\/\/twitter.com\/geeksforgeeks",
                "profile": "geeksforgeeks"
            },
            {
                "network": "facebook",
                "url": "https:\/\/www.facebook.com\/geeksforgeeks.org\/",
                "profile": "geeksforgeeks.org"
            },
            {
                "network": "instagram",
                "url": "https:\/\/www.instagram.com\/geeks_for_geeks\/",
                "profile": "geeks_for_geeks"
            },
            {
                "network": "linkedin",
                "url": "https:\/\/in.linkedin.com\/in\/sandeep-jain-b3940815",
                "profile": "sandeep-jain-b3940815"

            {
                "network": "linkedin",
                "url": "https:\/\/www.linkedin.com\/company\/1299009",
                "profile": "1299009"
            },
            {
                "network": "linkedin",
                "url": "https:\/\/in.linkedin.com\/company\/geeksforgeeks",
                "profile": "geeksforgeeks"
            }
        ]
    }
}
BenchmarkProcess-8   	       1	1965078107 ns/op
PASS

Process finished with exit code 0
```

### Pure "similartech.com" API 
```go
goos: darwin
goarch: amd64
pkg: awesomeProject20/similartech.com
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkProcess
{
    "technologies": [
        {
            "coverage": 1.0,
            "id": 1261,
            "name": "Amazon Route 53",
            "categories": [
                "Server",
                "Nameserver"
            ],
            "paying": "yes"
        },
        {
            "coverage": 1.0,
            "id": 1230,
            "name": "Amazon Hosting",
            "categories": [
                "Server",
                "Web Hosting"
            ],
            "paying": "yes"
        },
        {
            "coverage": 1.0,
            "id": 4021,
            "name": "Amazon Hosting EU (Frankfurt)",
            "categories": [
                "Server",
                "Web Hosting"
            ],
            "paying": "yes"
        },
        {
            "coverage": 1.0,
            "id": 4600,
            "name": "Amazon",
            "categories": [
                "Information Technology"
            ],
            "paying": "maybe"
        },
        {
            "coverage": 1.0,
            "id": 2363,
            "name": "Yandex Verification",
            "categories": [
                "Analytics"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.03,
            "id": 1,
            "name": "Google Analytics",
            "categories": [
                "Analytics"
            ],
            "paying": "maybe"
        },
        {
            "coverage": 0.03,
            "id": 5505,
            "name": "HubSpot Chat",
            "categories": [
                "Live Chat"
            ],
            "paying": "yes"
        },
        {
            "coverage": 0.19,
            "id": 1474,
            "name": "Theme Color Tag",
            "categories": [
                "Document Standard",
                "Meta Tags"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.16,
            "id": 2685,
            "name": "MasterCard",
            "categories": [
                "ECommerce",
                "Payment Acceptance"
            ],
            "paying": "yes"
        },
        {
            "coverage": 0.19,
            "id": 10,
            "name": "HubSpot",
            "categories": [
                "Marketing",
                "Marketing Automation"
            ],
            "paying": "yes"
        },
        {
            "coverage": 0.16,
            "id": 1227,
            "name": "Google Tag Manager",
            "categories": [
                "JavaScript",
                "Tag Management"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.96,
            "id": 3579,
            "name": "United States Dollar",
            "categories": [
                "ECommerce",
                "Payment Currency"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.16,
            "id": 1368,
            "name": "Meta http-equiv",
            "categories": [
                "Document Standard",
                "Meta Tags"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.16,
            "id": 1561,
            "name": "HTML5 SVG Tag",
            "categories": [
                "Document Standard",
                "HTML Tags"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.16,
            "id": 2270,
            "name": "React JS",
            "categories": [
                "JavaScript"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.16,
            "id": 2469,
            "name": "X-UA-Compatible",
            "categories": [
                "Document Standard"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.03,
            "id": 1831,
            "name": "DoubleClick Advertiser",
            "categories": [
                "Advertising",
                "Advertisers"
            ],
            "paying": "yes"
        },
        {
            "coverage": 1.0,
            "id": 1324,
            "name": "Apple Mobile Tags",
            "categories": [
                "Mobile"
            ],
            "paying": "no"
        },
        {
            "coverage": 1.0,
            "id": 1325,
            "name": "Meta Viewport",
            "categories": [
                "Mobile"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.16,
            "id": 3965,
            "name": "Gatsby",
            "categories": [
                "JavaScript"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.03,
            "id": 239,
            "name": "DoubleClick",
            "categories": [
                "Advertising",
                "Publisher Ad Server"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.16,
            "id": 5492,
            "name": "Google Marketing Platform",
            "categories": [
                "Marketing"
            ],
            "paying": "yes"
        },
        {
            "coverage": 0.19,
            "id": 4854,
            "name": "Cookie Consent",
            "categories": [
                "Privacy"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.16,
            "id": 4601,
            "name": "Google",
            "categories": [
                "Information Technology"
            ],
            "paying": "maybe"
        },
        {
            "coverage": 0.16,
            "id": 2682,
            "name": "Visa",
            "categories": [
                "ECommerce",
                "Payment Acceptance"
            ],
            "paying": "yes"
        },
        {
            "coverage": 0.03,
            "id": 315,
            "name": "Google AdWords Advertiser",
            "categories": [
                "Advertising",
                "Advertisers"
            ],
            "paying": "yes"
        },
        {
            "coverage": 0.03,
            "id": 5885,
            "name": "Hubspot Ads",
            "categories": [
                "Advertising"
            ],
            "paying": "yes"
        },
        {
            "coverage": 1.0,
            "id": 2238,
            "name": "HTTPS",
            "categories": [
                "Document Standard"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.96,
            "id": 2466,
            "name": "HTML5 DocType",
            "categories": [
                "Document Standard"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.8,
            "id": 1408,
            "name": "OpenGraph Locale Entity",
            "categories": [
                "Document Standard",
                "Open Graph Tags"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.8,
            "id": 1409,
            "name": "OpenGraph Type Entity",
            "categories": [
                "Document Standard",
                "Open Graph Tags"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.8,
            "id": 1410,
            "name": "OpenGraph Url Entity",
            "categories": [
                "Document Standard",
                "Open Graph Tags"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.64,
            "id": 1411,
            "name": "OpenGraph Description Entity",
            "categories": [
                "Document Standard",
                "Open Graph Tags"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.8,
            "id": 3204,
            "name": "Yoast SEO",
            "categories": [
                "Widget"
            ],
            "paying": "maybe"
        },
        {
            "coverage": 0.8,
            "id": 2376,
            "name": "RSS",
            "categories": [
                "Blog",
                "Blogging Standards"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.8,
            "id": 2766,
            "name": "LightBox",
            "categories": [
                "JavaScript"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.8,
            "id": 2382,
            "name": "Live Writer Support",
            "categories": [
                "Blog",
                "Blogging Standards"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.8,
            "id": 2448,
            "name": "Meta Robot",
            "categories": [
                "Document Standard",
                "Meta Tags"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.8,
            "id": 2385,
            "name": "Really Simple Discovery",
            "categories": [
                "Blog",
                "Blogging Standards"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.8,
            "id": 2451,
            "name": "Open Graph Protocol",
            "categories": [
                "Document Standard",
                "Open Graph Tags"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.8,
            "id": 5589,
            "name": "UPS",
            "categories": [
                "ECommerce",
                "Shipping"
            ],
            "paying": "yes"
        },
        {
            "coverage": 0.06,
            "id": 1366,
            "name": "Meta Description",
            "categories": [
                "Document Standard",
                "Meta Tags"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.8,
            "id": 1371,
            "name": "Canonical Tag",
            "categories": [
                "Document Standard",
                "Meta Tags"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.8,
            "id": 1372,
            "name": "Meta Rel Home",
            "categories": [
                "Document Standard",
                "Meta Tags"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.8,
            "id": 1122,
            "name": "jQuery",
            "categories": [
                "JavaScript"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.8,
            "id": 1126,
            "name": "UTF-8",
            "categories": [
                "Document Standard",
                "Encoding"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.8,
            "id": 1128,
            "name": "PHP",
            "categories": [
                "Server",
                "Framework"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.8,
            "id": 1148,
            "name": "WordPress",
            "categories": [
                "Content Management System"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.8,
            "id": 1407,
            "name": "OpenGraph Site Name Entity",
            "categories": [
                "Document Standard",
                "Open Graph Tags"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.61,
            "id": 3595,
            "name": "Facebook",
            "categories": [
                "Information Technology"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.61,
            "id": 723,
            "name": "Facebook Social Plugins",
            "categories": [
                "Social"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.67,
            "id": 1405,
            "name": "OpenGraph Image Entity",
            "categories": [
                "Document Standard",
                "Open Graph Tags"
            ],
            "paying": "no"
        },
        {
            "coverage": 0.03,
            "id": 2445,
            "name": "Meta Keywords",
            "categories": [
                "Document Standard",
                "Meta Tags"
            ],
            "paying": "no"
        }
    ],
    "site": "wtotem.com",
    "found": true
}BenchmarkProcess-8   	       1	1643105660 ns/op
PASS

Process finished with exit code 0
```

