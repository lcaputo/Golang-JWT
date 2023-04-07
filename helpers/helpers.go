package helpers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"reflect"
	"strings"
)

type InvalidTypeResponse struct {
	Field    string `json:"field"`
	Expected string `json:"expected"`
	Got      string `json:"got"`
}

type SyntaxErrorResponse struct {
	Error string `json:"error"`
}

func PrintJSONTag(ptr interface{}, fieldName string) string {
	// Obtener el valor de la estructura a través del puntero
	value := reflect.ValueOf(ptr).Elem()
	// Obtener el tipo de la estructura
	typ := value.Type()
	// Buscar el índice del campo con el nombre especificado
	fieldIndex := -1
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Name == fieldName {
			fieldIndex = i
			break
		}
	}
	// Si no se encontró el campo, imprimir un mensaje de error y salir
	if fieldIndex == -1 {
		return ""
	}
	// Obtener el nombre del campo y su etiqueta JSON correspondiente
	field := typ.Field(fieldIndex)
	jsonTag := field.Tag.Get("json")

	// Imprimir el resultado
	return jsonTag
}

func ParseKeyValueString(str string) string {
	// Separamos la cadena en palabras utilizando los separadores "&" y "="
	splitByComma := strings.Split(str, ",")
	arr := []string{}
	for _, elm := range splitByComma {
		splitByColon := strings.Split(elm, ":")
		arr = append(arr, splitByColon...)
	}
	for index, elm := range arr {
		arr[index] = strings.TrimSpace(elm)
	}
	// Objeto para guardar el resultado final
	result := make(map[string]string)
	// Iterar sobre cada elemento del arreglo
	for _, item := range arr {
		// Separar la cadena por el símbolo "="
		parts := strings.Split(item, "=")

		// Validar que tenga un par llave-valor válido
		if len(parts) != 2 {
			continue
		}
		// Asignar la llave y el valor al objeto resultado
		key := parts[0]
		value := parts[1]
		result[key] = value
	}
	// Convertir el objeto resultado a formato JSON
	jsonData, err := json.Marshal(result)
	if err != nil {
		return ""
	}
	// Imprimir el resultado final en formato JSON
	return string(jsonData)
}

func BindAndValidate(c echo.Context, u interface{}) any {
	if err := c.Bind(u); err != nil {
		errors := ParseKeyValueString(err.Error())
		syntaxErrorResponse := SyntaxErrorResponse{}
		if err := json.Unmarshal([]byte(errors), &syntaxErrorResponse); err != nil {
			return err
		}
		if syntaxErrorResponse.Error != "" {
			return &syntaxErrorResponse
		}
		invalidTypeResponse := InvalidTypeResponse{}
		if err := json.Unmarshal([]byte(errors), &invalidTypeResponse); err != nil {
			return err
		}
		if invalidTypeResponse.Field != "" || invalidTypeResponse.Got != "" || invalidTypeResponse.Expected != "" {
			return &invalidTypeResponse
		}
	}
	if err := c.Validate(u); err != nil {
		return err
	}
	return nil
}
