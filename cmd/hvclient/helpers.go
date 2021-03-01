/*
Copyright (C) GMO GlobalSign, Inc. 2019 - All Rights Reserved.

Unauthorized copying of this file, via any medium is strictly prohibited.
No distribution/modification of whole or part thereof is allowed.

Proprietary and confidential.
*/

package main

import (
	"encoding/asn1"
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/globalsign/hvclient"
	"github.com/globalsign/hvclient/internal/oids"
)

// getPasswordFromTerminal does exactly what it says on the tin. If confirm
// is true, the user will be prompted to enter the password again to confirm
// it.
func getPasswordFromTerminal(prompt string, confirm bool) (string, error) {
	fmt.Fprintf(os.Stderr, "%s: ", prompt)

	var password []byte
	var err error

	password, err = terminal.ReadPassword(int(syscall.Stdin))
	fmt.Fprintf(os.Stderr, "\n")

	if err != nil {
		return "", err
	}

	if confirm {
		fmt.Fprintf(os.Stderr, "Enter again to confirm: ")

		var confirmation []byte

		confirmation, err = terminal.ReadPassword(int(syscall.Stdin))
		fmt.Fprintf(os.Stderr, "\n")

		if err != nil {
			return "", err
		}

		if string(confirmation) != string(password) {
			return "", errors.New("passwords don't match")
		}
	}

	return string(password), nil
}

// checkOneValue returns true if exactly one of the provided strings is not
// the empty string.
func checkOneValue(s ...string) bool {
	var count = 0

	for _, item := range s {
		if item != "" {
			count++
		}
	}

	return count == 1
}

// checkAllEmpty returns true if all of the provided strings are the empty
// string.
func checkAllEmpty(s ...string) bool {
	for _, item := range s {
		if item != "" {
			return false
		}
	}

	return true
}

// stringToOIDs converts a comma-separated list of string representations
// of OIDs to a slice of asn1.ObjectIdentifier objects.
func stringToOIDs(s string) ([]asn1.ObjectIdentifier, error) {
	var result = []asn1.ObjectIdentifier{}

	for _, stroid := range strings.Split(s, ",") {
		var oid asn1.ObjectIdentifier
		var err error

		if oid, err = oids.StringToOID(stroid); err != nil {
			return nil, err
		}

		result = append(result, oid)
	}

	return result, nil
}

// stringToIPs converts a comma-separated list of string representations of
// IP addresses to a slice of net.IP objects.
func stringToIPs(s string) ([]net.IP, error) {
	var ips []net.IP

	for _, strip := range strings.Split(s, ",") {
		var ip net.IP

		if ip = net.ParseIP(strings.TrimSpace(strip)); ip == nil {
			return nil, fmt.Errorf("invalid IP address: %s", strings.TrimSpace(strip))
		}

		ips = append(ips, ip)
	}

	return ips, nil
}

// stringToURIs converts a comma-separated list of string representations of
// URIs to a slice of *url.URL objects.
func stringToURIs(s string) ([]*url.URL, error) {
	var uris []*url.URL

	for _, struri := range strings.Split(s, ",") {
		var trimmed = strings.TrimSpace(struri)

		if len(trimmed) == 0 {
			return nil, fmt.Errorf("missing URI: %q", s)
		}

		var uri *url.URL
		var err error

		if uri, err = url.Parse(trimmed); err != nil {
			return nil, err
		}

		uris = append(uris, uri)
	}

	return uris, nil
}

// stringToOIDAndStrings converts a comma-separated list of string
// representations of OIDs and string values to a slice of OIDAndString
// objects.
func stringToOIDAndStrings(s string) ([]hvclient.OIDAndString, error) {
	var result []hvclient.OIDAndString

	for _, pair := range strings.Split(s, ",") {
		pair = strings.TrimSpace(pair)

		if len(pair) == 0 {
			return nil, fmt.Errorf("missing OID and value: %q", s)
		}

		var cmps = strings.Split(pair, "=")
		if len(cmps) < 2 {
			return nil, fmt.Errorf("value missing for OID: %q", pair)
		} else if len(cmps) > 2 {
			return nil, fmt.Errorf("extraneous value(s) for OID: %q", pair)
		} else if len(cmps[0]) == 0 {
			return nil, fmt.Errorf("missing OID: %q", pair)
		} else if len(cmps[1]) == 0 {
			return nil, fmt.Errorf("missing value for OID: %q", pair)
		}

		var oid asn1.ObjectIdentifier
		var err error
		if oid, err = oids.StringToOID(cmps[0]); err != nil {
			return nil, fmt.Errorf("invalid OID: %v", err)
		}

		result = append(
			result,
			hvclient.OIDAndString{
				OID:   oid,
				Value: strings.TrimSpace(cmps[1]),
			},
		)
	}

	return result, nil
}