<!-- web/src/components/sim/RotationCard.svelte -->
<!-- What the engine was told to press, in words (spec 2026-11-04-sim-results 3.11): always open, the
     curated APL's own notes, this run's casts per fight beside each line, and a line the run never
     cast dimmed. The live /sim page keeps LiveRotationCard and its drawer. -->
<script lang="ts">
  import type { ActionNames } from '../../lib/sim/action-names';
  import { createCardTables } from '../../lib/sim/card-tables.svelte';
  import { classSpecPath } from '../../lib/sim/gear-card';
  import { rotationCardCopy as copy } from '../../lib/sim/input-cards-copy';
  import { rotationActionKeys, rotationRows } from '../../lib/sim/rotation-card';
  import { loadRotationLines, type RotationLines } from '../../lib/sim/rotations';
  import { specDisplayName } from '../../lib/sim/spec-label';
  import type { SimResult } from '../../lib/sim/types';
  import { isStale } from '../../lib/sim/version';
  import ActionLink from '../ui/ActionLink.svelte';
  import Card from '../ui/Card.svelte';
  import ActionMark from './ActionMark.svelte';

  let { result, names }: { result: SimResult; names: ActionNames | null } = $props();

  let lines = $state.raw<RotationLines | null | undefined>(undefined);
  $effect(() => {
    const { spec } = result.request;
    loadRotationLines().then(
      (all) => (lines = all[spec] ?? null),
      (error: unknown) => {
        console.error('sim rotation: the rotation did not load', error);
        lines = null;
      },
    );
  });

  const tables = createCardTables(
    () => result,
    () => names,
    () => [
      ...(lines === null || lines === undefined ? [] : rotationActionKeys(lines)),
      ...result.summary.casts.map((row) => row.spell_name),
    ],
  );

  const spec = $derived(specDisplayName(result.request.spec));
  const rows = $derived(lines === null || lines === undefined ? null : rotationRows(result, lines, tables.names));
  const drifted = $derived(
    lines !== null &&
      lines !== undefined &&
      (result.rotation_sha === undefined
        ? isStale(result.engine_version)
        : lines.sha !== null && result.rotation_sha !== lines.sha),
  );
  const meta = $derived(
    rows === null ? undefined : `${copy.meta(rows.priority.length)}${drifted ? copy.today : ''}`,
  );

  function rankOf(action: string | null): number {
    return action === null ? 0 : tables.visuals(action).rank;
  }
</script>

{#snippet title(name: string, action: string | null)}
  <span class="text-strong text-[15px] font-semibold">{name}</span>
  {#if rankOf(action) > 0}<span class="text-muted ml-1 text-[13px] whitespace-nowrap">{copy.rank(rankOf(action))}</span>{/if}
{/snippet}

<Card
  glyph="list"
  label={copy.label(spec, result.request.character.level)}
  title={copy.title}
  {meta}
  body="none"
  class="mx-[18px] md:mx-0"
  testid="sim-rotation-saved-card"
>
  {#if rows === null}
    {#if lines === null}
      <p class="text-muted border-line-soft border-t px-5 py-4 text-[14px]">{copy.noRotation(spec)}</p>
    {/if}
  {:else}
    {#if rows.prepull.length > 0}
      <h3 class="quiet-label border-line-soft border-t px-5 py-3.5">{copy.beforePull}</h3>
      <ol data-testid="rotation-prepull">
        {#each rows.prepull as row (row.key)}
          <li
            class={`border-line-soft grid grid-cols-[52px_minmax(0,1fr)_auto] items-center gap-3 border-t px-5 py-3 ${row.cast === false ? 'opacity-50' : ''}`}
            data-testid="rotation-prepull-row"
            data-cast={row.cast}
          >
            <span class="text-muted numeric text-[14px] font-semibold whitespace-nowrap">{row.at}</span>
            <ActionMark action={row.action} {tables} size={28}>{@render title(row.name, row.action)}</ActionMark>
            {#if row.cast !== null}
              <span class="text-muted text-right text-[13px]">{row.cast ? copy.cast : copy.notCast}</span>
            {/if}
          </li>
        {/each}
      </ol>
    {/if}
    <h3 class="quiet-label border-line-soft border-t px-5 py-3.5">{copy.thenPriority}</h3>
    <ol data-testid="rotation-priority">
      {#each rows.priority as row (row.key)}
        <li
          class={`border-line-soft grid grid-cols-[24px_minmax(0,1fr)] gap-x-3 border-t px-5 py-3 md:grid-cols-[44px_minmax(0,1fr)_140px] ${row.dimmed ? 'opacity-50' : ''}`}
          data-testid="rotation-line"
          data-dimmed={row.dimmed}
        >
          <span class="text-muted numeric pt-1 text-right text-[14px] font-semibold">{row.number}</span>
          <div class="min-w-0">
            <div class="flex items-start justify-between gap-3">
              <ActionMark action={row.action} {tables} size={28}>{@render title(row.name, row.action)}</ActionMark>
              {#if row.casts !== null}
                <span class="text-muted shrink-0 pt-1 text-[13px] md:hidden">
                  {row.dimmed ? copy.notCast : copy.casts(row.casts)}
                </span>
              {/if}
            </div>
            {#if row.note !== ''}
              <p class="text-text mt-1 max-w-[860px] text-[14px] leading-[21px] md:pl-[38px]">{row.note}</p>
            {/if}
          </div>
          {#if row.casts !== null}
            <span class={`hidden pt-1 text-right md:block ${row.dimmed ? 'text-muted text-[13px]' : 'text-text numeric text-[14px]'}`}>
              {row.dimmed ? copy.notCast : copy.casts(row.casts)}
            </span>
          {/if}
        </li>
      {/each}
    </ol>
    <div class="border-line-soft flex flex-wrap items-center justify-between gap-x-4 border-t px-5 py-1.5">
      <p class="text-muted text-[13px]">{copy.foot(spec)}</p>
      <ActionLink href={`/guides/${classSpecPath(result.request.spec)}`}>{copy.guide(spec)}</ActionLink>
    </div>
  {/if}
</Card>
