"""Beta logs: per healer class, overheal share and the top target's (tank proxy) share of raw healing.
Encounters where the healer healed >= 3000 raw and >= 3 distinct targets."""
import collections, glob, json, statistics
SP = '/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad'
TL = SP + '/r10/healers/fl/tl'
cls_of = {}
for f in glob.glob(SP + '/fl/rep/*.roster.json'):
    rid = int(f.split('/')[-1].split('.')[0])
    for c in json.load(open(f))['roster']:
        for s in (c.get('snapshots') or [])[:1]:
            cls_of[(rid, c['character_id'])] = s['player']['class']
per = collections.defaultdict(lambda: collections.defaultdict(list))
spell = collections.defaultdict(lambda: [0.0, 0.0])
tot = collections.defaultdict(lambda: [0.0, 0.0])
for f in glob.glob(TL + '/*.json'):
    rid, enc, cid = map(int, f.split('/')[-1].split('.')[:3])
    cl = cls_of.get((rid, cid))
    try: ev = json.load(open(f)).get('events') or []
    except Exception: continue
    eff = collections.defaultdict(float); over = collections.defaultdict(float); tgtcls = {}
    for e in ev:
        if e['source_id'] != cid or e['event_type'] not in ('heal', 'overheal'): continue
        a = float(e['amount']); t = e['target_id']; tgtcls[t] = e.get('target_class')
        (eff if e['event_type'] == 'heal' else over)[t] += a
        k = (cl, e['spell_name'])
        spell[k][0 if e['event_type'] == 'heal' else 1] += a
    raw = {t: eff[t] + over[t] for t in set(eff) | set(over)}
    R = sum(raw.values())
    if R < 3000 or len(raw) < 3: continue
    top = max(raw, key=raw.get)
    selfraw = raw.get(cid, 0)
    per[cl]['top'].append(raw[top] / R)
    per[cl]['self'].append(selfraw / R)
    per[cl]['over'].append(sum(over.values()) / R)
    per[cl]['topover'].append(over[top] / raw[top])
    tot[cl][0] += sum(eff.values()); tot[cl][1] += sum(over.values())
for cl, d in per.items():
    print(f"{cl}: encounters {len(d['top'])}; top-target share median {statistics.median(d['top']):.2f} (IQR {statistics.quantiles(d['top'])[0]:.2f}-{statistics.quantiles(d['top'])[2]:.2f}); self {statistics.median(d['self']):.2f}; overheal median {statistics.median(d['over']):.2f}, pooled {tot[cl][1]/(tot[cl][0]+tot[cl][1]):.2f}; on top target {statistics.median(d['topover']):.2f}")
print()
for (cl, sp), (h, o) in sorted(spell.items(), key=lambda kv: -(kv[1][0] + kv[1][1])):
    if h + o > 40000: print(f"  {cl:8} {sp:24} raw {h+o:9.0f} overheal {o/(h+o):.2f}")
