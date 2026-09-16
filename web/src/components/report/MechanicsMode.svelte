<!-- web/src/components/report/MechanicsMode.svelte -->
<!-- Who stood in what, and what it cost: the problems list a raid leader reads after a
     wipe, from the encounter's curated mechanics table, with a card per player and a
     roll-up over the night. Whole-fight figures: neither the Analyze window nor the source
     scope applies, so the list is always the whole raid's whole fight. That is the spec's
     own ruling, and it goes further than Compare, which rescopes to a brushed window and
     only ignores the source: half a raid's mistakes, or the ten seconds someone happened
     to brush, is not the answer to "what went wrong".
     The arithmetic lives in lib/report/mechanics.ts; this file is the words and the links. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { formatAmount, formatDuration, classColorVar } from '../../lib/report/format';
  import {
    cleanMechanicRows,
    mechanicRowKey,
    mostOftenHit,
    nightBossGroups,
    playerMechanics,
    unclassifiedAbilities,
  } from '../../lib/report/mechanics';
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

  /** No `mechanics` key at all: the report was parsed by an engine that had none. That is
      not the same answer as "this boss has no table yet", and saying the wrong one sends a
      reader hunting for a table that exists. */
  const beforeMechanics = $derived(summary.mechanics === undefined);
  const block = $derived(summary.mechanics ?? { table_found: false, rows: [] });

  /**
   * What each kind means in a sentence, for the rows nobody failed. The table is curated
   * JSON from outside this build, so a kind this version does not know is named as such
   * rather than rendered as a blank: the spec's honesty rule, at the one layer that can
   * see an unknown kind at all.
   */
  const KIND_WORDS: Record<string, string> = {
    avoidable: 'avoidable',
    unavoidable: 'unavoidable, the fight’s own damage',
    interrupt: 'to interrupt',
    dispel: 'to dispel',
  };
  function kindWords(kind: MechanicKind): string {
    return KIND_WORDS[kind] ?? 'a kind this page does not know';
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
            key: `${mechanicRowKey(row)}-${hit.guid}`,
            row,
            hit,
            cost: hit.damage + (hit.killed ? DEATH_COST : 0),
            subject: `${splitUnitName(hit.name).name}, ${row.name}`,
            text: `${splitUnitName(hit.name).name} took ${row.name} ${hit.hits === 1 ? 'once' : `${hit.hits} times`} for ${formatAmount(hit.damage)} damage${hit.killed ? ' and died to it' : ''}`,
          });
        }
      } else if (row.kind === 'interrupt' && (row.casts ?? 0) > (row.stopped ?? 0)) {
        out.push({
          key: `${mechanicRowKey(row)}-through`,
          row,
          cost: (row.casts ?? 0) - (row.stopped ?? 0),
          subject: row.name,
          text: `${row.name} went through ${(row.casts ?? 0) - (row.stopped ?? 0)} of ${row.casts} casts`,
        });
      } else if (row.kind === 'dispel' && (row.applied ?? 0) > (row.dispelled ?? 0)) {
        out.push({
          key: `${mechanicRowKey(row)}-uncured`,
          row,
          cost: (row.applied ?? 0) - (row.dispelled ?? 0),
          subject: row.name,
          text: `${row.name} ran its course ${(row.applied ?? 0) - (row.dispelled ?? 0)} of ${row.applied} times it landed`,
        });
      }
    }
    return out.sort((a, b) => b.cost - a.cost);
  });

  /** The night's roll-up, per boss: "hit someone on 6 of 18 pulls" needs that boss's pulls. */
  const nightBosses = $derived(nightBossGroups(block));

  /**
   * Per player: their avoidable hits, most damage first, the unavoidable damage the table
   * accounts for, and the mechanics that never touched them. Every half is the card's
   * point -- "what to tell them next week" is as much the list they stayed out of as the
   * list they stood in, and a card with only the failures reads as an accusation.
   */
  const players = $derived(playerMechanics(block.rows, summary.roster));
  /** An interrupt-only table judges nothing about any one player, so it draws no cards. */
  const showPlayerCards = $derived(
    block.rows.some(
      (row) => row.kind === 'avoidable' || (row.kind === 'unavoidable' && (row.players?.length ?? 0) > 0),
    ),
  );

  /**
   * The rows nobody failed. The table lists unavoidable damage so a reader does not
   * mistake it for a gap in the table, and an avoidable ability nobody stood in is worth
   * the same sentence: silence here would read as "not covered" rather than "nothing
   * went wrong".
   */
  const clean = $derived(
    cleanMechanicRows(
      block.rows,
      problems.map((problem) => mechanicRowKey(problem.row)),
    ),
  );

  /** What hit a player and the table does not list. The spec's honesty rule: an ability
      the curator has not reached is unjudged, not absent. */
  const unclassified = $derived(unclassifiedAbilities(summary));

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
  {#if beforeMechanics}
    <p class="text-muted text-[14px]" data-testid="mechanics-not-parsed">
      This report was parsed before Mechanics mode existed; parse it again to see mechanics.
    </p>
  {:else if !block.table_found}
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
              {#if nightMode && problem.row.encounter}<span class="text-muted text-[12px]"
                  >on {problem.row.encounter}</span
                >{/if}
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
      <section class="flex flex-col gap-2" data-testid="mechanics-night">
        <h2 class="label text-muted">Over the night</h2>
        {#if nightBosses.length === 0}
          <p class="text-[14px]">Nothing the table lists hit anyone on any pull.</p>
        {:else}
          {#each nightBosses as boss (boss.key)}
            <div class="flex flex-col gap-1">
              <h3 class="text-strong text-[13px] font-semibold">{boss.name}</h3>
              <ul class="flex flex-col">
                {#each boss.rows as row (mechanicRowKey(row))}
                  {@const worst = mostOftenHit(row.players)}
                  <li class="border-line-soft border-b px-2 py-2 text-[14px]">
                    <span class="font-semibold">{row.name}</span>
                    hit someone on <span class="tabular font-mono">{row.pulls_hit ?? 0}</span>
                    of <span class="tabular font-mono">{boss.pulls}</span>
                    {boss.pulls === 1 ? 'pull' : 'pulls'}
                    {#if worst !== undefined}
                      · most often {splitUnitName(worst.name).name}
                    {/if}
                  </li>
                {/each}
              </ul>
            </div>
          {/each}
        {/if}
      </section>
    {/if}

    {#if showPlayerCards}
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
                {#each player.hits as entry (mechanicRowKey(entry.row))}
                  <li class:text-death={entry.hit.killed}>
                    {entry.row.name} · <span class="tabular font-mono">{entry.hit.hits}</span>
                    {entry.hit.hits === 1 ? 'hit' : 'hits'} ·
                    <span class="tabular font-mono">{formatAmount(entry.hit.damage)}</span> damage
                    {#if nightMode && entry.row.encounter}· on {entry.row.encounter}{/if}
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
            {#if player.unavoidable.damage > 0}
              <p class="text-muted text-[13px]">
                Unavoidable: <span class="tabular font-mono">{formatAmount(player.unavoidable.damage)}</span>
                damage from {player.unavoidable.names.join(', ')}.
              </p>
            {/if}
            {#if player.avoided.length > 0}
              <p class="text-muted text-[13px]">Never hit by: {player.avoided.join(', ')}.</p>
            {/if}
          </article>
        {/each}
      </section>
    {/if}

    {#if clean.length > 0}
      <section class="flex flex-col gap-1" data-testid="mechanics-clean">
        <h2 class="label text-muted">Also in the table, nothing to answer for</h2>
        <ul class="flex flex-col">
          {#each clean as row (mechanicRowKey(row))}
            <li class="border-line-soft text-muted border-b px-2 py-2 text-[13px]">
              <span class="text-text font-semibold">{row.name}</span>
              {#if nightMode && row.encounter}<span>· {row.encounter}</span>{/if}
              <span>· {kindWords(row.kind)}</span>
              {#if row.note}<span>· {row.note}</span>{/if}
            </li>
          {/each}
        </ul>
      </section>
    {/if}

    {#if unclassified.length > 0}
      <section class="flex flex-col gap-1" data-testid="mechanics-unclassified">
        <h2 class="label text-muted">Not yet classified</h2>
        <p class="text-muted text-[12px]">
          These enemy abilities hit someone and the table does not list them, so it does not say whether they
          were avoidable. Run <code class="font-mono text-[12px]">forever-logs mechanics-draft</code> on the
          log for a draft that classifies them.{nightMode
            ? ' Over the night they are pooled across every boss: a folded ability no longer says which pull it came from.'
            : ''}
        </p>
        <ul class="flex flex-col">
          {#each unclassified as ability (ability.spell_id)}
            <li
              class="border-line-soft text-muted flex flex-wrap items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[13px]"
            >
              <span class="text-text font-semibold">{ability.name}</span>
              <span
                >· <span class="tabular font-mono">{formatAmount(ability.damage)}</span> damage to players ·
                <span class="tabular font-mono">{ability.players}</span>
                {ability.players === 1 ? 'player' : 'players'} hit</span
              >
              <button
                type="button"
                class={linkClass}
                aria-label={`Damage Taken for ${ability.name}`}
                onclick={() => openDamageTaken(ability.spell_id)}>Damage Taken</button
              >
            </li>
          {/each}
        </ul>
      </section>
    {/if}
  {/if}
</div>
