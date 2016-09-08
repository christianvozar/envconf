// MIT Licensed.
// Christian R. Vozar <christian@rogueethic.com>
// Fabriqué en Nouvelle Orléans ⚜

package envconf

var (
	// GitCommit is the git commit that was compiled. This will be filled in by
	// the compiler.
	GitCommit string
	// GitDescribe is the git description that was compiled. This will be filled
	// in by the compiler.
	GitDescribe string
)

const (
	// Version is the semantic version number being executed.
	Version = "1.1.0"

	// VersionPrerelease marks the pre-release version. If this is ""
	// (empty string) then it is a final release. Otherwise, this is a pre-release
	// such as "dev" (in development), "beta", "rc1", etc.
	VersionPrerelease = ""
)
