package gg.foreversixty.app.planner.ui

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.selection.selectable
import androidx.compose.foundation.selection.selectableGroup
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Add
import androidx.compose.material.icons.filled.Check
import androidx.compose.material3.AlertDialog
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.OutlinedTextFieldDefaults
import androidx.compose.material3.Text
import androidx.compose.material3.TextButton
import androidx.compose.material3.rememberModalBottomSheetState
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.vector.rememberVectorPainter
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.Role
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.MicroLabel
import gg.foreversixty.app.model.RaceInfo
import gg.foreversixty.app.planner.logic.PlannerEditor
import gg.foreversixty.app.planner.logic.PlannerRules
import gg.foreversixty.app.planner.logic.PlannerShare
import java.util.Locale

/** The sheets' frame: the app's bottom sheet on the raised ground. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun PlannerSheet(onDismiss: () -> Unit, content: @Composable () -> Unit) {
    ModalBottomSheet(onDismissRequest = onDismiss, sheetState = rememberModalBottomSheetState(skipPartiallyExpanded = true), containerColor = FsColor.raised) {
        content()
    }
}

@Composable
private fun SheetTitle(text: String) {
    Text(text.uppercase(Locale.ROOT), style = FsType.sec, color = FsColor.strong, modifier = Modifier.semantics { heading() })
}

/** The Race chip's sheet: the class's legal races (combos.json), the current one ticked. */
@Composable
fun RaceSheetContent(races: List<RaceInfo>, current: String?, choose: (String) -> Unit) {
    Column(Modifier.fillMaxWidth().padding(start = 16.dp, end = 16.dp, bottom = 24.dp).verticalScroll(rememberScrollState())) {
        SheetTitle(stringResource(R.string.plan_raceSheet))
        Column(Modifier.padding(top = 8.dp).selectableGroup()) {
            races.forEach { race ->
                val selected = race.slug == current
                Row(
                    Modifier.fillMaxWidth().heightIn(min = 52.dp).selectable(selected, role = Role.RadioButton) { choose(race.slug) },
                    verticalAlignment = Alignment.CenterVertically,
                ) {
                    Text(race.name, style = FsType.body, color = if (selected) FsColor.strong else FsColor.text, modifier = Modifier.weight(1f))
                    if (selected) Icon(Icons.Filled.Check, contentDescription = null, tint = FsColor.gold, modifier = Modifier.size(20.dp))
                }
            }
        }
    }
}

/** The Level chip's sheet: a 10 to 60 stepper that never goes below the spent points' level. */
@Composable
fun LevelSheetContent(level: Int, minLevel: Int, change: (Int) -> Unit) {
    Column(Modifier.fillMaxWidth().padding(start = 16.dp, end = 16.dp, bottom = 24.dp), verticalArrangement = Arrangement.spacedBy(12.dp)) {
        SheetTitle(stringResource(R.string.plan_levelSheet))
        Row(verticalAlignment = Alignment.CenterVertically, horizontalArrangement = Arrangement.spacedBy(20.dp)) {
            StepButton(painterResource(R.drawable.ic_minus), stringResource(R.string.plan_level, level - 1), level > minLevel, { change(level - 1) })
            Text(level.toString(), style = FsType.monoDay, color = FsColor.strong)
            StepButton(rememberVectorPainter(Icons.Filled.Add), stringResource(R.string.plan_level, level + 1), level < PlannerEditor.MAX_LEVEL, { change(level + 1) })
        }
        Text(stringResource(R.string.plan_levelSheetLine, PlannerRules.pointsAvailable(level)), style = FsType.meta, color = FsColor.muted)
    }
}

/** Reset asks once: `Clear all {n} points?`, Clear in the error colour, and Cancel. */
@Composable
fun ResetDialog(points: Int, clear: () -> Unit, cancel: () -> Unit) {
    AlertDialog(
        onDismissRequest = cancel,
        containerColor = FsColor.raised,
        title = { Text(stringResource(R.string.plan_resetConfirm, points), style = FsType.lead, color = FsColor.strong) },
        confirmButton = { TextButton(onClick = clear) { Text(stringResource(R.string.plan_resetProceed), style = FsType.bodyStrong, color = FsColor.wipe) } },
        dismissButton = { TextButton(onClick = cancel) { Text(stringResource(R.string.common_cancel), style = FsType.bodyStrong, color = FsColor.text) } },
    )
}

/** The share sheet's state: the typed title and whether the POST is in flight. */
data class ShareSheetState(val title: String, val sharing: Boolean, val signedIn: Boolean)

/** F13.3: the web's confirm step, verbatim (spec F13.3). */
@Composable
fun ShareSheetContent(state: ShareSheetState, onTitle: (String) -> Unit, proceed: () -> Unit, cancel: () -> Unit) {
    Column(
        Modifier.fillMaxWidth().padding(start = 16.dp, end = 16.dp, bottom = 24.dp).verticalScroll(rememberScrollState()),
        verticalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        SheetTitle(stringResource(R.string.plan_shareTitle))
        Text(stringResource(R.string.plan_shareIntro), style = FsType.body, color = FsColor.text)
        val title = PlannerShare.cleanTitle(state.title)
        Column(verticalArrangement = Arrangement.spacedBy(4.dp)) {
            Bullet(stringResource(R.string.plan_shareFieldClassRace))
            Bullet(stringResource(R.string.plan_shareFieldTalents))
            title?.let { Bullet(stringResource(R.string.plan_shareFieldTitle, it)) }
        }
        MicroLabel(stringResource(R.string.plan_shareTitleField))
        OutlinedTextField(
            value = state.title,
            onValueChange = { onTitle(it.take(PlannerShare.TITLE_MAX)) },
            singleLine = true,
            textStyle = FsType.body,
            label = null,
            trailingIcon = {
                Text(
                    stringResource(R.string.plan_shareTitleCount, state.title.length, PlannerShare.TITLE_MAX),
                    style = FsType.monoMeta, color = FsColor.muted, modifier = Modifier.padding(end = 12.dp),
                )
            },
            colors = OutlinedTextFieldDefaults.colors(
                focusedTextColor = FsColor.strong, unfocusedTextColor = FsColor.strong,
                focusedBorderColor = FsColor.gold, unfocusedBorderColor = FsColor.line, cursorColor = FsColor.gold,
            ),
            modifier = Modifier.fillMaxWidth(),
        )
        BrandButton(stringResource(R.string.plan_shareProceed), BrandButtonKind.Primary, proceed, busy = state.sharing, height = 48.dp)
        BrandButton(stringResource(R.string.common_cancel), BrandButtonKind.Secondary, cancel, enabled = !state.sharing, height = 48.dp)
        if (state.signedIn) Text(stringResource(R.string.plan_shareSignedIn), style = FsType.meta, color = FsColor.muted)
    }
}

@Composable
private fun Bullet(text: String) {
    Row {
        Text(BULLET, style = FsType.body, color = FsColor.muted)
        Spacer(Modifier.width(6.dp))
        Text(text, style = FsType.body, color = FsColor.text)
    }
}

private const val BULLET = "·"
