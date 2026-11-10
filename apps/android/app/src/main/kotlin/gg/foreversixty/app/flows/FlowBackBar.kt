package gg.foreversixty.app.flows

import androidx.compose.ui.unit.dp
import gg.foreversixty.app.designsystem.surfaceBand
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.Row
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
 * A pushed screen's bar on the painted band (round 5, as `FsTopBar`): Back to [backTo] at the left, Share at the
 * right when [share] is set. The page's head names the screen, so the bar carries no title.
 */
@Composable
fun FlowBackBar(backTo: String, back: () -> Unit, share: (() -> Unit)? = null) {
    Row(Modifier.fillMaxWidth().surfaceBand().heightIn(min = BAR_HEIGHT).padding(horizontal = FsSpace.space2), verticalAlignment = Alignment.CenterVertically) {
        IconButton(onClick = back) {
            Icon(Icons.AutoMirrored.Filled.ArrowBack, contentDescription = stringResource(R.string.flows_backTo, backTo), tint = FsColor.strong)
        }
        Spacer(Modifier.weight(1f))
        if (share != null) {
            IconButton(onClick = share) { Icon(Icons.Filled.Share, contentDescription = stringResource(R.string.share_action), tint = FsColor.strong) }
        }
    }
}

/** Material 3's small top app bar height, as `FsTopBar` draws it. */
private val BAR_HEIGHT = 64.dp
