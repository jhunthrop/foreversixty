package gg.foreversixty.app.sims.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
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
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.KeyboardArrowRight
import androidx.compose.material3.Icon
import androidx.compose.material3.LinearProgressIndicator
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalDensity
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.semantics.stateDescription
import androidx.compose.ui.text.SpanStyle
import androidx.compose.ui.text.buildAnnotatedString
import androidx.compose.ui.text.style.TextDecoration
import androidx.compose.ui.text.withStyle
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsRadius
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.items.SpecIcon
import gg.foreversixty.app.model.SimProgress
import gg.foreversixty.app.model.SimRow
import gg.foreversixty.app.sims.logic.RerunStatus
import gg.foreversixty.app.sims.logic.ServerRunState
import gg.foreversixty.app.sims.logic.ServerRuns
import gg.foreversixty.app.sims.logic.SimsCopy
import java.time.Instant

/** The words and marks one row draws, worked out once for both sections. */
class SimRowFace(val row: SimRow, specs: SpecCatalog, now: Instant, copy: SimsCopy) {
    val state: ServerRunState = ServerRuns.state(row, now)
    private val spec = specs.spec(row.spec.orEmpty())
    val specName: String = spec?.name ?: row.spec.orEmpty()
    val icon: String? = spec?.icon
    val kind: String = copy.kindLabel(row.kind)
    val title: String = row.title.orEmpty().ifEmpty { specName }

    /** `{Spec} · {Kind}` with a title, `{Kind}` without (the site's row meta). */
    val meta: String = if (row.title.isNullOrEmpty()) kind else "$specName · $kind"
    val headline: String = row.headline.orEmpty()
    val isUnseen: Boolean get() = state == ServerRunState.DoneUnseen
}

object SimsLayout {
    val serverRowHeight = 88.dp
    val mineRowHeight = 72.dp
    val serverIcon = 36.dp
    val mineIcon = 32.dp
    val runAgainHeight = 48.dp

    /** From this font scale the age stacks under the state (spec F5 Largest text). */
    const val STACK_SCALE = 1.3f

    /** A headline that is a plain figure (`1,043 DPS`) sets in mono; a sentence does not. */
    fun isFigure(headline: String): Boolean = headline.firstOrNull()?.isDigit() == true
}

@Composable
fun SimSpecMark(icon: String?, build: String?, size: Dp) {
    if (icon != null && build != null) {
        SpecIcon(icon, build, size)
    } else {
        Box(Modifier.size(size).background(FsColor.cardTop, RoundedCornerShape(FsRadius.pill)))
    }
}

@Composable
private fun SimRowTitle(face: SimRowFace) {
    Row(verticalAlignment = Alignment.CenterVertically, horizontalArrangement = Arrangement.spacedBy(8.dp)) {
        if (face.isUnseen) Box(Modifier.size(8.dp).background(FsColor.gold, CircleShape))
        Text(face.title, style = FsType.bodyStrong, color = FsColor.strong, maxLines = 2)
    }
}

@Composable
fun SimStateText(face: SimRowFace, copy: SimsCopy, singleLine: Boolean = false, onOverflow: () -> Unit = {}) {
    val lines = if (singleLine) 1 else Int.MAX_VALUE
    val layout: (androidx.compose.ui.text.TextLayoutResult) -> Unit = { if (it.hasVisualOverflow) onOverflow() }
    when (face.state) {
        ServerRunState.Queued -> Row(verticalAlignment = Alignment.CenterVertically, horizontalArrangement = Arrangement.spacedBy(6.dp)) {
            Box(Modifier.size(8.dp).background(FsColor.gold, CircleShape))
            Text(stringResource(R.string.sims_queued), style = FsType.meta, color = FsColor.strong, maxLines = lines, onTextLayout = layout)
        }
        ServerRunState.Running, ServerRunState.RunningOverdue, ServerRunState.RunningNoEstimate -> Text(
            buildAnnotatedString {
                withStyle(SpanStyle(color = FsColor.gold)) { append(stringResource(R.string.sims_running)) }
                withStyle(SpanStyle(color = FsColor.muted)) { append(copy.stateSuffix(face.row, face.state)) }
            },
            style = FsType.meta, maxLines = lines, softWrap = !singleLine, onTextLayout = layout,
        )
        ServerRunState.Done, ServerRunState.DoneUnseen ->
            Text(face.headline, style = if (SimsLayout.isFigure(face.headline)) FsType.mono else FsType.bodyStrong, color = FsColor.strong, maxLines = lines, softWrap = !singleLine, onTextLayout = layout)
        ServerRunState.Failed -> Text(copy.failedLine(face.row), style = FsType.meta, color = FsColor.wipe, maxLines = lines, softWrap = !singleLine, onTextLayout = layout)
    }
}

/** State left and age right; at large text the age stacks under the state. */
@Composable
private fun SimStateLine(face: SimRowFace, copy: SimsCopy, now: Instant) {
    val age = @Composable { Text(copy.ageLine(face.row, face.state, now), style = FsType.monoMeta, color = FsColor.muted) }
    // The age drops under the state when the two do not fit on one line (spec F5.1 item 3).
    var overflows by remember(face.row, now) { mutableStateOf(false) }
    if (overflows || LocalDensity.current.fontScale >= SimsLayout.STACK_SCALE) {
        Column(verticalArrangement = Arrangement.spacedBy(2.dp)) { SimStateText(face, copy); age() }
    } else {
        Row(verticalAlignment = Alignment.CenterVertically) {
            Box(Modifier.weight(1f)) { SimStateText(face, copy, singleLine = true) { overflows = true } }
            Spacer(Modifier.width(8.dp))
            age()
        }
    }
}

/** A bulk run's 8 dp bar and `40 of 64 combinations` (only when the progress read counts them). */
@Composable
private fun SimProgressBar(progress: SimProgress) {
    val fraction = ServerRuns.fraction(progress) ?: return
    val words = stringResource(R.string.sims_combos, progress.combosDone, progress.combosTotal)
    Column(Modifier.padding(top = 8.dp).semantics { stateDescription = words }, verticalArrangement = Arrangement.spacedBy(4.dp)) {
        LinearProgressIndicator(
            progress = { fraction },
            modifier = Modifier.fillMaxWidth().heightIn(min = 8.dp, max = 8.dp),
            color = FsColor.gold,
            trackColor = FsColor.lineSoft,
            gapSize = 0.dp,
            drawStopIndicator = {},
        )
        Text(words, style = FsType.monoMeta, color = FsColor.muted)
    }
}

@Composable
private fun SimChevron() {
    Icon(Icons.AutoMirrored.Filled.KeyboardArrowRight, contentDescription = null, tint = FsColor.muted, modifier = Modifier.size(24.dp))
}

/** One Server runs row (spec F5.1): done rows open the result; a failed row's one control is Run again. */
@Composable
fun ServerRunRowView(
    face: SimRowFace,
    copy: SimsCopy,
    now: Instant,
    build: String?,
    progress: SimProgress?,
    rerun: RerunStatus,
    offline: Boolean,
    mayLinkToPlans: Boolean,
    actions: SimsActions,
) {
    val res = LocalResources.current
    val label = SimsAccessibility.serverRow(face, copy, now, res)
    val tap = if (face.state.isDone) Modifier.clickable(onClickLabel = stringResource(R.string.sims_openHint)) { actions.open(face.row) } else Modifier
    Row(
        modifier = Modifier.fillMaxWidth().then(tap).heightIn(min = SimsLayout.serverRowHeight).padding(horizontal = 16.dp, vertical = 12.dp),
        horizontalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        SimSpecMark(face.icon, build, SimsLayout.serverIcon)
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
            Column(Modifier.semantics(mergeDescendants = true) { contentDescription = label }) {
                SimRowTitle(face)
                Text(face.meta, style = FsType.meta, color = FsColor.muted)
                Box(Modifier.padding(top = 2.dp)) { SimStateLine(face, copy, now) }
            }
            if (progress != null && face.state.isInFlight) SimProgressBar(progress)
            if (face.state == ServerRunState.Failed) FailedControls(face, rerun, offline, mayLinkToPlans, actions)
        }
        if (face.state.isDone) Box(Modifier.heightIn(min = 36.dp), contentAlignment = Alignment.Center) { SimChevron() }
    }
}

@Composable
private fun FailedControls(face: SimRowFace, rerun: RerunStatus, offline: Boolean, mayLinkToPlans: Boolean, actions: SimsActions) {
    val title = stringResource(
        when (rerun) {
            RerunStatus.Pending -> R.string.sims_runAgainPending
            RerunStatus.Started -> R.string.sims_runAgainDone
            else -> R.string.sims_runAgainRow
        },
    )
    val description = "$title, ${face.title}"
    Column(Modifier.padding(top = 8.dp), verticalArrangement = Arrangement.spacedBy(4.dp)) {
        BrandButton(
            title, BrandButtonKind.Secondary, { face.row.simId?.let(actions.runAgain) },
            modifier = Modifier.semantics { contentDescription = description },
            enabled = !rerun.isLocked && !offline, height = SimsLayout.runAgainHeight,
        )
        when {
            offline -> Text(stringResource(R.string.sims_offline), style = FsType.meta, color = FsColor.muted)
            rerun == RerunStatus.Failed -> Text(stringResource(R.string.sims_runAgainFailed), style = FsType.meta, color = FsColor.wipe)
            rerun == RerunStatus.Lapsed -> LapsedLine(mayLinkToPlans, actions.openPlans)
        }
    }
}

@Composable
private fun LapsedLine(mayLinkToPlans: Boolean, openPlans: () -> Unit) {
    Row(verticalAlignment = Alignment.CenterVertically, horizontalArrangement = Arrangement.spacedBy(6.dp)) {
        Text(stringResource(R.string.sims_premiumLine), style = FsType.meta, color = FsColor.wipe, modifier = Modifier.weight(1f, fill = false))
        if (mayLinkToPlans) {
            Text(
                buildAnnotatedString { withStyle(SpanStyle(textDecoration = TextDecoration.Underline)) { append(stringResource(R.string.you_seePlans)) } },
                style = FsType.meta.copy(fontWeight = FsType.bodyStrong.fontWeight),
                color = FsColor.gold,
                modifier = Modifier.heightIn(min = SimsLayout.runAgainHeight).clickable(onClick = openPlans).padding(vertical = 14.dp),
            )
        }
    }
}

/** One Your sims row: the headline (or the state while in flight), and where it ran. */
@Composable
fun MySimRowView(face: SimRowFace, copy: SimsCopy, build: String?, open: (SimRow) -> Unit) {
    val tap = if (face.state.isDone) Modifier.clickable(onClickLabel = stringResource(R.string.sims_openHint)) { open(face.row) } else Modifier
    Row(
        modifier = Modifier.fillMaxWidth().then(tap).heightIn(min = SimsLayout.mineRowHeight).padding(horizontal = 16.dp, vertical = 12.dp)
            .semantics(mergeDescendants = true) {},
        horizontalArrangement = Arrangement.spacedBy(12.dp),
    ) {
        SimSpecMark(face.icon, build, SimsLayout.mineIcon)
        Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(2.dp)) {
            SimRowTitle(face)
            SimStateText(face, copy)
            Text(copy.origin(face.row), style = FsType.meta, color = FsColor.muted)
        }
        if (face.state.isDone) Box(Modifier.heightIn(min = 36.dp), contentAlignment = Alignment.Center) { SimChevron() }
    }
}

/** The row labels TalkBack reads (spec F5 Accessibility). */
object SimsAccessibility {
    fun serverRow(face: SimRowFace, copy: SimsCopy, now: Instant, res: android.content.res.Resources): String {
        val parts = mutableListOf(face.title, face.specName)
        when {
            face.state.isDone -> {
                parts += face.headline
                parts += copy.finishedLine(face.row, now)
                if (face.isUnseen) parts += res.getString(R.string.sims_unread)
            }
            face.state == ServerRunState.Failed -> {
                parts += copy.failedLine(face.row)
                parts += copy.ageLine(face.row, face.state, now)
            }
            else -> parts += res.getString(R.string.sims_running) + copy.stateSuffix(face.row, face.state)
        }
        return parts.filter { it.isNotEmpty() }.joinToString(", ")
    }
}
