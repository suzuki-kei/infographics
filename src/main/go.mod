module main

go 1.18

require (
    src/assert v0.0.0
    src/bigints v0.0.0
    src/infographics v0.0.0
    src/logging v0.0.0
    src/slices v0.0.0
)

replace src/assert => ../assert
replace src/bigints => ../bigints
replace src/infographics => ../infographics
replace src/logging => ../logging
replace src/slices => ../slices

