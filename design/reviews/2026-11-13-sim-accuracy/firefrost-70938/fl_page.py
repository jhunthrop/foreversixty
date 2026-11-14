import json, sys, time, urllib.parse, urllib.request
rid, cid, out = sys.argv[1], sys.argv[2], sys.argv[3]
evs=[]; cur=None
while True:
    p=f"https://foreverlogs.gg/api/reports/{rid}/meta_events?sourceCharacterIds={cid}&limit=2000"
    if cur: p+="&cursor="+urllib.parse.quote(str(cur))
    req=urllib.request.Request(p, headers={"User-Agent":"Mozilla/5.0","Referer":"https://foreverlogs.gg/"})
    d=json.load(urllib.request.urlopen(req, timeout=60)); evs+=d["events"]
    if not d.get("hasMore"): break
    cur=d["nextCursor"]; time.sleep(0.3)
json.dump(evs, open(out,"w")); print(len(evs))
