import json, collections, statistics as st
ev=json.load(open('logs/ev_2839_191953.json'))
ev.sort(key=lambda e:(e['encounterId'], e['tMs']))
by_enc=collections.defaultdict(list)
for e in ev: by_enc[e['encounterId']].append(e)
print('encounters', len(by_enc))
name=lambda e:(e.get('spell') or {}).get('name')
# 1. Touch of the Grave
tog=[e for e in ev if name(e)=='Touch of the Grave' and e['type'] in ('damage','immune')]
gaps=[]
for enc,L in by_enc.items():
    t=[e['tMs'] for e in L if name(e)=='Touch of the Grave' and e['type'] in ('damage','immune')]
    gaps+= [ (b-a)/1000 for a,b in zip(t,t[1:])]
print('ToG procs', len(tog), 'gaps min/median', (min(gaps) if gaps else None), (st.median(gaps) if gaps else None), 'n<15s', sum(1 for g in gaps if g<14.9))
print('ToG amounts by encounter (median):', [(enc, st.median([e['amount'] for e in L if name(e)=='Touch of the Grave' and e['type']=='damage'])) for enc,L in by_enc.items() if any(name(e)=='Touch of the Grave' and e['type']=='damage' for e in L)][:12])
# trigger: the closest earlier event by this priest within 200 ms
trig=collections.Counter()
for enc,L in by_enc.items():
    for i,e in enumerate(L):
        if name(e)=='Touch of the Grave' and e['type'] in ('damage','immune'):
            prev=[x for x in L[max(0,i-12):i] if name(x)!='Touch of the Grave' and e['tMs']-x['tMs']<=300 and x['type'] in ('damage','cast','debuff_applied')]
            trig[tuple(sorted({(name(x),x['type']) for x in prev}))]+=1
print('ToG triggers (events within 300 ms before):'); [print('  ',k,v) for k,v in trig.most_common(12)]
# 2. Shoot timing
for enc,L in by_enc.items(): pass
shoot_d=[e for e in ev if name(e)=='Shoot' and e['type'] in ('damage','miss')]
sg=[]; delays=[]
for enc,L in by_enc.items():
    S=[e for e in L if name(e)=='Shoot' and e['type'] in ('damage','miss')]
    ts=[e['tMs'] for e in S]; sg+=[(b-a)/1000 for a,b in zip(ts,ts[1:])]
    casts=[e for e in L if e['type']=='cast' and name(e)!='Shoot']
    for s in S:
        prior=[c for c in casts if c['tMs']<s['tMs']]
        if prior:
            c=prior[-1]; delays.append((name(c), round((s['tMs']-c['tMs'])/1000,2)))
sg.sort()
print('shot-to-shot gaps (s): n',len(sg),'min',sg[:8],'median',st.median(sg) if sg else None)
print('shot gaps histogram', collections.Counter(round(g*2)/2 for g in sg).most_common(10))
d=collections.defaultdict(list)
for n,x in delays: d[n].append(x)
for n,xs in d.items():
    xs.sort(); print('Shoot delay after', n, 'n',len(xs),'min',xs[:6],'median',st.median(xs))
# Shoot cast events: time from Shoot cast to Shoot damage
for enc,L in list(by_enc.items())[:1]: pass
sc=[]
for enc,L in by_enc.items():
    S=[e for e in L if name(e)=='Shoot']
    for i,e in enumerate(S):
        if e['type']=='cast':
            nxt=[x for x in S[i+1:i+3] if x['type'] in ('damage','miss')]
            if nxt: sc.append(round((nxt[0]['tMs']-e['tMs'])/1000,2))
sc.sort(); print('Shoot cast->damage delay n',len(sc),'min',sc[:8],'median',st.median(sc) if sc else None)
# 3. Mind Flay ticks per channel
mf=[]
for enc,L in by_enc.items():
    cur=None
    for e in L:
        if name(e)=='Mind Flay':
            if e['type']=='debuff_applied': 
                if cur is not None: mf.append(cur)
                cur=0
            elif e['type']=='damage' and cur is not None: cur+=1
    if cur is not None: mf.append(cur)
print('Mind Flay ticks per channel', collections.Counter(mf).most_common())
# 4. SWP ticks per application
sw=[]
for enc,L in by_enc.items():
    cur=None
    for e in L:
        if name(e)=='Shadow Word: Pain':
            if e['type'] in ('debuff_applied','debuff_refreshed'):
                if cur is not None: sw.append(cur)
                cur=0
            elif e['type']=='damage' and cur is not None: cur+=1
            elif e['type']=='debuff_removed' and cur is not None: sw.append(cur); cur=None
print('SWP ticks per application', collections.Counter(sw).most_common())
# SWP tick interval
ti=[]
for enc,L in by_enc.items():
    t=[e['tMs'] for e in L if name(e)=='Shadow Word: Pain' and e['type']=='damage']
    ti+=[(b-a)/1000 for a,b in zip(t,t[1:]) if (b-a)<5000]
print('SWP tick interval median', st.median(ti), 'mode', collections.Counter(round(x,1) for x in ti).most_common(3))
# 5. Mind Blast gaps
mb=[]
for enc,L in by_enc.items():
    t=[e['tMs'] for e in L if name(e)=='Mind Blast' and e['type']=='cast']
    mb+=[(b-a)/1000 for a,b in zip(t,t[1:])]
mb.sort(); print('MB cast gaps n',len(mb),'min',mb[:10])
# 6. mana gains
print('mana gains', collections.Counter((name(e), e.get('amount')) for e in ev if e['type']=='resource_gain_mana').most_common(5))
# 7. Mind Flay debuff stacks (slow) & Shadow Vulnerability stacks
print('SV stacks seen', collections.Counter(e.get('amount') for e in ev if name(e)=='Shadow Vulnerability' and e['type'] in ('debuff_applied','debuff_refreshed')))
