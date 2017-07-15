package deptestdos

import (
  "github.com/brendan-munro/deptest"
  "github.com/brendan-munro/deptest/floats"
  "github.com/brendan-munro/deptest/strings"
)

type other deptest.Foo
type another floats.CoolFloat
type yetAnother strings.CoolString
type Bar int
