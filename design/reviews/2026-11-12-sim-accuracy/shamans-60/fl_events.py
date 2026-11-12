import json, sys, os, time, urllib.request, urllib.parse, urllib.error
UA='Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0 Safari/537.36'
OUT='fl'
def get(path, params=None):
    url='https://foreverlogs.gg/api'+path+('?'+urllib.parse.urlencode(params) if params else '')
    for i in range(8):
        try:
            req=urllib.request.Request(url,headers={'User-Agent':UA,'Accept':'application/json','Referer':'https://foreverlogs.gg/'})
            r=json.load(urllib.request.urlopen(req,timeout=120)); time.sleep(0.8); return r
        except urllib.error.HTTPError as e:
            wait=int(e.headers.get('Retry-After') or 0) or 20*(i+1)
            print('retry',e.code,url[:120],wait,file=sys.stderr,flush=True); time.sleep(wait)
        except Exception as e:
            print('retry',url[:120],e,file=sys.stderr,flush=True); time.sleep(10)
    raise SystemExit('failed '+url)
for rid in sys.argv[1:]:
    fn=f'{OUT}/shaman-events-{rid}.json'
    if os.path.exists(fn): print('have',fn); continue
    ro=json.load(open(f'{OUT}/roster-{rid}.json'))
    ids=sorted({p['character_id'] for p in ro if p['class']=='Shaman'})
    events=[]
    for i in range(0,len(ids),40):
        cur=None
        while True:
            prm={'limit':500,'sourceCharacterIds':','.join(map(str,ids[i:i+40])),'streams':'damage,casts,avoidance,resources,buffs'}
            if cur: prm['cursor']=cur
            d=get(f'/reports/{rid}/meta_events',prm); events+=d['events']
            if not d.get('hasMore'): break
            cur=d['nextCursor']
    json.dump({'ids':ids,'events':events},open(fn+'.tmp','w')); os.rename(fn+'.tmp',fn)
    print(rid,len(ids),len(events),flush=True)
