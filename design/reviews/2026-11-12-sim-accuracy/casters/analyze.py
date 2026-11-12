import json,sys,collections,statistics as st
BUFF_IDS={400625:'Heating Up',12654:'Ignite',412538:'Ignite',11129:'Combustion',28682:'Combustion buff',12873:'Improved Scorch',22959:'Fire Vulnerability',400669:'Fingers of Frost',400670:'FoF ind',12579:"Winter's Chill",408248:'Eclipse',16886:"Nature's Grace",16870:'Clearcasting (druid)',12536:'Clearcasting (mage)',24858:'Moonkin Form',400589:'Missile Barrage',400573:'Arcane Blast',12494:'Frostbite',11426:'Ice Barrier',13031:'Ice Barrier',13032:'Ice Barrier',13033:'Ice Barrier',29166:'Innervate'}
for fn in sys.argv[1:]:
    ev=json.load(open(fn)); ev.sort(key=lambda e:(e['encounterId'],e['tMs'],int(e['id'])))
    name=fn.split('-')[-1].replace('.json','')
    print('=====',name,len(ev),'events, encounters',len(set(e['encounterId'] for e in ev)))
    types=collections.Counter(e['type'] for e in ev); print('types',dict(types))
    casts=collections.Counter((e['spell']['id'],e['spell']['name']) for e in ev if e['type']=='cast')
    print('casts:'); [print('  ',k,v) for k,v in casts.most_common(30)]
    dmg=collections.Counter(); hits=collections.Counter(); crits=collections.Counter(); misses=collections.Counter(); vals=collections.defaultdict(list); cvals=collections.defaultdict(list); per=collections.Counter()
    for e in ev:
        if e['type']=='damage':
            k=(e['spell']['id'],e['spell']['name']); dmg[k]+=e['amount'] or 0; hits[k]+=1
            if e.get('isCritical'): crits[k]+=1; cvals[k].append(e['amount'] or 0)
            else: vals[k].append(e['amount'] or 0)
        elif e['type'] in ('miss','resist','dodge','parry','immune','absorb','deflect','evade'):
            misses[(e['spell']['id'],e['spell']['name'])]+=1
    tot=sum(dmg.values()) or 1
    print('damage share:')
    for k,v in dmg.most_common(25):
        nv=vals[k]; cv=cvals[k]
        print('   %-48s %7d %5.1f%%  n %4d crit %4d (%.1f%%) miss %d | noncrit med %s range %s-%s | crit med %s'%(k,v,100*v/tot,hits[k],crits[k],100*crits[k]/max(1,hits[k]),misses.get(k,0),int(st.median(nv)) if nv else '-',min(nv) if nv else '-',max(nv) if nv else '-',int(st.median(cv)) if cv else '-'))
    b=collections.Counter((e['spell']['id'],e['spell']['name'],e['type']) for e in ev if e['type'].startswith('buff') or e['type'].startswith('debuff'))
    print('buffs/debuffs:'); [print('  ',k,v,BUFF_IDS.get(k[0],'')) for k,v in sorted(b.items(),key=lambda x:-x[1])[:40]]
    # sample keys
    if '--keys' in sys.argv[0:1]: pass
