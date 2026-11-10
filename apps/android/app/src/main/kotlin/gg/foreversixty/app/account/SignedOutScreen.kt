package gg.foreversixty.app.account

import gg.foreversixty.app.designsystem.panel
import gg.foreversixty.app.designsystem.screenGround
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.LinearProgressIndicator
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.heading
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import gg.foreversixty.app.R
import gg.foreversixty.app.common.Seal
import gg.foreversixty.app.designsystem.BrandButton
import gg.foreversixty.app.designsystem.BrandButtonKind
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.surfaceBand
import gg.foreversixty.app.designsystem.FsSpace
import gg.foreversixty.app.designsystem.ScaledType

/** Which tab draws the signed-out panel: Home's words, or Raids' (spec F1.1). */
enum class SignedOutVariant { Home, Raids }

/**
 * F1.1 signed out, with F1.2's spinner while the Custom Tab is open and F1.4's toast. Stateless:
 * the screenshot tests render it directly.
 */
@Composable
fun SignedOutScreen(
    notice: String?,
    signingIn: Boolean,
    onSignIn: () -> Unit,
    onPickClass: () -> Unit,
    variant: SignedOutVariant = SignedOutVariant.Home,
    noticeIsFailure: Boolean = true,
) {
    // The web's signed-out page: a band head, then one Plate on the ground with the one primary.
    Box(Modifier.fillMaxSize().screenGround()) {
        Column(Modifier.fillMaxSize().verticalScroll(rememberScrollState())) {
            val (heading, line) = when (variant) {
                SignedOutVariant.Home -> R.string.signin_line to R.string.signin_body
                SignedOutVariant.Raids -> R.string.raids_signedOut_title to R.string.raids_signedOut_body
            }
            Column(
                Modifier
                    .fillMaxWidth()
                    .surfaceBand()
                    .padding(horizontal = FsSpace.space4)
                    .padding(top = if (notice == null) 48.dp else 108.dp, bottom = FsSpace.space6),
                verticalArrangement = Arrangement.spacedBy(FsSpace.space4),
            ) {
                Seal(64.dp)
                Text(stringResource(R.string.app_name), style = ScaledType.wordmark, color = FsColor.strong)
                Text(
                    stringResource(heading),
                    style = ScaledType.title,
                    color = FsColor.strong,
                    modifier = Modifier.semantics { this.heading() },
                )
            }
            Column(
                Modifier.padding(FsSpace.space4).fillMaxWidth().panel().padding(FsSpace.space4),
                verticalArrangement = Arrangement.spacedBy(FsSpace.space4),
            ) {
                Text(stringResource(line), style = ScaledType.body, color = FsColor.text)
                BrandButton(stringResource(R.string.signin_button), BrandButtonKind.Primary, onSignIn, busy = signingIn)
                BrandButton(stringResource(R.string.signin_pickClass), BrandButtonKind.Quiet, onPickClass)
                Text(stringResource(R.string.signin_pickClassNote), style = ScaledType.meta, color = FsColor.muted)
            }
        }
        if (notice != null) {
            NoticeToast(notice, Modifier.padding(horizontal = FsSpace.space4, vertical = FsSpace.space3), isFailure = noticeIsFailure)
        }
    }
}

/** F1.3: after the redirect, until `GET /v1/me` answers. */
@Composable
fun ImportScreen() {
    Column(
        Modifier.fillMaxSize().surfaceBand(bottomLine = false).padding(horizontal = FsSpace.space4),
        verticalArrangement = Arrangement.spacedBy(FsSpace.space4, Alignment.CenterVertically),
        horizontalAlignment = Alignment.CenterHorizontally,
    ) {
        Seal(64.dp)
        Text(stringResource(R.string.signin_importTitle), style = ScaledType.h2, color = FsColor.strong, textAlign = TextAlign.Center)
        Text(stringResource(R.string.signin_importBody), style = ScaledType.body, color = FsColor.text, textAlign = TextAlign.Center)
        LinearProgressIndicator(Modifier.width(160.dp).height(8.dp), color = FsColor.gold, trackColor = FsColor.lineSoft)
    }
}

/** F1.3 with no characters: Home's no-characters state. */
@Composable
fun NoCharactersScreen(onGetAddon: () -> Unit) {
    Column(
        Modifier.fillMaxSize().screenGround().padding(FsSpace.space4),
        verticalArrangement = Arrangement.spacedBy(FsSpace.space4),
    ) {
        Text(
            stringResource(R.string.signin_noChars_title),
            style = ScaledType.h2,
            color = FsColor.strong,
            modifier = Modifier.semantics { heading() },
        )
        Text(stringResource(R.string.signin_noChars_body), style = ScaledType.body, color = FsColor.text)
        BrandButton(stringResource(R.string.shared_cta_getAddonAndCompanion), BrandButtonKind.Secondary, onGetAddon, Modifier.fillMaxWidth())
    }
}
