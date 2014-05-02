#!/bin/sh

export LANG=C LANGUAGE=C LC_ALL=C
for x in \
	"cacert_org_class1.pem|http://www.cacert.org/certs/root.crt|13:5C:EC:36:F4:9C:B8:E9:3B:1A:B2:70:CD:80:88:46:76:CE:8F:33" \
	"cacert_org_class3.pem|http://www.cacert.org/certs/class3.crt|AD:7C:3F:64:FC:44:39:FE:F4:E9:0B:E8:F4:7C:6C:FA:8A:AD:FD:CE" \
	; do
	
	f=$(echo "$x" | cut -d'|' -f1)
	url=$(echo "$x" | cut -d'|' -f2)
	sha1=$(echo "$x" | cut -d'|' -f3)

	if [ ! -s "$f" ]; then
		echo "$url => $f"
		curl -s -S -o "$f~" "$url" || exit $?
		if [ "$(openssl x509 -in "$f~" -sha1 -noout -fingerprint)" = "SHA1 Fingerprint=$sha1" ]; then
			mv "$f~" "$f"
		else
			echo "$f: invalid fingerprint" >&2
			rm -f "$f~"
		fi
	fi
done
exec c_rehash .
