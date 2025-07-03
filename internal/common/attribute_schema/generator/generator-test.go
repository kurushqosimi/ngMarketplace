package generator

import (
	"log"
	"ngMarketplace/internal/common/attribute_schema/parser"
)

func main() {
	path := "/home/kurush/Desktop/Projects/Production/ngMarketplace/test"
	packageName := "category"

	example1 := `{
	"type": "object",
	"title": "Test",
	"description": "Test description",
	"required": ["f1", "f2"],
	"properties": {
		"f1": {
			"type": "string",
			"default": "f1",
			"enum": ["enum1", "enum2"]
		},
		"f2": {
			"type": "int",
			"minLength": 1,
			"maxLength": 2,
			"description": "test"
		}
	}
}`

	test, err := parser.ExtractInformation([]byte(example1))
	if err != nil {
		log.Fatal(err)
	}

	err = Generate(path, packageName, test)
	if err != nil {
		log.Fatal(err)
	}

	example2 := `{
	"type": "object",
	"title": "Электроника",
	"description": "Атрибуты для категории электроники",
	"required": ["гарантия", "бренд"],
	"properties": {
		"гарантия": {
			"type": "integer",
			"description": "Срок гарантии в месяцах",
			"minimum": 0,
			"maximum": 60
		},
		"бренд": {
			"type": "string",
			"enum": ["Samsung", "Apple", "Xiaomi", "Другое"]
		}
	}
}`

	test2, err := parser.ExtractInformation([]byte(example2))
	if err != nil {
		log.Fatal(err)
	}

	err = Generate(path, packageName, test2)
	if err != nil {
		log.Fatal(err)
	}

	example3 := `{
    "type": "object",
    "title": "Смартфоны",
    "description": "Атрибуты для смартфонов",
    "required": ["cpu", "ram"],
    "properties": {
      "cpu": {
        "type": "string",
        "description": "Тип процессора",
        "enum": ["Snapdragon", "Exynos", "Apple A-series", "MediaTek"]
      },
      "ram": {
        "type": "integer",
        "description": "Объем оперативной памяти в ГБ",
        "minimum": 2,
        "maximum": 16
      },
      "цвет": {
        "type": "string",
        "description": "Цвет устройства",
        "enum": ["Черный", "Белый", "Синий", "Красный"]
      }
    }
  }`

	test3, err := parser.ExtractInformation([]byte(example3))
	if err != nil {
		log.Fatal(err)
	}

	err = Generate(path, packageName, test3)
	if err != nil {
		log.Fatal(err)
	}

	example4 := `{
	 "type": "object",
	 "title": "Пылесосы",
	 "description": "Атрибуты для пылесосов с использованием oneOf",
	 "oneOf": [
	   {
	     "properties": {
	       "тип": {
	         "type": "string",
	         "const": "Робот"
	       },
	       "батарея": {
	         "type": "integer",
	         "description": "Емкость батареи в мАч",
	         "minimum": 2000
	       }
	     },
	     "required": ["тип", "батарея"]
	   },
	   {
	     "properties": {
	       "тип": {
	         "type": "string",
	         "const": "Классический"
	       },
	       "мощность": {
	         "type": "integer",
	         "description": "Мощность в ваттах",
	         "minimum": 500
	       }
	     },
	     "required": ["тип", "мощность"]
	   }
	 ]
	}`

	test4, err := parser.ExtractInformation([]byte(example4))
	if err != nil {
		log.Fatal(err)
	}

	err = Generate(path, packageName, test4)
}
