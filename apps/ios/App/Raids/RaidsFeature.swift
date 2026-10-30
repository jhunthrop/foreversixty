import Networking

extension AppEnvironment {
    /// The Raids tab ships dark behind the `raids` switch; debug and fixture builds turn it on.
    var raidsEnabled: Bool {
        #if DEBUG
        return true
        #else
        return configuration.mode == .fixture || appConfig.isEnabled(.raids)
        #endif
    }
}
