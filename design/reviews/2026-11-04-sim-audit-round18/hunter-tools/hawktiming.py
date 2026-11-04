import json,glob,os,collections,statistics,sys
D=sys.argv[1]
first=[];last=[];counts=[];outcomes=collections.Counter();evtypes=collections.Counter()
tgtlvl=collections.Counter()
for f in sorted(glob.glob(D+'/*.json')):
    rid,eid,cid=os.path.basename(f).split('.')[:3]
    try: ev=json.load(open(f))['events']
    except Exception: continue
    ci=int(cid)
    ev.sort(key=lambda e:e['timestamp_ms'])
    hawks=collections.defaultdict(list)
    for e in ev:
        if e.get('source_owner_id')==ci and e['source_name']=='Hawk':
            hawks[e['source_id']].append(e)
            evtypes[(e['event_type'],e['spell_id'])]+=1
            if e['event_type']=='damage': outcomes['glancing' if e.get('is_glancing') else ('crit' if e['is_critical'] else 'hit')]+=1
    casts=[e for e in ev if e['source_id']==ci and e['spell_id']==1293241 and e['event_type']=='cast']
    for hid,es in hawks.items():
        t0=es[0]['timestamp_ms']
        c=[x['timestamp_ms'] for x in casts if x['timestamp_ms']<=t0]
        if not c: continue
        ct=c[-1]
        first.append(t0-ct); last.append(es[-1]['timestamp_ms']-ct)
        counts.append(len([x for x in es if x['spell_id']==-1]))
print('event types',evtypes.most_common(12))
print('outcomes',outcomes)
print('first swing after cast ms: median',statistics.median(first),'quartiles',statistics.quantiles(first,n=4))
print('last event after cast ms: median',statistics.median(last),'max',max(last),'quartiles',statistics.quantiles(last,n=4))
print('swings/summon dist',sorted(collections.Counter(counts).items()))
