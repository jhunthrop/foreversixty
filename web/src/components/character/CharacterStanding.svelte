<!-- web/src/components/character/CharacterStanding.svelte -->
<!-- The owner's standing (character spec 2026-11-06 section 6, C16, C19, C21), a rail panel of link
     rows since the h-audit review (2026-11-09; a bare link strip under the band read as loose
     nudges): the place in the guild's sim ladder, the spec's tier as a link to the tier
     list, and the next guild raid as a short clock to its raid page. Each piece shows only when it
     has something true to say; with none, the panel is left out. Home keeps the ladder, the tile and
     the countdown panels; this only links to them. -->
<script lang="ts">
  import type { Me, MeCharacter } from '../../lib/account/api';
  import { fetchGuildHome, type NextRaidView } from '../../lib/guild/api';
  import { fetchHomeAnswer, type HomeAnswer } from '../../lib/home/home-answer';
  import { guildLinks } from '../../lib/home/home-view';
  import { gradeLine, guildRankLine, raidLine, standingCopy } from '../../lib/character/standing';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';

  let { me, own, specName, now }: { me: Me; own: MeCharacter; specName: string | null; now: Date } = $props();

  let answer = $state<HomeAnswer | null>(null);
  let raid = $state<NextRaidView | null>(null);
  $effect(() => {
    const key = own.key;
    answer = null;
    raid = null;
    fetchHomeAnswer(key)
      .then((read) => {
        if (key === own.key) answer = read;
      })
      .catch(() => {
        answer = null;
      });
    const guildId = own.guild?.id;
    if (guildId !== undefined)
      fetchGuildHome(guildId)
        .then((home) => {
          if (key === own.key) raid = home.next_raid ?? null;
        })
        .catch(() => {
          raid = null;
        });
  });

  const links = $derived(guildLinks(me, own.guild?.id ?? -1, raid?.id ?? null));
  const pieces = $derived(
    [
      { key: 'rank', text: guildRankLine(answer?.guild_sims ?? null), href: links.guild ?? '/guild' },
      { key: 'grade', text: gradeLine(answer?.spec_week ?? null, specName), href: '/tiers' },
      { key: 'raid', text: raidLine(raid, now), href: links.raid ?? links.guild ?? '/guild' },
    ].filter((piece): piece is { key: string; text: string; href: string } => piece.text !== null),
  );
</script>

{#if pieces.length > 0}
  <Section title={standingCopy.standing} panel testid="character-standing">
    {#each pieces as piece (piece.key)}
      <Row
        title={piece.text}
        href={piece.href}
        density="slim"
        open
        testid={`character-standing-${piece.key}`}
      />
    {/each}
  </Section>
{/if}
