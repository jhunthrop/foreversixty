import re,sys,collections,json
data=collections.defaultdict(dict)
for f in sys.argv[1:]:
    for l in open(f):
        m=re.match(r'sv-14266-L(\d+)-(\S+)-r(\d+) .* DPS ([\d.]+)',l)
        if m: data[(int(m.group(1)),m.group(2))][int(m.group(3))]=float(m.group(4))
TH=float(__import__("os").environ.get("TH","0.002")); TOL=float(__import__("os").environ.get("TOL","0.0015"))
H=['orc','troll','tauren','windshaper-skyborne']; A=['human','dwarf','night-elf','high-order-skyborne']
def ok(lv,race):
    v=data[(lv,race)]; top=max(v); t=v[top]; best=max(v.values())
    if best/t-1<TH: return {top}
    return {r for r,d in v.items() if r!=top and d/t-1>=TH and (best-d)/best<=TOL}
levels=sorted(set(lv for lv,_ in data if lv>=24))
assign={}  # (lv,race)->rank
for lv in levels:
    for grp in (H,A):
        races=[r for r in grp if (lv,r) in data]
        if not races: continue
        if len(races)<4:  # orc/human only levels: apply to whole group
            s=ok(lv,races[0]); races_all=grp
        else:
            s=set.intersection(*[ok(lv,r) for r in races]); races_all=races
        if s:
            r=max(s)
            for race in races_all: assign[(lv,race)]=r
        else:
            for race in races:
                assign[(lv,race)]=max(ok(lv,race))
tops={lv:max(data[(lv,r)]) for (lv,r) in data}
for lv in levels:
    print(lv, 'top',tops[lv], ' '.join(f"{r[:5]}:{assign[(lv,r)]}" for r in H+A if (lv,r) in assign))
json.dump({f"{k[0]}|{k[1]}":v for k,v in assign.items()},open(M:='/tmp/x','w')) if False else None
# rows: per level, group races by assigned non-top rank
ALL=H+A
bylevel={}
for lv in levels:
    g=collections.defaultdict(list)
    for r in ALL:
        if (lv,r) in assign and assign[(lv,r)]!=tops[lv]: g[assign[(lv,r)]].append(r)
    bylevel[lv]={k:tuple(v) for k,v in g.items()}
rows=[]
for lv in levels:
    for rank,races in bylevel[lv].items():
        prev=[x for x in rows if x['rank']==rank and x['races']==races and x['to']==lv-1]
        if prev: prev[0]['to']=lv
        else: rows.append({'from':lv,'to':lv,'rank':rank,'races':races})
json.dump(rows,open(sys.argv[1]+'.rows.json','w'))
for r in rows: print('ROW',r['from'],r['to'],'rank',r['rank'],'all' if len(r['races'])==8 else ','.join(r['races']))
