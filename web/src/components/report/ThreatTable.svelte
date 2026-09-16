<!-- web/src/components/report/ThreatTable.svelte -->
<!-- Threat, with the model's own honesty on the page. The engine's BaseThreat reports
     itself incomplete until the per-class table lands, and a threat number nobody can
     trust is worse than one labelled untrustworthy.

     `row.threat` is a second, independent lie: window.ts's scopeThreat has no threat
     model in the browser to recompute from, so it scales each row by the same ratio the
     actor rows' damage-plus-healing split moved by. That is a `~`, the same mark
     ActorRow's per-ability and per-target splits carry, not the model-incompleteness note
     above -- the two are unrelated and both can apply to the same row at once. The
     per-target rows are scaled the same way and carry the same mark.

     The picker turns one question into two. "Every enemy" is the totals table: every
     unit's threat over the window, enemies folded by name. Pick an enemy and the table
     becomes that enemy's own ranking -- who it is looking at, and by how much -- which is
     the question a tank actually asks. It rides in the url's `target`, the same key the
     damage tabs' target filter uses, so one link carries one target. Over a whole night
     that target is the enemy's NAME rather than a GUID -- an add is a new GUID on every
     pull -- which is why url.ts reads `target` as any printable string and not as a GUID
     pattern. -->
<script lang="ts">
  import { splitUnitName } from '../../lib/characters';
  import {
    approximateAriaLabel,
    approximateMark,
    approximateTitle,
    classColorVar,
    formatAmount,
    formatDuration,
    formatPercent,
  } from '../../lib/report/format';
  import type { Taunt, ThreatPair, ThreatRow } from '../../lib/report/types';
  import { aroundWindow, type TimeWindow } from '../../lib/report/window';
  import CopyCsv from './CopyCsv.svelte';

  let {
    rows,
    pairs = [],
    everyonePairs = undefined,
    taunts = undefined,
    names = new Map<string, string>(),
    classOf = new Map<string, string>(),
    approximate = false,
    totalThreat = undefined,
    target = '',
    startMs,
    durationMs,
    onPatch,
    onWindow,
  }: {
    rows: ThreatRow[];
    /** One player's threat on one enemy. Absent from summaries before engine 0.3.1. */
    pairs?: ThreatPair[];
    /** The same pairs before the source scope, so a scoped row still shares against everyone. */
    everyonePairs?: ThreatPair[];
    /** Undefined means the report was parsed before taunts were kept; [] means none. */
    taunts?: Taunt[];
    /** GUID to unit name, the units table being the canonical spelling of a name. */
    names?: ReadonlyMap<string, string>;
    /** The whole window's threat, so a source scope's rows still share against everyone. */
    totalThreat?: number;
    classOf?: Map<string, string>;
    /** True when the window is brushed, so threat is a scaled share, not measured. */
    approximate?: boolean;
    /** The picked enemy, from the url's `target`; '' is every enemy. */
    target?: string;
    /** The window's start, so every time on this tab is read from the same zero. */
    startMs: number;
    /** The whole fight's length, which a taunt's window link is clamped against. */
    durationMs: number;
    onPatch: (patch: { target?: string }) => void;
    onWindow: (window: TimeWindow) => void;
  } = $props();

  /** A row of the table, whichever question it is answering. */
  interface ThreatLine {
    guid: string;
    name: string;
    threat: number;
  }

  /** One enemy, by name, and the players who built threat on it. */
  interface TargetGroup {
    /** The url value: the group's first GUID, which over a night is the enemy's name. */
    id: string;
    name: string;
    /** Every GUID folded under this name, so a link to any one of them picks the group. */
    guids: string[];
    lines: ThreatLine[];
    total: number;
  }

  /** The table as lines: each row's threat and its share of whatever the table totals. */
  function csvLines(): string[][] {
    return [
      picked === undefined
        ? ['Unit', 'Threat', 'Share %']
        : ['Player', `Threat on ${picked.name}`, 'Share %'],
      ...lines.map((line) => [
        splitUnitName(line.name).name,
        String(Math.round(line.threat)),
        (total === 0 ? 0 : (line.threat / total) * 100).toFixed(1),
      ]),
    ];
  }
  /** The environment (a fall, a fire) has the null GUID and holds no threat. */
  const NULL_GUID = /^0+$/;
  /**
   * One row per enemy name: six "General Kaal" units are one boss to the reader, the way
   * Damage Taken folds them; players keep their own rows.
   */
  const ordered = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const byName = new Map<string, ThreatRow>();
    const out: ThreatRow[] = [];
    for (const row of rows) {
      if (NULL_GUID.test(row.guid)) continue;
      if (classOf.has(row.guid)) {
        out.push(row);
        continue;
      }
      const key = splitUnitName(row.name).name;
      const found = byName.get(key);
      if (found === undefined) {
        const merged = { ...row, name: key };
        byName.set(key, merged);
        out.push(merged);
      } else {
        const merged = {
          ...found,
          threat: found.threat + row.threat,
          complete: found.complete && row.complete,
        };
        byName.set(key, merged);
        out[out.indexOf(found)] = merged;
      }
    }
    return out.sort((a, b) => b.threat - a.threat);
  });
  /**
   * The picker's options: one per enemy name, biggest first, each folding every unit of
   * that name -- an add is a new GUID on every pull, and six of them are one enemy to the
   * reader. A player's threat on all of them adds up into one row.
   */
  function groupPairs(list: ThreatPair[]): TargetGroup[] {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const byName = new Map<string, { guids: string[]; players: Map<string, ThreatLine> }>();
    for (const pair of list) {
      const key = splitUnitName(pair.target_name).name;
      let found = byName.get(key);
      if (found === undefined) {
        found = { guids: [], players: new Map<string, ThreatLine>() };
        byName.set(key, found);
      }
      if (!found.guids.includes(pair.target_guid)) found.guids.push(pair.target_guid);
      const have = found.players.get(pair.guid);
      found.players.set(pair.guid, {
        guid: pair.guid,
        name: splitUnitName(pair.name).name,
        threat: (have?.threat ?? 0) + pair.threat,
      });
    }
    const out: TargetGroup[] = [...byName.entries()].map(([name, group]) => {
      const lines = [...group.players.values()].sort((a, b) => b.threat - a.threat);
      return {
        id: group.guids[0],
        name,
        guids: group.guids,
        lines,
        total: lines.reduce((sum, line) => sum + line.threat, 0),
      };
    });
    return out.sort((a, b) => b.total - a.total);
  }
  const groups = $derived(groupPairs(pairs));
  /** The url's target names one of this table's enemies, or it names none of them. */
  const picked = $derived(target === '' ? undefined : groups.find((group) => group.guids.includes(target)));
  /**
   * The picked enemy's threat from every player, not only the ones in scope -- the same
   * job `totalThreat` does for the totals table, so narrowing the source to one player
   * shows their real share of what the boss is looking at rather than 100% of themselves.
   */
  const pickedTotal = $derived(
    picked === undefined
      ? 0
      : (groupPairs(everyonePairs ?? pairs).find((group) => group.name === picked.name)?.total ??
          picked.total),
  );
  const lines = $derived<ThreatLine[]>(picked?.lines ?? ordered);
  /** Six units named "General Kaal" are six rows; each after the first says which copy it is. */
  const copyOf = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const seen = new Map<string, number>();
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const out = new Map<string, number>();
    for (const row of ordered) {
      const n = (seen.get(row.name) ?? 0) + 1;
      seen.set(row.name, n);
      out.set(row.guid, n);
    }
    return out;
  });
  const copies = $derived(new Set([...copyOf.entries()].filter(([, n]) => n > 1).map(([guid]) => guid)));
  const peak = $derived(lines.reduce((highest, line) => Math.max(highest, line.threat), 0));
  // Both branches share against everyone, never against the rows in scope: a picked enemy
  // against every player's threat on it, the totals table against the whole window.
  const total = $derived(
    picked !== undefined ? pickedTotal : (totalThreat ?? lines.reduce((sum, line) => sum + line.threat, 0)),
  );
  const incomplete = $derived(ordered.some((row) => !row.complete));
  const modelVersion = $derived(ordered[0]?.model_version ?? '');
  const mark = $derived(approximateMark(approximate));
  const title = $derived(approximateTitle(approximate));
  const orderedTaunts = $derived([...(taunts ?? [])].sort((a, b) => a.at_ms - b.at_ms));
  /**
   * Everything that tells one taunt from another: two tanks can taunt two adds on the
   * same millisecond, and over a night the same instant recurs in every pull, which the
   * label separates.
   */
  const tauntKey = (taunt: Taunt): string =>
    `${taunt.label ?? ''}|${taunt.at_ms}|${taunt.source_guid}|${taunt.target_guid}|${taunt.spell_id}`;
  /** The units table spells a name; the event's own copy of it is the fallback. */
  const nameOf = (guid: string, recorded: string): string => splitUnitName(names.get(guid) ?? recorded).name;
  const selectClass =
    'border-line-warm bg-raised rounded-control text-text h-11 w-full max-w-full min-w-0 px-2 text-[13px] md:h-9 md:w-auto';
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
    {#if groups.length > 0}
      <label class="label text-muted flex flex-wrap items-center gap-2" for="threat-target">
        Enemy
        <select
          id="threat-target"
          class={selectClass}
          data-testid="threat-target"
          value={picked?.id ?? ''}
          onchange={(event) => onPatch({ target: (event.currentTarget as HTMLSelectElement).value })}
        >
          <option value="">Every enemy</option>
          {#each groups as group (group.id)}
            <option value={group.id}
              >{group.name}{group.guids.length > 1 ? ` ×${group.guids.length}` : ''}</option
            >
          {/each}
        </select>
      </label>
    {/if}
    <div
      class="text-muted label hidden grid-cols-[minmax(120px,1.2fr)_minmax(0,3fr)_96px_64px] gap-x-3 px-2 pb-1 md:grid"
    >
      <span>{picked === undefined ? 'Unit' : 'Player'}</span>
      <span title="Threat generated, against the highest row">Threat</span>
      <span
        class="text-right"
        title={picked === undefined
          ? 'Threat accumulated from damage and healing over this window'
          : `Threat this player built on ${picked.name} over this window`}>Total</span
      >
      <span
        class="text-right"
        title={picked === undefined
          ? 'Share of all the threat in this table'
          : `Share of every player's threat on ${picked.name}`}>Share</span
      >
    </div>
    <ul class="flex flex-col" data-testid={picked === undefined ? undefined : 'threat-on-target'}>
      {#each lines as row (row.guid)}
        <li
          class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px] md:grid-cols-[minmax(120px,1.2fr)_minmax(0,3fr)_96px_64px]"
          data-testid={`threat-${row.guid}`}
        >
          <span class="truncate font-semibold" style={`color: ${classColorVar(classOf.get(row.guid))}`}>
            {splitUnitName(row.name).name}{#if picked === undefined && copies.has(row.guid)}
              <span class="text-muted ml-1 font-mono text-[11px]" title="Another unit of the same name"
                >#{copyOf.get(row.guid)}</span
              >{/if}
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
          <span
            class="text-muted tabular col-span-2 text-right font-mono text-[13px] md:col-span-1"
            data-testid="threat-share">{formatPercent(total === 0 ? 0 : (row.threat / total) * 100)}</span
          >
        </li>
      {/each}
    </ul>
    <CopyCsv lines={csvLines} />
    {#if approximate}
      <p class="text-muted text-[12px]" data-testid="threat-approximate-note">
        Threat is marked {mark} because it is accumulated from damage and healing and scaled to this window's share
        of that total, not recomputed from the model directly.
      </p>
    {/if}
    <div class="flex flex-col gap-1 pt-2" data-testid="threat-taunts">
      <h3 class="label text-muted">Taunts</h3>
      {#if taunts === undefined}
        <p class="text-muted text-[13px]">This report was parsed before taunts were kept; parse it again.</p>
      {:else if orderedTaunts.length === 0}
        <p class="text-muted text-[13px]">No taunts in this window.</p>
      {:else}
        <ul class="flex flex-col">
          {#each orderedTaunts as taunt (tauntKey(taunt))}
            <li
              class="border-line-soft flex min-h-11 flex-wrap items-center gap-x-2 gap-y-1 border-b px-2 py-2 text-[14px]"
            >
              <span class="text-muted tabular font-mono text-[13px]"
                >{formatDuration(taunt.at_ms - startMs)}</span
              >
              <span aria-hidden="true" class="text-muted">·</span>
              <span>
                <span class="font-semibold" style={`color: ${classColorVar(classOf.get(taunt.source_guid))}`}
                  >{nameOf(taunt.source_guid, taunt.source_name)}</span
                >
                taunted {nameOf(taunt.target_guid, taunt.target_name)}
                <span class="text-muted">({taunt.spell_name})</span>
              </span>
              {#if taunt.label !== undefined}
                <span class="text-muted text-[12px]">{taunt.label}</span>
              {/if}
              <button
                type="button"
                class="text-nav ml-auto inline-flex min-h-11 items-center text-[12px] font-bold tracking-[0.06em] uppercase md:min-h-0"
                data-testid="taunt-window"
                title="Set the window to the ten seconds around this taunt"
                onclick={() => onWindow(aroundWindow(taunt.at_ms, durationMs))}>10s around it</button
              >
            </li>
          {/each}
        </ul>
      {/if}
    </div>
  </div>
{/if}
