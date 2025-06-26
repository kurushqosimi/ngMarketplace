package category

import (
	"encoding/json"
	"errors"
	"fmt"
	"ngMarketplace/internal/category/attribute_schema/parser"
	"ngMarketplace/pkg/validator"
	"time"
)

// Category represents a category in the marketplace
type Category struct {
	CategoryID      int             `json:"category_id"`
	CategoryName    string          `json:"category_name"`
	ParentID        *int            `json:"parent_id"`
	Language        string          `json:"language"`
	AttributeSchema json.RawMessage `json:"attribute_schema"`
	CreatedAt       time.Time       `json:"-"`
	Active          bool            `json:"-"`
	UpdatedAt       time.Time       `json:"-"`
	DeletedAt       time.Time       `json:"-"`
}

func validateCategory(v *validator.Validator, category *Category) {
	v.Check(len(category.CategoryName) <= 50, "category_name", "must not be more than 50 bytes long")

	if len(category.AttributeSchema) > 0 {
		parsedInfo, err := parser.ExtractInformation(category.AttributeSchema)
		if err != nil {
			v.AddError("failed to extract information from attribute_schema", err.Error())
			return
		}
		validateAttributeSchema(v, parsedInfo)
	}
}

func validateAttributeSchema(v *validator.Validator, info *parser.SchemaInformation) {
	if len(info.OneOf) > 0 {
		for i, oneOf := range info.OneOf {
			if err := validateProperties(&oneOf); err != nil {
				v.AddError("attribute_schema", fmt.Sprintf("failed to validate fields for oneOf[%d]: %v", i, err))
			}
		}
	} else {
		if err := validateProperties(&info.Fields); err != nil {
			v.AddError("attribute_schema", fmt.Sprintf("failed to validate fields: %v", err))
		}
	}
}

func validateProperties(fields *parser.Fields) error {
	for _, val := range fields.Properties {
		if val.FieldType != "string" && val.FieldType != "int" && val.FieldType != "double" && val.FieldType != "float" && val.FieldType != "number" && val.FieldType != "integer" {
			return fmt.Errorf("unsuppoted field type: %s for %s field", val.FieldType, val.FieldName)
		}
	}

	return nil
}

var (
	ErrUpdate   = errors.New("no active category to update")
	ErrDelete   = errors.New("no active category to delete")
	ErrNotFound = errors.New("category not found")
)

// Repository Errors
var (
	ErrDuplicateCategory = errors.New("category already exists")
	ErrInvalidParentID   = errors.New("parent category does not exist")
	ErrConnectionFailed  = errors.New("database connection failed")
)

// Service Errors
var (
	ErrInvalid = errors.New("validation failed")
)

// Handler Errors
var (
	ErrBindJSON = errors.New("failed binding json")
)
