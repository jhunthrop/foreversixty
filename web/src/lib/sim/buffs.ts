// web/src/lib/sim/buffs.ts
// The panel's sections, and nothing else. The id vocabulary itself is IDS.md's, generated
// by the engine from its own protobuf descriptors and copied into
// src/data/generated/sim-ids.json by scripts/sync-sim-ids.mjs -- so this file never holds
// a list of ids, only the table that says which section each engine message and each
// `Consumes` field belongs in. A buff the engine adds appears in the panel the next time
// the data is synced, with no diff here.
//
// Graded ids are contract 1.7: `<id>:improved` for every TristateEffect field. They are
// folded into the plain id as a three-way choice rather than shown as a second checkbox,
// because "Battle Shout" and "Battle Shout (improved)" both ticked is not a state the
// engine has.
import generated from '../../data/generated/sim-ids.json';

export const IMPROVED_SUFFIX = ':improved';

export type BuffGroupId =
  | 'raid-buffs'
  | 'party-buffs'
  | 'player-buffs'
  | 'world-buffs'
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
  'world-buffs',
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
  worldBuffs: string[];
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
  const world = new Set(ids.worldBuffs);
  const graded = new Set(
    ids.buffs
      .filter((row) => row.id.endsWith(IMPROVED_SUFFIX))
      .map((row) => row.id.slice(0, -IMPROVED_SUFFIX.length)),
  );

  const buffs: BuffRow[] = ids.buffs
    .filter((row) => !row.id.endsWith(IMPROVED_SUFFIX))
    .map((row) => ({
      id: row.id,
      group:
        row.message === 'IndividualBuffs' && world.has(row.id)
          ? ('world-buffs' as BuffGroupId)
          : (BY_MESSAGE[row.message] ?? 'player-buffs'),
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

  return [...buffs, ...consumables];
}

/**
 * IDS.md has no world-buff section yet (contract 1.7 adds one). Until it does, these are
 * the `IndividualBuffs` fields that are world buffs, read off IDS.md's own table by hand:
 * the four Dire Maul tribute buffs, the two capital-city buffs, Songflower and Sayge's.
 * `buildCatalogue` prefers the generated list whenever it is non-empty, so this disappears
 * the day the engine publishes the section, without a change to any caller.
 */
export const FALLBACK_WORLD_BUFFS: readonly string[] = [
  'fengus_ferocity',
  'moldars_moxie',
  'rallying_cry_of_the_dragonslayer',
  'sayges_fortune',
  'slipkiks_savvy',
  'songflower_serenade',
  'spirit_of_zandalar',
  'warchiefs_blessing',
];

const file = generated as SimIdsFile;

export const CATALOGUE: readonly BuffRow[] = buildCatalogue({
  ...file,
  worldBuffs: file.worldBuffs.length > 0 ? file.worldBuffs : [...FALLBACK_WORLD_BUFFS],
});

/** One group's rows, alphabetical by id so the panel's order never depends on IDS.md's. */
export function rowsIn(group: BuffGroupId): BuffRow[] {
  return CATALOGUE.filter((row) => row.group === group).sort((a, b) => a.id.localeCompare(b.id));
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
