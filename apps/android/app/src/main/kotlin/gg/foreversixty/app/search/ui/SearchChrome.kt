package gg.foreversixty.app.search.ui

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Search
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.ModalBottomSheet
import androidx.compose.material3.rememberModalBottomSheetState
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.talents.TalentRankContent
import gg.foreversixty.app.talents.TalentSheetModel

/**
 * A root bar (Home, Raids, BiS): [selector] in the leading slot and, while [onSearch] is set, the
 * 48 dp search button at the trailing edge (spec §3.2). Without it the bar is the selector alone.
 */
@Composable
fun RootSearchBar(onSearch: (() -> Unit)?, selector: @Composable () -> Unit) {
    if (onSearch == null) {
        selector()
        return
    }
    Row(Modifier.fillMaxWidth().padding(end = 4.dp), verticalAlignment = Alignment.CenterVertically) {
        Box(Modifier.weight(1f)) { selector() }
        SearchButton(onSearch)
    }
}

@Composable
fun SearchButton(onClick: () -> Unit) {
    IconButton(onClick = onClick) {
        Icon(Icons.Filled.Search, contentDescription = stringResource(R.string.search_button), tint = FsColor.text)
    }
}

/** A talent from Search: the v1 F6 rank sheet's body, then `Open in planner` (the web planner until F13). */
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun TalentSearchSheet(model: TalentSheetModel, build: String, onPlanner: () -> Unit, onDismiss: () -> Unit) {
    ModalBottomSheet(
        onDismissRequest = onDismiss,
        sheetState = rememberModalBottomSheetState(skipPartiallyExpanded = true),
        containerColor = FsColor.raised,
    ) {
        TalentSearchContent(model, build, onPlanner)
    }
}

/** The sheet's body, pure for the screenshot tests. */
@Composable
fun TalentSearchContent(model: TalentSheetModel, build: String, onPlanner: () -> Unit = {}) {
    Column(verticalArrangement = Arrangement.spacedBy(4.dp)) {
        TalentRankContent(model, build)
        BrandButton(
            stringResource(R.string.search_talent_openPlanner), BrandButtonKind.Secondary, onPlanner,
            modifier = Modifier.padding(horizontal = 20.dp).padding(bottom = 16.dp), height = 48.dp,
        )
    }
}
