package gg.foreversixty.app.guides.ui

import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.ClassCrest
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.MicroLabel
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.designsystem.primaryCard
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.guides.logic.ClassCell
import gg.foreversixty.app.guides.logic.GuideDates
import gg.foreversixty.app.guides.logic.Guides
import gg.foreversixty.app.guides.logic.GuidesCopy
import gg.foreversixty.app.guides.logic.Loadable
import gg.foreversixty.app.items.SpecIcon
import gg.foreversixty.app.model.GuideIndex
import gg.foreversixty.app.model.GuideIndexSpec
import gg.foreversixty.app.reports.ui.Chevron
import gg.foreversixty.app.reports.ui.ErrorRegion
import gg.foreversixty.app.reports.ui.SkeletonBar
import gg.foreversixty.app.reports.ui.isLargeText
import gg.foreversixty.app.reports.ui.scaledSize
import java.util.Locale

/** What the Guides segment draws (spec F9.1). */
data class GuidesHomeState(
    val index: Loadable<GuideIndex>,
    /** Saved guides by spec id, with their bytes. */
    val saved: Map<String, Long>,
    /** The current character's guide; null with no character (the card is then absent). */
    val yours: GuideIndexSpec?,
)

class GuidesHomeActions(
    val openGuide: (String) -> Unit = {},
    val openClass: (String) -> Unit = {},
    val retry: () -> Unit = {},
)

/** F9.1: `Your guide` (the one warm card), `All classes`, and what is saved on this phone. */
@Composable
fun GuidesHome(state: GuidesHomeState, actions: GuidesHomeActions) {
    when (val index = state.index) {
        Loadable.Loading -> GuidesHomeSkeleton()
        Loadable.NotCached -> Text(stringResource(R.string.guides_listNotCached), style = FsType.body, color = FsColor.muted, modifier = Modifier.screenMargins().padding(top = GuideSpace.section))
        Loadable.Failed -> ErrorRegion(stringResource(R.string.guides_error), actions.retry)
        is Loadable.Loaded -> Column(Modifier.screenMargins()) {
            state.yours?.let { spec ->
                val cls = Guides.classOf(index.value, spec.spec)
                YourGuideCard(spec, cls?.classSlug, cls?.className.orEmpty(), index.value.build, saved = spec.spec in state.saved) { actions.openGuide(spec.spec) }
            }
            GuideSectionBlock(stringResource(R.string.guides_classes)) { ClassGrid(Guides.classCells(index.value), actions.openClass) }
            val copy = GuidesCopy(LocalResources.current)
            Text(copy.footer(Guides.footer(index.value, state.saved)), style = FsType.meta, color = FsColor.muted, modifier = Modifier.padding(top = GuideSpace.rowGap, bottom = GuideSpace.section))
        }
    }
}

@Composable
private fun YourGuideCard(spec: GuideIndexSpec, classSlug: String?, className: String, build: String, saved: Boolean, open: () -> Unit) {
    val copy = GuidesCopy(LocalResources.current)
    Column(
        Modifier.fillMaxWidth().padding(top = GuideSpace.labelGap).primaryCard().heightIn(min = CARD_MIN)
            .clickable(role = Role.Button, onClick = open).cardInside(),
        verticalArrangement = Arrangement.spacedBy(GuideSpace.rowGap),
    ) {
        Row(horizontalArrangement = Arrangement.spacedBy(GuideSpace.inCard), verticalAlignment = Alignment.CenterVertically) {
            SpecIcon(spec.icon, build, scaledSize(HEAD_ICON, HEAD_ICON_CAP))
            Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
                MicroLabel(stringResource(R.string.guides_yours), color = FsColor.gold)
                Text(specTitle(spec.specName, className).uppercase(Locale.ROOT), style = FsType.h2.copy(fontSize = FsType.cardHead.fontSize), color = classColor(classSlug))
                Text(copy.cardMeta(spec.role, GuideDates.long(spec.updated)), style = FsType.meta, color = FsColor.muted)
            }
            Chevron(CHEVRON)
        }
        if (saved) SavedMark(saved = true, long = true)
    }
}

/** Three crests a row (two at the largest text), each with its class name in its colour. */
@Composable
private fun ClassGrid(cells: List<ClassCell>, open: (String) -> Unit) {
    val columns = if (isLargeText()) LARGE_COLUMNS else COLUMNS
    Column(Modifier.fillMaxWidth().quietCard().padding(vertical = GuideSpace.inCard)) {
        cells.chunked(columns).forEach { row ->
            Row(Modifier.fillMaxWidth()) {
                row.forEach { cell -> ClassCellView(cell, Modifier.weight(1f), open) }
                repeat(columns - row.size) { Spacer(Modifier.weight(1f)) }
            }
        }
    }
}

@Composable
private fun ClassCellView(cell: ClassCell, modifier: Modifier, open: (String) -> Unit) {
    Column(
        modifier.heightIn(min = CELL_MIN).clickable(role = Role.Button) { open(cell.classSlug) }.padding(vertical = GuideSpace.labelGap)
            .semantics { contentDescription = cell.className },
        horizontalAlignment = Alignment.CenterHorizontally,
        verticalArrangement = Arrangement.spacedBy(6.dp),
    ) {
        ClassCrest(cell.classSlug, CREST)
        Text(cell.className, style = FsType.meta.copy(fontWeight = androidx.compose.ui.text.font.FontWeight.SemiBold), color = classColor(cell.classSlug))
    }
}

/** F9.7: skeletons at the real heights of the card and the grid; the segments stay live. */
@Composable
private fun GuidesHomeSkeleton() {
    Column(Modifier.screenMargins().padding(top = GuideSpace.labelGap), verticalArrangement = Arrangement.spacedBy(GuideSpace.section)) {
        Column(Modifier.fillMaxWidth().quietCard().heightIn(min = CARD_MIN).cardInside(), verticalArrangement = Arrangement.spacedBy(10.dp)) {
            SkeletonBar(120.dp, 12.dp)
            SkeletonBar(220.dp, 20.dp)
            SkeletonBar(160.dp, 12.dp)
        }
        Box(Modifier.fillMaxWidth().quietCard().height(GRID_SKELETON))
        Text(stringResource(R.string.guides_loading), style = FsType.meta, color = FsColor.muted)
    }
}

/** `Elemental Shaman`. */
internal fun specTitle(spec: String, className: String) = "$spec $className"

private val CARD_MIN = 80.dp
private val HEAD_ICON = 56.dp
private val HEAD_ICON_CAP = 72.dp
private val CHEVRON = 24.dp
private val CREST = 44.dp
private val CELL_MIN = 96.dp
private val GRID_SKELETON = 330.dp
private const val COLUMNS = 3
private const val LARGE_COLUMNS = 2
