<!-- web/src/components/guild/RaidDoneSections.svelte -->
<!-- What a matched raid adds under the band (attendance spec 3.4, 3.5): the Report for everyone and,
     for officers, Attendance after it. A matched log that later turned private or flagged keeps its
     match but the raid view then carries no night, report or attendance: nothing renders and nothing
     says why (the spec has no copy for it). -->
<script lang="ts">
  import type { GuildRaidView } from '../../lib/guild/api';
  import type { NightActions } from '../../lib/guild/night-actions.svelte';
  import RaidAttendance from './RaidAttendance.svelte';
  import RaidReport from './RaidReport.svelte';

  let {
    raid,
    actions,
    guildName,
    guildPath,
    settingsHref,
    officer,
    officerTools,
    gateId,
    viewerTimezone,
    locale,
  }: {
    raid: GuildRaidView;
    actions: NightActions;
    guildName: string;
    guildPath: string;
    settingsHref: string;
    officer: boolean;
    officerTools: boolean;
    gateId: string;
    viewerTimezone: string;
    locale: string;
  } = $props();

  const report = $derived(raid.report);
  const night = $derived(raid.night);
  const match = $derived(raid.match);
  const attendance = $derived(raid.attendance);
</script>

{#if report !== null && night !== null && match !== null}
  <RaidReport
    {raid}
    {report}
    {night}
    reportId={match.report_id}
    {guildName}
    {guildPath}
    {settingsHref}
    {officer}
    {officerTools}
    {gateId}
    posting={actions.posting}
    postLine={actions.postLine}
    onPost={() => void actions.post()}
    {viewerTimezone}
    {locale}
  />
{/if}
{#if officer && night !== null && attendance !== null}
  <RaidAttendance
    {night}
    rows={attendance}
    busy={actions.excuseBusy}
    lines={actions.excuseLines}
    readOnly={!officerTools}
    {gateId}
    onExcuse={(key, excused) => void actions.excuse(key, excused)}
  />
{/if}
