"""scan.py <dl root>: every player casting a cat or bear ability: level range, casts by ability."""
import sys, os, glob, collections
import pyarrow.parquet as pq
FERAL = {'Rip','Ferocious Bite','Shred','Claw','Rake','Maul','Lacerate','Swipe','Primal Bite','Ravage','Pounce','Demoralizing Roar','Faerie Fire (Feral)','Shifting Power','Enrage','Frenzied Regeneration'}
COLS = ['event','source_guid','source_name','spell_id','spell_name','adv_present','adv_info_guid','adv_level']
root = sys.argv[1]
out = collections.defaultdict(lambda: {'name':'','lv':collections.Counter(),'sp':collections.Counter(),'f':set()})
for rep in sorted(os.listdir(root)):
    for fp in glob.glob(f'{root}/{rep}/*.parquet'):
        t = pq.read_table(fp, columns=COLS).to_pylist()
        fer = {e['source_guid'] for e in t if e['spell_name'] in FERAL and e['source_guid'].startswith('Player')}
        for e in t:
            g = e['source_guid']
            if g not in fer: continue
            o = out[(rep,g)]; o['name'] = e['source_name']
            if e['adv_present'] and e['adv_info_guid'] == g and e['adv_level']: o['lv'][e['adv_level']] += 1
            if e['event'] in ('SPELL_CAST_SUCCESS',) and e['spell_name'] in FERAL: o['sp'][e['spell_name']] += 1; o['f'].add(fp)
for (rep,g),o in sorted(out.items(), key=lambda kv: -max(kv[1]['lv'] or {0:0})):
    lv = sorted(o['lv'])
    if not o['sp']: continue
    print(rep, o['name'][:24], 'lvl', (lv[0],lv[-1]) if lv else None, 'fights', len(o['f']), dict(o['sp'].most_common(10)))
