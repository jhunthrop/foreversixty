import sys,re
def load(p):
    d={};order=None
    for l in open(p):
        m=re.match(r'(\S+) (horde|alliance) L(\d+) ([\d.]+) ([\d.]+)',l)
        if m: d[(m.group(2),int(m.group(3)))]=float(m.group(4))
        if l.startswith('ORDER'): order=l.split(' ',2)[2]
    return d,order
a,oa=load(sys.argv[1]); b,ob=load(sys.argv[2])
for k in sorted(a):
    if k in b: print(k[0],k[1],a[k],b[k],f"{(b[k]/a[k]-1)*100:+.2f}%")
import json
if oa and ob:
    oa=json.loads(oa); ob=json.loads(ob)
    for i,(x,y) in enumerate(zip(oa,ob)):
        if x!=y: print('lvl',10+i,x,'->',y)
