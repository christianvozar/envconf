// Copyright © 2015-2016, Rogue Ethic, LLC.
// Licensed under MIT. All rights reserved.
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

// Version is the semantic version number being executed.
const Version = "0.1.0"

// VersionPrerelease identifies a pre-release marker for the version. If ""
// (empty string) then this is a final release. Otherwise, this a pre-release
// such as "dev" (in development), "beta", "rc1", etc.
