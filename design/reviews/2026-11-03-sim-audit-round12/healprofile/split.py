"""Tank share of the damage players take, per fight, in foreverlogs.gg beta dungeon runs.

A fight's group is every player that deals or takes damage in it; only 3-to-6-player fights
count (five-man groups). Damage taken is every creature-to-player damage event (pets and
player-to-player damage are left out: the profile heals players). The tank is the player who
takes the most creature melee (Auto Attack) in the fight, and must be a warrior, paladin or
druid; a fight whose top melee taker is anything else is counted as "no tank" (a pet or a
caster held the pull) and reported apart. Tank share = the tank's damage taken over the group's.
"""
import collections, datetime, glob, json, statistics, sys
SP = '/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad'
DT = SP + '/r12/healprofile/fl/dt'
DUNGEON = {
    'Deadmines': ["Rhahk'zor", 'Sneed', 'Gilnid', 'Mr. Smite', 'Captain Greenskin', 'Cookie', 'Edwin VanCleef'],
    'Wailing Caverns': ['Kresh', 'Lady Anacondra', 'Lord Cobrahn', 'Lord Pythas', 'Lord Serpentis', 'Skum', 'Verdan the Everliving', 'Mutanus the Devourer'],
    'Razorfen Kraul': ['Agathelos the Raging', 'Aggem Thorncurse', 'Charlga Razorflank', 'Death Speaker Jargba', 'Overlord Ramtusk', 'Roogug'],
    'Gnomeregan': ['Crowd Pummeler 9-60', 'Electrocutioner 6000', 'Grubbis', 'Mekgineer Thermaplugg', 'Viscous Fallout'],
    'SM Graveyard': ['Bloodmage Thalnos', 'Interrogator Vishas'],
    'SM Armory': ['Herod', 'Houndmaster Loksey'],
    'Shadowfang Keep': ['Rethilgore', 'Razorclaw the Butcher', 'Baron Silverlaine', 'Commander Springvale', 'Odo the Blindwatcher', 'Fenrus the Devourer', 'Wolf Master Nandos', 'Archmage Arugal'],
    'Blackfathom Deeps': ['Ghamoo-ra', 'Lady Sarevess', 'Gelihast', 'Lorgus Jett', 'Baron Aquanis', 'Twilight Lord Kelris', "Old Serra'kis", "Aku'mai"],
    'The Stockade': ['Targorr the Dread', 'Kam Deepfury', 'Hamhock', 'Bazil Thredd', 'Dextren Ward'],
    'Ragefire Chasm': ['Oggleflint', 'Taragaman the Hungerer', 'Jergosh the Invoker', 'Bazzalan'],
    'Ruins of Lordaeron (Forever)': ['Bjork', "Rath'mael", 'The Abandoned', 'The Baron', 'Viktor the Vile', 'Witherfang', 'Lordaeron Captain'],
    'Excavation Site: Wetlands (Forever)': ['Saltspine', 'Shadetooth', 'Relic Guardian', 'Highland Horror'],
    'arcane wing (Forever)': ['Arcane Anomaly', 'Mana Devourer', 'Fel Ancient', 'Shade of the Archmage', 'Unstable Sentinel', 'Lyn the Ignored', 'Mana Wraith', 'Atrexis the Grave Knight'],
}
BOSS_DUNGEON = {b: d for d, bs in DUNGEON.items() for b in bs}
BAND = {'Deadmines': 20, 'Wailing Caverns': 20, 'Razorfen Kraul': 30, 'Gnomeregan': 30, 'SM Graveyard': 30, 'SM Armory': 40}
TANKS = {'Warrior', 'Paladin', 'Druid'}

def ts(s):
    return datetime.datetime.fromisoformat(s.replace('Z', '+00:00')).timestamp()

def fight_split(events):
    taken = collections.defaultdict(float); melee = collections.defaultdict(float); cls = {}; players = set(); mobs = set()
    for e in events:
        if e['source_type'] in ('player', 'pet') and e['target_type'] == 'creature':
            mobs.add(e['target_id'])
        if e['source_type'] == 'player':
            players.add(e['source_id']); cls[e['source_id']] = e['source_class']
        if e['target_type'] != 'player':
            continue
        players.add(e['target_id']); cls[e['target_id']] = e['target_class']
        if e['source_type'] != 'creature':
            continue
        a = float(e['amount'] or 0)
        taken[e['target_id']] += a
        if e['spell_id'] == -1:
            melee[e['target_id']] += a
    total = sum(taken.values())
    if not (3 <= len(players) <= 6) or total <= 0 or not melee:
        return None
    top = max(melee, key=melee.get)
    return {'players': len(players), 'total': total, 'tank_class': cls.get(top),
            'has_tank': cls.get(top) in TANKS, 'share': taken[top] / total,
            'melee_share': sum(melee.values()) / total, 'mobs': len(mobs), 'tank_id': top, 'group': sorted(players),
            'taken': dict(taken), 'melee': dict(melee), 'cls': cls}

def main():
    reps = {int(f.split('/')[-1].split('.')[0]): json.load(open(f)) for f in glob.glob(SP + '/fl/rep/*[0-9].json')}
    rows = []
    for f in glob.glob(DT + '/*.json'):
        rid, eid = map(int, f.split('/')[-1].split('.')[:2])
        encs = reps[rid]['encounters']; enc = next(e for e in encs if e['id'] == eid)
        kills = [e for e in encs if e['success'] and e['name'] in BOSS_DUNGEON]
        if enc['success']:
            if enc['name'] not in BOSS_DUNGEON:
                continue
            kind, dungeon = 'boss', BOSS_DUNGEON[enc['name']]
        else:
            if not kills:
                continue
            near = min(kills, key=lambda k: abs(ts(k['start_time']) - ts(enc['start_time'])))
            if abs(ts(near['start_time']) - ts(enc['start_time'])) > 1800:
                continue
            kind, dungeon = 'trash', BOSS_DUNGEON[near['name']]
        try:
            ev = json.load(open(f))['events'] or []
        except Exception:
            continue
        s = fight_split(ev)
        if s:
            rows.append(dict(s, rid=rid, eid=eid, name=enc['name'], kind=kind, dungeon=dungeon, band=BAND.get(dungeon)))
    json.dump(rows, open(SP + '/r12/healprofile/fl/split.json', 'w'))
    report(rows)

def summarise(label, rs):
    t = [r for r in rs if r['has_tank']]
    if len(t) < 3:
        print(f'{label:46} fights {len(rs):4}  (too few with a tank: {len(t)})'); return
    sh = sorted(r['share'] for r in t)
    q = statistics.quantiles(sh, n=4)
    pooled = sum(r['share'] * r['total'] for r in t) / sum(r['total'] for r in t)
    print(f'{label:46} fights {len(rs):4}  tanked {len(t):4} ({len(t)/len(rs):.0%})  median {statistics.median(sh):.3f}  IQR {q[0]:.2f}-{q[2]:.2f}  pooled {pooled:.3f}  melee share {statistics.median(r["melee_share"] for r in t):.2f}')

def run_tanks(rows):
    """The run's tank: in each report, for each group (the same players), the warrior, paladin or
    druid who takes the most creature melee over all the group's fights. Each fight's share is
    then that player's, so a fight where a loose mob or a pet held the pull still counts."""
    melee = collections.defaultdict(lambda: collections.defaultdict(float)); cls = {}
    for r in rows:
        key = (r['rid'], tuple(r['group']))
        for pid, a in r['melee'].items():
            melee[key][pid] += a
        cls.update(r['cls'])
    for r in rows:
        key = (r['rid'], tuple(r['group']))
        cands = {p: a for p, a in melee[key].items() if cls.get(p) in TANKS}
        if not cands:
            r['run_share'] = None; continue
        tank = max(cands, key=cands.get)
        r['run_share'] = r['taken'].get(tank, 0.0) / r['total']

def summarise_run(label, rs):
    t = [r for r in rs if r.get('run_share') is not None]
    if len(t) < 3:
        print(f'{label:46} run-tank fights {len(t)}'); return None
    sh = sorted(r['run_share'] for r in t); q = statistics.quantiles(sh, n=4)
    pooled = sum(r['run_share'] * r['total'] for r in t) / sum(r['total'] for r in t)
    print(f'{label:46} run-tank fights {len(t):4}  median {statistics.median(sh):.3f}  IQR {q[0]:.2f}-{q[2]:.2f}  pooled {pooled:.3f}')
    return statistics.median(sh)

def pull_sizes(label, rs):
    sizes = sorted(r['mobs'] for r in rs if r['mobs'] > 0)
    if not sizes:
        return
    weighted = sorted(x for s in sizes for x in [s] * s)
    pick = lambda q: weighted[min(len(weighted) - 1, int(q * len(weighted)))]
    print(f'{label:46} logged trash fights {len(sizes)}: mobs per fight median {statistics.median(sizes)}, mob-weighted median {pick(0.5)}, q75 {pick(0.75)}')

def report(rows):
    run_tanks(rows)
    for kind in ('boss', 'trash'):
        print(f'== run tank, {kind}')
        for band in (20, 30, 40):
            summarise_run(f'band {band} content', [r for r in rows if r['kind'] == kind and r['band'] == band])
        summarise_run('all beta dungeons', [r for r in rows if r['kind'] == kind])
    print('== run tank, boss and trash together')
    for band in (20, 30, 40):
        summarise_run(f'band {band} content', [r for r in rows if r['band'] == band])
    summarise_run('all beta dungeons', rows)
    print('== logged trash pull sizes')
    for band in (20, 30):
        pull_sizes(f'band {band} content', [r for r in rows if r['kind'] == 'trash' and r['band'] == band])
    for d in ('Deadmines', 'Wailing Caverns', 'Razorfen Kraul', 'Gnomeregan', 'SM Graveyard'):
        pull_sizes(d, [r for r in rows if r['kind'] == 'trash' and r['dungeon'] == d])
    for kind in ('boss', 'trash'):
        print(f'== {kind}')
        for band in (20, 30, 40):
            summarise(f'band {band} content', [r for r in rows if r['kind'] == kind and r['band'] == band])
        for d in DUNGEON:
            summarise(d, [r for r in rows if r['kind'] == kind and r['dungeon'] == d])
        summarise('all beta dungeons', [r for r in rows if r['kind'] == kind])

if __name__ == '__main__':
    main()
