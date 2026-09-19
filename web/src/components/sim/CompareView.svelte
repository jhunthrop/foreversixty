<!-- web/src/components/sim/CompareView.svelte -->
<!-- The sim beside the fight. The words come first and the tables second, because most
     people will read "Bloodthirst cast 33 times, the sim expects 41" and stop -- that is
     the answer, and the tables are the working.
     Nothing here is coloured. A red-and-green diff would be a fourth colour vocabulary on
     a page that already has class colours, rarity colours and percentile tokens, and the
     design system does not allow colour to carry meaning on its own anyway. -->
<script lang="ts">
  import { simCopy } from '../../lib/sim/copy';
  import { pct as pctOf, type Comparison } from '../../lib/sim/compare';

  let {
    comparison,
    simDuration,
    actualDuration,
  }: { comparison: Comparison; simDuration: number; actualDuration: number } = $props();

  // The rounding rule itself (whole percent, own duration) lives once, in compare.ts, so
  // this table's numbers and compareSummaries' own uptime lines can never round differently.
  const pct = (ms: number, of: number): string => `${pctOf(ms, of)}%`;
  const amount = (value: number): string => Math.round(value).toLocaleString('en-US');

  const abilityRow =
    'border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto_auto] items-center gap-x-3 gap-y-1 border-b px-3 py-2 text-[14px] last:border-b-0 md:grid-cols-[minmax(0,1.6fr)_72px_72px_88px_88px]';
</script>

<div class="flex flex-col gap-4" data-testid="sim-compare">
  <p class="text-strong px-[18px] text-[15px] md:px-0" role="status" data-testid="compare-headline">
    {comparison.headline}
  </p>

  {#if comparison.lines.length > 0}
    <ul class="flex flex-col gap-1 px-[18px] md:px-0" data-testid="compare-lines">
      {#each comparison.lines as line (line)}
        <li class="text-[14px]">{line}</li>
      {/each}
    </ul>
  {/if}

  <section
    class="border-line bg-raised rounded-panel mx-[18px] border md:mx-0"
    data-testid="compare-abilities"
  >
    <div class={`${abilityRow} label text-muted hidden md:grid`}>
      <span>{simCopy.compareAbility}</span>
      <span class="text-right">{simCopy.compareActualCasts}</span>
      <span class="text-right">{simCopy.compareSimCasts}</span>
      <span class="text-right">{simCopy.compareActualDamage}</span>
      <span class="text-right">{simCopy.compareSimDamage}</span>
    </div>
    <!-- Keyed and tested by id, never by name: a resolved name can hold a space ("Heroic
         Strike"), a colon ("spell:25286") or a slash, none of which belong in a test id or
         in a Svelte key. -->
    {#each comparison.abilities as row (row.spellId)}
      <div class={abilityRow} data-testid={`compare-ability-${row.spellId}`}>
        <span class="text-strong truncate font-semibold">{row.name}</span>
        <span class="tabular text-right font-mono">{row.actualCasts}</span>
        <span class="tabular text-muted text-right font-mono">{row.simCasts}</span>
        <span class="tabular hidden text-right font-mono md:inline">{amount(row.actualDamage)}</span>
        <span class="tabular text-muted hidden text-right font-mono md:inline">{amount(row.simDamage)}</span>
        <span class="text-muted label col-span-3 md:hidden">
          {amount(row.actualDamage)}
          {simCopy.compareActual} · {amount(row.simDamage)}
          {simCopy.compareSimulated}
        </span>
      </div>
    {/each}
  </section>

  <section class="border-line bg-raised rounded-panel mx-[18px] border md:mx-0" data-testid="compare-auras">
    <div
      class="border-line-soft label text-muted grid min-h-11 grid-cols-[minmax(0,1fr)_72px_72px] items-center gap-x-3 border-b px-3 py-2"
    >
      <span>{simCopy.compareBuff}</span>
      <span class="text-right">{simCopy.compareActual}</span>
      <span class="text-right">{simCopy.compareSimulated}</span>
    </div>
    <!-- Keyed and tested by id, never by name, for the same reason the ability table is:
         two real buffs at different spell ids (a rank mismatch between the sim's loadout
         and the logged fight's) can resolve to the same display name, and a repeated
         {#each} key is a Svelte 5 runtime error. -->
    {#each comparison.auras as row (row.spellId)}
      <div
        class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_72px_72px] items-center gap-x-3 border-b px-3 py-2 text-[14px] last:border-b-0"
        data-testid={`compare-aura-${row.spellId}`}
      >
        <span class="text-strong truncate font-semibold">{row.name}</span>
        <span class="tabular text-right font-mono">{pct(row.actualUptimeMs, actualDuration)}</span>
        <span class="tabular text-muted text-right font-mono">{pct(row.simUptimeMs, simDuration)}</span>
      </div>
    {/each}
  </section>

  <p class="text-muted px-[18px] text-[12px] md:px-0" data-testid="compare-footnote">
    {simCopy.compareFootnote}
  </p>
</div>
