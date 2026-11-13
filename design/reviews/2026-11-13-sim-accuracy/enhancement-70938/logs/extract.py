"""extract.py <own root> <out.parquet> <report>...: every event whose source or dest is a shaman (a unit that casts
any shaman class spell, by spell id) from our downloaded fight files, with the report and fight index."""
import sys, os, json, glob
import pyarrow.parquet as pq, pandas as pd
B = os.path.join(os.path.dirname(os.path.abspath(__file__)), '../../../../../data/builds/1.60.1.70938/spellconst/shaman.json')
SHAMAN = {int(k) for k in json.load(open(B))['spells']} | {17364, 408505}
COLS = ['time_unix_nano', 'event', 'source_guid', 'source_name', 'dest_guid', 'dest_name', 'dest_flags', 'spell_id', 'spell_name',
        'extra_spell_id', 'amount', 'base_amount', 'overkill', 'resisted', 'blocked', 'absorbed', 'critical', 'glancing',
        'off_hand', 'miss_type', 'stacks', 'aura_type', 'adv_present', 'adv_info_guid', 'adv_attack_power', 'adv_spell_power',
        'adv_current_power', 'adv_max_power', 'adv_power_cost', 'adv_level']
root, out = sys.argv[1], sys.argv[2]
frames = []
for rep in sys.argv[3:]:
    tabs = {}
    for fp in glob.glob(f'{root}/{rep}/fights/*/events.parquet'):
        tabs[int(fp.split('/')[-2])] = pq.read_table(fp, columns=COLS).to_pandas()
    sham = set()
    for t in tabs.values():
        c = t[(t.event == 'SPELL_CAST_SUCCESS') & t.spell_id.isin(SHAMAN) & t.source_guid.str.startswith('Player', na=False)]
        sham |= set(c.source_guid)
    for i, t in tabs.items():
        t = t[t.source_guid.isin(sham) | t.dest_guid.isin(sham)].copy()
        t['rep'] = rep; t['fight'] = i; frames.append(t)
    print(rep, len(sham), 'shamans', flush=True)
df = pd.concat(frames); df.to_parquet(out); print(len(df), 'events')
