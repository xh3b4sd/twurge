# twurge

Tweet purger for personal timelines. Twurge uses the Twitter V2 API. The
required endpoints make at least a "Basic" Subscription mandatory. Further note
that tweets can only be deleted one-by-one and that the API endpoints used are
heavily rate-limited. Therefore, twurge caches the user ID and tweet IDs to be
deleted in local files, in order to reduce request count.



### credentials

```
% tree .credential
.credential
├── twurge-acc-key
├── twurge-acc-secret
├── twurge-api-key
└── twurge-api-secret

1 directory, 4 files
```



### automation

```
#!/bin/bash

# This script is executed by a crontab every 5 minutes in order to automatically
# delete old tweets from a twitter user's timeline.
#
#     https://github.com/xh3b4sd/twurge
#

export GOTWI_ACCESS_TOKEN=$(        cat ~/.credential/twurge-acc-key    )
export GOTWI_ACCESS_TOKEN_SECRET=$( cat ~/.credential/twurge-acc-secret )
export GOTWI_API_KEY=$(             cat ~/.credential/twurge-api-key    )
export GOTWI_API_KEY_SECRET=$(      cat ~/.credential/twurge-api-secret )

/Users/xh3b4sd/project/xh3b4sd/twurge/twurge
```



```
export VISUAL=vim
crontab -e
```



```
$ crontab -l
*/5 * * * * /Users/xh3b4sd/.script/twurge.sh
```
