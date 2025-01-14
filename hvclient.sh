#!/bin/bash
hvclient -claimsubmit=exm.com


# Look for argument, if not set use default
if [[ $1 -eq 0 ]]; then
    domain="exm.com"
else
    domain="${1}"
fi

echo "Creating certificate for ${domain}"

hvclientoutput=`hvclient -claimsubmit=${domain}`

token=$(echo $hvclientoutput | awk 'BEGIN { FS = "," } ; { print $1 }' | awk 'BEGIN { FS = "=" } ; { print $2 }')
claim_id=$(echo $hvclientoutput | awk 'BEGIN { FS = "," } ; { print $3 }')

echo "Token to put into DNS: ${token}"
echo "The claimID: ${claim_id}"

hvclient -claimdns="${claim_id}"
# hvclient -generate -publickey="testdata/rsa_pub.key"
# hvclient -generate -duration="30d" -privatekey="testdata/rsa_priv.key"
# hvclient -generate -notbefore="2021-11-16T10:00:00IST" -notafter="2021-11-16T18:00:00IST" -csr="testdata/request.p10"
# hvclient -generate -notbefore="2021-11-16T10:00:00IST" -duration="90d" -privatekey="testdata/ec_priv.key" -gencsr
# openssl genrsa 2048 > test1.key

openssl genrsa 2048 > test.key
hvclient -privatekey test.key -commonname Demo_cert -csrout > csr.pem
openssl req -text -noout -in csr.pem
hvclient -commonname Demo_cert -csr csr.pem | openssl x509 -text -noout
