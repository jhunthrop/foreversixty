<!-- web/src/components/ScoreBadge.svelte -->
<!-- The Forever Score as one number in its tier colour -- the same number and the same client
     item-quality colour the addon's tooltip shows (Score.tierColor), with the tier's name as
     its title. The one badge for every surface that shows a score: the character page, the
     /ratings example, the "In game" tooltip mock, and later the guild roster. A null score
     draws nothing; the caller shows the no-logs line instead, never a zero.

     md, lg and xl are large bold text (20px, 32px, 36-48px) in the exact quality colour: rare and epic sit
     under 4.5:1 on the site's dark ground, and large text needs only 3:1. sm sits in a line of
     body text, so it takes the design system's lightened rare and epic text variants. -->
<script lang="ts">
  import { tierColor, tierLabel, tierTextColor } from '../lib/score/score';

  const SIZE = {
    sm: 'text-[14px]',
    md: 'text-[20px]',
    lg: 'text-[32px]',
    /** The profile card's hero number: 36px on a phone, 48px from md. */
    xl: 'text-[36px] md:text-[48px]',
  } as const;

  let {
    score,
    tier,
    size = 'md',
  }: { score: number | null; tier: string | null; size?: keyof typeof SIZE } = $props();

  const color = $derived(size === 'sm' ? tierTextColor(tier) : tierColor(tier));
  const label = $derived(tierLabel(tier));
</script>

{#if score !== null}
  <span
    class="numeric leading-none font-bold {SIZE[size]}"
    style={color === null ? undefined : `color: ${color}`}
    title={label ?? undefined}
    data-testid="score-badge">{score}</span
  >
{/if}
