"""hawk.py <fight parquet>...: Summon Hawk (1293241) cadence, summon spells, hawk unit lifetime, swing interval and hit sizes."""
import sys, collections, statistics as st
import pyarrow.parquet as pq
HAWK = {1293241, 1293525}  # Summon Hawk ranks seen in logs
casts, summons, hawkdmg, spelldmg, swings = [], collections.Counter(), [], [], collections.defaultdict(list)
for fp in sys.argv[1:]:
    t = pq.read_table(fp).to_pylist(); hawks = {}
    for e in t:
        if e['spell_id'] in HAWK and e['event'] == 'SPELL_CAST_SUCCESS': casts.append((fp, e['time_unix_nano'] / 1e9))
        if e['event'] == 'SPELL_SUMMON' and 'Hawk' in e['spell_name'] or (e['event'] == 'SPELL_SUMMON' and e['spell_id'] in (1293248, 1312637)):
            summons[(e['spell_id'], e['dest_name'])] += 1; hawks[e['dest_guid']] = e['time_unix_nano'] / 1e9
        if e['spell_id'] in HAWK and e['event'] == 'SPELL_DAMAGE': spelldmg.append((e['amount'], e['spell_id'], e['dest_name']))
    for e in t:
        if e['source_guid'] in hawks and e['event'] == 'SWING_DAMAGE_LANDED':
            hawkdmg.append(e['amount']); swings[e['source_guid']].append(e['time_unix_nano'] / 1e9)
life = {g: (ts[-1] - ts[0], len(ts)) for g, ts in swings.items()}
gaps = [b - a for ts in swings.values() for a, b in zip(ts, ts[1:])]
cg = [b[1] - a[1] for a, b in zip(casts, casts[1:]) if a[0] == b[0]]
print('Summon Hawk casts', len(casts), 'cast gaps min/med', min(cg) if cg else None, st.median(cg) if cg else None)
print('summons', summons.most_common())
print('summon-spell damage events', len(spelldmg), 'median', st.median([x[0] for x in spelldmg]) if spelldmg else None, spelldmg[:3])
print('hawk units', len(swings), 'swings per hawk med', st.median([v[1] for v in life.values()]) if life else None, 'first-last swing span med', st.median([v[0] for v in life.values()]) if life else None)
print('hawk swing gap med', st.median(gaps) if gaps else None, 'p10', sorted(gaps)[len(gaps)//10] if gaps else None, 'n', len(gaps), 'hit med', st.median(hawkdmg) if hawkdmg else None)
