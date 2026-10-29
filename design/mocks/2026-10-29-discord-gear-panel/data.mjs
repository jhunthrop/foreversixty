// Real data for the gear-panel boards (design/specs/2026-10-29-discord-gear-panel.md): the active
// build's Beast Mastery Hunter band at level 60 Phase 1 raid, Alliance, its 17 picks with the
// class item table's quality, item level and icon, and each source line resolved against the
// build's loot.json the way the bot does (api/internal/discordbot/item_source.go, ported here
// for the mock only). Nothing is invented except the example character's worn gear, which is
// drawn from the band's own alternatives so every gain is the band's measured delta.
import { readFileSync } from "node:fs";
import path from "node:path";

export const BUILD = "1.60.1.70658";
export const SPEC = "hunter-beast-mastery";
export const CLASS = "hunter";
export const FACTION = "alliance";

export function loadData(root) {
  const dir = path.join(root, "data", "builds", BUILD);
  const json = (p) => JSON.parse(readFileSync(path.join(dir, p), "utf8"));
  const band = json(path.join("bis", SPEC + ".json")).bands.find(
    (b) => b.band === 60 && b.preset === "raid" && b.faction === FACTION,
  );
  const items = new Map(json(path.join("items", CLASS + ".json")).items.map((i) => [i.id, i]));
  const loot = json("loot.json");
  return { band, items, loot, iconDir: path.join(dir, "icons") };
}

// ---- the bot's source line (item_source.go + copy_source.go), ported ----
const PVP_TITLES = {
  alliance: ["Private", "Corporal", "Sergeant", "Master Sergeant", "Sergeant Major", "Knight", "Knight-Lieutenant",
    "Knight-Captain", "Knight-Champion", "Lieutenant Commander", "Commander", "Marshal", "Field Marshal", "Grand Marshal"],
  horde: ["Scout", "Grunt", "Sergeant", "Senior Sergeant", "First Sergeant", "Stone Guard", "Blood Guard", "Legionnaire",
    "Centurion", "Champion", "Lieutenant General", "General", "Warlord", "High Warlord"],
};
const KIND_WORDS = { raid: "Raids", dungeon: "Dungeons", world: "World bosses", world_drop: "World drops",
  zone: "Zone drops", vendor: "Vendors", crafted: "Crafted", rep: "Reputation", pvp: "PvP" };
const cap = (s) => (s ? s[0].toUpperCase() + s.slice(1) : s);
const itemsOf = (s) => [...(s.items ?? []), ...(s.trash ?? []), ...(s.bosses ?? []).flatMap((b) => b.items ?? [])];
const names = (s, id) => itemsOf(s).includes(id);
const findLoot = (loot, kind, id) => loot.sources.find((s) => s.kind === kind && names(s, id));
const bossOf = (s, id) => (s.bosses ?? []).find((b) => (b.items ?? []).includes(id));
function chanceOf(s, id) {
  let ch = s.item_chances;
  const b = bossOf(s, id);
  if (b && b.item_chances) ch = b.item_chances;
  return Math.max((ch ?? {})[String(id)] ?? 0, 0);
}
function bestChance(loot, kind, id) {
  let best, bestAt = -1, found = false;
  for (const s of loot.sources) {
    if (s.kind !== kind || !names(s, id)) continue;
    let at = chanceOf(s, id);
    if (at === 0) at = -1;
    if (!found || at > bestAt) [best, bestAt, found] = [s, at, true];
  }
  return best;
}
const chanceLine = (c, from) => (c < 1 ? "<1" : String(Math.round(c))) + "% from " + from;
function pvpLine(rank, faction) {
  const t = PVP_TITLES[faction][rank - 5];
  return t ? `PvP rank ${rank} · ${t} · ${cap(faction)}` : `PvP rank ${rank} · ${cap(faction)}`;
}
const realFaction = (f) => f === "alliance" || f === "horde";
function resolved(loot, id, kind, faction) {
  switch (kind) {
    case "quest": {
      const qs = loot.quests[String(id)] ?? [];
      if (!qs.length) return undefined;
      const q = qs.find((o) => o.faction === faction) ?? qs[0];
      // The level a player must be to take the quest (min_level), not the quest's own level.
      return "Quest: " + q.name + (q.min_level > 0 ? " · requires level " + q.min_level : "");
    }
    case "dungeon":
    case "raid": {
      const s = bestChance(loot, kind, id);
      if (!s) return undefined;
      const b = bossOf(s, id);
      const bossName = b ? b.name || "Unnamed source in " + s.name : "";
      const c = chanceOf(s, id);
      if (c > 0) return chanceLine(c, b ? bossName : s.name + " trash");
      return b ? s.name + " · " + bossName : s.name;
    }
    case "crafted": {
      const s = findLoot(loot, "crafted", id);
      if (!s) return undefined;
      return !s.profession || s.profession.toLowerCase() === s.name.toLowerCase()
        ? "Crafted: " + s.name
        : "Crafted: " + cap(s.profession) + " · " + s.name;
    }
    case "vendor": {
      const p = findLoot(loot, "pvp", id);
      if (p && p.rank != null && realFaction(p.faction)) return pvpLine(p.rank, p.faction);
      const s = findLoot(loot, "vendor", id);
      if (!s) return undefined;
      if (s.rank != null && realFaction(s.faction)) return pvpLine(s.rank, s.faction);
      return "Vendor: " + s.name;
    }
    case "rep": {
      const s = findLoot(loot, "rep", id);
      if (!s) return undefined;
      return s.standing ? `${s.name} (${s.standing})` : s.name;
    }
    case "zone":
    case "world": {
      const s = findLoot(loot, kind, id);
      if (!s) return undefined;
      const c = chanceOf(s, id);
      return c > 0 ? chanceLine(c, s.name) : s.name;
    }
    case "pvp": {
      const s = findLoot(loot, "pvp", id);
      if (!s || s.rank == null || !realFaction(s.faction)) return undefined;
      return pvpLine(s.rank, s.faction);
    }
    case "world_drop": {
      const s = findLoot(loot, "world_drop", id);
      if (!s) return undefined;
      return s.level_min != null && s.level_max != null
        ? `World drop (BoE) · levels ${s.level_min}-${s.level_max}`
        : "World drop (BoE)";
    }
  }
  return undefined;
}
export function sourceLine(loot, id, kind, faction, raw) {
  if (!kind) return raw;
  return resolved(loot, id, kind, faction) ??
    (kind === "quest" ? (faction === "horde" ? "Horde quest" : "Alliance quest") : KIND_WORDS[kind] ?? kind);
}

// ---- the sheet's short form (spec §B.3 fit rule, step 2): real abbreviations only ----
// The instance names players shorten in chat and on every guide; nothing else is abbreviated.
export const INSTANCE_SHORT = {
  "Blackrock Depths": "BRD", "Lower Blackrock Spire": "LBRS", "Upper Blackrock Spire": "UBRS",
  "Stratholme": "Strat", "Scholomance": "Scholo", "Dire Maul": "Dire Maul", "Molten Core": "MC",
  "Blackwing Lair": "BWL", "Onyxia's Lair": "Ony", "Zul'Gurub": "ZG", "Ruins of Ahn'Qiraj": "AQ20",
  "Temple of Ahn'Qiraj": "AQ40", "Naxxramas": "Naxx", "Sunken Temple": "ST", "Zul'Farrak": "ZF",
  "Maraudon": "Mara", "Scarlet Monastery": "SM", "Razorfen Downs": "RFD", "Razorfen Kraul": "RFK",
  "Shadowfang Keep": "SFK", "Blackfathom Deeps": "BFD", "Wailing Caverns": "WC", "Ragefire Chasm": "RFC",
  "Uldaman": "Uldaman", "Gnomeregan": "Gnomer",
};
/** The short form of a source line: the quest's requirement without its title, an instance by
 *  its abbreviation, a PvP rank without its title, a recipe by its profession. */
export function sourceShort(loot, id, kind, faction, full) {
  if (kind === "quest") {
    const m = / · requires level (\d+)$/.exec(full ?? "");
    return m ? "Quest · requires level " + m[1] : "Quest";
  }
  if (kind === "pvp" || (full ?? "").startsWith("PvP rank ")) {
    const m = /^PvP rank (\d+) · .* · (Alliance|Horde)$/.exec(full ?? "");
    return m ? `PvP rank ${m[1]} · ${m[2]}` : full;
  }
  if (kind === "crafted") return (full ?? "").replace(/^Crafted: ([^·]+) · .*$/, "Crafted: $1");
  let out = full ?? "";
  for (const [long, short] of Object.entries(INSTANCE_SHORT)) out = out.replace(long, short);
  return out;
}

// ---- slots ----
export const SLOT_LABELS = {
  head: "Head", neck: "Neck", shoulder: "Shoulder", back: "Back", chest: "Chest", wrist: "Wrist",
  hands: "Hands", waist: "Waist", legs: "Legs", feet: "Feet", finger1: "Ring 1", finger2: "Ring 2",
  trinket1: "Trinket 1", trinket2: "Trinket 2", main_hand: "Main hand", off_hand: "Off hand", ranged: "Ranged",
};
// The in-game character frame's columns (shirt and tabard have no BiS row; the sheet draws
// their sockets empty, spec §B.3).
export const LEFT = ["head", "neck", "shoulder", "back", "chest", "shirt", "tabard", "wrist"];
export const RIGHT = ["hands", "waist", "legs", "feet", "finger1", "finger2", "trinket1", "trinket2"];
export const BOTTOM = ["main_hand", "off_hand", "ranged"];
export const PAGE1 = ["head", "neck", "shoulder", "back", "chest", "wrist", "hands", "waist", "legs"];
export const PAGE2 = ["feet", "finger1", "finger2", "trinket1", "trinket2", "main_hand", "off_hand", "ranged"];
export const QUALITY_WORDS = ["Poor", "Common", "Uncommon", "Rare", "Epic", "Legendary"];

const MINUS = "−";
export function signedDPS(v) {
  const r = Math.round(v * 10) / 10;
  return (r < 0 ? MINUS : "+") + Math.abs(r).toFixed(1) + " DPS";
}

/** One item as the panel and the sheet see it. */
export function itemFacts(d, id, fallbackName) {
  const it = d.items.get(id);
  return {
    id,
    name: it?.name ?? fallbackName,
    quality: it?.quality ?? -1,
    ilvl: it?.item_level ?? 0,
    icon: it?.icon ?? "inv_misc_questionmark",
  };
}

/** The band's 17 rows: pick, gain over the next best, source line; off hand two-handed. */
export function bisRows(d) {
  const out = new Map();
  for (const s of d.band.slots) {
    if (s.empty_reason === "two_hand_equipped" || !s.item_id) {
      out.set(s.slot, { slot: s.slot, empty: s.empty_reason ?? "none" });
      continue;
    }
    const alt = (s.alternatives ?? [])[0];
    // A gain shows only when there is one: a tie (under 0.05) or no alternative shows none.
    const gain = alt && Math.abs(alt.dps_delta) >= 0.05 ? signedDPS(-alt.dps_delta) : undefined;
    const source = sourceLine(d.loot, s.item_id, s.source_kind, FACTION, s.source);
    out.set(s.slot, {
      slot: s.slot,
      ...itemFacts(d, s.item_id, s.item_name),
      gain,
      unmodelled: s.effect_unmodelled === true,
      source,
      sourceShort: sourceShort(d.loot, s.item_id, s.source_kind, FACTION, source),
      alternatives: (s.alternatives ?? []).map((a) => ({
        ...itemFacts(d, a.item_id, a.item_name),
        delta: a.dps_delta,
        unmodelled: a.effect_unmodelled === true,
        source: sourceLine(d.loot, a.item_id, a.source_kind, FACTION, a.source),
      })),
    });
  }
  return out;
}

/** Zulmara's worn gear: each slot the pick, or one of the band's alternatives (its delta is the
 *  pick's measured gain over it), Trinket 2 empty. */
export const WORN_CHOICE = {
  head: 2, back: 1, chest: 2, wrist: 2, legs: 2, feet: 2, finger2: 1, ranged: 2,
};
export const WORN_EMPTY = ["trinket2"];

export function charRows(d, rows) {
  const out = new Map();
  for (const [slot, r] of rows) {
    if (r.empty) {
      out.set(slot, { slot, twoHand: r.empty === "two_hand_equipped" });
      continue;
    }
    if (WORN_EMPTY.includes(slot)) {
      out.set(slot, { slot, empty: true, pick: r });
      continue;
    }
    const k = WORN_CHOICE[slot];
    if (k === undefined) {
      out.set(slot, { slot, worn: r, pick: r });
      continue;
    }
    const alt = r.alternatives[k];
    // A pick whose effect the sim cannot model has no trustworthy gain: it says so instead.
    const gainShown = r.unmodelled ? "Gain not modelled yet" : signedDPS(-alt.delta);
    out.set(slot, { slot, worn: alt, pick: r, gainValue: -alt.delta, gain: signedDPS(-alt.delta), gainShown });
  }
  return out;
}
