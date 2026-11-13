<!-- web/src/components/guild/GuildRaiderRanks.svelte -->
<!-- The Settings tab's raider ranks row (addon hub L9): one checkbox per guild rank, highest first,
     each with its holders from the synced roster (rank names are not synced), for "counts as a
     raider". The addon's Guild and Officer tabs read the saved list; until it is set they count
     every level 60. The line over the boxes follows them as they change. Saves through the
     settings PATCH the page runs (its busy and error line); Save shows the shared spinner and
     "Saving…", and "Saved." stays until the next change. The page is officer-only, so this row is
     too. -->
<script lang="ts">
  import { CHECK_ROW } from '../../lib/guild/field-styles';
  import {
    raiderRankRows,
    raiderRanksCopy,
    raiderRanksSummary,
    sameRanks,
    toggleRank,
  } from '../../lib/guild/raider-ranks';
  import type { SetupRankLine } from '../../lib/guild/setup-api';
  import { busyCopy } from '../../lib/ui/copy';
  import Button from '../ui/Button.svelte';
  import SettingsRow from '../ui/SettingsRow.svelte';

  let {
    lines,
    officerMax,
    saved,
    busy,
    disabled,
    justSaved,
    onSave,
    onEdit,
  }: {
    /** The roster's rank lines; null while they load, or `failed` when they did not. */
    lines: readonly SetupRankLine[] | null | 'failed';
    officerMax: number;
    /** The guild's raider ranks as the API holds them; null until set. */
    saved: readonly number[] | null;
    /** A settings request is in flight (this one or another on the page). */
    busy: boolean;
    /** Frozen by a contested claim. */
    disabled: boolean;
    /** The last save was this row's and it took. */
    justSaved: boolean;
    onSave: (ranks: number[]) => void;
    onEdit: () => void;
  } = $props();

  let draft = $derived<number[]>([...(saved ?? [])]);
  const changed = $derived(!sameRanks(draft, saved ?? []));
  const known = $derived(Array.isArray(lines) ? lines : null);
  const rows = $derived(raiderRankRows(known ?? [], officerMax, draft));

  function tick(rankIndex: number, on: boolean): void {
    draft = toggleRank(draft, rankIndex, on);
    onEdit();
  }
</script>

<SettingsRow
  open
  label={raiderRanksCopy.label}
  sub={raiderRanksCopy.explain}
  loading={lines === null}
  error={lines === 'failed' ? raiderRanksCopy.ranksFailed : undefined}
  testid="guild-raider-ranks"
>
  {#snippet below()}
    {#if Array.isArray(lines)}
      {#if rows.length === 0}
        <p class="text-muted pt-1 text-[13px]" data-testid="guild-raider-ranks-none">
          {raiderRanksCopy.noRanks}
        </p>
      {:else}
        <p class="text-strong pt-2 text-[14px]" aria-live="polite" data-testid="guild-raider-ranks-summary">
          {raiderRanksSummary(known, draft)}
        </p>
        <fieldset class="pt-1" data-testid="guild-raider-ranks-list">
          <legend class="sr-only">{raiderRanksCopy.legend}</legend>
          {#each rows as row (row.rankIndex)}
            <label class="{CHECK_ROW} py-1" data-testid={`guild-raider-rank-${row.rankIndex}`}>
              <input
                type="checkbox"
                checked={row.checked}
                disabled={busy || disabled}
                onchange={(event) => tick(row.rankIndex, event.currentTarget.checked)}
                data-testid={`guild-raider-rank-${row.rankIndex}-check`}
              />
              <span class="flex min-w-0 flex-col">
                <span class="text-strong font-semibold"
                  >{row.title}{#if row.role !== ''}<span class="text-muted font-normal"
                      >{` · ${row.role}`}</span
                    >{/if}</span
                >
                <span class="text-muted text-[13px] [overflow-wrap:anywhere]">{row.sub}</span>
              </span>
            </label>
          {/each}
        </fieldset>
        <div class="flex flex-wrap items-center gap-3 pt-2">
          <Button
            size="sm"
            onclick={() => onSave(draft)}
            busy={busy && changed}
            busyLabel={busyCopy.saving}
            disabled={busy || disabled || !changed}
            testid="guild-raider-ranks-save">{raiderRanksCopy.save}</Button
          >
          {#if justSaved && !changed}
            <span class="text-muted text-[13px]" role="status" data-testid="guild-raider-ranks-saved"
              >{raiderRanksCopy.saved}</span
            >
          {/if}
        </div>
      {/if}
    {/if}
  {/snippet}
</SettingsRow>
