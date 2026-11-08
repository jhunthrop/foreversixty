package gg.foreversixty.app.guides.logic

import android.net.Uri

/** The site pages the tier list links to that the app has no screen for (the web's tier-load.ts and tier-copy.ts). */
object GuideLinks {
    private const val SITE = "https://foreversixty.gg"

    /** `How we check the sim ›`. */
    const val SIM_CHECKS = "$SITE/sim/specs"

    /** A spec's ranking sim, run on open: `/sim?code=warrior-fury-alliance&run=1`. */
    fun rankingSim(spec: String, faction: String): String = "$SITE/sim?code=${Uri.encode("$spec-$faction")}&run=1"
}
