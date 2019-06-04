package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema should not include `Required: false,`"
		Required: false,
	}

	_ = schema.Schema{
		Required: true,
	}

}
