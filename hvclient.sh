#!/bin/bash
hvclient -claimsubmit=example.com.com
# hvclient -trustchain
# openssl genrsa 2048 > test.key
# hvclient -privatekey test.key -commonname Demo_cert -csrout > csr.pem
# openssl req -text -noout -in csr.pem
# hvclient -claimdns="3AEA9E9966E8774203A20C574B4C5C5D"
# hvclient -commonname Demo_cert -dnsnames example.com -csr csr.pem > cert1.pem
# cat cert.pem
