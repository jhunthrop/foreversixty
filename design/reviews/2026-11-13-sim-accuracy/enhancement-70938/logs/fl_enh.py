"""fl_enh.py <fl dir> <report>...: per Stormstrike shaman in foreverlogs.gg reports (public names): swing speed from
white swing gaps, Stormstrike gaps and size against white hits, Stormstrike outcome table, Windfury Weapon and Windfury
Totem procs, Flametongue Totem procs, Maelstrom Weapon, casts a minute of fight time, damage shares."""
import sys, json, collections, statistics as st
SS, WFW, MW = 17364, 439440, 408505
WFT = {8516, 10608, 10610}; FTT = {8253, 8248, 10523, 16389}
AVOID = ('miss', 'dodge', 'parry', 'block', 'immune')
def q(xs, i): return st.quantiles(xs, n=10)[i] if len(xs) >= 10 else None
root = sys.argv[1]
for rid in sys.argv[2:]:
    ev = json.load(open(f'{root}/shaman-events-{rid}.json'))['events']
    by = collections.defaultdict(list)
    for e in ev: by[e['source']['name']].append(e)
    for nm, es in by.items():
        sid = lambda e: (e.get('spell') or {}).get('id')
        if not any(e['type'] == 'cast' and sid(e) == SS for e in es): continue
        enc = collections.defaultdict(list)
        for e in es: enc[e['encounterId']].append(e)
        ftime = sum((max(x['tMs'] for x in v) - min(x['tMs'] for x in v)) / 1000 for v in enc.values())
        sw, ssg, wfw_t, wft_t = [], [], [], []
        both = 0
        for v in enc.values():
            v.sort(key=lambda e: e['tMs'])
            w = [e['tMs'] for e in v if sid(e) == -1 and e['type'] in ('damage',) + AVOID]
            sw += [(b - a) / 1000 for a, b in zip(w, w[1:]) if b - a < 6000]
            s = [e['tMs'] for e in v if sid(e) == SS and e['type'] == 'cast']
            ssg += [(b - a) / 1000 for a, b in zip(s, s[1:]) if b - a < 30000]
            p1 = sorted({e['tMs'] for e in v if sid(e) == WFW}); p2 = sorted({e['tMs'] for e in v if sid(e) in WFT})
            wfw_t += [(b - a) / 1000 for a, b in zip(p1, p1[1:]) if b - a < 30000]; wft_t += p2; both += len(set(p1) & set(p2))
        white = [e for e in es if sid(e) == -1 and e['type'] in ('damage',) + AVOID]
        wh = [e['amount'] for e in white if e['type'] == 'damage' and not e['isCritical'] and e['target']['isPlayer'] is False]
        ssd = [e for e in es if sid(e) == SS and e['type'] != 'cast']
        ssn = [e['amount'] for e in ssd if e['type'] == 'damage' and not e['isCritical']]
        wfw = [e for e in es if sid(e) == WFW]
        procs_wfw = len({(e['encounterId'], e['tMs']) for e in wfw}); procs_wft = len({(e['encounterId'], e['tMs']) for e in es if sid(e) in WFT})
        landed_mh = sum(1 for e in white if e['type'] == 'damage') + sum(1 for e in ssd if e['type'] == 'damage')
        print(f'== {rid} {nm}: fight time {ftime:.0f} s over {len(enc)} encounters')
        print(f'   swing gap median {st.median(sw):.3f} p90 {q(sw, -1):.3f} (n {len(sw)})')
        print(f'   Stormstrike casts {sum(1 for e in es if sid(e) == SS and e["type"] == "cast")}, gaps n {len(ssg)} min {min(ssg) if ssg else None} p10 {q(ssg, 0)} median {st.median(ssg) if ssg else None}; '
              f'gaps <7.9 s {sum(g < 7.9 for g in ssg)}, 7.9-8.6 {sum(7.9 <= g < 8.6 for g in ssg)}')
        oc = collections.Counter(e['type'] + ('+crit' if e['isCritical'] else '') for e in ssd)
        print(f'   Stormstrike outcomes {dict(oc)}; white outcomes {dict(collections.Counter(e["type"] + ("+crit" if e["isCritical"] else "") for e in white))}')
        if ssn and wh:
            hi = sorted(wh)[len(wh) // 2:]; hs = sorted(ssn)[len(ssn) // 2:]
            print(f'   SS/white non-crit: median {st.median(ssn) / st.median(wh):.3f} (SS {st.median(ssn)} n {len(ssn)}, white {st.median(wh)} n {len(wh)}); upper-half medians {st.median(hs) / st.median(hi):.3f}')
        print(f'   Windfury Weapon: procs {procs_wfw} strikes {len(wfw)} {dict(collections.Counter(e["type"] for e in wfw))}, min proc gap {min(wfw_t) if wfw_t else None}, '
              f'procs per landed main-hand hit {procs_wfw / landed_mh if landed_mh else 0:.3f}; Windfury Totem extra attacks {procs_wft}; same-ms both {both}')
        print(f'   Flametongue Totem procs {dict(collections.Counter(sid(e) for e in es if sid(e) in FTT))}; Maelstrom events {sum(1 for e in es if sid(e) == MW)}')
        cs = collections.Counter(((e['spell'] or {}).get('name'), sid(e)) for e in es if e['type'] == 'cast')
        print('   casts/min', {f'{k[0]}:{k[1]}': round(60 * v / ftime, 2) for k, v in cs.most_common(14)})
        dm = collections.Counter()
        for e in es:
            if e['type'] == 'damage' and not e['target']['isPlayer']: dm[(e['spell'] or {}).get('name')] += e['amount']
        tot = sum(dm.values())
        print('   damage share', {k: f'{100 * v / tot:.0f}%' for k, v in dm.most_common(10)}, f'total {tot}, dps over fight time {tot / ftime:.1f}')
