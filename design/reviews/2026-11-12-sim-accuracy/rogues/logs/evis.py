import sys,collections,statistics as st
sys.path.insert(0,'.')
from load import players
P=players(int(sys.argv[1]))
FIN={'Eviscerate','Slice and Dice','Rupture','Kidney Shot','Expose Armor'}
BUILD={'Sinister Strike','Backstab','Gouge','Garrote','Ambush','Cheap Shot','Hemorrhage','Ghostly Strike','Mutilate','Shank'}
pts=collections.defaultdict(list)
for k,p in P.items():
    by=collections.defaultdict(list)
    for e in p['ev']: by[e['encounterId']].append(e)
    for enc,evs in by.items():
        cp=0; tgt=None; lf=None
        for e in evs:
            n=e['spell']['name']
            if e['type']=='cast' and n in BUILD and e['target']['id']!=tgt:
                tgt=e['target']['id']; cp=0
            if e['type']=='resource_gain_happiness': cp=min(5,cp+e['amount'])
            elif e['type']=='cast' and n in FIN:
                ok = e['target']['id']==tgt or n=='Slice and Dice'
                lf=(n,cp if ok else None,e['tMs']); cp=0
            elif e['type']=='damage' and n=='Eviscerate' and lf and lf[0]=='Eviscerate' and lf[1] and 0<=e['tMs']-lf[2]<=50 and not e['isCritical']:
                pts[(k,e['spell']['id'])].append((lf[1],e['amount'],e['target']['name']))
for (k,sid),L in sorted(pts.items()):
    if len(L)<20: continue
    xs=[c for c,a,t in L]; ys=[a for c,a,t in L]
    mx,my=st.mean(xs),st.mean(ys)
    b=sum((x-mx)*(y-my) for x,y,t in L)/sum((x-mx)**2 for x in xs); a=my-b*mx
    d=collections.defaultdict(list)
    for c,am,t in L: d[c].append(am)
    print(k,P[k]['name'],sid,'n',len(L),'slope',round(b,1),'intercept',round(a,1),{c:(len(v),round(st.median(v))) for c,v in sorted(d.items())})
