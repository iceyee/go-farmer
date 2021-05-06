#!/bin/bash                                                 
# **********************************************************
# *  Author: Farmer                                        *
# *  Mail: iceyee.studio@qq.com                            *
# *  Git: https://github.com/iceyee                        *
# **********************************************************
#                                                           
#farmer:arguments 
#farmer:flags 
#                                                           
case $1 in
    doc)
        mkdir -p /tmp/go-farmer/src
        cp -r * /tmp/go-farmer
        SCREEN_NAME="godoc-develop"
        SCREEN_COMMAND="godoc -http=:8081 -goroot=/tmp/go-farmer \n"
        SCREEN_ID=$(screen -ls | awk "/$SCREEN_NAME/ {print \$1}" | awk -F '.' '{print $1}')
        [ "$SCREEN_ID" ] && kill $SCREEN_ID && sleep 1
        screen -dmS "$SCREEN_NAME" && sleep 1
        screen -S "$SCREEN_NAME" -x -p 0 -X stuff "$SCREEN_COMMAND"
        ;;
    *)
        echo "未识别参数."
        exit 1
        ;;
esac
