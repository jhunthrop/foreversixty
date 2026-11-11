package gg.foreversixty.app.raids.flows

import androidx.compose.foundation.background
import gg.foreversixty.app.designsystem.FsSize
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.onClick
import androidx.compose.ui.semantics.role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.platform.LocalResources
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.selected
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsTarget
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.PanelFrame
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ScaledType
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPanel
import gg.foreversixty.app.flows.FlowRule
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.flows.SUB
import gg.foreversixty.app.flows.tallyTitle
import gg.foreversixty.app.designsystem.classColor
import androidx.compose.ui.text.AnnotatedString
import gg.foreversixty.app.home.flows.HomeActions
import gg.foreversixty.app.home.flows.RaidCardModel
import gg.foreversixty.app.home.flows.RaidAnswers
import gg.foreversixty.app.home.flows.RaidLiveModel
import gg.foreversixty.app.raids.ui.LocalRaidIconBuild
import gg.foreversixty.app.items.GameIconImage
import gg.foreversixty.app.reports.ui.Chevron

/**
 * The Raid tab's head (spec 4.2, owner 2026-11-09): one compact line, the zone's icon, the raid, its
 * time and the pull, who the answer is for (E7); the four answers under it are the answer (AS5), drawn
 * as one 48 dp segmented pill, and with Late the minutes inline, the chosen one selected (F1); the
 * officer [notes] close the card (row 9).
 */
@Composable
internal fun RaidHead(raid: RaidCardModel, actions: HomeActions, notes: String? = null) {
    FlowPanel {
        Column(Modifier.padding(vertical = FsSpace.space2)) {
            Row(Modifier.fillMaxWidth().clickable(role = Role.Button) { actions.onRaid(raid.raidId) }, verticalAlignment = Alignment.CenterVertically) {
                Box(
                    Modifier
                        .size(HEAD_ICON)
                        .parityItem(ParityItemKind.Image)
                        .clip(RoundedCornerShape(FsRadius.control))
                        .border(1.dp, FsColor.line, RoundedCornerShape(FsRadius.control)),
                ) { GameIconImage(raid.icon, LocalRaidIconBuild.current, HEAD_ICON) }
                Spacer(Modifier.width(FsSpace.space3))
                Column(Modifier.weight(1f)) {
                    Text(raid.zone, style = FsType.item, color = FsColor.strong, modifier = Modifier.semantics { heading() })
                    Text(raid.headLine, style = SUB, color = FsColor.muted)
                    raid.facts?.let { Text(it, style = SUB, color = FsColor.muted) }
                    raid.who?.let { Text(it, style = SUB, color = FsColor.muted) }
                }
            }
            RaidAnswers(raid, actions, Modifier.padding(top = FsSpace.space3))
            // Row 9: the officer notes (`Bring flasks`) close the card, on both apps.
            notes?.let { RaidNotes(it) }
        }
    }
}

/**
 * Who is coming's rows: the role's tally (S6, its count in its tone), Damage's melee and ranged split,
 * then the names in their class colours; [footer] closes the list (`Bench 1 · Out 4 · 14 not answered`,
 * the Covered line, an officer's readiness).
 */
@Composable
internal fun RoleRows(
    roles: List<RoleLine>,
    footer: List<RoleFoot> = emptyList(),
    onRosterOf: (Set<String>) -> Unit = {},
    onCover: () -> Unit = {},
) {
    roles.forEachIndexed { index, line ->
        Column(Modifier.fillMaxWidth().parityItem(ParityItemKind.Row).padding(vertical = ROLE_PAD)) {
            Text(roleTitle(line), style = FsType.item, color = FsColor.strong)
            Text(roleNames(line), style = SUB, color = FsColor.muted)
        }
        if (index < roles.lastIndex || footer.isNotEmpty()) FlowRule()
    }
    footer.forEach { foot ->
        val cover = foot.cover
        when {
            cover != null -> CoverRow(cover, foot.text, onCover)
            foot.onlyKeys.isNotEmpty() -> NotSyncedRow(foot.text) { onRosterOf(foot.onlyKeys) }
            else -> Text(foot.text, style = SUB, color = FsColor.muted, modifier = Modifier.padding(vertical = FsSpace.space2))
        }
    }
}

/**
 * An officer's not-synced line (round 5 row 1, as iOS): the names in the primary text tone, a trailing
 * chevron, the whole 48 dp row the tap that opens the roster filtered to them.
 */
@Composable
private fun NotSyncedRow(text: String, onClick: () -> Unit) {
    Row(
        RAID_TARGET.fillMaxWidth().clickable(role = Role.Button, onClick = onClick),
        verticalAlignment = Alignment.CenterVertically,
        horizontalArrangement = Arrangement.spacedBy(FsSpace.space1),
    ) {
        Text(text, style = SUB, color = FsColor.text, modifier = Modifier.weight(1f))
        Chevron(FsSize.icon20)
    }
}

/**
 * The Covered line (row 4, round 4 row 1): a neutral `Missing`, then each name in wipe red, with its
 * spell icon when it has one, joined by the one separator, then `and N more`. A name never breaks
 * across lines; the line wraps between names. Read aloud as [spoken], one sentence. With names left
 * out (round 5 row 4) the tail is a gold link and the 48 dp line opens the raid page's cover list.
 */
@Composable
private fun CoverRow(line: CoverLine, spoken: String, onMore: () -> Unit) {
    val res = LocalResources.current
    val names = RaidTabFacts(res).names(line)
    val more = line.more > 0
    val tap = if (more) RAID_TARGET.clickable(role = Role.Button, onClick = onMore) else Modifier
    FlowRow(
        Modifier
            .fillMaxWidth()
            .then(tap)
            .padding(vertical = FsSpace.space2)
            .clearAndSetSemantics {
                contentDescription = spoken
                if (more) {
                    role = Role.Button
                    onClick { onMore(); true }
                }
            },
        horizontalArrangement = Arrangement.spacedBy(FsSpace.space1),
        verticalArrangement = Arrangement.spacedBy(FsSpace.space1, Alignment.CenterVertically),
    ) {
        // Every piece sits on the text baseline, so a name beside its icon lines up with one without.
        val onBaseline = Modifier.alignByBaseline()
        if (line.missing.isNotEmpty()) Text(stringResource(R.string.raid_tab_coverLabel), onBaseline, style = SUB, color = FsColor.muted, maxLines = 1, softWrap = false)
        names.forEachIndexed { index, name -> CoverNameView(name, separated = index < names.lastIndex, onBaseline) }
        if (more) {
            Text(
                res.getString(R.string.raid_tab_andMore, "", line.more).trim(),
                onBaseline,
                style = SUB,
                color = FsColor.gold,
                fontWeight = FontWeight.SemiBold,
                maxLines = 1,
                softWrap = false,
            )
        }
    }
}

/**
 * One name, its icon only when it has one that loads (a missing file closes the gap, round 4 row 1),
 * then the separator in the muted tone when another name follows.
 */
@Composable
private fun CoverNameView(name: CoverName, separated: Boolean, modifier: Modifier = Modifier) {
    val build = LocalRaidIconBuild.current
    var iconShown by remember(name.icon, build) { mutableStateOf(name.hasIcon && build != null) }
    Row(modifier, verticalAlignment = Alignment.CenterVertically) {
        if (iconShown) {
            Box(Modifier.size(COVER_ICON).clip(RoundedCornerShape(FsRadius.control))) {
                GameIconImage(name.icon, build, COVER_ICON, onUnavailable = { iconShown = false })
            }
            Spacer(Modifier.width(FsSpace.space1))
        }
        Text(
            buildAnnotatedString {
                withStyle(SpanStyle(color = FsColor.wipe)) { append(name.text) }
                if (separated) withStyle(SpanStyle(color = FsColor.muted)) { append(RaidTabFacts.SEPARATOR.trimEnd()) }
            },
            style = SUB,
            maxLines = 1,
            softWrap = false,
        )
    }
}

/** `Damage 11 of 30 · 5 melee, 6 ranged`: the tally, then the split. */
@Composable
private fun roleTitle(line: RoleLine): AnnotatedString {
    val title = tallyTitle(line.tally)
    val split = line.split ?: return title
    return buildAnnotatedString {
        append(title)
        withStyle(SpanStyle(color = FsColor.muted, fontWeight = FontWeight.Normal)) { append(" · $split") }
    }
}

/** The names, each in its class colour, then `and 8 more`. */
private fun roleNames(line: RoleLine): AnnotatedString = buildAnnotatedString {
    line.names.forEachIndexed { index, name ->
        if (index > 0) append(", ")
        withStyle(SpanStyle(color = classColor(name.className))) { append(name.text) }
    }
    line.more?.let(::append)
}

/** While the raid runs (spec 4.2): Raid night leads, the pull and your answer, then who is in by role. */
@Composable
internal fun RaidNightBlock(live: RaidLiveModel, roles: List<RoleLine>) {
    FlowSection(stringResource(R.string.raid_tab_raidNight), top = 0.dp, panel = PanelFrame.Plain) {
        Row(Modifier.padding(top = FsSpace.space3), verticalAlignment = Alignment.CenterVertically) {
            Box(Modifier.size(LIVE_DOT).background(FsColor.kill, CircleShape))
            Spacer(Modifier.width(FsSpace.space2))
            Text(live.eyebrow, style = FsType.bodyStrong, color = FsColor.strong, modifier = Modifier.semantics { heading() })
        }
        FlowNote(live.line, Modifier.padding(top = FsSpace.space1, bottom = FsSpace.space2), color = FsColor.text)
        if (roles.isNotEmpty()) {
            FlowRule()
            RoleRows(roles)
        }
    }
}

/** A full-width 48 dp button under a panel's line: the flat teal primary, else the secondary. */
@Composable
internal fun RaidWideButton(label: String, primary: Boolean, onClick: () -> Unit) {
    val kind = if (primary) BrandButtonKind.Primary else BrandButtonKind.Secondary
    BrandButton(label, kind, onClick, Modifier.padding(top = FsSpace.space3, bottom = FsSpace.space3).parityItem(ParityItemKind.Button), height = FsTarget.min)
}

/** A row's button at the 48 dp target. */
internal val RAID_TARGET = Modifier.heightIn(min = FsTarget.min)

private val HEAD_ICON = FsSize.icon40
private val COVER_ICON = FsSize.icon20
private val ROLE_PAD = 10.dp
private val LIVE_DOT = 8.dp
