<!-- web/src/components/character/CharacterActivity.svelte -->
<!-- The character's own activity (character spec 2026-11-06 section 6, C18, C20): its rows of the
     site feed, newest first, five at most, each an icon, the sentence the API wrote, its age and
     its tag. A session adds its guild's rows (a guild sim). No rows, or a failed read, and the
     section is left out: never an empty box, never a made-up row. -->
<script lang="ts">
  import { feedCopy, whenText, type FeedRow } from '../../lib/home/feed';
  import { fetchActivity, standingCopy } from '../../lib/character/standing';
  import FeedSentence from '../home/FeedSentence.svelte';
  import FeedRowUi from '../ui/FeedRow.svelte';
  import ItemIcon from '../ui/ItemIcon.svelte';
  import Section from '../ui/Section.svelte';
  import ClassCrestRing from './ClassCrestRing.svelte';

  let {
    characterKey,
    signedIn,
    build,
    now,
  }: { characterKey: string; signedIn: boolean; build: string; now: Date } = $props();

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
  <Section title={standingCopy.activity} testid="character-activity">
    {#each rows as row (row.id)}
      <FeedRowUi
        href={row.href}
        when={whenText(row, now)}
        kind={feedCopy.tags[row.tag]}
        testid="character-activity-row"
      >
        {#snippet lead()}
          {#if row.crest !== null}
            <ClassCrestRing characterClass={row.crest} size={40} />
          {:else}
            <ItemIcon icon={row.icon ?? undefined} quality={row.quality ?? 1} {build} size={40} />
          {/if}
        {/snippet}
        <FeedSentence segments={row.segments} />
      </FeedRowUi>
    {/each}
  </Section>
{/if}
