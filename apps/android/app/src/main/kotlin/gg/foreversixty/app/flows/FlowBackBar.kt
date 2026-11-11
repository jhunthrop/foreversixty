package gg.foreversixty.app.flows

import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Share
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.runtime.Composable
import androidx.compose.ui.res.stringResource
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsTopBar

/**
 * A raid flow's bar: [FsTopBar] on the painted band with Back to [backTo], and Share when [share] is
 * set. The page's head names the screen (the raid, the applicant), so the bar carries no title.
 */
@Composable
fun FlowBackBar(backTo: String, back: () -> Unit, share: (() -> Unit)? = null) {
    FsTopBar(null, stringResource(R.string.flows_backTo, backTo), back) {
        if (share != null) {
            IconButton(onClick = share) { Icon(Icons.Filled.Share, contentDescription = stringResource(R.string.share_action), tint = FsColor.strong) }
        }
    }
}
