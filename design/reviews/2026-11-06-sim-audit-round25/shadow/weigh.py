import sys,collections
XP={10:7600,11:8800,12:10100,13:11400,14:12900,15:14400,16:16000,17:17700,18:19400,19:21300,20:23200,21:25200,22:27300,23:29400,24:31700,25:34000,26:36400,27:38900,28:41400,29:44300,30:47400,31:50800,32:54500,33:58600,34:62800,35:67100,36:71600,37:76100,38:80800,39:85700,40:90700,41:95800,42:101000,43:106300,44:111800}
O=['a','b','d','f']
r=collections.defaultdict(dict)
for line in open(sys.argv[1]):
    o,l,b,f,s,d,e=line.rstrip('\n').split('\t'); r[(int(l),f,int(s))][o]=(float(d),e,int(b))
print("| Lvl | Band | Faction | Seed | (a) | (b) | (d) | (f) | b/a | d/a | f/a |\n|---|---|---|---|---|---|---|---|---|---|---|")
tot=collections.defaultdict(lambda: collections.defaultdict(float))
for k in sorted(r):
    l,f,s=k; v=r[k]; a=v["a"][0]
    if "f" not in v: v["f"]=v["b"]
    print(f"| {l} | {v['a'][2]} | {f[0].upper()} | {s} | "+" | ".join(f"{v[o][0]:.1f}" for o in O)+" | "+" | ".join(f"{(v[o][0]/a-1)*100:+.1f}%" for o in O[1:])+" |")
    if l>41: continue
    kills=XP[l]/(45+5*l)
    for o in O:
        tot[(f,s)]['L'+o]+=a/v[o][0]; tot[(f,s)]['K'+o]+=kills*a/v[o][0]
print("\n| Faction | Seed | rule | (b) | (d) | (f) |\n|---|---|---|---|---|---|")
for k,t in sorted(tot.items()):
    for rule,p in (("one weight per level",'L'),("weighted by kills per level",'K')):
        print(f"| {k[0]} | {k[1]} | {rule} | "+" | ".join(f"{(t[p+o]/t[p+'a']-1)*100:+.2f}%" for o in 'bdf')+" |")
