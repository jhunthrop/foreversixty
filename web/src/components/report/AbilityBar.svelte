<!-- web/src/components/report/AbilityBar.svelte -->
<!-- The amount bar, segmented by ability. Colour comes from one token -- the actor's class
     colour -- at five decreasing opacities, largest ability first, with a hairline of page
     background between segments. Using a palette here would either invent hex values the
     design system forbids or repurpose the class colours, which it forbids harder. -->
<script lang="ts">
  import { formatAmount, schoolName, schoolToken } from '../../lib/report/format';
  import { abilityKey, type Ability } from '../../lib/report/types';

  let {
    abilities,
    total,
    peak,
    color: _color = '',
  }: { abilities: Ability[]; total: number; peak: number; color?: string } = $props();
  // The class colour used to tint the bar; kept in the props so callers need not change,
  // unused now that every segment carries its school's colour.
  void _color;

  // Segments of one school step down in opacity so two fire spells still read as two;
  // the school sets the hue.
  const OPACITIES = [1, 0.8, 0.62, 0.5, 0.42];
  const ordered = $derived([...abilities].sort((a, b) => b.total - a.total));
  const schoolRank = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const seen = new Map<string, number>();
    return ordered.map((ability) => {
      const token = schoolToken(ability.school);
      const rank = seen.get(token) ?? 0;
      seen.set(token, rank + 1);
      return rank;
    });
  });
  const widthPct = $derived(peak <= 0 ? 0 : Math.max((total / peak) * 100, 0.5));
</script>

<div class="bg-line-soft h-[8px] w-full" style={`max-width: ${widthPct}%`} aria-hidden="true">
  <div class="flex h-full w-full">
    {#each ordered as ability, index (abilityKey(ability))}
      <span
        class="h-full"
        style={`flex: ${Math.max(ability.total, 0)} 0 0; background: ${schoolToken(ability.school)}; opacity: ${OPACITIES[Math.min(schoolRank[index] ?? 0, OPACITIES.length - 1)]}; border-right: 1px solid var(--color-bg)`}
        title={`${ability.name} · ${schoolName(ability.school)} · ${formatAmount(ability.total)}`}
      ></span>
    {/each}
  </div>
</div>
