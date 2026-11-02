"""Every published healer item (all bands) and its client item-effect spells, with each spell's effects."""
import csv, collections, json, re
SP = '/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad'
RAW = SP + '/r9/casters/raw/'
S = '/Users/jh/code/forever/.claude/worktrees/r11-healers/data/builds/1.60.1.70707/bis/'
def rows(t): return list(csv.DictReader(open(RAW + t + '.csv', newline='')))
x = collections.defaultdict(list)
for r in rows('ItemXItemEffect'): x[r['ItemID']].append(r['ItemEffectID'])
ie = {r['ID']: r for r in rows('ItemEffect')}
eff = collections.defaultdict(list)
for r in rows('SpellEffect'): eff[r['SpellID']].append(r)
names = {r['ID']: r['Name_lang'] for r in rows('SpellName')}
sets = collections.defaultdict(list)
item_set = {}
for r in rows('ItemSparse'):
    if r.get('ItemSet') and r['ItemSet'] != '0': item_set[r['ID']] = r['ItemSet']
for r in rows('ItemSetSpell'): sets[r['ItemSetID']].append((r['Threshold'], r['SpellID']))
used = collections.defaultdict(set)
for spec in ['druid-restoration', 'shaman-restoration', 'priest-holy', 'priest-discipline', 'paladin-holy']:
    d = json.load(open(S + spec + '.json'))
    for b in d['bands']:
        for s in b['slots']:
            if s.get('item_id'): used[(str(s['item_id']), s.get('item_name'))].add(f"{spec.split('-')[0][:3]}{b['band']}{b['preset'][0]}")
TRIG = {'0': 'use', '1': 'equip', '2': 'proc', '5': 'learn', '6': 'soulbound-learn'}
seen_sets = set()
for (iid, nm), where in sorted(used.items(), key=lambda kv: kv[0][1] or ''):
    fx = [ie[e] for e in x.get(iid, []) if e in ie]
    st = item_set.get(iid)
    line = []
    for f in fx:
        sid = f['SpellID']
        es = '; '.join(f"e{e['EffectIndex']} eff{e['Effect']} aura{e['EffectAura']} bp{e['EffectBasePointsF']} misc{e['EffectMiscValue_0']} trig{e['EffectTriggerSpell']}" for e in eff.get(sid, []))
        line.append(f"{TRIG.get(f['TriggerType'], f['TriggerType'])} {sid} {names.get(sid)} cd{f.get('CoolDownMSec')} cat{f.get('CategoryCoolDownMSec')} [{es}]")
    if line or st:
        print(f"{iid} {nm} ({','.join(sorted(where))}) set={st}")
        for l in line: print('    ', l)
    if st and st not in seen_sets:
        seen_sets.add(st)
        for th, sid in sorted(sets[st]):
            es = '; '.join(f"e{e['EffectIndex']} eff{e['Effect']} aura{e['EffectAura']} bp{e['EffectBasePointsF']} trig{e['EffectTriggerSpell']}" for e in eff.get(sid, []))
            print(f"      set {st} {th}p {sid} {names.get(sid)} [{es}]")
