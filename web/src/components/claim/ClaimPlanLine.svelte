<!-- web/src/components/claim/ClaimPlanLine.svelte -->
<!-- Spec 2026-10-27 §6.2: what a claim link made for a plan is for. 14 px --text on a 2 px
     --gold left rule, padding-left 14. -->
<script lang="ts">
  let { line }: { line: string } = $props();
</script>

<p class="text-text border-gold border-l-2 pl-[14px] text-[14px]" data-testid="guild-claim-plan-line">
  {line}
</p>
