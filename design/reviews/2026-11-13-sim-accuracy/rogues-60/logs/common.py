import pandas as pd, json
df=pd.read_parquet('rogue_events.parquet')
rog={}
for rep in df.rep.unique():
    for x in json.load(open(f'dl/{rep}/report.json'))['units']:
        if x.get('class')=='Rogue': rog[x['guid']]=x
s=df[df.source_guid.isin(rog)].copy()
lv=s[s.adv_present&(s.adv_info_guid==s.source_guid)].groupby('source_guid').adv_level.max()
s['lvl']=s.source_guid.map(lv)
IP={8680,8685,8689,11335,11336,11337}
DP={2818,2819,11353,11354,25349}
import csv as _csv
SL={}
for _r in _csv.DictReader(open('/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/wago/SpellLevels_70888h.csv')):
    if _r['DifficultyID']=='0': SL[int(_r['SpellID'])]=int(_r['SpellLevel'])
