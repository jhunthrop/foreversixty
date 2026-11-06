<!-- web/src/components/sim/SimLandingPage.svelte -->
<!-- /sim with no character, the whole page (spec 2026-11-06-sim §4.2): PageHeader "Simulator" with
     its one meta line and the tool tabs, then the landing body. sim.astro renders it statically
     (the first paint and the LCP) and the island renders the same component, so hydration moves
     nothing. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { simPageCopy as copy } from '../../lib/sim/page-copy';
  import type { BisSetClass } from '../../lib/sim/bis-set-rail';
  import PageHeader from '../ui/PageHeader.svelte';
  import SimLanding from './SimLanding.svelte';
  import SimToolTabs from './SimToolTabs.svelte';

  let {
    bisSets,
    busy = false,
    pasteError = '',
    message = null,
    onpaste,
    onbuild,
    onfight,
    onsignin,
    notice,
    railTail,
  }: {
    bisSets: readonly BisSetClass[];
    busy?: boolean;
    pasteError?: string;
    message?: string | null;
    onpaste?: (code: string) => void;
    onbuild?: (ref: string) => void;
    onfight?: (ref: string) => void;
    onsignin?: () => void;
    /** A line between the head and the body (signed in with no characters yet). */
    notice?: Snippet;
    /** The rail's tail (Your sims, signed in). */
    railTail?: Snippet;
  } = $props();
</script>

<PageHeader title={copy.title} testid="sim-head">
  {#snippet meta()}<span>{copy.landingMeta}</span>{/snippet}
  {#snippet tabs()}<SimToolTabs />{/snippet}
</PageHeader>
{#if notice}{@render notice()}{/if}
<SimLanding {bisSets} {busy} {pasteError} {message} {onpaste} {onbuild} {onfight} {onsignin} {railTail} />
