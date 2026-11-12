#!/usr/bin/env bash
cd "$(dirname "$0")"
while read -r rid cid name; do
  python3 -I fetch.py events "$rid" "$cid" "$name" 2>>fetch.log
done <<'LIST'
2853 208013 Briana
2853 209354 Ariwen
2853 208213 Weeman
2834 188396 Tirian
2834 186613 Krackzilla
2834 185122 Xyvre
2834 188694 Losemid
2834 185894 Aneon
2834 187526 Euge
2834 185871 Zwicc
LIST
ls -la events-* | awk '{print $5,$9}'
