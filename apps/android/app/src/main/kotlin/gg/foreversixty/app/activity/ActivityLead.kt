package gg.foreversixty.app.activity

import androidx.compose.runtime.Composable
import androidx.compose.runtime.remember
import androidx.compose.ui.platform.LocalContext
import gg.foreversixty.app.bis.bundledSpecCatalog
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.model.ActivityRow

/**
 * An Activity row's lead, on the tab and on Home's New for you and Last night alike: the class crest
 * for a row about a class ([Activity.crestClass]), else the server's game icon (the item, the
 * instance, the pocketwatch).
 */
@Composable
fun activityLead(row: ActivityRow, build: String?): FlowLead {
    val context = LocalContext.current
    val specs = remember(context) { bundledSpecCatalog(context) }
    return Activity.crestClass(row, specs)?.let { FlowLead.Crest(it) } ?: FlowLead.Icon(row.icon, build)
}
