"""fl_ssratio.py <fl dir> <report> <name> <weapon min> <weapon max> <speed>: Stormstrike against white hits on the same
target. Quantiles of both non-crit distributions, the per-target median ratio, and the attack power that the white hits
imply under each Stormstrike model (100% or 150% of normalized weapon damage, 3.3 for a two-hander)."""
import sys, json, collections, statistics as st
root, rid, nm = sys.argv[1], sys.argv[2], sys.argv[3]; wmin, wmax, spd = map(float, sys.argv[4:7])
ev = [e for e in json.load(open(f'{root}/shaman-events-{rid}.json'))['events'] if e['source']['name'] == nm and e['type'] == 'damage' and not e['isCritical'] and not e['target']['isPlayer']]
sid = lambda e: (e.get('spell') or {}).get('id')
wh = collections.defaultdict(list); ss = collections.defaultdict(list)
for e in ev:
    k = (e['encounterId'], e['target']['id'])
    if sid(e) == -1: wh[k].append(e['amount'])
    elif sid(e) == 17364: ss[k].append(e['amount'])
W = sorted(x for v in wh.values() for x in v); S = sorted(x for v in ss.values() for x in v)
qs = lambda xs: [round(st.quantiles(xs, n=20)[i]) for i in (4, 9, 14, 18)] + [max(xs)]
print('white q25 q50 q75 q95 max', qs(W), 'n', len(W)); print('SS    q25 q50 q75 q95 max', qs(S), 'n', len(S))
print('quantile ratios', [round(a / b, 3) for a, b in zip(qs(S), qs(W))])
pr = [st.median(ss[k]) / st.median(wh[k]) for k in ss if wh.get(k) and len(wh[k]) >= 3]
print('per-target median ratio', round(st.median(pr), 3), 'targets', len(pr))
# The top of the white range is the max roll on the least armored target; solve AP from the white q95, then predict SS q95.
for pct in (1.0, 1.5):
    wq = st.quantiles(W, n=20)[18]; ap = (wq / 1 - wmax) * 14 / spd   # before armor; armor ignored, so this is the ratio test
    pred = pct * (wmax + ap * 3.3 / 14)
    print(f'model {int(pct * 100)}%: AP from white q95 {ap:.0f}, predicted SS q95 {pred:.0f} vs logged {st.quantiles(S, n=20)[18]:.0f}')
