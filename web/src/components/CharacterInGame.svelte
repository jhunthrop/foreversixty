<!-- web/src/components/CharacterInGame.svelte -->
<!-- The character page's "In game" section (Task L8.2 ruling 5, fix round 1 item 5). A mock of
     the game's tooltip, matched to the 1.60 client (L8.4 fix round 2 item 6, from a real client
     capture): its own frame capped at the client's width (a long line wraps, never clips), the
     name in its class colour, then in the client's light grey-white "<Guild>" when known,
     "Level 60 Night Elf (Player)", the class and the faction ("Alliance", when known) on lines
     of their own; then a blank line (left out when there is neither a guild nor a level line)
     and the addon's own line "Forever Score 1830 · Epic · as of 4 Nov" (the fixture's `line` and
     `asOf`), or the fixture's no-logs line muted -- with one line beside it saying where it
     shows and linking setup. For the owner only (signed in and theirs, or their active
     character): how to copy the character in. When the site last synced it is the header's sync
     line (fix round 2 item 5), not repeated here. Plain data in: Character.svelte works out
     ownership and the faction. -->
<script lang="ts">
  import { factionName } from '../lib/faction-mark';
  import { classColorVar } from '../lib/report/format';
  import { scoreCopy } from '../lib/score/copy';
  import { hasScore, shortDate, tierLabel, tierTextColor, type ForeverScore } from '../lib/score/score';
  import ScoreBadge from './ScoreBadge.svelte';

  let {
    forever,
    name,
    className = undefined,
    owned = false,
    level = null,
    race = null,
    guild = null,
    faction = null,
  }: {
    forever: ForeverScore | null;
    name: string;
    className?: string;
    /** The viewer is signed in and owns this character, or it is their active character. */
    owned?: boolean;
    /** Level and race (the character endpoint's `CharacterIdentity`), for the tooltip's
     *  "Level 60 Night Elf (Player)" line; null when unknown, and that line, the class and the
     *  faction are left out with neither (Task L8.4 ruling 2, fix round 2 item 6). */
    level?: number | null;
    race?: string | null;
    /** The guild's name, the game's line under the character's name; null with no guild. */
    guild?: string | null;
    /** The faction, when the API knows it (the owner's /v1/me row); null leaves its line out. */
    faction?: 'alliance' | 'horde' | null;
  } = $props();

  const descriptor = $derived(scoreCopy.tooltipLevelLine(level, race));
  const classLine = $derived(
    descriptor !== null && className !== undefined && className !== '' ? className : null,
  );
  const factionLine = $derived(descriptor !== null && faction !== null ? factionName(faction) : null);
  const guildLine = $derived(
    typeof guild === 'string' && guild.trim() !== '' ? scoreCopy.tooltipGuild(guild.trim()) : null,
  );

  const line = scoreCopy.lineParts();
  /** In-text links are underlined: colour alone does not mark a link (WCAG 1.4.1). */
  /** An inline link, so the paragraph keeps the body line-height on a phone (fix round 3 item
   *  4); its vertical padding still gives it a 44px hit area without opening up the lines. The
   *  gold is the base `a` rule's, so hover and pressed take --gold-hover (L8.4 fix round 5
   *  item 2). */
  const LINK = 'underline underline-offset-2 max-md:py-[14px]';
  /** The client's light grey-white for every line under the name (sampled from the reference
   *  capture). */
  const LINE = 'text-[#dfe0e5]';
</script>

<section class="flex flex-col gap-3" data-testid="character-in-game">
  <h2 class="section-title text-[18px]">{scoreCopy.inGameHeading}</h2>
  <!-- Below 900px the line sits under the mock; from 900px beside it. -->
  <div class="flex flex-col gap-3 min-[900px]:flex-row min-[900px]:items-center min-[900px]:gap-6">
    <!-- The game's own tooltip frame: a dark translucent fill, a 1px grey-blue border, a 4px
         radius, 10px padding, at most the client's 320px wide (a long guild name wraps rather
         than pushing the score line out); the name 14px bold in its class colour, then in the
         client's order the guild, "Level 60 Night Elf (Player)", the class and the faction when
         known, a blank line, then the addon's line. -->
    <div
      class="flex w-fit max-w-[min(100%,320px)] min-w-[240px] flex-col rounded-[4px] border border-[#5d6b82] bg-[rgba(8,10,16,0.92)] p-[10px] text-[14px] leading-[20px] break-words shadow-[0_2px_10px_rgba(0,0,0,.6)]"
      data-testid="character-in-game-tooltip"
    >
      <span class="text-[14px] font-bold" style={`color: ${classColorVar(className)}`}>{name}</span>
      {#if guildLine !== null}<span class={LINE} data-testid="character-in-game-guild">{guildLine}</span>{/if}
      {#if descriptor !== null}<span class={LINE} data-testid="character-in-game-descriptor"
          >{descriptor}</span
        >{/if}
      {#if classLine !== null}<span class={LINE} data-testid="character-in-game-class">{classLine}</span>{/if}
      {#if factionLine !== null}<span class={LINE} data-testid="character-in-game-faction">{factionLine}</span
        >{/if}
      {#if guildLine !== null || descriptor !== null}<span class="h-[20px]" aria-hidden="true"></span>{/if}
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
</section>
