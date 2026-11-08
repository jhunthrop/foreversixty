package gg.foreversixty.app.rankings.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.offset
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.layout.widthIn
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.clearAndSetSemantics
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.ParityItemKind
import gg.foreversixty.app.designsystem.ParityRegion
import gg.foreversixty.app.designsystem.parityItem
import gg.foreversixty.app.designsystem.parityRegion
import gg.foreversixty.app.flows.FlowGap
import gg.foreversixty.app.flows.FlowPlate
import gg.foreversixty.app.flows.FlowSection
import gg.foreversixty.app.flows.FlowSkeleton
import gg.foreversixty.app.flows.SUB
import gg.foreversixty.app.items.ItemIcon
import gg.foreversixty.app.items.ItemQuality
import gg.foreversixty.app.paperdoll.Paperdoll
import gg.foreversixty.app.rankings.logic.CellGain
import gg.foreversixty.app.rankings.logic.CharacterLink
import gg.foreversixty.app.rankings.logic.CharacterPageView
import gg.foreversixty.app.rankings.logic.GearCell

/** C4: the worn gear as the phone paperdoll, one cell a slot, the owner's gain to the band's pick on the right. */
@Composable
fun CharacterGear(view: CharacterPageView, build: String, actions: CharacterActions) {
    FlowSection(
        stringResource(R.string.char_gear),
        action = stringResource(R.string.char_page_bisAction),
        onAction = { actions.open(CharacterLink.Bis) },
        meta = view.gearMeta,
        top = FsSpace.space5,
    ) {
        FlowGap(FsSpace.space3)
        FlowPlate {
            Column(Modifier.fillMaxWidth().parityRegion(ParityRegion.SLOTS)) {
                view.cells.forEach { GearCellRow(it, build, actions) }
            }
        }
    }
}

@Composable
private fun GearCellRow(cell: GearCell, build: String, actions: CharacterActions) {
    val item = cell.slot.item
    val label = Paperdoll.label(cell.slot, LocalResources.current)
    val tap = if (item != null) Modifier.clickable(role = Role.Button, onClickLabel = label) { actions.openSocket(cell.slot) } else Modifier
    Row(Modifier.fillMaxWidth().parityItem(ParityItemKind.Row).then(tap).padding(start = FsSpace.space4, end = FsSpace.space4, top = CELL_TOP, bottom = CELL_BOTTOM)) {
        Box(Modifier.padding(top = ICON_DROP)) {
            if (item != null) {
                ItemIcon(item.icon, item.quality, build, ICON, Modifier.parityItem(ParityItemKind.Image))
            } else {
                Box(Modifier.size(ICON).border(1.dp, FsColor.lineSoft, RoundedCornerShape(ICON_RADIUS)))
            }
        }
        Spacer(Modifier.width(FsSpace.space3))
        Column(Modifier.weight(1f)) {
            Text(cell.label, style = SLOT_LABEL, color = FsColor.muted)
            Text(
                cell.slot.name,
                style = if (item == null) FsType.body.copy(lineHeight = 20.sp) else FsType.item,
                color = if (item == null) FsColor.muted else ItemQuality(item.quality).text,
            )
            cell.level?.let {
                Text(it, style = SUB, color = FsColor.muted)
            }
        }
        cell.gain?.let {
            Spacer(Modifier.width(FsSpace.space3))
            Gain(it)
        }
    }
}

/** `+18.4 HPS` over `vs Devout Crown` (the pick in its quality colour), or `Best in slot`. */
@Composable
private fun Gain(gain: CellGain) {
    Column(Modifier.widthIn(max = GAIN_WIDTH), horizontalAlignment = Alignment.End) {
        Text(
            buildAnnotatedString {
                append(gain.value)
                gain.unit?.takeIf { it.isNotEmpty() }?.let { withStyle(SpanStyle(fontSize = 13.sp, fontWeight = FontWeight.SemiBold, color = FsColor.muted)) { append(" $it") } }
            },
            style = FsType.figureSm,
            color = if (gain.muted) FsColor.muted else FsColor.strong,
            maxLines = 1,
        )
        gain.versus?.let { pick ->
            Text(
                buildAnnotatedString {
                    append(stringResource(R.string.shared_paperdoll_versus))
                    append(" ")
                    withStyle(SpanStyle(color = ItemQuality(pick.quality).text)) { append(pick.name.orEmpty()) }
                },
                style = VERSUS,
                color = FsColor.muted,
                maxLines = 1,
                overflow = TextOverflow.Ellipsis,
            )
        }
    }
}

/** The loading page (`.head-sk`, `.cell.sk`): the crest and name blocks, then seventeen cells at their real height. */
@Composable
fun CharacterLoading() {
    Box(Modifier.fillMaxWidth().height(SKELETON_HEAD).parityRegion(ParityRegion.HEAD).clearAndSetSemantics {}) {
        Box(Modifier.offset(y = SKELETON_CREST_TOP).size(SKELETON_CREST).background(FsColor.raised, CircleShape))
        FlowSkeleton(SKELETON_NAME_WIDTH, SKELETON_NAME_HEIGHT, Modifier.offset(x = SKELETON_CREST + FsSpace.space4, y = SKELETON_NAME_TOP))
    }
    FlowSection(stringResource(R.string.char_gear), top = FsSpace.space5) {
        FlowGap(FsSpace.space3)
        FlowPlate {
            Column(Modifier.fillMaxWidth().parityRegion(ParityRegion.SLOTS).clearAndSetSemantics {}) {
                repeat(Paperdoll.LEFT.size + Paperdoll.RIGHT.size + Paperdoll.WEAPONS.size) { SkeletonCell() }
            }
        }
    }
}

@Composable
private fun SkeletonCell() {
    Row(Modifier.fillMaxWidth().height(SKELETON_CELL).parityItem(ParityItemKind.Row).padding(horizontal = FsSpace.space4)) {
        FlowSkeleton(ICON, ICON, Modifier.padding(top = SKELETON_ICON_TOP))
        Spacer(Modifier.width(FsSpace.space3))
        Column(Modifier.padding(top = SKELETON_LABEL_TOP)) {
            FlowSkeleton(SKELETON_LABEL_WIDTH, SKELETON_LINE)
            FlowSkeleton(SKELETON_TITLE_WIDTH, SKELETON_TITLE)
            FlowSkeleton(SKELETON_SUB_WIDTH, SKELETON_LINE, Modifier.padding(top = SKELETON_SUB_GAP))
        }
    }
}

private val CELL_TOP = 9.dp
private val CELL_BOTTOM = 10.dp
private val ICON_DROP = 2.dp
private val ICON = 40.dp
private val ICON_RADIUS = 5.dp
private val GAIN_WIDTH = 104.dp
private val SLOT_LABEL = FsType.small.copy(fontSize = 12.sp, lineHeight = 17.sp, fontWeight = FontWeight.SemiBold)
private val VERSUS = FsType.small.copy(fontSize = 12.sp, lineHeight = 16.sp)
private val SKELETON_HEAD = 96.dp
private val SKELETON_CREST = 56.dp
private val SKELETON_CREST_TOP = 22.dp
private val SKELETON_NAME_TOP = 34.dp
private val SKELETON_NAME_WIDTH = 280.dp
private val SKELETON_NAME_HEIGHT = 32.dp
private val SKELETON_CELL = 68.dp
private val SKELETON_ICON_TOP = 15.dp
private val SKELETON_LABEL_TOP = 10.dp
private val SKELETON_LABEL_WIDTH = 48.dp
private val SKELETON_LINE = 12.dp
private val SKELETON_TITLE_WIDTH = 160.dp
private val SKELETON_TITLE = 16.dp
private val SKELETON_SUB_WIDTH = 224.dp
private val SKELETON_SUB_GAP = 8.dp
