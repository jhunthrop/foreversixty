package gg.foreversixty.app.planner

import gg.foreversixty.app.planner.logic.SharedCompare
import android.content.Context
import gg.foreversixty.app.PrefsValue
import gg.foreversixty.app.bis.createBisDataSource
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.AppConfigRepository
import gg.foreversixty.app.planner.logic.EditorRequest
import gg.foreversixty.app.planner.logic.LivePlannerSource
import gg.foreversixty.app.planner.logic.PlannerDrafts
import gg.foreversixty.app.planner.logic.PlannerNavigator
import gg.foreversixty.app.planner.logic.PlannerScreen
import gg.foreversixty.app.planner.logic.PlannerSource
import gg.foreversixty.app.talents.LiveTalentsSource
import gg.foreversixty.app.network.StaticDataStore

/**
 * Lane F13's dependencies, built once in the app graph behind the `planner` switch ([create]); tests
 * pass a fixture source. Every entry point (Talents, a guide, rankings, a link) goes through
 * [navigator], and the scaffold draws its top screen.
 */
class PlannerFeature(
    val source: PlannerSource,
    val drafts: PlannerDrafts,
    /** Whether a device token is held: Your builds and its row show only then. */
    val signedIn: () -> Boolean,
    /** The session's current character, for `Plan a new build`. */
    val character: () -> AccountCharacter?,
    /** Lane F15: the account's characters, for `Compare with mine` on a shared build of another class. */
    val characters: () -> List<AccountCharacter> = { listOfNotNull(character()) },
) {
    /** The export `Compare with mine` holds against a shared build of [classSlug]; null hides the button. */
    fun compareExport(classSlug: String): AccountCharacter? = SharedCompare.exportFor(classSlug, character(), characters())

    val navigator = PlannerNavigator()

    /** A planner link (spec §3.4); false when it is not one, or needs a sign-in it does not have. */
    fun open(link: DeepLink): Boolean = navigator.open(link, signedIn())

    /** `Plan a build` for a class: its draft, else empty at the character's race and level. */
    fun plan(classSlug: String, raceSlug: String? = null, level: Int? = null) =
        navigator.push(PlannerScreen.Editor(EditorRequest.Draft(classSlug, raceSlug, level)))

    /** A guide's `Open this build in the planner`: its FS1 code as one legal order. */
    fun openCode(classSlug: String, code: String) = navigator.push(PlannerScreen.Editor(EditorRequest.Code(classSlug, code)))

    fun openBuilds() = navigator.push(PlannerScreen.Builds)

    companion object {
        fun create(
            context: Context,
            statics: StaticDataStore,
            appConfig: AppConfigRepository,
            api: ApiClient,
            signedIn: () -> Boolean,
            character: () -> AccountCharacter?,
            characters: () -> List<AccountCharacter>,
        ) = PlannerFeature(
            source = LivePlannerSource(
                files = statics,
                talentsSource = LiveTalentsSource(statics, createBisDataSource(context, statics, appConfig)),
                api = api,
            ),
            drafts = PlannerDrafts { slug -> PrefsValue(context, PlannerDrafts.key(slug)) },
            signedIn = signedIn,
            character = character,
            characters = characters,
        )
    }
}
