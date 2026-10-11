// design/mocks/shoot_character_header_built.mjs
// Acceptance captures of the BUILT shared character header (spec 2026-10-11-character-header.md):
//
//   states   every state of the real component at 1440 and 390, one file per state and one sheet per
//            width, from a throwaway harness page (ch_harness/) the script copies into web/src, serves
//            with `astro dev`, and removes again;
//   controls the action and the identity link at rest, hover, focus-visible and active;
//   pages    the home page and the BiS landing signed in at 1440 (and the home at 390, the BiS landing
//            signed out), from `astro preview` of a built dist with `/v1/me` stubbed.
//
//   (cd web && FOREVER_DATA=fixture npm run build)
//   node design/mocks/shoot_character_header_built.mjs [states|controls|pages ...]
//     -> design/mocks/renders/character-header-built-*.png   (git-ignored)
//   python3 design/mocks/compose_character_header_built.py   -> the built-versus-mock composite
//
// The character is the mock's invented example (Obnoxious Yell, a Gnome Fury Warrior); nothing here is
// a real player.
import { spawn } from "node:child_process";
import { copyFileSync, mkdirSync, readFileSync, rmSync } from "node:fs";
import { createRequire } from "node:module";

const require = createRequire(new URL("../../web/package.json", import.meta.url));
const { chromium } = require("@playwright/test");

const WEB = new URL("../../web/", import.meta.url).pathname;
const REPO = new URL("../../", import.meta.url).pathname;
const OUT = new URL("./renders/", import.meta.url).pathname;
const HARNESS = new URL("./ch_harness/", import.meta.url).pathname;
const DEV_PORT = process.env.CH_DEV_PORT ?? "4459";
const PREVIEW_PORT = process.env.CH_PREVIEW_PORT ?? "4458";
const SETTLE_MS = 3500;
const MIN = 60_000;
const BUILD = JSON.parse(readFileSync(`${WEB}src/data/active-build.json`, "utf8")).build;

const slug = (text) =>
  text
    .toLowerCase()
    .replace(/[^a-z0-9]+/g, "-")
    .replace(/^-|-$/g, "")
    .slice(0, 48);

async function waitFor(url, attempts = 90) {
  for (let i = 0; i < attempts; i += 1) {
    try {
      if ((await fetch(url)).ok) return;
    } catch {
      /* not up yet */
    }
    await new Promise((resolve) => setTimeout(resolve, 1000));
  }
  throw new Error(`${url} never came up`);
}

async function withHarness(run) {
  const component = `${WEB}src/components/character/__HeaderStatesHarness.svelte`;
  const page = `${WEB}src/pages/ch-states.astro`;
  copyFileSync(`${HARNESS}HeaderStatesHarness.svelte`, component);
  copyFileSync(`${HARNESS}ch-states.astro`, page);
  const server = spawn("npx", ["astro", "dev", "--port", DEV_PORT], {
    cwd: WEB,
    env: { ...process.env, FOREVER_DATA: "fixture", ASTRO_DEV_BACKGROUND: "1" },
    stdio: "ignore",
  });
  try {
    await waitFor(`http://localhost:${DEV_PORT}/ch-states`);
    await run(`http://localhost:${DEV_PORT}/ch-states`);
  } finally {
    server.kill("SIGTERM");
    rmSync(component, { force: true });
    rmSync(page, { force: true });
  }
}

async function states(browser) {
  await withHarness(async (url) => {
    for (const width of [1440, 390]) {
      const context = await browser.newContext({ viewport: { width, height: 900 } });
      await context.addCookies([{ name: "fs_csrf", value: "token", domain: "localhost", path: "/" }]);
      const page = await context.newPage();
      await page.goto(url);
      await page.waitForTimeout(SETTLE_MS);
      const sheet = `${OUT}character-header-built-states-${width}.png`;
      await page.locator(".sheet").screenshot({ path: sheet });
      console.log(sheet);
      // One file per state in the first section (the real 1440 and the real 390 render of each).
      const cells = await page.locator(".sheet > .cell").all();
      for (const cell of cells) {
        const caption = (await cell.locator(".cap").first().textContent()) ?? "state";
        const out = `${OUT}character-header-built-${slug(caption)}-${width}.png`;
        await cell.screenshot({ path: out });
        console.log(out);
      }
      await context.close();
    }
  });
}

async function controls(browser) {
  await withHarness(async (url) => {
    const context = await browser.newContext({ viewport: { width: 1440, height: 900 } });
    await context.addCookies([{ name: "fs_csrf", value: "token", domain: "localhost", path: "/" }]);
    const page = await context.newPage();
    await page.goto(url);
    await page.waitForTimeout(SETTLE_MS);
    const header = page.getByTestId("character-header").first();
    const shot = async (name, prepare) => {
      await prepare();
      await page.waitForTimeout(250);
      const out = `${OUT}character-header-built-control-${name}.png`;
      await header.screenshot({ path: out });
      console.log(out);
    };
    const action = header.getByTestId("character-header-action");
    const link = header.getByTestId("character-header-link");
    await shot("action-rest", async () => page.mouse.move(0, 0));
    await shot("action-hover", () => action.hover());
    await shot("action-active", async () => {
      await action.hover();
      await page.mouse.down();
    });
    await page.mouse.up();
    await page.mouse.move(0, 0);
    await shot("action-focus", async () => {
      await page.keyboard.press("Tab");
      await action.focus();
      await page.keyboard.press("Shift+Tab");
      await page.keyboard.press("Tab");
    });
    await shot("identity-hover", () => link.hover());
    await shot("identity-focus", async () => {
      await page.mouse.move(0, 0);
      await link.focus();
      await page.keyboard.press("Shift+Tab");
      await page.keyboard.press("Tab");
    });
    const sec = page.locator(".ch-secondary").first();
    await sec.scrollIntoViewIfNeeded();
    for (const [name, prepare] of [
      ["secondary-rest", async () => page.mouse.move(0, 0)],
      ["secondary-hover", () => sec.getByTestId("character-header-action").hover()],
    ]) {
      await prepare();
      await page.waitForTimeout(250);
      const out = `${OUT}character-header-built-control-${name}.png`;
      await sec.screenshot({ path: out });
      console.log(out);
    }
    await context.close();
  });
}

const fury = JSON.parse(readFileSync(`${REPO}data/builds/${BUILD}/bis/warrior-fury.json`, "utf8"));
const band = fury.bands.find((b) => b.faction === "alliance" && b.band === 40 && (b.preset ?? "bare") !== "raid");
const GEAR = Object.fromEntries(
  band.slots
    .filter((s) => typeof s.item_id === "number")
    .slice(0, 6)
    .map((s) => [s.slot, s.item_id]),
);

const me = (character) => ({
  ok: true,
  data: {
    user: { id: 1, battletag: "Fixture#1", email: null, role: "user", anonymize: false },
    characters: [character],
    guilds: [],
  },
  error: null,
  request_id: "r",
});

const OBNOXIOUS = {
  key: "us/normal/obnoxious-yell",
  region: "us",
  ruleset: "normal",
  name: "Obnoxious Yell",
  class: "Warrior",
  spec: "Fury",
  race: "Gnome",
  level: 47,
  realm: "Living Flame",
  faction: "alliance",
  build: { source: "addon", captured_at: new Date(Date.now() - 12 * MIN).toISOString(), gear: GEAR },
};

/** name -> { path, width, signedIn, clip? } */
const PAGES = {
  "home-1440": { path: "/", width: 1440, signedIn: true },
  "home-390": { path: "/", width: 390, signedIn: true },
  "bis-landing-1440": { path: "/bis", width: 1440, signedIn: true },
  "bis-landing-signed-out-1440": { path: "/bis", width: 1440, signedIn: false },
  "bis-landing-390": { path: "/bis", width: 390, signedIn: true },
};

async function stubBis(context) {
  for (const spec of ["warrior-fury"]) {
    await context.route(`**/data/*/bis/${spec}.json`, (route) =>
      route.fulfill({ status: 200, contentType: "application/json", body: readFileSync(`${REPO}data/builds/${BUILD}/bis/${spec}.json`, "utf8") }),
    );
  }
  await context.route("**/v1/sims**", (route) =>
    route.fulfill({ status: 200, contentType: "application/json", body: JSON.stringify({ ok: true, data: { rows: [], total: 0, page: 1, per_page: 20 }, error: null, request_id: "r" }) }),
  );
}

async function pages(browser) {
  const preview = spawn("npx", ["astro", "preview", "--port", PREVIEW_PORT], {
    cwd: WEB,
    env: { ...process.env, FOREVER_DATA: "fixture", ASTRO_PREVIEW_BACKGROUND: "1" },
    stdio: "ignore",
  });
  try {
    await waitFor(`http://localhost:${PREVIEW_PORT}/bis`);
    for (const [name, shot] of Object.entries(PAGES)) {
      const context = await browser.newContext({ viewport: { width: shot.width, height: 900 } });
      if (shot.signedIn) {
        await context.addCookies([{ name: "fs_csrf", value: "token", domain: "localhost", path: "/" }]);
        await context.route("**/v1/me", (route) =>
          route.fulfill({ status: 200, contentType: "application/json", body: JSON.stringify(me(OBNOXIOUS)) }),
        );
      }
      await stubBis(context);
      const page = await context.newPage();
      await page.goto(`http://localhost:${PREVIEW_PORT}${shot.path}`);
      await page.waitForTimeout(SETTLE_MS);
      const out = `${OUT}character-header-built-page-${name}.png`;
      await page.screenshot({ path: out, fullPage: false });
      console.log(out);
      await context.close();
    }
  } finally {
    preview.kill("SIGTERM");
  }
}

mkdirSync(OUT, { recursive: true });
const wanted = process.argv.slice(2).length ? process.argv.slice(2) : ["states", "controls", "pages"];
const browser = await chromium.launch();
if (wanted.includes("states")) await states(browser);
if (wanted.includes("controls")) await controls(browser);
if (wanted.includes("pages")) await pages(browser);
await browser.close();
