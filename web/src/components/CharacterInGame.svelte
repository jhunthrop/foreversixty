<!-- web/src/components/CharacterInGame.svelte -->
<!-- The character page's "In game" section (Task L8.2 ruling 5, fix round 1 item 5). A mock of
     the game's tooltip -- dark ground, gold border, the name in its class colour, then the
     addon's own line "Forever Score 1830 · Epic · as of 4 Nov" (the fixture's `line` and
     `asOf`), or the fixture's no-logs line muted -- with one line beside it saying where it
     shows and linking setup. For the owner only (signed in and theirs, or their active
     character): how to copy the character in, and when the site last saw its export.
     Plain data in: Character.svelte works out ownership and the export time. -->
<script lang="ts">
  import { classColorVar } from '../lib/report/format';
  import { scoreCopy } from '../lib/score/copy';
  import {
    hasScore,
    lastSeenText,
    shortDate,
    tierLabel,
    tierTextColor,
    type ForeverScore,
  } from '../lib/score/score';
  import ScoreBadge from './ScoreBadge.svelte';

  let {
    forever,
    name,
    className = undefined,
    owned = false,
    lastSeenAt = null,
  }: {
    forever: ForeverScore | null;
    name: string;
    className?: string;
    /** The viewer is signed in and owns this character, or it is their active character. */
    owned?: boolean;
    /** `build.captured_at` of the viewer's own addon export; null for anyone else. */
    lastSeenAt?: string | null;
  } = $props();

  const lastSeen = $derived(owned && lastSeenAt !== null ? lastSeenText(lastSeenAt) : null);
  const line = scoreCopy.lineParts();
  /** In-text links are underlined: colour alone does not mark a link (WCAG 1.4.1). */
  const LINK = 'text-gold inline-flex min-h-11 items-center underline underline-offset-2 md:min-h-0';
</script>

<section class="flex flex-col gap-3" data-testid="character-in-game">
  <h2 class="section-title text-[18px]">{scoreCopy.inGameHeading}</h2>
  <div class="flex flex-col gap-3 md:flex-row md:items-center md:gap-6">
    <div
      class="border-gold/70 bg-bg flex w-fit max-w-full min-w-[240px] flex-col gap-0.5 rounded-[4px] border px-3 py-2 text-[14px] leading-snug shadow-[0_2px_10px_rgba(0,0,0,.6)]"
      data-testid="character-in-game-tooltip"
    >
      <span class="font-semibold" style={`color: ${classColorVar(className)}`}>{name}</span>
      {#if hasScore(forever)}
        {@const tier = tierLabel(forever.tier)}
        {@const date = shortDate(forever.as_of)}
        <span data-testid="character-in-game-score"
          >{line.before}<ScoreBadge
            score={forever.score}
            tier={forever.tier}
            size="sm"
          />{line.after}{#if tier !== null}{scoreCopy.separator}<span
              style={`color: ${tierTextColor(forever.tier)}`}>{tier}</span
            >{/if}{#if date !== null}{scoreCopy.separator}<span class="text-muted"
              >{scoreCopy.asOf(date)}</span
            >{/if}</span
        >
      {:else}
        <span class="text-muted" data-testid="character-in-game-score">{scoreCopy.noLogs}</span>
      {/if}
    </div>
    <p class="text-[14px]">
      {scoreCopy.inGameShownOn}
      <a class={LINK} href={scoreCopy.inGameSetupHref}>{scoreCopy.inGameSetupLink}</a>
    </p>
  </div>
  {#if owned}
    <p class="text-[14px]" data-testid="character-in-game-planner">
      {scoreCopy.inGamePlannerLead}
      <a class={LINK} href={scoreCopy.inGamePasteHref}>{scoreCopy.inGamePasteLink}</a>.
    </p>
  {/if}
  {#if lastSeen !== null}
    <p class="text-muted text-[14px]" data-testid="character-in-game-last-seen">{lastSeen}</p>
  {/if}
</section>
