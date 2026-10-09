<!-- web/src/components/planner/BandCompare.svelte -->
<!-- Rebuild spec §4.D (answers review findings 1 and 3): the new rail panel comparing a
     build against its own level band's published build -- the band's own set DPS (an
     always-available number, no engine run), a differs-count, a load action, and a link to
     the BiS page's own best-in-slot list for the band (owner ruling §11).

     Fetches `bis/<spec>.json` through the exact `fetchBisFile` surface `BisSlotPopover.
     svelte` already imports -- the planner island never bundles this data, and a spec
     nobody has hovered or compared yet costs exactly one small fetch, cached in memory for
     the life of the page. The per-cell diff and the "load this band" reconstruction are the
     new `lib/planner/band-compare.ts` adapter (`diffAgainstBand`/`loadFromBand`), both thin
     wrappers over existing, unchanged functions (see that module's own header comment).

     Owner ruling §12: this panel always compares against the band of the build's OWN level
     (`bandForLevel(levelReached(order))`) -- there is no band picker here; previewing a
     different bracket is the BiS page's own job, one click away through "Best in slot
     list" below. -->
<script lang="ts">
  import { bandForLevel, bandEntryFor, bisPageHref, fetchBisFile } from '../../lib/bis/hover';
  import type { BisFile, Faction } from '../../lib/bis/types';
  import { diffAgainstBand, loadFromBand, type BandDiffView } from '../../lib/planner/band-compare';
  import { bandCompareCopy } from '../../lib/planner/copy';
  import { slotScoreUnitFor, tankHeadlineForBand } from '../../lib/bis/tank-view';
  import { plannerScoreCopy } from '../../lib/planner/score-unit';
  import { levelReached } from '../../lib/planner/derive';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import { SECONDARY_BUTTON } from '../../lib/planner/styles';
  import { specDisplayName } from '../../lib/sim/spec-label';
  import { specKeyFor } from '../../lib/addon/score';
  import Skeleton from '../ui/Skeleton.svelte';

  let {
    store,
    specOverride,
    onbanddiff,
    onloaded,
    class: className = '',
  }: {
    store: PlannerStore;
    /** `?spec=` (spec §4.I): names which spec's `bis/<spec>.json` this panel reads,
     *  overriding the points-based inference for the life of this mount. */
    specOverride?: string;
    /** Fires with the panel's own diff view (or `null` once there is none) so the tree row
     *  can draw the per-cell "differs from this band" marker (spec §4.E.4) from the same
     *  computation, rather than a second fetch/derivation of its own. */
    onbanddiff?: (diff: BandDiffView | null) => void;
    /** Fires once a load actually applies, with the band's own label -- fix round 2
     *  (ux-designer review finding 2): a band load is never a character/addon import and
     *  must never claim to be one in Planner.svelte's own reconstruction note. */
    onloaded?: (bandLabel: string) => void;
    /** Merged onto the root element -- spec §6's phone placement (above the tree row,
     *  never collapsed) needs an `order-first` the rail's own desktop position does not. */
    class?: string;
  } = $props();

  const faction = $derived<Faction>(store.raceRow?.faction === 'horde' ? 'horde' : 'alliance');
  const band = $derived(bandForLevel(levelReached(store.order)));
  const specKey = $derived(
    specOverride ?? (store.talentIndex ? specKeyFor(store.classSlug, store.split) : null),
  );
  const bandLabel = $derived(band === 60 ? '60' : `${band} to ${band + 9}`);

  // undefined: still loading (or not started). null: fetched, and there is no file for this
  // spec yet -- the honest empty state, never a guess (BisSlotPopover.svelte's own pattern).
  let file = $state<BisFile | null | undefined>(undefined);
  let loadedFor = $state('');

  $effect(() => {
    const build = store.treeVersion;
    const spec = specKey;
    if (spec === null || spec === '' || build === '') {
      file = null;
      loadedFor = '';
      return;
    }
    const key = `${build}::${spec}`;
    if (loadedFor === key) return;
    loadedFor = key;
    file = undefined;
    void fetchBisFile(build, spec).then((result) => {
      if (loadedFor === key) file = result;
    });
  });

  const bandEntry = $derived(
    file !== null && file !== undefined ? bandEntryFor(file, band, faction) : undefined,
  );

  /** The band's headline: the set's DPS, or a tank band's first headline figure (its own
   *  damage is not what a tank is judged on). */
  const bandFigureLine = $derived.by(() => {
    if (bandEntry === undefined) return '';
    const tankFigure = tankHeadlineForBand(bandEntry)?.figures[0];
    return tankFigure === undefined
      ? plannerScoreCopy.bandFigure(bandEntry.set_dps, slotScoreUnitFor(bandEntry))
      : `${tankFigure.value} ${tankFigure.label.toLowerCase()}`;
  });

  const diff = $derived.by(() => {
    if (store.talentIndex === null || bandEntry === undefined) return null;
    const result = diffAgainstBand(store.talentIndex, store.ranks, bandEntry.talents);
    return { ...result, bandLabel } as BandDiffView;
  });

  $effect(() => {
    onbanddiff?.(diff);
  });

  const gearHref = $derived(
    specKey !== null && specKey !== '' ? bisPageHref(specKey, faction, band) : '/bis',
  );

  let confirming = $state(false);

  /** The band's talents, applied in tree/tier/column order -- the same order the
   *  reconstruction itself walks -- so "see what differs" always jumps to a real,
   *  deterministic first cell rather than whichever `Set` iteration happens to yield. */
  function firstDiffingTalentId(): number | null {
    if (store.talentIndex === null || diff === null) return null;
    for (const tree of store.talentIndex.trees) {
      const byTier = [...tree.talents].sort((a, b) => a.tier - b.tier || a.column - b.column);
      for (const talent of byTier) {
        if (diff.diffTalentIds.has(talent.id)) return talent.id;
      }
    }
    return null;
  }

  function scrollToFirstDiff(): void {
    const id = firstDiffingTalentId();
    if (id === null) return;
    document
      .querySelector(`[data-testid="talent-${id}"]`)
      ?.scrollIntoView({ behavior: 'smooth', block: 'center' });
  }

  function applyLoad(): void {
    if (store.talentIndex === null || bandEntry === undefined) return;
    const result = loadFromBand(store.talentIndex, bandEntry.talents);
    store.applyOrder(result.order, {});
    onloaded?.(bandLabel);
  }

  function requestLoad(): void {
    if (store.order.length === 0) {
      applyLoad();
      return;
    }
    confirming = true;
  }

  function confirmLoad(): void {
    confirming = false;
    applyLoad();
  }
</script>

<section
  class="border-line bg-raised rounded-panel flex w-full flex-col gap-3 border p-4 {className}"
  data-testid="band-compare"
>
  <!-- `whitespace-nowrap` on the heading: its own text is the one thing in this panel that
       changes length as a build levels (e.g. "50 to 59 build" -> "60 build"), and letting it
       wrap internally at a narrow rail width shifted everything below it by a text line on
       the exact click that crosses a band boundary (planner-dps.spec.ts's own "never moves
       the trees" case) -- a real, if small, CLS regression this panel must not introduce. -->
  <header class="flex items-baseline justify-between gap-3">
    <h2 class="section-title text-[15px] whitespace-nowrap">{bandCompareCopy.heading(bandLabel)}</h2>
    <span class="text-muted text-[12px] whitespace-nowrap">{bandCompareCopy.headingAside}</span>
  </header>

  {#if file === undefined}
    <!-- Fix round 4 (CI: the facts-rail/footer position tests this panel's own loading-to-
         ready swap was moving, 18-33.5px, now that it reliably finishes mid-test instead of
         before it): this panel's own ready state measures 88px total, at every width and
         points-spent state checked, against this skeleton's old `lines={3}` default at
         140.5px -- a real 52.5px the outer reserve divs (Planner.svelte, tuned for the
         bigger talent-data swap) never accounted for, because this panel's own `bis/<spec>.
         json` fetch resolves on its own schedule, independent of and after that swap. One
         row at 20px plus this section's own header/gap/padding overhead (68.5px, measured
         the same way) lands within a pixel of that same 88px -- close enough that nothing
         downstream moves when this panel settles. -->
    <Skeleton lines={1} rowHeight="h-5" testid="band-compare-skeleton" />
  {:else if file === null || bandEntry === undefined}
    <p class="text-muted text-[13px]" data-testid="band-compare-empty">
      {bandCompareCopy.noLevelingList(specKey !== null ? specDisplayName(specKey) : '')}
    </p>
  {:else if diff !== null}
    <div class="flex flex-col gap-1">
      <span class="tabular text-gold font-mono text-[22px]" data-testid="band-compare-set-dps">
        {bandFigureLine}
      </span>
      <span class="text-muted text-[12px]">{bandCompareCopy.setDpsCaption(bandLabel)}</span>
    </div>

    <p class="text-[13px]" data-testid="band-compare-differs">
      {#if diff.diffCount === 0}
        <span class="font-semibold" style="color: var(--color-kill)">{bandCompareCopy.matches}</span>
      {:else}
        <span class="tabular text-gold font-mono">{diff.diffCount}</span>
        {bandCompareCopy.differsSuffix}
      {/if}
    </p>

    {#if confirming}
      <div class="flex flex-col gap-3" data-testid="band-compare-load-confirm">
        <p class="text-muted text-[13px]">{bandCompareCopy.loadConfirmMessage}</p>
        <div class="flex flex-wrap gap-3">
          <button
            type="button"
            class="{SECONDARY_BUTTON} border-line-warm-strong text-gold px-4"
            data-testid="band-compare-load-confirm-proceed"
            onclick={confirmLoad}
          >
            {bandCompareCopy.loadConfirmProceed}
          </button>
          <button
            type="button"
            class="{SECONDARY_BUTTON} border-line-warm text-text px-4"
            data-testid="band-compare-load-confirm-cancel"
            onclick={() => (confirming = false)}
          >
            {bandCompareCopy.loadConfirmCancel}
          </button>
        </div>
      </div>
    {:else}
      <div class="flex flex-wrap gap-3">
        {#if !store.readOnly}
          <button
            type="button"
            class="{SECONDARY_BUTTON} border-line-warm-strong text-gold px-4"
            data-testid="band-compare-load"
            onclick={requestLoad}
          >
            {bandCompareCopy.loadButton(bandLabel)}
          </button>
        {/if}
        {#if diff.diffCount > 0}
          <button
            type="button"
            class="{SECONDARY_BUTTON} border-line-warm text-text px-4"
            data-testid="band-compare-see-differs"
            onclick={scrollToFirstDiff}
          >
            {bandCompareCopy.seeWhatDiffers}
          </button>
        {/if}
      </div>
    {/if}

    <a
      class="border-line-warm rounded-control flex min-h-11 items-center justify-between border px-3 text-[13px]"
      href={gearHref}
      data-testid="band-compare-gear-link"
    >
      <span class="text-muted">{bandCompareCopy.gearRowLabel(bandLabel)}</span>
      <span class="text-gold flex items-center gap-1 font-semibold">
        {bandCompareCopy.bestInSlotList}
        <span aria-hidden="true">→</span>
      </span>
    </a>
  {/if}
</section>
