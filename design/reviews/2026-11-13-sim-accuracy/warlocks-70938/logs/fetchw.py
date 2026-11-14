import json, sys, os, time, urllib.request, urllib.parse, urllib.error
UA='Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0 Safari/537.36'
OUT=sys.argv[1]; REPS=[int(x) for x in sys.argv[2].split(',')]
def get(path, params=None):
    url='https://foreverlogs.gg/api'+path+('?'+urllib.parse.urlencode(params) if params else '')
    for i in range(10):
        try:
            req=urllib.request.Request(url,headers={'User-Agent':UA,'Accept':'application/json','Referer':'https://foreverlogs.gg/'})
            r=json.load(urllib.request.urlopen(req,timeout=90)); time.sleep(0.8); return r
        except urllib.error.HTTPError as e:
            wait=int(e.headers.get('Retry-After') or 0) or 30*(i+1)
            print('retry',e.code,url[:120],wait,file=sys.stderr,flush=True); time.sleep(wait)
        except Exception as e:
            print('retry',url[:120],e,file=sys.stderr,flush=True); time.sleep(10)
    raise SystemExit('failed '+url)
pub=get('/reports/public',{'limit':100})
for r in [x for x in pub['reports'] if x['id'] in REPS]:
    rid=r['id']; fn=os.path.join(OUT,f'{rid}.json')
    if os.path.exists(fn): continue
    det=get(f'/reports/{rid}')
    encs=[e['id'] for e in det['encounters']]
    players={}; roster=[]
    for i in range(0,len(encs),60):
        ro=get(f'/reports/{rid}/combatants-roster',{'encounterIds':','.join(map(str,encs[i:i+60]))})
        roster+=ro.get('roster',[])
        for p in ro.get('roster',[]):
            if p['class']=='Warlock': players[p['character_id']]=p
    ma=get(f'/reports/{rid}/meta_actors')
    names={p['name'] for p in players.values()}
    pets={}
    for a in ma['actors']:
        if a['isPet'] and (a.get('ownerName') in names or not a.get('ownerName')):
            for x in a['ids']: pets[x]=a
    ids=list(players)+list(pets)
    events=[]
    for i in range(0,len(ids),40):
        cur=None
        while True:
            prm={'limit':500,'sourceCharacterIds':','.join(map(str,ids[i:i+40])),'streams':'damage,casts,resources,buffs,healing'}
            if cur: prm['cursor']=cur
            d=get(f'/reports/{rid}/meta_events',prm)
            events+=d['events']
            if not d.get('hasMore'): break
            cur=d['nextCursor']
    tmp=fn+'.tmp'
    json.dump({'report':r,'encounters':det['encounters'],'players':players,'roster':roster,'pets':pets,'actors':ma['actors'],'events':events},open(tmp,'w')); os.rename(tmp,fn)
    print(rid,len(players),len(pets),len(events),flush=True)
