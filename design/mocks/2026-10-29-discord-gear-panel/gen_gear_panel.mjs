// Boards for design/specs/2026-10-29-discord-gear-panel.md: the /bis and /char answers as one
// Components V2 panel, and Share in channel posting the character sheet PNG. Mock only.
// The reviewers' render kit (web/tests/captures/discord-bot/discord-render.ts) draws legacy
// embeds only, so this file carries its own Components V2 drawing in the kit's dark-theme
// tokens: the container with its accent bar, sections with their right-hand thumbnail,
// text displays with Discord's markdown (### heading, -# subtext, bold, masked links,
// <t:…:R>), separators, the media gallery and the button and select rows.
// Data is real (data.mjs: the active build's Beast Mastery Hunter band, level 60 Phase 1 raid,
// Alliance; item icons from data/builds/1.60.1.70658/icons; the HD crest and faction mark from
// web/public/icons/hd). Payloads land in payloads/, the sheets at their real size, the boards
// at the Discord desktop-client (640) and phone-client (390) widths, as the kit does.
// Run with Node 22 or later from anywhere: node design/mocks/2026-10-29-discord-gear-panel/gen_gear_panel.mjs
import { writeFileSync } from "node:fs";
import path from "node:path";
import { fileURLToPath, pathToFileURL } from "node:url";
import { BUILD, bisRows, loadData } from "./data.mjs";
import {
  API, SITE, T, bisPanel, bisPanelShareFailed, bisShare, bisShareNoSheet, bisSlotPanel, charPanel, charShare,
  countComponents, textLength, upgradesOf, averageItemLevel,
} from "./payloads.mjs";
import { SHEET_H, SHEET_W, assets, bisSheetModel, charSheetModel, sheetHtml } from "./sheet.mjs";

const HERE = path.dirname(fileURLToPath(import.meta.url));
const ROOT = path.resolve(HERE, "../../..");
const { chromium } = await import(pathToFileURL(path.join(ROOT, "web/node_modules/playwright/index.mjs")).href);

const CLOCK = new Date("2026-12-10T22:00:00Z");
const CLIENTS = [
  { name: "desktop", width: 640 },
  { name: "phone", width: 390 },
];
const MAX_COMPONENTS = 38; // the spec's own ceiling (§A.6): 40 less two in hand
const TEXT_BUDGET = 3800; // our guard under Discord's 4000 (spec §A.6)

const d = loadData(ROOT);
const a = assets(ROOT, d.iconDir);

// ---- payloads, with the budget asserted ----
const frames = [
  { id: "bis-panel-p1", title: "/bis: Head to Legs", usedCommand: null, message: bisPanel(d, 1) },
  { id: "bis-panel-p2", title: "/bis: Feet to Ranged", message: bisPanel(d, 2) },
  { id: "bis-panel-p1-sharefail", title: "/bis page 1 after a failed share: the failure line, Share live again", message: bisPanelShareFailed(d) },
  { id: "bis-panel-slot", title: "/bis slot:legs", message: bisSlotPanel(d, "legs") },
  { id: "char-panel", title: "/char, linked member: top upgrades", message: charPanel(d) },
  { id: "char-panel-nogear", title: "/char, no gear shared", message: charPanel(d, { noGear: true }) },
  { id: "share-bis", title: "Share in channel: /bis with the sheet", usedCommand: "Zulmara used /bis", sheet: "bis", message: bisShare(d) },
  { id: "share-char", title: "Share in channel: /char with the sheet", usedCommand: "Zulmara used /char", sheet: "char", message: charShare(d) },
  { id: "share-bis-nosheet", title: "Share in channel: the sheet failed to render", usedCommand: "Zulmara used /bis", message: bisShareNoSheet(d) },
];
for (const f of frames) {
  const n = countComponents(f.message.components);
  const t = textLength(f.message.components);
  if (n > MAX_COMPONENTS) throw new Error(`${f.id}: ${n} components, over ${MAX_COMPONENTS}`);
  if (t > TEXT_BUDGET) throw new Error(`${f.id}: ${t} text characters, over ${TEXT_BUDGET}`);
  f.budget = { components: n, text: t };
  writeFileSync(path.join(HERE, "payloads", f.id + ".json"), JSON.stringify({ id: f.id, title: f.title, budget: f.budget, message: f.message }, null, 2) + "\n");
  console.log(f.id.padEnd(20), "components", String(n).padStart(2), "text", t);
}

// ---- the V2 drawing ----
const esc = (s) => String(s).replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;").replace(/"/g, "&quot;");
const rel = (unix) => {
  const days = Math.round((CLOCK.getTime() / 1000 - unix) / 86400);
  return days <= 1 ? "a day ago" : `${days} days ago`;
};
function inline(raw) {
  return esc(raw)
    .replace(/&lt;:fs_q_(\w+):\d+&gt; ?/g, (_m, q) => `<span class="qd"><i style="background:${QUALITY_DOT[q]}"></i></span>`)
    .replace(/&lt;t:(\d+):R&gt;/g, (_m, u) => `<span class="ts">${rel(Number(u))}</span>`)
    .replace(/\[([^\]]+)\]\((https?:\/\/[^)\s]+)\)/g, '<a class="ml">$1</a>')
    .replace(/\*\*(.+?)\*\*/g, "<strong>$1</strong>");
}
function markdown(content) {
  return content
    .split("\n")
    .map((line) => {
      if (line.startsWith("### ")) return `<div class="h3">${inline(line.slice(4))}</div>`;
      if (line.startsWith("-# ")) return `<div class="sub">${inline(line.slice(3))}</div>`;
      return `<div>${inline(line)}</div>`;
    })
    .join("");
}

const QUALITY_DOT = { poor: "#9d9d9d", common: "#ffffff", uncommon: "#1eff00", rare: "#0070dd", epic: "#a335ee", legendary: "#ff8000" };
const QUALITY_RING = ["#9d9d9d", "#ffffff", "#1eff00", "#0070dd", "#a335ee", "#ff8000"];
const TILE = new RegExp("^" + API.replace(/\./g, "\\.") + "/v1/discord/tiles/" + BUILD.replace(/\./g, "\\.") + "/(\\d+)\\.png$");
const itemById = (id) => d.items.get(id);

/** The tile endpoint's image (spec §A.4): the icon at 1:1 inside a 4 px quality ring, 72 x 72. */
function thumbHtml(c) {
  const m = TILE.exec(c.media.url);
  if (m) {
    const it = itemById(Number(m[1]));
    return `<div class="thumb tile" style="background:${QUALITY_RING[it.quality]}"><img src="${a.icon(it.icon)}" alt=""></div>`;
  }
  const crest = /\/icons\/hd\/crests\/(\w+)\.webp$/.exec(c.media.url);
  return `<img class="thumb" src="${crest ? a.crest(crest[1]) : ""}" alt="">`;
}

const BUTTON = { 2: "secondary", 3: "success", 5: "link" };
function component(c, sheets) {
  switch (c.type) {
    case T.CONTAINER: {
      const bar = c.accent_color != null ? "#" + c.accent_color.toString(16).padStart(6, "0") : "transparent";
      return `<div class="ctn" style="border-left-color:${bar}">${c.components.map((k) => component(k, sheets)).join("")}</div>`;
    }
    case T.SECTION:
      return `<div class="sec"><div class="stx">${c.components.map((k) => component(k, sheets)).join("")}</div>${thumbHtml(c.accessory)}</div>`;
    case T.TEXT:
      return `<div class="td">${markdown(c.content)}</div>`;
    case T.SEP:
      return `<div class="sep s${c.spacing ?? 1}${c.divider === false ? " nodiv" : ""}"></div>`;
    case T.GALLERY:
      return `<div class="gal">${c.items.map((i) => `<img src="${sheets[i.media.url] ?? ""}" alt="">`).join("")}</div>`;
    case T.ROW:
      return `<div class="row">${c.components.map((k) => component(k, sheets)).join("")}</div>`;
    case T.BUTTON:
      return `<button class="btn ${BUTTON[c.style] ?? "secondary"}" ${c.disabled ? "disabled" : ""}>${esc(c.label)}${c.style === 5 ? '<span class="ext">↗</span>' : ""}</button>`;
    case T.SELECT: {
      const chosen = (c.options ?? []).find((o) => o.default);
      return `<div class="select"><span class="${chosen ? "chosen" : ""}">${esc(chosen ? chosen.label : c.placeholder)}</span><span>⌄</span></div>`;
    }
  }
  return "";
}

const clockText = new Intl.DateTimeFormat("en-US", { timeZone: "America/Chicago", hour: "numeric", minute: "2-digit" }).format(CLOCK);

function page(f, width, sheets) {
  const m = f.message;
  const ephemeral = (m.flags & 64) !== 0;
  const used = f.usedCommand ? `<div class="used">↱ <b>${esc(f.usedCommand)}</b></div>` : "";
  return `<!doctype html><meta charset="utf-8"><style>
body{margin:0;background:#313338;color:#dbdee1;font:16px/1.375 "gg sans","Segoe UI",Helvetica,Arial,sans-serif}
.wrap{box-sizing:border-box;width:${width}px;padding:16px}
.used{font-size:13px;color:#949ba4;margin:0 0 4px 52px}.used b{color:#c4c9ce;font-weight:500}
.msg{display:flex;gap:12px}
.avatar{flex:none;width:40px;height:40px;border-radius:50%;background:#e5b955;color:#07090d;font-weight:700;display:grid;place-items:center}
.main{min-width:0;flex:1}
.head b{color:#f2f3f5}.app{background:#5865f2;color:#fff;font-size:10px;padding:1px 5px;border-radius:3px;margin:0 4px}
.head small{color:#949ba4;font-size:12px}
.ctn{box-sizing:border-box;max-width:560px;margin-top:4px;background:#2b2d31;border:1px solid #3f4147;border-left:4px solid;border-radius:8px;padding:16px;display:flex;flex-direction:column;gap:8px}
.sec{display:flex;gap:12px;align-items:flex-start}
.stx{min-width:0;flex:1;display:flex;flex-direction:column;gap:8px}
.thumb{flex:none;width:${width < 500 ? 64 : 72}px;height:${width < 500 ? 64 : 72}px;border-radius:8px;object-fit:contain}
.thumb.tile{box-sizing:border-box;padding:${width < 500 ? 3.5 : 4}px;display:block}
.thumb.tile img{display:block;width:100%;height:100%;border-radius:2px}
.td{font-size:16px;line-height:22px;overflow-wrap:anywhere}
.td .h3{font-size:20px;line-height:26px;font-weight:700;color:#f2f3f5;margin-bottom:2px}
.td .sub{font-size:13px;line-height:18px;color:#949ba4}
.td strong{font-weight:700}
.ml{color:#00a8fc;text-decoration:none}
.qd{display:inline-flex;width:22px;height:22px;align-items:center;justify-content:center;vertical-align:-5px;margin-right:4px}
.qd i{display:block;width:12px;height:12px;border-radius:50%;box-shadow:0 0 0 1px rgba(0,0,0,.45)}
.ts{background:#4e5058;border-radius:3px;padding:0 2px}
.sep{border-top:1px solid #3f4147;margin:0}.sep.s2{margin:8px 0}.sep.nodiv{border-top:0}
.gal img{display:block;width:100%;border-radius:8px}
.row{display:flex;flex-wrap:wrap;gap:8px;margin-top:8px}
.btn{font:500 14px/1 inherit;font-family:inherit;border:0;border-radius:8px;padding:10px 16px;color:#fff;min-height:32px}
.btn.secondary,.btn.link{background:#4e5058}.btn.success{background:#248046}.btn[disabled]{opacity:.5}.ext{margin-left:6px}
.select{display:flex;justify-content:space-between;align-items:center;background:#1e1f22;border:1px solid #1e1f22;border-radius:8px;padding:10px 12px;width:100%;max-width:400px;box-sizing:border-box;color:#b5bac1}
.select .chosen{color:#dbdee1}
.only{margin-top:6px;font-size:12px;color:#949ba4}.only a{color:#00a8fc}
</style><div class="wrap" id="msg">${used}<div class="msg"><div class="avatar">FS</div><div class="main">
<div class="head"><b>Forever Sixty</b><span class="app">APP</span><small>Today at ${clockText}</small></div>
${m.components.map((c) => component(c, sheets)).join("")}
${ephemeral ? '<div class="only">Only you can see this • <a href="#">Dismiss message</a></div>' : ""}
</div></div></div>`;
}

// ---- render ----
const browser = await chromium.launch();

// The sheets first, at their real size: 540 x 808 logical, device scale 2, 1080 x 1616 PNG.
const sheetPage = await browser.newPage({ deviceScaleFactor: 2, viewport: { width: SHEET_W, height: SHEET_H } });
const { crows } = upgradesOf(d);
const sheetModels = {
  bis: bisSheetModel(bisRows(d)),
  char: charSheetModel(crows, averageItemLevel(d, crows)),
};
const sheetFiles = {};
for (const [key, model] of Object.entries(sheetModels)) {
  const html = sheetHtml(a, model);
  await sheetPage.setContent(html, { waitUntil: "load" });
  await sheetPage.waitForSelector("body[data-fit]");
  const file = path.join(HERE, `sheet-${key}.png`);
  await sheetPage.locator("#sheet").screenshot({ path: file });
  sheetFiles[key] = file;
}
await sheetPage.close();

const { readFileSync } = await import("node:fs");
const dataPng = (file) => "data:image/png;base64," + readFileSync(file).toString("base64");

const boardPage = await browser.newPage({ deviceScaleFactor: 2 });
for (const client of CLIENTS)
  for (const f of frames) {
    const sheets = f.sheet ? { "attachment://sheet.png": dataPng(sheetFiles[f.sheet]) } : {};
    await boardPage.setViewportSize({ width: client.width, height: 800 });
    await boardPage.setContent(page(f, client.width, sheets), { waitUntil: "load" });
    await boardPage.locator("#msg").screenshot({ path: path.join(HERE, `${f.id}-${client.name}.png`) });
  }
await browser.close();
console.log("wrote", frames.length * CLIENTS.length, "boards and 2 sheets to", HERE, "for", SITE);
