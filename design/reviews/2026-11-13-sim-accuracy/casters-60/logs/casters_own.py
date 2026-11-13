"""casters_own.py <dl root> <min level>: every druid and shaman at or above min level (adv_level) in our
reports who casts a Balance or Elemental damage spell: per spell, casts, landed hits, crits, misses
and resists; Nature's Grace (16886), Clearcasting (16870, 16246) and Lightning Overload hits; their
mana at their own casts. Players are numbered, never named (private reports)."""
import sys, os, json, glob, collections
import pyarrow.parquet as pq
root, minlv = sys.argv[1], int(sys.argv[2])
B = os.path.join(os.path.dirname(os.path.abspath(__file__)), '../../../../../data/builds/1.60.1.70938/spellconst')
NAMES = {'Wrath', 'Starfire', 'Moonfire', 'Insect Swarm', 'Lightning Bolt', 'Chain Lightning', 'Earth Shock', 'Flame Shock', 'Frost Shock'}
IDS = {}
for c in ('druid', 'shaman'):
    for sid, sp in json.load(open(f'{B}/{c}.json'))['spells'].items():
        if sp['name'] in NAMES: IDS[int(sid)] = sp['name'] + (' overload' if int(sid) >= 408000 and sp['name'] in ('Lightning Bolt', 'Chain Lightning') else '')
AURAS = {16886: "Nature's Grace", 16870: 'Clearcasting (druid)', 16246: 'Clearcasting (shaman)'}
COLS = ['event', 'source_guid', 'dest_guid', 'spell_id', 'critical', 'miss_type', 'amount', 'adv_present', 'adv_info_guid', 'adv_level', 'adv_current_power', 'adv_max_power']
n = 0
for rep in sorted(os.listdir(root)):
    rp = f'{root}/{rep}/report.json'
    if not os.path.exists(rp): continue
    cls = {u['guid']: u['class'] for u in json.load(open(rp))['units'] if u.get('class') in ('Druid', 'Shaman')}
    st = collections.defaultdict(lambda: collections.defaultdict(collections.Counter)); lv = collections.defaultdict(set); mana = collections.defaultdict(list)
    for fp in glob.glob(f'{root}/{rep}/f*.parquet'):
        for e in pq.read_table(fp, columns=COLS).to_pylist():
            g = e['source_guid']
            if g not in cls: continue
            if e['adv_present'] and e['adv_info_guid'] == g and e['adv_level']: lv[g].add(e['adv_level'])
            sid = e['spell_id']
            if sid in AURAS and e['event'] == 'SPELL_AURA_APPLIED' and e['dest_guid'] == g: st[g][AURAS[sid]]['applied'] += 1
            if sid not in IDS: continue
            c = st[g][IDS[sid]]; ev = e['event']
            if ev == 'SPELL_CAST_SUCCESS':
                c['cast'] += 1
                if e['adv_info_guid'] == g and e['adv_max_power']: mana[g].append(e['adv_current_power'] / e['adv_max_power'])
            elif ev in ('SPELL_DAMAGE', 'SPELL_PERIODIC_DAMAGE'):
                k = 'tick' if 'PERIODIC' in ev else 'hit'; c[k] += 1; c[k + '_crit'] += bool(e['critical'])
            elif ev == 'SPELL_MISSED': c['miss_' + (e['miss_type'] or '?')] += 1
    for g in st:
        if not lv[g] or max(lv[g]) < minlv: continue
        n += 1; m = sorted(mana[g])
        print(f'{rep} {cls[g]} #{n} levels {sorted(lv[g])} mana at cast min {m[0]:.2f} median {m[len(m)//2]:.2f}' if m else f'{rep} {cls[g]} #{n} levels {sorted(lv[g])}')
        for k, c in sorted(st[g].items()): print('   ', k, dict(c))
