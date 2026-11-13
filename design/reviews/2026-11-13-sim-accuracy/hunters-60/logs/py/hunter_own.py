"""hunter_own.py <data/own> <min_level> [name-substring]: per-hunter aggregates over every fight of our own parquet logs.
Damage by spell id (hunter + pets/summons owned by it), casts per minute, Auto Shot intervals, crit and miss rates, AP, level."""
import sys, os, glob, json, collections, statistics as st
import pyarrow.parquet as pq
root, minlvl = sys.argv[1], int(sys.argv[2]); want = sys.argv[3] if len(sys.argv) > 3 else ''
DMG = {'SPELL_DAMAGE', 'SPELL_PERIODIC_DAMAGE', 'RANGE_DAMAGE', 'SWING_DAMAGE_LANDED', 'DAMAGE_SHIELD'}
MISS = {'SPELL_MISSED', 'RANGE_MISSED', 'SWING_MISSED', 'SPELL_PERIODIC_MISSED'}
def med(xs): return round(st.median(xs), 3) if xs else None
for rep in sorted(d for d in os.listdir(root) if os.path.exists(f'{root}/{d}/report.json')):
    rj = json.load(open(f'{root}/{rep}/report.json'))
    hunters = {u['guid'] for u in rj['units'] if u.get('class') == 'Hunter'}
    agg = collections.defaultdict(lambda: {'name': '', 'lv': collections.Counter(), 'ap': [], 'ms': 0, 'dmg': collections.Counter(),
                                          'pet': collections.Counter(), 'casts': collections.Counter(), 'hit': collections.Counter(),
                                          'crit': collections.Counter(), 'miss': collections.Counter(), 'gaps': [], 'fights': 0})
    for fp in sorted(glob.glob(f'{root}/{rep}/fights/*/events.parquet')):
        t = pq.read_table(fp).to_pylist(); owner = {}
        for e in t:
            if e['adv_present'] and e['adv_owner_guid'] in hunters and e['adv_info_guid'] == e['source_guid']: owner[e['source_guid']] = e['adv_owner_guid']
            if e['event'] == 'SPELL_SUMMON' and e['source_guid'] in hunters: owner[e['dest_guid']] = e['source_guid']
        seen = collections.defaultdict(list)
        for e in t:
            g = e['source_guid']; h = g if g in hunters else owner.get(g)
            if not h: continue
            a = agg[h]; key = (e['spell_id'] or 0, e['spell_name'] or 'Melee')
            if g == h:
                a['name'] = e['source_name']; seen[h].append(e['time_unix_nano'])
                if e['adv_present'] and e['adv_info_guid'] == h:
                    if e['adv_level']: a['lv'][e['adv_level']] += 1
                    if e['adv_attack_power']: a['ap'].append(e['adv_attack_power'])
                if e['event'] == 'SPELL_CAST_SUCCESS': a['casts'][key] += 1
                if e['event'] in DMG and e['amount']:
                    a['dmg'][key] += e['amount']; a['hit'][key] += 1; a['crit'][key] += bool(e['critical'])
                if e['event'] in MISS: a['miss'][(key, e['miss_type'])] += 1
            elif e['event'] in DMG and e['amount']:
                a['pet'][(e['source_name'], key[1])] += e['amount']
        for h, ts in seen.items():
            agg[h]['ms'] += (max(ts) - min(ts)) / 1e6; agg[h]['fights'] += 1
        for h in seen:
            au = sorted(e['time_unix_nano'] for e in t if e['source_guid'] == h and e['event'] == 'SPELL_CAST_SUCCESS' and e['spell_id'] == 75)
            agg[h]['gaps'] += [(b - a) / 1e9 for a, b in zip(au, au[1:]) if (b - a) / 1e9 < 5]
    for h, a in agg.items():
        if not a['lv'] or max(a['lv']) < minlvl or want not in a['name']: continue
        mins = a['ms'] / 60000; tot = sum(a['dmg'].values()); pet = sum(a['pet'].values())
        print(f"== {rep} {a['name']} lvl {min(a['lv'])}-{max(a['lv'])} fights {a['fights']} active-span {mins:.1f} min; AP med {med(a['ap'])}; own dmg {tot}, pet/summon dmg {pet} ({100*pet/max(1,tot+pet):.1f}% share)")
        for k, v in a['dmg'].most_common(12):
            print(f"  {k[1]:28s} {k[0]:>8} dmg {v:7d} {100*v/max(1,tot+pet):5.1f}%  hits {a['hit'][k]:4d} crit {100*a['crit'][k]/max(1,a['hit'][k]):5.1f}%  casts/min {a['casts'][k]/max(mins,1e-9):.2f}")
        print('  pets:', [(k, v) for k, v in a['pet'].most_common(8)])
        print('  casts:', [(k[1], k[0], v) for k, v in a['casts'].most_common(14)])
        print('  misses:', [(k[0][1], k[1], v) for k, v in a['miss'].most_common(8)])
        print('  auto-shot gaps (<5s):', len(a['gaps']), 'median', med(a['gaps']))
