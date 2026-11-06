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
    shell = false,
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
    /** The static render in sim.astro: carries the payer's reserve box. */
    shell?: boolean;
  } = $props();
</script>

<PageHeader title={copy.title} testid="sim-head">
  {#snippet meta()}<span>{copy.landingMeta}</span>{/snippet}
  {#snippet tabs()}<SimToolTabs />{/snippet}
</PageHeader>
<!-- Hidden unless html[data-session='1'][data-server-sims='1'] (global.css): a payer's "Your
     server runs" box, reserved before the island paints (spec 2026-10-24 3.7). The island's own
     list takes this place. -->
{#if shell}<div class="server-runs-reserve server-runs-shell-reserve" aria-hidden="true"></div>{/if}
{#if notice}{@render notice()}{/if}
<SimLanding {bisSets} {busy} {pasteError} {message} {onpaste} {onbuild} {onfight} {onsignin} />
