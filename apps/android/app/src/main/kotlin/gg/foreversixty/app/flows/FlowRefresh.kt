package gg.foreversixty.app.flows

import androidx.compose.foundation.ScrollState
import androidx.compose.foundation.layout.Box
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.pulltorefresh.PullToRefreshBox
import androidx.compose.material3.pulltorefresh.PullToRefreshDefaults
import androidx.compose.material3.pulltorefresh.rememberPullToRefreshState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import gg.foreversixty.app.common.Haptics
import gg.foreversixty.app.common.rememberHaptics
import gg.foreversixty.app.designsystem.FsColor
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.launch

/** Pull to refresh on a tab root (spec R12): the spinner shows while [refreshing]; [onRefresh] starts the reload. */
class FlowRefresh(val refreshing: Boolean, val onRefresh: () -> Unit)

/**
 * A tab root's page state: its [scroll], kept by the route while a page is pushed over the root and
 * popped (the default lives as long as the page does), and its pull to refresh, if it has one.
 */
class FlowPageState(val scroll: ScrollState, val refresh: FlowRefresh? = null)

/**
 * The tab root's refresh: R14's light haptic at the trigger point, then [reload] with the spinner up
 * until it returns. Home, the Raid tab and Activity use it.
 */
@Composable
fun rememberFlowRefresh(reload: suspend () -> Unit): FlowRefresh {
    val scope = rememberCoroutineScope()
    val haptics = rememberHaptics()
    var refreshing by remember { mutableStateOf(false) }
    return FlowRefresh(refreshing) { startRefresh(scope, haptics, { refreshing = it }, reload) }
}

private fun startRefresh(scope: CoroutineScope, haptics: Haptics, show: (Boolean) -> Unit, reload: suspend () -> Unit) {
    haptics.selection()
    scope.launch {
        show(true)
        try {
            reload()
        } finally {
            show(false)
        }
    }
}

/** The system control, tinted gold on the raised ground (spec 3.3), over [content]; none without [refresh]. */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
internal fun FlowRefreshBox(refresh: FlowRefresh?, modifier: Modifier = Modifier, content: @Composable () -> Unit) {
    if (refresh == null) {
        Box(modifier) { content() }
        return
    }
    val state = rememberPullToRefreshState()
    PullToRefreshBox(
        isRefreshing = refresh.refreshing,
        onRefresh = refresh.onRefresh,
        modifier = modifier,
        state = state,
        indicator = {
            PullToRefreshDefaults.Indicator(
                state = state,
                isRefreshing = refresh.refreshing,
                modifier = Modifier.align(Alignment.TopCenter),
                containerColor = FsColor.raised,
                color = FsColor.gold,
            )
        },
    ) { content() }
}
