import sys, time, urllib.request, urllib.parse, os
UA='Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 Chrome/130.0 Safari/537.36'
BUILD='1.60.1.70888'
def get(table, sid, hot=True):
    fn=f'wago/{table}-{sid}-{"hot" if hot else "base"}.csv'
    if os.path.exists(fn): return open(fn).read()
    q={'build':BUILD,'filter[SpellID]':str(sid)}
    if hot: q['useHotfixes']='1'
    url=f'https://wago.tools/db2/{table}/csv?'+urllib.parse.urlencode(q)
    for i in range(5):
        try:
            req=urllib.request.Request(url,headers={'User-Agent':UA})
            t=urllib.request.urlopen(req,timeout=90).read().decode(); time.sleep(1.2)
            open(fn,'w').write(t); return t
        except Exception as e:
            print('retry',url,e,file=sys.stderr); time.sleep(8)
    return ''
tables=sys.argv[1].split(',')
for sid in sys.argv[2:]:
    for t in tables:
        txt=get(t,sid)
        rows=txt.strip().split('\n')
        print(f'## {t} {sid}: {len(rows)-1} rows')
