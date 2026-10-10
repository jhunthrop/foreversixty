// design/mocks/shoot_nav_built.mjs
// Acceptance captures of the BUILT nav bar and character selector (spec 2026-10-09-nav-
// character-selector.md section 11), taken from `astro preview` of a FOREVER_DATA=fixture
// build with `/v1/me` stubbed, at the viewports and states the spec names.
//
//   (cd web && FOREVER_DATA=fixture npm run build && npm run preview -- --port 4325)
//   node design/mocks/shoot_nav_built.mjs [name ...]      -> design/mocks/renders/nav-built-*.png
//   python3 design/mocks/compose_nav_built.py             -> the side-by-sides and the states sheet
//
// The characters are the same invented examples the mock uses (gen_nav.py); "Obnoxious Yell" is
// the owner's own. Playwright is found through web/.
import { createRequire } from "node:module";
import { mkdirSync } from "node:fs";

const require = createRequire(
  new URL("../../web/package.json", import.meta.url),
);
const { chromium } = require("@playwright/test");

const BASE = process.env.NAV_BASE ?? "http://localhost:4325";
const OUT = new URL("./renders/", import.meta.url).pathname;
const PAGE_PATH = "/planner";
const DAY = 86_400_000;
const MIN = 60_000;

const slugOf = (name) => name.toLowerCase().replace(/\s+/g, "-");
const keyOf = (c) => `us/normal/${slugOf(c.name)}`;

const ch = (name, cls, spec, level, faction, ageMs, extra = {}) => ({
  name,
  cls,
  spec,
  level,
  faction,
  ageMs,
  realm: "Living Flame",
  source: "addon",
  ...extra,
});
const OBNOXIOUS = ch(
  "Obnoxious Yell",
  "Warrior",
  "Fury",
  60,
  "alliance",
  12 * MIN,
);
const FROSTBYTE = ch("Frostbyte", "Mage", "Frost", 42, "horde", 3 * DAY, {
  realm: "Stonehearth",
  ruleset: "hardcore",
  source: "blizzard",
});
const SHADOWMEND = ch("Shadowmend", "Priest", "Shadow", 35, "horde", 2 * DAY);
const OAKHEART = ch(
  "Oakheart",
  "Paladin",
  "Protection",
  60,
  "alliance",
  3 * DAY,
  {
    source: "blizzard",
    syncError: "refresh failed",
  },
);
const QUICKSHOT = ch(
  "Quickshot",
  "Hunter",
  "Beast Mastery",
  60,
  "horde",
  5 * DAY,
);
const NIGHTSNARE = ch("Nightsnare", "Rogue", "Combat", 60, "horde", 6 * DAY, {
  realm: "Stonehearth",
});
const CINDERMAW = ch(
  "Cindermaw",
  "Warlock",
  "Affliction",
  51,
  "horde",
  8 * DAY,
);
const DAWNBRINGER = ch(
  "Dawnbringer",
  "Paladin",
  "Retribution",
  60,
  "alliance",
  12 * DAY,
  {
    realm: "Stonehearth",
  },
);
const TREEWALKER = ch(
  "Treewalker",
  "Druid",
  "Balance",
  60,
  "alliance",
  19 * DAY,
);
const SIX = [OBNOXIOUS, FROSTBYTE, SHADOWMEND, OAKHEART, QUICKSHOT, TREEWALKER];
const EIGHT = [
  OBNOXIOUS,
  FROSTBYTE,
  SHADOWMEND,
  OAKHEART,
  QUICKSHOT,
  NIGHTSNARE,
  CINDERMAW,
  TREEWALKER,
];
const NINE = [...EIGHT.slice(0, 7), DAWNBRINGER, TREEWALKER];
const CHOSEN = [FROSTBYTE, SHADOWMEND].map(keyOf);

function meBody(characters) {
  return {
    ok: true,
    error: null,
    request_id: "r",
    data: {
      user: {
        id: 1,
        battletag: "Fixture#1",
        email: null,
        role: "user",
        anonymize: false,
      },
      guilds: [],
      characters: characters.map((c) => ({
        key: keyOf(c),
        region: "us",
        ruleset: c.ruleset ?? "normal",
        name: c.name,
        class: c.cls,
        spec: c.spec,
        level: c.level,
        realm: c.realm,
        faction: c.faction,
        build: {
          source: c.source,
          captured_at: new Date(Date.now() - c.ageMs).toISOString(),
          ...(c.syncError ? { sync_error: c.syncError } : {}),
        },
      })),
    },
  };
}

const pointerFor = (c) => ({
  source: "armory",
  ref: keyOf(c),
  label: `${c.name} · ${c.spec} ${c.cls}`,
  classSlug: c.cls.toLowerCase(),
});

let browser;

/** opts: { width, height, chars, current, out: 'signed-out' | 'expired' | 'failed' | 'slow', pasted, blockJs, order } */
async function open(opts) {
  const context = await browser.newContext({
    viewport: { width: opts.width, height: opts.height ?? 900 },
    deviceScaleFactor: 1,
  });
  const page = await context.newPage();
  const signedOut = opts.out === "signed-out" || opts.out === "expired";
  if (!signedOut)
    await context.addCookies([{ name: "fs_csrf", value: "token", url: BASE }]);
  await page.route("**/v1/me", async (route) => {
    if (signedOut) {
      return route.fulfill({
        status: 401,
        contentType: "application/json",
        body: '{"ok":false,"data":null,"error":null,"request_id":"r"}',
      });
    }
    if (opts.out === "failed") return route.abort();
    if (opts.out === "slow")
      await new Promise((resolve) => setTimeout(resolve, 60_000));
    return route.fulfill({
      status: 200,
      contentType: "application/json",
      body: JSON.stringify(meBody(opts.chars ?? [])),
    });
  });
  const pointer =
    opts.pasted ?? (opts.current ? pointerFor(opts.current) : null);
  await page.addInitScript(
    ({ pointer, order }) => {
      if (
        pointer &&
        window.localStorage.getItem("fs.currentCharacter") === null
      ) {
        const savedAt = new Date(Date.now() - 2 * 86_400_000).toISOString();
        window.localStorage.setItem(
          "fs.currentCharacter",
          JSON.stringify({ ...pointer, savedAt }),
        );
      }
      if (order && window.localStorage.getItem("fs.characterOrder") === null) {
        window.localStorage.setItem("fs.characterOrder", JSON.stringify(order));
      }
    },
    { pointer, order: opts.order ?? null },
  );
  if (opts.blockJs)
    await page.route("**/_astro/**/*.js", (route) => route.abort());
  await page.goto(`${BASE}${PAGE_PATH}`);
  if (!opts.blockJs) await page.getByTestId("character-selector").waitFor();
  await page.waitForTimeout(400);
  return { page, context };
}

async function snap(page, name, clip) {
  const path = `${OUT}nav-built-${name}.png`;
  await page.screenshot({ path, clip: clip ?? undefined });
  console.log(path);
}

const selector = (page) => page.getByTestId("character-selector");
const bar = (width, height = 140) => ({ x: 0, y: 0, width, height });

async function barShot(width, name, opts = {}) {
  const { page, context } = await open({
    width,
    chars: SIX,
    current: OBNOXIOUS,
    ...opts,
  });
  const box = await page.getByRole("banner").boundingBox();
  const sideScroll = await page.evaluate(
    () =>
      document.documentElement.scrollWidth -
      document.documentElement.clientWidth,
  );
  const nameEl = page
    .locator(width >= 1024 ? "[data-testid=selector-name]" : ".csel-name-bar")
    .first();
  const clipped = await nameEl.evaluate(
    (el) => el.scrollWidth > el.clientWidth,
  );
  const doorTops = await page
    .getByTestId("primary-nav")
    .getByRole("link")
    .evaluateAll(
      (links) =>
        [
          ...new Set(
            links.map((a) => Math.round(a.getBoundingClientRect().top)),
          ),
        ].length,
    );
  console.log(
    `  ${name}: bar height ${box.height}, sideways overflow ${sideScroll}, name clipped ${clipped}, door rows ${doorTops}`,
  );
  await snap(page, name, bar(width, Math.ceil(box.height) + 60));
  await context.close();
}

const SHOTS = {};

for (const width of [2000, 1440, 1439, 1280, 1279, 1024, 1023, 768]) {
  SHOTS[`bar-${width}`] = () => barShot(width, `bar-${width}`);
}

for (const [label, longName] of [
  ["22", "Bartholomewthe Unbearab"],
  ["16", "Sir Reginald Bunt"],
]) {
  for (const width of [1440, 1024]) {
    SHOTS[`longname-${label}-${width}`] = async () => {
      const long = { ...OBNOXIOUS, name: longName };
      await barShot(width, `longname-${label}-${width}`, {
        chars: [long],
        current: long,
      });
    };
  }
}

for (const width of [390, 360]) {
  SHOTS[`phone-${width}`] = () => barShot(width, `phone-${width}`);
  SHOTS[`phone-${width}-menu`] = async () => {
    const { page, context } = await open({
      width,
      height: 844,
      chars: SIX,
      current: OBNOXIOUS,
    });
    await page.getByTestId("menu-button").click();
    await snap(page, `phone-${width}-menu`, { x: 0, y: 0, width, height: 420 });
    await context.close();
  };
  SHOTS[`phone-${width}-sheet`] = async () => {
    const { page, context } = await open({
      width,
      height: 844,
      chars: SIX,
      current: OBNOXIOUS,
      order: CHOSEN,
    });
    await selector(page).click();
    await page.waitForTimeout(300);
    await snap(page, `phone-${width}-sheet`);
    await context.close();
  };
}

async function openList(name, chars, extra = {}) {
  const { page, context } = await open({
    width: 1440,
    height: 900,
    chars,
    current: OBNOXIOUS,
    order: CHOSEN,
    ...extra,
  });
  await selector(page).click();
  await page.waitForTimeout(300);
  await snap(page, name, { x: 0, y: 0, width: 1440, height: 700 });
  await context.close();
}
SHOTS["open-six"] = () => openList("open-six", SIX);
SHOTS["open-eight"] = () => openList("open-eight", EIGHT);
SHOTS["open-nine"] = () => openList("open-nine", NINE);

// Full-viewport captures for the same-scale side-by-sides with the mock (compose_nav_built.py).
SHOTS["full-bar-1440"] = async () => {
  const { page, context } = await open({
    width: 1440,
    height: 900,
    chars: SIX,
    current: OBNOXIOUS,
  });
  await snap(page, "full-bar-1440");
  await context.close();
};
SHOTS["full-open-1440"] = async () => {
  const { page, context } = await open({
    width: 1440,
    height: 900,
    chars: SIX,
    current: OBNOXIOUS,
    order: CHOSEN,
  });
  await selector(page).click();
  await page.waitForTimeout(300);
  await snap(page, "full-open-1440");
  await context.close();
};
SHOTS["full-bar-390"] = async () => {
  const { page, context } = await open({
    width: 390,
    height: 844,
    chars: SIX,
    current: OBNOXIOUS,
  });
  await snap(page, "full-bar-390");
  await context.close();
};
SHOTS["full-menu-390"] = async () => {
  const { page, context } = await open({
    width: 390,
    height: 844,
    chars: SIX,
    current: OBNOXIOUS,
  });
  await page.getByTestId("menu-button").click();
  await snap(page, "full-menu-390");
  await context.close();
};

async function signedOutShot(name, width, height, act) {
  const { page, context } = await open({ width, height, out: "signed-out" });
  await act?.(page);
  await page.waitForTimeout(300);
  await snap(page, name, { x: 0, y: 0, width, height: Math.min(height, 700) });
  await context.close();
}
SHOTS["signed-out-1440"] = () => signedOutShot("signed-out-1440", 1440, 900);
SHOTS["signed-out-1440-open"] = () =>
  signedOutShot("signed-out-1440-open", 1440, 900, (p) => selector(p).click());
SHOTS["signed-out-1024"] = () => signedOutShot("signed-out-1024", 1024, 900);
SHOTS["signed-out-1024-open"] = () =>
  signedOutShot("signed-out-1024-open", 1024, 900, (p) => selector(p).click());
SHOTS["signed-out-390"] = () => signedOutShot("signed-out-390", 390, 844);
SHOTS["signed-out-390-menu"] = () =>
  signedOutShot("signed-out-390-menu", 390, 844, (p) =>
    p.getByTestId("menu-button").click(),
  );
SHOTS["signed-out-390-sheet"] = () =>
  signedOutShot("signed-out-390-sheet", 390, 844, (p) => selector(p).click());
SHOTS["signed-out-pasted"] = async () => {
  const { page, context } = await open({
    width: 1440,
    height: 900,
    out: "signed-out",
    pasted: {
      source: "code",
      ref: "FS1:x",
      label: "Shadowmend · Shadow Priest",
      classSlug: "priest",
    },
  });
  await selector(page).click();
  await page.waitForTimeout(300);
  await snap(page, "signed-out-pasted", {
    x: 0,
    y: 0,
    width: 1440,
    height: 700,
  });
  await context.close();
};
SHOTS["expired"] = async () => {
  const { page, context } = await open({
    width: 1440,
    height: 900,
    out: "expired",
    current: OBNOXIOUS,
  });
  await selector(page).click();
  await page.waitForTimeout(300);
  await snap(page, "expired", { x: 0, y: 0, width: 1440, height: 700 });
  await context.close();
};

async function stateShot(
  name,
  opts,
  act,
  clip = { x: 880, y: 0, width: 560, height: 640 },
) {
  const { page, context } = await open({
    width: 1440,
    height: 900,
    chars: SIX,
    current: OBNOXIOUS,
    order: CHOSEN,
    ...opts,
  });
  await act?.(page);
  await page.waitForTimeout(300);
  await snap(page, `state-${name}`, clip);
  await context.close();
}
const TIP_CLIP = { x: 880, y: 0, width: 560, height: 180 };
SHOTS["state-hover"] = () =>
  stateShot("hover", {}, (p) => selector(p).hover(), TIP_CLIP);
SHOTS["state-focus"] = () =>
  stateShot(
    "focus",
    {},
    async (p) => {
      await p.keyboard.press("Shift");
      await selector(p).focus();
    },
    TIP_CLIP,
  );
SHOTS["state-open"] = () =>
  stateShot("open", {}, (p) => selector(p).click());
SHOTS["state-row-hover"] = () =>
  stateShot("row-hover", {}, async (p) => {
    await selector(p).click();
    await p.getByTestId(`selector-row-${keyOf(SHADOWMEND)}`).hover();
  });
SHOTS["state-row-focus"] = () =>
  stateShot("row-focus", {}, async (p) => {
    await p.keyboard.press("Shift");
    await selector(p).click();
    await p.keyboard.press("ArrowDown");
  });
SHOTS["state-mark-stale"] = () =>
  stateShot(
    "mark-stale",
    { current: TREEWALKER },
    (p) => selector(p).hover(),
    TIP_CLIP,
  );
SHOTS["state-mark-failed"] = () =>
  stateShot(
    "mark-failed",
    { current: OAKHEART },
    (p) => selector(p).hover(),
    TIP_CLIP,
  );
SHOTS["state-mark-session"] = async () => {
  const { page, context } = await open({
    width: 1440,
    height: 900,
    out: "expired",
    current: OBNOXIOUS,
  });
  await selector(page).hover();
  await page.waitForTimeout(300);
  await snap(page, "state-mark-session", TIP_CLIP);
  await context.close();
};
SHOTS["state-discord-tip"] = async () => {
  const { page, context } = await open({
    width: 1279,
    height: 900,
    chars: SIX,
    current: OBNOXIOUS,
  });
  await page.getByTestId("discord-link").hover();
  await page.waitForTimeout(300);
  await snap(page, "state-discord-tip", {
    x: 779,
    y: 0,
    width: 500,
    height: 150,
  });
  await context.close();
};
SHOTS["state-stale-row"] = () =>
  stateShot("stale-row", { chars: [OBNOXIOUS, TREEWALKER] }, (p) =>
    selector(p).click(),
  );
SHOTS["state-failed-row"] = () =>
  stateShot("failed-row", { chars: [OBNOXIOUS, OAKHEART] }, (p) =>
    selector(p).click(),
  );
SHOTS["state-one"] = () =>
  stateShot("one", { chars: [OBNOXIOUS], order: [] }, (p) =>
    selector(p).click(),
  );
SHOTS["state-none"] = () =>
  stateShot("none", { chars: [], current: null, order: [] }, (p) =>
    selector(p).click(),
  );
SHOTS["state-loading"] = () =>
  stateShot("loading", { out: "slow", current: OBNOXIOUS }, (p) =>
    selector(p).click(),
  );
SHOTS["state-load-error"] = () =>
  stateShot("load-error", { out: "failed", current: OBNOXIOUS }, (p) =>
    selector(p).click(),
  );
SHOTS["state-paste"] = () =>
  stateShot("paste", { chars: [OBNOXIOUS], order: [] }, async (p) => {
    await selector(p).click();
    await p.getByTestId("selector-paste").click();
  });
SHOTS["state-paste-error"] = () =>
  stateShot("paste-error", { chars: [OBNOXIOUS], order: [] }, async (p) => {
    await selector(p).click();
    await p.getByTestId("selector-paste").click();
    await p.getByTestId("selector-paste-input").fill("not an export");
    await p.getByTestId("selector-paste-use").click();
  });

for (const [name, width, height] of [
  ["1440", 1440, 360],
  ["390", 390, 420],
]) {
  SHOTS[`prehydration-${name}`] = async () => {
    const { page, context } = await open({
      width,
      height: 900,
      chars: SIX,
      current: OBNOXIOUS,
      blockJs: true,
    });
    const box = await page.getByRole("banner").boundingBox();
    console.log(`  prehydration-${name}: bar height ${box.height}`);
    await snap(page, `prehydration-${name}`, { x: 0, y: 0, width, height });
    await context.close();
  };
}

SHOTS["chosen-sequence"] = async () => {
  const { page, context } = await open({
    width: 1440,
    height: 900,
    chars: SIX,
    current: OBNOXIOUS,
    order: CHOSEN,
  });
  await snap(page, "chosen-1-before", { x: 0, y: 0, width: 1440, height: 420 });
  await selector(page).click();
  await snap(page, "chosen-2-open", { x: 0, y: 0, width: 1440, height: 700 });
  await page
    .getByTestId(`selector-row-${keyOf(QUICKSHOT)}`)
    .getByRole("button")
    .first()
    .click();
  await page.waitForLoadState("load");
  await page.getByTestId("character-selector").waitFor();
  await page.waitForTimeout(500);
  await snap(page, "chosen-3-after", { x: 0, y: 0, width: 1440, height: 420 });
  await page.goto(`${BASE}/tiers`);
  await page.getByTestId("character-selector").waitFor();
  await page.waitForTimeout(800);
  await snap(page, "chosen-4-tiers", { x: 0, y: 0, width: 1440, height: 900 });
  await context.close();
};

const wanted = process.argv.slice(2);
mkdirSync(OUT, { recursive: true });
browser = await chromium.launch();
for (const name of wanted.length > 0 ? wanted : Object.keys(SHOTS)) {
  if (SHOTS[name] === undefined) throw new Error(`no shot named ${name}`);
  await SHOTS[name]();
}
await browser.close();
