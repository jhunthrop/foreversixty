exec(open('common.py').read())
import collections
res=collections.Counter()
for gid,x in s.groupby('source_guid'):
    x=x.sort_values(['rep','fight','time_unix_nano'])
    hits=[];ip=[]
    for r in x.itertuples():
        if r.event=='SWING_DAMAGE' or (r.event=='SPELL_DAMAGE' and r.spell_id in (1752,1757,1758,1759,1760,8621,11293,11294,2098,6760,6761,6762,8623,8624,11299,11300,31016,53,2589,2590,2591,8721,11279,11280,11281,16511,14278,1766,1767,1768,1769,1776,1777,8629,11285,11286)):
            hits.append(r)
        if r.spell_id in IP and r.event in ('SPELL_DAMAGE','SPELL_MISSED'): ip.append(r)
    if len(ip)<20: continue
    # attribute each proc to a landed hit on the same target within 0 ns..1ms before
    import bisect
    keyed=collections.defaultdict(list)
    for h in hits: keyed[(h.rep,h.fight,h.dest_guid)].append(h.time_unix_nano)
    used=set();kind=collections.Counter();tot=collections.Counter()
    for h in hits:
        k='oh' if (h.event=='SWING_DAMAGE' and h.off_hand) else ('mh' if h.event=='SWING_DAMAGE' else 'yellow')
        tot[k]+=1
    for p in ip:
        best=None
        for h in hits:
            if (h.rep,h.fight,h.dest_guid)==(p.rep,p.fight,p.dest_guid) and 0<=p.time_unix_nano-h.time_unix_nano<=50_000_000 and id(h) not in used:
                best=h
        if best is None: kind['none']+=1; continue
        used.add(id(best))
        k='oh' if (best.event=='SWING_DAMAGE' and best.off_hand) else ('mh' if best.event=='SWING_DAMAGE' else 'yellow')
        kind[k]+=1
    print(gid[-8:],x.lvl.iloc[0],'procs',dict(kind),'hits',dict(tot),{k:round(kind[k]/tot[k],3) for k in tot if tot[k]})
