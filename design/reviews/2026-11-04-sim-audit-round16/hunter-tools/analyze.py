import json, glob, os, collections, statistics
SP='/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad'
D=SP+'/r16/hunter/full'
lvl={}
for f in glob.glob(SP+'/fl/rep/*.roster.json'):
    rid=os.path.basename(f).split('.')[0]
    for c in json.load(open(f))['roster']:
        for s in (c.get('snapshots') or [])[:1]:
            tabs=s['talents']['talent_groups'][str(s['talents']['active_group'])]['tabs']
            lvl[(rid,str(c['character_id']))]=sum(v['points'] for v in tabs.values())+9
hawk=collections.defaultdict(list)   # level -> (dmg,crit)
hawkints=collections.defaultdict(list)
hawk_per_summon=collections.defaultdict(list)
dive=collections.defaultdict(list)
qs=collections.Counter(); qs_ex=[]
denom=collections.Counter()
for f in sorted(glob.glob(D+'/*.json')):
    rid,eid,cid=os.path.basename(f).split('.')[:3]
    try: ev=json.load(open(f))['events']
    except Exception: continue
    L=lvl.get((rid,cid))
    cidi=int(cid)
    hawks=collections.defaultdict(list)
    for e in ev:
        if e.get('source_owner_id')==cidi and e['source_name']=='Hawk' and e['spell_id']==-1:
            hawks[e['source_id']].append(e)
        if e['source_id']==cidi and e['spell_id']==1293241 and e['event_type']=='damage':
            dive[L].append(int(e['amount']))
    for hid,es in hawks.items():
        es.sort(key=lambda e:e['timestamp_ms'])
        hawk_per_summon[L].append((len(es), es[-1]['timestamp_ms']-es[0]['timestamp_ms']))
        for a,b in zip(es,es[1:]): hawkints[L].append(b['timestamp_ms']-a['timestamp_ms'])
        for e in es:
            if e['event_type']=='damage': hawk[L].append((int(e['amount']),e['is_critical']))
    # Quick Shots attribution
    mine=[e for e in ev if e['source_id']==cidi]
    bydt=collections.defaultdict(list)
    for e in mine:
        if e['event_type']=='damage': bydt[e['timestamp_ms']].append(e)
    active=False
    for e in sorted(mine,key=lambda e:e['timestamp_ms']):
        if e['spell_id']==6150 and e['event_type']=='buff_applied':
            t=e['timestamp_ms']
            near=[x for dt in range(t-30,t+1) for x in bydt.get(dt,[])]
            names=sorted(set(x['spell_name'] for x in near))
            qs[tuple(names)]+=1
            if len(qs_ex)<5: qs_ex.append((rid,eid,cid,t,names))
        if e['event_type']=='damage' and e['spell_id'] in (75,2643,14287,14286,14285,14284,14283,14282,14281,3044,1130,20900,20901,20902,20903,20904,19434,1293241):
            denom[e['spell_name']]+=1
print('Quick Shots gains by same-ms hunter damage:'); [print(' ',v,k) for k,v in qs.most_common()]
print('hunter landed shots:',dict(denom))
for L in sorted(k for k in hawk if k):
    d=[a for a,c in hawk[L] if not c]; cr=[a for a,c in hawk[L] if c]
    ints=[i for i in hawkints[L] if i<4000]
    print(f'level~{L}: hawk hits {len(hawk[L])} noncrit mean {statistics.mean(d) if d else 0:.1f} crit% {100*len(cr)/max(1,len(hawk[L])):.1f} median interval {statistics.median(ints) if ints else 0} ms; swings/summon {statistics.mean(n for n,_ in hawk_per_summon[L]):.1f} span {statistics.mean(s for _,s in hawk_per_summon[L])/1000:.1f}s; dive mean {statistics.mean(dive[L]) if dive[L] else 0:.0f} n={len(dive[L])}')
