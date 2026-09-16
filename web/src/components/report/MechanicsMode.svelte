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
    meleeBucket,
    unjudgedDeaths,
  } from '../../lib/report/mechanics';
  import type { ReportState } from '../../lib/report/url';
  import type { MechanicHit, MechanicKind, MechanicRow, Summary } from '../../lib/report/types';

  let {
    summary,
    trash = false,
    classOf,
    nightMode,
    onPatch,
    hrefFor = () => '',
  }: {
    summary: Summary;
    /** True for a trash segment: there is no boss to have a table for. */
    trash?: boolean;
    classOf: Map<string, string>;
    nightMode: boolean;
    onPatch: (patch: Partial<ReportState>) => void;
    /** The url a patch would land on, so a link can be opened in a new tab as well as clicked. */
    hrefFor?: (patch: Partial<ReportState>) => string;
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
  // Over the night only the players who were in a judged pull get a card: a clean sheet
  // for someone who was never in the room is not a clean sheet.
  const judgedRoster = $derived(
    block.judged_players === undefined
      ? summary.roster
      : summary.roster.filter((row) => block.judged_players?.includes(row.guid)),
  );
  const players = $derived(playerMechanics(block.rows, judgedRoster));
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
  /** A name two unclassified spells share, or one that reads as the melee swing, gets its id. */
  function unclassifiedNameShared(ability: { spell_id: number; name: string }): boolean {
    return (
      ability.name === 'Melee' ||
      unclassified.some((other) => other.name === ability.name && other.spell_id !== ability.spell_id)
    );
  }
  const melee = $derived(meleeBucket(summary));
  /** Over the night: which bosses have a table, over how many of the night's pulls. */
  const judgedSummary = $derived.by(() => {
    const bosses = (block.bosses ?? []).filter((boss) =>
      block.rows.some((row) => row.encounter_id === boss.encounter_id),
    );
    const judgedPulls = bosses.reduce((sum, boss) => sum + boss.pulls, 0);
    const nightPulls = summary.pulls?.length ?? judgedPulls;
    if (bosses.length === 0) return 'No boss of the night has a table yet.';
    const names = bosses
      .map((boss) => `${boss.name} (${boss.pulls} ${boss.pulls === 1 ? 'pull' : 'pulls'})`)
      .join(', ');
    const rest = nightPulls - judgedPulls;
    return `Judged: ${names}, ${judgedPulls} of the night's ${nightPulls} pulls${rest > 0 ? `; the other ${rest} ${rest === 1 ? 'pull has' : 'pulls have'} no table yet` : ''}.`;
  });
  const unjudged = $derived(unjudgedDeaths(summary));

  function takenOf(guid: string): number {
    return summary.damage_taken.find((actor) => actor.guid === guid)?.effective ?? 0;
  }
  function damageTakenPatch(spellId: number, guid?: string, encounter?: string): Partial<ReportState> {
    // On the night a row belongs to one boss; the target filter carries that boss's name,
    // so the link lands on that boss's pulls rather than the whole night.
    return {
      mode: 'analyze',
      view: 'tables',
      tab: 'damage-taken',
      ability: spellId,
      ...(guid ? { source: guid } : {}),
      ...(nightMode && encounter ? { target: encounter } : {}),
    };
  }
  /**
   * The Deaths tab, scoped to the player and with their card open: the Deaths tab keys an
   * open card by guid and the death's instant on the summary's own clock.
   */
  function deathsPatch(guid: string, nearMs?: number, spellId?: number): Partial<ReportState> {
    // The player's death by this mechanic nearest the instant given (a hit's last instant,
    // or the death's own): a player killed twice by the same mechanic gets the right card.
    const death = summary.deaths
      .filter(
        (entry) => entry.guid === guid && (spellId === undefined || entry.killing_blow?.spell_id === spellId),
      )
      .sort((a, b) => Math.abs(a.at_ms - (nearMs ?? a.at_ms)) - Math.abs(b.at_ms - (nearMs ?? b.at_ms)))[0];
    return {
      mode: 'analyze',
      view: 'tables',
      tab: 'deaths',
      source: guid,
      ...(death ? { openDeaths: [`${death.guid}-${death.at_ms}`] } : {}),
    };
  }
  /** A link that patches the page in place and still opens in a new tab from a middle click. */
  function follow(event: MouseEvent, patch: Partial<ReportState>): void {
    if (event.button !== 0 || event.metaKey || event.ctrlKey || event.shiftKey || event.altKey) return;
    event.preventDefault();
    onPatch(patch);
  }

  const linkClass = 'text-gold min-h-11 text-[12px] underline-offset-2 hover:underline md:min-h-0';
</script>

<div class="flex flex-col gap-4" data-testid="mechanics-mode">
  {#if beforeMechanics}
    <p class="text-muted text-[14px]" data-testid="mechanics-not-parsed">
      This report was parsed before Mechanics mode existed; parse it again to see mechanics.
    </p>
  {:else if !block.table_found && trash}
    <p class="text-muted text-[14px]" data-testid="mechanics-trash">
      Mechanics are judged per boss, and this is a trash segment: pick a boss pull, or the whole night.
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
        ? `${judgedSummary} A source pick does not apply here.`
        : 'Whole fight: a brushed time window does not apply here, so none is offered.'}
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
                <a
                  href={hrefFor(
                    damageTakenPatch(problem.row.spell_id, problem.hit?.guid, problem.row.encounter),
                  )}
                  class={linkClass}
                  aria-label={`Damage Taken for ${problem.subject}`}
                  onclick={(event) =>
                    follow(
                      event,
                      damageTakenPatch(problem.row.spell_id, problem.hit?.guid, problem.row.encounter),
                    )}>Damage Taken</a
                >
                {#if problem.hit?.killed}
                  <a
                    href={hrefFor(deathsPatch(problem.hit.guid, problem.hit.last_ms, problem.row.spell_id))}
                    class={linkClass}
                    aria-label={`Deaths for ${problem.subject}`}
                    onclick={(event) =>
                      follow(
                        event,
                        deathsPatch(problem.hit?.guid ?? '', problem.hit?.last_ms, problem.row.spell_id),
                      )}>Deaths</a
                  >
                {/if}
              {:else if problem.row.kind === 'interrupt'}
                <a
                  href={hrefFor({ mode: 'analyze', view: 'tables', tab: 'interrupts' })}
                  class={linkClass}
                  aria-label={`Interrupts for ${problem.subject}`}
                  onclick={(event) => follow(event, { mode: 'analyze', view: 'tables', tab: 'interrupts' })}
                  >Interrupts</a
                >
              {:else}
                <a
                  href={hrefFor({ mode: 'analyze', view: 'tables', tab: 'dispels' })}
                  class={linkClass}
                  aria-label={`Dispels for ${problem.subject}`}
                  onclick={(event) => follow(event, { mode: 'analyze', view: 'tables', tab: 'dispels' })}
                  >Dispels</a
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

    {#if melee.damage > 0}
      <section class="flex flex-col gap-1" data-testid="mechanics-melee">
        <h2 class="label text-muted">Melee swings</h2>
        <p class="text-muted text-[13px]">
          <span class="text-text tabular font-mono">{formatAmount(melee.damage)}</span> damage to
          <span class="tabular font-mono">{melee.players}</span>
          {melee.players === 1 ? 'player' : 'players'} from the enemies’ swings{#if melee.most}, most of it on
            <span class="text-text">{splitUnitName(melee.most.name).name}</span> (<span
              class="tabular font-mono">{formatAmount(melee.most.damage)}</span
            >){#if melee.others.players > 0}, the other
              <span class="tabular font-mono">{formatAmount(melee.others.damage)}</span> on
              <span class="tabular font-mono">{melee.others.players}</span>
              {melee.others.players === 1 ? 'other player' : 'other players'}{/if}{/if}. A swing lands on
          whoever holds the enemy, so no table judges it; the Damage Taken tab splits it by source.
          <a
            href={hrefFor(damageTakenPatch(0))}
            class={linkClass}
            aria-label="Damage Taken, melee"
            onclick={(event) => follow(event, damageTakenPatch(0))}>Damage Taken</a
          >
        </p>
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
              {#if nightMode}
                <span
                  class="text-muted"
                  title="Over the night the share would divide by every pull's damage taken, judged or not, so it is not shown; open a pull for it"
                  >· share per pull only</span
                >
              {:else if takenOf(player.guid) > 0}
                · <span class="tabular font-mono"
                  >{Math.round((player.damage / takenOf(player.guid)) * 100)}%</span
                > of what they took
              {/if}
            </p>
            {#if player.hits.length === 0}
              <p class="text-muted text-[13px]">
                {nightMode
                  ? 'Nothing avoidable landed on them on the judged pulls, or they were not in those pulls.'
                  : 'Clean: nothing avoidable landed on them.'}
              </p>
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
              <span class="text-text font-semibold"
                >{ability.name}{#if unclassifiedNameShared(ability)}
                  <span
                    class="text-muted ml-1 font-mono text-[11px]"
                    title="Two things share this name; this is spell id {ability.spell_id}"
                    >#{ability.spell_id}</span
                  >{/if}</span
              >
              <span
                >· <span class="tabular font-mono">{formatAmount(ability.damage)}</span> damage to players ·
                <span class="tabular font-mono">{ability.players}</span>
                {ability.players === 1 ? 'player' : 'players'} hit</span
              >
              <a
                href={hrefFor(damageTakenPatch(ability.spell_id))}
                class={linkClass}
                aria-label={`Damage Taken for ${ability.name}`}
                onclick={(event) => follow(event, damageTakenPatch(ability.spell_id))}>Damage Taken</a
              >
            </li>
          {/each}
        </ul>
      </section>
    {/if}
    {#if unjudged.length > 0}
      <section class="flex flex-col gap-1" data-testid="mechanics-unjudged-deaths">
        <h2 class="label text-muted">Deaths the table does not explain</h2>
        <p class="text-muted text-[12px]">
          Killed by a swing, by an ability the table does not list, or by nothing the log named: the problems
          above cannot say whether these were avoidable. The Deaths tab has each one in full.
        </p>
        <ul class="flex flex-col">
          {#each unjudged as death (`${death.guid}-${death.at_ms}`)}
            <li
              class="border-line-soft text-muted flex flex-wrap items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[13px]"
            >
              <span class="text-text font-semibold">{splitUnitName(death.name).name}</span>
              <span
                >{death.label ? `${death.label} · ` : ''}at
                <span class="tabular font-mono">{formatDuration(death.at_ms)}</span>
                · {death.by}</span
              >
              <a
                href={hrefFor(deathsPatch(death.guid, death.night_ms))}
                class={linkClass}
                aria-label={`Deaths, ${splitUnitName(death.name).name} at ${formatDuration(death.at_ms)}`}
                onclick={(event) => follow(event, deathsPatch(death.guid, death.night_ms))}>Deaths</a
              >
            </li>
          {/each}
        </ul>
      </section>
    {/if}
  {/if}
</div>
