// Boards for design/specs/2026-10-28-discord-guided-cards.md: the /bis guided card at each step,
// the result it becomes, the raid post with its roster and day-before reminder, and the spec card, at the Discord
// desktop-client (640) and phone-client (390) widths. Mock only. It draws through the reviewers'
// render kit (web/tests/captures/discord-bot/discord-render.ts) and extends it here, in the mock,
// for what the kit lacks: a select showing its chosen option (Discord's default: true), a
// disabled select, one select opened at a time, and the reply chip over a reminder. The result
// frame is the real handler payload bis-top from the kit's payloads.json with the amendment's
// Change picks button added; the post frames start from the real post-open and post-webhook
// payloads with the roster of spec §P.9 (spec-icon application emoji drawn from the real icon files), and
// the Which spec tonight? card of §P.10 closes the set.
// Run this file with Node 22 or later (it strips the kit's TypeScript types); boards land here.
import {
  existsSync,
  readdirSync,
  readFileSync,
  writeFileSync,
} from "node:fs";
import path from "node:path";
import { fileURLToPath, pathToFileURL } from "node:url";

const HERE = path.dirname(fileURLToPath(import.meta.url));
const ROOT = path.resolve(HERE, "../../..");
const WEB = path.join(ROOT, "web");
const KIT = path.join(WEB, "tests/captures/discord-bot");
const { renderMessage } = await import(
  pathToFileURL(path.join(KIT, "discord-render.ts")).href
);
const { chromium } = await import(
  pathToFileURL(path.join(WEB, "node_modules/playwright/index.mjs")).href
);

const SITE = "https://foreversixty.gg";
const SITE_URL = /^https:\/\/foreversixty\.gg\//;
const CLOCK = "2026-12-10T22:00:00Z";
const CLIENTS = [
  { name: "desktop", width: 640 },
  { name: "phone", width: 390 },
];
const GOLD = 15055189;
const HUNTER = 0xaad372;
const EPHEMERAL = 64;

// Sentinels the mock's post-pass turns into the states the kit does not draw.
const CHOSEN = "⁣";
const DISABLED = "⁤";
const OPEN = "⁥";

const CLASSES = [
  ["Druid", "Balance, Feral, Feral Bear, Restoration"],
  ["Hunter", "Beast Mastery, Marksmanship, Survival"],
  ["Mage", "Arcane, Fire, Frost"],
  ["Paladin", "Holy, Protection, Retribution"],
  ["Priest", "Discipline, Holy, Shadow"],
  ["Rogue", "Assassination, Combat, Subtlety"],
  ["Shaman", "Elemental, Enhancement, Restoration"],
  ["Warlock", "Affliction, Demonology, Destruction"],
  ["Warrior", "Arms, Fury, Protection"],
];
const HUNTER_SPECS = [
  ["Beast Mastery", "DPS"],
  ["Marksmanship", "DPS"],
  ["Survival", "DPS"],
];
const RANGES = [
  ["Level 20", "Levels 20 to 29"],
  ["Level 30", "Levels 30 to 39"],
  ["Level 40", "Levels 40 to 49"],
  ["Level 50", "Levels 50 to 59"],
  ["Level 60", "Level 60, no buffs or consumables"],
  ["Phase 1 raid", "Level 60, raid buffs and consumables"],
];

const options = (rows, chosen) =>
  rows.map(([label, description]) => ({
    label,
    value: label.toLowerCase().replace(/ /g, "-"),
    description: label === chosen ? CHOSEN + description : description,
    default: label === chosen,
  }));

/** A string select: chosen shows its option, disabled greys it, open lists it (one per frame). */
function select(
  id,
  placeholder,
  rows,
  { chosen, disabled = false, open = false } = {},
) {
  let shown = chosen ? CHOSEN + chosen : placeholder;
  if (disabled) shown = DISABLED + placeholder;
  if (open) shown = OPEN + shown;
  return {
    type: 3,
    custom_id: id,
    placeholder: shown,
    disabled,
    options: disabled ? [] : options(rows, chosen),
  };
}
const row = (...components) => ({ type: 1, components });
const button = (label, style, id) => ({ type: 2, style, label, custom_id: id });
const link = (label, url) => ({ type: 2, style: 5, label, url });

const crest = (cls) => ({ url: SITE + "/icons/hd/crests/" + cls + ".webp" });

function card({
  title,
  description,
  color = GOLD,
  thumb,
  classSel,
  specSel,
  rangeSel,
  buttons,
}) {
  const embed = {
    author: { name: "Best in slot" },
    title,
    description,
    color,
    footer: { text: "Forever Sixty" },
    timestamp: CLOCK,
  };
  if (thumb) embed.thumbnail = crest(thumb);
  return {
    embeds: [embed],
    components: [row(classSel), row(specSel), row(rangeSel), row(...buttons)],
    flags: EPHEMERAL,
  };
}

const id = (part) => (state) => "gbis:1:" + part + ":" + state;
const C = id("c");
const S = id("s");
const R = id("r");
const GO = id("go");
const ME = id("me");
const LINK = SITE + "/v1/discord/link/start";

const LEAD_UNLINKED =
  "Choose a class, a spec and a range, then Show BiS. Link Discord and this starts from your character.";
const LEAD_READY = "Ready. Show BiS, or change a pick.";

const picked = (state, cls, spec, range, open) => ({
  classSel: select(C(state), "Choose a class", CLASSES, { chosen: cls }),
  specSel: select(S(state), "Choose a spec", HUNTER_SPECS, {
    chosen: spec,
    open: open === "spec",
  }),
  rangeSel: select(R(state), "Choose a range", RANGES, {
    chosen: range,
    open: open === "range",
  }),
});

const BOLD = "\x2a\x2a";
const bold = (s) => BOLD + s + BOLD;
const STALE_LINE =
  "This card is out of date. Run " + String.fromCharCode(47) + "bis again.";
const NUDGE = bold("Choose a spec first.") + " Then Show BiS.";
const MAIN_LEAD =
  "Picked from your main, " + bold("Zulmara") + ". Show BiS, or change a pick.";
const HUNTER_SLUG = "hunter";
const P1 = "Phase 1 raid";
const showRow = (state, linked) =>
  linked
    ? [button("Show BiS", 3, GO(state))]
    : [button("Show BiS", 3, GO(state)), link("Link Discord", LINK)];
const frame = (fid, title, message) => ({ id: fid, title, message });

const HUNTER_NAME = "Hunter";
const hunterCard = (title, description, state, spec, range, open, linked) =>
  card(
    Object.assign(picked(state, HUNTER_NAME, spec, range, open), {
      title,
      description,
      color: HUNTER,
      thumb: HUNTER_SLUG,
      buttons: showRow(state, linked),
    }),
  );

const PH_CLASS = "Choose a class";
const PH_CLASS_FIRST = "Choose a class first";
const PH_RANGE = "Choose a range";
const STATE0 = "::p1:";
const unlinkedFirst = card({
  title: PH_CLASS,
  description: LEAD_UNLINKED,
  classSel: select(C(STATE0), PH_CLASS, CLASSES, { open: true }),
  specSel: select(S(STATE0), PH_CLASS_FIRST, HUNTER_SPECS, { disabled: true }),
  rangeSel: select(R(STATE0), PH_RANGE, RANGES, { chosen: P1 }),
  buttons: showRow(STATE0, false),
});

const BM = "Beast Mastery";
const MM = "Marksmanship";
const BM_P1 = "Beast Mastery Hunter · level 60 Phase 1 raid";
const HUNTER_PICK = "Hunter · choose a spec";
const S_CLASS = "hunter::p1:";
const S_BM = "hunter:beast-mastery:p1:";
const S_MM = "hunter:marksmanship:50:";
const changed = hunterCard(
  "Marksmanship Hunter · level 50",
  LEAD_READY,
  S_MM,
  MM,
  "Level 50",
  "",
  true,
);
changed.components[3].components.push(button("Use Zulmara", 2, ME(S_MM)));
const stale = Object.assign(Object.create(null), {
  content: STALE_LINE,
  flags: EPHEMERAL,
});

const frames = [
  frame(
    "bis-card-1-class",
    "not linked: the card, the class list open",
    unlinkedFirst,
  ),
  frame(
    "bis-card-2-spec",
    "after the class pick: the spec list open",
    hunterCard(
      HUNTER_PICK,
      "Now choose a Hunter spec.",
      S_CLASS,
      undefined,
      P1,
      "spec",
      false,
    ),
  ),
  frame(
    "bis-card-3-range",
    "after the spec pick: the range list open",
    hunterCard(BM_P1, LEAD_READY, S_BM, BM, P1, "range", false),
  ),
  frame(
    "bis-card-4-ready",
    "every pick made: Show BiS",
    hunterCard(BM_P1, LEAD_READY, S_BM, BM, P1, "", false),
  ),
  frame(
    "bis-card-nudge",
    "Show BiS before a spec: the nudge, picks kept",
    hunterCard(HUNTER_PICK, NUDGE, S_CLASS, undefined, P1, "", false),
  ),
  frame(
    "bis-card-main",
    "linked: the card picked from the main",
    hunterCard(BM_P1, MAIN_LEAD, S_BM, BM, P1, "", true),
  ),
  frame(
    "bis-card-main-changed",
    "linked, a pick changed: Use Zulmara appears",
    changed,
  ),
  frame("bis-card-stale", "a press on a card of an older shape", stale),
];

const payloads = JSON.parse(
  readFileSync(path.join(KIT, "payloads.json"), "utf8"),
);
const payload = (name) =>
  structuredClone(payloads.find((p) => p.id === name).message);

// The result: the real handler's bis-top payload, plus Change picks before the link.
const result = payload("bis-top");
const [share, allSlots] = result.components[0].components;
result.components = [
  row(
    share,
    button("Change picks", 2, "gbis:1:chg:hunter:beast-mastery:p1:"),
    allSlots,
  ),
];
frames.push({
  id: "bis-result",
  title: "Show BiS: the card becomes the typed answer, in place",
  message: result,
});

// ---------------------------------------------------------------------------------------------
// Raid posts with the roster (spec §P.9). The roster is drawn by the rules the builder follows:
// four groups (Tanks, Healers, Melee, Ranged) as inline fields, each line "{spec icon emoji}
// {Name}", sorted by class in the game's order then by name, at most 16 lines a field
// and two fields a group; then Late and Bench inline, and Out full width as names A to Z.

/** The game's class order (classChoices); the roster sorts by it, then by name. */
const CLASS_ORDER = [
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
/** §P.9.1: the specs, from data/curated/specs.json; the emoji is the spec's own talent-tab icon. */
const SPECS = JSON.parse(
  readFileSync(path.join(ROOT, "data/curated/specs.json"), "utf8"),
);
/** The build whose icon folder is the source (web/public/data/{build}/icons/{icon}.webp). */
const ICON_BUILD = "1.60.1.70658";
const ICON_ROUTE = "/spec-icons/";
const emojiName = (sp) => "fs_" + sp.spec.replaceAll("-", "_");
const specRow = (cls, name) =>
  SPECS.find((sp) => sp.class_slug === cls && sp.name === name);
/** Mock application-emoji ids; the bot reads the real ones from GET /applications/{id}/emojis. */
const EMOJI_ID = Object.fromEntries(
  SPECS.map((sp, i) => [emojiName(sp), String(1301000000000000001n + BigInt(i))]),
);
const ICON_OF = Object.fromEntries(SPECS.map((sp) => [emojiName(sp), sp.icon]));
const MIDDOT = " · ";
const FIELD_LINES = 16;
const LIST_LINES = 10;
const ZWSP = String.fromCharCode(0x200b);
const NL = String.fromCharCode(10);
const LT = String.fromCharCode(60);
const GT = String.fromCharCode(62);
const titleCase = (s) => s[0].toUpperCase() + s.slice(1);
const nameOf = (r) => r.at(0);
const classOf = (r) => r.at(1);
const specOf = (r) => r.at(2);
const roleOf = (r) => r.at(3);
const stateOf = (r) => r.at(4);
const markOf = (r) => {
  const sp = specRow(classOf(r), specOf(r));
  return sp ? emojiName(sp) : undefined;
};

/** One roster line: the spec icon then the name (§P.9.3). */
const emojiLine = (r) =>
  LT + ":" + markOf(r) + ":" + EMOJI_ID[markOf(r)] + GT + " " + nameOf(r);
/** No emoji (a webhook, a missing emoji): the name, the full spec name and the class; no spec, the class. */
const textLine = (r) =>
  nameOf(r) +
  MIDDOT +
  (specOf(r) ? specOf(r) + " " : "") +
  titleCase(classOf(r));
const rosterLine = (r, emoji) =>
  emoji && markOf(r) ? emojiLine(r) : textLine(r);
const classRank = (r) => CLASS_ORDER.indexOf(classOf(r));
const byName = (a, b) =>
  nameOf(a).localeCompare(nameOf(b), "en", { sensitivity: "base" });
const byClassThenName = (a, b) =>
  classRank(a) === classRank(b) ? byName(a, b) : classRank(a) - classRank(b);
const sorted = (rows) => [...rows].sort(byClassThenName);
const field = (name, value, inline) => ({ name, value, inline });
const inState = (rows, state) =>
  sorted(rows.filter((r) => stateOf(r) === state));

/** Lines into n balanced chunks, the earlier chunks one longer. */
const chunks = (lines, n) =>
  Array.from({ length: n }, (_, i) =>
    lines.slice(
      Math.ceil((i * lines.length) / n),
      Math.ceil(((i + 1) * lines.length) / n),
    ),
  );
/** A capped list: at most max lines, the last becoming "and {n} more". */
const capped = (lines, max) =>
  lines.length <= max
    ? lines
    : [...lines.slice(0, max - 1), "and " + (lines.length - max + 1) + " more"];

/** §P.9.4: DPS specs that fight in melee (Survival is melee in Forever); a DPS line with no spec goes by class. */
const MELEE_SPECS = new Set([
  "Feral",
  "Survival",
  "Retribution",
  "Assassination",
  "Combat",
  "Subtlety",
  "Enhancement",
  "Arms",
  "Fury",
]);
const MELEE_CLASSES = new Set(["paladin", "rogue", "warrior"]);
const isMelee = (r) =>
  specOf(r) ? MELEE_SPECS.has(specOf(r)) : MELEE_CLASSES.has(classOf(r));
const GROUPS = [
  ["Tanks", (r) => roleOf(r) === "tank"],
  ["Healers", (r) => roleOf(r) === "healer"],
  ["Melee", (r) => (roleOf(r) === "dps" ? isMelee(r) : false)],
  ["Ranged", (r) => (roleOf(r) === "dps" ? !isMelee(r) : false)],
];
const FIELDS_PER_GROUP = 2;

/** One group as inline fields: its header on the first, U+200B on a continuation, capped. */
function groupFields(label, rows, emoji) {
  const lines = capped(
    rows.map((r) => rosterLine(r, emoji)),
    FIELD_LINES * FIELDS_PER_GROUP,
  );
  const header = label + MIDDOT + rows.length;
  return lines.length === 0
    ? [field(header, "None yet", true)]
    : chunks(lines, Math.ceil(lines.length / FIELD_LINES)).map((part, j) =>
        field(j === 0 ? header : ZWSP, part.join(NL), true),
      );
}

/** The In roster: Tanks, Healers, Melee, Ranged, each inline, sorted by class then name. */
const roleFields = (roster, emoji) =>
  GROUPS.flatMap(([label, inGroup]) =>
    groupFields(label, inState(roster.filter(inGroup), "in"), emoji),
  );

/** Late and Bench inline as capped lines, after the groups; Out full width as names A to Z. */
function offFields(roster, emoji) {
  const listed = (label, rows) =>
    field(
      label + MIDDOT + rows.length,
      capped(
        rows.map((r) => rosterLine(r, emoji)),
        LIST_LINES,
      ).join(NL),
      true,
    );
  const late = inState(roster, "late");
  const bench = inState(roster, "bench");
  const gone = roster.filter((r) => stateOf(r) === "out").sort(byName);
  return [
    ...(late.length ? [listed("Late", late)] : []),
    ...(bench.length ? [listed("Bench", bench)] : []),
    ...(gone.length
      ? [
          field(
            "Out" + MIDDOT + gone.length,
            gone.map(nameOf).join(", "),
            false,
          ),
        ]
      : []),
  ];
}

const START = 1796954400;
const SIZE = 40;
const stamp = (style) => LT + "t:" + START + ":" + style + GT;
const WHEN = stamp("F") + " (" + stamp("R") + ")";

/** The post's fields: When, Guild time, Signed up inline, then the roster. */
function rosterEmbedFields(roster, emoji) {
  const signed = roster.filter((r) =>
    ["in", "late"].includes(stateOf(r)),
  ).length;
  const head = [
    field("When", WHEN, true),
    field("Guild time", "8:00 pm CST", true),
    field(
      "Signed up",
      signed ? bold(signed + " of " + SIZE) : "None yet",
      true,
    ),
  ];
  const roles = signed ? roleFields(roster, emoji) : [];
  return [...head, ...roles, ...offFields(roster, emoji)];
}

// The example raid: Onyxia's Lair, Thursday 10 December, 8:00 pm guild time, 40 players.
const T = "tank";
const H = "healer";
const D = "dps";
/** [name, class, spec, role, state]: 33 of 40 (31 In, 2 Late), 3 Bench, 4 Out. */
const ROSTER_33 = [
  ["Kraggor", "warrior", "Protection", T, "in"],
  ["Ironvale", "warrior", "Protection", T, "in"],
  ["Bramblehide", "druid", "Feral Bear", T, "in"],
  ["Lightwell", "priest", "Holy", H, "in"],
  ["Dawnmere", "priest", "Holy", H, "in"],
  ["Embertide", "priest", "Discipline", H, "in"],
  ["Aurelith", "paladin", "Holy", H, "in"],
  ["Halcyra", "paladin", "Holy", H, "in"],
  ["Mossgrove", "druid", "Restoration", H, "in"],
  ["Vaelora", "druid", "Restoration", H, "in"],
  ["Sablewind", "shaman", "Restoration", H, "in"],
  ["Zulmara", "hunter", "Beast Mastery", D, "in"],
  ["Windtalon", "hunter", "Marksmanship", D, "in"],
  ["Brightshot", "hunter", "Marksmanship", D, "in"],
  ["Kelthar", "hunter", "Survival", D, "in"],
  ["Pyrewisp", "mage", "Fire", D, "in"],
  ["Cindrel", "mage", "Fire", D, "in"],
  ["Frostwhisk", "mage", "Frost", D, "in"],
  ["Stormvale", "mage", "Frost", D, "in"],
  ["Orivane", "mage", "Frost", D, "in"],
  ["Sundermark", "paladin", "Retribution", D, "in"],
  ["Nightbrand", "priest", "Shadow", D, "in"],
  ["Corvane", "rogue", "Combat", D, "in"],
  ["Duskmantle", "rogue", "Combat", D, "in"],
  ["Ravenmoor", "rogue", "Assassination", D, "in"],
  ["Ashenra", "shaman", "Enhancement", D, "in"],
  ["Grimhollow", "warlock", "Destruction", D, "in"],
  ["Talvorn", "warlock", "Destruction", D, "in"],
  ["Vexlash", "warlock", "Affliction", D, "in"],
  ["Wyrmbane", "warrior", "Fury", D, "in"],
  ["Thornquill", "druid", "Balance", D, "in"],
  ["Glimmerfen", "warlock", "Destruction", D, "late"],
  ["Hollowmere", "rogue", "Combat", D, "late"],
  ["Quillfeather", "hunter", "Marksmanship", D, "bench"],
  ["Marrowind", "hunter", "Marksmanship", D, "bench"],
  ["Selwyn", "mage", "Frost", D, "bench"],
  ["Fennrick", "warlock", "Affliction", D, "out"],
  ["Ysolde", "priest", "Shadow", D, "out"],
  ["Garrowe", "warrior", "Arms", D, "out"],
  ["Lunetta", "mage", "Fire", D, "out"],
];
/** Partly signed, the evening after it went up: 8 of 40, no healer yet. */
const ROSTER_8 = [
  ["Kraggor", "warrior", "Protection", T, "in"],
  ["Zulmara", "hunter", "Beast Mastery", D, "in"],
  ["Pyrewisp", "mage", "Fire", D, "in"],
  ["Corvane", "rogue", "Combat", D, "in"],
  ["Frostwhisk", "mage", "Frost", D, "in"],
  ["Vexlash", "warlock", "Affliction", D, "in"],
  ["Wyrmbane", "warrior", "Fury", D, "in"],
  ["Glimmerfen", "warlock", "Destruction", D, "late"],
  ["Garrowe", "warrior", "Arms", D, "out"],
];
/** Full: 40 of 40 (39 In, 1 Late), 4 tanks, 10 healers, 25 DPS; 3 Bench, 2 Out. */
const ROSTER_40 = [
  ...ROSTER_33.filter((r) => stateOf(r) === "in"),
  ["Stonebrow", "warrior", "Protection", T, "in"],
  ["Lumen", "priest", "Holy", H, "in"],
  ["Tidecall", "shaman", "Restoration", H, "in"],
  ["Hollowmere", "rogue", "Combat", D, "in"],
  ["Selwyn", "mage", "Frost", D, "in"],
  ["Ashgrave", "warlock", "Destruction", D, "in"],
  ["Quillfeather", "hunter", "Marksmanship", D, "in"],
  ["Briarthorn", "warrior", "Fury", D, "in"],
  ["Glimmerfen", "warlock", "Destruction", D, "late"],
  ["Marrowind", "hunter", "Marksmanship", D, "bench"],
  ["Fennrick", "warlock", "Affliction", D, "bench"],
  ["Embervale", "mage", "Fire", D, "bench"],
  ["Ysolde", "priest", "Shadow", D, "out"],
  ["Lunetta", "mage", "Fire", D, "out"],
];

const POSTED_CLOCK = "2026-12-06T03:00:00Z";
const PARTIAL_CLOCK = "2026-12-07T02:00:00Z";
const REMINDER_CLOCK = "2026-12-10T02:00:00Z";
const FULL_CLOCK = "2026-12-10T22:00:00Z";
const TOMORROW_TITLE = "Tomorrow: Onyxia's Lair · Thursday 10 December";
const TOMORROW_DESC = "The raid is tomorrow. Sign up with the buttons below.";

/** The post-open (or webhook) payload with the roster's fields, the clock and the Tomorrow state. */
function rosterPost(roster, clock, { tomorrow = false, webhook = false } = {}) {
  const post = payload(webhook ? "post-webhook" : "post-open");
  post.nonce = undefined;
  post.enforce_nonce = undefined;
  const e = post.embeds[0];
  e.fields = rosterEmbedFields(roster, webhook ? false : true);
  e.timestamp = clock;
  e.title = tomorrow ? TOMORROW_TITLE : e.title;
  e.description = tomorrow
    ? webhook
      ? e.description
      : TOMORROW_DESC
    : e.description;
  return post;
}
const timed = (id, title, clock, message) =>
  Object.assign(frame(id, title, message), { clock });

frames.push(
  timed(
    "post-scheduled",
    "posted on scheduling: nobody yet",
    POSTED_CLOCK,
    rosterPost([], POSTED_CLOCK),
  ),
);
frames.push(
  timed(
    "post-partial",
    "partly signed: 8 of 40, no healer yet",
    PARTIAL_CLOCK,
    rosterPost(ROSTER_8, PARTIAL_CLOCK),
  ),
);
frames.push(
  timed(
    "post-tomorrow",
    "the reminder as an edit of the post",
    REMINDER_CLOCK,
    rosterPost(ROSTER_33, REMINDER_CLOCK, { tomorrow: true }),
  ),
);
frames.push(
  timed(
    "post-full",
    "full: 40 of 40, bench and out",
    FULL_CLOCK,
    rosterPost(ROSTER_40, FULL_CLOCK),
  ),
);

// The reminder, 24 h before the start, when the channel has moved on: the post (Tomorrow, with
// its roster) above, and the reply under it with the post's own row.
const reminderText =
  bold("Onyxia's Lair is tomorrow") +
  ", " +
  WHEN +
  ". 33 of 40 signed up so far.";
frames.push({
  id: "post-reminder",
  title: "the day-before reminder, a reply under the post",
  clock: REMINDER_CLOCK,
  above: rosterPost(ROSTER_33, REMINDER_CLOCK, { tomorrow: true }),
  reply: "Forever Sixty  Tomorrow: Onyxia's Lair · Thursday 10 December",
  message: {
    content: reminderText,
    components: payload("post-open").components,
  },
});

// The reminder through a webhook (no bot, never edited): the line and the masked link, then the
// roster as it stands, the full spec name because a webhook draws no application emoji.
const raidURL = SITE + "/guild/us/normal/olympus-xxvii/raids/1";
const webhookReminder = rosterPost(ROSTER_33, REMINDER_CLOCK, {
  tomorrow: true,
  webhook: true,
});
webhookReminder.content =
  reminderText + " [Sign up on Forever Sixty](" + raidURL + ")";
frames.push(
  timed(
    "post-reminder-webhook",
    "the day-before reminder through a webhook",
    REMINDER_CLOCK,
    webhookReminder,
  ),
);

// Which spec tonight? (§P.10): Aurelith, a Paladin with no spec known, pressed In on the post.
const PALADIN = 0xf48cba;
const PALADIN_SPECS = [
  ["Holy", "Healer"],
  ["Protection", "Tank"],
  ["Retribution", "DPS"],
];
const SPEC_KEY = "us/normal/aurelith";
const specAsk = {
  embeds: [
    {
      title: "Which spec tonight?",
      description:
        "The pick signs " +
        bold("Aurelith") +
        " up as In for Onyxia's Lair, Thu 10 Dec, and is kept for the next raid.",
      color: PALADIN,
      thumbnail: crest("paladin"),
      footer: { text: "Forever Sixty" },
      timestamp: CLOCK,
    },
  ],
  components: [
    row(
      select("gspc:1:ask:1:in:" + SPEC_KEY, "Choose a spec", PALADIN_SPECS, {
        open: true,
      }),
    ),
  ],
  flags: EPHEMERAL,
};
frames.push(
  frame(
    "spec-ask-card",
    "In pressed, the spec unknown: the one-select card",
    specAsk,
  ),
);

const specDone = {
  content:
    "Aurelith is In as Healer for Onyxia's Lair, Thu 10 Dec. 34 of 40 signed up. Holy is kept for Aurelith.",
  components: [row(button("Change spec", 2, "gspc:1:chg:1:in:" + SPEC_KEY))],
  flags: EPHEMERAL,
};
frames.push(
  frame(
    "spec-ask-done",
    "the pick: the sign-up outcome in place, Change spec under it",
    specDone,
  ),
);

// ---------------------------------------------------------------------------------------------
// The real frames. payloads/<frame>.json is what the bot's own handlers and drainer answered
// (api/internal/discordbot/payload_dump_test.go, run with GUIDED_PAYLOADS_DIR); each is drawn
// again as <frame>-real through the same kit and post-pass, so a board and its real payload
// can be laid side by side. Frames with no board (char, signup, officer cards) are drawn too.
const PAYLOADS = path.join(HERE, "payloads");
/** Which select the board shows open, by frame: the real frame opens the same one. */
const OPEN_SELECT = {
  "bis-card-1-class": 0,
  "bis-card-2-spec": 1,
  "bis-card-3-range": 2,
};

/** Discord's own look of a select: the default option shown, a disabled select greyed, one list open. */
function realSelect(c, open) {
  const chosen = (c.options ?? []).find((o) => o.default);
  let shown = c.placeholder ?? "";
  if (c.disabled) shown = DISABLED + shown;
  else if (chosen) shown = CHOSEN + chosen.label;
  if (open) shown = OPEN + shown;
  const opts = c.disabled
    ? []
    : (c.options ?? []).map((o) =>
        o.default ? { ...o, description: CHOSEN + (o.description ?? "") } : o,
      );
  return { ...c, placeholder: shown, options: opts };
}

function realMessage(message, openIndex) {
  let seen = 0;
  const components = (message.components ?? []).map((r) => ({
    ...r,
    components: r.components.map((c) =>
      c.type === 3 ? realSelect(c, seen++ === openIndex) : c,
    ),
  }));
  return { ...message, components };
}

function realFrame(file) {
  const raw = JSON.parse(readFileSync(path.join(PAYLOADS, file), "utf8"));
  if (!raw.message) return undefined; // a modal: nothing to draw
  const board = frames.find((f) => f.id === raw.id);
  const open = OPEN_SELECT[raw.id] ?? -1;
  return {
    id: raw.id + "-real",
    title: "real payload: " + raw.title,
    clock: board?.clock ?? raw.clock,
    above: raw.above ? realMessage(raw.above, -1) : undefined,
    reply: board?.reply,
    message: realMessage(raw.message, open),
  };
}

if (existsSync(PAYLOADS))
  for (const file of readdirSync(PAYLOADS)
    .filter((n) => n.endsWith(".json"))
    .sort()) {
    const real = realFrame(file);
    if (real) frames.push(real);
  }

const EXTRA_CSS =
  "<style>" +
  ".select .chosen{color:#dbdee1}" +
  ".select.dis{opacity:.5;cursor:not-allowed}" +
  ".option.sel{background:#404249;border-radius:3px}" +
  '.option.sel b::after{content:"  \\2713";color:#00a8fc}' +
  ".emoji{width:20px;height:20px;vertical-align:-5px;object-fit:contain}" +
  "#msg >.wrap + .wrap{padding-top:4px}" +
  "</style>";

const CLOSED_LIST =
  /<div class="select open"><span>(?!⁥)([^<]*)<\/span><span>⌃<\/span><\/div><div class="options">(?:<div class="option">.*?<\/div>)*<\/div>/g;

const OPEN_MARK = /<span>⁥/g;
const DISABLED_SELECT = /<div class="select(?: open)?"><span>⁤([^<]*)<\/span>/g;
const CHOSEN_VALUE = /<span>⁣([^<]*)<\/span>/g;
const CHOSEN_OPTION = /<div class="option"><b>([^<]*)<\/b><small>⁣/g;
const EMOJI = /&lt;:(fs_[a-z_]+):\d+&gt;/g;
const emojiImg = (_m, name) =>
  '<img class="emoji" src="' +
  SITE +
  ICON_ROUTE +
  ICON_OF[name] +
  '.webp" alt="">';
const ICON_PATH = /^\/data\/([^/]+)\/(icons\/.+)/;
const SPAN = "<span>";
const SPAN_END = "</span>";
const STYLE_END = "</style>";
const closedSelect = (_m, label) =>
  '<div class="select"><span>' + label + "</span><span>⌄</span></div>";
const disabledSelect = (_m, label) =>
  '<div class="select dis"><span>' + label + SPAN_END;
const chosenValue = (_m, label) => '<span class="chosen">' + label + SPAN_END;
const chosenOption = (_m, label) =>
  '<div class="option sel"><b>' + label + "</b><small>";

/** Draws what the kit lacks: the chosen value, the disabled select, one list open, the spec-icon emoji. */
const postPass = (html) =>
  html
    .replace(CLOSED_LIST, closedSelect)
    .replace(OPEN_MARK, SPAN)
    .replace(DISABLED_SELECT, disabledSelect)
    .replace(CHOSEN_VALUE, chosenValue)
    .replace(CHOSEN_OPTION, chosenOption)
    .replace(EMOJI, emojiImg)
    .replace(STYLE_END, STYLE_END + EXTRA_CSS);

const BUILDS = path.join(ROOT, "data", "builds");
const PUBLIC = path.join(WEB, "public");
const iconFile = (match) => path.join(BUILDS, match[1], match[2]);
const localFile = (url) =>
  url.pathname.startsWith(ICON_ROUTE)
    ? path.join(PUBLIC, "data", ICON_BUILD, "icons", url.pathname.slice(ICON_ROUTE.length))
    : ICON_PATH.test(url.pathname)
    ? iconFile(ICON_PATH.exec(url.pathname))
    : path.join(PUBLIC, url.pathname);

const serve = async (route) => {
  const file = localFile(new URL(route.request().url()));
  return existsSync(file)
    ? route.fulfill({ path: file, contentType: IMAGE_TYPE })
    : route.abort();
};
const IMAGE_TYPE = "image/webp";

const WRAP_OPEN = '<div class="wrap" id="msg">';
const PLAIN_WRAP = '<div class="wrap">';
/** A frame with a message above it (the post over its reply): both bodies under one #msg. */
const stacked = (top, bottom) =>
  top.replace(WRAP_OPEN, '<div id="msg">' + PLAIN_WRAP) +
  bottom.slice(bottom.indexOf(WRAP_OPEN)).replace(WRAP_OPEN, PLAIN_WRAP) +
  "</div>";
const render = (f, client, message, reply) =>
  renderMessage(
    {
      id: f.id,
      row: ROW,
      title: f.title,
      surface: SURFACE,
      clock: f.clock ?? CLOCK,
      message,
    },
    { width: client.width, selectOpen: true, usedCommand: reply },
  );
const frameHtml = (f, client) =>
  f.above
    ? stacked(render(f, client, f.above), render(f, client, f.message, f.reply))
    : render(f, client, f.message, f.reply);

const shoot = async (page, client, f) => {
  const html = postPass(frameHtml(f, client));
  const name = f.id + DASH + client.name;
  writeFileSync(path.join(HERE, name + HTML_EXT), html);
  await page.setViewportSize({ width: client.width, height: 600 });
  await page.setContent(html, { waitUntil: LOAD });
  await page.locator(MSG).screenshot({ path: path.join(HERE, name + PNG_EXT) });
};
const ROW = "G";
const SURFACE = "ephemeral";
const DASH = "-";
const HTML_EXT = ".html";
const PNG_EXT = ".png";
const LOAD = "load";
const MSG = "#msg";

// MOCK_JSON_DIR=<dir> also writes each frame's message as JSON, for diffing board against payload.
if (process.env.MOCK_JSON_DIR)
  for (const f of frames)
    writeFileSync(
      path.join(process.env.MOCK_JSON_DIR, f.id + ".json"),
      JSON.stringify({ message: f.message, above: f.above }, null, 2),
    );

const browser = await chromium.launch();
const page = await browser.newPage({ deviceScaleFactor: 2 });
await page.route(SITE_URL, serve);
for (const client of CLIENTS)
  for (const f of frames) await shoot(page, client, f);
await browser.close();
console.log("wrote", frames.length * CLIENTS.length, "boards to", HERE);
