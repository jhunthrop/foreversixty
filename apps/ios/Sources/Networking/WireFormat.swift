import Foundation

/// The API's JSON conventions in one place: snake_case keys and RFC 3339
/// timestamps, with or without fractional seconds (Go's time.Time).
public enum WireFormat {
    public static func decoder() -> JSONDecoder {
        let decoder = JSONDecoder()
        decoder.keyDecodingStrategy = .convertFromSnakeCase
        decoder.dateDecodingStrategy = .custom(decodeDate)
        return decoder
    }

    public static func encoder() -> JSONEncoder {
        let encoder = JSONEncoder()
        encoder.keyEncodingStrategy = .convertToSnakeCase
        encoder.dateEncodingStrategy = .iso8601
        encoder.outputFormatting = [.sortedKeys]
        return encoder
    }

    static func decodeDate(_ decoder: any Decoder) throws -> Date {
        let container = try decoder.singleValueContainer()
        let text = try container.decode(String.self)
        if let date = parseDate(text) {
            return date
        }
        throw DecodingError.dataCorruptedError(
            in: container,
            debugDescription: "not an RFC 3339 timestamp: \(text)"
        )
    }

    static func parseDate(_ text: String) -> Date? {
        let fractional = Date.ISO8601FormatStyle(includingFractionalSeconds: true)
        let whole = Date.ISO8601FormatStyle()
        return (try? fractional.parse(text)) ?? (try? whole.parse(text))
    }
}
