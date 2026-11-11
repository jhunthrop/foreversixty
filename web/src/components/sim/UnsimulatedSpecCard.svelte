<!-- web/src/components/sim/UnsimulatedSpecCard.svelte -->
<!-- One of /sim/specs' healer and tank specs (mock design/mocks/ws_simtools.py `web-sim-specs`): the
     spec's own icon (the class crest when the build has none, as the damage rows lead) and its
     name in its class colour, a link to its guide, set in one column on a phone (name left, role
     right, 44 px rows) and four across from 768, under their heading (h-audit 2026-11-09: eight
     rows of "Guide ›" said nothing a link line does not). No state tag, engine stamp or rotation. -->
<script lang="ts">
  import { classSpecPath } from '../../lib/sim/gear-card';
  import { classOfSpec, specLabel } from '../../lib/sim/spec-label';
  import { specLeadIcon } from '../../lib/sim/spec-icon';
  import Crest from '../ui/Crest.svelte';
  import type { Spec } from '../../lib/sim/specs';
  import { classColorVar } from '../../lib/report/format';
  import { simCopy } from '../../lib/sim/copy';

  let { spec }: { spec: Spec } = $props();
  let iconFailed = $state(false);
  const lead = $derived(specLeadIcon(spec.spec, iconFailed));
</script>

<span id={spec.spec} class="flex min-h-11 min-w-0 items-center" data-testid={`spec-unsimulated-${spec.spec}`}>
  <a
    href={`/guides/${classSpecPath(spec.spec)}`}
    class="flex w-full min-w-0 items-center gap-2 text-[15px] font-bold whitespace-nowrap hover:underline md:inline-flex md:w-auto"
    style:color={classColorVar(classOfSpec(spec.spec))}
    data-testid={`spec-guide-${spec.spec}`}
    >{#if lead === null}<Crest cls={classOfSpec(spec.spec)} size={24} />{:else}<img
        src={lead}
        alt=""
        width="24"
        height="24"
        class="rounded-[5px]"
        onerror={() => (iconFailed = true)}
      />{/if}{specLabel(spec.spec)}<span class="text-muted ml-auto text-[13px] font-normal md:ml-0"
      >{simCopy.specsRole[spec.role]}</span
    ></a
  >
</span>
