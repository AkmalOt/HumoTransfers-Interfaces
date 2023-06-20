package Agent

//	type MockAgentRepository interface {
//		AddAgent(agents *models.Agents) error
//		GetAgent(pagination *models.Pagination) ([]models.Agents, error)
//		UpdateAgents(agent *models.Agents) error
//		DeleteAgents(agent *models.Agents) error
//		AgentStatus(agent *models.Agents) error
//		TotalPageAgents(limit int64) (int64, error)
////	}
//var MockAgentRepository = AgentInterface
//
//func (r *MockAgentRepository) AddAgent(agents *models.Agents) error {
//	// Возвращаем nil, чтобы имитировать успешное добавление агента
//	return nil
//}
//
//func TestAddAgent(t *testing.T) {
//	agent := models.Agents{
//		Name:      "Test",
//		LegalName: "Testovich",
//		Active:    true,
//	}
//
//	//MockAgentRepository := AgentInterface{}
//	//paginate := Paginate
//	handler := Agent{
//		repo: MockAgentRepository,
//		//pagination: Paginate,
//	}
//	// Создаем экземпляр тестового роутера Gin
//	router := gin.Default()
//	router.POST("/test", handler.AddAgent)
//
//	// Создаем тестовый контекст Gin и привязываем JSON-данные к телу запроса
//	w := httptest.NewRecorder()
//
//	req, _ := http.NewRequest(http.MethodPost, "/get_agents", toJSON(agent))
//	//req.Header.Set("Content-Type", "application/json")
//	router.ServeHTTP(w, req)
//
//	// Проверяем код состояния ответа
//	if w.Code != http.StatusOK {
//		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
//	}
//
//	// Проверяем тело ответа
//	expectedResponse := "Agent added!"
//	if w.Body.String() != expectedResponse {
//		t.Errorf("Expected response body '%s', but got '%s'", expectedResponse, w.Body.String())
//	}
//
//}
//
//// Функция для сериализации структуры в JSON
//func toJSON(data interface{}) *bytes.Buffer {
//	buf := new(bytes.Buffer)
//	json.NewEncoder(buf).Encode(data)
//	return buf
//}
