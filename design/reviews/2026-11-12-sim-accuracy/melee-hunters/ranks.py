import re,sys,collections
rows=collections.OrderedDict()
for l in open(sys.argv[1]):
    m=re.match(r'(\S+)-r(\d+) .* DPS ([\d.]+) \+/- ([\d.]+)(.*)',l)
    if not m: continue
    key,r,d,e,rest=m.group(1),int(m.group(2)),float(m.group(3)),float(m.group(4)),m.group(5)
    oom=re.search(r'out of mana: ([\d.]+)',rest)
    rows.setdefault(key,[]).append((r,d,e,oom.group(1) if oom else ''))
for k,v in rows.items():
    top=max(v,key=lambda x:x[0])
    best=max(v,key=lambda x:x[1])
    s=' | '.join(f"r{r} {d:.1f} ({(d/top[1]-1)*100:+.2f}%)" for r,d,e,o in v)
    print(f"{k:28s} best r{best[0]} {(best[1]/top[1]-1)*100:+.2f}%  : {s}")
