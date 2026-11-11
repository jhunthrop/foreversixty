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
  import CharacterName from '../ui/CharacterName.svelte';
  import { splitUnitName } from '../../lib/characters';
  import {
    cleanMechanicRows,
    mechanicRowKey,
    mostOftenHit,
    nightBossGroups,
    playerMechanics,
    unclassifiedAbilities,
    abilityHitsByPlayer,
    meleeBucket,
    unjudgedDeaths,
    rolePlacements,
    judgedRoster as judgedPlayers,
    judgesPlayers,
    mechanicTable,
  } from '../../lib/report/mechanics';
  import { classSlugOf } from '../../lib/report/planner-link';
  import { SpellIcons } from '../../lib/report/spell-icons.svelte';
  import Crest from '../ui/Crest.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import SpellIcon from '../ui/SpellIcon.svelte';
  import Tag, { type TagTone } from '../ui/Tag.svelte';
  import type { ReportState } from '../../lib/report/url';
  import type { MechanicHit, MechanicKind, MechanicRow, Summary } from '../../lib/report/types';
  import { clockDuration } from '../../lib/format/duration';
  import { compact } from '../../lib/format/number';

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
  /**
   * What a cast or debuff that went through cost: the damage its effects dealt and the
   * healing they gave the enemies (engine 0.3.4), so seven unkicked drains that healed
   * the boss rank above a single hit for less. A row parsed before that carries neither
   * and falls back to the count.
   */
  const throughCost = (row: MechanicRow, count: number): number =>
    row.damage === undefined && row.healed === undefined ? count : (row.damage ?? 0) + (row.healed ?? 0);
  const throughWords = (row: MechanicRow): string => {
    const parts = [
      ...((row.damage ?? 0) > 0 ? [`dealt ${compact(row.damage ?? 0)} damage`] : []),
      ...((row.healed ?? 0) > 0 ? [`healed the enemies for ${compact(row.healed ?? 0)}`] : []),
    ];
    return parts.length === 0 ? '' : `; it ${parts.join(' and ')}`;
  };

  /** Who holds which role, so an ability assigned to a role is not charged to its holder. */
  const roleOf = $derived(new Map(summary.roster.map((row) => [row.guid, row.role])));
  /** One line per failure, most costly first: a death outranks any amount of damage. */
  const problems = $derived.by<Problem[]>(() => {
    const out: Problem[] = [];
    for (const row of block.rows) {
      if (row.kind === 'avoidable') {
        for (const hit of row.players ?? []) {
          // The tank standing in the tank's cleave is not a problem; the card says whose
          // it was to place. Anyone else in it is.
          if (row.role !== undefined && roleOf.get(hit.guid) === row.role) continue;
          out.push({
            key: `${mechanicRowKey(row)}-${hit.guid}`,
            row,
            hit,
            cost: hit.damage + (hit.killed ? DEATH_COST : 0),
            subject: `${splitUnitName(hit.name).name}, ${row.name}`,
            text: `${splitUnitName(hit.name).name} took ${row.name} ${hit.hits === 1 ? 'once' : `${hit.hits} times`} for ${compact(hit.damage)} damage${hit.killed ? ' and died to it' : ''}`,
          });
        }
        // The role's own line: the cleave placed on four people is the tank's failure as
        // much as theirs, and the list names it, ranked by everything it cost.
        for (const placement of rolePlacements([row], summary.roster)) {
          out.push({
            key: `${mechanicRowKey(row)}-placed-${placement.hit.guid}`,
            row,
            hit: placement.hit,
            cost: placement.hit.damage + (placement.hit.killed ? DEATH_COST : 0),
            subject: `${splitUnitName(placement.hit.name).name}, ${row.name} placed`,
            text: `${splitUnitName(placement.hit.name).name} placed ${row.name} on ${placement.others} ${placement.others === 1 ? 'player' : 'players'} ${placement.hit.hits === 1 ? 'once' : `${placement.hit.hits} times`} for ${compact(placement.hit.damage)} damage${placement.hit.killed ? ', and someone died to it' : ''}; the ${row.role}'s to place, and the same hits as their rows below, charged here as well`,
          });
        }
      } else if (row.kind === 'interrupt' && (row.casts ?? 0) > (row.stopped ?? 0)) {
        out.push({
          key: `${mechanicRowKey(row)}-through`,
          row,
          cost: throughCost(row, (row.casts ?? 0) - (row.stopped ?? 0)),
          subject: row.name,
          text: `${row.name} went through ${(row.casts ?? 0) - (row.stopped ?? 0)} of ${row.casts} casts${throughWords(row)}`,
        });
      } else if (row.kind === 'dispel' && (row.applied ?? 0) > (row.dispelled ?? 0)) {
        out.push({
          key: `${mechanicRowKey(row)}-uncured`,
          row,
          cost: throughCost(row, (row.applied ?? 0) - (row.dispelled ?? 0)),
          subject: row.name,
          text: `${row.name} ran its course ${(row.applied ?? 0) - (row.dispelled ?? 0)} of ${row.applied} times it landed${throughWords(row)}`,
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
  const players = $derived(playerMechanics(block.rows, judgedPlayers(block, summary.roster)));
  /** An interrupt-only table judges nothing about any one player, so it draws no cards. */
  const showPlayerCards = $derived(judgesPlayers(block));
  /** The Plate's rows: every mechanic the table lists, the ones that went wrong first. */
  const table = $derived(mechanicTable(block.rows, summary.roster));
  const icons = new SpellIcons();
  const KIND_TAG: Record<string, string> = {
    avoidable: 'Avoidable',
    unavoidable: 'Unavoidable',
    interrupt: 'Interrupt',
    dispel: 'Dispel',
  };
  const tagTone = (row: MechanicRow, failed: number): TagTone =>
    row.kind === 'avoidable' && failed > 0 ? 'danger' : 'muted';

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
  const cleanKeys = $derived(new Set(clean.map((row) => mechanicRowKey(row))));

  /** What hit a player and the table does not list. The spec's honesty rule: an ability
      the curator has not reached is unjudged, not absent. */
  const unclassified = $derived(unclassifiedAbilities(summary));
  /** On a boss with no table: the enemy abilities that did damage, each read player by player. */
  const bossAbilities = $derived(unclassified.filter((ability) => ability.damage > 0));
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
    // A player who was only ever in the unjudged pulls has no card and no line, and the
    // silence would otherwise read as a clean sheet or an omission.
    const absent =
      block.judged_players === undefined
        ? []
        : summary.roster
            .filter((row) => !block.judged_players?.includes(row.guid))
            .map((row) => splitUnitName(row.name).name);
    return `Judged: ${names}, ${judgedPulls} of the night's ${nightPulls} pulls${rest > 0 ? `; the other ${rest} ${rest === 1 ? 'pull has' : 'pulls have'} no table yet` : ''}.${
      absent.length > 0 ? ` Not in a judged pull, so not judged here: ${absent.join(', ')}.` : ''
    }`;
  });
  const unjudged = $derived(unjudgedDeaths(summary));

  function takenOf(guid: string): number {
    return summary.damage_taken.find((actor) => actor.guid === guid)?.effective ?? 0;
  }
  function damageTakenPatch(spellId: number, guid?: string): Partial<ReportState> {
    // No target scope on the night: the night's table cannot measure, so a target filter
    // there prorates a total by the target's share and lands on a figure the row never
    // showed. The ability filter alone returns the row's real total.
    return {
      mode: 'analyze',
      view: 'tables',
      tab: 'damage-taken',
      ability: spellId,
      ...(guid ? { source: guid } : {}),
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

  const linkClass =
    'text-gold inline-flex min-h-11 items-center text-[13px] underline-offset-2 hover:underline md:min-h-0';
</script>

{#snippet line(text: string, testid: string)}
  <p class="text-muted px-4 py-4 text-[14px]" data-testid={testid}>{text}</p>
{/snippet}

{#snippet problemLinks(problem: Problem)}
  {#if problem.row.kind === 'avoidable'}
    <a
      href={hrefFor(damageTakenPatch(problem.row.spell_id, problem.hit?.guid))}
      class={linkClass}
      aria-label={`Damage Taken for ${problem.subject}`}
      onclick={(event) => follow(event, damageTakenPatch(problem.row.spell_id, problem.hit?.guid))}
      >Damage taken</a
    >
    {#if problem.hit?.killed}
      <a
        href={hrefFor(deathsPatch(problem.hit.guid, problem.hit.last_ms, problem.row.spell_id))}
        class={linkClass}
        aria-label={`Deaths for ${problem.subject}`}
        onclick={(event) =>
          follow(event, deathsPatch(problem.hit?.guid ?? '', problem.hit?.last_ms, problem.row.spell_id))}
        >Deaths</a
      >
    {/if}
  {:else if problem.row.kind === 'interrupt'}
    <a
      href={hrefFor({ mode: 'analyze', view: 'tables', tab: 'interrupts' })}
      class={linkClass}
      aria-label={`Interrupts for ${problem.subject}`}
      onclick={(event) => follow(event, { mode: 'analyze', view: 'tables', tab: 'interrupts' })}>Interrupts</a
    >
  {:else}
    <a
      href={hrefFor({ mode: 'analyze', view: 'tables', tab: 'dispels' })}
      class={linkClass}
      aria-label={`Dispels for ${problem.subject}`}
      onclick={(event) => follow(event, { mode: 'analyze', view: 'tables', tab: 'dispels' })}>Dispels</a
    >
  {/if}
{/snippet}

{#snippet figure(value: string, label: string)}
  <span class="flex min-w-[64px] flex-col items-end">
    <span class="text-strong text-[15px] leading-5 font-bold tabular-nums">{value}</span>
    <span class="text-muted text-[12px] leading-4">{label}</span>
  </span>
{/snippet}

<div class="flex flex-col gap-[var(--section-gap)]" data-testid="mechanics-mode">
  <Section title="Mechanics" panel flush>
    <div>
      {#if beforeMechanics}
        {@render line(
          'This report was parsed before Mechanics mode existed; parse it again to see mechanics.',
          'mechanics-not-parsed',
        )}
      {:else if !block.table_found && trash}
        {@render line(
          'Mechanics are judged per boss, and this is a trash segment: pick a boss pull, or the whole night.',
          'mechanics-trash',
        )}
      {:else if !block.table_found && nightMode}
        {@render line('No boss of the night has a table yet.', 'mechanics-no-table')}
      {:else if !block.table_found && bossAbilities.length === 0}
        <p class="text-muted px-4 py-4 text-[14px]" data-testid="mechanics-no-table">
          No mechanics for this boss yet.
        </p>
      {:else if !block.table_found}
        <!-- A boss no table covers yet: what its side did to whom, ability by ability, from damage taken. -->
        <div class="px-5 py-2" data-testid="mechanics-boss-abilities">
          {#each bossAbilities as ability (ability.spell_id)}
            <Row
              title={unclassifiedNameShared(ability) ? `${ability.name} #${ability.spell_id}` : ability.name}
              open
              wrapTrail
              testid={`mechanics-ability-${ability.spell_id}`}
            >
              {#snippet lead()}<SpellIcon
                  socket
                  url={icons.of(ability.spell_id, ability.name)}
                  size={24}
                  tip={ability.spell_id}
                />{/snippet}
              {#snippet subContent()}{#each abilityHitsByPlayer(summary, ability.spell_id) as one, i (one.guid)}{i >
                  0
                    ? ' '
                    : ''}<span class="whitespace-nowrap" data-testid="mechanics-hit"
                    >{i > 0 ? '· ' : ''}<CharacterName
                      name={splitUnitName(one.name).name}
                      cls={classOf.get(one.guid)}
                      weight="inherit"
                    />
                    <b class="text-text font-semibold tabular-nums">{compact(one.damage)}</b>
                    <span class="tabular-nums">({one.hits} {one.hits === 1 ? 'hit' : 'hits'})</span></span
                  >{/each}{/snippet}
              {#snippet trail()}
                <span class="text-strong text-[15px] leading-5 font-bold tabular-nums"
                  >{compact(ability.damage)}</span
                >
                <a
                  href={hrefFor(damageTakenPatch(ability.spell_id))}
                  class={linkClass}
                  aria-label={`Damage Taken for ${ability.name}`}
                  onclick={(event) => follow(event, damageTakenPatch(ability.spell_id))}>Damage taken</a
                >
              {/snippet}
            </Row>
          {/each}
        </div>
      {:else if table.length === 0}
        {@render line('Nothing the table lists went wrong.', 'mechanics-table-empty')}
      {:else}
        <div data-testid="mechanics-table">
          {#each table as { row, tally } (mechanicRowKey(row))}
            <Row
              title={row.name}
              sub={[
                cleanKeys.has(mechanicRowKey(row))
                  ? `Nothing to answer for · ${kindWords(row.kind)}`
                  : tally.who,
                nightMode && row.encounter ? `on ${row.encounter}` : '',
                row.note ?? '',
              ]
                .filter((part) => part !== '')
                .join(' · ')}
              testid={`mechanics-row-${mechanicRowKey(row)}`}
            >
              {#snippet lead()}<SpellIcon
                  socket
                  url={icons.of(row.spell_id, row.name)}
                  tip={row.spell_id}
                />{/snippet}
              {#snippet trail()}
                <Tag tone={tagTone(row, tally.failed)}>{KIND_TAG[row.kind] ?? row.kind}</Tag>
                {@render figure(String(tally.failed), 'Failed')}
                {@render figure(compact(tally.damage), 'Damage')}
              {/snippet}
            </Row>
          {/each}
        </div>
      {/if}
    </div>
    {#if !beforeMechanics && block.table_found}
      <p class="text-muted px-5 pt-3 pb-4 text-[13px]">
        {nightMode
          ? `Whole night, from the encounter’s mechanics table. ${judgedSummary} A source pick does not apply here.`
          : 'Whole fight only, from the encounter’s mechanics table.'}
      </p>
    {/if}
  </Section>

  {#if !beforeMechanics && block.table_found}
    <Section title="Problems" meta="Deaths first, then most costly" testid="mechanics-problems" panel>
      {#if problems.length === 0}
        <p class="text-[14px]">Nothing the table lists went wrong.</p>
      {:else}
        <ol class="flex flex-col">
          {#each problems as problem (problem.key)}
            <li>
              <Row
                title={problem.text}
                sub={[
                  nightMode && problem.row.encounter ? `on ${problem.row.encounter}` : '',
                  problem.row.note ?? '',
                ]
                  .filter((part) => part !== '')
                  .join(' · ')}
                open
                wrapTrail
              >
                {#snippet lead()}<SpellIcon
                    socket
                    tip={problem.row.spell_id}
                    url={icons.of(problem.row.spell_id, problem.row.name)}
                    size={24}
                  />{/snippet}
                {#snippet titleContent()}<span class={problem.hit?.killed ? 'text-death' : 'text-strong'}
                    >{problem.text}</span
                  >{/snippet}
                {#snippet trail()}{@render problemLinks(problem)}{/snippet}
              </Row>
            </li>
          {/each}
        </ol>
      {/if}
    </Section>

    {#if nightMode}
      <Section title="Over the night" testid="mechanics-night" panel>
        {#if nightBosses.length === 0}
          <p class="text-[14px]">Nothing the table lists hit anyone on any pull.</p>
        {:else}
          {#each nightBosses as boss (boss.key)}
            <h3 class="text-strong pt-2 text-[15px] font-bold">{boss.name}</h3>
            {#each boss.rows as row (mechanicRowKey(row))}
              {@const worst = mostOftenHit(row.players)}
              <Row
                title={row.name}
                sub={[
                  `hit someone on ${row.pulls_hit ?? 0} of ${boss.pulls} ${boss.pulls === 1 ? 'pull' : 'pulls'}`,
                  worst === undefined ? '' : `most often ${splitUnitName(worst.name).name}`,
                  row.role === undefined
                    ? ''
                    : [
                        `the ${row.role}'s to place`,
                        ...rolePlacements([row], summary.roster).map(
                          (placement) => splitUnitName(placement.hit.name).name,
                        ),
                      ].join(', '),
                ]
                  .filter((part) => part !== '')
                  .join(' · ')}
                open
              >
                {#snippet lead()}<SpellIcon
                    socket
                    tip={row.spell_id}
                    url={icons.of(row.spell_id, row.name)}
                    size={24}
                  />{/snippet}
              </Row>
            {/each}
          {/each}
        {/if}
      </Section>
    {/if}

    {#if melee.damage > 0}
      <Section title="Melee swings" testid="mechanics-melee" panel>
        <p class="text-muted text-[14px] leading-6">
          <span class="text-text tabular-nums">{compact(melee.damage)}</span> damage to
          <span class="tabular-nums">{melee.players}</span>
          {melee.players === 1 ? 'player' : 'players'} from the enemies’ swings{#if melee.most}, most of it on
            <span class="text-text">{splitUnitName(melee.most.name).name}</span> (<span class="tabular-nums"
              >{compact(melee.most.damage)}</span
            >){#if melee.others.players > 0}, the other
              <span class="tabular-nums">{compact(melee.others.damage)}</span> on
              <span class="tabular-nums">{melee.others.players}</span>
              {melee.others.players === 1 ? 'other player' : 'other players'}{/if}{/if}. No table judges
          melee; the Damage taken tab splits it by source.
          <a
            href={hrefFor(damageTakenPatch(0))}
            class={linkClass}
            aria-label="Damage Taken, melee"
            onclick={(event) => follow(event, damageTakenPatch(0))}>Damage taken</a
          >
        </p>
      </Section>
    {/if}

    {#if showPlayerCards}
      <Section title="By player" meta="What hit them, and what did not" testid="mechanics-players" panel>
        {#each players as player (player.guid)}
          <Row title={splitUnitName(player.name).name} open testid={`mechanics-player-${player.guid}`}>
            {#snippet lead()}<Crest cls={classSlugOf(classOf.get(player.guid)) ?? ''} size={28} />{/snippet}
            {#snippet titleContent()}<CharacterName
                name={splitUnitName(player.name).name}
                cls={classOf.get(player.guid)}
                weight="inherit"
              />{/snippet}
            {#snippet subContent()}
              <span class="block">
                <span class="tabular-nums">{compact(player.damage)}</span> avoidable damage
                {#if nightMode}
                  <span
                    title="Over the night the share would divide by every pull's damage taken, judged or not, so it is not shown; open a pull for it"
                    >· share per pull only</span
                  >
                {:else if takenOf(player.guid) > 0}
                  · <span class="tabular-nums"
                    >{Math.round((player.damage / takenOf(player.guid)) * 100)}%</span
                  > of what they took
                {/if}
                {#if player.placed > 0}
                  · <span class="tabular-nums">{compact(player.placed)}</span> placed on others
                {/if}
              </span>
              {#if player.hits.length === 0}
                <span class="block">
                  {nightMode
                    ? 'Nothing avoidable landed on them on the judged pulls, or they were not in those pulls.'
                    : 'Clean: nothing avoidable landed on them.'}
                </span>
              {:else}
                <!-- A tank both hit by the cleave and placing it on others has two entries for
                     one row; a key on the row alone threw on the duplicate and blanked the page. -->
                {#each player.hits as entry (`${mechanicRowKey(entry.row)}${entry.others === undefined ? '' : '-placed'}`)}
                  <span class={['block', entry.hit.killed ? 'text-death' : 'text-text']}>
                    {entry.row.name} · <span class="tabular-nums">{entry.hit.hits}</span>
                    {entry.hit.hits === 1 ? 'hit' : 'hits'}{entry.others === undefined ? '' : ' on'}
                    {#if entry.others !== undefined}
                      <span class="tabular-nums">{entry.others}</span>
                      {entry.others === 1 ? 'player' : 'players'} who are not the {entry.row.role}{/if} ·
                    <span class="tabular-nums">{compact(entry.hit.damage)}</span>
                    damage{#if (entry.hit.absorbed ?? 0) > 0}
                      (<span class="tabular-nums">{compact(entry.hit.absorbed ?? 0)}</span> absorbed){/if}
                    {#if nightMode && entry.row.encounter}· on {entry.row.encounter}{/if}
                    {#if !nightMode}· first at <span class="tabular-nums"
                        >{clockDuration(entry.hit.first_ms)}</span
                      >{/if}
                    {#if entry.hit.killed}· {entry.others === undefined
                        ? 'died to it'
                        : 'someone died to it'}{/if}
                    {#if entry.others !== undefined}· <span
                        title="The table says this is the {entry.row
                          .role}'s to take; where it lands is theirs to set"
                        >the {entry.row.role}'s to place</span
                      >{/if}
                    {#if entry.hit.pulls}· on <span class="tabular-nums">{entry.hit.pulls}</span>
                      {entry.hit.pulls === 1 ? 'pull' : 'pulls'}{/if}
                  </span>
                {/each}
              {/if}
              {#if player.unavoidable.damage > 0}
                <span class="block">
                  Unavoidable: <span class="tabular-nums">{compact(player.unavoidable.damage)}</span>
                  damage from {player.unavoidable.names.join(', ')}.
                </span>
              {/if}
              {#if player.avoided.length > 0}
                <span class="block">Never hit by: {player.avoided.join(', ')}.</span>
              {/if}
            {/snippet}
          </Row>
        {/each}
      </Section>
    {/if}

    {#if unclassified.length > 0}
      <Section title="Not yet classified" testid="mechanics-unclassified" panel>
        <p class="text-muted text-[13px]">
          These enemy abilities hit someone and the table does not list them yet.{nightMode
            ? ' Over the night they are pooled across every boss: a folded ability no longer says which pull it came from.'
            : ''}
        </p>
        {#each unclassified as ability (ability.spell_id)}
          <Row
            title={unclassifiedNameShared(ability) ? `${ability.name} #${ability.spell_id}` : ability.name}
            sub={`${compact(ability.damage)} damage to players · ${ability.players} ${ability.players === 1 ? 'player' : 'players'} hit`}
            open
          >
            {#snippet lead()}<SpellIcon
                socket
                tip={ability.spell_id}
                url={icons.of(ability.spell_id, ability.name)}
                size={24}
              />{/snippet}
            {#snippet trail()}
              <a
                href={hrefFor(damageTakenPatch(ability.spell_id))}
                class={linkClass}
                aria-label={`Damage Taken for ${ability.name}`}
                onclick={(event) => follow(event, damageTakenPatch(ability.spell_id))}>Damage taken</a
              >
            {/snippet}
          </Row>
        {/each}
      </Section>
    {/if}

    {#if unjudged.length > 0}
      <Section title="Deaths the table does not explain" testid="mechanics-unjudged-deaths" panel>
        <p class="text-muted text-[13px]">
          Killed by a swing, by an ability the table does not list, or by nothing the log named: the problems
          above cannot say whether these were avoidable. The Deaths tab has each one in full.
        </p>
        {#each unjudged as death (`${death.guid}-${death.at_ms}`)}
          <Row
            title={splitUnitName(death.name).name}
            sub={`${death.label ? `${death.label} · ` : ''}at ${clockDuration(death.at_ms)} · ${death.by}`}
            open
          >
            {#snippet trail()}
              <a
                href={hrefFor(deathsPatch(death.guid, death.night_ms))}
                class={linkClass}
                aria-label={`Deaths, ${splitUnitName(death.name).name} at ${clockDuration(death.at_ms)}`}
                onclick={(event) => follow(event, deathsPatch(death.guid, death.night_ms))}>Deaths</a
              >
            {/snippet}
          </Row>
        {/each}
      </Section>
    {/if}
  {/if}
</div>
