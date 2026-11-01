package gg.foreversixty.app.home

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.FsType
import gg.foreversixty.app.designsystem.quietCard
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.paperdoll.Paperdoll
import gg.foreversixty.app.paperdoll.PaperdollFrame
import java.util.Locale

/**
 * F2.7 Your gear: all 17 slots on the Paperdoll (lane F11, tenet 10), from GearModel's rows, with
 * the gold dot where Home has an upgrade. [open] shows an item's tooltip; its second value is the
 * item worn in that slot. Scrolling is the caller's.
 */
@Composable
fun GearList(rows: List<GearRow>, intro: String, build: String, open: (ItemLook, Int?) -> Unit = { _, _ -> }) {
    val slots = Paperdoll.fromGear(rows, LocalResources.current)
    Column(Modifier.fillMaxWidth().padding(16.dp), verticalArrangement = Arrangement.spacedBy(16.dp)) {
        Text(
            stringResource(R.string.gear_title).uppercase(Locale.ROOT),
            style = FsType.title,
            color = FsColor.strong,
            modifier = Modifier.semantics { heading() },
        )
        Text(intro, style = FsType.body, color = FsColor.muted)
        Column(Modifier.fillMaxWidth().quietCard().padding(12.dp), verticalArrangement = Arrangement.spacedBy(10.dp)) {
            PaperdollFrame(slots, build, open = { slot -> slot.item?.let { open(it, it.id) } })
            if (slots.any { it.better }) Text(stringResource(R.string.char_goldDot), style = FsType.meta, color = FsColor.muted)
        }
    }
}
