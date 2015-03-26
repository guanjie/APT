#!/bin/bash
cd /Users/guanjiehe/gocode/src/github.com/guanjie/erictools/web_hello/spiders/financialtimes_sample
scrapy crawl finance -o ~/Desktop/daily.csv
