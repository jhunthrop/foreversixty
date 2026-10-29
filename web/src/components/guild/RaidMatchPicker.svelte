<!-- web/src/components/guild/RaidMatchPicker.svelte -->
<!-- The match picker (attendance spec 3.2, 4; officers): opens in place under the band's match row
     in the strip's panel. Rows form one radiogroup of the guild's recent logs, the nearest to the
     raid's window first with `Closest` chosen; a log another raid holds stays listed, disabled,
     with the raid it went to linked. State lives in MatchPicker (lib/guild/match-picker.svelte.ts);
     arrows move focus across the free rows, Space and Enter select. `Match this log` is the page's
     one fill while the picker is open. -->
<script lang="ts">
  import type { MatchCandidate } from '../../lib/guild/attendance-contract';
  import {
    attendanceCopy,
    byName,
    matchedElsewhereParts,
    pickerEmpty,
    pickerFigures,
    pickerHelper,
    durationLabel,
  } from '../../lib/guild/attendance-copy';
  import { PHONE_HIT_AREA } from '../../lib/guild/field-styles';
  import { focusOnMount } from '../../lib/guild/focus-now';
  import { MatchPicker, pickerKeyAction } from '../../lib/guild/match-picker.svelte';
  import { RAID_TEXT_BUTTON } from '../../lib/guild/raid-buttons';
  import { SEP } from '../../lib/guild/raid-copy';
  import { stampParts } from '../../lib/guild/raid-time';
  import { guildRaidHref } from '../../lib/characters';
  import { PRIMARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import LoadError from '../ui/LoadError.svelte';
  import Figures from './Figures.svelte';
  import Skeleton from '../ui/Skeleton.svelte';

  let {
    picker,
    zoneName,
    guildPath,
    viewerTimezone,
    guildTimezone,
    locale,
    onMatched,
    onCancel,
  }: {
    picker: MatchPicker;
    zoneName: string;
    guildPath: string;
    viewerTimezone: string;
    guildTimezone: string;
    locale: string;
    /** The log was matched; the page swaps this panel for the night. */
    onMatched: () => void;
    onCancel: () => void;
  } = $props();

  const copy = attendanceCopy.picker;
  const ROW =
    'grid grid-cols-[24px_minmax(0,1fr)] items-start gap-x-3 gap-y-1 px-3 py-2 md:items-center md:min-h-14 md:grid-cols-[24px_minmax(0,1fr)_200px_140px_120px] focus-visible:outline-gold focus-visible:outline-2 focus-visible:-outline-offset-2';
  const rows: Record<string, HTMLElement | undefined> = {};
  let heading = $state<HTMLElement | undefined>();

  // Spec 3.2: once the list is ready, focus still on the heading moves to the chosen row (or the
  // first free one), so the arrows work at once; focus the officer moved elsewhere stays put.
  $effect(() => {
    if (picker.load !== 'ready') return;
    const stop = picker.selected ?? picker.freeIds[0];
    if (stop === undefined || heading === undefined || document.activeElement !== heading) return;
    rows[stop]?.focus();
  });
  const SENTENCE_ID = (id: string): string => `raid-match-elsewhere-${id}`;

  function rowClass(row: MatchCandidate): string {
    return picker.isTaken(row) ? `${ROW} opacity-50` : `${ROW} hover:bg-card-top cursor-pointer`;
  }

  function tabStop(row: MatchCandidate): number {
    const stop = picker.selected ?? picker.freeIds[0];
    return row.report_id === stop ? 0 : -1;
  }

  function onKeydown(event: KeyboardEvent, row: MatchCandidate): void {
    const action = pickerKeyAction(event.key, row.report_id, picker.freeIds);
    if (action === null) return;
    event.preventDefault();
    if (action.kind === 'move') rows[action.reportId]?.focus();
    else picker.select(action.reportId);
  }

  async function submit(): Promise<void> {
    if (await picker.submit()) onMatched();
  }
</script>

{#snippet when(row: MatchCandidate)}
  {@const stamp = stampParts(new Date(row.uploaded_at), viewerTimezone, guildTimezone, locale)}
  <span class="text-[13px]"
    ><span class="text-muted">{copy.uploaded}</span>&#32;<span class="numeric whitespace-nowrap"
      >{stamp.when}</span
    >{#if stamp.guildTime !== null}&#32;<span class="text-muted whitespace-nowrap">{stamp.guildTime}</span
      >{/if}</span
  >
{/snippet}

<section
  class="bg-raised border-line rounded-panel flex flex-col gap-3 border p-4 shadow-[inset_0_1px_0_rgba(229,185,85,.35)] md:p-5"
  aria-labelledby="raid-match-heading"
  data-testid="raid-match-picker"
>
  <!-- Focus lands here on open, whatever state the list is in (the button that opened it is gone). -->
  <h2
    id="raid-match-heading"
    class="section-title"
    tabindex="-1"
    bind:this={heading}
    use:focusOnMount
    data-testid="raid-match-heading"
  >
    {copy.head}
  </h2>
  <!-- The empty list says the zone and the window itself (UX 22), so the helper steps aside. -->
  {#if !(picker.load === 'ready' && picker.rows.length === 0)}
    <p class="text-muted text-[13px]" data-testid="raid-match-helper">
      {pickerHelper(zoneName, picker.zoneKnown)}
    </p>
  {/if}

  {#if picker.load === 'loading'}
    <Skeleton lines={3} rowHeight="h-14" testid="raid-match-skeleton" />
  {:else if picker.load === 'failed'}
    <LoadError message={copy.loadFailed} onRetry={() => void picker.start()} testid="raid-match-load-error" />
  {:else if picker.rows.length === 0}
    <p class="text-text text-[14px]" data-testid="raid-match-empty">{pickerEmpty(zoneName)}</p>
  {:else}
    <div
      role="radiogroup"
      aria-label={copy.radiogroupLabel}
      class="flex flex-col"
      data-testid="raid-match-rows"
    >
      {#each picker.rows as row (row.report_id)}
        {@const taken = picker.isTaken(row)}
        {@const elsewhere =
          row.matched_raid === null ? null : matchedElsewhereParts(row.matched_raid, guildTimezone, locale)}
        <div class="border-line-soft border-t first:border-t-0">
          <div
            bind:this={rows[row.report_id]}
            role="radio"
            aria-checked={picker.selected === row.report_id ? 'true' : 'false'}
            aria-disabled={taken ? 'true' : undefined}
            aria-describedby={elsewhere === null ? undefined : SENTENCE_ID(row.report_id)}
            tabindex={taken ? undefined : tabStop(row)}
            class={rowClass(row)}
            onclick={() => picker.select(row.report_id)}
            onkeydown={(event) => onKeydown(event, row)}
            data-testid={`raid-match-row-${row.report_id}`}
          >
            <span
              class="border-line-warm relative mt-1 size-4 rounded-full border md:mt-0"
              aria-hidden="true"
              data-testid="raid-match-radio"
            >
              {#if picker.selected === row.report_id}<i class="bg-gold absolute inset-[3px] rounded-full"
                ></i>{/if}
            </span>
            <!-- Phone, three lines (UX 15): zone, Closest and the duration; the upload time; then
                 `by Kraggor · 4 pulls · 1 kill`. From md the time, figures and duration are cells. -->
            <span class="flex min-w-0 flex-col">
              <span class="flex items-center gap-2">
                <span class="text-strong text-[14px] font-semibold whitespace-nowrap">{row.zone_name}</span>
                {#if row.closest}<span class="pill pill-site" data-testid="raid-match-closest"
                    >{copy.closest}</span
                  >{/if}
              </span>
              <!-- Phone (final check): line 2 the upload time, never broken; line 3 who, the figures
                   and the duration, figures in mono. -->
              <span class="md:hidden">{@render when(row)}</span>
              <span class="text-muted text-[12px]"
                >{byName(row.uploaded_by_name)}<span class="md:hidden"
                  >{SEP}<Figures text={pickerFigures(row.pulls, row.kills)} />{SEP}<Figures
                    text={durationLabel(row.duration_ms)}
                  /></span
                ></span
              >
            </span>
            <span class="hidden md:inline">{@render when(row)}</span>
            <span class="hidden text-[13px] md:inline">{pickerFigures(row.pulls, row.kills)}</span>
            <span class="numeric hidden text-[13px] md:inline">{durationLabel(row.duration_ms)}</span>
          </div>
          {#if elsewhere !== null && row.matched_raid !== null}
            <!-- Outside the radio, so its link is not a control inside one. -->
            <p
              id={SENTENCE_ID(row.report_id)}
              class="text-muted px-3 pb-2 pl-[48px] text-[12px]"
              data-testid="raid-match-elsewhere"
            >
              {elsewhere.lead}<a
                href={guildRaidHref(guildPath, row.matched_raid.id)}
                class="text-text underline {PHONE_HIT_AREA}">{elsewhere.title}</a
              >{elsewhere.tail}
            </p>
          {/if}
        </div>
      {/each}
    </div>
  {/if}

  <div class="flex flex-wrap items-center gap-3">
    {#if picker.load === 'ready' && picker.rows.length > 0}
      <button
        type="button"
        class="{PRIMARY_BUTTON_FIXED} focus-visible:outline-gold px-4 focus-visible:outline-2 focus-visible:outline-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
        disabled={!picker.canMatch}
        onclick={() => void submit()}
        data-testid="raid-match-submit">{picker.busy ? copy.matching : copy.match}</button
      >
    {/if}
    <button type="button" class={RAID_TEXT_BUTTON} onclick={onCancel} data-testid="raid-match-cancel"
      >{copy.cancel}</button
    >
  </div>
  <div class="min-h-5">
    {#if picker.line !== null}
      {#key picker.line}
        <p
          class="text-wipe text-[13px]"
          role="alert"
          tabindex="-1"
          use:focusOnMount
          data-testid="raid-match-status"
        >
          {picker.line}
        </p>
      {/key}
    {/if}
  </div>
</section>
