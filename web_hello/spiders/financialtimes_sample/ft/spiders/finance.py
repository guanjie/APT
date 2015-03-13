# -*- coding: utf-8 -*-
import scrapy
# import datetime
# import chetEmail
from ft.items import FtItem


class FinanceSpider(scrapy.Spider):
    name = "finance"
    allowed_domains = ["http://www.ft.com/home/us"]
    start_urls = (
        'http://www.ft.com/home/us/',
        'http://www.ft.com/intl/world',
        'http://www.ft.com/intl/companies',
        'http://www.ft.com/intl/markets',
        'http://www.ft.com/intl/global-economy',
    )

    def parse(self, response):
        sel = scrapy.Selector(response)

        # 可以直接用response.xpath("//li")
        titles = sel.xpath("//li[@class='ft-list-item']")
        items = []
        for title in titles:
            item = FtItem()

            item["title"] = title.xpath("a/text()").extract()
            items.append(item)
        print items
        return items
        # sendmail(items)


""" Tried sendmail function below, but it seems will require gmail service"""
# def sendmail(items):
#     SEND_EMAIL = True
#     emailTos = ['humancool@gmail.com']
#
#     # generate an email!
#     text = 'Found ' + str(len(items)) + ' news feeds.'
#     for item in items:
#         text = text + str(item["title"])
#     text = text + '\n This wraps up todays news, Muahh'
#     print text
#
#     if SEND_EMAIL:
#         today = datetime.date.today().strftime("%B %d, %Y")
#         chetEmail.send(emailTos, 'Robot Humancool Result - '+today, text)
