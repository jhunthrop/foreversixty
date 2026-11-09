<!-- web/src/components/ui/Frame.svelte -->
<!-- A featured panel (round 5, The frame): the rounded panel with a 1 px cream line inset 6 px on
     its radius, their bordered square elements made soft. `ornate` adds the corner flourishes and
     is the page's one lead panel (home's raid plaque, BiS's paperdoll, the tier list's spec
     panel). Two or three framed panels a page, never more; everything else is a plain Plate.
     `art` puts a painted background behind the content under its veil (the raid plaque's cave). -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { frameClass } from '../../lib/ui/primitives';
  import FrameCorners from './FrameCorners.svelte';

  let {
    ornate = false,
    art,
    label,
    testid,
    region,
    children,
  }: {
    ornate?: boolean;
    /** A painted background (lib/ui/assets.ts), cover, under the cave veil. */
    art?: 'cave';
    /** The panel's accessible name, when it is a landmark of its own. */
    label?: string;
    testid?: string;
    /** The parity region it carries, when it is one (lib/ui/regions.ts). */
    region?: string;
    children: Snippet;
  } = $props();
</script>

<div
  class="ui-frame ui-surface {frameClass(ornate ? 'lead' : 'framed')}"
  data-art={art}
  role={label ? 'region' : undefined}
  aria-label={label}
  data-testid={testid}
  data-region={region}
>
  {#if ornate}<FrameCorners />{/if}
  {@render children()}
</div>

<style>
  .ui-frame {
    min-width: 0;
    overflow: hidden;
    padding: 22px 24px;
  }
  .ui-frame[data-art='cave'] {
    background:
      linear-gradient(var(--color-veil-cave), var(--color-veil-cave)),
      var(--art-cave, none) 35% 40% / 160% auto no-repeat,
      var(--color-raised);
  }
  @media (max-width: 767px) {
    .ui-frame {
      padding: 18px 16px;
    }
    .ui-frame[data-art='cave'] {
      background-image:
        linear-gradient(var(--color-veil-cave), var(--color-veil-cave)), var(--art-cave-phone, none);
    }
  }
</style>
