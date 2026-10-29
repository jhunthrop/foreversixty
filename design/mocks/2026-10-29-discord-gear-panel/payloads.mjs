// The Components V2 payloads of design/specs/2026-10-29-discord-gear-panel.md, built from the
// real band (data.mjs). These are the reference payloads: the handler tests compare the bot's
// own output against the same shapes, and the boards draw exactly these.
import {
  BUILD, PAGE1, PAGE2, QUALITY_WORDS, SLOT_LABELS, bisRows, charRows, signedDPS,
} from "./data.mjs";

export const SITE = "https://foreversixty.gg";
export const API = "https://api.foreversixty.gg";
export const IS_COMPONENTS_V2 = 1 << 15;
export const EPHEMERAL = 1 << 6;
export const HUNTER = 0xaad372;
export const T = { ROW: 1, BUTTON: 2, SELECT: 3, SECTION: 9, TEXT: 10, THUMB: 11, GALLERY: 12, SEP: 14, CONTAINER: 17 };

const crest = (cls) => SITE + "/icons/hd/crests/" + cls + ".webp";
export const tileURL = (id) => `${API}/v1/discord/tiles/${BUILD}/${id}.png`;
const PAGE = SITE + "/bis/hunter/beast-mastery?faction=alliance&preset=raid#band-alliance-60";
const itemLink = (id) =>
  `${SITE}/bis/hunter/beast-mastery?faction=alliance&item=${id}&preset=raid#band-alliance-60`;
const CHAR_URL = SITE + "/character/us/normal/zulmara";

const text = (content) => ({ type: T.TEXT, content });
const sep = (spacing = 1) => ({ type: T.SEP, divider: true, spacing });
const thumb = (url, description) => ({ type: T.THUMB, media: { url }, description });
const section = (content, accessory) => ({ type: T.SECTION, components: [text(content)], accessory });
const row = (...components) => ({ type: T.ROW, components });
const button = (label, style, custom_id) => ({ type: T.BUTTON, style, label, custom_id });
const link = (label, url) => ({ type: T.BUTTON, style: 5, label, url });

const tileAlt = (it, slot) => `${it.name}, ${QUALITY_WORDS[it.quality] ?? ""} ${slot.toLowerCase()} item`.replace("  ", " ");
/** The subtext line: quality word, item level, source (spec §A.3, line 3). */
const facts = (it, source) => "-# " + [QUALITY_WORDS[it.quality], it.ilvl > 0 ? "item level " + it.ilvl : "", source].filter(Boolean).join(" · ");
/** The slot line: slot, then the gain (spec §A.3, line 2). */
const slotLine = (slot, gain) => [slot, gain].filter(Boolean).join(" · ");

/** One slot of the /bis panel: the item link, slot · quality · item level, gain · source. */
function bisSlot(r, rows) {
  const label = SLOT_LABELS[r.slot];
  if (r.empty === "two_hand_equipped") {
    const main = rows.get("main_hand");
    return text(`**Off hand** · none: ${main.name} is two-handed.`);
  }
  if (r.empty) return text(`**${label}** · no pick for this range.`);
  const gain = r.gain === "tie" ? "ties the next best" : r.gain;
  return section(`**[${r.name}](${itemLink(r.id)})**\n${slotLine(label, gain)}\n${facts(r, r.source)}`, thumb(tileURL(r.id), tileAlt(r, label)));
}

const BIS_HEAD = "### Beast Mastery Hunter\nBest in slot · level 60 Phase 1 raid · Alliance";
const PAGE_LINE = {
  1: "-# Head to Legs, 9 of 17 slots. Each gain is over the next best item.",
  2: "-# Feet to Ranged, 8 of 17 slots. Each gain is over the next best item.",
};
const SHARE = (token) => button("Share in channel", 2, "share:" + token);
const TOKEN = "ETsckh0MUHCXRpYAXQoDPcxTEKAbZTDvN6CED4pY-QQ";

export function bisPanel(d, page) {
  const rows = bisRows(d);
  const slots = (page === 1 ? PAGE1 : PAGE2).map((s) => bisSlot(rows.get(s), rows));
  const other = page === 1 ? "Feet to Ranged" : "Head to Legs";
  return {
    flags: IS_COMPONENTS_V2 | EPHEMERAL,
    components: [
      {
        type: T.CONTAINER,
        accent_color: HUNTER,
        components: [
          section(BIS_HEAD + "\n" + PAGE_LINE[page], thumb(crest("hunter"), "Hunter crest")),
          sep(),
          ...slots,
        ],
      },
      row(
        SHARE(TOKEN),
        button(other, 2, `gbis:1:pg:hunter:beast-mastery:p1::${page === 1 ? 2 : 1}`),
        button("Change picks", 2, "gbis:1:chg:hunter:beast-mastery:p1:"),
        link("All 17 slots", PAGE),
      ),
    ],
    allowed_mentions: { parse: [] },
  };
}

/** /bis slot:legs: the pick, then the next two with their deltas against it. */
export function bisSlotPanel(d, slotKey) {
  const rows = bisRows(d);
  const r = rows.get(slotKey);
  const label = SLOT_LABELS[slotKey];
  const alts = r.alternatives.slice(0, 2).map((a) => {
    const delta = Math.abs(a.delta) < 0.05 ? "same DPS as the pick" : signedDPS(a.delta);
    return section(
      `**[${a.name}](${itemLink(a.id)})**\n${slotLine(label, delta)}\n${facts(a, a.source)}`,
      thumb(tileURL(a.id), tileAlt(a, label)),
    );
  });
  const pickGain = r.gain === "tie" ? "ties the next best" : r.gain;
  return {
    flags: IS_COMPONENTS_V2 | EPHEMERAL,
    components: [
      {
        type: T.CONTAINER,
        accent_color: HUNTER,
        components: [
          section(`### ${label} · Beast Mastery Hunter\nBest in slot · level 60 Phase 1 raid · Alliance`, thumb(crest("hunter"), "Hunter crest")),
          sep(),
          section(`**[${r.name}](${itemLink(r.id)})**\n${slotLine(label, pickGain)}\n${facts(r, r.source)}`, thumb(tileURL(r.id), tileAlt(r, label))),
          sep(),
          text("-# The next best"),
          ...alts,
        ],
      },
      row(SHARE(TOKEN), button("Change picks", 2, "gbis:1:chg:hunter:beast-mastery:p1:" + slotKey), link("Open on Forever Sixty", PAGE)),
    ],
    allowed_mentions: { parse: [] },
  };
}

const CHAR_HEAD = `### [Zulmara](${CHAR_URL})\nLevel 60 Beast Mastery Hunter · Olympus XXVII`;
export const SYNCED_UNIX = Date.UTC(2026, 11, 8, 22) / 1000; // two days before the boards' clock

export function averageItemLevel(d, crows) {
  let total = 0, n = 0;
  for (const [, c] of crows) if (c.worn) { total += c.worn.ilvl; n++; }
  return (total / n).toFixed(1);
}

function upgradeSection(c) {
  const label = SLOT_LABELS[c.slot];
  const p = c.pick;
  const url = CHAR_URL + "#gear-slot-" + c.slot;
  const gain = c.empty ? "fills an empty slot" : c.gain;
  const replaces = c.empty ? "" : `\n-# Replaces ${c.worn.name}, item level ${c.worn.ilvl}`;
  return section(`**[${p.name}](${url})**\n${slotLine(label, gain)}\n${facts(p, p.source)}${replaces}`, thumb(tileURL(p.id), tileAlt(p, label)));
}

/** Upgrades in the order the panel lists them: empty slots first, then the largest gain. */
export function upgradesOf(d) {
  const crows = charRows(d, bisRows(d));
  const ups = [...crows.values()].filter((c) => c.empty || (c.gainValue ?? 0) >= 0.05);
  ups.sort((a, b) => (a.empty ? -1 : b.empty ? 1 : b.gainValue - a.gainValue));
  return { crows, ups };
}

const OWN_SELECT = {
  type: T.SELECT,
  custom_id: "gchar:1:pick:",
  placeholder: "Another of your characters",
  options: [
    { label: "Zulmara", value: "Zulmara", description: "Beast Mastery Hunter · Level 60", default: true },
    { label: "Duskstrider", value: "Duskstrider", description: "Fury Warrior · Level 58" },
  ],
};
const CHAR_ROW = () =>
  row(SHARE("6C6UoNxef7TrDI2_1c5droEAp4_C6QFKGg22VqAquZ4"), button("Someone else", 2, "gchar:1:ask:"), link("Character page", CHAR_URL));

export function charPanel(d, { noGear = false } = {}) {
  const { crows, ups } = upgradesOf(d);
  const shown = ups.slice(0, 5);
  const more = ups.length - shown.length;
  const facts = noGear
    ? "-# Readiness: Not synced"
    : `-# Item level ${averageItemLevel(d, crows)} · Readiness: Ready · Synced <t:${SYNCED_UNIX}:R>`;
  const body = noGear
    ? [text("**Top upgrades** · level 60 Phase 1 raid\nNo gear shared yet. Upgrades show once the character is synced.")]
    : [
        text("**Top upgrades** · level 60 Phase 1 raid\n-# By DPS gained, against the guide build."),
        ...shown.map(upgradeSection),
        ...(more > 0 ? [text(`-# ${more} more ${more === 1 ? "upgrade" : "upgrades"} on the character page.`)] : []),
      ];
  return {
    flags: IS_COMPONENTS_V2 | EPHEMERAL,
    components: [
      {
        type: T.CONTAINER,
        accent_color: HUNTER,
        components: [section(CHAR_HEAD + "\n" + facts, thumb(crest("hunter"), "Hunter crest")), sep(), ...body],
      },
      CHAR_ROW(),
      row(OWN_SELECT),
    ],
    allowed_mentions: { parse: [] },
  };
}

// ---- the shared messages: the sheet attached, the link to the full page ----
export function bisSheetAlt(d) {
  const rows = bisRows(d);
  const parts = Object.keys(SLOT_LABELS).map((s) => {
    const r = rows.get(s);
    if (r.empty === "two_hand_equipped") return "Off hand: none, two-handed main hand";
    return `${SLOT_LABELS[s]}: ${r.name}`;
  });
  return "Best in slot sheet, Beast Mastery Hunter, level 60 Phase 1 raid, Alliance. " + parts.join("; ") + ".";
}

export function charSheetAlt(d) {
  const { crows } = upgradesOf(d);
  const parts = Object.keys(SLOT_LABELS).map((s) => {
    const c = crows.get(s);
    if (c.twoHand) return "Off hand: none, two-handed main hand";
    if (c.empty) return `${SLOT_LABELS[s]}: empty`;
    return `${SLOT_LABELS[s]}: ${c.worn.name}${c.gain ? ", " + c.gain + " available" : ""}`;
  });
  return "Character sheet, Zulmara, level 60 Beast Mastery Hunter, Olympus XXVII. " + parts.join("; ") + ".";
}

const SHEET = "attachment://sheet.png";

export function bisShare(d) {
  const alt = bisSheetAlt(d);
  return {
    flags: IS_COMPONENTS_V2,
    attachments: [{ id: 0, filename: "sheet.png", description: alt }],
    components: [
      {
        type: T.CONTAINER,
        accent_color: HUNTER,
        components: [text(BIS_HEAD), { type: T.GALLERY, items: [{ media: { url: SHEET }, description: alt }] }],
      },
      row(link("All 17 slots", PAGE)),
    ],
    allowed_mentions: { parse: [] },
  };
}

export function charShare(d) {
  const alt = charSheetAlt(d);
  return {
    flags: IS_COMPONENTS_V2,
    attachments: [{ id: 0, filename: "sheet.png", description: alt }],
    components: [
      {
        type: T.CONTAINER,
        accent_color: HUNTER,
        components: [
          text(CHAR_HEAD),
          { type: T.GALLERY, items: [{ media: { url: SHEET }, description: alt }] },
          text("-# Gold figures are the DPS the level 60 Phase 1 raid pick adds in that slot."),
        ],
      },
      row(link("Character page", CHAR_URL)),
    ],
    allowed_mentions: { parse: [] },
  };
}

/** The share's fallback when the sheet cannot render: the visitor's panel, no image. */
export function bisShareNoSheet(d) {
  const p = bisPanel(d, 1);
  return { ...p, flags: IS_COMPONENTS_V2, components: [p.components[0], row(link("All 17 slots", PAGE))] };
}

// ---- budget ----
/** Every component in the tree, nested ones included (Discord's 40 counts them all). */
export function countComponents(list) {
  let n = 0;
  const walk = (c) => {
    n++;
    for (const k of c.components ?? []) walk(k);
    if (c.accessory) walk(c.accessory);
  };
  list.forEach(walk);
  return n;
}
export function textLength(list) {
  let n = 0;
  const walk = (c) => {
    if (c.type === T.TEXT) n += c.content.length;
    for (const k of c.components ?? []) walk(k);
  };
  list.forEach(walk);
  return n;
}
