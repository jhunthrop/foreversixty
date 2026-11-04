import json,glob,os,collections
att=collections.Counter(); casts=collections.Counter(); castsInactive=collections.Counter(); hunters=set()
for f in sorted(glob.glob('full/*.json')):
    cid=int(os.path.basename(f).split('.')[2])
    try: ev=json.load(open(f))['events']
    except Exception: continue
    mine=sorted([e for e in ev if e['source_id']==cid],key=lambda e:(e['timestamp_ms'], e['event_type']!='cast'))
    active=False; bycast=collections.defaultdict(list)
    for e in mine:
        if e['event_type']=='cast': bycast[e['timestamp_ms']].append(e['spell_name'])
    for e in mine:
        if e['spell_id']==6150:
            if e['event_type']=='buff_applied':
                if e['timestamp_ms']==0: active=True; continue
                t=e['timestamp_ms']; names=set()
                for dt in range(t-3,t+1): names|=set(bycast.get(dt,[]))
                att[tuple(sorted(names))]+=1; hunters.add(cid)
                active=True
            elif e['event_type']=='buff_removed': active=False
        elif e['event_type']=='cast':
            casts[e['spell_name']]+=1
            if not active: castsInactive[e['spell_name']]+=1
print('hunters',len(hunters))
print('Quick Shots gains attributed to same-ms casts:'); [print('  ',v,k) for k,v in att.most_common()]
print('casts while Quick Shots down:', {k:v for k,v in castsInactive.most_common(12)})
