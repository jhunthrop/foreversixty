<!-- web/src/components/sim/tools/DropResults.svelte -->
<!-- Design 6.3: by source, with the best per boss, then a flat "every upgrade" list.
     Nothing here is a probability -- neither database records drop rates -- so a boss reads
     "3 of the 11 drops here are upgrades", which is what the data supports.

     No `loot` prop: a boss's name comes off `Substitution.SourceName` (contract 10.1 A6,
     resolved once by drop-picks.ts when the request was built), never a second join back to
     loot.json here. -->
<script lang="ts">
  import { humaniseKey } from '../../../lib/sim/humanise';
  import type { Item } from '../../../lib/planner/types';
  import type { BulkResult, Combo } from '../../../lib/sim/bulk-types';
  import {
    comboKey,
    comboRows,
    deltaLabel,
    planItHref,
    rowName,
    sourceNameOfCombo,
    type ComboRow,
  } from '../../../lib/sim/combos';
  import { bulkCopy, toolFixCopy } from '../../../lib/sim/copy';
  import type { UntriedPick } from '../../../lib/sim/drop-picks';
  import { handoffCopy } from '../../../lib/sim/handoff-copy';
  import { confidenceBand, formatMargin } from '../../../lib/sim/estimate';
  import { button } from '../../../lib/ui/button';
  import type { StatColumn } from '../../../lib/ui/stat-table';
  import Button from '../../ui/Button.svelte';
  import Row from '../../ui/Row.svelte';
  import StatRow from '../../ui/StatRow.svelte';
  import StatTable from '../../ui/StatTable.svelte';
  import SubstitutionChips from './SubstitutionChips.svelte';
  import ToolSection from './ToolSection.svelte';

  let {
    result,
    items,
    treeVersion,
    untried,
    onpin,
    plate = false,
  }: {
    result: BulkResult;
    items: ReadonlyMap<number, Item>;
    treeVersion: string;
    /**
     * Ticked picks that contributed zero tried items (drop-picks.ts's
     * `pickedWithNothingTried`) -- named here rather than left unmentioned (newcomer MAJOR,
     * review.md:291-298; dps D34).
     */
    untried: readonly UntriedPick[];
    /** The item, its `drop:<source-id>` origin and the boss/source name it carried in. */
    onpin: (itemId: number, origin: string, sourceName: string) => void;
    /** The results sit on the page's one Plate. */
    plate?: boolean;
  } = $props();

  /** A boss's rows: the gain (primary), then the share. */
  const COLUMNS: readonly StatColumn[] = [
    { id: 'delta', label: bulkCopy.resultsDelta },
    { id: 'percent', label: bulkCopy.resultsPercent },
  ];

  /** A block's subheading inside the one results Section: the body face 700, not caps. */
  const SUBHEAD = 'text-strong text-[15px] leading-5 font-bold';

  const rows = $derived(comboRows(result));

  /** A drops run is one substitution per combination, so a row's origin names its boss. */
  function originOf(combo: Combo): string {
    return combo.substitutions[0]?.origin ?? '';
  }

  /**
   * Contract 10.1 A6: the name rode in on `Candidate.SourceName` and came back on the
   * substitution, so this reads it rather than joining the origin id to loot.json a second
   * time. A result saved before the field existed carries no source name, and the id
   * itself (with its `drop:` prefix stripped) is never shown raw for it -- D48, the
   * newcomer review's own repro was exactly this, a raw `dungeon:blackrock-spire:175245`
   * as a boss name -- so the fallback is humaniseKey's reading of that id instead.
   */
  function bossName(combo: Combo, origin: string): string {
    const sourceName = sourceNameOfCombo(combo);
    if (sourceName !== '') return sourceName;
    const id = origin.replace(/^drop:/, '');
    return id === '' ? origin : humaniseKey(id);
  }

  /**
   * Grouped by origin, in the order each boss's first drop was ticked -- plain arrays, not
   * a `Map`, so nothing here builds a mutable instance of a built-in reactivity has its own
   * opinion about (`svelte/prefer-svelte-reactivity`); this is a throwaway grouping local to
   * one render, never read back reactively.
   */
  const byBoss = $derived.by(() => {
    const origins: string[] = [];
    for (const row of rows) {
      const origin = originOf(row.combo);
      if (!origins.includes(origin)) origins.push(origin);
    }
    return origins.map((origin) => [origin, rows.filter((row) => originOf(row.combo) === origin)] as const);
  });

  /**
   * A genuine upgrade, strictly -- design 6.3's own predicate. A tie (delta exactly 0, e.g.
   * a drop that is already equipped) is not an upgrade and does not belong in "Every
   * upgrade" or count toward a boss's "Best here" line. (Fix round 1, Finding 1: this used
   * to read `>= 0` to work around the checked-in fake engine tying every combination with
   * the equipped baseline exactly -- a real production-semantics change to dodge a test
   * fixture's determinism, and the wrong fix. The e2e spec now stubs a premium server-run
   * result with genuine positive deltas instead, so this predicate can stay what the design
   * says.)
   */
  function isUpgrade(row: ComboRow): boolean {
    return row.combo.delta.mean > 0;
  }

  const upgrades = $derived(rows.filter(isUpgrade));

  function pin(row: ComboRow): void {
    const sub = row.combo.substitutions[0];
    if (sub?.item_id === undefined) return;
    onpin(sub.item_id, sub.origin ?? '', sub.source_name ?? '');
  }
</script>

<ToolSection title={bulkCopy.resultsTitle} testid="sim-combos" {plate}>
  <div class="flex flex-col gap-3">
    <p class="text-muted text-[13px]" data-testid="sim-drops-note">{bulkCopy.dropsNoChance}</p>

    <p class="text-strong text-[15px] tabular-nums" data-testid="sim-equipped-line">
      {bulkCopy.resultsEquipped}: {Math.round(result.equipped.mean).toLocaleString('en-US')}
      ± {formatMargin(confidenceBand(result.equipped))}
    </p>

    <div class="flex flex-col gap-5 pt-2" data-testid="sim-drops-by-boss">
      <h3 class={SUBHEAD}>{bulkCopy.dropsByBoss}</h3>
      {#each byBoss as [origin, group] (origin)}
        {@const wins = group.filter(isUpgrade)}
        <div class="flex flex-col">
          <header class="flex flex-wrap items-baseline justify-between gap-2">
            <h4 class={SUBHEAD}>{bossName(group[0].combo, origin)}</h4>
            <span class="text-muted text-[13px]">{bulkCopy.dropsUpgrades(wins.length, group.length)}</span>
          </header>
          {#if wins.length > 0}
            <p class="text-gold text-[13px] font-bold tabular-nums" data-testid="sim-drops-best">
              {bulkCopy.dropsBest}: {deltaLabel(wins[0].combo.delta)}
            </p>
          {/if}
          <StatTable columns={COLUMNS} bar={false} open label={bossName(group[0].combo, origin)}>
            {#each group as row (comboKey(row))}
              <StatRow
                bar={false}
                figures={[
                  { value: deltaLabel(row.combo.delta) },
                  { value: `${row.percent.toFixed(1)}%`, muted: true },
                ]}
                testid="sim-combo-row"
              >
                {#snippet name()}
                  <SubstitutionChips
                    substitutions={row.combo.substitutions}
                    {items}
                    {treeVersion}
                    classSlug={result.request.character.class}
                  />
                {/snippet}
              </StatRow>
            {/each}
          </StatTable>
        </div>
      {/each}
      {#each untried as pick (pick.key)}
        <div class="flex flex-col" data-testid="sim-drops-untried">
          <!-- Fix round, Minor 4: the by-boss groups beside this one use the <h4> for the name
               and the line under it for the count, never repeating the name in the
               below. -->
          <h4 class={SUBHEAD}>{pick.name}</h4>
          <p class="text-muted text-[13px]">{toolFixCopy.dropsNothingTried}</p>
        </div>
      {/each}
    </div>

    <div class="flex flex-col pt-3" data-testid="sim-drops-flat">
      <h3 class={SUBHEAD}>{bulkCopy.dropsEveryUpgrade}</h3>
      {#if upgrades.length === 0}
        <p class="text-muted pt-1 text-[13px]">{bulkCopy.noGain}</p>
      {:else}
        <div class="pt-1">
          {#each upgrades as row (comboKey(row))}
            {@const pinId = comboKey(row)}
            {@const href = planItHref(result, row.combo, treeVersion)}
            <Row title={rowName(row)} open wrapTrail>
              {#snippet titleContent()}
                <SubstitutionChips
                  substitutions={row.combo.substitutions}
                  {items}
                  {treeVersion}
                  classSlug={result.request.character.class}
                />
              {/snippet}
              {#snippet trail()}
                <span class="text-gold text-[15px] font-bold tabular-nums">{deltaLabel(row.combo.delta)}</span
                >
                {#if href !== null}
                  <a class={button('ghost', 'sm')} {href} data-testid={`sim-drops-plan-it-${pinId}`}
                    >{handoffCopy.planIt}</a
                  >
                {/if}
                <Button size="sm" testid={`sim-drops-pin-${pinId}`} onclick={() => pin(row)}
                  >{bulkCopy.dropsPin}</Button
                >
              {/snippet}
            </Row>
          {/each}
        </div>
      {/if}
    </div>
  </div>
</ToolSection>
