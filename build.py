#!/usr/bin/env python

import time
import os
import platform

if __name__ == "__main__":
    x = time.localtime()
    y = x[3] * 60 * 60 + x[4] * 60 + x[5]
    nv = "{0}.{1}".format(time.strftime("%y.%m.%d", time.localtime()), y)
    ov = "\"0.0.0\""
    # os.system("sed -i 's/{0}/{1}/g' tcs.go".format(ov, nv))
    print("start build ...")
    r = os.popen('go version')
    gv = r.read().strip().replace("go version ", "")
    pl = "{0}({1})".format(platform.platform(), platform.node())
    r.close()
    outname = "mywebtools"
    # -w -s
    buildcmd = 'go build -ldflags="-X main.version={1} -X \'main.buildDate={2}\' -X \'main.goVersion={3}\' -X \'main.platform={4}\'" -o {0} main.go'.format(
        outname, nv, time.ctime(time.time()), gv, pl)
    print(buildcmd)
    os.system(buildcmd)
    os.system("rm -f mywebtools.tar.gz")
    os.system(
        "tar zcf mywebtools.tar.gz --exclude *.conf -C .. myvps/conf myvps/static myvps/views myvps/mywebtools"
    )
    os.system("rm -f mywebtools")
