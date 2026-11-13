"""Paced foreverlogs fetch: hunters (+ all pets) for the given report ids. Resumable."""
import json, sys, os, time, urllib.request, urllib.parse, urllib.error
UA='Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0 Safari/537.36'
OUT=sys.argv[1]; IDS=[int(x) for x in sys.argv[2:]]
CLASSES={'Hunter'}
def get(path, params=None):
    url='https://foreverlogs.gg/api'+path+('?'+urllib.parse.urlencode(params) if params else '')
    for i in range(12):
        try:
            req=urllib.request.Request(url,headers={'User-Agent':UA,'Accept':'application/json','Referer':'https://foreverlogs.gg/'})
            r=json.load(urllib.request.urlopen(req,timeout=90)); time.sleep(0.7); return r
        except urllib.error.HTTPError as e:
            wait=int(e.headers.get('Retry-After') or 0) or 60*(i+1)
            print('retry',e.code,url[:120],wait,file=sys.stderr,flush=True); time.sleep(wait+5)
        except Exception as e:
            print('retry',url[:120],e,file=sys.stderr,flush=True); time.sleep(15)
    raise SystemExit('failed '+url)
for rid in IDS:
    fn=os.path.join(OUT,f'{rid}.json')
    if os.path.exists(fn): continue
    det=get(f'/reports/{rid}')
    encs=[e['id'] for e in det['encounters']]
    players={}
    for i in range(0,len(encs),60):
        ro=get(f'/reports/{rid}/combatants-roster',{'encounterIds':','.join(map(str,encs[i:i+60]))})
        for p in ro.get('roster',[]):
            players.setdefault(p['character_id'],{'name':p['name'],'class':p['class'],'specs':p.get('specs')})
    ma=get(f'/reports/{rid}/meta_actors')
    pets={}
    for a in ma['actors']:
        if a['isPet']:
            for x in a['ids']: pets[x]={'name':a['name']}
    ids=[k for k,v in players.items() if v['class'] in CLASSES]+list(pets)
    events=[]
    for i in range(0,len(ids),40):
        cur=None
        while True:
            prm={'limit':500,'sourceCharacterIds':','.join(map(str,ids[i:i+40])),'streams':'damage,casts,avoidance,resources,buffs'}
            if cur: prm['cursor']=cur
            d=get(f'/reports/{rid}/meta_events',prm)
            events+=d['events']
            if not d.get('hasMore'): break
            cur=d['nextCursor']
    json.dump({'report':rid,'encounters':det['encounters'],'players':players,'actors':ma['actors'],'pets':pets,'events':events},open(fn+'.tmp','w')); os.rename(fn+'.tmp',fn)
    print(rid,len(players),len(pets),len(events),flush=True)
