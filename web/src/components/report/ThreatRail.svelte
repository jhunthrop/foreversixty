<!-- web/src/components/report/ThreatRail.svelte -->
<!-- The Threat tab's rail (logs spec 2026-11-07 section 6.2): Section `Taunts`, one open Row per
     taunt in pull order: the time on the pull's clock and the spell's icon, `{name} taunted
     {enemy}` with the taunter in their class colour, the spell under it, and `Around it`, which
     sets the window to the span around the taunt (not on the night, which has no one clock). With
     an enemy picked, its taunts alone: the list answers "who took this one off me". -->
<script lang="ts">
  import CharacterName from '../ui/CharacterName.svelte';
  import EmptyState from '../ui/EmptyState.svelte';
  import { splitUnitName } from '../../lib/characters';
  import { tauntKey } from '../../lib/report/format';
  import { SpellIcons } from '../../lib/report/spell-icons.svelte';
  import { noTauntsLine, orderedTaunts, pickedEnemyName } from '../../lib/report/threat-taunts';
  import type { Taunt, ThreatPair } from '../../lib/report/types';
  import { aroundWindow, type TimeWindow } from '../../lib/report/window';
  import { clockDuration } from '../../lib/format/duration';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import SpellIcon from '../ui/SpellIcon.svelte';

  let {
    taunts = undefined,
    enemy = undefined,
    target = '',
    pairs = [],
    names = new Map<string, string>(),
    classOf = new Map<string, string>(),
    sourceName = undefined,
    approximate = false,
    scopeNoun = 'pull',
    durationMs,
    onWindow,
  }: {
    /** Undefined means the report was parsed before taunts were kept; [] means none. */
    taunts?: readonly Taunt[];
    /** The picked enemy's name, when the caller already knows it (the table does). */
    enemy?: string;
    /** Else the url's `target` and the pairs that name it (the frame's rail). */
    target?: string;
    pairs?: readonly ThreatPair[];
    names?: ReadonlyMap<string, string>;
    classOf?: ReadonlyMap<string, string>;
    /** The one player the source scope narrows to, by name. */
    sourceName?: string;
    /** True under a brushed window. */
    approximate?: boolean;
    scopeNoun?: 'pull' | 'night';
    /** The whole fight's length, which a taunt's window is clamped against. */
    durationMs: number;
    onWindow: (window: TimeWindow | null) => void;
  } = $props();

  const icons = new SpellIcons();
  const enemyName = $derived(enemy ?? pickedEnemyName(pairs, target));
  const rows = $derived(orderedTaunts(taunts ?? [], enemyName));
  const empty = $derived(noTauntsLine({ sourceName, enemy: enemyName, windowed: approximate, scopeNoun }));
  /** The units table spells a name; the event's own copy of it is the fallback. */
  const nameOf = (guid: string, recorded: string): string => splitUnitName(names.get(guid) ?? recorded).name;
  const subOf = (taunt: Taunt): string =>
    [
      taunt.spell_name,
      taunt.pre_pull === true ? 'cast before the pull, so Casts does not count it' : '',
      taunt.label ?? '',
    ]
      .filter((part) => part !== '')
      .join(' · ');
</script>

<Section title="Taunts" testid="threat-taunts" panel>
  {#if taunts === undefined}
    <p class="text-small text-muted">This report was parsed before taunts were kept; parse it again.</p>
  {:else if rows.length === 0}
    <EmptyState message={empty} testid="threat-taunts-empty" />
  {:else}
    <div>
      {#each rows as taunt (tauntKey(taunt))}
        <Row title="" sub={subOf(taunt)} open density="slim" testid="threat-taunt">
          {#snippet lead()}
            <span class="text-muted w-9 text-right text-[13px] tabular-nums" title="On the pull's clock"
              >{clockDuration(taunt.at_ms)}</span
            >
            <SpellIcon
              socket
              url={icons.of(taunt.spell_id, taunt.spell_name)}
              size={24}
              tip={taunt.spell_id}
            />
          {/snippet}
          {#snippet titleContent()}
            <CharacterName
              name={nameOf(taunt.source_guid, taunt.source_name)}
              cls={classOf.get(taunt.source_guid)}
              weight="inherit"
            />
            <span class="text-text font-semibold">taunted {nameOf(taunt.target_guid, taunt.target_name)}</span
            >
          {/snippet}
          {#snippet trail()}
            {#if scopeNoun !== 'night'}
              <!-- A quiet link, not a button: the taunt's sentence keeps the rail's width. -->
              <button
                type="button"
                class="text-gold shrink-0 text-[13px] font-semibold whitespace-nowrap hover:underline [@media(pointer:coarse)]:-my-3 [@media(pointer:coarse)]:inline-flex [@media(pointer:coarse)]:min-h-11 [@media(pointer:coarse)]:items-center"
                data-testid="taunt-window"
                aria-label="Set the window to the span around this taunt"
                onclick={() => onWindow(aroundWindow(taunt.at_ms, durationMs))}>Around it ›</button
              >
            {/if}
          {/snippet}
        </Row>
      {/each}
    </div>
  {/if}
</Section>
