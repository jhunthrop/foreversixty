import { chromium } from '/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/rr-frame/web/node_modules/playwright/index.mjs';
const [url, sel, depth='3'] = process.argv.slice(2);
const b = await chromium.launch(); const p = await b.newPage({ viewport: { width: 1440, height: 1000 } });
await p.route('**/v1/**', r => r.fulfill({status:401, contentType:'application/json', body:'{"ok":false,"data":null,"error":null}'}));
await p.goto(url); await p.waitForTimeout(3000);
const out = await p.evaluate(([s,d]) => { const res=[]; const walk=(e,l)=>{ if(l>d) return; const r=e.getBoundingClientRect(); res.push('  '.repeat(l)+e.tagName.toLowerCase()+'.'+[...e.classList].slice(0,3).join('.')+' '+Math.round(r.top)+' h'+Math.round(r.height)); for(const c of e.children) walk(c,l+1);}; walk(document.querySelector(s),0); return res.join('\n'); }, [sel, Number(depth)]);
console.log(out); await b.close();
