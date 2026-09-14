// web/src/lib/planner/store.svelte.ts
// The single source of truth for the planner island. `.svelte.ts` so Svelte 5 runes compile
// here; the module has no DOM dependency, so it is unit-tested like any other module.
//
// Every mutation goes through rules.ts. A refused move sets `refusal` and changes nothing
// else, which is exactly what the UI shows inline.
import {
  activeSetBonuses,
  equippedItems,
  levelReached,
  pointsPerTree,
  ranksByTalent,
  splitLabel as deriveSplitLabel,
  statTotals as deriveStatTotals,
  type ActiveSet,
} from './derive';
import {
  canAddPoint,
  canEquip,
  canRemovePoint,
  comboIsLegal,
  indexItems,
  indexTalents,
  withPoint,
  withoutLastPoint,
  type TalentIndex,
} from './rules';
import type {
  BuildDraft,
  ClassRow,
  Combo,
  Gear,
  Item,
  ItemFile,
  ItemSet,
  RaceRow,
  Slot,
  StatKey,
  TalentFile,
} from './types';

/** Shown when someone tries to edit a build opened from a share link. */
export const READ_ONLY_REASON = 'This build was opened from a share link; fork it to edit';
/** The contract caps a title at 60 characters. */
export const MAX_TITLE_LENGTH = 60;

export interface PlannerInit {
  treeVersion: string;
  classSlug: string;
  raceSlug: string;
  order?: number[];
  gear?: Gear;
  title?: string;
  sourceId?: string | null;
  readOnly?: boolean;
}

export interface ReferenceData {
  classes: ClassRow[];
  races: RaceRow[];
  combos: Combo[];
}

export function createPlannerStore(init: PlannerInit) {
  let classSlug = $state(init.classSlug);
  let raceSlug = $state(init.raceSlug);
  let order = $state<number[]>(init.order ? [...init.order] : []);
  let gear = $state<Gear>({ ...(init.gear ?? {}) });
  let title = $state((init.title ?? '').trim().slice(0, MAX_TITLE_LENGTH));
  let readOnly = $state(init.readOnly ?? false);
  let sourceId = $state<string | null>(init.sourceId ?? null);
  let refusal = $state<string | null>(null);

  let talents = $state<TalentFile | null>(null);
  let itemFile = $state<ItemFile | null>(null);
  let sets = $state<ItemSet[]>([]);
  let classes = $state<ClassRow[]>([]);
  let races = $state<RaceRow[]>([]);
  let combos = $state<Combo[]>([]);

  const talentIndex = $derived<TalentIndex | null>(talents ? indexTalents(talents) : null);
  // A derived lookup, not mutable state: indexItems returns a fresh Map and the whole thing is
  // rebuilt whenever itemFile changes, so reactivity flows through the derivation. Every consumer
  // only reads it -- get(), values(), size -- and nothing ever sets, deletes or clears a key, so
  // SvelteMap's per-key tracking would be machinery for mutations that never happen.
  // eslint-disable-next-line svelte/prefer-svelte-reactivity
  const itemIndex = $derived<Map<number, Item>>(itemFile ? indexItems(itemFile) : new Map());
  const classRow = $derived<ClassRow | null>(classes.find((c) => c.slug === classSlug) ?? null);
  const raceRow = $derived<RaceRow | null>(races.find((r) => r.slug === raceSlug) ?? null);
  const legalRaces = $derived<RaceRow[]>(
    classRow === null ? races : races.filter((race) => comboIsLegal(combos, race.id, classRow.id)),
  );
  const ranks = $derived<Map<number, number>>(ranksByTalent(order));
  const split = $derived<number[]>(talentIndex ? pointsPerTree(talentIndex, order) : []);
  const splitLabel = $derived<string>(talentIndex ? deriveSplitLabel(talentIndex, order) : '');
  const equipped = $derived<Item[]>(equippedItems(itemIndex, gear));
  const statTotals = $derived<Partial<Record<StatKey, number>>>(deriveStatTotals(equipped));
  const activeSets = $derived<ActiveSet[]>(activeSetBonuses(equipped, sets));

  /** True when the caller may change the build; otherwise records the read-only refusal. */
  function editable(): boolean {
    if (!readOnly) return true;
    refusal = READ_ONLY_REASON;
    return false;
  }

  /**
   * Moves `raceSlug` to the first race legal for `classId` if the current one is not (or is
   * unresolvable). Leaves `raceSlug` alone when no legal race exists. Used by data loading
   * (`setReference`) and by class switches (`selectClass`) so the two never drift apart.
   */
  function repairRaceForClass(classId: number): void {
    const currentRaceId = races.find((r) => r.slug === raceSlug)?.id ?? -1;
    if (comboIsLegal(combos, currentRaceId, classId)) return;
    const first = races.find((race) => comboIsLegal(combos, race.id, classId));
    if (first) raceSlug = first.slug;
  }

  return {
    get treeVersion() {
      return init.treeVersion;
    },
    get classSlug() {
      return classSlug;
    },
    get raceSlug() {
      return raceSlug;
    },
    get order() {
      return order;
    },
    get gear() {
      return gear;
    },
    get title() {
      return title;
    },
    get readOnly() {
      return readOnly;
    },
    get sourceId() {
      return sourceId;
    },
    get refusal() {
      return refusal;
    },
    get talents() {
      return talents;
    },
    get talentIndex() {
      return talentIndex;
    },
    get itemIndex() {
      return itemIndex;
    },
    get sets() {
      return sets;
    },
    get classes() {
      return classes;
    },
    get races() {
      return races;
    },
    get combos() {
      return combos;
    },
    get classRow() {
      return classRow;
    },
    get raceRow() {
      return raceRow;
    },
    get legalRaces() {
      return legalRaces;
    },
    get ranks() {
      return ranks;
    },
    get spent() {
      return order.length;
    },
    get level() {
      return levelReached(order);
    },
    get split() {
      return split;
    },
    get splitLabel() {
      return splitLabel;
    },
    get equipped() {
      return equipped;
    },
    get statTotals() {
      return statTotals;
    },
    get activeSets() {
      return activeSets;
    },

    setReference(data: ReferenceData): void {
      classes = data.classes;
      races = data.races;
      combos = data.combos;
      const current = classes.find((c) => c.slug === classSlug);
      if (current) repairRaceForClass(current.id);
    },

    setTalents(file: TalentFile): void {
      talents = file;
    },

    /**
     * Replaces the build with one reconstructed from an FS1 code. Separate from the
     * constructor because an FS1 code names talents by tab position, which needs the class's
     * talent file -- and that arrives one fetch after the store is created.
     */
    applyOrder(nextOrder: number[], nextGear: Gear): void {
      if (!editable()) return;
      order = [...nextOrder];
      gear = { ...nextGear };
      refusal = null;
    },

    setItems(file: ItemFile): void {
      itemFile = file;
    },

    setSets(value: ItemSet[]): void {
      sets = value;
    },

    /** Switching class discards the build: its talent ids belong to the old class. */
    selectClass(slug: string): void {
      if (!editable()) return;
      if (slug === classSlug) return;
      classSlug = slug;
      talents = null;
      itemFile = null;
      order = [];
      gear = {};
      refusal = null;
      const next = classes.find((c) => c.slug === slug);
      if (next) repairRaceForClass(next.id);
    },

    selectRace(slug: string): void {
      if (!editable()) return;
      raceSlug = slug;
    },

    addPoint(talentId: number): void {
      if (!editable()) return;
      if (!talentIndex) return;
      const decision = canAddPoint(talentIndex, order, talentId);
      if (!decision.ok) {
        refusal = decision.reason;
        return;
      }
      order = withPoint(order, talentId);
      refusal = null;
    },

    removePoint(talentId: number): void {
      if (!editable()) return;
      if (!talentIndex) return;
      const decision = canRemovePoint(talentIndex, order, talentId);
      if (!decision.ok) {
        refusal = decision.reason;
        return;
      }
      order = withoutLastPoint(order, talentId);
      refusal = null;
    },

    equip(slot: Slot, itemId: number): void {
      if (!editable()) return;
      const decision = canEquip(itemIndex, gear, slot, itemId);
      if (!decision.ok) {
        refusal = decision.reason;
        return;
      }
      gear = { ...gear, [slot]: itemId };
      refusal = null;
    },

    unequip(slot: Slot): void {
      if (!editable()) return;
      const { [slot]: _dropped, ...rest } = gear;
      gear = rest;
      refusal = null;
    },

    setTitle(text: string): void {
      if (!editable()) return;
      title = text.trim().slice(0, MAX_TITLE_LENGTH);
    },

    clearRefusal(): void {
      refusal = null;
    },

    reset(): void {
      if (!editable()) return;
      order = [];
      gear = {};
      refusal = null;
    },

    /** Makes a shared build editable under a new identity. */
    fork(): void {
      readOnly = false;
      sourceId = null;
      title = '';
      refusal = null;
    },

    toDraft(): BuildDraft {
      if (!classRow) throw new Error(`toDraft: no class found for slug "${classSlug}"`);
      if (!raceRow) throw new Error(`toDraft: no race found for slug "${raceSlug}"`);
      const draft: BuildDraft = {
        class_id: classRow.id,
        race_id: raceRow.id,
        tree_version: init.treeVersion,
        point_order: [...order],
        gear: { ...gear },
      };
      if (title.length > 0) draft.title = title;
      return draft;
    },
  };
}

export type PlannerStore = ReturnType<typeof createPlannerStore>;
