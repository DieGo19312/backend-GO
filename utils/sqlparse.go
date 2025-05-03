package utils

//el sqlparse solo se ocupa para generar consultas sql de tipo insert y update
import (
	"fmt"
	"strconv"
	"strings"
)

// ParseInsertArray genera dinámicamente una consulta SQL de tipo INSERT
func ParseInsertArray(tableName string, data map[string]interface{}) (string, []interface{}) {
	columns := []string{}
	placeholders := []string{}
	values := []interface{}{}

	i := 1
	for key, value := range data {
		columns = append(columns, key)                             // Agrega el nombre de la columna
		placeholders = append(placeholders, fmt.Sprintf("$%d", i)) // Genera los placeholders ($1, $2, ...)
		values = append(values, value)                             // Agrega el valor correspondiente
		i++
	}

	// Genera la consulta SQL
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	return query, values

	/*ejemplo de uso:
			data := map[string]interface{}{
			"name":  user.Name,
			"email": "john@example.com",
			"age":   25,
	}

	query, values := utils.ParseInsertArray("users", data)
	fmt.Println(query)  // INSERT INTO users (name, email, age) VALUES ($1, $2, $3)
	fmt.Println(values) // [John john@example.com 25]
	*/

}
//Genera dinámicamente una consulta SQL de tipo UPDATE
// ParseUpdateArray genera una consulta SQL de actualización y una lista de valores
func ParseUpdateArray(table string, data map[string]interface{}, condition map[string]interface{}) (string, []interface{}) {
	setClauses := []string{}
	values := []interface{}{}
	i := 1

	// Construir las cláusulas SET
	for key, value := range data {
			if value != nil && value != "" { // Filtrar valores vacíos o nulos
					setClauses = append(setClauses, key+" = $"+strconv.Itoa(i))
					values = append(values, value)
					i++
			}
	}

	// Construir las condiciones WHERE
	whereClauses := []string{}
	for key, value := range condition {
			whereClauses = append(whereClauses, key+" = $"+strconv.Itoa(i))
			values = append(values, value)
			i++
	}

	// Generar la consulta SQL
	query := "UPDATE " + table + " SET " + strings.Join(setClauses, ", ") + " WHERE " + strings.Join(whereClauses, " AND ")
	return query, values
}
	/* Ejemplo de uso:
			data := map[string]interface{}{
    "name": "John Doe",
    "age":  26,
}
condition := map[string]interface{}{
    "id": 1,
}
query, values := utils.ParseUpdateArray("users", data, condition)
fmt.Println(query)  // UPDATE users SET name = $1, age = $2 WHERE id = $3
fmt.Println(values) // [John Doe 26 1]

	*/

