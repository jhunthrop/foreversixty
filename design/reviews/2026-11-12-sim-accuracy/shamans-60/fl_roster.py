import json, sys, os, time, urllib.request, urllib.parse, urllib.error
UA='Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0 Safari/537.36'
OUT='fl'
def get(path, params=None):
    url='https://foreverlogs.gg/api'+path+('?'+urllib.parse.urlencode(params) if params else '')
    for i in range(6):
        try:
            req=urllib.request.Request(url,headers={'User-Agent':UA,'Accept':'application/json','Referer':'https://foreverlogs.gg/'})
            r=json.load(urllib.request.urlopen(req,timeout=90)); time.sleep(0.8); return r
        except urllib.error.HTTPError as e:
            wait=int(e.headers.get('Retry-After') or 0) or 20*(i+1)
            print('retry',e.code,url[:120],wait,file=sys.stderr,flush=True); time.sleep(wait)
        except Exception as e:
            print('retry',url[:120],e,file=sys.stderr,flush=True); time.sleep(10)
    raise SystemExit('failed '+url)
for rid in sys.argv[1:]:
    fn=f'{OUT}/report-{rid}.json'
    if not os.path.exists(fn): json.dump(get(f'/reports/{rid}'),open(fn,'w'))
    det=json.load(open(fn)); encs=[e['id'] for e in det['encounters']]
    fn2=f'{OUT}/roster-{rid}.json'
    if not os.path.exists(fn2):
        ro=[]
        for i in range(0,len(encs),60):
            ro+=get(f'/reports/{rid}/combatants-roster',{'encounterIds':','.join(map(str,encs[i:i+60]))}).get('roster',[])
        json.dump(ro,open(fn2,'w'))
    seen=set()
    for p in json.load(open(fn2)):
        if p['class']=='Shaman' and p['character_id'] not in seen:
            seen.add(p['character_id'])
            print(rid, p['character_id'], p['name'], p.get('specs'), json.dumps(p.get('snapshots'))[:300])
