<!-- web/src/components/home/HomeRail.svelte -->
<!-- The signed-in home's rail (home spec 2026-11-06 H7): Guild when the character is in one, the
     character's own logs, and Set up (the gear source, Battle.net first; the addon is an optional extra).
     Open Sections of open Rows, no surface. -->
<script lang="ts">
  import type { Me, MeCharacter, MyReport } from '../../lib/account/api';
  import { homeCopy, homeHrefs, logSub, setUpView } from '../../lib/home/home-view';
  import { LINK } from '../../lib/ui/primitives';
  import Checklist from '../ui/Checklist.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';

  let {
    me,
    character,
    logs,
    guildHref,
  }: {
    me: Me;
    character: MeCharacter;
    logs: readonly MyReport[];
    guildHref: string | null;
  } = $props();

  const setUp = $derived(setUpView(me, character));

  const guild = $derived(character.guild);
  const guildSub = $derived.by(() => {
    if (guild === undefined) return '';
    const parts = [guild.rank ?? ''];
    if (guild.attendance !== undefined)
      parts.push(homeCopy.attendance(guild.attendance.present, guild.attendance.nights));
    return parts.filter(Boolean).join(' · ');
  });
</script>

{#if guild !== undefined}
  <Section title={homeCopy.guild} testid="home-guild">
    <Row title={guild.name} sub={guildSub} href={guildHref ?? undefined} open />
  </Section>
{/if}
<Section title={homeCopy.logs} action={{ label: homeCopy.allLogs, href: homeHrefs.logs }} testid="home-logs">
  {#if logs.length === 0}
    <Row title={homeCopy.noLogs} sub={homeCopy.noLogsSub} href={homeHrefs.logsUpload} open />
  {:else}
    {#each logs as report (report.id)}
      <Row title={report.title} sub={logSub(report)} href={homeHrefs.report(report.id)} open />
    {/each}
  {/if}
</Section>
<Section title={homeCopy.setUp} testid="home-set-up">
  <Checklist items={setUp.checks} testid="home-set-up-list" />
  {#if setUp.addonOptional}
    <Row title={homeCopy.addonOptional} sub={homeCopy.addonOptionalSub} open testid="home-addon-optional">
      {#snippet foot()}<a class={LINK} href={homeHrefs.addon}>{homeCopy.addonOptionalLink} ›</a>{/snippet}
    </Row>
  {/if}
</Section>
