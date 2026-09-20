<!-- web/src/components/sim/tools/DropResults.svelte -->
<!-- Design 6.3: by source, with the best per boss, then a flat "every upgrade" list.
     Nothing here is a probability -- neither database records drop rates -- so a boss reads
     "3 of the 11 drops here are upgrades", which is what the data supports.

     No `loot` prop: a boss's name comes off `Substitution.SourceName` (contract 10.1 A6,
     resolved once by drop-picks.ts when the request was built), never a second join back to
     loot.json here. -->
<script lang="ts">
  import { SECONDARY_BUTTON } from '../../../lib/planner/styles';
  import type { Item } from '../../../lib/planner/types';
  import type { BulkResult, Combo } from '../../../lib/sim/bulk-types';
  import { comboRows, deltaLabel, sourceNameOfCombo, type ComboRow } from '../../../lib/sim/combos';
  import { bulkCopy, toolFixCopy } from '../../../lib/sim/copy';
  import type { UntriedPick } from '../../../lib/sim/drop-picks';
  import { confidenceBand } from '../../../lib/sim/estimate';
  import SubstitutionChips from './SubstitutionChips.svelte';

  let {
    result,
    items,
    treeVersion,
    untried,
    onpin,
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
  } = $props();

  const rows = $derived(comboRows(result));

  /**
   * `slot:item_id`, not the item id alone. A candidate fitting more than one slot carries
   * `Candidate.Slot === ""` (contract 1.3 -- rings, trinkets, weapons) and the planner may
   * try the same item in either of its slots, which is two combinations with one item id:
   * keyed on the id alone Svelte sees a duplicate key, and the pin button's test id is
   * duplicated with it (final whole-branch review, Minor 7). The rank is the fallback for a
   * combination with no substitution at all.
   */
  function comboKey(row: ComboRow): string {
    const sub = row.combo.substitutions[0];
    return sub?.item_id === undefined ? String(row.rank) : `${sub.slot ?? ''}:${sub.item_id}`;
  }

  /** A drops run is one substitution per combination, so a row's origin names its boss. */
  function originOf(combo: Combo): string {
    return combo.substitutions[0]?.origin ?? '';
  }

  /**
   * Contract 10.1 A6: the name rode in on `Candidate.SourceName` and came back on the
   * substitution, so this reads it rather than joining the origin id to loot.json a second
   * time. The id (with its `drop:` prefix stripped) is the fallback for a result saved
   * before the field existed.
   */
  function bossName(combo: Combo, origin: string): string {
    return sourceNameOfCombo(combo) || origin.replace(/^drop:/, '');
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

<section class="mx-[18px] flex flex-col gap-3 md:mx-0" data-testid="sim-combos">
  <p class="text-muted text-[12px]" data-testid="sim-drops-note">{bulkCopy.dropsNoChance}</p>

  <p class="tabular text-strong font-mono text-[14px]" data-testid="sim-equipped-line">
    {bulkCopy.resultsEquipped}: {Math.round(result.equipped.mean).toLocaleString('en-US')}
    ± {Math.round(confidenceBand(result.equipped)).toLocaleString('en-US')}
  </p>

  <section class="flex flex-col gap-3" data-testid="sim-drops-by-boss">
    <h3 class="section-title text-[14px]">{bulkCopy.dropsByBoss}</h3>
    {#each byBoss as [origin, group] (origin)}
      {@const wins = group.filter(isUpgrade)}
      <div class="border-line rounded-panel border p-3">
        <header class="flex flex-wrap items-baseline justify-between gap-2">
          <h4 class="text-strong text-[13px]">{bossName(group[0].combo, origin)}</h4>
          <span class="text-muted text-[12px]">{bulkCopy.dropsUpgrades(wins.length, group.length)}</span>
        </header>
        {#if wins.length > 0}
          <p class="tabular text-gold font-mono text-[13px]" data-testid="sim-drops-best">
            {bulkCopy.dropsBest}: {deltaLabel(wins[0].combo.delta)}
          </p>
        {/if}
        <ul class="flex flex-col">
          {#each group as row (comboKey(row))}
            <li
              class="border-line-soft flex min-h-11 items-center gap-3 border-b px-2 py-1 last:border-b-0"
              data-testid="sim-combo-row"
            >
              <SubstitutionChips substitutions={row.combo.substitutions} {items} {treeVersion} />
              <span class="tabular text-gold ml-auto font-mono text-[13px]">
                {deltaLabel(row.combo.delta)}
              </span>
              <span class="tabular text-muted font-mono text-[12px]">{row.percent.toFixed(1)}%</span>
            </li>
          {/each}
        </ul>
      </div>
    {/each}
    {#each untried as pick (pick.key)}
      <div class="border-line rounded-panel border p-3" data-testid="sim-drops-untried">
        <!-- Fix round, Minor 4: the by-boss cards beside this one use the <h4> for the name
             and the line under it for the count, never repeating the name in the sentence
             below -- this card used to name it twice. -->
        <h4 class="text-strong text-[13px]">{pick.name}</h4>
        <p class="text-muted text-[13px]">{toolFixCopy.dropsNothingTried}</p>
      </div>
    {/each}
  </section>

  <section class="border-line rounded-panel border p-3" data-testid="sim-drops-flat">
    <h3 class="section-title text-[14px]">{bulkCopy.dropsEveryUpgrade}</h3>
    {#if upgrades.length === 0}
      <p class="text-muted text-[13px]">{bulkCopy.noGain}</p>
    {:else}
      <ul class="flex flex-col">
        {#each upgrades as row (comboKey(row))}
          {@const pinId = comboKey(row)}
          <li class="border-line-soft flex min-h-11 items-center gap-3 border-b px-2 py-1 last:border-b-0">
            <SubstitutionChips substitutions={row.combo.substitutions} {items} {treeVersion} />
            <span class="tabular text-gold ml-auto font-mono text-[13px]">
              {deltaLabel(row.combo.delta)}
            </span>
            <button
              type="button"
              class="{SECONDARY_BUTTON} border-line-warm text-nav px-3"
              data-testid={`sim-drops-pin-${pinId}`}
              onclick={() => pin(row)}>{bulkCopy.dropsPin}</button
            >
          </li>
        {/each}
      </ul>
    {/if}
  </section>
</section>
