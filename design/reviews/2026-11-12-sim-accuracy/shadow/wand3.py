import json, collections, statistics as st, sys
name=lambda e:(e.get('spell') or {}).get('name')
INST={'Shadow Word: Pain','Devouring Plague','Vampiric Embrace','Inner Fire','Corruption','Curse of Agony','Bane of Agony','Curse of Weakness','Siphon Life','Life Tap','Fire Blast','Arcane Explosion','Immolate'}
for f in sys.argv[1:]:
    ev=json.load(open(f)); ev.sort(key=lambda e:(e['encounterId'], e['tMs']))
    by_enc=collections.defaultdict(list)
    for e in ev: by_enc[e['encounterId']].append(e)
    pair_gaps=[]; fires=[]; starts=[]
    after_inst_start=[]; after_inst_fire=[]; after_hard_start=[]; after_hard_fire=[]; after_chan_fire=[]; fire_gaps=[]
    for enc,L in by_enc.items():
        S=[e for e in L if name(e)=='Shoot' and e['type']=='cast']
        # classify: a Shoot cast event followed by another Shoot cast within 0.7 s is a start; else a fire
        kinds=[]
        for i,e in enumerate(S):
            nxt=S[i+1]['tMs']-e['tMs'] if i+1<len(S) else 9e9
            prv=e['tMs']-S[i-1]['tMs'] if i>0 else 9e9
            if nxt<=700: kinds.append('start'); pair_gaps.append(nxt/1000)
            elif prv<=700: kinds.append('fire')
            else: kinds.append('lone')
        ft=[e['tMs'] for e,k in zip(S,kinds) if k in ('fire','lone')]
        fire_gaps+=[(b-a)/1000 for a,b in zip(ft,ft[1:])]
        casts=[e for e in L if e['type']=='cast' and name(e)!='Shoot']
        chan_end=[e['tMs'] for e in L if name(e) in ('Mind Flay','Drain Life','Drain Soul','Arcane Missiles','Wrack') and e['type']=='debuff_removed']
        for e,k in zip(S,kinds):
            prev=[c for c in casts if c['tMs']<e['tMs']]
            if not prev: continue
            p=prev[-1]; dt=(e['tMs']-p['tMs'])/1000
            # only the first shoot event after that cast
            earlier_shoot=[s for s in S if p['tMs']<s['tMs']<e['tMs']]
            if earlier_shoot: continue
            n=name(p)
            if n in INST:
                (after_inst_start if k=='start' else after_inst_fire).append((n, round(dt,2)))
            elif n in ('Mind Blast','Shadow Bolt','Frostbolt','Fireball','Immolate','Searing Pain','Arcane Blast'):
                (after_hard_start if k=='start' else after_hard_fire).append((n, round(dt,2)))
            if chan_end:
                ce=[c for c in chan_end if p['tMs']<=c<e['tMs']]
                if ce and k in ('fire','lone'): after_chan_fire.append(round((e['tMs']-ce[-1])/1000,2))
    print('=====', f, 'shoot cast events', sum(1 for e in ev if name(e)=='Shoot' and e['type']=='cast'))
    print(' start->fire pair gap median', round(st.median(pair_gaps),2) if pair_gaps else None, 'n', len(pair_gaps))
    fg=sorted(fire_gaps); print(' fire-to-fire gaps hist', sorted(collections.Counter(round(g*4)/4 for g in fg if g<6).items())[:12])
    ais=sorted(after_inst_start,key=lambda x:x[1]); print(' first START after an instant: n',len(ais),'<1.45 s:',sum(1 for x in ais if x[1]<1.45), ais[:12])
    aif=sorted(after_inst_fire,key=lambda x:x[1]); print(' first FIRE after an instant (lone): n',len(aif),'<1.45 s:',sum(1 for x in aif if x[1]<1.45), aif[:12])
    ahs=sorted(after_hard_start,key=lambda x:x[1]); print(' first START after a hardcast event: n',len(ahs), ahs[:10])
    acf=sorted(after_chan_fire); print(' first FIRE after a channel end: n',len(acf), acf[:12])
