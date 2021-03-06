package a

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func f() {
	_ = schema.Schema{ // want "schema of TypeList or TypeSet should include Elem"
		Type: schema.TypeList,
	}

	_ = schema.Schema{ // want "schema of TypeList or TypeSet should include Elem"
		Type: schema.TypeSet,
	}

	_ = schema.Schema{
		Type: schema.TypeList,
		Elem: &schema.Schema{Type: schema.TypeString},
	}

	_ = schema.Schema{
		Type: schema.TypeSet,
		Elem: &schema.Schema{Type: schema.TypeString},
	}
}
