import pyarrow.parquet as pq,glob,collections,sys
KEY={'Arcane Blast','Arcane Missiles','Arcane Wrath','Arcane Power','Presence of Mind','Evocation','Mind Flay','Shadowform','Vampiric Embrace','Mind Blast','Shadow Word: Pain','Devouring Plague','Shadow Word: Death','Arcane Explosion'}
cnt=collections.defaultdict(collections.Counter); lvl=collections.defaultdict(int); names={}
for f in glob.glob('dl/*/f*.parquet'):
    rep=f.split('/')[1]
    t=pq.read_table(f,columns=['event','source_guid','source_name','spell_name','adv_level','adv_info_guid']).to_pandas()
    t=t[t.source_guid.str.startswith('Player')]
    for (g,n,s),c in t[t.event=='SPELL_CAST_SUCCESS'].groupby(['source_guid','source_name','spell_name']).size().items():
        if s in KEY: cnt[(rep,g)][s]+=c; names[(rep,g)]=n
    a=t[(t.adv_info_guid==t.source_guid)&t.adv_level.notna()].groupby('source_guid').adv_level.max()
    for g,l in a.items(): lvl[(rep,g)]=max(lvl[(rep,g)],int(l))
for k,c in sorted(cnt.items(), key=lambda kv:-lvl[kv[0]]):
    if not ({'Arcane Blast','Arcane Missiles','Mind Flay','Shadowform','Mind Blast'} & set(c)): continue
    print(k[0], names[k][:12], lvl[k], dict(c))
