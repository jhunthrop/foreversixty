<!-- web/src/components/ScoreBadge.svelte -->
<!-- The Forever Score as one number in its tier colour -- the same number and the same client
     item-quality colour the addon's tooltip shows (Score.tierColor), with the tier's name as
     its title. The one badge for every surface that shows a score: the character page, the
     /ratings example, and later the guild roster. A null score draws nothing; the caller
     shows the no-logs line instead, never a zero.

     Always bold and never under 19px: rare and epic sit under 4.5:1 on the site's dark
     ground, so the coloured number has to be large text (WCAG's 3:1) to stay readable. -->
<script lang="ts">
  import { tierColor, tierLabel } from '../lib/score/score';

  const SIZE = { md: 'text-[20px]', lg: 'text-[32px]' } as const;

  let {
    score,
    tier,
    size = 'md',
  }: { score: number | null; tier: string | null; size?: keyof typeof SIZE } = $props();

  const color = $derived(tierColor(tier));
  const label = $derived(tierLabel(tier));
</script>

{#if score !== null}
  <span
    class="tabular font-mono leading-none font-bold {SIZE[size]}"
    style={color === null ? undefined : `color: ${color}`}
    title={label ?? undefined}
    data-testid="score-badge">{score}</span
  >
{/if}
