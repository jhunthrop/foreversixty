import json,glob,collections,statistics,sys
S='/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/wl2/fl/data'
MIN=int(sys.argv[1]) if len(sys.argv)>1 else 200
for fn in sorted(glob.glob(S+'/*.json')):
    d=json.load(open(fn)); rid=d['report']['id']
    byname=collections.defaultdict(list)
    for e in d['events']: byname[e['source']['name']].append(e)
    locks={p['name'] for p in d['players'].values()}
    for name in sorted(locks):
        ev=byname.get(name,[])
        casts=collections.Counter((e['spell']['id'],e['spell']['name']) for e in ev if e['type']=='cast' and e['target'] is not None)
        dmg=collections.defaultdict(list)
        for e in ev:
            if e['type']=='damage': dmg[(e['spell']['id'],e['spell']['name'])].append(e)
        tot=sum(e['amount'] for rows in dmg.values() for e in rows)
        if tot<MIN: continue
        # pets owned
        petdmg=collections.defaultdict(int); petn=collections.Counter()
        for pn,pev in byname.items():
            for e in pev:
                if e['type']=='damage' and e['source'].get('isPet') and e['source'].get('ownerName')==name:
                    petdmg[(pn,e['spell']['name'] if e['spell'] else 'melee')]+=e['amount']; petn[(pn,e['spell']['name'] if e['spell'] else 'melee')]+=1
        pt=sum(petdmg.values())
        taps=[e for e in ev if e['type']=='resource_gain_mana' and e['spell'] and e['spell']['name']=='Life Tap']
        encs=sorted({e['encounterId'] for e in ev}); 
        encdur={en['id']:(en['end_time'],en['start_time']) for en in d['encounters']}
        print('==',rid,name,'player',tot,'pet',pt,'share %.1f%%'%(100*pt/(tot+pt)),'encounters',len(encs),'casts',sum(casts.values()),'taps',len(taps), 'tap amounts',sorted(collections.Counter(e['amount'] for e in taps).items())[:6])
        print('  casts:',', '.join('%s(%d) %d'%(n,i,c) for (i,n),c in casts.most_common(14)))
        for sp,rows in sorted(dmg.items(),key=lambda x:-sum(e['amount'] for e in x[1])):
            n=len(rows); crit=sum(1 for e in rows if e['isCritical']); nc=[e['amount'] for e in rows if not e['isCritical']]; cr=[e['amount'] for e in rows if e['isCritical']]
            # tick intervals within same target/encounter
            byt=collections.defaultdict(list)
            for e in rows: byt[(e['encounterId'],e['target']['id'])].append(e['tMs'])
            gaps=[]
            for ts in byt.values():
                ts.sort(); gaps+=[round((b-a)/1000,1) for a,b in zip(ts,ts[1:]) if 0<b-a<20000]
            g=collections.Counter(gaps).most_common(3)
            print('   %-32s n=%4d crit=%4.1f%% nc med %s [%s-%s] crit med %s share %.0f%% gaps %s'%(str(sp)[:32],n,100*crit/n,statistics.median(nc) if nc else '-',min(nc) if nc else '-',max(nc) if nc else '-',statistics.median(cr) if cr else '-',100*sum(e['amount'] for e in rows)/tot,g))
        for k,v in sorted(petdmg.items(),key=lambda x:-x[1])[:6]: print('   pet',k,v,'n',petn[k])
