<!-- web/src/components/report/ClassIcon.svelte -->
<!-- A class's icon beside a name, from the build's icon set; a class the build has no
     icon for (a log from another client) shows a class-coloured dot instead of nothing. -->
<script lang="ts">
  import { classSlugOf } from '../../lib/report/planner-link';
  import { classColorVar } from '../../lib/report/format';
  import activeBuild from '../../data/active-build.json';

  let { className, size = 18 }: { className: string | undefined; size?: number } = $props();

  const slug = $derived(classSlugOf(className));
  let failed = $state(false);
</script>

{#if slug !== null && !failed}
  <img
    class="inline-block shrink-0 rounded-[3px] align-middle"
    style={`width: ${size}px; height: ${size}px`}
    src={`/data/${activeBuild.build}/icons/classicon_${slug}.webp`}
    alt=""
    aria-hidden="true"
    loading="lazy"
    onerror={() => (failed = true)}
  />
{:else}
  <span
    class="inline-block shrink-0 rounded-full align-middle"
    style={`width: ${Math.round(size / 2)}px; height: ${Math.round(size / 2)}px; background: ${classColorVar(className)}`}
    aria-hidden="true"
  ></span>
{/if}
