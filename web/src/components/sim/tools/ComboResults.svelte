<!-- web/src/components/sim/tools/ComboResults.svelte -->
<!-- Design 3.3. Every number here comes from the result: the ranking, the within-error
     grouping and each delta's own error are the planner's, and this component neither sorts
     nor recomputes any of them.

     "Open in planner" goes through `codeForCharacterSpec` (character.ts), not a hand-built
     `encodeFS1` call: that helper exists precisely so a winning combination's enchant and
     suffix survive into the planner link (it encodes FS1 version 2), where a manual
     `encodeFS1(...)` over `gearFromSlots(winner)` would silently drop both -- `gearFromSlots`
     returns a bare slot->item_id map with no room for either. -->
<script lang="ts">
  import { SECONDARY_BUTTON } from '../../../lib/planner/styles';
  import { SLOT_LABELS, type Item, type ItemSet, type Slot } from '../../../lib/planner/types';
  import { addonStringFor } from '../../../lib/sim/addon-export';
  import type { BulkResult } from '../../../lib/sim/bulk-types';
  import { codeForCharacterSpec } from '../../../lib/sim/character';
  import { confidenceBand, formatMargin } from '../../../lib/sim/estimate';
  import {
    anyKeepsSetBonus,
    collapsedComboCount,
    comboKey,
    comboRows,
    deltaLabel,
    exactTieGroups,
    headlineFor,
    keepsSetBonus,
    planItHref,
    signedGainLabel,
    slotSummary,
    winningGear,
  } from '../../../lib/sim/combos';
  import { bulkCopy, toolFixCopy } from '../../../lib/sim/copy';
  import { handoffCopy } from '../../../lib/sim/handoff-copy';
  import { poolQualityCopy } from '../../../lib/sim/pool-quality-copy';
  import SaveSimForm from './SaveSimForm.svelte';
  import SubstitutionChips from './SubstitutionChips.svelte';

  let {
    result,
    items,
    sets,
    treeVersion,
    partial,
    onsave,
  }: {
    result: BulkResult;
    items: ReadonlyMap<number, Item>;
    sets: readonly ItemSet[];
    treeVersion: string;
    partial: boolean;
    /** `store.save`, kept a callback rather than a `store` prop -- this component only
     *  ever needs the one method. */
    onsave: (title: string) => Promise<string | null>;
  } = $props();

  const FOUR_PIECE = 4;
  let keepSet = $state(false);
  let addonCopied = $state(false);

  const rows = $derived(
    comboRows(result).filter((row) => !keepSet || keepsSetBonus(row.combo, result, items, sets, FOUR_PIECE)),
  );
  /**
   * How many of `result.combos` `comboRows`' own de-dupe folded away -- computed off the
   * result itself, not off `rows` above: `rows` also runs the "keep 4-piece" filter, an
   * unrelated, player-chosen reduction that has nothing to do with the ring/trinket
   * both-slots collapse this note explains (final whole-branch review, Important 2). Tying
   * the note to that filter too would make it flicker (or misreport "0 collapsed") every
   * time the player ticks the box, for a reason the note does not describe.
   */
  const collapsed = $derived(collapsedComboCount(result));
  const summary = $derived(slotSummary(result));
  const winner = $derived(winningGear(result));

  /**
   * dps D24: "Only combinations keeping a 4-piece set bonus" and its checkbox used to show
   * over every result, including a character with no 4-piece set anywhere in play, where
   * ticking it can only empty the table. Shown only when the result proves it can matter.
   */
  const canKeepSet = $derived(anyKeepsSetBonus(result, items, sets, FOUR_PIECE));

  /** dps D36: rows tied to the decimal on genuinely different items, explained once. */
  const tieGroups = $derived(exactTieGroups(rows));

  const equippedFigure = $derived(Math.round(result.equipped.mean).toLocaleString('en-US'));
  const equippedBand = $derived(formatMargin(confidenceBand(result.equipped)));

  /**
   * The winning set into the planner, through the same ?code= bootstrap /sim already uses.
   * The leader's own row, through `planItHref` -- the identical conversion every per-row
   * "Plan it" link below uses, so this page's one "Open in planner" link can never disagree
   * with a row's own link about what a code encodes. Falls back to `winner` (`winningGear`)
   * when there is no leader to ask (an empty run) or the leader is one `planItHref` cannot
   * honestly open (a named set or consumables alone, `canPlanCombo`'s own rule) -- this
   * button has always shown a link regardless, and still does.
   */
  const plannerHref = $derived(
    (result.combos.length > 0 ? planItHref(result, result.combos[0], treeVersion) : null) ??
      `/planner?code=${encodeURIComponent(
        codeForCharacterSpec({ ...result.request.character, gear: winner }, treeVersion),
      )}`,
  );

  /**
   * `store.save` (bulk-store.svelte.ts) has no fallback title of its own -- unlike `/sim`'s
   * `save()`, which falls back to the settings clause -- so this page always passes one,
   * composed from the result's own headline (combos.ts's `headlineFor`) with `bulkCopy`'s
   * own "nothing beat what you have on" line as the one case that has no winner to name.
   */
  const saveTitleFor = $derived(headlineFor(result) === '' ? bulkCopy.noGain : headlineFor(result));

  /** A stopped run has nothing finished worth naming and saving. */
  const canSave = $derived(!partial);

  async function copyAddon(): Promise<void> {
    try {
      await navigator.clipboard.writeText(
        addonStringFor({
          dataBuild: treeVersion,
          classSlug: result.request.character.class,
          raceSlug: result.request.character.race,
          talents: result.request.character.talents,
          gear: winner,
        }),
      );
      addonCopied = true;
    } catch {
      addonCopied = false;
    }
  }
</script>

<section class="mx-[18px] flex flex-col gap-3 md:mx-0" data-testid="sim-combos">
  {#if partial}
    <p class="text-strong text-[13px]" role="status">{bulkCopy.partial}</p>
  {/if}

  <p class="tabular text-strong font-mono text-[14px]" data-testid="sim-equipped-line">
    {bulkCopy.resultsEquipped}: {equippedFigure} ± {equippedBand}
    <span class="text-muted">{bulkCopy.ranAtStages(result.stages)}</span>
  </p>

  {#if canKeepSet}
    <label class="text-muted flex min-h-11 w-fit items-center gap-2 text-[12px]">
      <input type="checkbox" class="h-5 w-5" data-testid="sim-keep-set" bind:checked={keepSet} />
      {bulkCopy.keepFourPiece}
    </label>
  {/if}

  {#if rows.length === 0}
    <p class="text-muted text-[13px]">{bulkCopy.noGain}</p>
  {:else}
    {#if tieGroups.length > 0}
      <p class="text-muted text-[12px]" data-testid="sim-exact-ties">
        {poolQualityCopy.exactTieNote(tieGroups.reduce((total, group) => total + group.length, 0))}
      </p>
    {/if}
    <!-- A generic ARIA table, not <ul>/<li>: the grid already carries five columns and
         needs their headers announced (finding 4, fix round 1) -- role="row"/"columnheader"/
         "cell" on plain divs, rather than fighting <ul>'s own implicit list semantics.

         Task 7's sixth column ("Plan it") is `auto`-wide and `md:`-only: the five columns
         above already fit a phone tightly (this file's own longstanding note), so a sixth
         fixed-width track there would either force horizontal scroll or squeeze the
         substitution column to nothing. Below `md` the action cell instead spans every
         column of its OWN row (`col-span-full`) -- CSS grid auto-flow cannot fit a
         full-span item into a row the other five cells already fill, so it drops to a new
         row of its own, full width, rather than overlapping them. The header row carries
         the identical sixth cell (empty, `aria-label`d) so every `role="row"` keeps the
         same cell count as there are `columnheader`s at every breakpoint -- this only
         changes CSS placement, never how many cells exist. -->
    <div role="table" class="flex flex-col">
      <div
        class="text-muted grid grid-cols-[28px_minmax(0,2fr)_84px_96px_56px] items-center gap-x-3 gap-y-1 px-2 pb-1 text-[11px] tracking-wide uppercase md:grid-cols-[28px_minmax(0,2fr)_84px_96px_56px_auto]"
        role="row"
      >
        <span role="columnheader">{bulkCopy.resultsRank}</span>
        <span role="columnheader">{bulkCopy.resultsChange}</span>
        <span role="columnheader" class="ml-auto">{bulkCopy.resultsDps}</span>
        <span role="columnheader" class="ml-auto">{bulkCopy.resultsDelta}</span>
        <span role="columnheader" class="ml-auto">{bulkCopy.resultsPercent}</span>
        <!-- No visible label: an action column names itself by its own link text, so this
             carries only the accessible name a screen reader needs for the column. Same
             breakpoint classes as the data cell below, so an empty header never desyncs
             from where the link itself sits. -->
        <span role="columnheader" class="col-span-full md:col-span-1" aria-label={handoffCopy.planIt}></span>
      </div>
      {#each rows as row, index (comboKey(row))}
        <!-- Review fix round 1: this each-key used to hash the substitutions inline here,
             a second copy of what `comboKey` (combos.ts) already computes for the row's own
             `data-testid` below. `comboKey` is at least as discriminating as the old inline
             hash -- it keys every substitution the same way (kind-qualified for a
             non-item one, slot+item_id for an item one) and additionally tells apart two
             rows carrying the same item at different enchants/suffixes, which the old hash
             never did -- so one function now serves both the each-key and the test id. -->
        <!-- Fix round 1, minor 2: a real visual boundary at the end of the leader's
             within-error group, not only a repeated rank digit in a 28px column. -->
        {@const groupEnd = index === rows.length - 1 || rows[index + 1].combo.group !== row.combo.group}
        {@const href = planItHref(result, row.combo, treeVersion)}
        <div
          class="grid min-h-11 grid-cols-[28px_minmax(0,2fr)_84px_96px_56px] items-center gap-x-3 gap-y-1 border-b px-2 py-2 md:grid-cols-[28px_minmax(0,2fr)_84px_96px_56px_auto] {groupEnd
            ? 'border-line-warm'
            : 'border-line-soft'}"
          role="row"
          aria-label={row.withinError ? bulkCopy.withinError : undefined}
          data-testid="sim-combo-row"
        >
          <span
            class="tabular text-muted font-mono text-[12px]"
            role="cell"
            aria-label={`${bulkCopy.resultsRank} ${row.rank}`}
          >
            {row.rank}
          </span>
          <span role="cell">
            <SubstitutionChips substitutions={row.combo.substitutions} {items} {treeVersion} />
          </span>
          <span
            class="tabular text-strong ml-auto font-mono text-[13px]"
            role="cell"
            aria-label={`${bulkCopy.resultsDps} ${Math.round(row.combo.dps.mean).toLocaleString('en-US')}`}
          >
            {Math.round(row.combo.dps.mean).toLocaleString('en-US')}
          </span>
          <span
            class="tabular text-gold ml-auto font-mono text-[13px]"
            role="cell"
            aria-label={`${bulkCopy.resultsDelta} ${deltaLabel(row.combo.delta)}`}
          >
            {deltaLabel(row.combo.delta)}
          </span>
          <span
            class="tabular text-muted ml-auto font-mono text-[12px]"
            role="cell"
            aria-label={`${bulkCopy.resultsPercent} ${row.percent.toFixed(1)}%`}
          >
            {row.percent.toFixed(1)}%
          </span>
          <span role="cell" class="col-span-full md:col-span-1 md:ml-auto">
            {#if href !== null}
              <a
                class="{SECONDARY_BUTTON} border-line-warm text-nav px-3"
                {href}
                data-testid={`sim-combo-plan-it-${comboKey(row)}`}>{handoffCopy.planIt}</a
              >
            {/if}
          </span>
        </div>
      {/each}
    </div>
    <!-- Fix round 1, minor 1: nothing to "tell apart" with exactly one finalist. -->
    {#if rows.length > 1}
      <p class="text-muted text-[12px]">{bulkCopy.withinErrorNote}</p>
    {/if}
    <!-- Final whole-branch review, Important 2: the run bar's own count (store.combinations)
         is the engine's simCount over the submitted request, before comboRows' de-dupe
         collapses a ring or trinket tried in both slots into one row -- so this table can
         show fewer rows than that count said. A companion to the "tried in both slots" rule
         bullet below (bulkCopy.rules), not a contradiction of it: that bullet says why the
         engine tries both slots, this says why the table shows one row for it. -->
    {#if collapsed > 0}
      <p class="text-muted text-[12px]" data-testid="sim-combos-collapsed">
        {toolFixCopy.combosCollapsedNote(collapsed)}
      </p>
    {/if}
  {/if}

  {#if summary.length > 0}
    <section class="border-line rounded-panel border p-3" data-testid="sim-slot-summary">
      <h3 class="section-title text-[14px]">{bulkCopy.slotSummary}</h3>
      <p class="text-muted text-[12px]">{bulkCopy.slotSummaryNote}</p>
      <ul class="flex flex-col">
        {#each summary as row (`${row.slot}-${row.item_id}`)}
          <li class="border-line-soft flex min-h-11 items-center gap-3 border-b px-2 last:border-b-0">
            <span class="text-muted w-24 text-[12px]">{SLOT_LABELS[row.slot as Slot] ?? row.slot}</span>
            <span class="text-text flex-1 text-[13px]">{row.name}</span>
            <span class="tabular text-gold font-mono text-[13px]">
              {signedGainLabel(row.gain)}
            </span>
          </li>
        {/each}
      </ul>
    </section>
  {/if}

  <div class="flex flex-wrap gap-3">
    <a
      class="{SECONDARY_BUTTON} border-line-warm text-nav px-4"
      href={plannerHref}
      data-testid="sim-open-in-planner">{bulkCopy.openInPlanner}</a
    >
    <button
      type="button"
      class="{SECONDARY_BUTTON} border-line-warm text-nav px-4"
      data-testid="sim-copy-addon"
      onclick={() => void copyAddon()}
    >
      {addonCopied ? bulkCopy.copiedToAddon : bulkCopy.copyToAddon}
    </button>
  </div>

  <!-- The controller's own ruling on this task: the bulk store's `save()` gained a state
       (Task 10) so a browser-lane result can get a `sim_id` and `/sim/<id>` (Task 20) can
       render it. `save()` has no fallback title of its own, so `saveTitleFor` above is
       always passed explicitly. SaveSimForm's three states (closed, open, saved) mirror
       `/sim`'s own save form (SimView.svelte) with the same generic vocabulary (`simCopy`)
       -- saving a sim is the identical concept on every page, so it reuses that copy rather
       than typing four more strings that would say the same thing; the follow-ups worktree
       lifted the form itself out to SaveSimForm.svelte so StatWeights.svelte's save control
       is the same component and flow rather than a second copy of it. -->
  <SaveSimForm {onsave} titleFor={saveTitleFor} {canSave} />

  <section class="border-line rounded-panel border p-3" data-testid="sim-rules">
    <h3 class="section-title text-[14px]">{bulkCopy.rulesTitle}</h3>
    <ul class="text-muted flex list-disc flex-col gap-1 pl-5 text-[12px]">
      {#each bulkCopy.rules as rule (rule)}
        <li>{rule}</li>
      {/each}
    </ul>
  </section>
</section>
