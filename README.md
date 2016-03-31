# Gratisddns
A simple program to update your ip-address for your gratisdns domain.

## Building it
As with any go program, it is as simple as

    go install

To cross compile it for a raspberry pi, do

    GOOS=linux GOARCH=arm go build

instead

## Using it
You are required to provide gratisdns username, ddns password, one of your domains and a host on these (one of the entries in your A-records). The gratisdns service will automatically register the ip-address the update is coming from, so no need to figure this out yourself.

    gratisddns -u <username> -p <ddns password> -h <domain> -p <host>

You can run this periodically either by adding it to your crontab or by having the program sending an update itself every nth hour by adding ``-s <n>``.
