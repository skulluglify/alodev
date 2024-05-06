module alodev

go 1.22

require (
	dvwk v1.0.0
	dvwk/singletons v1.0.0
)

replace (
	dvwk v1.0.0 => ./workspaces/dvwk
)
