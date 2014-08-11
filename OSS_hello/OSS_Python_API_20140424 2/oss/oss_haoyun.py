# Haoyunmommy's haoyun.py file, to upload the files into the OSS cloud.

from oss.oss_api import *

my_oss = OssAPI("oss.aliyuncs.com", "cJMXwh8e9bXJmgRQ", "CUKvIlpmW3QId63EeEEVfxwXYQsiR3")

res = my_oss.put_object_from_file("from-terminal", "HarvardStestDot.pdf", "./Harvard.pdf")

print "%s\n%s" % (res.status, res.read())
