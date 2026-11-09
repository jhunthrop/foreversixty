// Studies Blizzard's WoW: Forever page for the round-3 directions (design/specs/2026-11-09-directions.md section 7).
//   node design/mocks/directions/shoot_blizzard.cjs <node_modules with playwright-core>
// -> design/mocks/directions/blizzard/full-<w>.png, screen-<w>-<n>.png while scrolling, hover-*.png, and study.json
//    (computed type, colours, backgrounds, borders and transitions of the page's visible elements). Reference only:
//    nothing captured here is used in a mock.
const path = require("node:path");
const fs = require("node:fs");
const { chromium } = require(path.join(process.argv[2], "playwright-core"));

const URL = "https://worldofwarcraft.blizzard.com/en-us/forever";
const OUT = path.join(__dirname, "blizzard");
const SCREENS = 8;

function study() {
  const seen = new Map();
  const pick = (el) => {
    const s = getComputedStyle(el);
    const r = el.getBoundingClientRect();
    return { tag: el.tagName.toLowerCase(), cls: String(el.className).slice(0, 60), text: (el.innerText || "").trim().slice(0, 60),
      font: s.fontFamily, size: s.fontSize, weight: s.fontWeight, lh: s.lineHeight, ls: s.letterSpacing, tt: s.textTransform,
      color: s.color, bg: s.backgroundColor, bgimg: s.backgroundImage.slice(0, 160), border: s.border, radius: s.borderRadius,
      shadow: s.boxShadow, tshadow: s.textShadow, transition: s.transition.slice(0, 120), w: Math.round(r.width), h: Math.round(r.height) };
  };
  const sel = "h1,h2,h3,h4,p,a,button,li,span,img,section,header,footer,nav,[class*=button],[class*=Button],[class*=frame],[class*=Frame],[class*=divider]";
  for (const el of document.querySelectorAll(sel)) {
    const r = el.getBoundingClientRect();
    if (r.width < 8 || r.height < 8) continue;
    const p = pick(el);
    const key = [p.tag, p.font, p.size, p.weight, p.color, p.tt, p.bgimg.slice(0, 40)].join("|");
    if (!seen.has(key)) seen.set(key, p);
  }
  const fonts = [...document.fonts].map((f) => `${f.family} ${f.weight} ${f.style} ${f.status}`);
  return { fonts: [...new Set(fonts)], elements: [...seen.values()].slice(0, 400), height: document.documentElement.scrollHeight };
}

async function capture(browser, width) {
  const pg = await browser.newPage({ viewport: { width, height: width < 500 ? 844 : 900 } });
  // The page keeps streaming (video, analytics), so networkidle never comes: wait for load, then let it settle.
  await pg.goto(URL, { waitUntil: "load", timeout: 90000 });
  await pg.waitForTimeout(5000);
  const total = await pg.evaluate(() => document.documentElement.scrollHeight);
  const step = Math.max(400, Math.floor(total / SCREENS));
  for (let i = 0; i < SCREENS; i++) {
    await pg.evaluate((y) => window.scrollTo(0, y), i * step);
    await pg.waitForTimeout(900); // scroll-triggered motion settles
    await pg.screenshot({ path: path.join(OUT, `screen-${width}-${i}.png`) });
  }
  await pg.evaluate(() => window.scrollTo(0, 0));
  await pg.waitForTimeout(800);
  await pg.screenshot({ path: path.join(OUT, `full-${width}.png`), fullPage: true });
  if (width > 500) {
    const targets = await pg.$$("a[class*=button], a[class*=Button], button, [class*=Button]");
    for (const [i, t] of targets.slice(0, 4).entries()) {
      try {
        await t.scrollIntoViewIfNeeded();
        await t.hover();
        await pg.waitForTimeout(500);
        const box = await t.boundingBox();
        if (box) await pg.screenshot({ path: path.join(OUT, `hover-${i}.png`), clip: { x: Math.max(0, box.x - 40), y: Math.max(0, box.y - 40), width: box.width + 80, height: box.height + 80 } });
      } catch (e) {
        console.error("hover", i, e.message);
      }
    }
    fs.writeFileSync(path.join(OUT, "study.json"), JSON.stringify(await pg.evaluate(study), null, 1));
  }
  await pg.close();
  console.log("ok", width, total);
}

(async () => {
  fs.mkdirSync(OUT, { recursive: true });
  const browser = await chromium.launch();
  for (const w of [1440, 390]) await capture(browser, w);
  await browser.close();
})();
