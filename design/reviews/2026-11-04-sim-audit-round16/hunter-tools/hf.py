import sys,re,html,json,urllib.request,urllib.parse
def q(search):
    u='https://wago.tools/hotfixes?'+urllib.parse.urlencode({'search':search})
    t=urllib.request.urlopen(urllib.request.Request(u,headers={'User-Agent':'Mozilla/5.0'})).read().decode()
    d=json.loads(html.unescape(re.search(r'data-page="([^"]*)"',t).group(1)))
    return d['props']['hotfixes']
for s in sys.argv[1:]:
    h=q(s)
    rows=h['data']
    seen=set()
    for r in rows:
        k=(r['push_id'],r['table_name'],r['record_id'],r['build'],r['status'])
        if k in seen: continue
        seen.add(k); print(s,'|',k, json.dumps(r.get('data'))[:300])
    print(s,'total',h.get('total'))
