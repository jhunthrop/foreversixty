// The character sheet (spec Part B), drawn in HTML at its logical size, 600 x 844, and shot at
// device scale 2 into the real 1200 x 1688 PNG the bot attaches. Every coordinate here is the
// spec's §B.3 table, and the fit script at the bottom is §B.3's fit rule; the Go renderer draws
// the same boxes by the same rule.
import { readFileSync } from "node:fs";
import path from "node:path";
import { SLOT_LABELS } from "./data.mjs";

export const SHEET_W = 600;
export const SHEET_H = 844;
const PAD = 16, ROW0 = 92, PITCH = 72, SOCK = 36, GAP = 8, CENTRE = 72;
const COL = (SHEET_W - 2 * PAD - CENTRE - 2 * GAP) / 2; // 240: socket, gap, 196 of text

// The in-game frame's columns without Shirt and Tabard (spec §B.3): the left column's six sit
// on rows 2 to 7 of the right column's eight, so both centre on the crest.
const LEFT = ["head", "neck", "shoulder", "back", "chest", "wrist"];
const RIGHT = ["hands", "waist", "legs", "feet", "finger1", "finger2", "trinket1", "trinket2"];

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

const fontFaces = (a) => `@font-face{font-family:Cinzel;font-weight:700;src:url(${a.cinzel}) format('woff2')}
@font-face{font-family:Barlow;font-weight:400;src:url(${a.barlow4}) format('woff2')}
@font-face{font-family:Barlow;font-weight:600;src:url(${a.barlow6}) format('woff2')}`;

const esc = (s) => String(s).replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;").replace(/"/g, "&quot;");

function socket(a, c) {
  if (c.kind === "empty") return `<div class="sock none"></div>`;
  if (c.kind === "twohand") return `<div class="sock none"><img src="${a.icon(c.icon)}" style="opacity:.35"></div>`;
  return `<div class="sock" style="border-color:${RING[c.quality] ?? "#262e40"}"><img src="${a.icon(c.icon)}"></div>`;
}

/** The text block: the name, then the meta candidates (longest first) for the fit script, then the gold line. */
function textBlock(c) {
  const quiet = c.kind !== "item";
  const name = `<div class="nm${quiet ? " dim" : ""}" style="${quiet ? "" : "color:" + (NAME[c.quality] ?? "#e9e4d8")}">${esc(c.name)}</div>`;
  const meta = (c.meta ?? []).length ? `<div class="mt" data-alts="${esc(JSON.stringify(c.meta))}"></div>` : "";
  const gold = c.gold ? `<div class="gd">${esc(c.gold)}</div>` : "";
  return `<div class="tx">${name}${meta}${gold}</div>`;
}

const cell = (a, c, side, top, left) =>
  `<div class="cell ${side}" style="top:${top}px;left:${left}px">${side === "r" ? textBlock(c) + socket(a, c) : socket(a, c) + textBlock(c)}</div>`;

/** model: { kicker, title, titleColor, crest, faction, center, footer, footerRight, cells } */
export function sheetHtml(a, model) {
  const rightX = SHEET_W - PAD - COL;
  const left = LEFT.map((s, i) => cell(a, model.cells[s], "l", ROW0 + PITCH * (i + 1), PAD)).join("");
  const right = RIGHT.map((s, i) => cell(a, model.cells[s], "r", ROW0 + PITCH * i, rightX)).join("");
  const weapons =
    cell(a, model.cells.main_hand, "l", ROW0 + PITCH * 8, PAD) +
    cell(a, model.cells.off_hand, "r", ROW0 + PITCH * 8, rightX) +
    cell(a, model.cells.ranged, "l", ROW0 + PITCH * 9, (SHEET_W - COL) / 2);
  return `<!doctype html><meta charset="utf-8"><style>${fontFaces(a)}
*{box-sizing:border-box;margin:0}
body{background:#07090d}
#sheet{position:relative;width:${SHEET_W}px;height:${SHEET_H}px;overflow:hidden;font-family:Barlow,sans-serif;color:#e9e4d8;
 background:linear-gradient(180deg,rgba(229,185,85,.075),rgba(229,185,85,.012) 45%),#0d0f15;border:1px solid #3a3326}
.kick{position:absolute;left:16px;top:16px;font:600 12px/15px Barlow;letter-spacing:.12em;text-transform:uppercase;color:#e5b955;white-space:nowrap}
.title{position:absolute;left:16px;top:34px;font:700 24px/30px Cinzel;letter-spacing:.01em;white-space:nowrap}
.mark{position:absolute;right:16px;top:22px;display:flex;align-items:center;gap:8px;font:400 13px/16px Barlow;color:#9a9484}
.mark img{width:28px;height:28px}
.rule{position:absolute;left:16px;right:16px;top:76px;height:1px;background:#3a3326}
.cell{position:absolute;display:flex;gap:${GAP}px;width:${COL}px;height:64px}
.cell.r{text-align:right}
.sock{flex:none;width:${SOCK}px;height:${SOCK}px;border:2px solid #262e40;border-radius:4px;background:#000;overflow:hidden}
.sock img{display:block;width:32px;height:32px}
.sock.none{background:#0d111a}
.tx{min-width:0;flex:1}
.nm{font:600 20px/24px Barlow;white-space:nowrap;margin-top:-2px}
.nm.two{white-space:normal;font-size:14px;line-height:17px}
.nm.dim{color:#9a9484;font:400 16px/20px Barlow;white-space:normal}
.mt{font:400 16px/20px Barlow;color:#9a9484}
.gd{font:600 16px/20px Barlow;color:#e5b955}
.ctr{position:absolute;left:${(SHEET_W - CENTRE) / 2}px;top:${ROW0}px;width:${CENTRE}px;height:${PITCH * 7 + 64}px;display:flex;flex-direction:column;align-items:center;justify-content:center;text-align:center}
.ctr .crest{width:64px;height:64px}
.ctr .k{font:600 14px/17px Barlow;color:#f2eee4;margin-top:16px}
.ctr .s{font:400 13px/16px Barlow;color:#9a9484}
.ctr .lab{font:600 11px/13px Barlow;letter-spacing:.12em;text-transform:uppercase;color:#9a9484;margin-top:18px}
.ctr .big{font:700 22px/26px Cinzel;color:#e5b955}
.ctr .fac{width:32px;height:32px;margin-top:20px}
.ctr .fw{font:600 13px/16px Barlow;margin-top:2px}
.foot{position:absolute;left:16px;right:16px;top:${SHEET_H - 30}px;font:400 13px/16px Barlow;color:#9a9484;display:flex;justify-content:space-between}
</style><div id="sheet">
<div class="kick">${esc(model.kicker)}</div>
<div class="title" style="color:${model.titleColor}">${esc(model.title)}</div>
<div class="mark"><span>foreversixty.gg</span><img src="${a.seal}"></div>
<div class="rule"></div>
${left}${right}${weapons}
<div class="ctr"><img class="crest" src="${a.crest(model.crest)}">${model.center}
<img class="fac" src="${a.faction(model.faction)}"><div class="fw" style="color:${model.faction === "horde" ? "#ff6b5c" : "#6fb1ff"}">${model.faction === "horde" ? "Horde" : "Alliance"}</div></div>
<div class="foot"><span>${esc(model.footer)}</span><span>${esc(model.footerRight ?? "")}</span></div>
</div><script>
// The fit rule (spec §B.3). A name is one line at 20 px, shrinking 1 px at a time to 14 px;
// a name still wider takes two lines at 14/17. The meta takes the first candidate that fits
// in the lines left (two, or one under a two-line name, or one above a gold line); none fits,
// no meta. Nothing is ever cut with an ellipsis.
document.fonts.ready.then(() => {
  for (const tx of document.querySelectorAll(".tx")) {
    const nm = tx.querySelector(".nm:not(.dim)");
    let nameLines = 1;
    if (nm) {
      let size = 20;
      while (nm.scrollWidth > tx.clientWidth && size > 14) { size--; nm.style.fontSize = size + "px"; }
      if (nm.scrollWidth > tx.clientWidth) { nm.classList.add("two"); nm.style.fontSize = ""; nameLines = 2; }
    }
    const mt = tx.querySelector(".mt");
    if (!mt) continue;
    const lines = (nameLines === 2 ? 1 : 2) - (tx.querySelector(".gd") ? 1 : 0);
    const alts = JSON.parse(mt.dataset.alts);
    let chosen = null;
    for (const alt of alts) {
      mt.textContent = alt;
      if (mt.scrollHeight <= 20 * lines + 1 && mt.scrollWidth <= tx.clientWidth) { chosen = alt; break; }
    }
    if (chosen === null) mt.remove();
  }
  document.body.dataset.fit = "1";
});
</script>`;
}

const item = (it, meta, gold) => ({ kind: "item", icon: it.icon, quality: it.quality, name: it.name, meta, gold });

/** The /bis sheet: each slot's pick and its source line, full form first, then the short form. */
export function bisSheetModel(rows) {
  const cells = {};
  for (const s of Object.keys(SLOT_LABELS)) {
    const r = rows.get(s);
    if (r.empty === "two_hand_equipped") {
      cells[s] = { kind: "twohand", icon: rows.get("main_hand").icon, name: "Two-handed main hand" };
    } else if (r.empty) cells[s] = { kind: "empty", name: "No pick" };
    else cells[s] = item(r, [r.source, r.sourceShort].filter((v, i, l) => v && l.indexOf(v) === i));
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

/** The /char sheet: what the character wears, its item level, the gold gain where the pick adds. */
export function charSheetModel(crows, avgIlvl) {
  const cells = {};
  for (const s of Object.keys(SLOT_LABELS)) {
    const c = crows.get(s);
    if (c.twoHand) cells[s] = { kind: "twohand", icon: crows.get("main_hand").worn.icon, name: "Two-handed main hand" };
    else if (c.empty) cells[s] = { kind: "empty", name: "Empty", gold: "Upgrade available" };
    else cells[s] = item(c.worn, ["item level " + c.worn.ilvl], c.gainShown);
  }
  return {
    kicker: "Level 60 Beast Mastery Hunter · Olympus XXVII",
    title: "Zulmara",
    titleColor: "#aad372",
    crest: "hunter",
    faction: "alliance",
    center: `<div class="lab">Item level</div><div class="big">${avgIlvl}</div>`,
    footer: "Gold: gain over what is worn, level 60 Phase 1 raid",
    footerRight: "Synced 8 Dec 2026",
    cells,
  };
}
