# Convert IIS rewriteMap redirects to cloudflare functions

[![Release Go project](https://github.com/theztd/iis-redirect/actions/workflows/release.yml/badge.svg)](https://github.com/theztd/iis-redirect/actions/workflows/release.yml)

## Overview

You have a xml file rewriteMap.config which is something what IIS uses to configure rewrite rules (It seems to me like a htaccess for IIS, but I am definitely not an IIS expert). And you want to use cloudflare pages for hosting your modern web. This repository contains convertor from the IIS to the [cloudflare pages](https://pages.cloudflare.com/).

**In short:** Converting rewriteMap.config from IIS to cloudflare's [pages function](https://developers.cloudflare.com/pages/platform/functions/). The output from this script is structure generated in given directory.


**I recomend to run it first time out of your code to be sure what this script do, to prevent any accidents...**


## Usage

Parse rewriteMap.config on the **i**nput and generate functions with structure in **o**utput directory ./functions. 
```bash
iis2flare -i rewriteMap.config -o ./functions
```

**Before each run delete content of output directory. This binary does NOT delete anything.**



### Help
```bash
iis2flare -h
```



## Why???

**The long story short:** The [Cloudflare](https://cloudflare.com) is an awesome and I love all their products (allmost same like as the [Hashicorp's](https://www.hashicorp.com/#overview)). But even cloudflare have some annoying limitations and one of them is available amounth of [redirects](https://developers.cloudflare.com/pages/platform/redirects/) offered by "the pages" product (unfortunatelly it is increasable only in Enterprise plan). So even if the pages fullfill our requirements perfectly, we had to look around to choose: an another provider, or hack it somehow. 
:no_mouth:Our webs has hundreds of redirects due some kind of SEO reasons :no_mouth:...


We've found the [pages function](https://developers.cloudflare.com/pages/platform/functions/) as an another automatizable way, how to generate redirects and keep them verzioned in git.


------
I am not a coder and I'm only golang beginner, but it works for us and maybe it could work for you... :wink:
