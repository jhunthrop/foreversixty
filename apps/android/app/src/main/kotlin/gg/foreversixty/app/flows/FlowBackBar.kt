package gg.foreversixty.app.flows

import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.RowScope
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.ArrowBack
import androidx.compose.material.icons.filled.Share
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsSpace

/**
 * A pushed screen's bar (the mock's `.app-nav` with back): Back to [backTo] at the left, Share at the
 * right when [share] is set, after the screen's own [actions] (the planner's Undo). The page's head
 * names the screen, so the bar carries no title.
 */
@Composable
fun FlowBackBar(backTo: String, back: () -> Unit, share: (() -> Unit)? = null, actions: @Composable RowScope.() -> Unit = {}) {
    Row(Modifier.fillMaxWidth().padding(horizontal = FsSpace.space2), verticalAlignment = Alignment.CenterVertically) {
        IconButton(onClick = back) {
            Icon(Icons.AutoMirrored.Filled.ArrowBack, contentDescription = stringResource(R.string.flows_backTo, backTo), tint = FsColor.strong)
        }
        Spacer(Modifier.weight(1f))
        actions()
        if (share != null) {
            IconButton(onClick = share) { Icon(Icons.Filled.Share, contentDescription = stringResource(R.string.share_action), tint = FsColor.strong) }
        }
    }
}
