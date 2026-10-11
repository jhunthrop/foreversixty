<!-- web/src/components/CharacterInGame.svelte -->
<!-- The character page's "In game" section (Task L8.2 ruling 5): three lines of reference.
     What the addon shows for this character in game (the same score and tier the page leads
     with); how to put the character into the planner; and, for the signed-in owner only,
     when the site last saw its addon export. Plain data in: Character.svelte reads the score
     off the character answer and the export time off /v1/me (lib/score/score.ts's
     lastAddonExportAt), so this renders the same under svelte/server as in the page. -->
<script lang="ts">
  import { scoreCopy } from '../lib/score/copy';
  import { hasScore, lastSeenText, tierLabel, type ForeverScore } from '../lib/score/score';
  import ScoreBadge from './ScoreBadge.svelte';

  let {
    forever,
    lastSeenAt = null,
  }: {
    forever: ForeverScore | null;
    /** `build.captured_at` of the viewer's own addon export; null for anyone else. */
    lastSeenAt?: string | null;
  } = $props();

  const lastSeen = $derived(lastSeenAt === null ? null : lastSeenText(lastSeenAt));
</script>

<section class="flex flex-col gap-2" data-testid="character-in-game">
  <h2 class="section-title text-[18px]">{scoreCopy.inGameHeading}</h2>
  <ul class="flex flex-col gap-2 text-[14px] leading-relaxed">
    <li class="flex flex-wrap items-baseline gap-x-2" data-testid="character-in-game-score">
      {#if hasScore(forever)}
        {@const tier = tierLabel(forever.tier)}
        <span>{scoreCopy.inGameScoreLead}</span>
        <ScoreBadge score={forever.score} tier={forever.tier} />
        {#if tier !== null}<span class="font-semibold">{tier}</span>{/if}
      {:else}
        <span class="text-muted">{scoreCopy.inGameNoScore}</span>
      {/if}
    </li>
    <li data-testid="character-in-game-planner">
      {scoreCopy.inGamePlannerLead}
      <a
        class="text-text inline-flex min-h-11 items-center underline md:min-h-0"
        href={scoreCopy.inGamePasteHref}>{scoreCopy.inGamePasteLink}</a
      >.
    </li>
    {#if lastSeen !== null}
      <li class="text-muted" data-testid="character-in-game-last-seen">{lastSeen}</li>
    {/if}
  </ul>
</section>
