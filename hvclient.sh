#!/bin/bash
hvclient -claimsubmit=examp3.com
hvclient -generate -publickey="testdata/rsa_pub.key"
hvclient -generate -duration="30d" -privatekey="testdata/rsa_priv.key"
hvclient -generate -notbefore="2021-10-19T14:30:00IST" -notafter="2021-10-29T18:00:00IST" -csr="testdata/request.p10"
hvclient -generate -notbefore="2021-10-18T14:30:00IST" -duration="90d" -privatekey="testdata/ec_priv.key" -gencsr
# hvclient -trustchain
# openssl genrsa 2048 > test.key
# hvclient -privatekey test.key -commonname Demo_cert -csrout > csr.pem
# openssl req -text -noout -in csr.pem
# hvclient -claimdns="3AEA9E9966E8774203A20C574B4C5C5D"
# hvclient -commonname Demo_cert -dnsnames example.com -csr csr.pem > cert1.pem
# cat cert.pem
