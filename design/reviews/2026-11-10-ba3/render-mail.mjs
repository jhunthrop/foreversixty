// Captures for lane BA3: every email the server sends (dumped by `go test -tags payloaddump
// -run 'TestDumpEmails|TestDumpConfirmationEmail' ./internal/auth/ ./internal/subscribe/` with
// REVIEW_PAYLOADS_DIR=design/reviews/2026-11-10-ba3), drawn by a browser at a desktop mail
// client's reading width and at a phone's, plus one board of every inbox line (subject and
// preheader) as a client lists them.
//
// Run with Node 22+: node design/reviews/2026-11-10-ba3/render-mail.mjs <web/node_modules>
import { mkdirSync, readdirSync, readFileSync } from "node:fs";
import path from "node:path";
import { fileURLToPath, pathToFileURL } from "node:url";

const HERE = path.dirname(fileURLToPath(import.meta.url));
const MAIL = path.join(HERE, "mail");
const OUT = path.join(HERE, "captures");
const NODE_MODULES = process.argv[2] ?? path.join(HERE, "../../../web/node_modules");
const { chromium } = await import(pathToFileURL(path.join(NODE_MODULES, "playwright/index.mjs")).href);

const WIDTHS = [
  { name: "desktop", width: 1440 },
  { name: "phone", width: 390 },
];

mkdirSync(OUT, { recursive: true });
const names = readdirSync(MAIL).filter((f) => f.endsWith(".html")).map((f) => f.slice(0, -5)).sort();
const browser = await chromium.launch();
for (const { name: client, width } of WIDTHS) {
  const page = await browser.newPage({ viewport: { width, height: 900 }, deviceScaleFactor: 2 });
  for (const name of names) {
    await page.goto(pathToFileURL(path.join(MAIL, `${name}.html`)).href, { waitUntil: "networkidle" });
    await page.screenshot({ path: path.join(OUT, `mail-${name}-${client}.png`), fullPage: true });
  }
  await page.close();
}

// The inbox board: sender, subject and preheader, the way a client's list shows a message.
const rows = names.map((name) => {
  const html = readFileSync(path.join(MAIL, `${name}.html`), "utf8");
  const subject = /<title>([^<]*)<\/title>/.exec(html)?.[1] ?? "";
  const preheader = /<div style="display:none[^>]*>([^<]*)<\/div>/.exec(html)?.[1] ?? "";
  return `<tr><td class="from">Forever Sixty</td><td><b>${subject}</b> <span>${preheader}</span></td></tr>`;
});
const inbox = await browser.newPage({ viewport: { width: 1100, height: 400 }, deviceScaleFactor: 2 });
await inbox.setContent(`<!doctype html><meta charset="utf-8"><style>
body{margin:0;font:14px/1.4 Arial,sans-serif;background:#fff;color:#202124}
table{border-collapse:collapse;width:100%}td{padding:10px 16px;border-bottom:1px solid #e0e0e0;white-space:nowrap;overflow:hidden;text-overflow:ellipsis;max-width:820px}
.from{width:160px;font-weight:700}span{color:#5f6368}</style><table>${rows.join("")}</table>`);
await inbox.screenshot({ path: path.join(OUT, "mail-inbox.png"), fullPage: true });

// The push examples (notify's TestDumpPushExamples writes push.json): each as a phone's lock
// screen stacks a notification, the app name, then the title and the body.
const pushes = JSON.parse(readFileSync(path.join(HERE, "push.json"), "utf8"));
const SEAL = `data:image/png;base64,${readFileSync(path.join(HERE, "../../../web/public/seal-400.png")).toString("base64")}`;
const escape = (s) => s.replace(/&/g, "&amp;").replace(/</g, "&lt;");
const cards = pushes.map((p) => `<div class="n"><div class="app"><img src="${SEAL}">FOREVER SIXTY<span>now</span></div>
<b>${escape(p.Title)}</b><p>${escape(p.Body)}</p></div>`);
const lock = await browser.newPage({ viewport: { width: 390, height: 600 }, deviceScaleFactor: 2 });
await lock.setContent(`<!doctype html><meta charset="utf-8"><style>
body{margin:0;padding:24px 10px;background:linear-gradient(#2b3a4a,#14181f);font:15px/1.3 -apple-system,Helvetica,sans-serif;color:#fff}
.n{background:rgba(245,245,245,.18);border-radius:20px;padding:12px 14px;margin-bottom:10px}
.app{display:flex;align-items:center;gap:6px;font-size:12px;color:rgba(255,255,255,.7)}.app img{width:20px;height:20px;border-radius:5px}
.app span{margin-left:auto}b{display:block;margin-top:6px}p{margin:2px 0 0}</style>${cards.join("")}`);
await lock.screenshot({ path: path.join(OUT, "push-lockscreen.png"), fullPage: true });
await browser.close();
console.log(`captured ${names.length} emails in ${OUT}`);
