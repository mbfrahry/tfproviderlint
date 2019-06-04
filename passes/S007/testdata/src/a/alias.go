package a

import (
	s "github.com/hashicorp/terraform/helper/schema"
)

func falias() {
	_ = s.Schema{ // want "schema should not include `Required: false,`"
		Required: false,
	}
}
