package bytefmt

// Binary units (IEC 60027).
const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
)

// Decimal units (SI international system of units).
const (
	KB = 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
)

// Encode formats bytes integer to human-readable string according to IEC 60027.
//
// For example, 31323 bytes will return "30.59 KB".
//
// The smallest supported value is "1 B", so precision will be ignored
// below "1 KiB" or "1 KB".
func Encode(val int64) string {
	return NewEncoder().EncodeBinary(val)
}

// EncodeBinary formats bytes integer to human-readable string according to IEC 60027.
//
// For example, 31323 bytes will return "30.59 KB".
//
// The smallest supported value is "1 B", so precision will be ignored
// below "1 KB".
func EncodeBinary(val int64) string {
	return NewEncoder().EncodeBinary(val)
}

// EncodeDecimal formats bytes integer to human-readable string according to SI international system of units.
//
// For example, 31323 bytes will return "31.32 KB".
//
// The smallest supported value is "1 B", so precision will be ignored
// below "1 KiB".
func EncodeDecimal(val int64) string {
	return NewEncoder().EncodeDecimal(val)
}

// Decode parses human-readable bytes string to bytes integer.
//
// For example, "6 GiB" ("6 Gi" is also valid) will return 6442450944, and
// "6 GB" ("6 G" is also valid) will return 6000000000.
// The input case-insensitive, and the space is optional,
// so "6GiB" or "6GB" would produce the same output.
func Decode(text string) (int64, error) {
	return NewDecoder().Decode(text)
}

// DecodeBinary parses human-readable bytes string to bytes integer.
//
// For example, "6 GiB" ("6 Gi" is also valid) will return 6442450944.
// The input case-insensitive, and the space is optional,
// so "6GiB" would produce the same output.
func DecodeBinary(text string) (int64, error) {
	return NewDecoder().DecodeBinary(text)
}

// DecodeDecimal parses human-readable bytes string to bytes integer.
//
// For example, "6 GB" ("6 G" is also valid) will return 6000000000.
// The input case-insensitive, and the space is optional,
// so "6GB" would produce the same output.
func DecodeDecimal(text string) (int64, error) {
	return NewDecoder().DecodeDecimal(text)
}
