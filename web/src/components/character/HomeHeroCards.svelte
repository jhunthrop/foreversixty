<!-- web/src/components/character/HomeHeroCards.svelte -->
<!-- Home rebuild spec §3.B.2: the signed-in hero's three next-action cards (Best in slot /
     Talents / Simulator). Extracted out of HomeAccountPanel.svelte (review round 3,
     "islands" item): the Simulator card's own fetch (lib/sim/api.ts's `listMySims`) and its
     dependencies (lib/home/next-steps.ts's `simCardLine`, ui/Skeleton.svelte) were
     previously imported at HomeAccountPanel's own top level, so every signed-out visitor's
     browser downloaded that code too, even though the component never reaches the branch
     that uses it -- Vite/Rollup bundles a file's static imports into its own chunk
     regardless of which conditional branch actually runs. Mounted by HomeAccountPanel.svelte
     into its own `home-hero-cards-slot` only once `ready && me !== null && hero !== null`,
     the same dynamic-import-once-signed-in trick it already uses for
     HomeSwitchCharacterPanel, so a signed-out page never fetches this module or its
     dependencies either. -->
<script lang="ts">
  import type { MeCharacter } from '../../lib/account/api';
  import { homePanelCopy, homeHeroCardsCopy } from '../../lib/home-panel-copy';
  import { armorySimHref } from '../../lib/sim/url';
  import { listMySims } from '../../lib/sim/api';
  import { simCardLine } from '../../lib/home/next-steps';
  import type { SimListRow } from '../../lib/sim/types';
  import Skeleton from '../ui/Skeleton.svelte';

  let { hero }: { hero: MeCharacter } = $props();

  // The Simulator card's own data: the visitor's latest saved sim, the one figure of the
  // three next-action cards this site can compute honestly today (§3.B.2's own ruling).
  // Independent per-card loading, the same rule HomeNextSteps.svelte's own cards used to
  // follow before that component was removed.
  let latestSim = $state<SimListRow | null>(null);
  let simStatus = $state<'loading' | 'ready' | 'failed'>('loading');

  $effect(() => {
    const current = hero;
    simStatus = 'loading';
    void listMySims(1)
      .then((page) => {
        if (current !== hero) return;
        latestSim = page.rows[0] ?? null;
        simStatus = 'ready';
      })
      .catch(() => {
        if (current !== hero) return;
        simStatus = 'failed';
      });
  });

  const CARD_CLASS =
    'flex flex-col gap-[6px] p-[14px_16px] rounded-panel border border-line bg-gradient-to-b from-card-top to-raised shadow-[inset_0_-1px_0_rgba(229,185,85,.35)] text-strong hover:border-gold-deep transition-colors duration-150';
</script>

<!-- The three next-action cards (§3.B.2): Best in slot and Talents are not yet computable
     against live data (no worn-gear or talent-compare source exists), so they show one
     settled, honest line instead of a figure that would have to be invented -- see
     homeHeroCardsCopy's own doc. Simulator alone carries real data. -->
<div class="reveal grid grid-cols-1 gap-3 sm:grid-cols-3" data-testid="home-hero-cards">
  <a href="#upgrades" class={CARD_CLASS} data-testid="home-hero-card-bis">
    <span class="label text-gold">{homeHeroCardsCopy.bestInSlotLabel}</span>
    <span class="text-muted text-[13px]">{homeHeroCardsCopy.bestInSlotNotAvailable}</span>
  </a>
  <a href="/planner" class={CARD_CLASS} data-testid="home-hero-card-talents">
    <span class="label text-gold">{homeHeroCardsCopy.talentsLabel}</span>
    <span class="text-muted text-[13px]">{homeHeroCardsCopy.talentsNotAvailable}</span>
  </a>
  <a href={armorySimHref(hero.key)} class={CARD_CLASS} data-testid="home-hero-card-sim">
    <span class="label text-gold">{homeHeroCardsCopy.simulatorLabel}</span>
    {#if simStatus === 'loading'}
      <Skeleton lines={2} rowHeight="h-3" testid="home-hero-card-sim-skeleton" />
    {:else if simStatus === 'failed'}
      <span class="text-muted text-[12px]">{homePanelCopy.noSimYet}</span>
    {:else if latestSim !== null}
      <span class="text-strong text-[13px] font-semibold" data-testid="home-hero-card-sim-value"
        >{simCardLine(latestSim)}</span
      >
    {:else}
      <span class="text-muted text-[13px]">{homePanelCopy.noSimYet}</span>
      <span class="text-nav text-[12px] font-semibold">{homePanelCopy.runAction}</span>
    {/if}
  </a>
</div>
