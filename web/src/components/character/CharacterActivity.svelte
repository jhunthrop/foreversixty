<!-- web/src/components/character/CharacterActivity.svelte -->
<!-- The character's own activity (character spec 2026-11-06 section 6, C18, C20), in the page's
     rail: its rows of the site feed, newest first, five at most, each an icon, the sentence the API
     wrote with the character's own name taken out (ownActivitySegments), its age. A session adds its guild's rows (a guild sim). No rows, or a failed read, and the
     section is left out: never an empty box, never a made-up row. -->
<script lang="ts">
  import { whenText, type FeedRow } from '../../lib/home/feed';
  import { fetchActivity, ownActivitySegments, standingCopy } from '../../lib/character/standing';
  import FeedSentence from '../home/FeedSentence.svelte';
  import FeedRowUi from '../ui/FeedRow.svelte';
  import Section from '../ui/Section.svelte';

  let {
    characterKey,
    characterName,
    signedIn,
    build,
    now,
  }: { characterKey: string; characterName: string; signedIn: boolean; build: string; now: Date } = $props();

  let rows = $state<FeedRow[]>([]);
  $effect(() => {
    const key = characterKey;
    rows = [];
    fetchActivity(key, signedIn)
      .then((page) => {
        if (key === characterKey) rows = page.rows;
      })
      .catch(() => {
        rows = [];
      });
  });
</script>

{#if rows.length > 0}
  <Section title={standingCopy.activity} panel testid="character-activity">
    {#each rows as row (row.id)}
      <FeedRowUi href={row.href} when={whenText(row, now)} lead={row} {build} testid="character-activity-row">
        <FeedSentence segments={ownActivitySegments(row.segments, characterName)} />
      </FeedRowUi>
    {/each}
  </Section>
{/if}
