// web/src/lib/sim/buffs.ts
// The panel's sections, and nothing else. The id vocabulary itself is IDS.md's, generated
// by the engine from its own protobuf descriptors and copied into
// src/data/generated/sim-ids.json by scripts/sync-sim-ids.mjs -- so this file holds no
// list of ids to offer, only the table that says which section each engine message and each
// `Consumes` field belongs in, and the few ids the engine accepts but ignores (INERT_BUFFS). A buff the engine adds appears in the panel the next time
// the data is synced, with no diff here.
//
// Graded ids are contract 1.7: `<id>:improved` for every TristateEffect field. They are
// folded into the plain id as a three-way choice rather than shown as a second checkbox,
// because "Battle Shout" and "Battle Shout (improved)" both ticked is not a state the
// engine has.
import generated from '../../data/generated/sim-ids.json';
import { IMPROVED_SUFFIX } from './buff-id';

export { IMPROVED_SUFFIX };

export type BuffGroupId =
  | 'raid-buffs'
  | 'party-buffs'
  | 'player-buffs'
  | 'debuffs'
  | 'flask'
  | 'battle-elixir'
  | 'guardian-elixir'
  | 'food'
  | 'weapon-imbue'
  | 'potion'
  | 'explosive';

/** Panel order, top to bottom: what a raid gives you, then what you bring yourself. */
export const BUFF_GROUPS: readonly BuffGroupId[] = [
  'raid-buffs',
  'party-buffs',
  'player-buffs',
  'debuffs',
  'flask',
  'battle-elixir',
  'guardian-elixir',
  'food',
  'weapon-imbue',
  'potion',
  'explosive',
];

export interface SimIdsFile {
  buffs: { id: string; message: string }[];
  consumables: { id: string; sets: string }[];
  professions: string[];
  /** Contract A7's Stats section. Read by stats.ts, not by the catalogue. */
  stats: string[];
}

export interface BuffRow {
  id: string;
  group: BuffGroupId;
  /** True when IDS.md also publishes `<id>:improved`, so the row is a three-way choice. */
  graded: boolean;
  kind: 'buff' | 'consumable';
}

const BY_MESSAGE: Record<string, BuffGroupId> = {
  RaidBuffs: 'raid-buffs',
  PartyBuffs: 'party-buffs',
  IndividualBuffs: 'player-buffs',
  Debuffs: 'debuffs',
};

/**
 * The `Consumes` field each consumable sets, to the kind the design groups by. The engine
 * has one field per effect and the panel has one section per shelf in the bank, so several
 * fields land in one section -- every stat-boosting elixir is a battle elixir, every
 * defensive or regenerative one a guardian elixir, which is the vanilla client's own split.
 */
const BY_CONSUMES_FIELD: Record<string, BuffGroupId> = {
  flask: 'flask',
  agility_elixir: 'battle-elixir',
  strength_buff: 'battle-elixir',
  attack_power_buff: 'battle-elixir',
  spell_power_buff: 'battle-elixir',
  fire_power_buff: 'battle-elixir',
  frost_power_buff: 'battle-elixir',
  shadow_power_buff: 'battle-elixir',
  holy_power_buff: 'battle-elixir',
  hit_consumable: 'battle-elixir',
  bogling_root: 'battle-elixir',
  dragon_breath_chili: 'battle-elixir',
  armor_elixir: 'guardian-elixir',
  health_elixir: 'guardian-elixir',
  mana_regen_elixir: 'guardian-elixir',
  zanza_buff: 'guardian-elixir',
  food: 'food',
  alcohol: 'food',
  main_hand_imbue: 'weapon-imbue',
  off_hand_imbue: 'weapon-imbue',
  // A stone beside a rogue's poison (IDS.md: a poisoned hand also takes a stone or oil).
  main_hand_stone: 'weapon-imbue',
  off_hand_stone: 'weapon-imbue',
  default_potion: 'potion',
  default_conjured: 'potion',
  filler_explosive: 'explosive',
  sapper_explosive: 'explosive',
};

/** "Consumes.main_hand_imbue" is the field `main_hand_imbue`. */
function consumesField(sets: string): string {
  const dot = sets.lastIndexOf('.');
  return dot === -1 ? sets : sets.slice(dot + 1);
}

export function buildCatalogue(ids: SimIdsFile): BuffRow[] {
  const graded = new Set(
    ids.buffs
      .filter((row) => row.id.endsWith(IMPROVED_SUFFIX))
      .map((row) => row.id.slice(0, -IMPROVED_SUFFIX.length)),
  );

  const buffs: BuffRow[] = ids.buffs
    .filter((row) => !row.id.endsWith(IMPROVED_SUFFIX))
    .map((row) => ({
      id: row.id,
      group: BY_MESSAGE[row.message] ?? 'player-buffs',
      graded: graded.has(row.id),
      kind: 'buff' as const,
    }));

  const consumables: BuffRow[] = ids.consumables.map((row) => ({
    id: row.id,
    // An unmapped field lands with the potions rather than vanishing: a consumable the
    // panel cannot file is still a consumable the engine accepts, and hiding it would make
    // it unreachable except through the request drawer.
    group: BY_CONSUMES_FIELD[consumesField(row.sets)] ?? 'potion',
    graded: false,
    kind: 'consumable' as const,
  }));

  // Deduped by id, last write wins: IDS.md is generated in-repo so a repeat is unlikely,
  // but it is the one input this module parses and trusts without a uniqueness check, and
  // BuffPanel.svelte's `{#each rows as row (row.id)}` throws Svelte 5's own
  // `each_key_duplicate` on a repeat -- a dead settings panel rather than a doubled row
  // (finding 4, final whole-branch review).
  return [...new Map([...buffs, ...consumables].map((row) => [row.id, row])).values()];
}

const file = generated as SimIdsFile;

// Forever has no world buffs: sim/request's vocabulary leaves IndividualBuffs' world-buff
// section out (isWorldBuff), so IDS.md, and this catalogue, never carry one.
export const CATALOGUE: readonly BuffRow[] = buildCatalogue(file);

/**
 * Ids the request accepts but the engine applies nothing for, so the panel does not offer
 * them. Round 8 D3 (engine a56f0d992): Forever made Shadow Weaving, Improved Scorch's Fire
 * Vulnerability, Improved Shadow Bolt's Shadow Vulnerability and Stormstrike the caster's own
 * (aura 270/271, damage taken from the caster), so the raid flags give the raid nothing and
 * each class keeps its own from its talents. Curse of Shadow (casters lane, engine 35fa13a5a,
 * round 9): Forever deleted it from the client (no SpellName or SpellEffect row for 17862 or
 * 17937 in 70707); Curse of the Elements covers every school. Elixir of Greater Firepower (holypower lane, engine 46be83bcd, round 12): Forever made its item, 21546, the Elixir of Holy Power, so the fire id gives nothing; Lesser Arcane Elixir is likewise deprecated in the engine. Winter's Chill (auras lane, engine d1dc542c7,
 * round 9): 12579 is aura 271, the applying mage's own, so the raid flag applies nothing. They stay in the catalogue and the request
 * vocabulary, so a saved sim that named one still runs.
 */
export const INERT_BUFFS: readonly string[] = [
  'curse_of_shadow',
  'elixir_of_greater_firepower',
  'improved_scorch',
  'improved_shadow_bolt',
  'lesser_arcane_elixir',
  'shadow_weaving',
  'stormstrike',
  'winters_chill',
];

/** One group's offered rows, alphabetical by id so the panel's order never depends on IDS.md's. */
export function rowsIn(group: BuffGroupId): BuffRow[] {
  return CATALOGUE.filter((row) => row.group === group && !INERT_BUFFS.includes(row.id)).sort((a, b) =>
    a.id.localeCompare(b.id),
  );
}

export type BuffGrade = 'off' | 'on' | 'improved';

export function gradeOf(selected: readonly string[], id: string): BuffGrade {
  if (selected.includes(`${id}${IMPROVED_SUFFIX}`)) return 'improved';
  return selected.includes(id) ? 'on' : 'off';
}

/**
 * Both forms of an id are removed and at most one is added back, so the list can never
 * carry `battle_shout` and `battle_shout:improved` at once -- a state `sim/request` would
 * resolve to whichever it read last, silently.
 */
export function setGrade(selected: readonly string[], id: string, grade: BuffGrade): string[] {
  const improved = `${id}${IMPROVED_SUFFIX}`;
  const without = selected.filter((entry) => entry !== id && entry !== improved);
  if (grade === 'off') return without;
  return [...without, grade === 'improved' ? improved : id];
}

/** The two id lists a `CharacterSpec` carries, and the shape the panel edits. */
export interface Selection {
  buffs: string[];
  consumables: string[];
}

/**
 * Which of the selection's two lists a row of this kind lives in. The one place the
 * kind-to-list rule exists -- `withGrade`, `selectedIn` and the panel's own grade lookup
 * all read it from here rather than each re-deriving the same branch.
 */
export function listFor(selection: Selection, kind: BuffRow['kind']): string[] {
  return kind === 'consumable' ? selection.consumables : selection.buffs;
}

/**
 * One row's grade, written into whichever of the two lists that row belongs to. The panel
 * spreads the answer over the settings object, so a component never decides which list an
 * id goes in -- the catalogue's `kind` does, and it came from IDS.md.
 */
export function withGrade(
  selection: Selection,
  row: Pick<BuffRow, 'id' | 'kind'>,
  grade: BuffGrade,
): Selection {
  if (row.kind === 'consumable') {
    return { ...selection, consumables: setGrade(listFor(selection, row.kind), row.id, grade) };
  }
  return { ...selection, buffs: setGrade(listFor(selection, row.kind), row.id, grade) };
}

/** The plain ids of one group that are on, for the group heading's count. */
export function selectedIn(selection: Selection, group: BuffGroupId): string[] {
  return rowsIn(group)
    .filter((row) => gradeOf(listFor(selection, row.kind), row.id) !== 'off')
    .map((row) => row.id);
}
