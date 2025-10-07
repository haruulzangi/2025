#!/bin/sh
echo 0 > /proc/sys/kernel/randomize_va_space

socat TCP-LISTEN:1337,reuseaddr,fork EXEC:/app/challenge,pty,raw,echo=0,stderr
