#!/bin/sh

su -c 'iptables -I INPUT -p tcp -m tcp --dport 8092 -j ACCEPT && service iptables save'
