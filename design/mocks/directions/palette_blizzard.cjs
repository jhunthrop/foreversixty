// Reads the exact colours of Blizzard's WoW: Forever page for the round-4 tokens (design/specs/2026-11-09-directions.md).
//   node design/mocks/directions/palette_blizzard.cjs <node_modules with playwright-core>
// -> design/mocks/directions/blizzard/palette.json: every computed colour on the page with its use count (text, background,
//    border) across all elements and shadow roots, and every CSS custom property whose value is a colour.
const path = require("node:path");
const fs = require("node:fs");
const { chromium } = require(path.join(process.argv[2], "playwright-core"));

const URL = "https://worldofwarcraft.blizzard.com/en-us/forever";
const OUT = path.join(__dirname, "blizzard", "palette.json");

function collect() {
  const count = { text: {}, background: {}, border: {} };
  const add = (k, v) => {
    if (!v || v === "rgba(0, 0, 0, 0)" || v === "transparent") return;
    count[k][v] = (count[k][v] || 0) + 1;
  };
  const walk = (root) => {
    for (const el of root.querySelectorAll("*")) {
      const s = getComputedStyle(el);
      if ((el.innerText || "").trim()) add("text", s.color);
      add("background", s.backgroundColor);
      if (parseFloat(s.borderTopWidth) > 0) add("border", s.borderTopColor);
      if (el.shadowRoot) walk(el.shadowRoot);
    }
  };
  walk(document);
  const vars = {};
  for (const sheet of document.styleSheets) {
    let rules;
    try { rules = sheet.cssRules; } catch { continue; }
    for (const r of rules) {
      if (!r.style) continue;
      for (const p of r.style) {
        if (!p.startsWith("--")) continue;
        const v = r.style.getPropertyValue(p).trim();
        if (/^(#[0-9a-f]{3,8}|rgba?\(|hsla?\()/i.test(v)) vars[p] = v;
      }
    }
  }
  const sort = (o) => Object.entries(o).sort((a, b) => b[1] - a[1]);
  return { text: sort(count.text), background: sort(count.background), border: sort(count.border), vars };
}

(async () => {
  const browser = await chromium.launch();
  const pg = await browser.newPage({ viewport: { width: 1440, height: 900 } });
  await pg.goto(URL, { waitUntil: "load", timeout: 90000 });
  await pg.waitForTimeout(5000);
  for (let y = 0; y < 11000; y += 900) { await pg.evaluate((v) => window.scrollTo(0, v), y); await pg.waitForTimeout(250); }
  fs.writeFileSync(OUT, JSON.stringify(await pg.evaluate(collect), null, 1));
  await browser.close();
  console.log("ok", OUT);
})();
