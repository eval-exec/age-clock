#! /bin/bash
for fontname in `ls /usr/share/figlet` ; do
        echo "found a font :  "${fontname}
        toilet -f ${fontname} -w 20000 123456789.0
done
