<!-- web/src/components/CharacterInGame.svelte -->
<!-- The character page's "In game" section (Task L8.2 ruling 5, fix round 1 item 5). A mock of
     the game's tooltip (fix round 2 item 3: its own frame, the name in its class colour, a
     descriptor line when known, a blank line), then the
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
    level = undefined,
    race = undefined,
  }: {
    forever: ForeverScore | null;
    name: string;
    className?: string;
    /** The viewer is signed in and owns this character, or it is their active character. */
    owned?: boolean;
    /** `build.captured_at` of the viewer's own addon export; null for anyone else. */
    lastSeenAt?: string | null;
    /** Level and race, for the tooltip's "Level 60 Human Priest" line; the character endpoint
     *  carries neither yet (API lane), so the line is left out until it does. */
    level?: number;
    race?: string;
  } = $props();

  const descriptor = $derived(
    level !== undefined && race !== undefined && className !== undefined
      ? scoreCopy.tooltipDescriptor(level, race, className)
      : null,
  );

  const lastSeen = $derived(owned && lastSeenAt !== null ? lastSeenText(lastSeenAt) : null);
  const line = scoreCopy.lineParts();
  /** In-text links are underlined: colour alone does not mark a link (WCAG 1.4.1). */
  /** An inline link, so the paragraph keeps the body line-height on a phone (fix round 3 item
   *  4); its vertical padding still gives it a 44px hit area without opening up the lines. */
  const LINK = 'text-gold underline underline-offset-2 max-md:py-[14px]';
</script>

<section class="flex flex-col gap-3" data-testid="character-in-game">
  <h2 class="section-title text-[18px]">{scoreCopy.inGameHeading}</h2>
  <!-- Below 900px the line sits under the mock; from 900px beside it. -->
  <div class="flex flex-col gap-3 min-[900px]:flex-row min-[900px]:items-center min-[900px]:gap-6">
    <!-- The game's own tooltip frame: a dark translucent fill, a 1px grey-blue border, a 4px
         radius, 10px padding; the name 14px bold in its class colour, the descriptor line when
         it is known, a blank line, then the addon's line. -->
    <div
      class="flex w-fit max-w-full min-w-[240px] flex-col rounded-[4px] border border-[#5d6b82] bg-[rgba(8,10,16,0.92)] p-[10px] text-[14px] leading-[20px] shadow-[0_2px_10px_rgba(0,0,0,.6)]"
      data-testid="character-in-game-tooltip"
    >
      <span class="text-[14px] font-bold" style={`color: ${classColorVar(className)}`}>{name}</span>
      {#if descriptor !== null}<span class="text-text" data-testid="character-in-game-descriptor"
          >{descriptor}</span
        >{/if}
      <span class="h-[20px]" aria-hidden="true"></span>
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
    <p class="text-[14px] leading-[22px]">
      {scoreCopy.inGameShownOn}
      <a class={LINK} href={scoreCopy.inGameSetupHref}>{scoreCopy.inGameSetupLink}</a>
    </p>
  </div>
  {#if owned}
    <p class="text-[14px] leading-[22px]" data-testid="character-in-game-planner">
      {scoreCopy.inGamePlannerLead}
      <a class={LINK} href={scoreCopy.inGamePasteHref}>{scoreCopy.inGamePasteLink}</a>.
    </p>
  {/if}
  {#if lastSeen !== null}
    <p class="text-muted text-[14px]" data-testid="character-in-game-last-seen">{lastSeen}</p>
  {/if}
</section>
