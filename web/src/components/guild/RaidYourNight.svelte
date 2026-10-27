<!-- web/src/components/guild/RaidYourNight.svelte -->
<!-- Your night (attendance spec 3.3): on a done raid it takes the sign-up strip's slot for a member
     with an own character in the log or signed up. One row per own character, the header character
     first: who, the outcome sentence with the muted facts under it (deaths, the move since the last
     raid, Excused), and the loot (each award an ItemHover) or `No loot this raid.` A character still
     waiting for approval keeps the sign-up strip's muted row. Officers never get this panel. -->
<script lang="ts">
  import type { GuildRaidView } from '../../lib/guild/api';
  import { attendanceCopy } from '../../lib/guild/attendance-copy';
  import { yourNightFacts, yourNightRows, yourNightSentence } from '../../lib/guild/night';
  import { SEP, raidCopy } from '../../lib/guild/raid-copy';
  import { ROW_DIVIDER } from '../../lib/guild/raid-report-styles';
  import { classColorVar } from '../../lib/report/format';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import Figures from './Figures.svelte';
  import RaidLootItem from './RaidLootItem.svelte';

  let {
    raid,
    currentKey = null,
    locale,
  }: {
    raid: GuildRaidView;
    /** The header selector's current character, drawn first. */
    currentKey?: string | null;
    locale: string;
  } = $props();

  const copy = attendanceCopy.yourNight;
  const rows = $derived(yourNightRows(raid.attendance ?? [], currentKey));
  const waiting = $derived(raid.viewer.characters.filter((character) => !character.verified));
  const report = $derived(raid.report);

  const awardsOf = (characterKey: string) =>
    (report?.loot ?? []).flatMap((group) =>
      group.awards.filter((award) => award.character_key === characterKey),
    );
</script>

<section
  class="bg-raised border-line rounded-panel border p-4 shadow-[inset_0_1px_0_rgba(229,185,85,.35)] md:p-5"
  aria-labelledby="raid-your-night-heading"
  data-testid="raid-your-night"
>
  <h2 id="raid-your-night-heading" class="section-title text-[15px]">{copy.head}</h2>
  <ul class="mt-1">
    {#each rows as row (row.character_key)}
      {@const facts = yourNightFacts(row, report, raid.guild_timezone, locale)}
      {@const awards = awardsOf(row.character_key)}
      <li class={ROW_DIVIDER} data-testid="raid-your-night-row">
        <div
          class="grid grid-cols-1 gap-3 py-3 md:min-h-[68px] md:grid-cols-[minmax(180px,1fr)_minmax(0,2fr)_minmax(160px,auto)] md:items-center md:gap-4"
        >
          <div class="flex min-w-0 items-center gap-3">
            <ClassCrestRing characterClass={row.class} size={32} />
            <div class="min-w-0">
              <p class="font-display truncate text-[14px]" style={`color:${classColorVar(row.class)}`}>
                {row.name}
              </p>
              <p class="text-muted truncate text-[12px]">{row.spec}</p>
            </div>
          </div>
          <div class="min-w-0">
            <p class="text-text text-[14px]" data-testid="raid-your-night-sentence">
              {yourNightSentence(row)}
            </p>
            {#if facts.length > 0}
              <p class="text-muted text-[13px]" data-testid="raid-your-night-facts">
                {#each facts as fact, index (index)}{#if index > 0}{SEP}{/if}{#if fact.kind === 'move'}<Figures
                      text={fact.lead}
                    /><Figures text={fact.change} class={fact.rising ? 'text-kill' : ''} /><Figures
                      text={fact.tail}
                    />{:else}<Figures text={fact.text} />{/if}{/each}
              </p>
            {/if}
          </div>
          <div class="flex min-w-0 flex-col gap-1.5" data-testid="raid-your-night-loot">
            {#each awards as award (`${award.item_id}-${award.at}`)}
              <RaidLootItem {award} iconSize={20} nameSize="text-[13px]" />
            {:else}
              <p class="text-muted text-[13px]">{copy.noLoot}</p>
            {/each}
          </div>
        </div>
      </li>
    {/each}
    {#each waiting as character (character.character_key)}
      <li
        class="{ROW_DIVIDER} flex min-h-[68px] items-center gap-3 py-3"
        data-testid="raid-your-night-waiting"
      >
        <span class="opacity-60"><ClassCrestRing characterClass={character.class} size={32} /></span>
        <div class="min-w-0">
          <p class="font-display text-muted text-[14px]">{character.name}</p>
          <p class="text-muted text-[13px]">{raidCopy.waitingCharacter}</p>
        </div>
      </li>
    {/each}
  </ul>
</section>
