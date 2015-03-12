# This package will contain the spiders of your Scrapy project
#
# Please refer to the documentation for information on how to create and manage
# your spiders.
from scrapy.spider import BaseSpider
from scrapy.selector import Selector
from items import CraigslistSampleItem


class MySpider(BaseSpider):
    name = "craig"
    allowed_domains = ["craigslist.org"]
    start_urls = ["http://sfbay.craigslist.org/sfc/npo/"]

    def parse(self, response):
        sel = Selector(response)

        # or titles = response.selector.xpath("//p")
        titles = sel.xpath("//p")
        items = []
        for title in titles:
            item = CraigslistSampleItem()
            item["title"] = title.xpath("a/text()").extract()
            item["link"] = title.xpath("a/@href").extract()
            items.append(item)
        return items
