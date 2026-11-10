// Review captures for the icon lane (design/reviews/2026-11-09-nav-icons): the site header at each
// desktop width, the phone Menu open, each door's panel open, the compact lookup open, and the pages
// whose icons changed most, signed in with a long name so the bar is at its fullest. Prints the bar's
// fit per width: rows (1 = no wrap), overflow past the bar, and the slack between the last door and
// the account side.
//   node design/mocks/shoot_nav_icons.mjs <port> <out dir> <before|after>
import { createRequire } from "node:module";
import { mkdirSync } from "node:fs";
import path from "node:path";
const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");

const [port = "4377", outDir = "design/reviews/2026-11-09-nav-icons", prefix = "after"] = process.argv.slice(2);
const BASE = `http://localhost:${port}`;
mkdirSync(outDir, { recursive: true });

const ME = {
  ok: true,
  error: null,
  request_id: "r",
  data: {
    user: { id: 1, battletag: "Fixture#1", email: null, role: "user", anonymize: false },
    characters: [
      {
        key: "us/normal/frickin-laserbeams",
        region: "us",
        ruleset: "normal",
        name: "Frickin Laserbeams",
        class: "Hunter",
        spec: "Marksmanship",
        level: 60,
        realm: "Living Flame",
        faction: "alliance",
      },
    ],
    guilds: [],
  },
};

async function open(browser, width, height, url) {
  const context = await browser.newContext({ viewport: { width, height }, deviceScaleFactor: 1 });
  const page = await context.newPage();
  await context.addCookies([{ name: "fs_csrf", value: "token", domain: "localhost", path: "/" }]);
  await page.route("**/v1/me", (route) =>
    route.fulfill({ status: 200, contentType: "application/json", body: JSON.stringify(ME) }),
  );
  await page.goto(`${BASE}${url}`, { waitUntil: "networkidle" });
  await page.locator("[data-site-header][data-header-wired]").waitFor({ state: "attached" });
  await page.waitForTimeout(300);
  return { context, page };
}

/** The bar's fit: one row of doors, nothing past the bar's edge, the slack before the account side. */
async function fit(page) {
  return page.evaluate(() => {
    const bar = document.querySelector("[data-site-header]");
    const box = bar.getBoundingClientRect();
    const shown = (el) => el.getBoundingClientRect().width > 0;
    const doors = [...bar.querySelectorAll(".site-door:not(.site-door-premium)")].filter(shown);
    const parts = [...bar.querySelectorAll(".site-wordmark, .site-door, .site-right")].filter(shown);
    const rows = new Set(parts.map((el) => Math.round(el.getBoundingClientRect().top / 8))).size;
    const overflow = Math.max(0, ...parts.map((el) => el.getBoundingClientRect().right - box.right));
    const slack = bar.querySelector(".site-right").getBoundingClientRect().left - doors.at(-1).getBoundingClientRect().right;
    const icon = bar.querySelector(".site-door [data-icon]");
    const iconShown = icon !== null && icon.getBoundingClientRect().width > 0;
    return { bar: Math.round(box.width), rows, overflow: Math.round(overflow), slack: Math.round(slack), iconShown };
  });
}

const browser = await chromium.launch();
const shot = (name) => path.join(outDir, `${prefix}-${name}.png`);
try {
  for (const width of [1024, 1280, 1440, 1760, 2000]) {
    const { context, page } = await open(browser, width, 900, "/bis");
    console.log(width, JSON.stringify(await fit(page)));
    await page.screenshot({ path: shot(`header-${width}`), clip: { x: 0, y: 0, width, height: 104 } });
    if (width === 2000) {
      // The door icons' states: hover, keyboard focus, and the Guild panel open.
      const clip = { x: 0, y: 0, width, height: 104 };
      await page.locator('[data-door="sim"]').hover();
      await page.waitForTimeout(150);
      await page.screenshot({ path: shot("header-2000-hover-sim"), clip });
      await page.mouse.move(5, 600);
      await page.locator('[data-door="planner"]').focus();
      await page.keyboard.press("Shift+Tab");
      await page.keyboard.press("Tab");
      await page.waitForTimeout(150);
      await page.screenshot({ path: shot("header-2000-focus-planner"), clip });
      await page.locator('[data-door-toggle="guild"]').click();
      await page.waitForTimeout(200);
      await page.screenshot({ path: shot("panel-guild-2000"), clip: { x: 0, y: 0, width, height: 460 } });
      await page.keyboard.press("Escape");
    }
    if (width === 1440 || width === 1024) {
      for (const door of ["guides", "logs", "guild"]) {
        await page.locator(`[data-door-toggle="${door}"]`).click();
        await page.waitForTimeout(200);
        const height = door === "guides" ? 600 : 460;
        await page.screenshot({ path: shot(`panel-${door}-${width}`), clip: { x: 0, y: 0, width, height } });
        await page.keyboard.press("Escape");
        await page.waitForTimeout(150);
      }
      await page.locator('[data-testid="site-search-toggle"]').click();
      await page.waitForTimeout(200);
      await page.screenshot({ path: shot(`search-${width}`), clip: { x: width - 640, y: 0, width: 640, height: 200 } });
    }
    await context.close();
  }
  {
    const { context, page } = await open(browser, 390, 844, "/bis");
    await page.locator("[data-menu-button]").click();
    await page.waitForTimeout(250);
    await page.screenshot({ path: shot("menu-390"), fullPage: true });
    await context.close();
  }
  for (const [name, url] of [
    ["page-design-sheet-1440", "/design-sheet"],
    ["page-bis-1440", "/bis/hunter/marksmanship"],
  ]) {
    const { context, page } = await open(browser, 1440, 1100, url);
    await page.screenshot({ path: shot(name), fullPage: name.includes("design-sheet") });
    await context.close();
  }
  {
    // The account menu open: its tiles, chevrons and close are all from the one set.
    const { context, page } = await open(browser, 1440, 900, "/bis");
    await page.locator(".csel-btn").first().click();
    await page.waitForTimeout(300);
    await page.screenshot({ path: shot("account-menu-1440"), clip: { x: 640, y: 0, width: 800, height: 700 } });
    await context.close();
  }
  if (prefix === "after") {
    // The contact sheet: every glyph in the set at 16, 20 and 24 px on the dark ground.
    const context = await browser.newContext({ viewport: { width: 1100, height: 900 }, deviceScaleFactor: 2 });
    const page = await context.newPage();
    const sheet = new URL("../assets/icons/ui/preview.html", import.meta.url).href;
    await page.goto(sheet);
    await page.screenshot({ path: path.join(outDir, "contact-sheet.png"), fullPage: true });
    await context.close();
  }
} finally {
  await browser.close();
}
