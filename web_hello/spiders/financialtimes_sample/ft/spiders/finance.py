# -*- coding: utf-8 -*-
import scrapy
from items import FtItem


class FinanceSpider(scrapy.Spider):
    name = "finance"
    allowed_domains = ["http://www.ft.com/home/us"]
    start_urls = (
        'http://www.ft.com/home/us/',
    )

    def parse(self, response):
        sel = scrapy.Selector(response)

        titles = sel.xpath("//li[@class='ft-list-item']")
        items = []
        for title in titles:
            item = FtItem()

            item["title"] = title.xpath("a/text()").extract()
            items.append(item)
        return items
