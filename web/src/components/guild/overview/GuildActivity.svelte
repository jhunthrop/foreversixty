<!-- web/src/components/guild/overview/GuildActivity.svelte -->
<!-- The guild's activity on the Overview (home B spec 2026-11-07 P3, P7, P8): the Your guild rows of
     GET /v1/feed (raid posts and guild sims), each line's time updating each minute. Your guild never
     floors with site events: with no rows, or a read that fails, the section is not drawn. -->
<script lang="ts">
  import activeBuild from '../../../data/active-build.json';
  import { minuteClock } from '../../../lib/ui/minute-clock.svelte';
  import { feedCopy, fetchFeed, whenText, type FeedRow } from '../../../lib/home/feed';
  import FeedSentence from '../../home/FeedSentence.svelte';
  import FeedRowUi from '../../ui/FeedRow.svelte';
  import Section from '../../ui/Section.svelte';

  let { now }: { now?: Date } = $props();

  const TITLE = 'Guild activity';
  const clock = minuteClock(() => now);
  let rows = $state<FeedRow[]>([]);

  $effect(() => {
    void fetchFeed({ filter: 'guild', signedIn: true })
      .then((page) => (rows = page.rows))
      .catch(() => (rows = []));
  });
</script>

{#if rows.length > 0}
  <Section panel title={TITLE} testid="guild-overview-activity">
    <div class="flex flex-col">
      {#each rows as row (row.id)}
        <FeedRowUi
          href={row.href}
          when={whenText(row, clock.now)}
          kind={feedCopy.tags[row.tag]}
          lead={row}
          build={activeBuild.build}
          leadSize={40}
          testid="guild-overview-activity-row"
        >
          <FeedSentence segments={row.segments} />
        </FeedRowUi>
      {/each}
    </div>
  </Section>
{/if}
