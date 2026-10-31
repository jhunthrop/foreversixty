package gg.foreversixty.app.sims.ui

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Refresh
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.LiveRegionMode
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.liveRegion
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.Dp
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.common.OfflineStamp
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.MicroLabel
import gg.foreversixty.app.designsystem.OfflineStrip
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.model.SimProgress
import gg.foreversixty.app.model.SimRow
import gg.foreversixty.app.sims.logic.RerunStatus
import gg.foreversixty.app.sims.logic.ServerRuns
import gg.foreversixty.app.sims.logic.SimsContent
import gg.foreversixty.app.sims.logic.SimsCopy
import gg.foreversixty.app.sims.logic.SimsPhase
import java.time.Instant
import java.util.Locale

/** What the list hands back to its screen. */
data class SimsActions(
    val open: (SimRow) -> Unit = {},
    val runAgain: (String) -> Unit = {},
    val retry: () -> Unit = {},
    val loadMore: () -> Unit = {},
    val openPlans: () -> Unit = {},
)

/** What the list draws besides its phase. */
data class SimsListState(
    val entitled: Boolean,
    val progress: Map<String, SimProgress> = emptyMap(),
    val rerun: Map<String, RerunStatus> = emptyMap(),
    val pageFailed: Boolean = false,
    val mayLinkToPlans: Boolean = false,
)

/** F5.1 as drawn from a phase: stateless, so screenshots drive every state. */
@Composable
fun SimsList(phase: SimsPhase, state: SimsListState, specs: SpecCatalog, build: String?, now: Instant, actions: SimsActions = SimsActions()) {
    val copy = SimsCopy(LocalResources.current)
    Column(Modifier.fillMaxWidth().padding(bottom = 24.dp)) {
        when (phase) {
            SimsPhase.Loading -> { ScreenTitle(); SimsSkeleton() }
            SimsPhase.Failed -> { ScreenTitle(); ErrorBlock(stringResource(R.string.sims_failedLoad), actions.retry) }
            is SimsPhase.Loaded -> {
                ScreenTitle()
                phase.content.offlineSince?.let { OfflineStrip(stringResource(R.string.offline_strip, OfflineStamp.format(it, now))) }
                Loaded(phase.content, state, specs, build, now, copy, actions)
            }
        }
    }
}

@Composable
fun ScreenTitle(text: String = stringResource(R.string.sims_title)) {
    Text(
        text.uppercase(Locale.ROOT), style = FsType.title, color = FsColor.strong,
        modifier = Modifier.padding(start = 16.dp, end = 16.dp, top = 16.dp, bottom = 4.dp).semantics { heading() },
    )
}

@Composable
private fun Loaded(content: SimsContent, state: SimsListState, specs: SpecCatalog, build: String?, now: Instant, copy: SimsCopy, actions: SimsActions) {
    Column(Modifier.padding(horizontal = 16.dp).padding(top = 12.dp), verticalArrangement = Arrangement.spacedBy(22.dp)) {
        if (ServerRuns.showsSection(state.entitled, content.serverRows) || content.serverFailed) {
            ServerSection(content, state, specs, build, now, copy, actions)
        }
        MineSection(content, state, specs, build, now, copy, actions)
        if (content.mine.isNotEmpty()) {
            Text(
                stringResource(if (state.entitled) R.string.sims_footer_premium else R.string.sims_footer_free),
                style = FsType.meta, color = FsColor.muted,
            )
        }
    }
}

@Composable
private fun ServerSection(content: SimsContent, state: SimsListState, specs: SpecCatalog, build: String?, now: Instant, copy: SimsCopy, actions: SimsActions) {
    Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
        MicroLabel(stringResource(R.string.sims_server).uppercase(Locale.ROOT))
        when {
            content.serverFailed -> Box(Modifier.fillMaxWidth().quietCard()) { ErrorBlock(stringResource(R.string.sims_serverFailed), actions.retry) }
            content.serverRows.isEmpty() -> EmptyCard(stringResource(R.string.sims_serverEmpty))
            else -> {
                val (rows, hasMore) = ServerRuns.visible(content.serverRows)
                Column(Modifier.fillMaxWidth().quietCard()) {
                    rows.forEachIndexed { index, row ->
                        if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
                        val id = row.simId
                        ServerRunRowView(
                            SimRowFace(row, specs, now, copy), copy, now, build,
                            progress = id?.let { state.progress[it] }, rerun = id?.let { state.rerun[it] } ?: RerunStatus.Idle,
                            offline = content.offlineSince != null, mayLinkToPlans = state.mayLinkToPlans, actions = actions,
                        )
                    }
                    if (hasMore) {
                        Text(stringResource(R.string.sims_overflow), style = FsType.meta, color = FsColor.muted, modifier = Modifier.padding(16.dp))
                    }
                }
            }
        }
    }
}

@Composable
private fun MineSection(content: SimsContent, state: SimsListState, specs: SpecCatalog, build: String?, now: Instant, copy: SimsCopy, actions: SimsActions) {
    Column(verticalArrangement = Arrangement.spacedBy(8.dp)) {
        MicroLabel(stringResource(R.string.sims_mine).uppercase(Locale.ROOT))
        when {
            content.mineFailed -> Box(Modifier.fillMaxWidth().quietCard()) { ErrorBlock(stringResource(R.string.sims_failedLoad), actions.retry) }
            content.mine.isEmpty() -> EmptyCard(stringResource(R.string.sims_empty))
            else -> Column(Modifier.fillMaxWidth().quietCard()) {
                content.mine.forEachIndexed { index, row ->
                    if (index > 0) HorizontalDivider(color = FsColor.lineSoft)
                    MySimRowView(SimRowFace(row, specs, now, copy), copy, build, actions.open)
                }
                MineFooter(content, state, actions)
            }
        }
    }
}

@Composable
private fun MineFooter(content: SimsContent, state: SimsListState, actions: SimsActions) {
    when {
        state.pageFailed -> Row(Modifier.fillMaxWidth().padding(16.dp), verticalAlignment = Alignment.CenterVertically) {
            Text(stringResource(R.string.sims_failedLoad), style = FsType.meta, color = FsColor.wipe, modifier = Modifier.weight(1f))
            BrandButton(stringResource(R.string.common_tryAgain), BrandButtonKind.Secondary, actions.loadMore, height = 48.dp, fill = false)
        }
        content.hasMorePages && content.offlineSince == null -> {
            // The spinner row loads the next page as it scrolls into view.
            LaunchedEffect(content.page) { actions.loadMore() }
            Box(Modifier.fillMaxWidth().heightIn(min = 56.dp), contentAlignment = Alignment.Center) {
                CircularProgressIndicator(color = FsColor.gold, modifier = Modifier.size(24.dp), strokeWidth = 2.dp)
            }
        }
    }
}

@Composable
private fun EmptyCard(text: String) {
    Text(text, style = FsType.meta, color = FsColor.muted, modifier = Modifier.fillMaxWidth().quietCard().padding(16.dp))
}

@Composable
fun ErrorBlock(text: String, retry: () -> Unit) {
    Column(Modifier.fillMaxWidth().padding(16.dp), verticalArrangement = Arrangement.spacedBy(12.dp)) {
        Text(text, style = FsType.lead, color = FsColor.wipe, modifier = Modifier.semantics { liveRegion = LiveRegionMode.Polite })
        BrandButton(stringResource(R.string.common_tryAgain), BrandButtonKind.Secondary, retry, height = 48.dp, fill = false, leadingIcon = Icons.Filled.Refresh)
    }
}

/** F5.6: three rows of each section at their real heights. */
@Composable
private fun SimsSkeleton() {
    Column(Modifier.padding(horizontal = 16.dp).padding(top = 12.dp), verticalArrangement = Arrangement.spacedBy(22.dp)) {
        SkeletonSection(SimsLayout.serverRowHeight, SimsLayout.serverIcon)
        SkeletonSection(SimsLayout.mineRowHeight, SimsLayout.mineIcon)
    }
}

@Composable
private fun SkeletonSection(height: Dp, icon: Dp) {
    Column(Modifier.fillMaxWidth().quietCard()) {
        repeat(3) {
            Row(Modifier.fillMaxWidth().height(height).padding(horizontal = 16.dp, vertical = 14.dp), horizontalArrangement = Arrangement.spacedBy(12.dp)) {
                Bar(icon, icon)
                Column(Modifier.weight(1f), verticalArrangement = Arrangement.spacedBy(8.dp)) { Bar(180.dp, 15.dp); Bar(120.dp, 12.dp) }
                Bar(64.dp, 14.dp)
            }
        }
    }
}

@Composable
fun Bar(width: Dp, height: Dp) {
    Box(Modifier.width(width).height(height).background(FsColor.cardTop, RoundedCornerShape(4.dp)))
}
