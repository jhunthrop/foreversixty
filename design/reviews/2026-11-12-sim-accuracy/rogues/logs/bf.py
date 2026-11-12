import sys,collections,statistics
sys.path.insert(0,'.')
from load import players
P=players(2815)
for k,p in P.items():
    ev=p['ev']
    bfd=[e for e in ev if e['type']=='damage' and e['spell']['id']==1230700]
    if len(bfd)<20: continue
    ratios=[];src=collections.Counter(); crit=collections.Counter()
    idx={}
    for i,e in enumerate(ev):
        if e['type']=='damage' and e['spell']['id']==1230700:
            # nearest preceding damage event same tMs, different target, not BF
            j=i-1
            while j>=0 and ev[j]['tMs']==e['tMs'] and (ev[j]['type']!='damage' or ev[j]['spell']['id']==1230700): j-=1
            if j>=0 and ev[j]['tMs']==e['tMs'] and ev[j]['type']=='damage':
                s=ev[j]; src[s['spell']['name']]+=1
                crit[(s['isCritical'],e['isCritical'])]+=1
                if s['amount']>0: ratios.append(e['amount']/s['amount'])
    # BF windows: applied/removed
    app=[e['tMs'] for e in ev if e['type']=='buff_applied' and e['spell']['id']==13877]
    rem=[e['tMs'] for e in ev if e['type']=='buff_removed' and e['spell']['id']==13877]
    print(k,p['name'],'BF hits',len(bfd),'paired',len(ratios),'median ratio',round(statistics.median(ratios),3) if ratios else None,'q',[round(x,2) for x in statistics.quantiles(ratios,n=10)] if len(ratios)>10 else '', 'src',dict(src.most_common(5)),'crit pairs',dict(crit))
