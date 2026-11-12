import sys,re
def load(fn):
    d={}
    for l in open(fn):
        m=re.match(r"(\S+) (\S+) L(\d+) ([\d.]+) ([\d.]+) ([\d.]+)",l.strip())
        if m: d[(m.group(2),int(m.group(3)))]=(float(m.group(4)),float(m.group(5)),float(m.group(6)))
        if l.startswith('ORDER'): print(fn.split('/')[-1],l.strip()[:400])
    return d
a=load(sys.argv[1]); b=load(sys.argv[2])
for k in sorted(a, key=lambda k:(k[0],k[1])):
    if k in b:
        x,y=a[k],b[k]
        print("%-9s L%-3d base %8.2f var %8.2f  %+6.2f%% (err %.2f%%) oom %4.1f -> %4.1f"%(k[0],k[1],x[0],y[0],100*(y[0]/x[0]-1),100*(x[1]+y[1])/x[0],x[2],y[2]))
