package gg.foreversixty.app.sims.ui

import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.ExperimentalLayoutApi
import androidx.compose.foundation.layout.FlowRow
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.GenericShape
import androidx.compose.foundation.background
import androidx.compose.material3.Icon
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.dp
import androidx.compose.ui.unit.sp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.MicroLabel
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.designsystem.classColor
import gg.foreversixty.app.items.ItemIcon
import gg.foreversixty.app.items.ItemVocabulary
import gg.foreversixty.app.sims.logic.SimGearCell
import gg.foreversixty.app.sims.logic.SimGearLook
import gg.foreversixty.app.sims.logic.SimResultModel
import gg.foreversixty.app.sims.logic.SimResultPhase
import java.time.Instant
import java.util.Locale

object SimResultLayout {
    const val GRID_COLUMNS = 6
    val gearIcon = 40.dp
    val gridGap = 8.dp
    val figureSize = 52.sp

    /** The game's question-mark icon, for an id the build's items file does not list. */
    const val UNKNOWN_ICON = "inv_misc_questionmark"
}

/** F5.3 as drawn from a phase: stateless, so screenshots drive every state. */
@Composable
fun SimResultContent(
    phase: SimResultPhase,
    looks: Map<Int, SimGearLook>,
    build: String?,
    now: Instant,
    openItem: (SimGearCell) -> Unit = {},
    openBreakdown: (String) -> Unit = {},
    retry: () -> Unit = {},
) {
    Column(Modifier.fillMaxWidth().padding(bottom = 24.dp)) {
        when (phase) {
            SimResultPhase.Loading -> { ScreenTitle(); ResultSkeleton() }
            SimResultPhase.Failed -> { ScreenTitle(); ErrorBlock(stringResource(R.string.sims_resultFailed), retry) }
            is SimResultPhase.Loaded -> {
                ScreenTitle(phase.model.title)
                phase.offlineSince?.let { OfflineStrip(stringResource(R.string.offline_strip, OfflineStamp.format(it, now))) }
                Loaded(phase.model, looks, build, openItem, openBreakdown)
            }
        }
    }
}

@Composable
private fun Loaded(model: SimResultModel, looks: Map<Int, SimGearLook>, build: String?, openItem: (SimGearCell) -> Unit, openBreakdown: (String) -> Unit) {
    val res = LocalResources.current
    Column(Modifier.padding(horizontal = 16.dp).padding(top = 4.dp)) {
        MicroLabel(model.kind.uppercase(Locale.ROOT))
        Text(
            buildAnnotatedString {
                append("${model.specName} ")
                withStyle(SpanStyle(color = classColor(model.classSlug))) { append(model.className) }
                model.level?.let { append(" · " + res.getString(R.string.sims_level, it)) }
            },
            style = FsType.meta, color = FsColor.muted, modifier = Modifier.padding(top = 4.dp),
        )
        Answer(model.answer)
        Text(model.errorLine, style = FsType.monoMeta, color = FsColor.muted, modifier = Modifier.padding(top = 4.dp))
        if (model.gear.isNotEmpty()) Gear(model, looks, build, openItem)
        Row(
            Modifier.padding(top = 8.dp).heightIn(min = 48.dp).clickable { openBreakdown(model.breakdownUrl) },
            verticalAlignment = Alignment.CenterVertically,
            horizontalArrangement = Arrangement.spacedBy(6.dp),
        ) {
            Text(stringResource(R.string.sims_breakdown), style = FsType.bodyStrong, color = FsColor.gold, modifier = Modifier.weight(1f, fill = false))
            Icon(painterResource(R.drawable.ic_open_external), contentDescription = null, tint = FsColor.gold, modifier = Modifier.size(16.dp))
        }
    }
}

@OptIn(ExperimentalLayoutApi::class)
@Composable
private fun Answer(answer: SimResultModel.Answer) {
    Box(Modifier.padding(top = 12.dp)) {
        when (answer) {
            is SimResultModel.Answer.Figure -> FlowRow(
                verticalArrangement = Arrangement.Bottom,
                horizontalArrangement = Arrangement.spacedBy(6.dp),
                itemVerticalAlignment = Alignment.Bottom,
            ) {
                Text(answer.value, style = FsType.title.copy(fontSize = SimResultLayout.figureSize), color = FsColor.strong)
                Text(stringResource(R.string.bis_unit_dps), style = FsType.body, color = FsColor.muted, modifier = Modifier.padding(bottom = 8.dp))
            }
            is SimResultModel.Answer.Sentence -> Text(answer.text, style = FsType.stat, color = FsColor.strong)
        }
    }
}

@OptIn(ExperimentalLayoutApi::class)
@Composable
private fun Gear(model: SimResultModel, looks: Map<Int, SimGearLook>, build: String?, openItem: (SimGearCell) -> Unit) {
    val res = LocalResources.current
    Column(Modifier.padding(top = 22.dp), verticalArrangement = Arrangement.spacedBy(8.dp)) {
        MicroLabel(stringResource(R.string.sims_gear).uppercase(Locale.ROOT))
        FlowRow(
            Modifier.width(SimResultLayout.gearIcon * SimResultLayout.GRID_COLUMNS + SimResultLayout.gridGap * (SimResultLayout.GRID_COLUMNS - 1)),
            horizontalArrangement = Arrangement.spacedBy(SimResultLayout.gridGap),
            verticalArrangement = Arrangement.spacedBy(SimResultLayout.gridGap),
            maxItemsInEachRow = SimResultLayout.GRID_COLUMNS,
        ) {
            model.gear.forEach { cell ->
                val look = looks[cell.itemId]
                val slot = ItemVocabulary.slotLabel(res, cell.slot)
                val label = listOfNotNull(look?.name, slot, if (cell.changed) res.getString(R.string.sims_changedLabel) else null).joinToString(", ")
                val tap = if (look != null) Modifier.clickable { openItem(cell) } else Modifier
                Box(Modifier.then(tap).semantics { contentDescription = label }) {
                    ItemIcon(look?.icon ?: SimResultLayout.UNKNOWN_ICON, look?.quality, build.orEmpty(), SimResultLayout.gearIcon)
                    if (cell.changed) ChangedCorner(Modifier.align(Alignment.TopEnd))
                }
            }
        }
        if (model.showsChangedLegend) Text(stringResource(R.string.sims_changed), style = FsType.meta, color = FsColor.muted)
    }
}

private val CornerShape = GenericShape { size, _ ->
    moveTo(0f, 0f)
    lineTo(size.width, 0f)
    lineTo(size.width, size.height)
    close()
}

/** The 12 dp gold corner on a slot the result changes. */
@Composable
private fun ChangedCorner(modifier: Modifier) {
    Box(modifier.size(12.dp).background(FsColor.gold, CornerShape))
}

@Composable
private fun ResultSkeleton() {
    Column(Modifier.padding(16.dp), verticalArrangement = Arrangement.spacedBy(10.dp)) {
        Bar(80.dp, 12.dp)
        Bar(160.dp, 14.dp)
        Bar(140.dp, 52.dp)
        Bar(240.dp, 12.dp)
        Row(Modifier.padding(top = 16.dp), horizontalArrangement = Arrangement.spacedBy(SimResultLayout.gridGap)) {
            repeat(SimResultLayout.GRID_COLUMNS) { Bar(SimResultLayout.gearIcon, SimResultLayout.gearIcon) }
        }
    }
}
