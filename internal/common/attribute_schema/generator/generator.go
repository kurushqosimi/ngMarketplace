package generator

import (
	"fmt"
	"ngMarketplace/internal/common/attribute_schema/parser"
	"ngMarketplace/internal/common/attribute_schema/translit"
	"os"
	"strings"
)

func Generate(path string, packageName string, information *parser.SchemaInformation) error {
	// packName - название пакета
	var packName = "package %s\n\n"
	// baseType - основной шаблон структуры
	var baseType = "type %s struct{\n%s}\n"

	fields, err := generateFields(information.Fields)
	if err != nil {
		return fmt.Errorf("failed to generate fields: %w", err)
	}

	var oneOfStructs []string
	if len(information.OneOf) > 0 {
		for i, oneOf := range information.OneOf {
			oneOfFields, err := generateFields(oneOf)
			if err != nil {
				return fmt.Errorf("failed to generate fields for oneOf[%d]: %w", i, err)
			}
			oneOfStructName := fmt.Sprintf("%sVariant%d", translit.TranslitFieldName(information.Title), i+1)
			oneOfStructs = append(oneOfStructs, fmt.Sprintf(baseType, oneOfStructName, oneOfFields))
		}
	}

	finalType := fmt.Sprintf(baseType, translit.TranslitFieldName(information.Title), fields)

	if len(oneOfStructs) > 0 {
		finalType += "\n" + strings.Join(oneOfStructs, "\n")
	}

	packName = fmt.Sprintf(packName, packageName)

	fileContent := packName + finalType

	filePath := path + "/" + translit.TranslitFieldName(information.Title) + ".go"

	err = os.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", path, err)
	}

	return nil
}

func generateFields(fields parser.Fields) (string, error) {
	var result strings.Builder

	for _, val := range fields.Properties {
		goFieldName := translit.TranslitFieldName(val.FieldName)
		goFieldName = strings.ToUpper(string(goFieldName[0])) + goFieldName[1:]

		var goType string
		switch val.FieldType {
		case "string":
			goType = "string"
			break
		case "int", "integer":
			goType += "int"
			break
		case "double", "float", "number":
			goType += "float64"
			break
		default:
			return "", fmt.Errorf("unsuppoted field type: %s for %s field", val.FieldType, val.FieldName)
		}

		tags := fmt.Sprintf(`json:"%s"`, val.FieldName)

		var bindingTags []string
		if inString(fields.RequiredFields, val.FieldName) {
			bindingTags = append(bindingTags, "required")
		}
		if val.MinLength != nil {
			bindingTags = append(bindingTags, fmt.Sprintf("min=%d", *val.MinLength))
		}
		if val.MaxLength != 0 {
			bindingTags = append(bindingTags, fmt.Sprintf("max=%d", val.MaxLength))
		}
		if val.Minimum != nil {
			bindingTags = append(bindingTags, fmt.Sprintf("min=%.0f", *val.Minimum))
		}
		if val.Maximum != 0 {
			bindingTags = append(bindingTags, fmt.Sprintf("max=%.0f", val.Maximum))
		}
		if len(val.Enum) > 0 {
			bindingTags = append(bindingTags, fmt.Sprintf("oneof=%s", strings.Join(val.Enum, " ")))
		}

		if len(bindingTags) > 0 {
			tags += fmt.Sprintf(` binding:"%s"`, strings.Join(bindingTags, ","))
		}

		field := fmt.Sprintf("\t%s %s `%s`", goFieldName, goType, tags)
		result.WriteString(field + "\n")
	}

	return result.String(), nil
}

func inString(arr []string, el string) bool {
	for _, arrEl := range arr {
		if arrEl == el {
			return true
		}
	}
	return false
}
