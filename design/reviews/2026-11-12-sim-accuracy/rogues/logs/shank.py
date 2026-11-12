import sys,collections,statistics
sys.path.insert(0,'.')
from load import players
P=players(int(sys.argv[1]))
for k,p in P.items():
    ev=p['ev']
    casts=[e for e in ev if e['type']=='cast' and e['spell']['name']=='Shank']
    if len(casts)<10: continue
    en=collections.Counter(e['amount'] for e in ev if e['type']=='resource_gain_energy' and e['spell']['name']=='Shank')
    gaps=[]
    by=collections.defaultdict(list)
    for e in casts: by[e['encounterId']].append(e['tMs'])
    for L in by.values():
        L.sort(); gaps+= [(b-a)/1000 for a,b in zip(L,L[1:])]
    cpS=collections.Counter(e['amount'] for e in ev if e['type']=='resource_gain_happiness' and e['spell']['name'] in('Shank','Boot Knife'))
    dmg=[e for e in ev if e['type']=='damage' and e['spell']['name']=='Shank']
    ip=sum(1 for e in ev if e['type']=='damage' and 'Poison' in e['spell']['name'])
    print(k,p['name'],'casts',len(casts),'energy',dict(en),'min gaps',sorted(gaps)[:6],'cp',dict(cpS),'dmg n',len(dmg),'crit',sum(e['isCritical'] for e in dmg))
