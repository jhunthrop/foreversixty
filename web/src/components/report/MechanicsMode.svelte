<!-- web/src/components/report/MechanicsMode.svelte -->
<!-- Who stood in what, and what it cost: the problems list a raid leader reads after a
     wipe, from the encounter's curated mechanics table, with a card per player and a
     roll-up over the night. Whole-fight figures: neither the Analyze window nor the source
     scope applies, so the list is always the whole raid's whole fight. That is the spec's
     own ruling, and it goes further than Compare, which rescopes to a brushed window and
     only ignores the source: half a raid's mistakes, or the ten seconds someone happened
     to brush, is not the answer to "what went wrong". -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { classColorVar, formatAmount, formatDuration } from '../../lib/report/format';
  import type { ReportState } from '../../lib/report/url';
  import type { MechanicHit, MechanicKind, MechanicRow, Summary } from '../../lib/report/types';

  let {
    summary,
    classOf,
    nightMode,
    onPatch,
  }: {
    summary: Summary;
    classOf: Map<string, string>;
    nightMode: boolean;
    onPatch: (patch: Partial<ReportState>) => void;
  } = $props();

  const block = $derived(summary.mechanics ?? { table_found: false, rows: [] });

  /**
   * What each kind means in a sentence, for the rows nobody failed. The table is curated
   * JSON from outside this build, so a kind this version does not know is named "not yet
   * classified" rather than rendered as a blank: the spec's honesty rule, at the one
   * layer that can see an unknown kind at all.
   */
  const KIND_WORDS: Record<string, string> = {
    avoidable: 'avoidable',
    unavoidable: 'unavoidable, the fight’s own damage',
    interrupt: 'to interrupt',
    dispel: 'to dispel',
  };
  function kindWords(kind: MechanicKind): string {
    return KIND_WORDS[kind] ?? 'not yet classified';
  }

  interface Problem {
    key: string;
    row: MechanicRow;
    hit?: MechanicHit;
    cost: number;
    text: string;
    /** Who and what the line is about, so each link has an accessible name of its own:
        a column of buttons all called "Damage Taken" names none of them. */
    subject: string;
  }

  /** A death outranks any amount of damage, so it sorts above one rather than beside it. */
  const DEATH_COST = 1e12;

  /** One line per failure, most costly first: a death outranks any amount of damage. */
  const problems = $derived.by<Problem[]>(() => {
    const out: Problem[] = [];
    for (const row of block.rows) {
      if (row.kind === 'avoidable') {
        for (const hit of row.players ?? []) {
          out.push({
            key: `${row.spell_id}-${hit.guid}`,
            row,
            hit,
            cost: hit.damage + (hit.killed ? DEATH_COST : 0),
            subject: `${splitUnitName(hit.name).name}, ${row.name}`,
            text: `${splitUnitName(hit.name).name} took ${row.name} ${hit.hits === 1 ? 'once' : `${hit.hits} times`} for ${formatAmount(hit.damage)} damage${hit.killed ? ' and died to it' : ''}`,
          });
        }
      } else if (row.kind === 'interrupt' && (row.casts ?? 0) > (row.stopped ?? 0)) {
        out.push({
          key: `${row.spell_id}-through`,
          row,
          cost: (row.casts ?? 0) - (row.stopped ?? 0),
          subject: row.name,
          text: `${row.name} went through ${(row.casts ?? 0) - (row.stopped ?? 0)} of ${row.casts} casts`,
        });
      } else if (row.kind === 'dispel' && (row.applied ?? 0) > (row.dispelled ?? 0)) {
        out.push({
          key: `${row.spell_id}-uncured`,
          row,
          cost: (row.applied ?? 0) - (row.dispelled ?? 0),
          subject: row.name,
          text: `${row.name} ran its course ${(row.applied ?? 0) - (row.dispelled ?? 0)} of ${row.applied} times it landed`,
        });
      }
    }
    return out.sort((a, b) => b.cost - a.cost);
  });

  const avoidableRows = $derived(block.rows.filter((row) => row.kind === 'avoidable'));
  /** The night's roll-up is about what happened: a mechanic that hit nobody all night is
      not a "0 pulls" line here, it is one of the rows under "also in the table". */
  const nightRows = $derived(avoidableRows.filter((row) => (row.pulls_hit ?? 0) > 0));

  /**
   * Per player: their avoidable hits, most damage first, and the mechanics that never
   * touched them. Both halves are the card's point -- "what to tell them next week" is
   * as much the list they stayed out of as the list they stood in, and a card with only
   * the failures reads as an accusation rather than a record.
   */
  const players = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const byGuid = new Map<
      string,
      { guid: string; name: string; hits: { row: MechanicRow; hit: MechanicHit }[]; damage: number }
    >();
    for (const row of block.rows) {
      if (row.kind !== 'avoidable') continue;
      for (const hit of row.players ?? []) {
        const found = byGuid.get(hit.guid) ?? { guid: hit.guid, name: hit.name, hits: [], damage: 0 };
        found.hits.push({ row, hit });
        found.damage += hit.damage;
        byGuid.set(hit.guid, found);
      }
    }
    for (const row of summary.roster) {
      if (!byGuid.has(row.guid))
        byGuid.set(row.guid, { guid: row.guid, name: row.name, hits: [], damage: 0 });
    }
    return [...byGuid.values()]
      .map((player) => ({
        ...player,
        avoided: avoidableRows
          .filter((row) => !player.hits.some((entry) => entry.row.spell_id === row.spell_id))
          .map((row) => row.name),
      }))
      .sort((a, b) => b.damage - a.damage);
  });
  /**
   * The rows nobody failed. The table lists unavoidable damage so a reader does not
   * mistake it for a gap in the table, and an avoidable ability nobody stood in is worth
   * the same sentence: silence here would read as "not covered" rather than "nothing
   * went wrong".
   */
  const clean = $derived(block.rows.filter((row) => !problems.some((problem) => problem.row === row)));

  function takenOf(guid: string): number {
    return summary.damage_taken.find((actor) => actor.guid === guid)?.effective ?? 0;
  }
  function openDamageTaken(spellId: number, guid?: string): void {
    onPatch({
      mode: 'analyze',
      view: 'tables',
      tab: 'damage-taken',
      ability: spellId,
      ...(guid ? { source: guid } : {}),
    });
  }

  const linkClass = 'text-gold min-h-11 text-[12px] underline-offset-2 hover:underline md:min-h-0';
</script>

<div class="flex flex-col gap-4" data-testid="mechanics-mode">
  {#if !block.table_found}
    <p class="text-muted text-[14px]" data-testid="mechanics-no-table">
      No mechanics table for this boss yet. A table says which of its abilities are avoidable, which casts to
      interrupt and which debuffs to dispel; without one this page has nothing to judge. Run
      <code class="font-mono text-[13px]">forever-logs mechanics-draft</code> on the log for a draft to review.
    </p>
  {:else}
    <p class="text-muted text-[12px]">
      Whole {nightMode ? 'night' : 'fight'}, from the encounter’s mechanics table. {nightMode
        ? 'Every pull of the night, and the source above does not narrow it.'
        : 'The time window above does not apply here.'}
    </p>

    <section class="flex flex-col gap-1" data-testid="mechanics-problems">
      <h2 class="label text-muted">Problems, most costly first</h2>
      {#if problems.length === 0}
        <p class="text-[14px]">Nothing the table lists went wrong.</p>
      {:else}
        <ol class="flex flex-col">
          {#each problems as problem (problem.key)}
            <li
              class="border-line-soft flex flex-wrap items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px]"
            >
              <span class:text-death={problem.hit?.killed}>{problem.text}</span>
              {#if problem.row.note}<span class="text-muted text-[12px]">{problem.row.note}</span>{/if}
              {#if problem.row.kind === 'avoidable'}
                <button
                  type="button"
                  class={linkClass}
                  aria-label={`Damage Taken for ${problem.subject}`}
                  onclick={() => openDamageTaken(problem.row.spell_id, problem.hit?.guid)}
                  >Damage Taken</button
                >
                {#if problem.hit?.killed}
                  <button
                    type="button"
                    class={linkClass}
                    aria-label={`Deaths for ${problem.subject}`}
                    onclick={() => onPatch({ mode: 'analyze', view: 'tables', tab: 'deaths' })}>Deaths</button
                  >
                {/if}
              {:else if problem.row.kind === 'interrupt'}
                <button
                  type="button"
                  class={linkClass}
                  aria-label={`Interrupts for ${problem.subject}`}
                  onclick={() => onPatch({ mode: 'analyze', view: 'tables', tab: 'interrupts' })}
                  >Interrupts</button
                >
              {:else}
                <button
                  type="button"
                  class={linkClass}
                  aria-label={`Dispels for ${problem.subject}`}
                  onclick={() => onPatch({ mode: 'analyze', view: 'tables', tab: 'dispels' })}>Dispels</button
                >
              {/if}
            </li>
          {/each}
        </ol>
      {/if}
    </section>

    {#if nightMode}
      <section class="flex flex-col gap-1" data-testid="mechanics-night">
        <h2 class="label text-muted">Over the night</h2>
        {#if nightRows.length === 0}
          <p class="text-[14px]">Nothing the table lists hit anyone on any pull.</p>
        {:else}
          <ul class="flex flex-col">
            {#each nightRows as row (row.spell_id)}
              {@const worst = [...(row.players ?? [])].sort((a, b) => (b.pulls ?? 0) - (a.pulls ?? 0))[0]}
              <li class="border-line-soft border-b px-2 py-2 text-[14px]">
                <span class="font-semibold">{row.name}</span>
                hit someone on <span class="tabular font-mono">{row.pulls_hit ?? 0}</span>
                {row.pulls_hit === 1 ? 'pull' : 'pulls'}
                {#if worst !== undefined}
                  · most often {splitUnitName(worst.name).name}
                {/if}
              </li>
            {/each}
          </ul>
        {/if}
      </section>
    {/if}

    <section class="grid grid-cols-1 gap-3 md:grid-cols-2" data-testid="mechanics-players">
      {#each players as player (player.guid)}
        <article
          class="border-line rounded-panel bg-raised flex flex-col gap-2 border p-3"
          data-testid={`mechanics-player-${player.guid}`}
        >
          <h3 class="font-semibold" style={`color: ${classColorVar(classOf.get(player.guid))}`}>
            {splitUnitName(player.name).name}
          </h3>
          <p class="text-[13px]">
            <span class="tabular font-mono">{formatAmount(player.damage)}</span> avoidable damage
            {#if takenOf(player.guid) > 0}
              · <span class="tabular font-mono"
                >{Math.round((player.damage / takenOf(player.guid)) * 100)}%</span
              > of what they took
            {/if}
          </p>
          {#if player.hits.length === 0}
            <p class="text-muted text-[13px]">Clean: nothing avoidable landed on them.</p>
          {:else}
            <ul class="flex flex-col gap-1 text-[13px]">
              {#each player.hits as entry (entry.row.spell_id)}
                <li class:text-death={entry.hit.killed}>
                  {entry.row.name} · <span class="tabular font-mono">{entry.hit.hits}</span>
                  {entry.hit.hits === 1 ? 'hit' : 'hits'} ·
                  <span class="tabular font-mono">{formatAmount(entry.hit.damage)}</span> damage
                  {#if !nightMode}· first at <span class="tabular font-mono"
                      >{formatDuration(entry.hit.first_ms)}</span
                    >{/if}
                  {#if entry.hit.killed}· died to it{/if}
                  {#if entry.hit.pulls}· on <span class="tabular font-mono">{entry.hit.pulls}</span>
                    {entry.hit.pulls === 1 ? 'pull' : 'pulls'}{/if}
                </li>
              {/each}
            </ul>
          {/if}
          {#if player.avoided.length > 0}
            <p class="text-muted text-[13px]">Never hit by: {player.avoided.join(', ')}.</p>
          {/if}
        </article>
      {/each}
    </section>

    {#if clean.length > 0}
      <section class="flex flex-col gap-1" data-testid="mechanics-clean">
        <h2 class="label text-muted">Also in the table, nothing to answer for</h2>
        <ul class="flex flex-col">
          {#each clean as row (row.spell_id)}
            <li class="border-line-soft text-muted border-b px-2 py-2 text-[13px]">
              <span class="text-text font-semibold">{row.name}</span>
              <span>· {kindWords(row.kind)}</span>
              {#if row.note}<span>· {row.note}</span>{/if}
            </li>
          {/each}
        </ul>
      </section>
    {/if}
  {/if}
</div>
