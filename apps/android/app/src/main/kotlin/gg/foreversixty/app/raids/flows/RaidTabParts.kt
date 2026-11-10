package gg.foreversixty.app.raids.flows

import androidx.compose.foundation.background
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
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.platform.LocalDensity
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

/**
 * The Raid tab's head (spec 4.2, owner 2026-11-09): one compact line, the zone's icon, the raid, its
 * time and the pull, who the answer is for (E7) and the answer as a chip; the four answers under it
 * as one 48 dp segmented pill, and with Late the minutes inline, the chosen one selected (F1).
 */
@Composable
internal fun RaidHead(raid: RaidCardModel, actions: HomeActions) {
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
                // At large text the chip drops under the lines, so the raid's name keeps the row's width.
                val chipBeside = LocalDensity.current.fontScale <= CHIP_BESIDE_MAX_SCALE
                Column(Modifier.weight(1f)) {
                    Text(raid.zone, style = FsType.item, color = FsColor.strong, modifier = Modifier.semantics { heading() })
                    Text(raid.headLine, style = SUB, color = FsColor.muted)
                    raid.facts?.let { Text(it, style = SUB, color = FsColor.muted) }
                    raid.who?.let { Text(it, style = SUB, color = FsColor.muted) }
                    if (!chipBeside) raid.chip?.let { AnswerChip(it, Modifier.padding(top = FsSpace.space1)) }
                }
                if (chipBeside) {
                    raid.chip?.let {
                        Spacer(Modifier.width(FsSpace.space2))
                        AnswerChip(it)
                    }
                }
            }
            RaidAnswers(raid, actions, Modifier.padding(top = FsSpace.space3))
        }
    }
}

/** Your answer as a chip at the head's end, teal as the selected segment: `In`, `Late · 30 min`. */
@Composable
private fun AnswerChip(text: String, modifier: Modifier = Modifier) {
    val shape = RoundedCornerShape(CHIP_RADIUS)
    Text(
        text,
        style = CHIP_TEXT,
        color = FsColor.strong,
        maxLines = 1,
        modifier = modifier
            .background(FsColor.teal, shape)
            .border(1.dp, FsColor.tealEdge, shape)
            .padding(horizontal = CHIP_PAD_X, vertical = CHIP_PAD_Y),
    )
}

/**
 * Who is coming's rows: the role's tally (S6, its count in its tone), Damage's melee and ranged split,
 * then the names in their class colours; [footer] closes the list (`Bench 1 · Out 4 · 14 not answered`,
 * the Covered line, an officer's readiness).
 */
@Composable
internal fun RoleRows(roles: List<RoleLine>, footer: List<RoleFoot> = emptyList()) {
    roles.forEachIndexed { index, line ->
        Column(Modifier.fillMaxWidth().parityItem(ParityItemKind.Row).padding(vertical = ROLE_PAD)) {
            Text(roleTitle(line), style = FsType.item, color = FsColor.strong)
            Text(roleNames(line), style = SUB, color = FsColor.muted)
        }
        if (index < roles.lastIndex || footer.isNotEmpty()) FlowRule()
    }
    footer.forEach { foot ->
        Text(foot.text, style = SUB, color = if (foot.warn) FsColor.wipe else FsColor.muted, modifier = Modifier.padding(vertical = FsSpace.space2))
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

private val HEAD_ICON = 40.dp
private const val CHIP_BESIDE_MAX_SCALE = 1.3f
private val CHIP_RADIUS = 4.dp
private val CHIP_PAD_X = 10.dp
private val CHIP_PAD_Y = 3.dp
private val ROLE_PAD = 10.dp
private val LIVE_DOT = 8.dp
private val CHIP_TEXT = FsType.answerChip
