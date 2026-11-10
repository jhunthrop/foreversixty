package gg.foreversixty.app.raids.flows

import gg.foreversixty.app.flows.FlowChip
import androidx.compose.foundation.clickable
import androidx.compose.foundation.horizontalScroll
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.drawBehind
import androidx.compose.ui.geometry.Offset
import androidx.compose.ui.geometry.Size
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.ParityScrollRegion
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.parityRegion
import gg.foreversixty.app.flows.FlowButton
import gg.foreversixty.app.flows.FlowHead
import gg.foreversixty.app.flows.FlowHeadIcon
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.flows.FlowNote
import gg.foreversixty.app.flows.FlowPage
import gg.foreversixty.app.flows.FlowRow
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.raids.logic.SignupState

/** What Not answered hands back; the route binds them. */
data class NotAnsweredActions(
    val select: (NotAnsweredFilter) -> Unit = {},
    val mentionRole: () -> Unit = {},
    val mentionAll: () -> Unit = {},
    val mark: (NotAnsweredRow, SignupState) -> Unit = { _, _ -> },
    val undo: (NotAnsweredRow) -> Unit = {},
    val message: (NotAnsweredRow) -> Unit = {},
    val editOnSite: () -> Unit = {},
)

/**
 * Keep the raid full (spec F3, mock raid-unanswered): the raid page filtered to Not answered, the short
 * role first; the mention that names it with its preview; per row Mark In (it confirms on the row with
 * Undo), Bench, Discord DM; the bench that can close the gap.
 */
@Composable
fun NotAnsweredScreen(content: NotAnsweredContent, build: String?, actions: NotAnsweredActions, error: String? = null, note: String? = null) {
    FlowPage {
        FlowHead(content.title, content.facts, lead = { FlowHeadIcon(content.icon, build) })
        Chips(content, actions.select)
        content.mention?.let { Mention(it, actions) }
        note?.let { FlowNote(it, Modifier.padding(top = FsSpace.space3)) }
        error?.let { FlowNote(it, Modifier.padding(top = FsSpace.space3), color = FsColor.wipe) }
        FlowSection(content.listTitle, meta = content.listMeta) {
            if (content.rows.isEmpty()) FlowNote(stringResource(R.string.raid_na_empty), Modifier.padding(top = FsSpace.space2))
            content.rows.forEachIndexed { index, row ->
                Raider(row, if (content.actionable) RowActions.All else RowActions.None, index == content.rows.lastIndex, actions)
            }
        }
        content.benchTitle?.let { title ->
            FlowSection(title) {
                content.bench.forEachIndexed { index, row -> Raider(row, RowActions.MarkIn, index == content.bench.lastIndex, actions) }
                content.benchNote?.let { FlowNote(it, Modifier.padding(top = FsSpace.space2)) }
            }
        }
        if (content.actionable) SiteLine(actions.editOnSite)
    }
}

/** The filters, Not answered preselected, scrolling sideways (`.chiprow`, 4 above, 8 below). */
@Composable
private fun Chips(content: NotAnsweredContent, select: (NotAnsweredFilter) -> Unit) {
    Row(
        Modifier
            .fillMaxWidth()
            .parityRegion(ParityRegion.CONTROLS)
            .horizontalScroll(rememberScrollState())
            .padding(top = FsSpace.space1, bottom = FsSpace.space2),
        horizontalArrangement = Arrangement.spacedBy(FsSpace.space2),
    ) {
        // The chips scroll sideways; every one counts in the region, on screen or not.
        ParityScrollRegion(ParityRegion.CONTROLS) {
            content.chips.forEach { chip ->
                val on = chip.filter == content.filter
                FlowChip(chip.title, on, { select(chip.filter) }, minHeight = FILTER_HEIGHT)
            }
        }
    }
}

/** The short role's mention (primary), Mention all, and the message they post (`.vstack`, `.preview`). */
@Composable
private fun Mention(mention: MentionContent, actions: NotAnsweredActions) {
    Column(Modifier.padding(top = FsSpace.space5).fillMaxWidth(), verticalArrangement = Arrangement.spacedBy(FsSpace.space3)) {
        mention.role?.let { FlowButton(it, actions.mentionRole, Modifier.fillMaxWidth(), primary = true, small = false) }
        FlowButton(mention.all, actions.mentionAll, Modifier.fillMaxWidth(), primary = mention.role == null, small = false)
    }
    val rule = FsColor.line
    Text(
        mention.preview,
        style = PREVIEW,
        color = FsColor.text,
        modifier = Modifier
            .padding(top = FsSpace.space3)
            .fillMaxWidth()
            .drawBehind { drawRect(rule, Offset.Zero, Size(PREVIEW_RULE.toPx(), size.height)) }
            .padding(horizontal = FsSpace.space3, vertical = PREVIEW_PAD),
    )
}

private enum class RowActions { All, MarkIn, None }

@Composable
private fun Raider(row: NotAnsweredRow, rowActions: RowActions, last: Boolean, actions: NotAnsweredActions) {
    val marked = row.marked
    val trail: (@Composable () -> Unit)? = when {
        marked != null -> { { Marked(marked) { actions.undo(row) } } }
        rowActions == RowActions.All -> { {
            Row(Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.spacedBy(FsSpace.space2)) {
                FlowButton(stringResource(R.string.raid_na_markIn), { actions.mark(row, SignupState.In) }, Modifier.weight(1f))
                FlowButton(stringResource(SignupState.Bench.label), { actions.mark(row, SignupState.Bench) }, Modifier.weight(1f))
                if (row.dmUrl != null) FlowButton(stringResource(R.string.raid_na_discordDM), { actions.message(row) }, Modifier.weight(1f))
            }
        } }
        rowActions == RowActions.MarkIn -> { { FlowButton(stringResource(R.string.raid_na_markIn), { actions.mark(row, SignupState.In) }) } }
        else -> null
    }
    FlowRow(row.name, row.sub, FlowLead.Crest(row.className), titleColor = classColor(row.className), last = last, trail = trail)
}

/** `Marked In · Undo`, the mark held on its row. */
@Composable
private fun Marked(state: SignupState, undo: () -> Unit) {
    Row {
        FlowNote(stringResource(R.string.raid_na_marked, stringResource(state.label)) + SEPARATOR)
        Text(stringResource(R.string.raid_na_undo), style = LINK, color = FsColor.gold, modifier = Modifier.clickable(role = Role.Button, onClick = undo))
    }
}

/** `Creating and editing raids stays on the site. Edit on the site ›`, 16 above. */
@Composable
private fun SiteLine(open: () -> Unit) {
    val site = stringResource(R.string.raid_na_site)
    val link = stringResource(R.string.raid_na_siteLink)
    Text(
        buildAnnotatedString {
            append(site)
            append(" ")
            withStyle(SpanStyle(color = FsColor.gold, fontWeight = FontWeight.SemiBold)) { append(link) }
        },
        style = FsType.small,
        color = FsColor.muted,
        modifier = Modifier.padding(top = FsSpace.space4).clickable(role = Role.Button, onClick = open),
    )
}

private const val SEPARATOR = " · "
private val PREVIEW = FsType.body.copy(fontSize = 14.sp, lineHeight = 20.sp)
private val PREVIEW_RULE = 3.dp
private val PREVIEW_PAD = 10.dp
private val LINK = FsType.small.copy(fontWeight = FontWeight.SemiBold)

/** The status filters at Activity's filter height (36, 48 to the touch), as the mock's chips. */
private val FILTER_HEIGHT = 36.dp
