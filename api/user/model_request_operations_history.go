/*
 * Информация о пользователе
 *
 * API для получения баланса кошелька пользователя, истории операций и инфорации об авторизованном пользователе **Авито API для бизнеса предоставляется согласно [Условиям использования](https://api.avito.ru/docs/public/APITermsOfServiceV1.pdf).**
 *
 * API version: 1
 * Contact: supportautoload@avito.ru
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package user

import (
	"time"
)

type RequestOperationsHistory struct {
	// Время выборки от. Не далее одного года в прошлое от текущего момента.
	DateTimeFrom time.Time `json:"dateTimeFrom"`
	// Время выборки до (диапазон между dateTimeFrom и dateTimeTo должен быть не более одной недели)
	DateTimeTo time.Time `json:"dateTimeTo"`
}
