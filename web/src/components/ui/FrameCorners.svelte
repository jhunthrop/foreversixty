<!-- web/src/components/ui/FrameCorners.svelte -->
<!-- The lead panel's four corner flourishes (round 5, design/mocks/directions/style.py CORNERS): the
     inset line turning on the radius, an inner arc ending in a small scroll, and a bead. Drawn in
     currentColor (panel.css: the cream at 80 %), rotated per corner. Decorative only. -->
<script lang="ts">
  const CORNERS = ['tl', 'tr', 'br', 'bl'] as const;
</script>

{#each CORNERS as corner (corner)}
  <svg class="ui-frame-corner" data-corner={corner} viewBox="0 0 26 26" aria-hidden="true" focusable="false">
    <g fill="none" stroke="currentColor" stroke-linecap="round">
      <path d="M6 15 Q6 6 15 6" stroke-width="1.1" />
      <path d="M15 6 q3 0 3 3 q0 2 -2 2" stroke-width="1" />
      <circle cx="6" cy="18.5" r="1.3" fill="currentColor" stroke="none" />
    </g>
  </svg>
{/each}
