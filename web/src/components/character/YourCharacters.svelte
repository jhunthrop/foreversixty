<!-- web/src/components/character/YourCharacters.svelte -->
<!-- /characters' rail (round 5, directions mock b2-characters): the signed-in account's characters,
     each its ringed crest, its name in the class colour and "Level 60 Fury Warrior", linked to its
     page. Signed out, or with no characters, nothing draws: never an empty panel. -->
<script lang="ts">
  import { fetchMeOnce } from '../../lib/account/api';
  import { yourCharacterRows, type YourCharacterRow } from '../../lib/character/your-characters';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import Crest from '../ui/Crest.svelte';

  let { title }: { title: string } = $props();

  let rows = $state<YourCharacterRow[]>([]);
  $effect(() => {
    void fetchMeOnce()
      .then((me) => {
        rows = yourCharacterRows(me?.characters ?? []);
      })
      .catch(() => {
        rows = [];
      });
  });
</script>

{#if rows.length > 0}
  <Section {title} panel testid="your-characters">
    {#each rows as row (row.key)}
      <Row title={row.name} titleTone={row.tone} sub={row.line} href={row.href} density="slim" open>
        {#snippet lead()}<Crest cls={row.classSlug} size={32} />{/snippet}
      </Row>
    {/each}
  </Section>
{/if}
