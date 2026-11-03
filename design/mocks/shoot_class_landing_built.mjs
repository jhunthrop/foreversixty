// design/mocks/shoot_class_landing_built.mjs
// Acceptance captures of the BUILT class landings, /guides/<class> (spec
// design/specs/2026-11-02-class-landing.md §13), taken from `astro preview` of a real-data build.
//
//   (cd web && FOREVER_DATA=real PUBLIC_API_BASE_URL=http://127.0.0.1:9 npm run build \
//      && npm run preview -- --port 4356)
//   LANDING_BASE=http://localhost:4356 LANDING_OUT=design/reviews/2026-11-02-class-landing \
//      node design/mocks/shoot_class_landing_built.mjs <prefix> [class ...]
//        -> <out>/<prefix>-<class>-<width>.png, signed out, full page, at 390/1024/1440/2000
//   ... node design/mocks/shoot_class_landing_built.mjs states
//        -> <out>/states-*.png: druid and hunter at 768, 1280 and 1920; above-the-fold crops at
//           1280x720 and 360x640; card hover, focus-visible and active at 1440; the "Your spec"
//           pill signed in as a Fury Warrior at 1440 and 390, with the card heights before and
//           after it printed.
import { createRequire } from "node:module";
import { mkdirSync } from "node:fs";
import { resolve } from "node:path";

const require = createRequire(
  new URL("../../web/package.json", import.meta.url),
);
const { chromium } = require("@playwright/test");

const BASE = process.env.LANDING_BASE ?? "http://localhost:4356";
const OUT = resolve(
  process.env.LANDING_OUT ?? "design/reviews/2026-11-02-class-landing",
);
const [PREFIX = "built", ...ONLY] = process.argv.slice(2);
const CLASSES = [
  "druid",
  "hunter",
  "mage",
  "paladin",
  "priest",
  "rogue",
  "shaman",
  "warlock",
  "warrior",
];
const WIDTHS = [390, 1024, 1440, 2000];
const STATE_WIDTHS = [768, 1280, 1920];
const HEIGHT = 900;
const FONT_SETTLE_MS = 1200;
const FURY_WARRIOR = {
  ok: true,
  data: {
    user: {
      id: 1,
      battletag: "Fixture#1",
      email: null,
      role: "user",
      anonymize: false,
    },
    characters: [
      {
        key: "us/normal/tester",
        region: "us",
        ruleset: "normal",
        name: "Tester",
        level: 60,
        class: "Warrior",
        spec: "Fury",
      },
    ],
    guilds: [],
  },
  error: null,
  request_id: "r",
};

async function open(browser, width, height, signedIn = false) {
  const context = await browser.newContext({
    viewport: { width, height },
    deviceScaleFactor: 1,
    isMobile: width < 768,
    hasTouch: width < 768,
  });
  const page = await context.newPage();
  await page.route("**/v1/**", (route) =>
    signedIn && route.request().url().endsWith("/v1/me")
      ? route.fulfill({
          status: 200,
          contentType: "application/json",
          body: JSON.stringify(FURY_WARRIOR),
        })
      : route.fulfill({ status: 401, body: "" }),
  );
  return { context, page };
}

async function visit(page, slug) {
  await page.goto(`${BASE}/guides/${slug}`, { waitUntil: "networkidle" });
  await page.waitForTimeout(FONT_SETTLE_MS);
  const overflow = await page.evaluate(
    () => document.documentElement.scrollWidth - window.innerWidth,
  );
  if (overflow > 0) console.warn(`${slug}: horizontal overflow ${overflow}px`);
}

async function shoot(page, name, fullPage = true) {
  const file = `${OUT}/${name}.png`;
  await page.screenshot({ path: file, fullPage });
  console.log(file);
}

async function widths(browser) {
  for (const width of WIDTHS) {
    const { context, page } = await open(browser, width, HEIGHT);
    for (const slug of ONLY.length > 0 ? ONLY : CLASSES) {
      await visit(page, slug);
      await shoot(page, `${PREFIX}-${slug}-${width}`);
    }
    await context.close();
  }
}

async function cardStates(browser) {
  const { context, page } = await open(browser, 1440, HEIGHT);
  for (const slug of ["hunter", "druid"]) {
    await visit(page, slug);
    const card = page
      .locator('[data-testid="guide-class-spec-cards"] > a')
      .nth(1);
    await card.hover();
    await page.waitForTimeout(300);
    await shoot(page, `states-${slug}-1440-hover`, false);
  }
  await visit(page, "hunter");
  // A key press first, so the focus that follows counts as keyboard focus (:focus-visible).
  await page.keyboard.press("Shift");
  await page
    .locator('[data-testid="guide-class-spec-cards"] > a')
    .first()
    .focus();
  await page.waitForTimeout(300);
  await shoot(page, "states-hunter-1440-focus", false);
  const box = await page
    .locator('[data-testid="guide-class-spec-cards"] > a')
    .nth(2)
    .boundingBox();
  await page.mouse.move(box.x + box.width / 2, box.y + box.height / 2);
  await page.mouse.down();
  await page.waitForTimeout(300);
  await shoot(page, "states-hunter-1440-active", false);
  await page.mouse.up();
  await context.close();
}

async function yourSpec(browser) {
  for (const width of [1440, 390]) {
    const out = await open(browser, width, HEIGHT);
    await visit(out.page, "warrior");
    const before = await out.page
      .getByTestId("guide-spec-card-fury")
      .boundingBox();
    await out.context.close();
    const signed = await open(browser, width, HEIGHT, true);
    await visit(signed.page, "warrior");
    const after = await signed.page
      .getByTestId("guide-spec-card-fury")
      .boundingBox();
    const pills = await signed.page.locator(".your-spec-pill").count();
    console.log(
      `your-spec @${width}: pills ${pills}, Fury card height ${before.height} -> ${after.height}`,
    );
    await shoot(signed.page, `states-warrior-your-spec-${width}`, false);
    await signed.context.close();
  }
}

async function states(browser) {
  for (const width of STATE_WIDTHS) {
    const { context, page } = await open(browser, width, HEIGHT);
    for (const slug of ["hunter", "druid"]) {
      await visit(page, slug);
      await shoot(page, `states-${slug}-${width}`);
    }
    await context.close();
  }
  for (const [width, height] of [
    [1280, 720],
    [360, 640],
  ]) {
    const { context, page } = await open(browser, width, height);
    await visit(page, "hunter");
    await shoot(page, `states-hunter-fold-${width}x${height}`, false);
    await context.close();
  }
  await cardStates(browser);
  await yourSpec(browser);
}

mkdirSync(OUT, { recursive: true });
const browser = await chromium.launch({ args: ["--font-render-hinting=none"] });
try {
  await (PREFIX === "states" ? states(browser) : widths(browser));
} finally {
  await browser.close();
}
