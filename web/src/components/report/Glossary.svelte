<!-- web/src/components/report/Glossary.svelte -->
<!-- The words the tables use, in one place a phone can open. Every column heading carries
     the same text as a hover title, which a touch screen never shows; this is the tap
     version, a native disclosure so it costs no script and reads to a screen reader as
     what it is. -->
<script lang="ts">
  const TERMS: { term: string; meaning: string }[] = [
    {
      term: 'Parse',
      meaning:
        'Where a kill sits among every ranked kill of the same boss by the same spec, on this ruleset, as a percentile: 100 is the best, 0 the lowest. A wipe is not ranked; a dash means nothing of this spec has been ranked on this boss yet; a question mark means the rankings could not be reached (Try again asks once more); the whole night has none, since a parse is per kill.',
    },
    {
      term: 'Parse, by role',
      meaning:
        'On the Summary tab a healer is placed by healing per second and everyone else, tanks included, by damage per second, as on Warcraft Logs. The Damage Done and Healing tabs place every row on that table’s own metric, each among their own spec.',
    },
    {
      term: 'Parse, of how many',
      meaning:
        'On a phone the number is followed by how many ranked kills it was placed among; on a desktop that count is in the hover text. Early on a bracket can hold one kill, which stands first of one and reads 100.',
    },
    {
      term: 'Parse colours',
      meaning:
        'The same ladder as Warcraft Logs, so a "grey parse" or an "orange parse" means what it does there: grey under 25, green from 25, blue from 50, purple from 75, orange from 95, pink at 99, gold at 100.',
    },
    {
      term: 'Aura · Buff · Debuff',
      meaning:
        'An effect sitting on a unit for a while: a buff helps the one it is on, a debuff hurts them. The Buffs tab lists the helpful auras on anyone in scope, the Debuffs tab the harmful ones: the raid’s on the enemies, and the enemies’ on the raid.',
    },
    {
      term: 'Uptime',
      meaning:
        'The share of the window an aura was up on that unit. A damage-over-time debuff at 95% uptime was on the boss nearly the whole fight.',
    },
    {
      term: 'Applied',
      meaning: 'How many times an aura was put on that unit inside the window, refreshes not counted.',
    },
    {
      term: 'Damage · DPS',
      meaning: 'Damage done in the window, and the same divided by the window’s length in seconds.',
    },
    {
      term: 'Healing · HPS',
      meaning:
        'Effective healing done: healing that landed on a missing health bar. Overhealing is the rest.',
    },
    { term: 'Taken · DTPS', meaning: 'Damage taken, and the same per second. A tank’s headline figure.' },
    {
      term: 'Active',
      meaning:
        'The share of the window the player spent casting or attacking. Time spent dead is not active. The second per-second figure, marked "while active", divides by that time instead of the whole window, so it is never smaller: it says what they did while they were doing anything.',
    },
    {
      term: 'Ignore events while dead',
      meaning:
        'Leaves out what each player did, and what hit them, from a death to their first cast after it (or the end of the fight if they stayed dead). On one pull the tables are then measured from the events; over the whole night it is not available.',
    },
    {
      term: '#number after a spell',
      meaning:
        'On the Casts tab, two different spells can share one name for the same player (a talent’s version and the base spell, say). Each keeps its own row, and the number after the # is the spell id that tells them apart. It is not the same spell counted twice.',
    },
    {
      term: 'Share (Summary panels)',
      meaning:
        'On the Summary’s damage, healing and damage-taken panels, a share of every player’s total in this window, whatever the Source scope shows: one player on their own still reads their real share of the raid, and an ability that hit only them reads its share of everything the raid took.',
    },
    {
      term: 'Killed by · last hit by',
      meaning:
        'On a death card, "killed by" names the hit that ended them. "Last hit by" means no damage line did: the death came from something the log does not write as damage (a fall, an instant kill, a timer running out), so the card shows the last hit before it instead.',
    },
    {
      term: 'Casting, total',
      meaning:
        'On the Casts tab, the time a player spent casting that spell across the whole fight, every cast added together. An instant spell has none. A dagger marks it as the whole fight’s figure under any window.',
    },
    {
      term: 'Cancelled casts',
      meaning:
        'On the Casts tab, casts with a cast bar that were started and never went off: the player moved, was interrupted, or cancelled it. Instants have none. A cancelled hardcast is time spent for nothing, the gap a rotation can close.',
    },
    {
      term: 'Failed casts',
      meaning:
        'On the Casts tab, presses the game refused: out of range, no target, not enough mana or energy. An instant pressed while out of range fails too, which is why an instant can show more failures than casts. A cast bar cut short is Cancelled, not Failed. Only the player whose client wrote the log has failed casts recorded; everyone else reads a dash.',
    },
    {
      term: 'Mitigated · absorbed · blocked · avoided',
      meaning:
        'Under the Damage Taken table: damage that did not land. Absorbed was soaked by a shield, blocked was cut by a shield block, and avoided hits missed outright: a parry, dodge, miss, or a hit a shield took whole. Measured from the fight’s events on one pull; marked ~ where it is prorated.',
    },
    {
      term: '(n over)',
      meaning:
        'Beside a heal: the part that landed on a full health bar, overhealing. A heal of "+0 (1,719 over)" found nothing to heal.',
    },
    {
      term: 'Up · Just lost',
      meaning:
        'On a death card: the buffs and debuffs on the player at the moment they died, and the ones that dropped in the seconds before.',
    },
    {
      term: 'Parry · dodge · miss · absorb · evade · immune',
      meaning:
        'The ways a hit fails to land, as the log names them. A parry, dodge or miss avoided it outright; absorb means a shield took the whole hit; evade and immune mean the target could not be hit at all just then.',
    },
    {
      term: 'Amount · Share · Hits · Crit · Avg · Max',
      meaning:
        'The columns of an expanded row’s abilities table: the damage or healing that landed, its share of the row’s total (the bar beside it is that share drawn), how many hits and ticks landed, the share of them that were critical, the amount per hit, and the largest single hit. The last column holds notes: what was absorbed, blocked or avoided, or how much of a heal was over.',
    },
    {
      term: 'Overkill',
      meaning:
        'The part of a killing hit beyond the health the player had left. A death line shows the hit whole and the overkill in brackets; the tables leave overkill out unless "Count overkill" is on.',
    },
    {
      term: 'Max health',
      meaning: 'The player’s full health bar at the time, so a hit can be read as a share of it.',
    },
    {
      term: 'Kill · Wipe 59% · pull 1 of 2',
      meaning:
        'How a boss pull ended, and which attempt it was when a boss took more than one. The percentage after a wipe is the boss’s health when the pull ended: how far the pull got.',
    },
    {
      term: 'Source',
      meaning:
        'Narrows every table to the friendlies, the enemies, or one player. Clicking a name in the summary does the same.',
    },
    {
      term: 'Window',
      meaning:
        'A slice of the fight. Drag across the chart, use the sliders, or pick a preset; every table follows.',
    },
    {
      term: '~',
      meaning:
        'A figure split in proportion to the window or a filter rather than measured directly. Totals are exact.',
    },
    {
      term: '9† in the fight list',
      meaning:
        'How many players died on that pull. Beside a figure in a table, † marks a whole-fight figure the summary cannot cut down to a window.',
    },
    { term: 'Split', meaning: 'Talent points per tree, in tree order.' },
    {
      term: 'Threat',
      meaning:
        'Threat accumulated from damage and healing under the named model; indicative until every class’s modifiers are in, so a tank without their stance and taunt multipliers can read below the damage dealers. Under a brushed window it is the fight’s total scaled by the window’s share.',
    },
    {
      term: 'Threat share and bar',
      meaning:
        'On the Threat tab, Share is a player’s part of every player’s threat in this window (or of every player’s threat on the picked enemy), whatever the Source scope shows, so one player on their own still reads their real share. The bar is drawn against the highest row, not against 100%. An enemy’s own row is outside the share.',
    },
    {
      term: 'Threat on a target',
      meaning:
        'The threat one player has built on one enemy, from the damage they did to it and their share of the raid’s healing while it was engaged. Once every class’s modifiers are in, whoever has the most is who it attacks; until then a tank without their stance and taunt multipliers can read below the damage dealers, so the order is not yet the aggro order.',
    },
    {
      term: 'Standing threat',
      meaning:
        'On the Threat tab under a window: the threat a player had built on that enemy by the window’s end, counted from the pull’s start. It is the number that decides who the enemy attacks, so the table sorts by it. Measured from the fight’s own seconds, never scaled.',
    },
    {
      term: 'Threat built in a window',
      meaning:
        'The threat a player made inside the window itself, the seconds of the window added up. A player can be top of the standing and bottom of the built: they came in with a lead.',
    },
    {
      term: 'Taunt',
      meaning:
        'A cast that forces an enemy onto the caster; the Threat tab lists every one, and the timeline marks them.',
    },
    {
      term: 'Mechanics table',
      meaning:
        'The per-boss list behind Mechanics mode: which abilities are avoidable, which casts to interrupt, which debuffs to dispel. A boss without one has no Mechanics page yet.',
    },
    {
      term: 'Avoidable damage',
      meaning:
        'Damage from an ability the boss’s mechanics table says a player should not have been standing in. The table is curated by a person; the log only says what hit whom.',
    },
    {
      term: 'Unavoidable',
      meaning:
        'Damage the fight deals regardless of where anyone stood. Listed on the mechanics page so a boss’s tank damage is not mistaken for a gap in its table.',
    },
    {
      term: 'Went through',
      meaning:
        'Casts of an interruptible spell that finished because nobody stopped them, out of the casts that started.',
    },
    {
      term: 'Ran their course',
      meaning:
        'Debuffs that expired on their own because nobody dispelled them, out of the times they landed.',
    },
    {
      term: 'Not yet classified',
      meaning:
        'An enemy ability that hit someone and the boss’s mechanics table does not list. Nobody has judged it avoidable or not, so the page names it rather than hiding it.',
    },
    {
      term: 'Difference (Compare)',
      meaning:
        'This fight’s figure less the compared one’s: a plus means this fight did more. Expanding a row gives the same three columns ability by ability, and an ability only one side used shows a dash on the other. Threat has no ability split, because the engine keeps it per player.',
    },
  ];

  let open = $state(false);
</script>

<details class="border-line-soft rounded-panel border" bind:open data-testid="glossary">
  <summary
    class="text-nav inline-flex min-h-11 cursor-pointer items-center px-3 text-[12px] font-bold tracking-[0.06em] uppercase select-none md:min-h-9"
  >
    {open ? 'Hide the glossary' : 'What do these words mean?'}
  </summary>
  <dl class="grid grid-cols-1 gap-x-6 gap-y-2 px-3 pb-3 text-[13px] md:grid-cols-2">
    {#each TERMS as entry (entry.term)}
      <div class="flex flex-col gap-0.5">
        <dt class="text-strong font-semibold">{entry.term}</dt>
        <dd class="text-muted m-0">{entry.meaning}</dd>
      </div>
    {/each}
  </dl>
</details>
