import json,os,glob
FL='/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/rogues/fl'
def players(min_report=2836):
    out={}
    R=json.load(open(FL+'/rogues_old.json'))
    for k,v in R.items():
        if int(k.split(':')[0])>=min_report: out[k]=v
    for f in glob.glob(FL+'/data/*.json'):
        d=json.load(open(f)); rid=d['report']['id']
        for cid,p in d['players'].items():
            ev=[e for e in d['events'] if e['source']['id']==int(cid)]
            out[f'{rid}:{cid}']={'name':p['name'],'class':p['class'],'ev':ev}
    for k,v in out.items():
        v['ev'].sort(key=lambda e:(e['encounterId'],e['tMs'],int(e['id'])))
    return out
