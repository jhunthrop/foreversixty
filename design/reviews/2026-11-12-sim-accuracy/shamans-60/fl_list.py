import json, sys, time, urllib.request, urllib.parse, urllib.error, datetime
UA='Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0 Safari/537.36'
def get(path, params=None):
    url='https://foreverlogs.gg/api'+path+('?'+urllib.parse.urlencode(params) if params else '')
    for i in range(6):
        try:
            req=urllib.request.Request(url,headers={'User-Agent':UA,'Accept':'application/json','Referer':'https://foreverlogs.gg/'})
            r=json.load(urllib.request.urlopen(req,timeout=90)); time.sleep(0.8); return r
        except urllib.error.HTTPError as e:
            wait=int(e.headers.get('Retry-After') or 0) or 20*(i+1)
            print('retry',e.code,url[:120],wait,file=sys.stderr,flush=True); time.sleep(wait)
        except Exception as e:
            print('retry',url[:120],e,file=sys.stderr,flush=True); time.sleep(10)
    raise SystemExit('failed '+url)
pub=get('/reports/public',{'limit':100})
json.dump(pub,open('fl/public.json','w'))
print(len(pub['reports']), 'reports; keys:', list(pub['reports'][0].keys()))
for r in sorted(pub['reports'], key=lambda x: x['id']):
    print(r['id'], {k:r[k] for k in r if k in ('title','name','createdAt','created_at','startTime','start_time','uploadedAt','uploaded_at','zone','encounterCount','owner','uploader')})
