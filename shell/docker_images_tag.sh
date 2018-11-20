#!/bin/bash
domain="www.chauncey.com"
function check_ip() {
    IP=$1
    VALID_CHECK=$(echo $IP|awk -F. '$1<=255&&$2<=255&&$3<=255&&$4<=255{print "yes"}')
    if echo $IP|grep -E "^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$">/dev/null; then
        if [ ${VALID_CHECK:-no} == "yes" ]; then
            echo "0"
        else
            echo "1"
        fi
    else
        echo "1"
    fi
}
for line in `docker images |awk '{print $1"?"$2}'|sed '1d' `
do
    imageName=`echo $line |awk -F "?" '{print $1}'`
    imageTag=`echo $line |awk -F "?" '{print $2}'`
    ip=`echo $imageName |awk -F'/' '{print $1}'`
    echo $imageTag
    result=`check_ip $ip`
    if [ $result == "0" ];then
        newImageName=`echo $imageName|sed "s/$ip/$domain/g" `
        docker tag $imageName $newImageName":"$imageTag
        docker rmi $imageName
    fi
done

