<!-- web/src/components/ui/GradeTile.svelte -->
<!-- A tier letter on its tier colour (home B spec 2026-11-07 R3): the tier list's letters and
     their item-quality tones (lib/tiers/tier-rules.ts), 56 px square, the letter read aloud with
     its meaning. -->
<script lang="ts">
  import { TIER_LETTER_TONES, type TierLetter } from '../../lib/tiers/tier-rules';

  let { letter, testid }: { letter: TierLetter; testid?: string } = $props();
</script>

<span
  class="ui-grade"
  style:background={TIER_LETTER_TONES[letter]}
  role="img"
  aria-label={`Tier ${letter}`}
  data-testid={testid}>{letter}</span
>

<style>
  .ui-grade {
    display: inline-flex;
    flex: none;
    align-items: center;
    justify-content: center;
    width: 56px;
    height: 56px;
    border-radius: 8px;
    font-weight: 700;
    font-size: 32px;
    line-height: 1;
    color: var(--color-bg);
  }
</style>
