package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPublishHandler(t *testing.T) {
	// 1. Подготовка данных для теста
	// Создаем фейковую конфигурацию
	testCfg := &Config{
		MQTTTopic: "test/topic",
	}
	// Сообщение, которое мы хотим отправить
	testMessage := "hello from test"

	// 2. Создание "фейкового" запроса
	// Запрос теперь GET на /publish с параметром ?msg=...
	req, err := http.NewRequest("GET", "/publish?msg="+testMessage, nil)
	if err != nil {
		t.Fatal(err)
	}

	// 3. Создание "фейкового" сборщика ответов
	rr := httptest.NewRecorder()

	// 4. Вызов тестируемого обработчика
	// Вызываем publishHandler, передавая nil вместо реального MQTT клиента.
	// В `publishHandler` мы добавили проверку `if client != nil`, чтобы тест не падал.
	handler := publishHandler(nil, testCfg)
	handler.ServeHTTP(rr, req)

	// 5. Проверка результата
	// Проверяем, что HTTP-статус код ответа - 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Проверяем тело ответа
	// Ожидаем, что в ответе будет строка "Published message: hello from test"
	expectedBody := "Published message: " + testMessage
	if !strings.Contains(rr.Body.String(), expectedBody) {
		t.Errorf("handler returned unexpected body: got %v want to contain %q",
			rr.Body.String(), expectedBody)
	}
}
