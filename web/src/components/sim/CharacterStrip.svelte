<!-- web/src/components/sim/CharacterStrip.svelte -->
<!-- Who is being simulated. Read-only by design: gear is changed in the planner, which the
     "Open in planner" link goes to, and a second editor here would be a second set of rules
     to keep in step with lib/planner/rules.ts.
     An open section headed "Character" with "Change source": the identity line (crest, name,
     descriptor, item level, source), then the worn set as the shared paperdoll's list form in
     `worn` mode (web layout spec 2026-11-05 §4.2; inputs only, no gains), then the race
     question, the level note and the talents row. An export that was not read, or a saved sim
     with no gear list, draws no doll: the identity stands over the "not read" panel or the line
     that says so. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { classColorVar } from '../../lib/report/format';
  import { button } from '../../lib/ui/button';
  import Crest from '../ui/Crest.svelte';
  import ExportNotRead from '../character/ExportNotRead.svelte';
  import BisTooltipHost from '../BisTooltipHost.svelte';
  import GearDoll from '../character/GearDoll.svelte';
  import Section from '../ui/Section.svelte';
  import { simCharacterHasGear } from '../../lib/addon-export';
  import { gearCopy } from '../../lib/gear/copy';
  import { averageItemLevel } from '../../lib/gear/paperdoll';
  import { slotsFromWorn } from '../../lib/gear/paperdoll-slots';
  import { loadItemNames, loadTalents } from '../../lib/planner/load';
  import { indexTalents, type TalentIndex } from '../../lib/planner/rules';
  import { SLOTS, type Item, type RaceRow, type Slot, type TalentFile } from '../../lib/planner/types';
  import type { ItemsStatus } from '../../lib/sim/store.svelte';
  import type { SimCharacter } from '../../lib/sim/character';
  import { needsRace, plannerHrefFor } from '../../lib/sim/character';
  import { simCopy } from '../../lib/sim/copy';
  import { loadEnchants } from '../../lib/sim/enchants';
  import { enchantNames, withEnchants } from '../../lib/sim/enchant-names';
  import { handoffCopy } from '../../lib/sim/handoff-copy';
  import { sourcePill } from '../../lib/sim/sources';
  import { specLabel } from '../../lib/sim/spec-label';
  import Select from '../ui/Select.svelte';

  /** "high-order-skyborne" -> "High Order Skyborne". */
  function raceWords(slug: string): string {
    return slug
      .split('-')
      .filter((word) => word !== '')
      .map((word) => word.charAt(0).toUpperCase() + word.slice(1))
      .join(' ');
  }

  let {
    character,
    items,
    itemsStatus = 'ready',
    races = [],
    gearKnown = true,
    readonly = false,
    plannerHref,
    talentPoints,
    onchange,
    onrace = () => {},
    plated = true,
  }: {
    character: SimCharacter;
    items: Map<number, Item>;
    /** Whether `items` is the class's file yet: the frame shimmers while it loads. */
    itemsStatus?: ItemsStatus;
    /** The build's races, for the one-time picker. Only the fight source ever needs it. */
    races?: RaceRow[];
    /** False on a saved sim, whose stored request has no gear list. */
    gearKnown?: boolean;
    /** True on /sim/<sim_id> (SavedSim.svelte): hides "Change source", which has nothing to
     *  change on a read-only result -- the button was rendering live and focusable but
     *  wired to a no-op (Task 17's review, MEDIUM). */
    readonly?: boolean;
    /**
     * The "Open in planner" link, precomputed by the caller -- SavedSim.svelte's own fix:
     * a saved sim's `character` carries no `point_order` (genuinely unknowable, the same
     * honest-empty rule `sources.ts` already follows for a combat log) and never will, so
     * this component's own `plannerHrefFor(character, talentIndex)` below -- which needs
     * one to reconstruct a talents string through `toCharacterSpec` -- can only ever encode
     * zeroed talents for it. A saved sim's stored request already carries the true, final
     * talents string with nothing to reconstruct (`plannerHrefForSpec`, character.ts), so
     * SavedSim builds the href itself and hands it straight through here rather than this
     * component trying to re-derive from a `point_order` its caller never had. Absent (the
     * live /sim page, SimView.svelte): unchanged, this component still computes its own
     * from `character`/`talentIndex` below.
     */
    plannerHref?: string;
    /**
     * The point count, precomputed by the caller -- SavedSim.svelte's own fix for the same
     * reason `plannerHref` above is: `character.point_order` is empty for a saved sim (the
     * click order is genuinely unknowable from the stored request), so this component's own
     * `character.point_order.length` reads 0 and the count used to be hidden entirely
     * (2026-09-21 result-page review round 3, newcomer's own finding: the live page shows
     * "51 points" beside Open in planner, the saved page showed nothing). A saved sim's
     * stored request already carries the true, final talents string, and the point count is
     * just the sum of its digits (`talentPointsFromString`, character.ts) -- no order
     * needed. Absent (the live /sim page): unchanged, this component still counts
     * `character.point_order.length`.
     */
    talentPoints?: number;
    onchange: () => void;
    /** The player answering the race question; the store replaces the character. */
    onrace?: (slug: string) => void;
    /** The doll sits on a Plate; false on a page whose one Plate is its result (a saved Top gear run). */
    plated?: boolean;
  } = $props();

  // "Open in planner" (below) needs the talent tree layout to encode the character's
  // points into an FS1 code -- the same file TalentCandidates.svelte:41-44 loads for its
  // own build list. Starts null so the link renders from first paint with the class+race
  // fallback `plannerHrefFor` gives a null index, and upgrades in place once this resolves.
  // Skipped entirely when the caller already precomputed `plannerHref` (a saved sim, whose
  // stored talents this component never needs to re-derive at all): nothing else here reads
  // `talentIndex`, so fetching this file for it would be a request this page has no use for.
  let talents = $state<TalentFile | null>(null);
  onMount(() => {
    if (plannerHref !== undefined) return;
    void loadTalents(character.tree_version, character.class_slug)
      .then((file) => (talents = file))
      .catch(() => (talents = null));
  });
  const talentIndex = $derived<TalentIndex | null>(talents === null ? null : indexTalents(talents));

  const colour = $derived(classColorVar(character.class_slug));
  /** An addon export that carried no gear and no talents (W.3): said as "not read yet" in place
   *  of the slot grid and the talent row, never drawn as seventeen empty slots. A saved sim
   *  (readonly) and a gearless weights result are not exports being read. */
  const unread = $derived(
    !readonly &&
      gearKnown &&
      character.source.kind === 'addon' &&
      !simCharacterHasGear(character) &&
      (talentPoints ?? 0) === 0 &&
      plannerHref === undefined,
  );

  // Names for worn items the per-class files leave out, fetched once and only when a slot
  // needs one (a quality-1 keepsake ring, a totem): the row then names the item and says it
  // is not simmed, instead of quoting an id.
  let outsideNames = $state<Record<string, string>>({});
  const needsOutsideNames = $derived(
    SLOTS.some((slot) => {
      const id = character.gear[slot as Slot];
      return id !== undefined && !items.has(id);
    }),
  );
  $effect(() => {
    if (!needsOutsideNames) return;
    const build = character.tree_version;
    void loadItemNames(build).then((file) => {
      outsideNames = file.names;
    });
  });
  // A combat log records no race, so a character from one arrives with PENDING_RACE and
  // the descriptor asks instead of asserting. Rendering an empty string there would be the
  // page quietly claiming a raceless character, which is the thing Task 7 refuses to do.
  const pending = $derived(needsRace(character));
  /** A fight lacks a race outright; an addon export that does is a bare "Skyborne"
   *  (`characterFromFs1`), so the question is which of the two Skyborne rows. */
  const skyborneOnly = $derived(pending && character.source.kind === 'addon');
  const raceOptions = $derived(
    (skyborneOnly ? races.filter((race) => race.slug.endsWith('-skyborne')) : races).map((race) => ({
      value: race.slug,
      label: race.name,
    })),
  );
  // A build with fewer points than character.level - BASE_LEVEL implies is a part-levelled
  // character: `character.level` is what actually gets simulated (there is no level
  // control on this page), so the line says how many points it actually spends rather than
  // letting "· 60" imply a full-levelled build that ran the tree dry -- and says so
  // honestly (Task 9, spec section 1; the level-aware sim design, 2026-09-27, generalised
  // this from a flat SIM_LEVEL comparison to the character's own level, since a hand-built
  // planner build is the only source that still always simulates at MAX_LEVEL regardless
  // of points -- an addon export with no level section now sets `level` to this same
  // talent-implied number, so there is no discrepancy left to call out for it).
  // Said as its own sentence-case line under the descriptor, never folded into it.
  const levelNote = $derived(
    character.talent_level < character.level
      ? `${character.point_order.length} talent points. ${handoffCopy.simmedAtLevel(character.level)}`
      : '',
  );
  const descriptor = $derived(
    pending
      ? `${specLabel(character.spec)} · ${character.level}`
      : `${raceWords(character.race_slug)} ${specLabel(character.spec)} · ${character.level}`,
  );
  const source = $derived(sourcePill(character.source, new Date(), character.name));
  /** The client's name for each enchant id, so the strip reads as the result's Gear list does. */
  let enchants = $state.raw<ReadonlyMap<number, string>>(new Map());
  $effect(() => {
    const build = character.tree_version;
    loadEnchants(build).then(
      (rows) => (enchants = enchantNames(rows)),
      (error: unknown) => console.error('sim character: the enchant table did not load', error),
    );
  });
  const paperdollSlots = $derived(
    withEnchants(
      slotsFromWorn(character.gear, {
        lookup: (id) => items.get(id),
        outsideName: (id) => outsideNames[String(id)],
        failed: itemsStatus === 'failed',
      }),
      character.gear_slots,
      enchants,
    ),
  );
  const itemLevel = $derived(itemsStatus === 'ready' ? averageItemLevel(paperdollSlots) : null);
  /** The worn items' real tooltips: one host for the doll (BisTooltipHost claims each event, so the
   *  result's own host on the same page never doubles it). */
  const resolveItem = async (id: number) => {
    const { fetchItemTooltipModel } = await import('../../lib/items/lookup');
    return fetchItemTooltipModel(character.tree_version, character.class_slug, id);
  };
  const split = $derived(
    character.point_order.length > 0 ? character.point_order.length : (talentPoints ?? 0),
  );
</script>

{#snippet changeSource()}
  {#if !readonly}
    <button
      type="button"
      class={button('secondary', 'sm')}
      onclick={onchange}
      data-testid="sim-change-source"
    >
      {simCopy.changeSource}
    </button>
  {/if}
{/snippet}

<div class="mx-[18px] md:mx-0">
  <Section title={gearCopy.paperdoll.simHead} head={changeSource} panel testid="sim-character">
    <div class="flex flex-col gap-4">
      <div class="flex flex-wrap items-center gap-3">
        <Crest cls={character.class_slug} size={40} testid="sim-character-class-icon" />
        {#if character.name !== ''}
          <span class="text-[17px] font-semibold" style={`color: ${colour}`} data-testid="paperdoll-name">
            {character.name}
          </span>
        {/if}
        <span class="text-muted text-[13px]" data-testid="paperdoll-descriptor">{descriptor}</span>
        {#if itemLevel !== null && !unread && gearKnown}
          <span class="text-muted text-[13px]" data-testid="sim-gear-ilvl"
            >{gearCopy.paperdoll.averageItemLevel(itemLevel)}</span
          >
        {/if}
        <span class="pill pill-site" data-testid="paperdoll-source">{source}</span>
      </div>
      {#if !unread && gearKnown}
        {#snippet doll()}
          <GearDoll
            slots={paperdollSlots}
            mode="worn"
            build={character.tree_version}
            classSlug={character.class_slug}
            loading={itemsStatus === 'loading'}
            failed={itemsStatus === 'failed'}
            title={character.name === '' ? gearCopy.listLabel : character.name}
            form="list"
            head={false}
            label={gearCopy.listLabel}
            testid="paperdoll"
          />
        {/snippet}
        {#if plated}
          <div role="region" aria-label={gearCopy.listLabel} data-testid="sim-gear">{@render doll()}</div>
        {:else}
          <div data-testid="sim-gear">{@render doll()}</div>
        {/if}
      {/if}
      {#if pending}
        <!-- One question, asked once. The run control stays disabled until it is answered,
           because sim/request.ParseRace refuses an empty race at the boundary and a failed
           run is a worse answer than a question. -->
        <div class="flex flex-wrap items-center gap-2">
          <label class="text-strong text-[13px]" for="sim-race-field" data-testid="sim-race-prompt"
            >{skyborneOnly ? simCopy.pickSkyborne : simCopy.pickRace}</label
          >
          <Select
            id="sim-race-field"
            class="w-[200px]"
            size="compact"
            value=""
            placeholder={skyborneOnly ? simCopy.pickSkybornePlaceholder : simCopy.pickRacePlaceholder}
            options={raceOptions}
            testid="sim-race-picker"
            onchange={(event) => onrace((event.currentTarget as HTMLSelectElement).value)}
          />
        </div>
      {/if}

      {#if levelNote !== ''}
        <p class="text-muted text-[13px]" data-testid="sim-level-note">{levelNote}</p>
      {/if}

      {#if unread}
        <ExportNotRead capturedAt={character.source.captured_at} testid="sim-export-not-read" />
      {:else if !gearKnown}
        <p class="text-muted text-[13px]" data-testid="sim-no-gear">{simCopy.savedNoGear}</p>
      {/if}

      {#if !unread}
        <div class="flex flex-wrap items-baseline gap-3 text-[13px]">
          {#if levelNote === ''}<span class="text-muted text-label font-bold">Talents</span>{/if}
          <!-- A saved sim's stored request carries the engine's talent *string*, not the
         planner's point order (Task 17), so its `point_order` is always empty --
         `talentPoints` (the caller's own digit-sum of that string, SavedSim.svelte) is what
         keeps `split` honest there instead (2026-09-21 result-page review round 3). Neither
         source ever produces a genuine zero for a level-60 character, so `split > 0` still
         means "a real count is known" here, not "the build spent something" specifically --
         showing "0 points" would claim a build spent none, which this guard still avoids
         for the one remaining case (character/talentPoints both absent) where the count is
         simply unknown. -->
          {#if split > 0 && levelNote === ''}
            <span class="text-strong numeric" data-testid="sim-talent-count">{split} points</span>
          {/if}
          <!-- 87px wide already clears the 44px hit-target floor, but its own line-height (~20px)
         does not. `h-11` gives the link itself a real 44px box (what the touch-target check
         measures); the matching negative `-my-3` cancels that growth back out of this row's
         own layout contribution, so the fix stays local to the link's hit area rather than
         growing the strip and shifting the /sim skeleton's CLS-tuned reserved height
         ([id].astro's own comment) or anything mounted below it. -->
          <a
            class="-my-3 ml-auto inline-flex h-11 items-center md:my-0 md:h-auto"
            href={plannerHref ?? plannerHrefFor(character, talentIndex)}
            data-testid="sim-open-planner"
          >
            {simCopy.openInPlanner}
          </a>
        </div>
      {/if}
    </div>
  </Section>
</div>
<BisTooltipHost
  build={character.tree_version}
  modelsElementId="sim-character-tooltip-models"
  resolve={resolveItem}
/>
