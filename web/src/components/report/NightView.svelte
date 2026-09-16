<!-- web/src/components/report/NightView.svelte -->
<!-- The whole night: every boss pull folded into one row per player and one row per boss.
     This is the page a raid leader opens first -- who carried, who died, which boss ate
     the evening -- and each figure links back to the pull it came from. Trash is left out
     for the reason the rankings leave it out. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import { classColorVar, formatAmount, formatDuration, formatPercent } from '../../lib/report/format';
  import type { Night, NightPlayer } from '../../lib/report/night';
  import ClassIcon from './ClassIcon.svelte';

  let {
    night,
    loading,
    onSelect,
    onSelectPlayer,
  }: {
    night: Night;
    loading: boolean;
    onSelect: (index: number) => void;
    /** Narrows the page to one player, the way the summary's names do. */
    onSelectPlayer: (guid: string) => void;
  } = $props();

  let open = $state<string | null>(null);

  const perSecond = (value: number): string => formatAmount(Math.round(value));

  /** The headline per-second figure for a row: what their role is there to do. */
  function roleFigure(player: NightPlayer): { label: string; value: number } {
    if (player.role === 'healer') return { label: 'HPS', value: player.hps };
    if (player.role === 'tank') return { label: 'DTPS', value: player.dtps };
    return { label: 'DPS', value: player.dps };
  }
</script>

<section class="flex flex-col gap-6" data-testid="night-view">
  <p class="text-muted text-[13px]" data-testid="night-summary">
    <span class="tabular font-mono">{night.expected}</span> boss pulls ·
    <span class="tabular font-mono">{night.bosses.reduce((sum, boss) => sum + boss.kills, 0)}</span> kills ·
    <span class="tabular font-mono">{night.bosses.reduce((sum, boss) => sum + boss.wipes, 0)}</span> wipes ·
    <span class="tabular font-mono">{formatDuration(night.time_ms)}</span> in combat ·
    <span class="text-death tabular font-mono">{night.deaths}</span> deaths
    {#if loading}
      · <span data-testid="night-loading">loading {night.fights} of {night.expected}</span>
    {/if}
  </p>

  <div class="flex flex-col gap-2">
    <h2 class="label text-muted">Bosses</h2>
    <div
      class="text-muted label hidden grid-cols-[minmax(160px,2fr)_64px_64px_64px_80px_64px_120px] gap-x-3 px-2 pb-1 md:grid"
    >
      <span>Boss</span>
      <span class="text-right">Pulls</span>
      <span class="text-right">Kills</span>
      <span class="text-right">Wipes</span>
      <span class="text-right" title="Time spent on this boss across every pull">Time</span>
      <span class="text-right">Deaths</span>
      <span class="text-right" title="The quickest kill; click to open it">Best kill</span>
    </div>
    <ul class="flex flex-col" data-testid="night-bosses">
      {#each night.bosses as boss (boss.name)}
        <li
          class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px] md:grid-cols-[minmax(160px,2fr)_64px_64px_64px_80px_64px_120px]"
          data-testid={`night-boss-${boss.name}`}
        >
          <span class="flex min-w-0 flex-col">
            <span class="text-strong font-semibold">{boss.name}</span>
            <span class="flex flex-wrap gap-1">
              {#each boss.fights as index, i (index)}
                <button
                  type="button"
                  class="tabular inline-flex min-h-11 items-center font-mono text-[11px] underline-offset-2 hover:underline md:min-h-6"
                  title="Open this pull"
                  onclick={() => onSelect(index)}
                >
                  pull {i + 1}
                </button>
              {/each}
            </span>
          </span>
          <span class="tabular text-right font-mono md:text-[14px]"
            >{boss.pulls}<span class="label font-body text-muted ml-1 md:hidden">pulls</span></span
          >
          <span class="text-muted label col-span-2 flex flex-wrap gap-x-3 gap-y-1 md:hidden">
            <span
              >Kills <span class="tabular font-mono" class:text-kill={boss.kills > 0}>{boss.kills}</span
              ></span
            >
            <span
              >Wipes <span class="tabular font-mono" class:text-wipe={boss.wipes > 0}>{boss.wipes}</span
              ></span
            >
            <span>Time <span class="tabular font-mono">{formatDuration(boss.time_ms)}</span></span>
            <span
              >Deaths <span class="tabular font-mono" class:text-death={boss.deaths > 0}>{boss.deaths}</span
              ></span
            >
            {#if boss.best}
              <span
                >Best kill <span class="text-kill tabular font-mono"
                  >{formatDuration(boss.best.duration_ms)}</span
                ></span
              >
            {:else if boss.lowest_wipe_pct !== undefined}
              <span
                >Best wipe <span class="text-wipe tabular font-mono">{Math.round(boss.lowest_wipe_pct)}%</span
                ></span
              >
            {/if}
          </span>
          <span class="tabular hidden text-right font-mono md:inline" class:text-kill={boss.kills > 0}
            >{boss.kills}</span
          >
          <span class="tabular hidden text-right font-mono md:inline" class:text-wipe={boss.wipes > 0}
            >{boss.wipes}</span
          >
          <span class="text-muted tabular hidden text-right font-mono md:inline"
            >{formatDuration(boss.time_ms)}</span
          >
          <span class="tabular hidden text-right font-mono md:inline" class:text-death={boss.deaths > 0}
            >{boss.deaths}</span
          >
          <span class="hidden text-right md:inline">
            {#if boss.best}
              <button
                type="button"
                class="text-kill tabular inline-flex min-h-11 items-center font-mono underline-offset-2 hover:underline md:min-h-6"
                onclick={() => onSelect(boss.best?.index ?? 0)}
              >
                {formatDuration(boss.best.duration_ms)}
              </button>
            {:else if boss.lowest_wipe_pct !== undefined}
              <span class="text-wipe" title="The lowest the boss was brought to"
                >best wipe {Math.round(boss.lowest_wipe_pct)}%</span
              >
            {:else}
              <span class="text-wipe">no kill</span>
            {/if}
          </span>
        </li>
      {/each}
    </ul>
  </div>

  <div class="flex flex-col gap-2">
    <h2 class="label text-muted">Players, across every boss pull</h2>
    <div
      class="text-muted label hidden grid-cols-[minmax(120px,1.4fr)_88px_56px_96px_96px_96px_72px_56px] gap-x-3 px-2 pb-1 md:grid"
    >
      <span>Name</span>
      <span>Spec</span>
      <span class="text-right" title="Boss pulls this player was in">Pulls</span>
      <span class="text-right" title="Damage over every pull, and per second of their own combat time"
        >Damage</span
      >
      <span class="text-right" title="Healing over every pull, and per second of their own combat time"
        >Healing</span
      >
      <span class="text-right" title="Damage taken over every pull, and per second">Taken</span>
      <span class="text-right" title="Share of their combat time spent casting or attacking">Active</span>
      <span class="text-right">Deaths</span>
    </div>
    <ul class="flex flex-col" data-testid="night-players">
      {#each night.players as player (player.guid)}
        {@const display = splitUnitName(player.name)}
        {@const headline = roleFigure(player)}
        <li class="border-line-soft border-b" data-testid={`night-player-${player.guid}`}>
          <button
            type="button"
            class="grid min-h-11 w-full grid-cols-[minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 px-2 py-2 text-left text-[14px] md:grid-cols-[minmax(120px,1.4fr)_88px_56px_96px_96px_96px_72px_56px]"
            aria-expanded={open === player.guid}
            onclick={() => (open = open === player.guid ? null : player.guid)}
          >
            <span
              class="flex min-w-0 items-center gap-2 truncate font-semibold"
              style={`color: ${classColorVar(player.class)}`}
              ><ClassIcon className={player.class} /><span
                class="mr-1 inline-block w-3 transition-transform"
                class:rotate-90={open === player.guid}
                aria-hidden="true">›</span
              ><span
                role="link"
                tabindex="0"
                class="underline-offset-2 hover:underline"
                title="Show only this player"
                data-testid="night-player-name"
                onclick={(event) => {
                  event.stopPropagation();
                  onSelectPlayer(player.guid);
                }}
                onkeydown={(event) => {
                  if (event.key === 'Enter') {
                    event.stopPropagation();
                    onSelectPlayer(player.guid);
                  }
                }}>{display.name}</span
              ></span
            >
            <span class="text-muted text-[13px]"
              >{player.spec ?? player.class ?? 'Unknown'}
              <span class="label ml-2 text-[10px]">{open === player.guid ? 'hide pulls' : 'per pull ›'}</span
              ></span
            >
            <span class="tabular hidden text-right font-mono md:inline">{player.fights}</span>
            <span class="tabular hidden flex-col text-right font-mono leading-tight md:flex">
              <span>{formatAmount(player.damage_done)}</span>
              <span class="text-muted text-[11px]">{perSecond(player.dps)}/s</span>
            </span>
            <span class="tabular hidden flex-col text-right font-mono leading-tight md:flex">
              <span>{formatAmount(player.healing_done)}</span>
              <span class="text-muted text-[11px]">{perSecond(player.hps)}/s</span>
            </span>
            <span class="tabular hidden flex-col text-right font-mono leading-tight md:flex">
              <span>{formatAmount(player.damage_taken)}</span>
              <span class="text-muted text-[11px]">{perSecond(player.dtps)}/s</span>
            </span>
            <span class="text-muted tabular hidden text-right font-mono text-[13px] md:inline"
              >{formatPercent(player.activity_pct)}</span
            >
            <span class="tabular hidden text-right font-mono md:inline" class:text-death={player.deaths > 0}
              >{player.deaths}</span
            >
            <span class="text-muted label col-span-2 flex flex-wrap gap-x-3 md:hidden">
              <span>{headline.label} <span class="tabular font-mono">{perSecond(headline.value)}</span></span>
              <span>Pulls <span class="tabular font-mono">{player.fights}</span></span>
              <span
                >Deaths <span class="tabular font-mono" class:text-death={player.deaths > 0}
                  >{player.deaths}</span
                ></span
              >
            </span>
          </button>
          {#if open === player.guid}
            <table class="bg-card-top w-full text-[13px]" data-testid="night-player-fights">
              <caption class="label text-muted px-2 pt-2 text-left">Each pull</caption>
              <tbody>
                {#each player.by_fight as entry (entry.index)}
                  <tr class="border-line-soft border-b">
                    <td class="px-2 py-1">
                      <button
                        type="button"
                        class="inline-flex min-h-11 items-center underline-offset-2 hover:underline md:min-h-6"
                        onclick={() => onSelect(entry.index)}
                        title="Open this pull"
                      >
                        {entry.boss}{#if entry.pull}
                          <span class="text-muted ml-1 text-[12px]">· {entry.pull}</span>{/if}
                      </button>
                    </td>
                    <td class="py-1 pr-3 font-semibold {entry.kill ? 'text-kill' : 'text-wipe'}"
                      >{entry.kill ? 'Kill' : 'Wipe'}</td
                    >
                    <td class="text-muted tabular py-1 pr-3 text-right font-mono"
                      >{formatDuration(entry.duration_ms)}</td
                    >
                    <td class="tabular py-1 pr-3 text-right font-mono">{perSecond(entry.dps)} dps</td>
                    <td class="tabular py-1 pr-3 text-right font-mono">{perSecond(entry.hps)} hps</td>
                    <td class="tabular py-1 pr-3 text-right font-mono" title="Damage taken on this pull"
                      >{perSecond(
                        entry.duration_ms > 0 ? entry.damage_taken / (entry.duration_ms / 1000) : 0,
                      )}/s taken</td
                    >
                    <td class="tabular py-1 pr-3 text-right font-mono" class:text-death={entry.deaths > 0}
                      >{entry.deaths} {entry.deaths === 1 ? 'death' : 'deaths'}</td
                    >
                  </tr>
                {/each}
              </tbody>
            </table>
          {/if}
        </li>
      {/each}
    </ul>
  </div>
</section>
