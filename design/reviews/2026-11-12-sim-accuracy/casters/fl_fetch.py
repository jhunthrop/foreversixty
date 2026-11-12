import json, os, sys, time, urllib.request, urllib.parse, urllib.error
UA='Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0 Safari/537.36'
OUT=os.path.dirname(os.path.abspath(__file__))
def get(path, params=None):
    url='https://foreverlogs.gg/api'+path+('?'+urllib.parse.urlencode(params) if params else '')
    for i in range(8):
        try:
            req=urllib.request.Request(url,headers={'User-Agent':UA,'Accept':'application/json','Referer':'https://foreverlogs.gg/'})
            d=json.load(urllib.request.urlopen(req,timeout=60)); time.sleep(0.8); return d
        except urllib.error.HTTPError as e:
            wait=int(e.headers.get('Retry-After') or 60) if e.code==429 else 5+i*5
            print('http',e.code,'wait',wait,url,file=sys.stderr,flush=True); time.sleep(min(wait+2,600))
        except Exception as e:
            print('retry',url,e,file=sys.stderr,flush=True); time.sleep(5+i*5)
    raise SystemExit('failed '+url)
cmd=sys.argv[1]
if cmd=='roster':
    for rid in sys.argv[2:]:
        fn=os.path.join(OUT,f'report-{rid}.json')
        if not os.path.exists(fn): json.dump(get(f'/reports/{rid}'),open(fn,'w'))
        det=json.load(open(fn)); encs=[e['id'] for e in det['encounters']]
        fn2=os.path.join(OUT,f'roster-{rid}.json')
        if not os.path.exists(fn2):
            ro=[]
            for i in range(0,len(encs),40):
                ro+=get(f'/reports/{rid}/combatants-roster',{'encounterIds':','.join(map(str,encs[i:i+40]))}).get('roster',[])
            json.dump(ro,open(fn2,'w'))
        for p in json.load(open(fn2)):
            if p['class']=='Mage': print(rid,p['character_id'],p['name'],p.get('specs'),str(p.get('snapshots'))[:200])
elif cmd=='events':
    rid,cid,name=sys.argv[2],sys.argv[3],sys.argv[4]
    fn=os.path.join(OUT,f'events-{rid}-{cid}-{name}.json')
    if os.path.exists(fn): print('have',fn); sys.exit()
    events=[]; cur=None
    while True:
        prm={'limit':500,'sourceCharacterIds':str(cid),'streams':'damage,casts,avoidance,resources,buffs'}
        if cur: prm['cursor']=cur
        d=get(f'/reports/{rid}/meta_events',prm); events+=d['events']
        if not d.get('hasMore'): break
        cur=d['nextCursor']
    json.dump(events,open(fn+'.tmp','w')); os.rename(fn+'.tmp',fn); print(rid,name,len(events))
    fn=os.path.join(OUT,f'report-{rid}.json')
    if not os.path.exists(fn): json.dump(get(f'/reports/{rid}'),open(fn,'w'))
