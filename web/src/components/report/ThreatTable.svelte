<!-- web/src/components/report/ThreatTable.svelte -->
<!-- Threat, with the model's own honesty on the page. The engine's BaseThreat reports
     itself incomplete until the per-class table lands, and a threat number nobody can
     trust is worse than one labelled untrustworthy.

     `row.threat` is a second, independent lie: window.ts's scopeThreat has no threat
     model in the browser to recompute from, so it scales each row by the same ratio the
     actor rows' damage-plus-healing split moved by. That is a `~`, the same mark
     ActorRow's per-ability and per-target splits carry, not the model-incompleteness note
     above -- the two are unrelated and both can apply to the same row at once. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import {
    approximateAriaLabel,
    approximateMark,
    approximateTitle,
    classColorVar,
    formatAmount,
  } from '../../lib/report/format';
  import type { ThreatRow } from '../../lib/report/types';

  let {
    rows,
    classOf = new Map<string, string>(),
    approximate = false,
  }: {
    rows: ThreatRow[];
    classOf?: Map<string, string>;
    /** True when the window is brushed, so threat is a scaled share, not measured. */
    approximate?: boolean;
  } = $props();

  const ordered = $derived([...rows].sort((a, b) => b.threat - a.threat));
  const peak = $derived(ordered.reduce((highest, row) => Math.max(highest, row.threat), 0));
  const incomplete = $derived(ordered.some((row) => !row.complete));
  const modelVersion = $derived(ordered[0]?.model_version ?? '');
  const mark = $derived(approximateMark(approximate));
  const title = $derived(approximateTitle(approximate));
</script>

{#if ordered.length === 0}
  <p class="text-muted text-[14px]" data-testid="table-empty">No threat in this window.</p>
{:else}
  <div class="flex flex-col gap-2" data-testid="threat-table">
    {#if incomplete}
      <p class="text-muted text-[12px]" data-testid="threat-incomplete">
        Threat model {modelVersion} does not yet carry every class's modifiers, so these figures are indicative.
        The per-class table lands with Forever's ability data.
      </p>
    {/if}
    <ul class="flex flex-col">
      {#each ordered as row (row.guid)}
        <li
          class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px] md:grid-cols-[minmax(120px,1.2fr)_minmax(0,3fr)_96px]"
          data-testid={`threat-${row.guid}`}
        >
          <span class="truncate font-semibold" style={`color: ${classColorVar(classOf.get(row.guid))}`}>
            {splitUnitName(row.name).name}
          </span>
          <span class="bg-line-soft col-span-2 block h-[6px] w-full md:col-span-1">
            <span class="bg-gold block h-full" style={`width: ${peak === 0 ? 0 : (row.threat / peak) * 100}%`}
            ></span>
          </span>
          <span
            class="tabular text-right font-mono"
            {title}
            aria-label={approximateAriaLabel(approximate, formatAmount(Math.round(row.threat)))}
          >
            {mark}{formatAmount(Math.round(row.threat))}
          </span>
        </li>
      {/each}
    </ul>
    {#if approximate}
      <p class="text-muted text-[12px]" data-testid="threat-approximate-note">
        Threat is marked {mark} because it is accumulated from damage and healing and scaled to this window's share
        of that total, not recomputed from the model directly.
      </p>
    {/if}
  </div>
{/if}
