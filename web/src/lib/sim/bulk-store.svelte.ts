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
import {
  BulkCapError,
  BulkRunError,
  countCombinations,
  runBulk,
  runWeightsRun,
  stageProgressLine,
  type BulkProgress,
  type BulkRunHandle,
} from './bulk-run';
import {
  SERVER_CAP,
  browserCap,
  finalIterations,
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
  addRow,
  buildBulkSpec,
  copyAndModify,
  removeRow,
  rowFor,
  toggleRow,
  uiSlotsOf,
  validateBulk,
  type CandidateRow,
  type Origin,
} from './candidates';
import { needsRace, toCharacterSpec, type SimCharacter } from './character';
import type { CharacterPath } from '../characters';
import { bulkCopy, simCopy } from './copy';
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
import { PRECISION_ITERATIONS } from './precision';
import { indexTalents } from '../planner/rules';
import type { Item, ItemSet, Slot, TalentFile } from '../planner/types';
import { defaultSettings, type SimSettings } from './settings';
import { loadSimBuffs, type SimBuffFile } from './sim-buffs';
import {
  fromAddonExport,
  fromLoggedFight,
  fromPlannerBuild,
  fromStoredCharacter,
  type LoadContext,
  type SourceResult,
} from './sources';
import type { CharacterSpec, SimResult, SourceKind, SpecFidelity } from './types';
import { ENGINE_VERSION } from './version';
import { defaultStatsFor, referenceFor } from './weights';
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
  let settings = $state<SimSettings>(defaultSettings());
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
    result = null;
    progress = null;
    capNotice = null;
    serverCapNotice = null;
    combinations = null;
    await loadDataFor(outcome.character);
    rows = seedRows(outcome.character);
    // `drops` mode's candidates are the picked sources alone (validateBulk rule 3) -- a
    // fresh load has no picks yet, so this replaces the seeded equipped/bag/bank rows with
    // the empty list rather than leaving them there for a mode that must refuse them.
    if (init.tool === 'drops') rows = rowsFromPicks(pickedBosses, loot, items);
    seedStatsIfNeeded();
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
    const [itemFile, setFile, talents, enchantRows, suffixRows, lootFile, buffFile, phaseRows] =
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
      for (const slot of uiSlotsOf(item)) {
        seeded.push({ ...rowFor(item, slot, origin), enchant, suffix });
      }
    };
    for (const slot of next.gear_slots) add(slot.item_id, 'equipped', slot.enchant ?? 0, slot.suffix ?? 0);
    // FS1Item carries camelCase `itemId`/`enchant`/`suffix` (planner/fs1.ts) -- the addon
    // export's own shape, not the wire's `GearSlot`.
    for (const entry of next.bags) add(entry.itemId, 'bag', entry.enchant ?? 0, entry.suffix ?? 0);
    for (const entry of next.bank) add(entry.itemId, 'bank', entry.enchant ?? 0, entry.suffix ?? 0);
    return seeded;
  }

  /** The character as the engine wants it, or null while a character or its talents are missing. */
  function characterSpecOrNull(): CharacterSpec | null {
    if (character === null || talentFile === null) return null;
    return toCharacterSpec(character, indexTalents(talentFile), settings.buffs, settings.consumables);
  }

  function currentSpec(): BulkRequest['bulk'] | null {
    if (mode === null) return null;
    return buildBulkSpec({
      mode,
      rows,
      locked,
      loadouts,
      sets: namedSets,
      precision,
      cap,
      // "Try each of these" is one alternative list per ticked consumable, not every
      // subset of them (contract 10.1 A5).
      consumables: consumableIds.map((id) => [id]),
    });
  }

  /**
   * The fields a bulk and a weights request share, or null while a character or its talent
   * file is missing. One function so `run()`'s envelope and `recount()`'s can never drift
   * apart (fix round 1, Important 4) -- only `iterations` and the trailing block differ.
   */
  function envelope(iterations: number): Omit<BulkRequest, 'bulk'> | null {
    const spec = characterSpecOrNull();
    if (spec === null || character === null) return null;
    return {
      engine_version: ENGINE_VERSION,
      spec: character.spec,
      source: character.source,
      character: spec,
      encounter: settings.encounter,
      iterations,
      random_seed: 0,
    };
  }

  /**
   * The weights picker's seed, run whenever a character or the spec list newly become
   * available -- `loadSpecs()` may resolve before or after `loadAddon()` (Task 11's own
   * `onMount` order), and `stats` must not stay empty forever either way (fix round 1,
   * Important 2). Guarded on `stats.length === 0`, so a player's own edits are never lost.
   */
  function seedStatsIfNeeded(): void {
    if (init.tool !== 'weights' || stats.length > 0 || character === null) return;
    stats = defaultStatsFor(character.spec, referenceFor(character.spec, specRows));
  }

  function baseRequest(): BulkRequest | WeightsRequest | null {
    if (character === null) {
      message = bulkCopy.needCharacter;
      return null;
    }
    if (init.tool === 'weights') {
      // Contract 10.8: `WeightsSpec.Reference` is required. An empty `stats` list has
      // nothing to send as one, and running would otherwise fail only after the engine
      // rejects the request (fix round 1, Important 2).
      if (stats.length === 0) {
        message = bulkCopy.weightsNeedStats;
        return null;
      }
      const base = envelope(PRECISION_ITERATIONS.normal);
      if (base === null) {
        message = simCopy.failed;
        return null;
      }
      return { ...base, weights: { stats: [...stats], reference: stats[0] } };
    }
    // Contract 10.1 A3: a bulk request's iterations ARE its precision's final stage, and
    // `Validate` refuses anything else.
    const base = envelope(finalIterations(precision));
    if (base === null) {
      message = simCopy.failed;
      return null;
    }
    const bulk = currentSpec();
    if (bulk === null) return null;
    const refusal = validateBulk(bulk);
    if (refusal !== null) {
      message = refusal;
      return null;
    }
    return { ...base, bulk };
  }

  /** The live count, debounced: it fires on every checkbox tick. */
  function scheduleCount(): void {
    if (countTimer !== null) clearTimeout(countTimer);
    countTimer = setTimeout(() => void recount(), COUNT_DEBOUNCE_MS);
  }

  async function recount(): Promise<void> {
    if (mode === null || character === null) return;
    const bulk = currentSpec();
    if (bulk === null || validateBulk(bulk) !== null) {
      combinations = null;
      capNotice = null;
      // The invalid-spec exit used to leave a stale server-cap notice on screen after the
      // player unticked everything past 5,000 (fix round 1, Important 3) -- every exit now
      // clears all three of the same fields.
      serverCapNotice = null;
      return;
    }
    const base = envelope(finalIterations(precision));
    if (base === null) return;
    const request: BulkRequest = { ...base, bulk };
    phase = 'counting';
    try {
      combinations = await countCombinations(poolOnce(), request);
      capNotice = null;
      // Derived directly from the count on the success path too, not only through a
      // `BulkCapError` (fix round 1, Minor): `setCap()` is public, so a browser cap raised
      // past 5,000 must not let a 6,000-combination count succeed without this notice.
      serverCapNotice = combinations > SERVER_CAP ? { cap: SERVER_CAP, combinations } : null;
    } catch (error) {
      if (error instanceof BulkCapError) {
        combinations = error.combinations;
        capNotice = { cap: error.cap, combinations: error.combinations };
        // Past the premium lane's own 5,000 too (contract 10.1 A2), so the page does not
        // offer a server run the API would refuse at submit with `cap_exceeded`.
        serverCapNotice =
          error.combinations > SERVER_CAP ? { cap: SERVER_CAP, combinations: error.combinations } : null;
      } else {
        // Not a cap refusal: a genuine engine error while merely counting. The count
        // blanks rather than showing a stale number; `detail` carries the reason for a
        // component that wants it -- `run()` raises the same failure, with `message` set,
        // the moment the player actually presses Run.
        combinations = null;
        capNotice = null;
        serverCapNotice = null;
        detail = error instanceof Error ? error.message : '';
      }
    } finally {
      if (phase === 'counting') phase = 'idle';
    }
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
    get stats() {
      return stats;
    },
    get referenceStat() {
      return stats[0] ?? '';
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
      seedStatsIfNeeded();
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
    /** From the item search and from a Droptimizer pin: added, ticked, and counted. */
    addSearchItem(itemId: number, origin: Origin = 'search'): void {
      const item = items.get(itemId);
      if (item === undefined) return;
      for (const slot of uiSlotsOf(item)) {
        rows = addRow(rows, { ...rowFor(item, slot, origin), checked: true });
      }
      scheduleCount();
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
    addNamedSet(set: GearSet): void {
      namedSets = [...namedSets.filter((entry) => entry.name !== set.name), set];
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
      rows = rowsFromPicks(pickedBosses, loot, items);
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
      const request = baseRequest();
      if (request === null) {
        phase = 'idle';
        return;
      }
      // Invalidates a premium poll that might still be resolving from an earlier
      // generation (fix round 1, Minor).
      serverGeneration += 1;
      message = null;
      detail = '';
      stopRequested = false;
      phase = 'running';
      progress = null;

      handle =
        init.tool === 'weights'
          ? runWeightsRun(poolOnce(), request as WeightsRequest, () => {})
          : runBulk(poolOnce(), request as BulkRequest, (next) => {
              if (stopRequested) return;
              progress = next;
            });

      try {
        const finished = await handle.result;
        result = finished;
        if (finished.aborted === true) message = bulkCopy.partial;
        phase = 'done';
      } catch (error) {
        if (error instanceof BulkCapError) {
          capNotice = { cap: error.cap, combinations: error.combinations };
          combinations = error.combinations;
          message = error.message;
          phase = 'idle';
          return;
        }
        const failure = error instanceof BulkRunError ? error : null;
        message = failure?.cancelled === true ? simCopy.stopped : (failure?.message ?? bulkCopy.bulkFailed);
        detail = failure?.detail ?? '';
        phase = failure?.cancelled === true && result !== null ? 'done' : 'error';
      } finally {
        handle = null;
        progress = null;
      }
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
      const request = baseRequest();
      if (request === null) return;
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
