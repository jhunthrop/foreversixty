package gg.foreversixty.app.raids.flows

import gg.foreversixty.app.raids.ui.LocalRaidIconBuild
import gg.foreversixty.app.items.ItemQuality
import gg.foreversixty.app.designsystem.GameColors
import androidx.compose.ui.draw.clip
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.border
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.testTag
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.FactionMark
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.PanelFrame
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.ScaledType
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.flows.FlowBlock
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPanel
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.model.RaidAssignedName
import gg.foreversixty.app.raids.logic.AssignmentEncounterLine
import gg.foreversixty.app.raids.logic.RaidJobLine

/** What the Guild tab's own sections do (spec 2026-11-10-guild-tab.md section 2). */
data class GuildTabActions(
    val onRoster: () -> Unit = {},
    val onLoot: () -> Unit = {},
    val onProgression: () -> Unit = {},
    val onReport: (String) -> Unit = {},
    val onPlanner: () -> Unit = {},
    val onOfficer: (OfficerKind) -> Unit = {},
)

/** The guild head: the crest (the faction mark, a guild with none), the name and its line; below 60 the level spread. */
@Composable
fun GuildHeadBlock(head: GuildTabHead, spread: GuildSpread?) {
    FlowBlock(ParityRegion.section(stringResource(R.string.tab_guild)), top = 0.dp) {
        FlowPanel {
            Row(Modifier.padding(vertical = FsSpace.space2), verticalAlignment = Alignment.Top) {
                GuildCrestMark(head.faction)
                Spacer(Modifier.width(FsSpace.space3))
                Column {
                    Text(head.name, style = ScaledType.sectionTitle, color = FsColor.cream)
                    Text(head.line, style = FsType.small, color = FsColor.muted)
                }
            }
            if (spread != null) SpreadFigures(spread)
        }
    }
}

@Composable
private fun SpreadFigures(spread: GuildSpread) {
    HorizontalDivider(color = FsColor.line)
    Row(Modifier.fillMaxWidth().padding(vertical = FsSpace.space2)) {
        Figure(spread.atCap.toString(), stringResource(R.string.guild_tab_atCap, GuildTabPresenter.MAX_LEVEL), FsColor.strong, Modifier.weight(1f))
        Figure(
            spread.close.toString(),
            stringResource(R.string.guild_tab_closeRange, GuildTabPresenter.CLOSE_FROM, GuildTabPresenter.MAX_LEVEL - 1),
            FsColor.strong,
            Modifier.weight(1f),
        )
        Figure(spread.third.toString(), spread.thirdLabel, if (spread.own) FsColor.gold else FsColor.strong, Modifier.weight(1f))
    }
}

@Composable
private fun Figure(value: String, label: String, color: Color, modifier: Modifier) {
    Column(modifier) {
        Text(value, style = FsType.itemLg.copy(fontFeatureSettings = "tnum"), color = color)
        Text(label, style = FsType.small, color = FsColor.muted)
    }
}

/** The guild's crest in the one ringed circle: the faction's mark inside, a guild with no crest of its own. */
@Composable
private fun GuildCrestMark(faction: String?) {
    Box(
        Modifier
            .size(CREST)
            .parityItem(ParityItemKind.Image)
            .clip(CircleShape)
            .background(FsColor.raised)
            .border(CREST_RING, GameColors.faction(faction.orEmpty()).copy(alpha = RING_ALPHA), CircleShape),
        contentAlignment = Alignment.Center,
    ) { FactionMark(faction.orEmpty(), CREST * MARK_SHARE) }
}

/** Roster: Recently on and, below 60, Close to 60; an officer's Missing raids; `Roster ›` to the roster page. */
@Composable
fun GuildRosterSection(sections: GuildTabSections, onRoster: () -> Unit) {
    if (sections.recent.isEmpty() && sections.close.isEmpty()) return
    FlowSection(
        stringResource(R.string.guild_tab_roster),
        action = stringResource(R.string.raid_tab_rosterLink),
        onAction = onRoster,
        meta = sections.rosterMeta,
        top = 0.dp,
        panel = PanelFrame.None,
    ) {
        if (sections.recent.isNotEmpty()) {
            GroupLabel(stringResource(R.string.guild_tab_recentlyOn))
            MemberRows(sections.recent, last = sections.close.isEmpty() && sections.missing.isEmpty() && sections.gearGaps.isEmpty() && sections.yourAttendance == null)
        }
        if (sections.close.isNotEmpty()) {
            GroupLabel(stringResource(R.string.guild_tab_closeTo, GuildTabPresenter.MAX_LEVEL))
            MemberRows(sections.close, last = sections.missing.isEmpty() && sections.gearGaps.isEmpty() && sections.yourAttendance == null)
        }
        if (sections.missing.isNotEmpty()) {
            GroupLabel(stringResource(R.string.guild_tab_missingRaids))
            MemberRows(sections.missing, last = sections.gearGaps.isEmpty() && sections.yourAttendance == null)
        }
        if (sections.gearGaps.isNotEmpty()) {
            GroupLabel(stringResource(R.string.guild_tab_gearGaps))
            MemberRows(sections.gearGaps, last = sections.yourAttendance == null)
        }
        sections.yourAttendance?.let {
            GroupLabel(stringResource(R.string.guild_tab_yourAttendance))
            MemberRows(listOf(it), last = true)
        }
    }
}

@Composable
private fun GroupLabel(text: String) {
    Text(text, style = FsType.small.copy(fontWeight = FontWeight.SemiBold), color = FsColor.muted, modifier = Modifier.padding(top = FsSpace.space2))
}

@Composable
private fun MemberRows(lines: List<GuildMemberLine>, last: Boolean) {
    lines.forEachIndexed { index, line ->
        FlowRow(
            title = line.name,
            sub = line.line,
            lead = line.className?.let { FlowLead.Crest(it) },
            titleColor = classColor(line.className),
            value = line.trail,
            last = last && index == lines.lastIndex,
        )
    }
}

/** Recent logs: the guild's last reports, each opening the report. */
@Composable
fun GuildLogsSection(logs: List<GuildLogLine>, actions: GuildTabActions) {
    if (logs.isEmpty()) return
    FlowSection(stringResource(R.string.guild_tab_logs), top = 0.dp, panel = PanelFrame.None) {
        logs.forEachIndexed { index, log ->
            FlowRow(title = log.title, sub = log.line, value = log.day, last = index == logs.lastIndex, onClick = { actions.onReport(log.id) }, chevron = true)
        }
    }
}

/** Progression: the tier's line, opening the guild page's Progression on the site. */
@Composable
fun GuildProgressSection(progression: GuildProgressLine?, onProgression: () -> Unit) {
    progression ?: return
    FlowSection(stringResource(R.string.guild_tab_progression), top = 0.dp, panel = PanelFrame.None) {
        FlowRow(title = progression.title, sub = progression.line, last = true, onClick = onProgression, chevron = true)
    }
}

/** Loot at 60: the newest night's drops in their quality colour with who won them; `Loot history ›` to the guild's loot. */
@Composable
fun GuildLootSection(night: GuildLootNight?, loot: String?, onLoot: () -> Unit) {
    if (night == null && loot == null) return
    FlowSection(
        stringResource(R.string.guild_tab_loot), action = stringResource(R.string.guild_tab_lootLink), onAction = onLoot,
        meta = night?.day, top = 0.dp, panel = PanelFrame.None,
    ) {
        if (night == null) {
            FlowRow(title = loot.orEmpty(), last = true)
        } else {
            val notAwarded = stringResource(R.string.guild_tab_notAwarded)
            night.drops.forEachIndexed { index, drop ->
                FlowRow(
                    title = drop.name,
                    titleColor = ItemQuality(drop.quality).text,
                    lead = FlowLead.Icon(drop.icon, LocalRaidIconBuild.current),
                    styledSub = buildAnnotatedString {
                        val winner = drop.winner
                        if (winner == null) {
                            append(notAwarded)
                        } else {
                            // Yours in gold (the web's Tag tone), anyone else's in class colour.
                            withStyle(SpanStyle(color = if (drop.yours) FsColor.gold else classColor(drop.winnerClass))) { append(winner) }
                        }
                    },
                    last = index == night.drops.lastIndex,
                )
            }
        }
    }
}

/** Officer (spec 2.3): Plan a raid, Applications with their count, Setup with its progress; one tap each. */
@Composable
fun OfficerSection(rows: List<OfficerShortcut>, onOfficer: (OfficerKind) -> Unit) {
    if (rows.isEmpty()) return
    FlowSection(stringResource(R.string.guild_tab_officer), top = 0.dp, panel = PanelFrame.None) {
        rows.forEachIndexed { index, row ->
            Column(Modifier.testTag("guild-officer-${row.kind.name.lowercase()}")) {
                FlowRow(title = row.title, sub = row.line, value = row.count, last = index == rows.lastIndex, onClick = { onOfficer(row.kind) }, chevron = true)
            }
        }
    }
}

/** Your job (spec 2.3a): every slot the viewer holds, with the partners, the boss and the purpose; left out with none. */
@Composable
fun YourJobSection(jobs: List<RaidJobLine>) {
    if (jobs.isEmpty()) return
    FlowSection(stringResource(R.string.shared_assign_yours), top = 0.dp, panel = PanelFrame.None) {
        jobs.forEachIndexed { index, job -> FlowRow(title = job.title, sub = job.line, last = index == jobs.lastIndex) }
    }
}

/**
 * An officer's read-only Assignments (spec 2.3a): each encounter's filled slots, the names (yours in gold), how full;
 * `Open in planner ›` to the raid's page on the site, where assigning stays.
 */
@Composable
fun AssignmentsSection(encounters: List<AssignmentEncounterLine>, onPlanner: () -> Unit) {
    FlowSection(
        stringResource(R.string.shared_assign_title), action = stringResource(R.string.guild_tab_assignPlanner), onAction = onPlanner,
        top = 0.dp, panel = PanelFrame.None,
    ) {
        Column(Modifier.testTag("guild-assignments")) {
            if (encounters.isEmpty()) FlowNote(stringResource(R.string.shared_assign_nothingYet), Modifier.padding(vertical = FsSpace.space2), FsColor.text)
            encounters.forEach { encounter ->
                GroupLabel(encounter.title)
                encounter.slots.forEach { slot ->
                    FlowRow(
                        title = slot.label,
                        titleColor = if (slot.names.isEmpty()) FsColor.muted else FsColor.strong,
                        styledSub = if (slot.names.isEmpty()) null else names(slot.names),
                        value = slot.filled,
                    )
                }
            }
        }
    }
}

/** A slot's names: the viewer's own in gold, the rest muted (the web planner's Tag tones). */
private fun names(names: List<RaidAssignedName>) = buildAnnotatedString {
    names.forEachIndexed { index, name ->
        if (index > 0) withStyle(SpanStyle(color = FsColor.muted)) { append(", ") }
        withStyle(SpanStyle(color = if (name.own) FsColor.gold else FsColor.muted)) { append(name.name) }
    }
}

private val CREST = 48.dp
private val CREST_RING = 2.dp
private const val RING_ALPHA = 0.6f
/** The faction mark's share of the ringed circle. */
private const val MARK_SHARE = 0.72f
