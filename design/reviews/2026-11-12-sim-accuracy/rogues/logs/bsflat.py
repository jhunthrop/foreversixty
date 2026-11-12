import sys,collections,statistics as st
sys.path.insert(0,'.')
from load import players
P=players(2815)
def pooled(k,a,b):
    p=P[k]; by=collections.defaultdict(lambda: collections.defaultdict(list))
    for e in p['ev']:
        if e['type']=='damage' and not e['isCritical'] and e['amount']>0: by[e['target']['name']][e['spell']['name']].append(e['amount'])
    num=den=0; rs=[]
    for t,d in by.items():
        if d[a] and d[b]:
            w=min(len(d[a]),len(d[b])); r=st.mean(d[a])/st.mean(d[b]); rs.append((t,len(d[a]),len(d[b]),round(r,3))); num+=w*r; den+=w
    return (round(num/den,3) if den else None), rs
for k,name in [('2818:175636','Katrina'),('2818:175918','Crimson')]:
    for a,b in [('Ambush','Backstab'),('Backstab','Sinister Strike'),('Ambush','Sinister Strike')]:
        r,rs=pooled(k,a,b); print(name,a,'/',b,'pooled',r,rs if len(rs)<10 else rs[:10])
# predictions
def pred(N,Fbs,Famb,Fss,o,c):
    bs=o*(1.5*N+c*Fbs); amb=o*2.5*(N+Famb); ss=N+Fss
    return round(amb/bs,3),round(bs/ss,3),round(amb/ss,3)
for AP in (140,176,200):
    N=16+AP*1.7/14
    print('Katrina AP',AP,'N',round(N,1),'c=1.5',pred(N,32,28,15,1.2,1.5),'c=1',pred(N,32,28,15,1.2,1.0))
