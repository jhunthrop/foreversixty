<!-- web/src/components/guild/report/ReportYourNight.svelte -->
<!-- Your night, for a member (attendance spec 3.3, web layout spec 2026-11-05 §7.2): one Row per
     own character, the header character first: who, the outcome sentence with the muted facts
     under it (deaths, the move since the last raid, Excused) and the loot (each award with its
     tooltip) or `No loot this raid.` A character still waiting for approval keeps its muted row.
     Officers never get this section: they read every row in the report. -->
<script lang="ts">
  import type { GuildRaidView } from '../../../lib/guild/api';
  import { attendanceCopy } from '../../../lib/guild/attendance-copy';
  import { yourNightFacts, yourNightRows, yourNightSentence } from '../../../lib/guild/night';
  import { SEP, raidCopy } from '../../../lib/guild/raid-copy';
  import { specName } from '../../../lib/guild/spec-name';
  import { classTone } from '../../../lib/ui/primitives';
  import Crest from '../../ui/Crest.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';
  import Figures from '../Figures.svelte';
  import RaidLootItem from '../RaidLootItem.svelte';

  let {
    raid,
    currentKey = null,
  }: {
    raid: GuildRaidView;
    /** The header selector's current character, drawn first. */
    currentKey?: string | null;
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

<Section title={copy.head} id="raid-your-night" panel testid="raid-your-night">
  <ul>
    {#each rows as row (row.character_key)}
      {@const facts = yourNightFacts(row, report, raid.guild_timezone)}
      {@const awards = awardsOf(row.character_key)}
      <li data-testid="raid-your-night-row">
        <Row title={row.name} titleTone={classTone(row.class)} open>
          {#snippet lead()}<Crest cls={row.class} size={32} />{/snippet}
          {#snippet subContent()}
            <span class="block">{specName(row.spec)}</span>
            <span class="text-text block text-[14px]" data-testid="raid-your-night-sentence"
              >{yourNightSentence(row)}</span
            >
            {#if facts.length > 0}
              <span class="block" data-testid="raid-your-night-facts"
                >{#each facts as fact, index (index)}{#if index > 0}{SEP}{/if}{#if fact.kind === 'move'}<Figures
                      text={fact.lead}
                    /><Figures text={fact.change} class={fact.rising ? 'text-kill' : ''} /><Figures
                      text={fact.tail}
                    />{:else}<Figures text={fact.text} />{/if}{/each}</span
              >
            {/if}
            <span class="mt-1 flex flex-col gap-1" data-testid="raid-your-night-loot">
              {#each awards as award (`${award.item_id}-${award.at}`)}
                <RaidLootItem {award} iconSize={20} nameSize="text-[13px]" />
              {:else}
                <span>{copy.noLoot}</span>
              {/each}
            </span>
          {/snippet}
        </Row>
      </li>
    {/each}
    {#each waiting as character (character.character_key)}
      <li data-testid="raid-your-night-waiting">
        <Row title={character.name} sub={raidCopy.waitingCharacter} open>
          {#snippet lead()}<span class="opacity-60"><Crest cls={character.class} size={32} /></span>{/snippet}
        </Row>
      </li>
    {/each}
  </ul>
</Section>
