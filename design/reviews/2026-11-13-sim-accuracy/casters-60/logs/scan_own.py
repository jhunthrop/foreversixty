"""scan_own.py <dl root>: every druid and shaman in our backfilled reports who casts a Balance or
Elemental damage spell: highest adv_level, highest rank cast, cast counts, fight kinds. Private
reports are printed in aggregate only (no names)."""
import sys, os, json, glob, collections
import pyarrow.parquet as pq
SPELLS = {'Wrath', 'Starfire', 'Moonfire', 'Insect Swarm', 'Lightning Bolt', 'Chain Lightning',
          'Earth Shock', 'Flame Shock', 'Frost Shock', 'Lava Burst', 'Hurricane'}
BUILD = os.path.join(os.path.dirname(os.path.abspath(__file__)), '../../../../../data/builds/1.60.1.70938/spellconst')
IDS = {}
for c in ('druid', 'shaman'):
    for sid, sp in json.load(open(f'{BUILD}/{c}.json'))['spells'].items():
        if sp['name'] in SPELLS: IDS[int(sid)] = sp['name']
COLS = ['event', 'source_guid', 'spell_id', 'spell_name', 'adv_present', 'adv_info_guid', 'adv_level']
root = sys.argv[1]
out = collections.defaultdict(lambda: {'levels': collections.Counter(), 'spells': collections.Counter(), 'fights': set(), 'cls': ''})
for rep in sorted(os.listdir(root)):
    rp = f'{root}/{rep}/report.json'
    if not os.path.exists(rp): continue
    rj = json.load(open(rp))
    cls = {u['guid']: u.get('class') for u in rj['units'] if u.get('class') in ('Druid', 'Shaman')}
    for fp in glob.glob(f'{root}/{rep}/f*.parquet'):
        for e in pq.read_table(fp, columns=COLS).to_pylist():
            g = e['source_guid']
            if g not in cls: continue
            o = out[(rep, g)]; o['cls'] = cls[g]
            if e['adv_present'] and e['adv_info_guid'] == g and e['adv_level']: o['levels'][e['adv_level']] += 1
            if e['event'] == 'SPELL_CAST_SUCCESS' and e['spell_id'] in IDS:
                o['spells'][(IDS[e['spell_id']], e['spell_id'])] += 1; o['fights'].add(fp)
rows = [(rep, o) for (rep, g), o in out.items() if o['spells']]
rows.sort(key=lambda r: -max(r[1]['levels'] or {0: 0}))
for rep, o in rows[:40]:
    lv = sorted(o['levels'])
    print(rep, o['cls'], 'lvl', (lv[0], lv[-1]) if lv else None, 'fights', len(o['fights']),
          [f'{k[0]}:{k[1]}x{v}' for k, v in o['spells'].most_common(6)])
