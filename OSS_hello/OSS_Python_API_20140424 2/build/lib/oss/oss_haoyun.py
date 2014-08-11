# Haoyunmommy's haoyun.py file, to upload the files into the OSS cloud.

import oss_api

my_oss = oss_api()

print my_oss.get_all_buckets(my_oss)
