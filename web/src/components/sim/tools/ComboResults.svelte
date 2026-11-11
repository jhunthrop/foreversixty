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
  import { buildCodeFor } from '../../../lib/addon/build-code';
  import { indexTalents } from '../../../lib/planner/rules';
  import type { Item, ItemSet, TalentFile } from '../../../lib/planner/types';
  import { slotLabelForClass } from '../../../lib/bis/slot-display-labels';
  import type { BulkResult } from '../../../lib/sim/bulk-types';
  import { codeForCharacterSpec } from '../../../lib/sim/character';
  import { confidenceBand, formatMargin } from '../../../lib/sim/estimate';
  import {
    anyKeepsSetBonus,
    closestOverlappingPair,
    collapsedComboCount,
    comboKey,
    comboRows,
    deltaLabel,
    exactTieGroups,
    headlineFor,
    keepsSetBonus,
    planItHref,
    rowName,
    signedGainLabel,
    slotSummary,
    winningGear,
    winningTalents,
  } from '../../../lib/sim/combos';
  import { bulkCopy, toolFixCopy } from '../../../lib/sim/copy';
  import { handoffCopy } from '../../../lib/sim/handoff-copy';
  import { poolQualityCopy } from '../../../lib/sim/pool-quality-copy';
  import { LINK } from '../../../lib/ui/primitives';
  import type { StatColumn } from '../../../lib/ui/stat-table';
  import CopyForGame from '../../CopyForGame.svelte';
  import Button from '../../ui/Button.svelte';
  import Row from '../../ui/Row.svelte';
  import StatRow from '../../ui/StatRow.svelte';
  import StatTable from '../../ui/StatTable.svelte';
  import SaveSimForm from './SaveSimForm.svelte';
  import SubstitutionChips from './SubstitutionChips.svelte';
  import ToolSection from './ToolSection.svelte';

  import { grouped } from '../../../lib/format/number';
  let {
    result,
    items,
    sets,
    treeVersion,
    partial,
    onsave,
    talents = null,
  }: {
    result: BulkResult;
    items: ReadonlyMap<number, Item>;
    sets: readonly ItemSet[];
    treeVersion: string;
    partial: boolean;
    /** `store.save`, kept a callback rather than a `store` prop -- this component only
     *  ever needs the one method. */
    onsave: (title: string) => Promise<string | null>;
    /** The class's talent file, already loaded by the bulk store: the layout the game
     *  code's talent cells are read from. Null (not loaded, or failed) leaves "Copy for the
     *  game" disabled rather than copying a build with no talents in it. */
    talents?: TalentFile | null;
  } = $props();

  /** The ranking's figure columns: the combination's DPS (primary), its gain, its share. */
  const COLUMNS: readonly StatColumn[] = [
    { id: 'dps', label: bulkCopy.resultsDps },
    { id: 'delta', label: bulkCopy.resultsDelta },
    { id: 'percent', label: bulkCopy.resultsPercent },
  ];

  function dpsOf(mean: number): string {
    return grouped(mean);
  }

  const FOUR_PIECE = 4;
  let keepSet = $state(false);

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

  /**
   * newcomer round 4 / dps D25-pattern: "too close to separate" shown only when it is
   * actually true of an adjacent pair in the ranking, named to that pair.
   */
  const overlap = $derived(closestOverlappingPair(rows));

  /**
   * newcomer round 4 (review.md:83-110): the same pasted build reads a different DPS in
   * the editor's own preview (that build's own gear) and in this table (the CURRENT
   * character's gear, only the talents substituted) -- both correct, never explained.
   * `bulkCopy.talentsIntro` says this once at the top of the page; this repeats it beside
   * the numbers themselves, since that is where the disagreement is actually read.
   */
  const isTalentsMode = $derived(result.request.bulk?.mode === 'talents');

  const equippedFigure = $derived(grouped(result.equipped.mean));
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

  /**
   * Task L8.1: the winning set as the FSB1 build code the addon's build dialog reads -- the
   * leader's gear (`winner`) with each item's own stats, and the leader's talents
   * (`winningTalents`, so a talents run copies the build that won). The old "Copy to addon"
   * copied an FS1 character export here, which the build dialog refuses by design; the
   * sim's own numbers reach the addon through the companion, never inside a build code.
   */
  const talentIndex = $derived(talents === null ? null : indexTalents(talents));
  const gameCode = $derived(
    talentIndex === null
      ? ''
      : buildCodeFor({
          kind: 'sim',
          dataBuild: treeVersion,
          classSlug: result.request.character.class,
          talents: winningTalents(result),
          gear: winner,
          talentIndex,
          items,
        }),
  );
</script>

<ToolSection title={bulkCopy.resultsTitle} testid="sim-combos">
  <div class="flex flex-col gap-3">
    {#if partial}
      <p class="text-strong text-[13px]" role="status">{bulkCopy.partial}</p>
    {/if}

    <p class="text-strong text-[15px] tabular-nums" data-testid="sim-equipped-line">
      {bulkCopy.resultsEquipped}: {equippedFigure} ± {equippedBand}
      <span class="text-muted text-[13px]">{bulkCopy.ranAtStages(result.stages)}</span>
    </p>

    {#if isTalentsMode}
      <p class="text-muted text-[13px]" data-testid="sim-talents-gear-locked">
        {poolQualityCopy.talentsGearLockedNote}
      </p>
    {/if}

    {#if canKeepSet}
      <label class="text-text flex min-h-11 w-fit items-center gap-2 text-[14px]">
        <input
          type="checkbox"
          class="accent-gold h-5 w-5"
          data-testid="sim-keep-set"
          bind:checked={keepSet}
        />
        {bulkCopy.keepFourPiece}
      </label>
    {/if}

    {#if rows.length === 0}
      <p class="text-muted text-[13px]">{bulkCopy.noGain}</p>
    {:else}
      {#if tieGroups.length > 0}
        <p class="text-muted text-[13px]" data-testid="sim-exact-ties">
          {poolQualityCopy.exactTieNote(tieGroups.reduce((total, group) => total + group.length, 0))}
        </p>
      {/if}
      <!-- The ranking as a ui/StatTable: the rank in the lead column, what the combination
           changed as the name, then DPS, gain and share. Each row's "Plan it" link sits on its
           sub line, under the items it would open, with the leader's within-error group named
           there too. Below 1024 the
           table keeps the DPS with the gain under it. -->
      <StatTable
        columns={COLUMNS}
        nameHead={bulkCopy.resultsChange}
        bar={false}
        open
        label={bulkCopy.resultsTitle}
      >
        {#each rows as row (comboKey(row))}
          <!-- Review fix round 1: one function (`comboKey`, combos.ts) serves both the
               each-key and the plan-it test id. -->
          {@const href = planItHref(result, row.combo, treeVersion)}
          <StatRow
            bar={false}
            figures={[
              { value: dpsOf(row.combo.dps.mean) },
              { value: deltaLabel(row.combo.delta) },
              { value: `${row.percent.toFixed(1)}%`, muted: true },
            ]}
            phone={deltaLabel(row.combo.delta)}
            testid="sim-combo-row"
          >
            {#snippet icon()}
              <span class="text-muted block text-right text-[15px] leading-7 font-bold tabular-nums"
                >{row.rank}</span
              >
            {/snippet}
            {#snippet name()}
              <span class="sr-only">{bulkCopy.resultsRank} {row.rank}</span>
              <SubstitutionChips
                substitutions={row.combo.substitutions}
                {items}
                {treeVersion}
                classSlug={result.request.character.class}
              />
            {/snippet}
            {#snippet sub()}
              {#if href !== null}
                <a class="{LINK} text-[13px]" {href} data-testid={`sim-combo-plan-it-${comboKey(row)}`}
                  >{handoffCopy.planIt} ›</a
                >
              {/if}
              {#if row.withinError}
                <span>{href !== null ? ' · ' : ''}{bulkCopy.withinError}</span>
              {/if}
            {/snippet}
          </StatRow>
        {/each}
      </StatTable>
      <!-- Fix round 1, minor 1: nothing to "tell apart" with exactly one finalist.
           newcomer round 4 / dps D25-pattern: shown only when `closestOverlappingPair`
           proves an adjacent pair actually overlaps, named to the pair it is true of. -->
      {#if overlap !== null}
        <p class="text-muted text-[13px]" data-testid="sim-within-error-note">
          {poolQualityCopy.withinErrorNoteNaming(rowName(overlap.a), rowName(overlap.b))}
        </p>
      {/if}
      <!-- Final whole-branch review, Important 2: the run bar's own count (store.combinations)
           is the engine's simCount over the submitted request, before comboRows' de-dupe
           collapses a ring or trinket tried in both slots into one row -- so this table can
           show fewer rows than that count said. -->
      {#if collapsed > 0}
        <p class="text-muted text-[13px]" data-testid="sim-combos-collapsed">
          {toolFixCopy.combosCollapsedNote(collapsed)}
        </p>
      {/if}
    {/if}

    {#if summary.length > 0}
      <div class="pt-3" data-testid="sim-slot-summary">
        <h3 class="text-strong text-[15px] leading-5 font-bold">{bulkCopy.slotSummary}</h3>
        <p class="text-muted text-[13px]">{bulkCopy.slotSummaryNote}</p>
        <div class="pt-1">
          {#each summary as row (`${row.slot}-${row.item_id}`)}
            <Row
              title={row.name}
              sub={slotLabelForClass(row.slot, result.request.character.class)}
              open
              density="slim"
            >
              {#snippet trail()}
                <span class="text-gold text-[15px] font-bold tabular-nums">{signedGainLabel(row.gain)}</span>
              {/snippet}
            </Row>
          {/each}
        </div>
      </div>
    {/if}

    <div class="flex flex-wrap gap-3 pt-2">
      <Button size="sm" href={plannerHref} testid="sim-open-in-planner">{bulkCopy.openInPlanner}</Button>
      <CopyForGame code={gameCode} testid="sim-copy-addon" />
    </div>

    <!-- The bulk store's `save()` has no fallback title of its own, so `saveTitleFor` above is
         always passed explicitly. SaveSimForm is the one save flow every bulk tool page and
         StatWeights share. -->
    <SaveSimForm {onsave} titleFor={saveTitleFor} {canSave} />

    <div class="pt-3" data-testid="sim-rules">
      <h3 class="text-strong text-[15px] leading-5 font-bold">{bulkCopy.rulesTitle}</h3>
      <ul class="text-muted flex list-disc flex-col gap-1 pt-1 pl-5 text-[13px]">
        {#each bulkCopy.rules as rule (rule)}
          <li>{rule}</li>
        {/each}
      </ul>
    </div>
  </div>
</ToolSection>
