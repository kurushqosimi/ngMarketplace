-- Вставка категорий для русского языка ('ru')
INSERT INTO categories (category_name, parent_id, language, attribute_schema, created_at, active, updated_at, deleted_at)
VALUES
    -- Корневая категория: Электроника
    ('Электроника', NULL, 'ru', '{
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
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: Смартфоны
    ('Смартфоны', 1, 'ru', '{
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
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: Ноутбуки
    ('Ноутбуки', 1, 'ru', '{
    "type": "object",
    "title": "Ноутбуки",
    "description": "Атрибуты для ноутбуков",
    "required": ["cpu", "экран"],
    "properties": {
      "cpu": {
        "type": "string",
        "description": "Тип процессора",
        "enum": ["Intel Core", "AMD Ryzen", "Apple M-series"]
      },
      "экран": {
        "type": "number",
        "description": "Диагональ экрана в дюймах",
        "minimum": 10,
        "maximum": 18
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Корневая категория: Одежда
    ('Одежда', NULL, 'ru', '{
    "type": "object",
    "title": "Одежда",
    "description": "Атрибуты для одежды",
    "required": ["размер", "материал"],
    "properties": {
      "размер": {
        "type": "string",
        "description": "Размер одежды",
        "enum": ["XS", "S", "M", "L", "XL"]
      },
      "материал": {
        "type": "string",
        "description": "Материал изделия",
        "enum": ["Хлопок", "Полиэстер", "Шерсть"]
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: Футболки
    ('Футболки', 4, 'ru', '{
    "type": "object",
    "title": "Футболки",
    "description": "Атрибуты для футболок",
    "required": ["цвет", "размер"],
    "properties": {
      "цвет": {
        "type": "string",
        "description": "Цвет футболки",
        "enum": ["Белый", "Черный", "Серый"]
      },
      "размер": {
        "type": "string",
        "description": "Размер футболки",
        "enum": ["S", "M", "L"]
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Корневая категория: Бытовая техника
    ('Бытовая техника', NULL, 'ru', '{
    "type": "object",
    "title": "Бытовая техника",
    "description": "Атрибуты для бытовой техники",
    "required": ["гарантия"],
    "properties": {
      "гарантия": {
        "type": "integer",
        "description": "Срок гарантии в месяцах",
        "minimum": 12,
        "maximum": 36
      },
      "мощность": {
        "type": "integer",
        "description": "Мощность в ваттах",
        "minimum": 100
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: Холодильники
    ('Холодильники', 6, 'ru', '{
    "type": "object",
    "title": "Холодильники",
    "description": "Атрибуты для холодильников",
    "required": ["объем"],
    "properties": {
      "объем": {
        "type": "integer",
        "description": "Объем в литрах",
        "minimum": 100,
        "maximum": 1000
      },
      "цвет": {
        "type": "string",
        "description": "Цвет холодильника",
        "enum": ["Белый", "Серебристый", "Черный"]
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: Пылесосы
    ('Пылесосы', 6, 'ru', '{
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
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: Аксессуары для электроники
    ('Аксессуары для электроники', 1, 'ru', '{
    "type": "object",
    "title": "Аксессуары для электроники",
    "description": "Атрибуты для аксессуаров",
    "required": ["тип_аксессуара"],
    "properties": {
      "тип_аксессуара": {
        "type": "string",
        "description": "Тип аксессуара",
        "enum": ["Чехол", "Зарядное устройство", "Наушники"]
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: Джинсы
    ('Джинсы', 4, 'ru', '{
    "type": "object",
    "title": "Джинсы",
    "description": "Атрибуты для джинсов",
    "required": ["размер", "цвет"],
    "properties": {
      "размер": {
        "type": "string",
        "description": "Размер джинсов",
        "enum": ["28", "30", "32", "34"]
      },
      "цвет": {
        "type": "string",
        "description": "Цвет джинсов",
        "enum": ["Синий", "Черный", "Серый"]
      }
    }
  }', NOW(), true, NOW(), NULL);

-- Вставка категорий для английского языка ('en')
INSERT INTO categories (category_name, parent_id, language, attribute_schema, created_at, active, updated_at, deleted_at)
VALUES
    -- Корневая категория: Electronics
    ('Electronics', NULL, 'en', '{
    "type": "object",
    "title": "Electronics",
    "description": "Attributes for electronics category",
    "required": ["warranty", "brand"],
    "properties": {
      "warranty": {
        "type": "integer",
        "description": "Warranty period in months",
        "minimum": 0,
        "maximum": 60
      },
      "brand": {
        "type": "string",
        "enum": ["Samsung", "Apple", "Xiaomi", "Other"]
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: Smartphones
    ('Smartphones', 11, 'en', '{
    "type": "object",
    "title": "Smartphones",
    "description": "Attributes for smartphones",
    "required": ["cpu", "ram"],
    "properties": {
      "cpu": {
        "type": "string",
        "description": "Processor type",
        "enum": ["Snapdragon", "Exynos", "Apple A-series", "MediaTek"]
      },
      "ram": {
        "type": "integer",
        "description": "RAM size in GB",
        "minimum": 2,
        "maximum": 16
      },
      "color": {
        "type": "string",
        "description": "Device color",
        "enum": ["Black", "White", "Blue", "Red"]
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: Laptops
    ('Laptops', 11, 'en', '{
    "type": "object",
    "title": "Laptops",
    "description": "Attributes for laptops",
    "required": ["cpu", "screen"],
    "properties": {
      "cpu": {
        "type": "string",
        "description": "Processor type",
        "enum": ["Intel Core", "AMD Ryzen", "Apple M-series"]
      },
      "screen": {
        "type": "number",
        "description": "Screen size in inches",
        "minimum": 10,

        "maximum": 18
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Корневая категория: Clothing
    ('Clothing', NULL, 'en', '{
    "type": "object",
    "title": "Clothing",
    "description": "Attributes for clothing",
    "required": ["size", "material"],
    "properties": {
      "size": {
        "type": "string",
        "description": "Clothing size",
        "enum": ["XS", "S", "M", "L", "XL"]
      },
      "material": {
        "type": "string",
        "description": "Material type",
        "enum": ["Cotton", "Polyester", "Wool"]
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: T-Shirts
    ('T-Shirts', 14, 'en', '{
    "type": "object",
    "title": "T-Shirts",
    "description": "Attributes for t-shirts",
    "required": ["color", "size"],
    "properties": {
      "color": {
        "type": "string",
        "description": "T-shirt color",
        "enum": ["White", "Black", "Gray"]
      },
      "size": {
        "type": "string",
        "description": "T-shirt size",
        "enum": ["S", "M", "L"]
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Корневая категория: Home Appliances
    ('Home Appliances', NULL, 'en', '{
    "type": "object",
    "title": "Home Appliances",
    "description": "Attributes for home appliances",
    "required": ["warranty"],
    "properties": {
      "warranty": {
        "type": "integer",
        "description": "Warranty period in months",
        "minimum": 12,
        "maximum": 36
      },
      "power": {
        "type": "integer",
        "description": "Power in watts",
        "minimum": 100
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: Refrigerators
    ('Refrigerators', 16, 'en', '{
    "type": "object",
    "title": "Refrigerators",
    "description": "Attributes for refrigerators",
    "required": ["capacity"],
    "properties": {
      "capacity": {
        "type": "integer",
        "description": "Capacity in liters",
        "minimum": 100,
        "maximum": 1000
      },
      "color": {
        "type": "string",
        "description": "Refrigerator color",
        "enum": ["White", "Silver", "Black"]
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: Vacuum Cleaners
    ('Vacuum Cleaners', 16, 'en', '{
    "type": "object",
    "title": "Vacuum Cleaners",
    "description": "Attributes for vacuum cleaners with oneOf",
    "oneOf": [
      {
        "properties": {
          "type": {
            "type": "string",
            "const": "Robot"
          },
          "battery": {
            "type": "integer",
            "description": "Battery capacity in mAh",
            "minimum": 2000
          }
        },
        "required": ["type", "battery"]
      },
      {
        "properties": {
          "type": {
            "type": "string",
            "const": "Classic"
          },
          "power": {
            "type": "integer",
            "description": "Power in watts",
            "minimum": 500
          }
        },
        "required": ["type", "power"]
      }
    ]
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: Electronics Accessories
    ('Electronics Accessories', 11, 'en', '{
    "type": "object",
    "title": "Electronics Accessories",
    "description": "Attributes for accessories",
    "required": ["accessory_type"],
    "properties": {
      "accessory_type": {
        "type": "string",
        "description": "Type of accessory",
        "enum": ["Case", "Charger", "Headphones"]
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: Jeans
    ('Jeans', 14, 'en', '{
    "type": "object",
    "title": "Jeans",
    "description": "Attributes for jeans",
    "required": ["size", "color"],
    "properties": {
      "size": {
        "type": "string",
        "description": "Jeans size",
        "enum": ["28", "30", "32", "34"]
      },
      "color": {
        "type": "string",
        "description": "Jeans color",
        "enum": ["Blue", "Black", "Gray"]
      }
    }
  }', NOW(), true, NOW(), NULL);

-- Вставка категорий для таджикского языка ('tj')
INSERT INTO categories (category_name, parent_id, language, attribute_schema, created_at, active, updated_at, deleted_at)
VALUES
    -- Корневая категория: Электроника
    ('Электроника', NULL, 'tj', '{
    "type": "object",
    "title": "Электроника",
    "description": "Хусусиятҳо барои категорияи электроника",
    "required": ["кафолат", "бренд"],
    "properties": {
      "кафолат": {
        "type": "integer",
        "description": "Муддати кафолат дар моҳҳо",
        "minimum": 0,
        "maximum": 60
      },
      "бренд": {
        "type": "string",
        "enum": ["Samsung", "Apple", "Xiaomi", "Дигар"]
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: Смартфонҳо
    ('Смартфонҳо', 21, 'tj', '{
    "type": "object",
    "title": "Смартфонҳо",
    "description": "Хусусиятҳо барои смартфонҳо",
    "required": ["cpu", "ram"],
    "properties": {
      "cpu": {
        "type": "string",
        "description": "Навъи протсессор",
        "enum": ["Snapdragon", "Exynos", "Apple A-series", "MediaTek"]
      },
      "ram": {
        "type": "integer",
        "description": "Ҳаҷми хотираи оперативӣ дар ГБ",
        "minimum": 2,
        "maximum": 16
      },
      "ранг": {
        "type": "string",
        "description": "Ранги дастгоҳ",
        "enum": ["Сиёҳ", "Сафед", "Кабуд", "Сурх"]
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: Ноутбукҳо
    ('Ноутбукҳо', 21, 'tj', '{
    "type": "object",
    "title": "Ноутбукҳо",
    "description": "Хусусиятҳо барои ноутбукҳо",
    "required": ["cpu", "экран"],
    "properties": {
      "cpu": {
        "type": "string",
        "description": "Навъи протсессор",
        "enum": ["Intel Core", "AMD Ryzen", "Apple M-series"]
      },
      "экран": {
        "type": "number",
        "description": "Андозаи экран дар дюйм",
        "minimum": 10,
        "maximum": 18
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Корневая категория: Либос
    ('Либос', NULL, 'tj', '{
    "type": "object",
    "title": "Либос",
    "description": "Хусусиятҳо барои либос",
    "required": ["андоза", "материал"],
    "properties": {
      "андоза": {
        "type": "string",
        "description": "Андозаи либос",
        "enum": ["XS", "S", "M", "L", "XL"]
      },
      "материал": {
        "type": "string",
        "description": "Навъи материал",
        "enum": ["Пахта", "Полиэстер", "Пашм"]
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: Футболкаҳо
    ('Футболкаҳо', 24, 'tj', '{
    "type": "object",
    "title": "Футболкаҳо",
    "description": "Хусусиятҳо барои футболкаҳо",
    "required": ["ранг", "андоза"],
    "properties": {
      "ранг": {
        "type": "string",
        "description": "Ранги футболка",
        "enum": ["Сафед", "Сиёҳ", "Хокистарӣ"]
      },
      "андоза": {
        "type": "string",
        "description": "Андозаи футболка",
        "enum": ["S", "M", "L"]
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Корневая категория: Техникаи хонагӣ
    ('Техникаи хонагӣ', NULL, 'tj', '{
    "type": "object",
    "title": "Техникаи хонагī",
    "description": "Хусусиятҳо барои техникаи хонагī",
    "required": ["кафолат"],
    "properties": {
      "кафолат": {
        "type": "integer",
        "description": "Муддати кафолат дар моҳҳо",
        "minimum": 12,
        "maximum": 36
      },
      "қувват": {
        "type": "integer",
        "description": "Қувват дар ватт",
        "minimum": 100
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: Яхдонҳо
    ('Яхдонҳо', 26, 'tj', '{
    "type": "object",
    "title": "Яхдонҳо",
    "description": "Хусусиятҳо барои яхдонҳо",
    "required": ["ҳаҷм"],
    "properties": {
      "ҳаҷм": {
        "type": "integer",
        "description": "Ҳаҷм дар литр",
        "minimum": 100,
        "maximum": 1000
      },
      "ранг": {
        "type": "string",
        "description": "Ранги яхдон",
        "enum": ["Сафед", "Нуқра", "Сиёҳ"]
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: Чангкашакҳо
    ('Чангкашакҳо', 26, 'tj', '{
    "type": "object",
    "title": "Чангкашакҳо",
    "description": "Хусусиятҳо барои чангкашакҳо бо истифодаи oneOf",
    "oneOf": [
      {
        "properties": {
          "навъ": {
            "type": "string",
            "const": "Робот"
          },
          "батарея": {
            "type": "integer",
            "description": "Иқтидори батарея дар мАч",
            "minimum": 2000
          }
        },
        "required": ["навъ", "батарея"]
      },
      {
        "properties": {
          "навъ": {
            "type": "string",
            "const": "Классик"
          },
          "қувват": {
            "type": "integer",
            "description": "Қувват дар ватт",
            "minimum": 500
          }
        },
        "required": ["навъ", "қувват"]
      }
    ]
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: Лавозимоти электроника
    ('Лавозимоти электроника', 21, 'tj', '{
    "type": "object",
    "title": "Лавозимоти электроника",
    "description": "Хусусиятҳо барои лавозимот",
    "required": ["навъи_лавозимот"],
    "properties": {
      "навъи_лавозимот": {
        "type": "string",
        "description": "Навъи лавозимот",
        "enum": ["Ғилоф", "Заряддиҳанда", "Гӯшмонакҳо"]
      }
    }
  }', NOW(), true, NOW(), NULL),
    -- Подкатегория: Ҷинсҳо
    ('Ҷинсҳо', 24, 'tj', '{
    "type": "object",
    "title": "Ҷинсҳо",
    "description": "Хусусиятҳо барои ҷинсҳо",
    "required": ["андоза", "ранг"],
    "properties": {
      "андоза": {
        "type": "string",
        "description": "Андозаи ҷинс",
        "enum": ["28", "30", "32", "34"]
      },
      "ранг": {
        "type": "string",
        "description": "Ранги ҷинс",
        "enum": ["Кабуд", "Сиёҳ", "Хокистарӣ"]
      }
    }
  }', NOW(), true, NOW(), NULL);