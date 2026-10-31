package gg.foreversixty.app.talents

import androidx.compose.foundation.background
import androidx.compose.foundation.border
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.navigationBarsPadding
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.Text
import androidx.compose.material3.rememberModalBottomSheetState
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.GameColors
import gg.foreversixty.app.designsystem.MicroLabel

/** From font scale 1.5 the sheet opens full height (spec F6 §7, point 3). */
private const val FULL_SHEET_SCALE = 1.5f

/**
 * The rank sheet (spec F6 §7): the game's tooltip content in a bottom sheet, so it never clips.
 * Partly expanded then expanded at default text; full height from 1.5x, scrolling inside.
 */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun TalentRankSheet(model: TalentSheetModel, build: String, onDismiss: () -> Unit) {
    val full = LocalDensity.current.fontScale >= FULL_SHEET_SCALE
    ModalBottomSheet(
        onDismissRequest = onDismiss,
        sheetState = rememberModalBottomSheetState(skipPartiallyExpanded = full),
        containerColor = FsColor.raised,
    ) {
        TalentRankContent(model, build)
    }
}

/** The sheet's body, pure so the screenshot tests draw it directly. */
@Composable
fun TalentRankContent(model: TalentSheetModel, build: String) {
    Column(
        Modifier
            .fillMaxWidth()
            .background(FsColor.raised)
            .verticalScroll(rememberScrollState())
            .padding(start = 20.dp, end = 20.dp, top = 8.dp, bottom = 16.dp)
            .navigationBarsPadding(),
        verticalArrangement = Arrangement.spacedBy(16.dp),
    ) {
        Row(horizontalArrangement = Arrangement.spacedBy(12.dp), verticalAlignment = Alignment.Top) {
            val face = if (model.state == TalentCellState.Locked) TalentCellState.Open else model.state
            TalentIconTile(model.talent, face, build, size = 56.dp)
            Column(verticalArrangement = Arrangement.spacedBy(2.dp)) {
                Text(model.talent.name, style = FsType.name, color = FsColor.strong, modifier = Modifier.semantics { heading() })
                Text(stringResource(R.string.talents_rank, model.rank, model.talent.maxRank), style = FsType.monoMeta, color = FsColor.muted)
            }
        }
        val panel = RoundedCornerShape(FsRadius.control)
        Column(
            Modifier.fillMaxWidth().background(GameColors.tooltipPanel, panel).border(1.dp, GameColors.iconLine, panel).padding(horizontal = 14.dp, vertical = 12.dp),
            verticalArrangement = Arrangement.spacedBy(12.dp),
        ) {
            model.current?.let { Text(it, style = FsType.body, color = FsColor.text) }
            model.next?.let { next ->
                Column(verticalArrangement = Arrangement.spacedBy(4.dp)) {
                    MicroLabel(stringResource(R.string.talents_nextRank))
                    Text(next, style = FsType.body, color = FsColor.muted)
                }
            }
            model.requirements.forEach { Text(it, style = FsType.body, color = FsColor.wipe) }
        }
        model.guideLine?.let { line ->
            Text(
                line,
                style = FsType.bodyStrong,
                color = FsColor.strong,
                modifier = Modifier
                    .fillMaxWidth()
                    .dashedBorder(if (model.isOver) FsColor.wipe else FsColor.gold, FsRadius.control, 1.5.dp)
                    .padding(horizontal = 12.dp, vertical = 10.dp),
            )
        }
    }
}
