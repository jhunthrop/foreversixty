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
  import {
    comboRows,
    deltaLabel,
    headlineFor,
    keepsSetBonus,
    slotSummary,
    winningGear,
  } from '../../../lib/sim/combos';
  import { bulkCopy, simCopy } from '../../../lib/sim/copy';
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

  let saveOpen = $state(false);
  let saveTitle = $state('');
  let saving = $state(false);
  let saveFailed = $state(false);
  let savedUrl = $state<string | null>(null);
  let savedLinkCopied = $state(false);

  const rows = $derived(
    comboRows(result).filter((row) => !keepSet || keepsSetBonus(row.combo, result, items, sets, FOUR_PIECE)),
  );
  const summary = $derived(slotSummary(result));
  const winner = $derived(winningGear(result));

  const equippedFigure = $derived(Math.round(result.equipped.mean).toLocaleString('en-US'));
  const equippedBand = $derived(Math.round(1.96 * result.equipped.error).toLocaleString('en-US'));

  /** The winning set into the planner, through the same ?code= bootstrap /sim already uses. */
  const plannerHref = $derived(
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

  function openSaveForm(): void {
    saveTitle = saveTitleFor;
    saveFailed = false;
    savedUrl = null;
    saveOpen = true;
  }

  function cancelSave(): void {
    saveOpen = false;
    saveFailed = false;
  }

  async function confirmSave(): Promise<void> {
    saving = true;
    saveFailed = false;
    const id = await onsave(saveTitle);
    saving = false;
    if (id === null) {
      saveFailed = true;
      return;
    }
    saveOpen = false;
    savedUrl = `${window.location.origin}/sim/${id}`;
  }

  async function copySavedLink(): Promise<void> {
    if (savedUrl === null) return;
    try {
      await navigator.clipboard.writeText(savedUrl);
      savedLinkCopied = true;
    } catch {
      savedLinkCopied = false;
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

  <label class="text-muted flex min-h-11 w-fit items-center gap-2 text-[12px]">
    <input type="checkbox" class="h-5 w-5" data-testid="sim-keep-set" bind:checked={keepSet} />
    {bulkCopy.keepFourPiece}
  </label>

  {#if rows.length === 0}
    <p class="text-muted text-[13px]">{bulkCopy.noGain}</p>
  {:else}
    <ul class="flex flex-col">
      {#each rows as row (row.combo.substitutions
        .map((sub) => `${sub.kind}:${sub.slot ?? ''}:${sub.item_id ?? sub.name ?? ''}`)
        .join('|'))}
        <li
          class="border-line-soft grid min-h-11 grid-cols-[28px_minmax(0,2fr)_84px_96px_56px] items-center gap-x-3 border-b px-2 py-2"
          data-testid="sim-combo-row"
        >
          <span class="tabular text-muted font-mono text-[12px]">{row.rank}</span>
          <SubstitutionChips substitutions={row.combo.substitutions} {items} {treeVersion} />
          <span class="tabular text-strong ml-auto font-mono text-[13px]">
            {Math.round(row.combo.dps.mean).toLocaleString('en-US')}
          </span>
          <span class="tabular text-gold ml-auto font-mono text-[13px]">
            {deltaLabel(row.combo.delta)}
          </span>
          <span class="tabular text-muted ml-auto font-mono text-[12px]">
            {row.percent.toFixed(1)}%
          </span>
        </li>
      {/each}
    </ul>
    <p class="text-muted text-[12px]">{bulkCopy.withinErrorNote}</p>
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
              {row.gain === null ? '—' : `+${Math.round(row.gain).toLocaleString('en-US')}`}
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
       always passed explicitly. The three states (closed, open, saved) mirror `/sim`'s own
       save form (SimView.svelte) with the same generic vocabulary (`simCopy`) -- saving a
       sim is the identical concept on every page, so this reuses that copy rather than
       typing four more strings that would say the same thing. -->
  <div class="flex flex-wrap items-center gap-3" data-testid="sim-save">
    {#if savedUrl !== null}
      <label class="sr-only" for="sim-save-link">{simCopy.savedLinkLabel}</label>
      <input
        id="sim-save-link"
        type="text"
        readonly
        value={savedUrl}
        class="border-line-warm rounded-control bg-raised text-text h-11 min-w-0 flex-1 border px-3 text-[14px] md:max-w-[420px]"
        data-testid="sim-save-link"
        onclick={(event) => event.currentTarget.select()}
      />
      <button
        type="button"
        class="{SECONDARY_BUTTON} border-line-warm text-nav px-4"
        onclick={() => void copySavedLink()}
        data-testid="sim-save-copy"
      >
        {savedLinkCopied ? simCopy.copied : simCopy.copyLink}
      </button>
      <a
        class="{SECONDARY_BUTTON} border-line-warm text-nav px-4"
        href={savedUrl}
        target="_blank"
        rel="noopener"
        data-testid="sim-open-new-tab"
      >
        {simCopy.openInNewTab}
      </a>
    {:else if saveOpen}
      <label class="flex flex-col gap-1">
        <span class="label text-muted">{simCopy.saveTitleLabel}</span>
        <input
          type="text"
          bind:value={saveTitle}
          class="border-line-warm rounded-control bg-raised text-text h-11 w-[260px] border px-3 text-[14px]"
          data-testid="sim-save-title"
        />
      </label>
      <button
        type="button"
        class="{SECONDARY_BUTTON} border-line-warm text-nav px-5 disabled:opacity-50"
        disabled={saving}
        onclick={() => void confirmSave()}
        data-testid="sim-save-confirm"
      >
        {saving ? simCopy.savingAction : simCopy.saveAction}
      </button>
      <button
        type="button"
        class="{SECONDARY_BUTTON} border-line-warm text-nav px-4"
        onclick={cancelSave}
        data-testid="sim-save-cancel"
      >
        {simCopy.cancel}
      </button>
      {#if saveFailed}
        <span role="alert" class="text-strong text-[13px]" data-testid="sim-save-error">
          {simCopy.saveFailed}
        </span>
      {/if}
    {:else}
      <button
        type="button"
        class="{SECONDARY_BUTTON} border-line-warm text-nav px-5 disabled:opacity-50"
        disabled={!canSave}
        title={!canSave ? simCopy.saveAbortedDisabled : undefined}
        onclick={openSaveForm}
        data-testid="sim-save-open"
      >
        {simCopy.saveThisSim}
      </button>
    {/if}
  </div>

  <section class="border-line rounded-panel border p-3" data-testid="sim-rules">
    <h3 class="section-title text-[14px]">{bulkCopy.rulesTitle}</h3>
    <ul class="text-muted flex list-disc flex-col gap-1 pl-5 text-[12px]">
      {#each bulkCopy.rules as rule (rule)}
        <li>{rule}</li>
      {/each}
    </ul>
  </section>
</section>
