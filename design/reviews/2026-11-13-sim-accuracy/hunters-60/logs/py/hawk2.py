"""hawk2.py <data/own>: per Hawk unit (summon 1293248/1312637): lifetime vs its cast target's survival, swings, swing damage
by hunter level and AP, dive hit (Summon Hawk SPELL_DAMAGE 1293241/1293525). Output: data/hawk_units.json + printed tables."""
import sys, os, glob, json, collections, statistics as st
import pyarrow.parquet as pq
CAST = {1293241, 1293525}; SUMMON = {1293248, 1312637}; HOTFIX = 1794355200  # 2026-11-11 00:00 UTC
root = sys.argv[1]; units = []; dives = []
def hunter_snap(t, g, ts):
    """nearest adv snapshot of the hunter: (level, attack_power)."""
    best = None
    for e in t:
        if e['adv_present'] and e['adv_info_guid'] == g and e['source_guid'] == g and e['adv_level']:
            d = abs(e['time_unix_nano'] - ts)
            if best is None or d < best[0]: best = (d, e['adv_level'], e['adv_attack_power'])
    return best[1:] if best else (None, None)
def unit_level(t, g):
    lv = [e['adv_level'] for e in t if e['adv_present'] and e['adv_info_guid'] == g and e['adv_level']]
    return collections.Counter(lv).most_common(1)[0][0] if lv else None
for fp in sorted(glob.glob(f'{root}/*/fights/*/events.parquet')):
    t = pq.read_table(fp).to_pylist()
    if not any(e['spell_id'] in CAST for e in t): continue
    t.sort(key=lambda e: (e['time_unix_nano'], e['line'])); end = t[-1]['time_unix_nano'] / 1e9
    last_seen = collections.defaultdict(float); died = {}
    for e in t:
        for k in ('source_guid', 'dest_guid'):
            if e[k]: last_seen[e[k]] = e['time_unix_nano'] / 1e9
        if e['event'] in ('UNIT_DIED', 'PARTY_KILL') and e['dest_guid']: died.setdefault(e['dest_guid'], e['time_unix_nano'] / 1e9)
    casts = [e for e in t if e['event'] == 'SPELL_CAST_SUCCESS' and e['spell_id'] in CAST]
    for e in t:
        if e['event'] == 'SPELL_DAMAGE' and e['spell_id'] in CAST:
            lvl, ap = hunter_snap(t, e['source_guid'], e['time_unix_nano'])
            dives.append({'lvl': lvl, 'ap': ap, 'amt': e['amount'], 'base': e['base_amount'], 'crit': bool(e['critical']),
                          'tlvl': unit_level(t, e['dest_guid']), 'ts': e['time_unix_nano'] / 1e9, 'rank': e['spell_id']})
    for s in (e for e in t if e['event'] == 'SPELL_SUMMON' and e['spell_id'] in SUMMON):
        ts = s['time_unix_nano'] / 1e9; h = s['source_guid']; hg = s['dest_guid']
        c = max((x for x in casts if x['source_guid'] == h and x['time_unix_nano'] / 1e9 <= ts + 0.01), key=lambda x: x['time_unix_nano'], default=None)
        tgt = c['dest_guid'] if c else None
        sw = [e for e in t if e['source_guid'] == hg and e['event'] in ('SWING_DAMAGE', 'SWING_MISSED')]
        other = [e for e in t if hg in (e['source_guid'], e['dest_guid']) and e['event'] not in ('SWING_DAMAGE', 'SWING_DAMAGE_LANDED', 'SWING_MISSED', 'SPELL_SUMMON')]
        tdeath = died.get(tgt); tlast = last_seen.get(tgt, 0)
        alive18 = bool(tgt) and (tdeath or tlast) >= ts + 18 if tgt else None
        lvl, ap = hunter_snap(t, h, s['time_unix_nano'])
        units.append({'fp': fp.split('own/')[1], 'ts': ts, 'hunter': h, 'hawk': hg, 'target': tgt, 'hawk_last_seen': round(last_seen.get(hg, ts) - ts, 2), 'hunter_lvl': lvl, 'hunter_ap': ap, 'target_lvl': unit_level(t, tgt) if tgt else None,
                      'alive18': alive18, 'target_dies_at': round(tdeath - ts, 2) if tdeath else None, 'fight_left': round(end - ts, 2),
                      'swings': [{'dt': round(e['time_unix_nano'] / 1e9 - ts, 3), 'amt': e['amount'], 'base': e['base_amount'], 'crit': bool(e['critical']),
                                  'glance': bool(e['glancing']), 'miss': e['miss_type'], 'same_target': e['dest_guid'] == tgt,
                                  'hawk_ap': e['adv_attack_power'] if e['adv_present'] and e['adv_info_guid'] == hg else None} for e in sw],
                      'other_events': collections.Counter(e['event'] for e in other).most_common()})
json.dump({'units': units, 'dives': dives}, open('data/hawk_units.json', 'w'))
def summ(xs): return f"n={len(xs)} min={min(xs)} med={st.median(xs)} max={max(xs)}" if xs else 'n=0'
full = [u for u in units if u['alive18'] and u['fight_left'] >= 18]
print('hawk units', len(units), '| target alive >=18 s after summon and fight lasts 18 s:', len(full))
print(' swings per hawk:', sorted(collections.Counter(len(u['swings']) for u in full).items()))
print(' last swing dt:', summ([u['swings'][-1]['dt'] for u in full if u['swings']]))
print(' first swing dt:', summ([u['swings'][0]['dt'] for u in full if u['swings']]))
early = [u for u in units if u['target_dies_at'] is not None and u['target_dies_at'] < 17]
print('target died <17 s:', len(early), '| hawks with swings after target death:',
      sum(any(s['dt'] > u['target_dies_at'] + 0.2 for s in u['swings']) for u in early),
      '| on another target:', sum(any(not s['same_target'] for s in u['swings']) for u in early))
print(' swings after death dt-death:', [round(s['dt'] - u['target_dies_at'], 2) for u in early for s in u['swings'] if s['dt'] > u['target_dies_at'] + 0.2][:30])
print('other hawk events:', collections.Counter(k for u in units for k, _ in u['other_events']))
gaps = [b['dt'] - a['dt'] for u in units for a, b in zip(u['swings'], u['swings'][1:])]
print('swing gaps:', summ(sorted(gaps)), 'p10', sorted(gaps)[len(gaps) // 10] if gaps else None)
rows = collections.defaultdict(list)
for u in units:
    for s in u['swings']:
        rows[(u['hunter_lvl'], 'post' if u['ts'] >= HOTFIX else 'pre')].append((s, u))
print('lvl period | swings landed nonCrit-med mean base-med | crit% glance miss | hunterAP med | hawkAP | target lvls')
for k in sorted(rows, key=str):
    xs = rows[k]; land = [s for s, _ in xs if not s['miss']]; nc = [s['amt'] for s in land if not s['crit']]
    print(k, len(xs), len(land), st.median(nc) if nc else None, round(st.mean(nc), 1) if nc else None,
          st.median([s['base'] for s in land if not s['crit'] and s['base']]) if any(s['base'] for s in land) else None,
          f"{100*sum(s['crit'] for s in land)/max(1,len(land)):.1f}%", sum(s['glance'] for s in land),
          collections.Counter(s['miss'] for s, _ in xs if s['miss']), st.median([u['hunter_ap'] for _, u in xs if u['hunter_ap']] or [0]),
          sorted({s['hawk_ap'] for s, _ in xs if s['hawk_ap']})[:5], sorted(collections.Counter(u['target_lvl'] for _, u in xs).items(), key=str))
print('dives: lvl period | n nonCrit med min max | crit | hunterAP med | target lvls')
dr = collections.defaultdict(list)
for d in dives: dr[(d['lvl'], d['rank'], 'post' if d['ts'] >= HOTFIX else 'pre')].append(d)
for k in sorted(dr, key=str):
    xs = dr[k]; nc = [d['amt'] for d in xs if not d['crit']]
    print(k, len(xs), st.median(nc) if nc else None, min(nc) if nc else None, max(nc) if nc else None, sum(d['crit'] for d in xs),
          st.median([d['ap'] for d in xs if d['ap']] or [0]), sorted(collections.Counter(d['tlvl'] for d in xs).items(), key=str))
