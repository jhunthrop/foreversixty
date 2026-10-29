// The character sheet (spec Part B), drawn in HTML at its logical size, 540 x 808, and shot at
// device scale 2 into the real 1080 x 1616 PNG the bot attaches. Every coordinate here is the
// spec's §B.3 table; the Go renderer draws the same boxes.
import { readFileSync } from "node:fs";
import path from "node:path";
import { BOTTOM, LEFT, RIGHT, SLOT_LABELS } from "./data.mjs";

export const SHEET_W = 540;
export const SHEET_H = 808;

const RING = ["#9d9d9d", "#ffffff", "#1eff00", "#0070dd", "#a335ee", "#ff8000"];
const NAME = ["#9d9d9d", "#ffffff", "#1eff00", "#3d94f0", "#b866f5", "#ff8000"];

const b64 = (file, type) => `data:${type};base64,` + readFileSync(file).toString("base64");

export function assets(root, iconDir) {
  const fonts = path.join(root, "web/src/assets/fonts");
  return {
    cinzel: b64(path.join(fonts, "cinzel-700-subset.woff2"), "font/woff2"),
    barlow4: b64(path.join(fonts, "barlow-400-subset.woff2"), "font/woff2"),
    barlow6: b64(path.join(fonts, "barlow-600-subset.woff2"), "font/woff2"),
    crest: (cls) => b64(path.join(root, "web/public/icons/hd/crests", cls + ".webp"), "image/webp"),
    faction: (f) => b64(path.join(root, "web/public/icons/hd/faction", f + ".webp"), "image/webp"),
    seal: b64(path.join(root, "web/public/seal-400.png"), "image/png"),
    icon: (name) => b64(path.join(iconDir, name + ".webp"), "image/webp"),
  };
}

export function fontFaces(a) {
  return `@font-face{font-family:Cinzel;font-weight:700;src:url(${a.cinzel}) format('woff2')}
@font-face{font-family:Barlow;font-weight:400;src:url(${a.barlow4}) format('woff2')}
@font-face{font-family:Barlow;font-weight:600;src:url(${a.barlow6}) format('woff2')}`;
}

const esc = (s) => String(s).replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;");

/** One socket: the icon at 1:1 inside its quality ring; empty and two-handed sockets dim. */
function socket(a, cell) {
  if (cell.kind === "blank" || cell.kind === "empty")
    return `<div class="sock none"></div>`;
  if (cell.kind === "twohand")
    return `<div class="sock none"><img src="${a.icon(cell.icon)}" style="opacity:.35"></div>`;
  return `<div class="sock" style="border-color:${RING[cell.quality] ?? "#262e40"}"><img src="${a.icon(cell.icon)}"></div>`;
}

function textBlock(cell, lines, goldLine = false) {
  if (cell.kind === "blank") return "";
  const name =
    cell.kind === "item"
      ? `<div class="nm" style="color:${NAME[cell.quality] ?? "#e9e4d8"};-webkit-line-clamp:${lines}">${esc(cell.name)}</div>`
      : `<div class="nm dim" style="-webkit-line-clamp:${lines}">${esc(cell.name)}</div>`;
  const gold = cell.gold && !goldLine ? `<b>${esc(cell.gold)}</b>` : "";
  const meta = [cell.meta ? esc(cell.meta) : "", gold].filter(Boolean).join(" · ");
  const own = cell.gold && goldLine ? `<div class="gd">${esc(cell.gold)}</div>` : "";
  return `<div class="tx">${name}${meta ? `<div class="mt">${meta}</div>` : ""}${own}</div>`;
}

/** model: { kicker, title, titleColor, crest, faction, center: html, footer, cells: {slot: cell} } */
export function sheetHtml(a, model) {
  const ROW0 = 88, PITCH = 72;
  const left = LEFT.map((s, i) => {
    const c = model.cells[s];
    return `<div class="cell l" style="top:${ROW0 + PITCH * i}px">${socket(a, c)}${textBlock(c, 2)}</div>`;
  }).join("");
  const right = RIGHT.map((s, i) => {
    const c = model.cells[s];
    return `<div class="cell r" style="top:${ROW0 + PITCH * i}px">${textBlock(c, 2)}${socket(a, c)}</div>`;
  }).join("");
  const bottom = BOTTOM.map((s, i) => {
    const c = model.cells[s];
    return `<div class="cell w" style="left:${16 + 176 * i}px">${socket(a, c)}${textBlock(c, 2, true)}</div>`;
  }).join("");
  return `<!doctype html><meta charset="utf-8"><style>${fontFaces(a)}
*{box-sizing:border-box;margin:0}
body{background:#07090d}
#sheet{position:relative;width:540px;height:808px;overflow:hidden;font-family:Barlow,sans-serif;color:#e9e4d8;
 background:linear-gradient(180deg,rgba(229,185,85,.075),rgba(229,185,85,.012) 45%),#0d0f15;border:1px solid #3a3326}
.kick{position:absolute;left:16px;top:16px;font:600 11px/14px Barlow;letter-spacing:.12em;text-transform:uppercase;color:#e5b955;width:360px;white-space:nowrap;overflow:hidden;text-overflow:ellipsis}
.title{position:absolute;left:16px;top:32px;font:700 22px/28px Cinzel;letter-spacing:.01em;width:360px;white-space:nowrap;overflow:hidden;text-overflow:ellipsis}
.mark{position:absolute;right:16px;top:20px;display:flex;align-items:center;gap:8px;font:400 12px/16px Barlow;color:#9a9484}
.mark img{width:28px;height:28px}
.rule{position:absolute;left:16px;right:16px;top:72px;height:1px;background:#3a3326}
.cell{position:absolute;display:flex;gap:8px;width:210px}
.cell.l{left:16px}.cell.r{right:16px;justify-content:flex-end;text-align:right}
.cell.w{top:680px;width:156px}
.sock{flex:none;width:36px;height:36px;border:2px solid #262e40;border-radius:4px;background:#000;overflow:hidden}
.sock img{display:block;width:32px;height:32px}
.sock.none{background:#0d111a}
.tx{min-width:0;flex:1}
.nm{font:600 20px/23px Barlow;display:-webkit-box;-webkit-box-orient:vertical;overflow:hidden;margin-top:-3px}
.nm.dim{color:#9a9484;font:400 14px/18px Barlow;margin-top:0}
.nm.fit{font-size:17px;line-height:20px}
.gd{font:600 14px/18px Barlow;color:#e5b955}
.mt{font:400 14px/18px Barlow;color:#9a9484;white-space:nowrap;overflow:hidden;text-overflow:ellipsis}
.mt b{color:#e5b955;font-weight:600}
.ctr{position:absolute;left:226px;top:88px;height:576px;justify-content:center;width:88px;display:flex;flex-direction:column;align-items:center;text-align:center}
.ctr .crest{width:64px;height:64px}
.ctr .k{font:600 13px/16px Barlow;color:#f2eee4;margin-top:16px}
.ctr .s{font:400 12px/15px Barlow;color:#9a9484}
.ctr .lab{font:600 10px/12px Barlow;letter-spacing:.12em;text-transform:uppercase;color:#9a9484;margin-top:20px}
.ctr .big{font:700 20px/24px Cinzel;color:#e5b955}
.ctr .fac{width:32px;height:32px;margin-top:22px}
.ctr .fw{font:600 12px/15px Barlow;margin-top:2px}
.foot{position:absolute;left:16px;right:16px;top:780px;font:400 11px/14px Barlow;color:#9a9484;display:flex;justify-content:space-between}
</style><div id="sheet">
<div class="kick">${esc(model.kicker)}</div>
<div class="title" style="color:${model.titleColor}">${esc(model.title)}</div>
<div class="mark"><span>foreversixty.gg</span><img src="${a.seal}"></div>
<div class="rule"></div>
${left}${right}
<div class="ctr"><img class="crest" src="${a.crest(model.crest)}">${model.center}
<img class="fac" src="${a.faction(model.faction)}"><div class="fw" style="color:${model.faction === "horde" ? "#ff6b5c" : "#6fb1ff"}">${model.faction === "horde" ? "Horde" : "Alliance"}</div></div>
${bottom}
<div class="foot"><span>${esc(model.footer)}</span><span>${esc(model.footerRight ?? "")}</span></div>
</div><script>document.fonts.ready.then(()=>{for(const n of document.querySelectorAll('.nm:not(.dim)')){
const lines=Number(getComputedStyle(n).webkitLineClamp)||2;n.style.webkitLineClamp='unset';
if(n.scrollHeight>23*lines+1)n.classList.add('fit');n.style.webkitLineClamp=String(lines);
if(n.scrollHeight>n.clientHeight+1){const w=n.textContent.split(' ');while(w.length>1&&n.scrollHeight>n.clientHeight+1){w.pop();n.textContent=w.join(' ')+'\u2026';}}}
document.body.dataset.fit='1';});</script>`;
}

const item = (it, meta, gold) => ({ kind: "item", icon: it.icon, quality: it.quality, name: it.name, meta, gold });

/** The /bis sheet: each slot's pick and its source line. */
export function bisSheetModel(rows) {
  const cells = { shirt: { kind: "blank" }, tabard: { kind: "blank" } };
  for (const s of Object.keys(SLOT_LABELS)) {
    const r = rows.get(s);
    if (r.empty === "two_hand_equipped") {
      const m = rows.get("main_hand");
      cells[s] = { kind: "twohand", icon: m.icon, name: "Two-handed main hand", meta: "" };
    } else if (r.empty) cells[s] = { kind: "empty", name: "No pick", meta: "" };
    else cells[s] = item(r, r.source);
  }
  return {
    kicker: "Best in slot",
    title: "Beast Mastery Hunter",
    titleColor: "#aad372",
    crest: "hunter",
    faction: "alliance",
    center: `<div class="k">Level 60</div><div class="s">Phase 1 raid</div>`,
    footer: "Updated 29 Oct 2026",
    footerRight: "",
    cells,
  };
}

/** The /char sheet: what the character wears, its item level, the gain where the pick adds. */
export function charSheetModel(crows, avgIlvl) {
  const cells = { shirt: { kind: "blank" }, tabard: { kind: "blank" } };
  for (const s of Object.keys(SLOT_LABELS)) {
    const c = crows.get(s);
    if (c.twoHand) {
      const m = crows.get("main_hand").worn;
      cells[s] = { kind: "twohand", icon: m.icon, name: "Two-handed main hand", meta: "" };
    } else if (c.empty) cells[s] = { kind: "empty", name: "Empty", meta: "", gold: "Upgrade available" };
    else cells[s] = item(c.worn, "item level " + c.worn.ilvl, c.gain);
  }
  return {
    kicker: "Level 60 Beast Mastery Hunter · Olympus XXVII",
    title: "Zulmara",
    titleColor: "#aad372",
    crest: "hunter",
    faction: "alliance",
    center: `<div class="lab">Item level</div><div class="big">${avgIlvl}</div>`,
    footer: "Gold: DPS the level 60 Phase 1 raid pick adds",
    footerRight: "Synced 8 Dec 2026",
    cells,
  };
}
