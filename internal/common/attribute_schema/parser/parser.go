package parser

import (
	"encoding/json"
	"errors"
	"fmt"
)

type SchemaInformation struct {
	Title       string
	Description string
	Fields
	OneOf []Fields
}

type Fields struct {
	RequiredFields []string
	Properties     []FieldInfo
}

type FieldInfo struct {
	FieldName   string
	FieldType   string
	Enum        []string
	Default     interface{}
	MinLength   *int
	MaxLength   int
	Description string
	Minimum     *float64
	Maximum     float64
}

func ExtractInformation(data []byte) (*SchemaInformation, error) {
	if len(data) == 0 {
		return nil, errors.New("schema is empty")
	}

	var schema map[string]interface{}
	if err := json.Unmarshal(data, &schema); err != nil {
		return nil, errors.New("invalid JSON: " + err.Error())
	}

	typeVal, ok := schema["type"]
	if !ok {
		return nil, errors.New("type is not defined")
	}
	if typeStr, ok := typeVal.(string); !ok || typeStr != "object" {
		return nil, errors.New("type must be 'object'")
	}

	titleVal, ok := schema["title"]
	if !ok {
		return nil, errors.New("title should be provided")
	}

	schemaInfo := &SchemaInformation{}
	if titleValStr, ok := titleVal.(string); !ok {
		return nil, errors.New("title should be string")
	} else {
		schemaInfo.Title = titleValStr
	}

	descriptionVal, ok := schema["description"]
	if ok {
		descriptionStr, ok := descriptionVal.(string)
		if !ok {
			return nil, errors.New("description should be string")
		}
		schemaInfo.Description = descriptionStr
	}

	oneOfVal, ok := schema["oneOf"]
	if ok {
		oneOf, ok := oneOfVal.([]interface{})
		if !ok {
			return nil, errors.New("oneOf should be an array")
		}

		var fields []Fields
		for _, el := range oneOf {
			mapEl, ok := el.(map[string]interface{})
			if !ok {
				return nil, errors.New("elements of oneOf should be a map")
			}

			fieldInfo, err := extractProperties(mapEl)
			if err != nil {
				return nil, fmt.Errorf("oneOf err: %w", err)
			}
			fields = append(fields, *fieldInfo)
		}
		schemaInfo.OneOf = fields
	} else {
		fields, err := extractProperties(schema)
		if err != nil {
			return nil, err
		}
		schemaInfo.Fields = *fields
	}

	if err := checkRequiredFields(schemaInfo); err != nil {
		return nil, err
	}

	return schemaInfo, nil
}

func extractRequiredFields(requiredVal interface{}) ([]string, error) {
	required, ok := requiredVal.([]interface{})
	if !ok {
		return nil, errors.New("required must be an array")
	}

	requiredStrings := make([]string, 0, len(required))
	for _, val := range required {
		str, ok := val.(string)
		if !ok {
			return nil, errors.New("required element must be a string")
		}
		if str == "" {
			return nil, errors.New("required element cannot be empty")
		}
		requiredStrings = append(requiredStrings, str)
	}

	return requiredStrings, nil
}

func extractEnum(enumVal interface{}) ([]string, error) {
	enum, ok := enumVal.([]interface{})
	if !ok {
		return nil, errors.New("enum must be an array")
	}

	enumStrings := make([]string, 0, len(enum))
	for _, val := range enum {
		str, ok := val.(string)
		if !ok {
			return nil, errors.New("enum element must be a string")
		}
		if str == "" {
			return nil, errors.New("enum element cannot be empty")
		}
		enumStrings = append(enumStrings, str)
	}

	return enumStrings, nil
}

func extractProperties(schema map[string]interface{}) (*Fields, error) {
	var fieldsInfo Fields

	if requiredVal, ok := schema["required"]; ok {
		requiredFields, err := extractRequiredFields(requiredVal)
		if err != nil {
			return nil, err
		}
		fieldsInfo.RequiredFields = requiredFields
	}

	props := make([]FieldInfo, 0, len(fieldsInfo.RequiredFields))

	if propertiesVal, ok := schema["properties"]; ok {
		properties, ok := propertiesVal.(map[string]interface{})
		if !ok {
			return nil, errors.New("properties must be an object")
		}
		for key, prop := range properties {
			var fieldInfo FieldInfo

			if key == "" {
				return nil, errors.New("property name cannot be empty")
			}

			fieldInfo.FieldName = key

			propMap, ok := prop.(map[string]interface{})
			if !ok {
				return nil, errors.New("property must be an object")
			}

			if typeVal, ok := propMap["type"]; ok {
				typeValStr, ok := typeVal.(string)
				if !ok {
					return nil, errors.New("property type must be a string")
				}
				fieldInfo.FieldType = typeValStr
			}

			if defaultVal, ok := propMap["default"]; ok {
				switch fieldInfo.FieldType {
				case "integer", "number", "int":
					if _, ok := defaultVal.(float64); !ok {
						return nil, errors.New("default value for integer/number must be a number")
					}
				case "string":
					if _, ok := defaultVal.(string); !ok {
						return nil, errors.New("default value for string must be a string")
					}
				}
				fieldInfo.Default = defaultVal
			}

			if minVal, ok := propMap["minLength"]; ok {
				minValFloat, ok := minVal.(float64)
				if !ok {
					return nil, errors.New("property minLength must be an integer")
				}
				minValInt := int(minValFloat)
				fieldInfo.MinLength = &minValInt
			}

			if maxVal, ok := propMap["maxLength"]; ok {
				maxValFloat, ok := maxVal.(float64)
				if !ok {
					return nil, errors.New("property maxLength must be an integer")
				}
				fieldInfo.MaxLength = int(maxValFloat)
			}

			if minVal, ok := propMap["minimum"]; ok {
				minValFloat, ok := minVal.(float64)
				if !ok {
					return nil, errors.New("property maxLength must be an integer")
				}
				fieldInfo.Minimum = &minValFloat
			}

			if maxVal, ok := propMap["maximum"]; ok {
				maxValFloat, ok := maxVal.(float64)
				if !ok {
					return nil, errors.New("property maxLength must be an integer")
				}
				fieldInfo.Maximum = maxValFloat
			}

			if descriptionVal, ok := propMap["description"]; ok {
				descriptionValStr, ok := descriptionVal.(string)
				if !ok {
					return nil, errors.New("property description must be a string")
				}
				fieldInfo.Description = descriptionValStr
			}

			if enumVal, ok := propMap["enum"]; ok {
				enum, err := extractEnum(enumVal)
				if err != nil {
					return nil, err
				}
				fieldInfo.Enum = enum
			}

			props = append(props, fieldInfo)
		}
	}

	fieldsInfo.Properties = props

	return &fieldsInfo, nil
}

func checkRequiredFields(schemaInfo *SchemaInformation) error {
	for _, reqFiledName := range schemaInfo.RequiredFields {
		found := false
		for _, prop := range schemaInfo.Properties {
			if prop.FieldName == reqFiledName {
				found = true
			}
		}
		if !found {
			return errors.New("not all required fields are provided")
		}
	}

	if len(schemaInfo.Properties) == 0 && len(schemaInfo.RequiredFields) > 0 {
		return errors.New("schema must define at least required properties")
	}

	return nil
}
