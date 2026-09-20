// web/src/lib/sim/bulk-store.svelte.ts
// The single source of truth for /sim/gear, /sim/talents, /sim/drops and /sim/weights.
//
// It is deliberately NOT the /sim store with more fields on it. The two share their
// character loading (sources.ts, which both call and neither owns) and nothing else: /sim
// runs one request and shows a damage breakdown, these four build a candidate list and show
// a ranking. Merging them would give every page the other's state to reason about, and
// would make the two web lanes of this spec edit one file.
//
// Two rules the components must not second-guess, both the same as /sim's:
//   * a failed load keeps the character already on screen;
//   * the pool is created on the first run, never on mount -- the wasm is 4 MB and these
//     pages have the site's Lighthouse budget.
import { fetchSpecs, saveSim } from './api';
import { stageProgressLine, type BulkProgress, type BulkRunHandle } from './bulk-run';
import {
  browserCap,
  type BulkMode,
  type BulkRequest,
  type GearSet,
  type Precision,
  type StatWeight,
  type TalentLoadout,
  type WeightsRequest,
} from './bulk-types';
import { MAX_SERVER_POLLS, runServerJob } from './bulk-server-run';
import {
  applyRequestFields,
  buildRequest,
  characterCode,
  previewRequest,
  recount as recountCombinations,
  runBulkAndSettle,
  seededStats,
  validateRequestJson,
  weightsSpecRefusal,
  type BulkRequestDeps,
  type BulkRunDeps,
  type RecountDeps,
} from './bulk-store-request';
import {
  addRow,
  copyAndModify,
  removeRow,
  rowFor,
  toggleRow,
  uiSlotsOf,
  type CandidateRow,
  type Origin,
} from './candidates';
import { needsRace, type SimCharacter } from './character';
import type { CharacterPath } from '../characters';
import { bulkCopy } from './copy';
import type { RequestValidation } from './engine';
import { loadEnchants, loadSuffixes, type EnchantRow, type SuffixRow } from './enchants';
import {
  initialShownKinds,
  rowsFromPicks,
  togglePick,
  toggleShownKind,
  visibleSources as visibleSourcesOf,
} from './drop-picks';
import { loadLoot, sourcesByItem, type LootFile } from './loot';
import { BUILT_IN_PHASES, fetchPhases, type PhaseRow } from './phase';
import { loadItems, loadSets, loadTalents } from '../planner/load';
import { isKnownItem, knownItemIds, loadSimItems } from './sim-items';
import { SLOTS, type Item, type ItemSet, type Slot, type TalentFile } from '../planner/types';
import { defaultSettings, withSpecForPreset, type SimSettings } from './settings';
import { loadSimBuffs, type SimBuffFile } from './sim-buffs';
import {
  fromAddonExport,
  fromLoggedFight,
  fromPlannerBuild,
  fromStoredCharacter,
  type LoadContext,
  type SourceResult,
} from './sources';
import type { SimResult, SourceKind, SpecFidelity } from './types';
import { weightStatsFor } from './weights';
import { createPool, type SimPool } from './worker';

export const TOOLS = ['gear', 'talents', 'drops', 'weights'] as const;
export type SimTool = (typeof TOOLS)[number];

/** The three tools that send a `bulk` block, and which mode each sends. */
export const MODE_OF_TOOL: Record<Exclude<SimTool, 'weights'>, BulkMode> = {
  gear: 'gear',
  talents: 'talents',
  drops: 'drops',
};

export type BulkPhase = 'idle' | 'loading-character' | 'counting' | 'running' | 'done' | 'error';

const DEFAULT_SERVER_POLL_MS = 2000;
const COUNT_DEBOUNCE_MS = 250;

/**
 * Plain `Map`s and `Date`s, not `SvelteMap`/`SvelteDate`: `items` and `sourceIndex` are
 * replaced wholesale on every load and only ever read by key, so per-entry tracking would be
 * reactivity machinery for mutations that never happen. Defined at module scope, the same as
 * `store.svelte.ts`'s own `toItemMap` -- `svelte/prefer-svelte-reactivity` flags a bare `new
 * Map()`/`new Date()` written directly inside a function that also calls `$state`, so these
 * live outside `createBulkStore` rather than as inline expressions within it.
 */
function toItemMap(items: readonly Item[]): Map<number, Item> {
  return new Map(items.map((item) => [item.id, item]));
}

function emptySourceIndex(): Map<number, string[]> {
  return new Map();
}

function now(): Date {
  return new Date();
}

export interface BulkStoreInit {
  tool: SimTool;
  treeVersion: string;
  apiBase?: string;
  /** Injected by tests; production creates one on the first run. */
  pool?: SimPool;
  /** Injected by tests; production reads `navigator.hardwareConcurrency`. */
  hardwareConcurrency?: number;
  source?: SourceKind | '';
  ref?: string;
  serverPollMs?: number;
  /** Injected by tests; production uses `MAX_SERVER_POLLS`. */
  serverPollLimit?: number;
  /** Injected by tests, so the phase gate is not a clock-dependent assertion. */
  now?: () => Date;
}

export function createBulkStore(init: BulkStoreInit) {
  const ctx: LoadContext = { treeVersion: init.treeVersion, apiBase: init.apiBase };
  const clock = init.now ?? now;
  const mode: BulkMode | null = init.tool === 'weights' ? null : MODE_OF_TOOL[init.tool];

  let phase = $state<BulkPhase>('idle');
  let character = $state<SimCharacter | null>(null);
  let settings = $state<SimSettings>(defaultSettings('attack_power'));
  let message = $state<string | null>(null);
  let detail = $state('');
  let premium = $state(false);

  let items = $state<Map<number, Item>>(toItemMap([]));
  let sets = $state<ItemSet[]>([]);
  let talentFile = $state<TalentFile | null>(null);
  let enchants = $state<EnchantRow[]>([]);
  let suffixes = $state<SuffixRow[]>([]);
  let loot = $state<LootFile>({ sources: [] });
  let sourceIndex = $state<Map<number, string[]>>(emptySourceIndex());
  // The engine's embedded item universe (sim-items.ts). `null` means the build ships no
  // simitems.json -- every candidate source below treats that the same as "nothing to
  // filter against" (isKnownItem's own default), not "everything is unknown".
  let knownItems = $state<Set<number> | null>(null);
  let specRows = $state<SpecFidelity[]>([]);
  let simBuffs = $state<SimBuffFile>({ entries: {} });
  // The build-time table until GET /v1/phases answers; see phase.ts's own header.
  let phases = $state<readonly PhaseRow[]>(BUILT_IN_PHASES);

  let rows = $state<CandidateRow[]>([]);
  let locked = $state<string[]>([]);
  let loadouts = $state<TalentLoadout[]>([]);
  let namedSets = $state<GearSet[]>([]);
  let precision = $state<Precision>('fast');
  let cap = $state(browserCap(init.hardwareConcurrency ?? globalThis.navigator?.hardwareConcurrency));

  let consumableIds = $state<string[]>([]);
  let combinations = $state<number | null>(null);
  let capNotice = $state<{ cap: number; combinations: number } | null>(null);
  // Set when the count is past the premium lane's own cap too, so the page does not offer
  // a server run the API would refuse at submit (contract 10.1 A2, 5,000).
  let serverCapNotice = $state<{ cap: number; combinations: number } | null>(null);
  let progress = $state<BulkProgress | null>(null);
  let result = $state<SimResult | null>(null);

  // Droptimizer state. `pickedBosses` is keyed `<source id>|<boss id or "">` (drop-picks.ts),
  // so a raid card ("every boss") and one boss of it are two different picks the player can
  // hold at once.
  let pickedBosses = $state<string[]>([]);
  /**
   * `pickedBosses` frozen at the last `run()`/`runOnServer()`/`runRequest()` -- Droptimizer.
   * svelte's `pickedWithNothingTried` argument, deliberately not the live `pickedBosses`:
   * a source ticked AFTER a run must not read as "untried" for a result it was never part
   * of (final whole-branch review, Important 3's deferred minor).
   */
  let submittedDropPicks = $state<string[]>([]);
  let showUpcoming = $state(false);
  let shownKinds = $state<string[]>([]);

  // Weights state.
  let stats = $state<string[]>([]);

  let pool: SimPool | null = init.pool ?? null;
  let handle: BulkRunHandle | null = null;
  let stopRequested = false;
  let serverRunning = $state(false);
  let serverGeneration = 0;
  let countTimer: ReturnType<typeof setTimeout> | null = null;

  function poolOnce(): SimPool {
    pool ??= createPool({});
    return pool;
  }

  /** Every source funnels through here, so the "a bad paste keeps the page" rule is once. */
  async function adopt(load: Promise<SourceResult>): Promise<void> {
    serverGeneration += 1;
    serverRunning = false;
    phase = 'loading-character';
    message = null;
    const outcome = await load;
    if (!outcome.ok) {
      message = outcome.message;
      phase = 'idle';
      return;
    }
    character = outcome.character;
    settings = withSpecForPreset(settings, outcome.character.spec);
    result = null;
    progress = null;
    capNotice = null;
    serverCapNotice = null;
    combinations = null;
    submittedDropPicks = [];
    await loadDataFor(outcome.character);
    rows = seedRows(outcome.character);
    // Design 3.4: talent compare is Top Gear with the gear locked. Locking every slot here
    // rather than hiding the grid means the invariant holds for the request too -- the
    // contract validates "no candidate on a locked slot", and a page that only hid the UI
    // could still send one through the request drawer.
    if (init.tool === 'talents') locked = [...SLOTS];
    // `drops` mode's candidates are the picked sources alone (validateBulk rule 3) -- a
    // fresh load has no picks yet, so this replaces the seeded equipped/bag/bank rows with
    // the empty list rather than leaving them there for a mode that must refuse them.
    if (init.tool === 'drops') rows = rowsFromPicks(pickedBosses, loot, items, knownItems);
    stats = seededStats(init.tool, character, specRows, stats);
    // Set only now, not before `loadDataFor`/`seedRows` above: setting it earlier left a
    // window where the page looked settled (`phase === 'idle'`) while `items` was still
    // empty and `talentFile` still null (fix round 1, Minor).
    phase = 'idle';
  }

  /**
   * Everything a tool page needs about the build, in one pass. Each file is optional in its
   * own way and each failure degrades that one feature rather than the page: no item file
   * means slot names without icons, no loot file means no source picker, no enchant file
   * means no enchant column.
   */
  async function loadDataFor(next: SimCharacter): Promise<void> {
    const [itemFile, setFile, talents, enchantRows, suffixRows, lootFile, buffFile, phaseRows, simItemFile] =
      await Promise.all([
        loadItems(next.tree_version, next.class_slug).catch(() => null),
        loadSets(next.tree_version).catch(() => []),
        loadTalents(next.tree_version, next.class_slug).catch(() => null),
        loadEnchants(next.tree_version).catch(() => []),
        loadSuffixes(next.tree_version).catch(() => []),
        loadLoot(next.tree_version).catch(() => ({ sources: [] }) as LootFile),
        loadSimBuffs(next.tree_version).catch(() => ({ entries: {} }) as SimBuffFile),
        // Never rejects: phase.ts falls back to the build-time table rather than leaving
        // the gate unknown.
        fetchPhases(init.apiBase),
        // A failed fetch degrades to "nothing to filter against" -- the same fallback
        // knownItemIds(null) already gives a build that ships no simitems.json at all.
        loadSimItems(next.tree_version).catch(() => null),
      ]);
    items = toItemMap(itemFile?.items ?? []);
    sets = setFile;
    talentFile = talents;
    enchants = enchantRows;
    suffixes = suffixRows;
    loot = lootFile;
    simBuffs = buffFile;
    phases = phaseRows;
    sourceIndex = sourcesByItem(lootFile);
    shownKinds = initialShownKinds(lootFile);
    knownItems = knownItemIds(simItemFile);
  }

  /**
   * The grid opens with what the character is wearing, plus whatever the addon export's
   * bags and bank carried (part A's FS1 v2 decoder). Nothing is ticked: a page that
   * pre-ticks is a page that runs something the player did not choose. Contract 10.5 puts
   * per-slot enchant and suffix on `gear_slots`, so an equipped row carries the enchant the
   * player actually has on. `gear_slots`/`bags`/`bank` are required, non-`undefined` arrays
   * on every source, so this reads them directly rather than falling back to `gear`.
   */
  function seedRows(next: SimCharacter): CandidateRow[] {
    const seeded: CandidateRow[] = [];
    const add = (itemId: number, origin: Origin, enchant = 0, suffix = 0): void => {
      const item = items.get(itemId);
      if (item === undefined) return;
      const known = isKnownItem(itemId, knownItems);
      for (const slot of uiSlotsOf(item)) {
        seeded.push({ ...rowFor(item, slot, origin, '', known), enchant, suffix });
      }
    };
    for (const slot of next.gear_slots) add(slot.item_id, 'equipped', slot.enchant ?? 0, slot.suffix ?? 0);
    // FS1Item carries camelCase `itemId`/`enchant`/`suffix` (planner/fs1.ts) -- the addon
    // export's own shape, not the wire's `GearSlot`.
    for (const entry of next.bags) add(entry.itemId, 'bag', entry.enchant ?? 0, entry.suffix ?? 0);
    for (const entry of next.bank) add(entry.itemId, 'bank', entry.enchant ?? 0, entry.suffix ?? 0);
    return seeded;
  }

  /** The live count, debounced: it fires on every checkbox tick. */
  function scheduleCount(): void {
    if (countTimer !== null) clearTimeout(countTimer);
    countTimer = setTimeout(() => void recount(), COUNT_DEBOUNCE_MS);
  }

  /**
   * The one stat every weight is normalised against. It always leads `stats`
   * (StatWeights.svelte keeps it there), and it is defined here once so the page's
   * "Reference: …" line and `WeightsSpec.Reference` on the wire cannot drift apart.
   */
  const referenceStat = (): string => stats[0] ?? '';

  /**
   * Everything `bulk-store-request.ts`'s functions need from this closure, built once:
   * `$state` cannot cross a module boundary, so every field is a getter or a setter here
   * rather than the module reading/writing `$state` itself (Task 15 -- this store's own
   * request-building surface outgrew the 800-line cap once the drawer's four methods were
   * added, the same seam `store.svelte.ts`/`store-request.ts` already use).
   */
  const requestDeps: BulkRequestDeps = {
    tool: init.tool,
    mode,
    treeVersion: init.treeVersion,
    getCharacter: () => character,
    getTalentFile: () => talentFile,
    getSettings: () => settings,
    getRows: () => rows,
    getLocked: () => locked,
    getLoadouts: () => loadouts,
    getNamedSets: () => namedSets,
    getPrecision: () => precision,
    getCap: () => cap,
    getConsumableIds: () => consumableIds,
    getStats: () => stats,
    getReferenceStat: () => referenceStat(),
    setPrecision: (value) => (precision = value),
    setCap: (value) => (cap = value),
    setLocked: (value) => (locked = value),
    setLoadouts: (value) => (loadouts = value),
    setNamedSets: (value) => (namedSets = value),
    setStats: (value) => (stats = value),
    scheduleCount,
    poolOnce,
  };

  /** Everything `runBulkAndSettle` needs from this closure -- same reasoning as `requestDeps`. */
  const runDeps: BulkRunDeps = {
    tool: init.tool,
    poolOnce,
    getStopRequested: () => stopRequested,
    setStopRequested: (value) => (stopRequested = value),
    setPhase: (value) => (phase = value),
    setProgress: (value) => (progress = value),
    setMessage: (value) => (message = value),
    setDetail: (value) => (detail = value),
    getResult: () => result,
    setResult: (value) => (result = value),
    setHandle: (value) => (handle = value),
    setCapNotice: (value) => (capNotice = value),
    setCombinations: (value) => (combinations = value),
    bumpServerGeneration: () => (serverGeneration += 1),
  };

  /**
   * Everything `recount` (bulk-store-request.ts) needs beyond `requestDeps` -- the phase
   * gate and the fields a count can change. Built once, the same seam as `requestDeps`/
   * `runDeps` above (fix round 2: `recount`'s own body moved out to keep this file under
   * the 800-line cap; this object is the only thing that replaced it here).
   */
  const recountDeps: RecountDeps = {
    ...requestDeps,
    getPhase: () => phase,
    setPhase: (value) => (phase = value),
    setMessage: (value) => (message = value),
    setDetail: (value) => (detail = value),
    setCombinations: (value) => (combinations = value),
    setCapNotice: (value) => (capNotice = value),
    setServerCapNotice: (value) => (serverCapNotice = value),
  };

  /** The live combination count. Thin wiring only -- see `bulk-store-request.ts`'s own
   *  `recount` for the actual decision, moved there in fix round 2. */
  function recount(): Promise<void> {
    return recountCombinations(recountDeps);
  }

  return {
    get tool() {
      return init.tool;
    },
    get phase() {
      return phase;
    },
    get character() {
      return character;
    },
    get needsRace() {
      return character !== null && needsRace(character);
    },
    get settings() {
      return settings;
    },
    get items() {
      return items;
    },
    get sets() {
      return sets;
    },
    get enchants() {
      return enchants;
    },
    get suffixes() {
      return suffixes;
    },
    get loot() {
      return loot;
    },
    get sourceIndex() {
      return sourceIndex;
    },
    get knownItems() {
      return knownItems;
    },
    get specRows() {
      return specRows;
    },
    get rows() {
      return rows;
    },
    get locked() {
      return locked;
    },
    get loadouts() {
      return loadouts;
    },
    get namedSets() {
      return namedSets;
    },
    get precision() {
      return precision;
    },
    get cap() {
      return cap;
    },
    get combinations() {
      return combinations;
    },
    get capNotice() {
      return capNotice;
    },
    /** Non-null only when even the premium lane's 5,000 would be exceeded. */
    get serverCapNotice() {
      return serverCapNotice;
    },
    get consumableIds() {
      return consumableIds;
    },
    /** The buff/consumable name table, for the candidate list's labels. */
    get simBuffs() {
      return simBuffs;
    },
    /** The phase table the source picker gates on: live where the API answered. */
    get phases() {
      return phases;
    },
    get progress() {
      return progress;
    },
    /** "" when nothing is running; the design's own sentence otherwise. */
    get progressLine() {
      return progress === null ? '' : stageProgressLine(progress);
    },
    get result() {
      return result;
    },
    get combos() {
      return result?.combos ?? [];
    },
    get weights(): StatWeight[] {
      return result?.weights ?? [];
    },
    get message() {
      return message;
    },
    get detail() {
      return detail;
    },
    get premium() {
      return premium;
    },
    get serverRunning() {
      return serverRunning;
    },
    get showUpcoming() {
      return showUpcoming;
    },
    get shownKinds() {
      return shownKinds;
    },
    get pickedBosses() {
      return pickedBosses;
    },
    get submittedDropPicks() {
      return submittedDropPicks;
    },
    get stats() {
      return stats;
    },
    get referenceStat() {
      return referenceStat();
    },
    /**
     * The current character's spec's own `weight_stats` from `GET /v1/specs`, or
     * `undefined` when the list has no row, or no column, for it -- `weights.ts`'s
     * `pickableStatsFor`/`WEIGHTS_STATS_FROM_ENGINE` both key off this same optionality, so
     * the picker and its explainer can never disagree about whether the list is curated.
     */
    get weightStats() {
      return weightStatsFor(character?.spec ?? '', specRows);
    },
    /** The sources the picker draws: kind ticked on, and released unless asked otherwise. */
    get visibleSources() {
      return visibleSourcesOf(loot, shownKinds, phases, showUpcoming, clock());
    },

    loadAddon: (code: string) => adopt(fromAddonExport(code, ctx)),
    loadBuild: (id: string) => adopt(fromPlannerBuild(id, ctx)),
    loadFight: (ref: string) => adopt(fromLoggedFight(ref, ctx)),
    loadStored: (path: CharacterPath) => adopt(fromStoredCharacter(path, ctx)),

    async loadSpecs(): Promise<void> {
      try {
        specRows = await fetchSpecs(init.apiBase);
      } catch {
        specRows = [];
      }
      stats = seededStats(init.tool, character, specRows, stats);
    },

    /** Exactly what a run would send, for part A's Advanced drawer (design 8). */
    get requestPreview() {
      return previewRequest(requestDeps);
    },
    /** A fresh FS1 code for the loaded character, or null -- the tab strip's own fallback
     *  for a `source.ref`-less character (bulk-store-request.ts's own comment). */
    get characterCode() {
      return characterCode(requestDeps);
    },
    /**
     * A request edited in the drawer, adopted whole. Only the two blocks this store owns
     * are read back -- the encounter and the buffs belong to `settings`, which part A's own
     * panel owns, and writing them from here would fight it.
     */
    applyRequest(next: unknown): void {
      applyRequestFields(requestDeps, next);
    },
    /** `api.SimRequest.Validate`, inside the wasm -- the drawer's own inline JSON check. */
    validateRequest(json: string): Promise<RequestValidation> {
      return validateRequestJson(requestDeps, json);
    },

    setPremium(value: boolean): void {
      premium = value;
    },
    setMessage(text: string | null): void {
      message = text;
    },
    setSettings(next: SimSettings): void {
      settings = next;
      scheduleCount();
    },
    setRace(slug: string): void {
      if (character === null) return;
      character = { ...character, race_slug: slug };
      result = null;
      message = null;
      phase = 'idle';
    },
    setPrecision(value: Precision): void {
      precision = value;
      scheduleCount();
    },
    setCap(value: number): void {
      cap = value;
      scheduleCount();
    },
    setStats(next: string[]): void {
      stats = [...next];
    },

    toggleRow(key: string): void {
      rows = toggleRow(rows, key);
      scheduleCount();
    },
    removeRow(key: string): void {
      rows = removeRow(rows, key);
      scheduleCount();
    },
    copyAndModify(key: string, patch: { enchant?: number; suffix?: number }): void {
      rows = copyAndModify(rows, key, patch);
      scheduleCount();
    },
    /**
     * From the item search and from a Droptimizer pin: added, ticked, and counted.
     * `sourceName` carries a pin's boss/source name through to the row (contract 10.1 A6),
     * so `addRow`'s own provenance merge can prefer it over a bare, unnamed origin already
     * ticked from bags or equipped.
     *
     * Returns whether the item was actually found and added. `ToolsView.svelte`'s own pin
     * effect gates on `phase === 'idle'`, which is only true once `items` is populated, so a
     * pin arriving through the normal flow always finds its item here -- but a stale or
     * cross-class item id in the URL is still a real, reachable case (fix round 1, Finding
     * 2's "say so" ask). A miss now sets `message` to `bulkCopy.itemNotAdded(itemId)` --
     * player-visible copy, not only the internal `detail` fix round 1 left this as (fix
     * round 2: a `detail`-only write is invisible, since `BulkRunBar` only renders `detail`
     * inside `{#if message !== null}`). Set here, once, rather than at each of the two call
     * sites (`ToolsView`'s pin effect, `TopGear`'s item-search "Add"): one place owns what
     * "the item was not found" means to the player, and neither caller has to remember to
     * check the boolean for the message to appear -- though both still can, for their own
     * reasons, since the boolean is still returned.
     */
    addSearchItem(itemId: number, origin: Origin = 'search', sourceName = ''): boolean {
      const item = items.get(itemId);
      if (item === undefined) {
        message = bulkCopy.itemNotAdded(itemId);
        return false;
      }
      // A row the engine does not know about is added unticked, not ticked-but-excluded:
      // `toCandidates` would drop it from the outgoing request either way (candidates.ts's
      // `known` guard), but a checked, disabled checkbox reads as "this is included and
      // you cannot change that" -- exactly backwards for a candidate that can never be
      // included. `CandidateRows.svelte` disables the checkbox regardless of `checked`, so
      // this only changes what it shows, not what it would do if it could be ticked.
      const known = isKnownItem(itemId, knownItems);
      for (const slot of uiSlotsOf(item)) {
        rows = addRow(rows, { ...rowFor(item, slot, origin, sourceName, known), checked: known });
      }
      scheduleCount();
      return true;
    },
    toggleLock(slot: Slot): void {
      locked = locked.includes(slot) ? locked.filter((entry) => entry !== slot) : [...locked, slot];
      scheduleCount();
    },
    addLoadout(loadout: TalentLoadout): void {
      if (loadouts.some((entry) => entry.name === loadout.name)) return;
      loadouts = [...loadouts, loadout];
      scheduleCount();
    },
    removeLoadout(name: string): void {
      loadouts = loadouts.filter((entry) => entry.name !== name);
      scheduleCount();
    },
    /**
     * A named set is one whole-outfit candidate (design 3.1.7), with no per-slot row in the
     * grid to disable -- so unlike a search, bag, bank or Droptimizer candidate, an unknown
     * item here cannot be shown disabled and left in. It is dropped from the set instead:
     * that slot simply is not substituted for this candidate (the same as a slot the export
     * never named), rather than refusing the whole set over one item the engine does not
     * carry, or sending a `simCount` refusal for the whole request.
     */
    addNamedSet(set: GearSet): void {
      const known = { ...set, gear: set.gear.filter((slot) => isKnownItem(slot.item_id, knownItems)) };
      namedSets = [...namedSets.filter((entry) => entry.name !== set.name), known];
      scheduleCount();
    },
    removeNamedSet(name: string): void {
      namedSets = namedSets.filter((entry) => entry.name !== name);
      scheduleCount();
    },
    toggleConsumable(id: string): void {
      consumableIds = consumableIds.includes(id)
        ? consumableIds.filter((entry) => entry !== id)
        : [...consumableIds, id];
      scheduleCount();
    },
    toggleKind(kind: string): void {
      shownKinds = toggleShownKind(shownKinds, kind);
    },
    setShowUpcoming(value: boolean): void {
      showUpcoming = value;
    },
    /** A boss, or a whole source when `bossId` is empty. */
    toggleSource(sourceId: string, bossId = ''): void {
      pickedBosses = togglePick(pickedBosses, sourceId, bossId);
      rows = rowsFromPicks(pickedBosses, loot, items, knownItems);
      scheduleCount();
    },

    recount,

    adoptResult(next: SimResult): void {
      result = next;
      phase = 'done';
    },

    async run(): Promise<void> {
      // A premium poll owns `result`/`phase` while it runs; refuse rather than race it,
      // the same guard `runOnServer()` puts on itself (fix round 1, Minor).
      if (serverRunning) return;
      // Unenforced invariant (fix round 3): `runBulk` never calls `pool.validate` itself.
      // `buildRequest` only runs `validateBulk`'s rules, not `encounter` bounds like
      // `targets` (1-10) -- those stay in range only because `SettingsBar`'s `withTargets`
      // clamps every write; `setSettings()` validates nothing. A control that skips that
      // clamp reaches here unvalidated and fails as a generic `bulkFailed`, not
      // `countCombinations`'s earlier `BulkValidationError`.
      const outcome = buildRequest(requestDeps);
      if ('error' in outcome) {
        message = outcome.error;
        phase = 'idle';
        return;
      }
      submittedDropPicks = pickedBosses;
      await runBulkAndSettle(runDeps, outcome.request);
    },

    /**
     * The edited request from part A's drawer, run exactly as written (design 8's escape
     * hatch) -- bypasses the ticked-candidates reconstruction `run()` does entirely, the
     * same way `applyRequest` bypasses it for Apply. The drawer's own `checkRequest` has
     * already run the engine's own Validate by the time it calls this, so this does not
     * re-validate -- except for the one thing `simValidate` cannot see at all (no spec-role
     * concept, on either the real engine or the fake): sub-item 5's honest refusal. Fix
     * round 1 gated only `buildRequest`, the ticked-candidates path; this path reached the
     * pool for an unsupported spec untouched, with `previewRequest` even seeding the
     * drawer's textarea with a full, valid-looking request for it. `weightsSpecRefusal` is
     * the one function both paths now call, so a third path cannot reopen this gap again.
     */
    async runRequest(request: unknown): Promise<void> {
      if (serverRunning) return;
      const parsed = request as Partial<BulkRequest & WeightsRequest>;
      if (parsed.bulk === undefined && parsed.weights === undefined) return;
      if (parsed.weights !== undefined && parsed.spec !== undefined) {
        const refusal = weightsSpecRefusal(parsed.spec);
        if (refusal !== null) {
          message = refusal;
          phase = 'idle';
          return;
        }
      }
      // Best-effort: Apply syncs the drawer to `pickedBosses` before this runs it.
      submittedDropPicks = pickedBosses;
      await runBulkAndSettle(runDeps, parsed as BulkRequest | WeightsRequest);
    },

    stop(): void {
      stopRequested = true;
      handle?.cancel();
      // One Stop button covers both lanes (fix round 1, Important 1): bumping the
      // generation unwinds the premium poll on its next `await`, and `serverRunning` flips
      // now rather than waiting for that poll to notice.
      serverGeneration += 1;
      serverRunning = false;
    },

    /**
     * The same envelope, on the premium lane. The dispatch-and-poll mechanics live in
     * `bulk-server-run.ts` (fix round 1's split, for the 800-line cap); this is only the
     * wiring from that module's updates onto this store's own `$state`.
     */
    async runOnServer(): Promise<void> {
      if (serverRunning) return;
      // Task 8, sub-item 2: the server lane keeps its own (unguarded) default -- only a
      // weights request reads `lane` at all, and only to pick its iteration base.
      const outcome = buildRequest(requestDeps, 'server');
      if ('error' in outcome) {
        message = outcome.error;
        return;
      }
      submittedDropPicks = pickedBosses;
      const request = outcome.request;
      serverRunning = true;
      const generation = ++serverGeneration;
      const current = (): boolean => generation === serverGeneration;
      message = null;
      detail = '';
      try {
        await runServerJob(
          request,
          {
            apiBase: init.apiBase,
            pollMs: init.serverPollMs ?? DEFAULT_SERVER_POLL_MS,
            pollLimit: init.serverPollLimit ?? MAX_SERVER_POLLS,
            isCurrent: current,
            hasPriorResult: () => result !== null,
          },
          (update) => {
            if (update.kind === 'progress') progress = update.progress;
            else if (update.kind === 'message') message = update.message;
            else if (update.kind === 'failed') {
              message = update.message;
              phase = update.phase;
            } else {
              result = update.result;
              progress = null;
              phase = 'done';
            }
          },
        );
      } finally {
        if (current()) serverRunning = false;
      }
    },

    /**
     * Saves a finished browser-lane result via `POST /v1/sims` (contract 10.6: that route
     * "remains the browser-result save" for every kind) -- mirrors `store.svelte.ts`'s own
     * `save()`: same signature, null on failure without touching `message` (a save failure
     * is the save form's own concern). No `result.lane` check: a server-lane result already
     * has its own `sim_id`, so wiring this to only a browser-lane result is Task 16's job.
     */
    async save(title?: string): Promise<string | null> {
      if (result === null) return null;
      try {
        return await saveSim(result, init.apiBase, title);
      } catch {
        return null;
      }
    },

    dispose(): void {
      handle?.cancel();
      if (countTimer !== null) clearTimeout(countTimer);
      serverGeneration += 1;
      serverRunning = false;
      pool?.terminate();
      pool = null;
    },
  };
}

export type BulkStore = ReturnType<typeof createBulkStore>;
