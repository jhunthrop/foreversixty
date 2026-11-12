import json, collections, statistics
OVL=set(range(408439,408478))|set(range(408479,408485))
for rid in (2847,2848,2853,2850,2849):
    d=json.load(open(f'fl/shaman-events-{rid}.json'))
    ev=d['events']
    if not ev: print(rid,'no events'); continue
    by=collections.defaultdict(list)
    for e in ev: by[e['source']['id']].append(e)
    for cid,es in sorted(by.items()):
        name=es[0]['source']['name']; spec=es[0]['source']['spec']
        abil=collections.Counter(); sids={}
        for e in es:
            sp=e.get('spell') or {}
            abil[(e['type'], sp.get('name'))]+=1; sids[sp.get('name')]=sp.get('id')
        ss=[e for e in es if (e.get('spell') or {}).get('id')==17364 and e['type']=='cast']
        gaps=[]
        byenc=collections.defaultdict(list)
        for e in ss: byenc[e['encounterId']].append(e['tMs'])
        for enc,ts in byenc.items():
            ts.sort(); gaps+=[(b-a)/1000 for a,b in zip(ts,ts[1:]) if b-a<30000]
        ovl=sum(1 for e in es if (e.get('spell') or {}).get('id') in OVL)
        mael=sum(1 for e in es if 'Maelstrom' in str((e.get('spell') or {}).get('name')))
        top=[f"{k[0][:6]}/{k[1]}:{v}" for k,v in abil.most_common(16)]
        print(f"{rid} {cid} {name} {spec} n={len(es)} SS={len(ss)} ssgap_min={min(gaps) if gaps else None} p10={statistics.quantiles(gaps,n=10)[0] if len(gaps)>10 else None} overload={ovl} maelstrom={mael}")
        print('    ', ', '.join(top))
