"""demons.py <dl dir>: every warlock demon in the downloaded fights (Pet guids of creature 416 Imp, 1860 Voidwalker,
1863 Succubus, 417 Felhunter), its melee swings and crits, its spells' crits, and its owner's real level floor."""
import glob, sys, collections
import pyarrow.parquet as pq, pandas as pd
KIND = {'416': 'Imp', '1860': 'Voidwalker', '1863': 'Succubus', '417': 'Felhunter'}
dl = sys.argv[1]
sw = collections.defaultdict(lambda: [0, 0]); sp = collections.defaultdict(lambda: [0, 0]); owners = collections.defaultdict(set)
cols = ['event', 'source_guid', 'spell_id', 'spell_name', 'critical', 'miss_type', 'adv_present', 'adv_info_guid', 'adv_owner_guid', 'dest_guid']
for f in glob.glob(f'{dl}/*/f*.parquet'):
    t = pq.read_table(f, columns=cols).to_pandas()
    pets = t[t.source_guid.astype(str).str.startswith(('Pet-', 'Creature-'))]
    for g, grp in pets.groupby('source_guid'):
        parts = str(g).split('-')
        kind = KIND.get(parts[5] if len(parts) > 5 else '')
        if not kind:
            continue
        own = grp[grp.adv_present & (grp.adv_info_guid == g)].adv_owner_guid.dropna()
        if len(own):
            owners[kind].add(own.iloc[0])
        m = grp[grp.event == 'SWING_DAMAGE']
        sw[kind][0] += len(m); sw[kind][1] += int((m.critical == True).sum())
        s = grp[grp.event == 'SPELL_DAMAGE']
        for nm, sg in s.groupby('spell_id'):
            sp[(kind, int(nm))][0] += len(sg); sp[(kind, int(nm))][1] += int((sg.critical == True).sum())
for k, (n, c) in sw.items():
    print(f'{k:10s} swings {n:5d} crits {c:4d} {100 * c / max(n, 1):.1f}%  owners {len(owners[k])}')
for k, (n, c) in sorted(sp.items()):
    print(f'{k[0]:10s} spell {k[1]:7d} hits {n:5d} crits {c:4d} {100 * c / max(n, 1):.1f}%')
