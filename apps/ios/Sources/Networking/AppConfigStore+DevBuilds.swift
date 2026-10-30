import Foundation

public extension AppConfigStore {
    /// A feature is on in debug and fixture builds whatever the config says, so every lane's
    /// screens can be built and captured before the API turns them on; a release build on the
    /// live API follows `isEnabled(_:)` alone (the brief: "defaults to on in debug and in
    /// fixture builds").
    func isAvailable(_ feature: Feature, mode: APIMode, isDebugBuild: Bool) -> Bool {
        isDebugBuild || mode == .fixture || isEnabled(feature)
    }
}
