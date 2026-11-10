// Review boards for lane f-discord (design/specs/2026-11-09-directions.md, round 5): every
// Discord-visible message the bot and the webhooks send, drawn from the real payloads the Go
// handlers produce (api/internal/discordbot/review_dump_test.go writes them to
// design/reviews/2026-11-09-f-discord/payloads), in a faithful copy of Discord's dark theme: the
// message header with the seal avatar, legacy embeds (colour bar, author, title, fields,
// thumbnail, footer icon and time) and Components V2 (container with its accent bar, sections
// with their thumbnail, text displays, separators, the media gallery), button and select rows,
// application emoji and <t:...> stamps. It draws what the payload says and nothing else; it is
// not the real client. Images are the repo's own files: web/public for the site, data/builds for
// the item icons, the dumped tiles and quality dots, and the rendered sheet PNGs for
// attachment://sheet.png.
//
// Run with Node 22+: node design/mocks/2026-11-09-f-discord/render.mjs <web/node_modules>
// It writes discord/<id>-desktop.png (640) and discord/<id>-phone.png (390) beside the payloads.
import { existsSync, mkdirSync, readFileSync, readdirSync } from "node:fs";
import path from "node:path";
import { fileURLToPath, pathToFileURL } from "node:url";

const HERE = path.dirname(fileURLToPath(import.meta.url));
const ROOT = path.resolve(HERE, "../../..");
const REVIEW = path.join(ROOT, "design/reviews/2026-11-09-f-discord");
const PAYLOADS = path.join(REVIEW, "payloads");
const OUT = path.join(REVIEW, "discord");
const NODE_MODULES = process.argv[2] ?? path.join(ROOT, "web/node_modules");
const { chromium } = await import(pathToFileURL(path.join(NODE_MODULES, "playwright/index.mjs")).href);

const CLIENTS = [
  { name: "desktop", width: 640 },
  { name: "phone", width: 390 },
];
const GEAR_BUILD = "1.60.1.70658";
const ZONE = "America/Chicago"; // the reader sits in the guild's raid zone
const EMOJI = JSON.parse(readFileSync(path.join(PAYLOADS, "emoji.json"), "utf8"));
const SHEETS = { "share-char": "sheet-char.png", "share-bis": "sheet-bis.png" };

// ---- images -------------------------------------------------------------------------------

/** The local file an image URL in a payload stands for. */
function localFile(raw) {
  const url = new URL(raw);
  const p = decodeURIComponent(url.pathname);
  const tile = /^\/v1\/discord\/tiles\/[^/]+\/(\d+)\.png$/.exec(p);
  if (tile) return path.join(PAYLOADS, "tiles", `${tile[1]}.png`);
  if (/^\/v1\/guilds\/\d+\/crest/.test(p)) return path.join(ROOT, "web/public/icons/hd/crests/warrior.webp");
  const icon = /^\/data\/([^/]+)\/icons\/(.+)$/.exec(p);
  if (icon) {
    const own = path.join(ROOT, "data/builds", icon[1], "icons", icon[2]);
    return existsSync(own) ? own : path.join(ROOT, "data/builds", GEAR_BUILD, "icons", icon[2]);
  }
  return path.join(ROOT, "web/public", p);
}

const MIME = { ".png": "image/png", ".webp": "image/webp", ".jpg": "image/jpeg", ".svg": "image/svg+xml" };

function dataUri(file) {
  if (!existsSync(file)) return null;
  return `data:${MIME[path.extname(file)] ?? "image/png"};base64,${readFileSync(file).toString("base64")}`;
}

/** An image URL as a data URI, or null when the repo has no file for it. */
function image(url, frameId) {
  if (url.startsWith("attachment://")) return dataUri(path.join(REVIEW, SHEETS[frameId] ?? "sheet-char.png"));
  return dataUri(localFile(url));
}

function emojiImage(name) {
  const src = EMOJI[name];
  if (!src) return null;
  if (src.startsWith("icon:")) return dataUri(path.join(ROOT, "data/builds", GEAR_BUILD, "icons", `${src.slice(5)}.webp`));
  return dataUri(path.join(PAYLOADS, src));
}

// ---- markdown -----------------------------------------------------------------------------

const esc = (s) => s.replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;").replace(/"/g, "&quot;");

function stamp(seconds, style, clock) {
  const at = new Date(seconds * 1000);
  if (style === "R") {
    const diff = (at - clock) / 1000;
    const units = [["year", 31536000], ["month", 2592000], ["day", 86400], ["hour", 3600], ["minute", 60]];
    const rtf = new Intl.RelativeTimeFormat("en", { numeric: "auto" });
    for (const [unit, s] of units) if (Math.abs(diff) >= s) return rtf.format(Math.round(diff / s), unit);
    return rtf.format(Math.round(diff), "second");
  }
  const opts = { timeZone: ZONE, weekday: "long", year: "numeric", month: "long", day: "numeric", hour: "numeric", minute: "2-digit" };
  if (style === "f") delete opts.weekday;
  if (style === "t") return at.toLocaleTimeString("en-US", { timeZone: ZONE, hour: "numeric", minute: "2-digit" });
  if (style === "D") return at.toLocaleDateString("en-US", { timeZone: ZONE, year: "numeric", month: "long", day: "numeric" });
  return at.toLocaleString("en-US", opts).replace(" at ", " ");
}

/** Discord markdown, the subset the bot writes, as HTML. */
function md(text, clock) {
  let s = esc(text ?? "");
  s = s.replace(/&lt;t:(\d+)(?::([tTdDfFR]))?&gt;/g, (_, n, st) => `<span class="stamp">${stamp(Number(n), st ?? "f", clock)}</span>`);
  s = s.replace(/&lt;a?:([a-z0-9_]+):\d+&gt;/gi, (_, name) => {
    const src = emojiImage(name);
    return src ? `<img class="emoji" src="${src}" alt=":${name}:">` : `:${name}:`;
  });
  s = s.replace(/\[([^\]]+)\]\((https?:[^)\s]+)\)/g, '<a href="#">$1</a>');
  s = s.replace(/\*\*([^*]+)\*\*/g, "<strong>$1</strong>");
  s = s.replace(/(^|[\s(])\*([^*\n]+)\*/g, "$1<em>$2</em>");
  s = s.replace(/`([^`]+)`/g, "<code>$1</code>");
  return s
    .split("\n")
    .map((line) => {
      if (line.startsWith("### ")) return `<div class="h3">${line.slice(4)}</div>`;
      if (line.startsWith("## ")) return `<div class="h2">${line.slice(3)}</div>`;
      if (line.startsWith("-# ")) return `<div class="sub">${line.slice(3)}</div>`;
      if (/^[-•] /.test(line)) return `<div class="li">${line.slice(2)}</div>`;
      return `<div class="ln">${line || "&nbsp;"}</div>`;
    })
    .join("");
}

// ---- components -----------------------------------------------------------------------------

const BUTTON_STYLE = { 1: "primary", 2: "secondary", 3: "success", 4: "danger", 5: "link" };

function emojiOf(c) {
  if (!c.emoji) return "";
  const src = c.emoji.id ? emojiImage(c.emoji.name) : null;
  return src ? `<img class="emoji" src="${src}">` : esc(c.emoji.name ?? "");
}

function rowHtml(row) {
  const parts = row.components.map((c) => {
    if (c.type === 2) {
      const ext = c.style === 5 ? '<span class="ext">↗</span>' : "";
      return `<span class="btn ${BUTTON_STYLE[c.style] ?? "secondary"}${c.disabled ? " disabled" : ""}">${emojiOf(c)}${esc(c.label ?? "")}${ext}</span>`;
    }
    if (c.type === 3) {
      const chosen = (c.options ?? []).filter((o) => o.default).map((o) => o.label);
      const text = chosen.length ? chosen.join(", ") : c.placeholder ?? "Make a selection";
      return `<span class="select${chosen.length ? " chosen" : ""}"><span>${esc(text)}</span><span class="chev">⌄</span></span>`;
    }
    return "";
  });
  return `<div class="row">${parts.join("")}</div>`;
}

function nodeHtml(n, f) {
  switch (n.type) {
    case 10:
      return `<div class="td">${md(n.content, f.clock)}</div>`;
    case 9: {
      const acc = n.accessory;
      let side = "";
      if (acc?.type === 11) {
        const src = image(acc.media.url, f.id);
        side = src ? `<img class="thumb" src="${src}">` : '<div class="thumb missing"></div>';
      } else if (acc?.type === 2) side = rowHtml({ components: [acc] });
      return `<div class="section"><div class="sbody">${n.components.map((c) => nodeHtml(c, f)).join("")}</div>${side}</div>`;
    }
    case 14:
      return `<div class="sep${n.spacing === 2 ? " big" : ""}${n.divider === false ? " nodiv" : ""}"></div>`;
    case 12:
      return `<div class="gallery">${n.items
        .map((it) => {
          const src = image(it.media.url, f.id);
          return src ? `<img src="${src}">` : '<div class="thumb missing"></div>';
        })
        .join("")}</div>`;
    case 17: {
      const bar = n.accent_color != null ? `#${n.accent_color.toString(16).padStart(6, "0")}` : "transparent";
      return `<div class="container" style="border-left-color:${bar}">${n.components.map((c) => nodeHtml(c, f)).join("")}</div>`;
    }
    case 1:
      return rowHtml(n);
    default:
      return "";
  }
}

// ---- embeds -------------------------------------------------------------------------------

function footerTime(iso, clock) {
  const at = new Date(iso);
  const t = at.toLocaleTimeString("en-US", { timeZone: ZONE, hour: "numeric", minute: "2-digit" });
  const day = (d) => d.toLocaleDateString("en-US", { timeZone: ZONE });
  return day(at) === day(clock) ? `Today at ${t}` : `${day(at)} ${t}`;
}

function embedHtml(e, f) {
  const bar = e.color ? `#${e.color.toString(16).padStart(6, "0")}` : "#1e1f22";
  const icon = (url, cls) => {
    const src = url ? image(url, f.id) : null;
    return src ? `<img class="${cls}" src="${src}">` : "";
  };
  const author = e.author ? `<div class="author">${icon(e.author.icon_url, "aicon")}${md(e.author.name, f.clock)}</div>` : "";
  const title = e.title ? `<div class="etitle${e.url ? " link" : ""}">${md(e.title, f.clock)}</div>` : "";
  const desc = e.description ? `<div class="edesc">${md(e.description, f.clock)}</div>` : "";
  const fields = (e.fields ?? [])
    .map((fl) => `<div class="field${fl.inline ? " inline" : ""}"><div class="fname">${md(fl.name, f.clock)}</div><div class="fval">${md(fl.value, f.clock)}</div></div>`)
    .join("");
  const thumb = e.thumbnail ? icon(e.thumbnail.url, "ethumb") : "";
  const img = e.image ? icon(e.image.url, "eimage") : "";
  const time = e.timestamp ? footerTime(e.timestamp, f.clock) : "";
  const footer =
    e.footer || time
      ? `<div class="footer">${e.footer ? icon(e.footer.icon_url, "ficon") : ""}<span>${e.footer ? md(e.footer.text, f.clock).replace(/<\/?div[^>]*>/g, "") : ""}${e.footer && time ? " • " : ""}${time}</span></div>`
      : "";
  return `<div class="embed" style="border-left-color:${bar}"><div class="egrid"><div class="emain">${author}${title}${desc}<div class="fields">${fields}</div></div>${thumb}</div>${img}${footer}</div>`;
}

// ---- the message ----------------------------------------------------------------------------

function messageHtml(m, f, { webhook = false } = {}) {
  const v2 = ((m.flags ?? 0) & 32768) !== 0;
  const ephemeral = ((m.flags ?? 0) & 64) !== 0;
  const seal = dataUri(path.join(ROOT, "web/public/seal-400.png"));
  const name = m.username ?? "Forever Sixty";
  const badge = webhook || m.username ? "" : '<span class="app">APP</span>';
  const time = footerTime(f.clock.toISOString(), f.clock);
  const body = v2
    ? (m.components ?? []).map((c) => nodeHtml(c, f)).join("")
    : `${m.content ? `<div class="content">${md(m.content, f.clock)}</div>` : ""}${(m.embeds ?? []).map((e) => embedHtml(e, f)).join("")}${(m.components ?? []).map(rowHtml).join("")}`;
  const eph = ephemeral ? '<div class="eph">👁 Only you can see this • <a href="#">Dismiss message</a></div>' : "";
  return `<div class="msg"><img class="avatar" src="${seal}"><div class="main"><div class="head"><span class="uname">${esc(name)}</span>${badge}<span class="when">${time}</span></div>${body}${eph}</div></div>`;
}

function modalHtml(mo) {
  const inputs = (mo.components ?? [])
    .flatMap((r) => r.components ?? [])
    .map((c) => `<div class="minput"><div class="mlabel">${esc(c.label ?? "")}${c.required ? ' <span class="req">*</span>' : ""}</div><div class="mfield">${esc(c.placeholder ?? c.value ?? "")}</div></div>`)
    .join("");
  return `<div class="modal"><div class="mtitle">${esc(mo.title ?? "")}</div>${inputs}<div class="mfoot"><span class="btn link-plain">Cancel</span><span class="btn primary">Submit</span></div></div>`;
}

const CSS = `
*{box-sizing:border-box}body{margin:0;background:#313338;color:#dbdee1;font:16px/1.375 "gg sans","Noto Sans","Helvetica Neue",Arial,sans-serif;-webkit-font-smoothing:antialiased}
.wrap{padding:16px 16px 20px 0}.frame-title{font-size:12px;color:#949ba4;padding:0 0 10px 16px;text-transform:uppercase;letter-spacing:.02em}
.msg{display:flex;gap:16px;padding:2px 16px 2px 16px}.avatar{width:40px;height:40px;border-radius:50%;flex:none;margin-top:2px}
.main{min-width:0;flex:1}.head{display:flex;align-items:center;gap:6px;margin-bottom:2px}.uname{color:#f2f3f5;font-weight:600}
.app{background:#5865f2;color:#fff;font-size:10px;font-weight:600;border-radius:3px;padding:1px 4px;line-height:15px}.when{color:#949ba4;font-size:12px;margin-left:2px}
.content{margin-bottom:4px}a{color:#00a8fc;text-decoration:none}strong{font-weight:700;color:#f2f3f5}code{background:#2b2d31;border-radius:3px;padding:0 3px;font-size:14px}
.h3{font-size:16px;font-weight:700;color:#f2f3f5;margin:4px 0 2px}.h2{font-size:20px;font-weight:700;color:#f2f3f5;margin:6px 0 2px}.sub{font-size:13px;color:#949ba4;line-height:1.3}.li{padding-left:14px;position:relative}.li:before{content:"•";position:absolute;left:2px}
.stamp{background:#3f4147;border-radius:3px;padding:0 2px}.emoji{width:20px;height:20px;vertical-align:-4px;object-fit:contain;margin:0 1px}
.embed{background:#2b2d31;border-left:4px solid;border-radius:4px;padding:8px 16px 16px 12px;margin-top:4px;max-width:520px}
.egrid{display:flex;gap:16px}.emain{min-width:0;flex:1}.author{display:flex;align-items:center;gap:8px;font-size:14px;font-weight:600;color:#f2f3f5;margin-top:8px}.aicon{width:24px;height:24px;border-radius:50%}
.etitle{font-weight:600;color:#f2f3f5;margin-top:8px}.etitle.link{color:#00a8fc}.edesc{font-size:14px;margin-top:8px;line-height:1.3}
.fields{display:flex;flex-wrap:wrap;gap:8px 16px;margin-top:8px}.field{flex:1 1 100%;font-size:14px;line-height:1.3}.field.inline{flex:1 1 30%;min-width:120px}.fname{font-weight:600;color:#f2f3f5;margin-bottom:2px}
.ethumb{width:80px;height:80px;object-fit:contain;border-radius:4px;flex:none;margin-top:8px}.eimage{max-width:100%;border-radius:4px;margin-top:16px}
.footer{display:flex;align-items:center;gap:8px;font-size:12px;color:#b5bac1;margin-top:8px}.ficon{width:20px;height:20px;border-radius:50%}
.row{display:flex;flex-wrap:wrap;gap:8px;margin-top:8px}
.btn{display:inline-flex;align-items:center;gap:6px;height:32px;padding:0 16px;border-radius:8px;font-size:14px;font-weight:500;color:#fff}
.btn.primary{background:#5865f2}.btn.secondary,.btn.link{background:#4e5058}.btn.success{background:#248046}.btn.danger{background:#da373c}.btn.disabled{opacity:.5}.ext{font-size:12px;opacity:.8}.link-plain{background:none}
.select{display:flex;justify-content:space-between;align-items:center;width:100%;max-width:400px;height:40px;padding:0 12px;background:#1e1f22;border:1px solid #1e1f22;border-radius:8px;font-size:14px;color:#949ba4}.select.chosen{color:#dbdee1}.chev{font-size:18px}
.container{background:#2b2d31;border:1px solid #3f4147;border-left:4px solid;border-radius:8px;padding:16px;margin-top:4px;max-width:540px}
.td{font-size:15px;line-height:1.35;margin:2px 0}.section{display:flex;gap:16px;align-items:flex-start;margin:2px 0}.sbody{flex:1;min-width:0}
.thumb{width:72px;height:72px;object-fit:contain;flex:none;border-radius:8px}.thumb.missing{background:#1e1f22}
.sep{border-top:1px solid #3f4147;margin:8px 0}.sep.big{margin:16px 0}.sep.nodiv{border:0}
.gallery{margin-top:8px}.gallery img{max-width:100%;border-radius:8px;display:block}
.eph{font-size:12px;color:#949ba4;margin-top:6px}
.above{opacity:.6;transform-origin:left;margin-bottom:8px}
.modal{background:#313338;border-radius:8px;margin:16px;max-width:440px;box-shadow:0 8px 24px rgba(0,0,0,.4);padding:16px}.mtitle{font-size:20px;font-weight:600;color:#f2f3f5;margin-bottom:16px}
.minput{margin-bottom:16px}.mlabel{font-size:12px;font-weight:700;text-transform:uppercase;color:#b5bac1;margin-bottom:8px}.req{color:#f23f42}.mfield{background:#1e1f22;border-radius:4px;height:40px;padding:10px;color:#6d6f78;font-size:14px}
.mfoot{display:flex;justify-content:flex-end;gap:8px;background:#2b2d31;margin:0 -16px -16px;padding:16px;border-radius:0 0 8px 8px}
`;

function frameHtml(f) {
  const webhook = f.id.startsWith("webhook-") || f.id.startsWith("post-reminder-webhook");
  let inner = "";
  if (f.modal) inner = modalHtml(f.modal);
  else {
    if (f.above) inner += `<div class="above">${messageHtml(f.above, f)}</div>`;
    inner += messageHtml(f.message, f, { webhook });
  }
  return `<!doctype html><meta charset="utf-8"><style>${CSS}</style><div class="wrap" id="frame"><div class="frame-title">${esc(f.id)}: ${esc(f.title ?? "")}</div>${inner}</div>`;
}

// ---- main -----------------------------------------------------------------------------------

mkdirSync(OUT, { recursive: true });
const only = process.argv.slice(3);
const files = readdirSync(PAYLOADS).filter((n) => n.endsWith(".json") && n !== "emoji.json");
const browser = await chromium.launch();
try {
  for (const client of CLIENTS) {
    const page = await browser.newPage({ viewport: { width: client.width, height: 800 }, deviceScaleFactor: 2 });
    for (const name of files) {
      const raw = JSON.parse(readFileSync(path.join(PAYLOADS, name), "utf8"));
      if (only.length && !only.includes(raw.id)) continue;
      const frame = { ...raw, clock: new Date(raw.clock ?? "2026-12-10T22:00:00Z") };
      await page.setContent(frameHtml(frame), { waitUntil: "load" });
      await page.locator("#frame").screenshot({ path: path.join(OUT, `${raw.id}-${client.name}.png`) });
    }
    await page.close();
  }
} finally {
  await browser.close();
}
console.log(`boards in ${path.relative(ROOT, OUT)}`);
