import json, sys, collections, statistics as st
for f in sys.argv[1:]:
    ev=json.load(open(f))
    print('=====', f, len(ev))
    types=collections.Counter(e['type'] for e in ev)
    print('types', dict(types))
    # spells cast / damage by name+rank
    casts=collections.Counter(); dmg=collections.Counter(); crit=collections.Counter(); amounts=collections.defaultdict(list); ticks=collections.Counter()
    for e in ev:
        sp=e.get('spell') or {}
        name=sp.get('name'); rank=sp.get('rank') if 'rank' in sp else None
        key=f"{name} r{rank} #{sp.get('id')}"
        t=e['type']
        if t in ('cast','cast_success','spell_cast'): casts[key]+=1
        if t in ('damage','periodic_damage','spell_damage','damage_periodic'):
            dmg[key]+=1
            if e.get('isCritical'): crit[key]+=1
            if e.get('amount') is not None: amounts[key].append((e['amount'], e.get('isCritical')))
    print('-- casts'); [print('  ',k,v) for k,v in casts.most_common(40)]
    print('-- damage events'); 
    for k,v in dmg.most_common(40):
        a=[x for x,c in amounts[k] if not c]; ac=[x for x,c in amounts[k] if c]
        print('  ',k,v,'crit',crit[k], 'median', (st.median(a) if a else None), 'critmed', (st.median(ac) if ac else None))
    # other streams
    other=collections.Counter((e['type'], (e.get('spell') or {}).get('name')) for e in ev if e['type'] not in ('cast','damage','periodic_damage','spell_cast'))
    print('-- other'); [print('  ',k,v) for k,v in other.most_common(45)]
