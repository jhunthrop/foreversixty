package gg.foreversixty.app.network

/**
 * A feature is on in debug and fixture builds whatever the config says, so every lane's screens
 * can be built and captured before the API turns them on; a release build on the live API
 * follows [AppConfigRepository.isEnabled] alone (the Wave 1 brief).
 */
fun AppConfigRepository.isAvailable(feature: Feature, mode: ApiMode, isDebugBuild: Boolean): Boolean =
    isDebugBuild || mode == ApiMode.Fixture || isEnabled(feature)
